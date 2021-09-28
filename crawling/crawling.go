package crawling

type CrawlerInterface interface {
	New() *struct{}
	Crawl(string, int)
}
