package server

import (
	"context"
	"fmt"
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
		Stream:  stream,
	}

	ebay := &crawling.CrawlClient{
		Keyword: keyword,
		Web:     crawling.TypeEbay,
		Stream:  stream,
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	resultCh := make(chan []productPB.ProductResponse, 2)

	go wp.QueueJob([]*crawling.CrawlClient{amazon, ebay}, wg, resultCh)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var combinedResult []productPB.ProductResponse

	for results := range resultCh {
		combinedResult = append(combinedResult, results...)
		fmt.Printf("result len: %d\n", len(results))
	}

	ps := redis.ProductSlice{Products: combinedResult}
	if err := redisClient.Insert(keyword, ps); err != nil {
		s.Fatalf("Failed to insert %s into redis: %v", keyword, err)
	}

	return nil
}

const (
	workersCount = 5
	jobsCount    = 10
)

func New() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize worker pool.
	wp = workers.New(workersCount, jobsCount)
	defer wp.Close()
	// Activate workers to listening for jobs channel.
	go wp.Run(ctx)

	s.Println("Starting gRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		s.Fatalf("Failed to create gRPC service: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	productPB.RegisterProductServiceServer(grpcServer, &Server{})

	// Graceful shutdown.
	sigCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		cancel()
		grpcServer.GracefulStop()
		done <- true
	}()

	fmt.Println("grpc server is listening...")
	if err := grpcServer.Serve(lis); err != nil {
		s.Fatalf("Failed to serve: %v \n", err)
	}

	<-done
	s.Println("Graceful shutdown")
}

func (*Server) SayHello(ctx context.Context, req *productPB.HelloRequest) (*productPB.HelloReply, error) {
	name := req.GetName()
	s.Println("Got a request, try to say hello")
	res := &productPB.HelloReply{
		Message: "hello, " + name,
	}

	return res, nil
}
