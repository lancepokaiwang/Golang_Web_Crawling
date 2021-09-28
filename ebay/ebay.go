/*
	This package contains all necessary functions for crawling Ebay website with given keyword and page.

	How-To:
	eb := ebay.New()
	eb.Crawl("<KEYWOOD>", <PAGE_NUM>)
*/
package ebay

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

var c = colly.NewCollector(
	colly.AllowedDomains("www.ebay.com"),
)

type Ebay struct{}

// New creates a new Ebay instance.
func New() *Ebay {
	return &Ebay{}
}

// Crawl performs crawling operations.
func (e *Ebay) Crawl(keyword string, page int) {
	url := fmt.Sprintf("https://www.ebay.com/sch/i.html?_nkw=%v&_pgn=%v", html.EscapeString(strings.Replace(keyword, " ", "+", -1)), page)

	c.OnHTML(".s-item", func(soup *colly.HTMLElement) {
		if err := e.extractContent(soup); err != nil {
			log.Println("Failed to extract ebay product content: ", err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	if err := c.Visit(url); err != nil {
		log.Fatalf("Failed to start scraping url %q: %v", url, err)
	}

}

// extractContent extract product information from HTML contents.
func (e *Ebay) extractContent(soup *colly.HTMLElement) error {
	id := e.parseID(soup.ChildAttr(".s-item__link", "href"))
	fmt.Println(id)

	name := soup.ChildText(".s-item__title")
	fmt.Println(name)

	price, err := e.parsePrice(soup.ChildText(".s-item__price"))
	if err != nil {
		return errors.Wrap(err, "failed to parse product price")
	}
	fmt.Println(price)

	productUrl := e.parseProductURL(soup.ChildAttr(".s-item__link", "href"))
	fmt.Println(productUrl)

	imageUrl := soup.ChildAttr(".s-item__image-img", "src")
	fmt.Println(imageUrl)

	fmt.Println("--------------------------------")

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
	priceWithoutComma := strings.Replace(priceWithoutRange, ",", "", -1)
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
