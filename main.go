package main

import (
	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	"github.com/lancepokaiwang/Golang_Web_Crawling/server"
)

// This is where the whole application start.
// http://localhost:8080/products
func main() {
	s.Println("Starting application")
	server.New()
}

// func main() {
// a := ebay.New("Apple")
// t := a.Crawl()
// fmt.Println(t)
// }
