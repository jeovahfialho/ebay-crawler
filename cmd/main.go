// Package main is the entry point for the eBay crawler application.
// This application is designed to crawl a specified eBay seller's page,
// extract information about listed items including title, price, URL, and condition,
// and save these details as JSON files in a specified directory.
package main

import (
	"fmt"
	"jeovahfialho/internal/crawler"
	"os"
)

func main() {
	// baseURL specifies the URL of the eBay seller's page to be crawled.
	// Replace the placeholder URL with the actual URL you intend to crawl.
	baseURL := "https://www.ebay.com/sch/i.html?_ssn=garlandcomputer"

	// dataFolder specifies the directory where the JSON files containing
	// the extracted product information will be saved. Ensure this directory
	// exists prior to running the crawler to prevent errors.
	dataFolder := "./data"

	// ebayCrawler is an instance of the Crawler type, initialized with the
	// baseURL of the eBay seller's page and the path to the directory where
	// the results will be stored.
	ebayCrawler := crawler.NewCrawler(baseURL, dataFolder)

	// conditionFilter allows specifying a condition to filter the items by,
	// such as "New" or "Pre-Owned". Leave this empty ("") to crawl items of
	// all conditions. Adjust this parameter based on the specific requirements.
	conditionFilter := "" // Example: "New", "Pre-Owned", or "" for all items.

	// Initiates the crawling process. The Crawl method visits the specified eBay page,
	// extracts details of each listed item, and saves the information as JSON files
	// in the specified data folder. It can filter items based on the condition if specified.
	fmt.Println("Starting the crawling process...")
	err := ebayCrawler.Crawl(conditionFilter)
	if err != nil {
		fmt.Printf("Error during crawling: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Crawling completed successfully.")
}
