# Golang Web Crawling Project

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
- [ ] 1. The model need to be able to query and search at least 2 online shopping platforms.
- [ ] 2. Unit test is required.
- [ ] 3. Using ***Worker package (multi-threading in Golang)*** to accelerate workflow. But there must be a maximum limit, though.
- [ ] 4. Multi-page results must be supported.
- [ ] 5. Colud use ***interface*** feature to make code more flexible and extendable.
- [ ] 6. When program is interrupted, ***worker*** can not be stopped until it complete its task.

### Optional (Plus)
- [ ] 1. Construct a proper front-end web application.
- [ ] 2. Support ***streaming*** feature to return query result asynchronously as it might take long time to complete the whole searching process.
- [ ] 3. Support database to support simulate cache.
- [ ] 4. Construct a proper ***log*** system for better error tracking and debuging.

## Libraries Used
### Web crawling:
- [colly](https://github.com/gocolly/colly): open source web crawling framework for golang.

### Web crawling:
- [RESTful in Golang](https://golang.org/doc/tutorial/web-service-gin): tutorial: Developing a RESTful API with Go and Gin

## Notes
1. Using existing gframework or libraries is a good approach.
2. Be aware of NOT sending query too frequent.
3. Query result threshold can be specified in advance. Out-of-stock merchandise can be ignored as well.

## Something Useful:
- [How to install Go on Mac via brew](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)
