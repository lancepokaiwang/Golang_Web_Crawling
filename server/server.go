package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
	"google.golang.org/grpc"
)

type Server struct{ productPB.ProductServiceServer }

func (*Server) Query(req *productPB.ProductRequest, stream productPB.ProductService_QueryServer) error {
	log.Printf("Query function is invoked with %v \n", req)

	keyword := req.GetQuery()

	for i := 0; i < 10; i++ {
		// Query keyword here...
		price, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 1.99+float32(i)), 32)

		res := &productPB.ProductResponse{
			Id:         "asd1234" + strconv.Itoa(i+1),
			Name:       keyword + strconv.Itoa(i+1),
			Price:      float32(price),
			ProductUrl: "https://amazon.com/" + keyword + strconv.Itoa(i+1),
			ImageUrl:   "https://image.amazon.com/" + keyword + strconv.Itoa(i+1),
		}

		if err := stream.Send(res); err != nil {
			log.Fatal("Failed to start streaming")
		}

		time.Sleep(time.Second)
	}

	return nil
}

func (*Server) SayHello(ctx context.Context, req *productPB.HelloRequest) (*productPB.HelloReply, error) {
	name := req.GetName()
	s.ContextLog("Got a request, try to say hello")
	res := &productPB.HelloReply{
		Message: "hello, " + name,
	}

	return res, nil
}

func New() {
	s.ContextLog("Starting gRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("Failed to create gRPC service: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	productPB.RegisterProductServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v \n", err)
	}
}
