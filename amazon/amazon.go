package amazon

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"

	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"

	"github.com/gocolly/colly"
)

type Amazon struct {
	c       *colly.Collector
	keyword string
	results []productPB.ProductResponse
}

// New creates a new Amazon instance.
func New(kw string) *Amazon {
	c := colly.NewCollector(colly.AllowedDomains("www.amazon.com", "amazon.com"))
	
	return &Amazon{
		c:       c,
		keyword: kw,
	}
}

func (a *Amazon) Crawl() []productPB.ProductResponse {

	a.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/93.0.961.44")
	})

	s := []string{""}
	a.c.OnHTML("div.sg-col-inner", func(h *colly.HTMLElement) {
		title := h.ChildText("span.a-size-medium")
		price := h.ChildText("span.a-price-whole") + "." + h.ChildText("span.a-price-fraction")
		image := h.ChildAttr("img.s-image", "src")
		link := h.ChildAttr("a.a-link-normal", "href")
		priceNum, err := strconv.ParseFloat(price, 32)
		if err != nil{
			log.Println("Failed to convert price to float32: ", err)
		}
		if !(s[len(s)-1] == title) && title != "" {
			s = append(s, title)
			if len(s) > 2 {
				res := productPB.ProductResponse{
					Platform:   "Amazon",
					Name:       title,
					Price:      float32(priceNum),
					ProductUrl: link,
					ImageUrl:   image,
				}
				a.results = append(a.results, res)
			}
		}
	})

	for i:=1; i<=5; i++{
		url := fmt.Sprintf("https://www.amazon.com/s?k=%v&page=%v&language=zh_TW&currency=TWD", html.EscapeString(strings.Replace(a.keyword, " ", "+", -1)), i)
		if err := a.c.Visit(url); err != nil {
			log.Fatalf("Failed to start scraping url %q: %v", url, err)
		}
	}

	return a.results
}
