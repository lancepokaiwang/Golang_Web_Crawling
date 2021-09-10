package main

import (
	"context"
	"io"
	"log"

	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := productPB.NewProductServiceClient(conn)

	sendRequest(client)
}

func sendRequest(client productPB.ProductServiceClient) {
	log.Print("Staring to do a Unary RPC")

	keyword := "iPhone_" // Let's say we want to query a product called `iphone_`

	req := &productPB.ProductRequest{
		Query: keyword,
	}

	stream, err := client.Query(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to sent query for keyword %q: %v \n", keyword, err)
	}

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Query(_) = _, %v", client, err)
		}
		log.Println(data)
	}
}
