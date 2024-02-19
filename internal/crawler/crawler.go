package crawler

import (
	"fmt"
	"net/http"
	"sync"

	"jeovahfialho/internal/model"
	"jeovahfialho/internal/parser"
	"jeovahfialho/internal/writer"

	"github.com/PuerkitoBio/goquery"
)

// Crawler represents the crawler for eBay listings, encapsulating the base URL to crawl,
// the target folder for data storage, and an HTTP client for web requests.
type Crawler struct {
	BaseURL    string
	DataFolder string
	Client     *http.Client
}

// NewCrawler creates and returns a new instance of Crawler, initialized with
// the provided base URL and data folder path. It sets up a new HTTP client for the crawler.
func NewCrawler(baseURL, dataFolder string) *Crawler {
	return &Crawler{
		BaseURL:    baseURL,
		DataFolder: dataFolder,
		Client:     &http.Client{},
	}
}

// Crawl initiates the crawling process on the eBay page specified by BaseURL.
// It extracts item details and stores them as JSON files in the DataFolder.
// The function can filter items by a specified condition if provided.
func (c *Crawler) Crawl(conditionFilter string) error {
	// Attempt to fetch the page content from the BaseURL.
	resp, err := c.Client.Get(c.BaseURL)
	if err != nil {
		return fmt.Errorf("failed to get page: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful HTTP status code.
	if resp.StatusCode != 200 {
		return fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Parse the fetched page content using goquery.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse page: %v", err)
	}

	// Extract product information from the parsed document.
	products, err := parser.ExtractProductInfo(doc)
	if err != nil {
		return fmt.Errorf("failed to extract product info: %v", err)
	}

	// Optionally filter extracted products by condition.
	if conditionFilter != "" {
		products = filterProductsByCondition(products, conditionFilter)
	}

	// Asynchronously write product data to JSON files in the DataFolder.
	var wg sync.WaitGroup
	for _, product := range products {
		wg.Add(1)
		go func(p model.Product) {
			defer wg.Done()
			if err := writer.WriteProductToFile(p, c.DataFolder); err != nil {
				fmt.Printf("Failed to write product data to file: %v\n", err)
			}
		}(product)
	}
	wg.Wait()

	return nil
}

// filterProductsByCondition filters a slice of products based on the specified condition
// (e.g., "New", "Pre-Owned") and returns a new slice containing only products that match the condition.
func filterProductsByCondition(products []model.Product, condition string) []model.Product {
	filtered := make([]model.Product, 0)
	for _, product := range products {
		if product.Condition == condition {
			filtered = append(filtered, product)
		}
	}
	return filtered
}
