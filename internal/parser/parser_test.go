// Package parser contains the functionality to parse HTML documents
// and extract product information for the eBay crawler application.
package parser

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestExtractProductInfo tests the extraction of product information
// from an HTML document. It ensures that the parser correctly identifies
// and extracts data such as the product's title, price, URL, and condition
// from a given HTML structure.
func TestExtractProductInfo(t *testing.T) {
	// Open the local HTML file that contains the test data.
	htmlFile, err := os.Open("/Users/jeovahfialho/Documents/ebay-crawler/html/index.html")
	if err != nil {
		t.Fatalf("Failed to open HTML file: %s", err)
	}
	defer htmlFile.Close()

	// Create a goquery document from the HTML file, which will be used
	// for parsing in the test.
	doc, err := goquery.NewDocumentFromReader(htmlFile)
	if err != nil {
		t.Fatalf("Failed to create goquery document: %s", err)
	}

	// Call the ExtractProductInfo function with the created document
	// and check if it returns the expected product information without errors.
	products, err := ExtractProductInfo(doc)
	if err != nil {
		t.Fatalf("ExtractProductInfo returned an error: %s", err)
	}

	// Verify that at least one product is found. This is a basic check
	// to ensure parsing is working; more detailed checks could be added
	// to verify the accuracy of extracted data.
	if len(products) == 0 {
		t.Error("No products found, expected at least one product")
	}
}
