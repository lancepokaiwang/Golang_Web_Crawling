package main

import ebay "github.com/lancepokaiwang/Golang_Web_Crawling/ebay"

// This is where the whole application start.
// http://localhost:8080/products
// func main() {
// 	s.ContextLog("Starting application")
// 	server.New()
// }

func main() {
	a := ebay.New()

	a.Crawl("bike", 1)
}
