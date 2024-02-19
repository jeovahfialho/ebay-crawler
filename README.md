
# eBay Crawler Project

## Description
This project is an eBay crawler developed in Go. It is designed to visit a specified eBay seller's page, extract information about listed items including title, price, URL, and condition, and save these details as JSON files in a specified directory.

## Installation

To run this project, you need to have Go installed on your system. If you don't have Go installed, you can download it from https://golang.org/dl/.

After installing Go, clone the project repository to your local machine.

## Running the Project

1. Navigate to the project directory in your terminal.
2. Build the project using the `go build` command.
3. Run the executable. For example, if your executable is named `ebayCrawler`, you would run `./ebayCrawler`.

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
- `internal/crawler/crawler.go`: Contains the logic for crawling web pages.
- `internal/parser/parser.go`: Parses HTML pages and extracts product information.
- `internal/model/product.go`: Defines the product structure.
- `internal/writer/writer.go`: Manages writing product data to JSON files.

## Testing

To run tests, use the `go test` command in the directory of the package you wish to test.

