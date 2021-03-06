/*
	This package contains all necessary functions for crawling Ebay website with given keyword and page.

	How-To:
	eb := ebay.New("<KEYWOOD>")
	eb.Crawl()
*/
package ebay

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"

	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

type Ebay struct {
	c       *colly.Collector
	keyword string
	results []productPB.ProductResponse
	stream  productPB.ProductService_QueryServer
}

// New creates a new Ebay instance.
func New(stream productPB.ProductService_QueryServer, kw string) *Ebay {
	c := colly.NewCollector(colly.AllowedDomains("www.ebay.com"))
	return &Ebay{
		c:       c,
		keyword: kw,
		stream:  stream,
	}
}

// Crawl performs crawling operations.
func (e *Ebay) Crawl() []productPB.ProductResponse {

	e.c.OnHTML(".s-item", func(soup *colly.HTMLElement) {
		if err := e.extractContent(soup); err != nil {
			log.Println("Failed to extract ebay product content: ", err)
		}
	})

	e.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for page := 1; page <= 5; page++ {
		url := fmt.Sprintf("https://www.ebay.com/sch/i.html?_nkw=%v&_pgn=%v", html.EscapeString(strings.Replace(e.keyword, " ", "+", -1)), page)
		if err := e.c.Visit(url); err != nil {
			log.Fatalf("Failed to start scraping url %q: %v", url, err)
		}
	}

	return e.results
}

// extractContent extract product information from HTML contents.
func (e *Ebay) extractContent(soup *colly.HTMLElement) error {
	id := e.parseID(soup.ChildAttr(".s-item__link", "href"))

	name := soup.ChildText(".s-item__title")

	price, err := e.parsePrice(soup.ChildText(".s-item__price"))
	if err != nil {
		return errors.Wrap(err, "failed to parse product price")
	}

	productUrl := e.parseProductURL(soup.ChildAttr(".s-item__link", "href"))

	imageUrl := soup.ChildAttr(".s-item__image-img", "src")

	res := productPB.ProductResponse{
		Platform:   "Ebay",
		Id:         id,
		Name:       name,
		Price:      price,
		ProductUrl: productUrl,
		ImageUrl:   imageUrl,
	}

	e.results = append(e.results, res)
	e.stream.Send(&res)

	return nil
}

func (e *Ebay) parseID(url string) string {
	urlClean := strings.Split(url, "?")[0]
	urlPieces := strings.Split(urlClean, "/")
	productId := urlPieces[len(urlPieces)-1]
	return productId
}

// parsePrice parses and return correct price format.
func (e *Ebay) parsePrice(priceOriginal string) (float32, error) {
	if strings.Replace(priceOriginal, " ", "", -1) == "" {
		return -1, errors.New("get empty price")
	}
	priceWithoutRange := strings.Split(priceOriginal, "NT$")[1]
	priceWithoutGap := strings.Fields(priceWithoutRange)[0]
	priceWithoutComma := strings.Replace(priceWithoutGap, ",", "", -1)
	priceWithoutSpace := strings.TrimSpace(priceWithoutComma)
	priceFinal, err := strconv.ParseFloat(priceWithoutSpace, 32)
	if err != nil {
		return -1, err
	}

	return float32(priceFinal), nil
}

func (e *Ebay) parseProductURL(url string) string {
	return strings.Split(url, "?")[0]
}
