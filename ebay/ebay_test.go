/*
	This package contains all necessary functions for crawling Ebay website with given keyword and page.

	How-To:
	eb := ebay.New("<KEYWOOD>")
	eb.Crawl()
*/
package ebay

import (
	"reflect"
	"testing"

	"github.com/gocolly/colly"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

func TestNew(t *testing.T) {
	type args struct {
		stream productPB.ProductService_QueryServer
		kw     string
	}
	tests := []struct {
		name string
		args args
		want *Ebay
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.stream, tt.args.kw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEbay_Crawl(t *testing.T) {
	tests := []struct {
		name string
		e    *Ebay
		want []productPB.ProductResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Crawl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ebay.Crawl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEbay_extractContent(t *testing.T) {
	type args struct {
		soup *colly.HTMLElement
	}
	tests := []struct {
		name    string
		e       *Ebay
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.extractContent(tt.args.soup); (err != nil) != tt.wantErr {
				t.Errorf("Ebay.extractContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEbay_parseID(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		e    *Ebay
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.parseID(tt.args.url); got != tt.want {
				t.Errorf("Ebay.parseID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEbay_parsePrice(t *testing.T) {
	type args struct {
		priceOriginal string
	}
	tests := []struct {
		name    string
		e       *Ebay
		args    args
		want    float32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.parsePrice(tt.args.priceOriginal)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ebay.parsePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ebay.parsePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEbay_parseProductURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		e    *Ebay
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.parseProductURL(tt.args.url); got != tt.want {
				t.Errorf("Ebay.parseProductURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
