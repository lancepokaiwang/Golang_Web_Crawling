package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lancepokaiwang/Golang_Web_Crawling/crawling"
	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
	"github.com/lancepokaiwang/Golang_Web_Crawling/redis"
	"github.com/lancepokaiwang/Golang_Web_Crawling/workers"

	"google.golang.org/grpc"
)

const (
	workersCount = 5
)

var wp *workers.WorkerPool

type Server struct{}

func (*Server) Query(req *productPB.ProductRequest, stream productPB.ProductService_QueryServer) error {
	s.Printf("Query function is invoked with %v \n", req)

	keyword := req.GetQuery()

	redisClient := redis.NewClient()
	products, err := redisClient.Query(keyword)
	if err != nil {
		s.Fatalf("Failed to query redis: %v", err)
	}

	if products != nil {
		for _, product := range products {
			if err := stream.Send(&product); err != nil {
				s.Fatalf("Failed to send response by stream: %v", err)
			}
		}

		return nil
	}

	amazon := &crawling.CrawlClient{
		Keyword: keyword,
		Web:     crawling.TypeAmazon,
	}

	ebay := &crawling.CrawlClient{
		Keyword: keyword,
		Web:     crawling.TypeEbay,
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	resultChan := make(chan []productPB.ProductResponse, 2)

	go wp.NewJob([]*crawling.CrawlClient{amazon, ebay}, wg, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var combinedResult []productPB.ProductResponse

	for results := range resultChan {
		combinedResult = append(combinedResult, results...)
		fmt.Printf("result len: %d\n", len(results))
	}

	ps := redis.ProductSlice{Products: combinedResult}
	if err := redisClient.Insert(keyword, ps); err != nil {
		s.Fatalf("Failed to insert %s into redis: %v", keyword, err)
	}

	return nil
}

func New() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init worker pool.
	wp = workers.New(workersCount)
	// Activate workers to listening for jobs channel.
	go wp.Run(ctx)

	s.Println("Starting gRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("Failed to create gRPC service: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	productPB.RegisterProductServiceServer(grpcServer, &Server{})

	// Graceful shutdown.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		cancel()
		grpcServer.GracefulStop()
		wg.Done()
	}()

	fmt.Println("grpc server is listening...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v \n", err)
	}

	wg.Wait()
	log.Println("Graceful shutdown")
}

func (*Server) SayHello(ctx context.Context, req *productPB.HelloRequest) (*productPB.HelloReply, error) {
	name := req.GetName()
	s.Println("Got a request, try to say hello")
	res := &productPB.HelloReply{
		Message: "hello, " + name,
	}

	return res, nil
}
