package crawling

import (
	amazon "github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	ebay "github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

type CrawlingProvider struct{}

func (cp CrawlingProvider) Process(keyword string) []productPB.ProductResponse {
	// Crawling Ebay.
	e := ebay.New(keyword)
	ebayResults := e.Crawl()

	// Crawling Amazon.
	a := amazon.New(keyword)
	amazonResults := a.Crawl()

	// Return combined results.
	return append(ebayResults, amazonResults...)
}

type Logic interface {
	Process(keyword string) []productPB.ProductResponse
}

type CrawlClient struct {
	L Logic
}

// PerformCrawling prodives an entry point for clients who want to perform crawling for both platforms.
func (cc CrawlClient) PerformCrawling(keyword string) []productPB.ProductResponse {
	return cc.L.Process(keyword)
}

/*
In main or any function where you want to perform crawling for both platforms:

func main() {
	cc := CrawlClient{
		L: CrawlingProvider{},
	}

	data := cc.PerformCrawling(<KEYWORD>)

	// Use `data` for further actions.
}
*/
