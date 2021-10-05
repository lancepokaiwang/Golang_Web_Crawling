package crawling

import (
	"github.com/lancepokaiwang/Golang_Web_Crawling/amazon"
	"github.com/lancepokaiwang/Golang_Web_Crawling/ebay"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

type CrawlClient struct {
	Keyword string
	Web     WebsiteType
	Stream  productPB.ProductService_QueryServer
}

// PerformCrawling prodives an entry point for clients who want to perform crawling for both platforms.
func (cc CrawlClient) PerformCrawling() []productPB.ProductResponse {
	switch cc.Web {
	case TypeAmazon:
		a := amazon.New(keyword)
		return a.Crawl()
	case TypeEbay:
		e := ebay.New(cc.Stream, cc.Keyword)
		return e.Crawl()
	default:
		return nil
	}
}

type WebsiteType int

const (
	TypeAmazon WebsiteType = iota
	TypeEbay
)

/*
In main or any function where you want to perform crawling for both platforms:

func main() {
	cc := CrawlClient{
		Keyword: keyword,
		Web:     crawling.TypeAmazon,
	}

	data := cc.PerformCrawling(cc.Keyword, cc.Web)
	if data == nil{
		// Error handling here.
	}

	// Use `data` for further actions.
}
*/
