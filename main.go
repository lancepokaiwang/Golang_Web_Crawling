package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	"github.com/lancepokaiwang/Golang_Web_Crawling/redis"
)

// This is where the whole application start.
// http://localhost:4000
func main() {
	// s.ContextLog("Starting application")
	// server.New()

	keyword := "bike"

	a := ebay.New(keyword)

	results := a.Crawl()

	r := redis.NewClient()
	if err := r.Insert(keyword, results); err != nil {
		log.Fatalf("Failed to insert result of keyword %q: %v", keyword, err)
	}

	queryResult, err := r.Query(keyword)
	if err != nil {
		log.Fatalf("Failed to query keyword %q: %v", err, err)
	}

	spew.Dump(queryResult)
}
