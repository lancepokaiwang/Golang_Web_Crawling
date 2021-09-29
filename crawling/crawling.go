package crawling

import (
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

type CrawlerInterface interface {
	Crawl(keyword string, page int) productPB.ProductResponse
}
