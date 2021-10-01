package crawling

import (
	amazon "github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	ebay "github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

type CrawlingProvider struct{}

func (cp CrawlingProvider) Process(keyword string, web websiteType) []productPB.ProductResponse {
	switch web {
	case TypeAmazon:
		a := amazon.New(keyword)
		return a.Crawl()
	case TypeEbay:
		e := ebay.New(keyword)
		return e.Crawl()
	default:
		return nil
	}
}

type Logic interface {
	Process(keyword string, web websiteType) []productPB.ProductResponse
}

type CrawlClient struct {
	L Logic
}

// PerformCrawling prodives an entry point for clients who want to perform crawling for both platforms.
func (cc CrawlClient) PerformCrawling(keyword string, web websiteType) []productPB.ProductResponse {
	return cc.L.Process(keyword, web)
}

type websiteType int

const (
	TypeAmazon websiteType = iota
	TypeEbay
)

/*
In main or any function where you want to perform crawling for both platforms:

func main() {
	cc := CrawlClient{
		L: CrawlingProvider{},
	}

	data := cc.PerformCrawling(<KEYWORD>, <WebsiteType>)
	if data == nil{
		// Error handling here.
	}

	// Use `data` for further actions.
}
*/
