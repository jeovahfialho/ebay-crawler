
# eBay Crawler Project

## Description
This project is an eBay crawler developed in Go. It is designed to visit a specified eBay seller's page, extract information about listed items including title, price, URL, and condition, and save these details as JSON files in a specified directory. The crawler supports both single-page and multi-page (pagination) crawling modes.

## Installation

To run this project, you need to have Go installed on your system. If you don't have Go installed, you can download it from https://golang.org/dl/.

After installing Go, clone the project repository to your local machine.

## Running the Project

1. Navigate to the project directory in your terminal.
2. Build the project using the `go build` command.
3. Run the executable. For example, if your executable is named `ebayCrawler`, you would run `./ebayCrawler`.

### Enabling Pagination

The crawler can operate in two modes: single-page mode and pagination mode. By default, the crawler operates in single-page mode, processing only the items found on the initial page specified by the `baseURL`.

To enable pagination, set the `enablePagination` variable to `true` in the `main.go` file before running the crawler. This instructs the crawler to follow pagination links and process items across multiple pages, up to a predefined limit.

```go
enablePagination := true // Enable pagination
```

## Dependencies

This project uses the following Go packages:
- `goquery` for parsing HTML documents.
- `net/http` for making HTTP requests.

You can install any necessary dependencies using `go get`. For example:
```
go get github.com/PuerkitoBio/goquery
```

## Structure

- `cmd/main.go`: Entry point of the program.
- `internal/crawler/crawler.go`: Contains the logic to crawl eBay pages, including pagination.
- `internal/parser/parser.go`: Parses HTML pages and extracts product information.
- `internal/model/product.go`: Defines the product structure.
- `internal/writer/writer.go`: Manages writing product data to JSON files.

## Testing

To run tests, use the `go test ./...` command in the directory of the package you wish to test.
