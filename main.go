package main

import (
	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	api "github.com/lancepokaiwang/Golang_Web_Crawling/rest_api"
)

// This is where the whole application start.
// http://localhost:8080/products
func main() {
	s.ContextLog("Starting REST server")
	if err := api.NewConnection(); err != nil {
		s.Fatal(err.Error())
	}
	// run this to get pre-defined objects:
	/*
		curl http://localhost:8080/products
	*/

	// try to run this:
	/*
		curl http://localhost:8080/products \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": "4", "name": "D", "price": 4.99, "rating": 3.5, "url": "https://carleton.ca"}'
	*/
}
