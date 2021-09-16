# Golang Web Crawling Project
This project aims at searching products across the biggest two obline shopping platforms (i.e., [Amazon](https://www.amazon.com/) and [Walmart](https://www.walmart.com/)) and find out the one with the lowest price or highest user rating by using Golang and gRPC service.

## Requirement
Construct a product price comparison service that enables customers to search and compare prices from two online shopping platforms simultaneously.

```
service YourService {
  rpc Search (Request) returns (stream Result) {}
}

// Where `Request` is search keyword and `Result` contins merchandise information (e.g., name, price, image, link, etc.) 
```

## Features
### Basic
- [ ] 1. The model need to be able to query and search at least 2 online shopping platforms **(Amazon and Walmart)**.
- [ ] 2. Unit test is required.
- [ ] 3. Using ***Worker package (multi-threading in Golang)*** to accelerate workflow. But there must be a maximum limit, though. (Use third-party tools)
- [ ] 4. Multi-page results must be supported.
- [ ] 5. Colud use ***interface*** feature to make code more flexible and extendable.
- [ ] 6. When program is interrupted, ***worker*** can not be stopped until it complete its task.

### Optional (Plus)
- [ ] 1. Construct a proper front-end web application.
- [ ] 2. Support ***streaming*** feature to return query result asynchronously as it might take long time to complete the whole searching process.
- [ ] 3. Support database to support simulate cache.
- [ ] 4. Construct a proper ***log*** system for better error tracking and debuging.

## First Time Use?
1. Make sure Golang is installed on your machine.
2. Clone the project under the proper path (outside of your $GOPATH)
3. Make sure GO111MODULE is "" or "on".
4. Under the project folder, use terminal command to download required packages:
``` 
// This command will update all required packages. 
$ go mod tidy 
```

## How to Run?
```
// Start gRPC server.
$ go run main.go
```
```
// Let's make a fake query.
$ go run client/client.go

// Outputs will look like:
// 2021/09/10 17:16:57 id:"asd1234"  name:"iPhone_1"  price:1.99  rating:4.5  product_url:"https://amazon.com/iPhone_1"  image_url:"https://image.amazon.com/iPhone_1"
// 2021/09/10 17:16:58 id:"asd1234"  name:"iPhone_2"  price:2.99  rating:4.5  product_url:"https://amazon.com/iPhone_2"  image_url:"https://image.amazon.com/iPhone_2"
// 2021/09/10 17:16:59 id:"asd1234"  name:"iPhone_3"  price:3.99  rating:4.5  product_url:"https://amazon.com/iPhone_3"  image_url:"https://image.amazon.com/iPhone_3"
......
// 2021/09/10 17:17:05 id:"asd1234"  name:"iPhone_9"  price:9.99  rating:4.5  product_url:"https://amazon.com/iPhone_9"  image_url:"https://image.amazon.com/iPhone_9"
// 2021/09/10 17:17:06 id:"asd1234"  name:"iPhone_10"  price:10.99  rating:4.5  product_url:"https://amazon.com/iPhone_10"  image_url:"https://image.amazon.com/iPhone_10"
```

## Libraries Used
### Web crawling:
- [colly](https://github.com/gocolly/colly): open source web crawling framework for golang.

### API:
- [gRPC Getting Started](https://pjchender.dev/golang/grpc-getting-started/)
- [gRPC Quick start from Google](https://grpc.io/docs/languages/go/quickstart/)
- [gRPC-Web](https://github.com/grpc/grpc-web)

## Notes
1. Using existing gframework or libraries is a good approach.
2. Be aware of NOT sending query too frequent.
3. Query result threshold can be specified in advance. Out-of-stock merchandise can be ignored as well.

## Something Useful:
- [How to install Go on Mac via brew](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)
