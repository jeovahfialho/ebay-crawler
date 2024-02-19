// Package parser provides the functionality to parse HTML documents and
// extract meaningful information from them, specifically targeting product
// listings on eBay pages. It leverages the goquery library to navigate
// and parse the structure of HTML documents.
package parser

import (
	"strings"

	"jeovahfialho/internal/model"

	"github.com/PuerkitoBio/goquery"
)

// ExtractProductInfo accepts an HTML document (represented by a goquery.Document)
// and extracts information about products listed on the page. It specifically looks
// for the title, price, URL, and condition of each product, organizing this data
// into a slice of Product structs.
func ExtractProductInfo(doc *goquery.Document) ([]model.Product, error) {
	var products []model.Product

	// Iterate over each element matched by the .s-item__wrapper selector,
	// which is assumed to be the container for individual product listings.
	doc.Find(".s-item__wrapper").Each(func(i int, s *goquery.Selection) {
		// Extract the title, price, URL, and condition from each listing.
		// These selectors are based on the assumed structure of eBay product listings.
		title := s.Find(".s-item__title span").Text()
		price := s.Find(".s-item__price").Text()
		url, _ := s.Find(".s-item__image a").Attr("href")
		condition := s.Find(".s-item__subtitle .SECONDARY_INFO").Text()

		// Clean up the extracted data by trimming any surrounding whitespace.
		title = strings.TrimSpace(title)
		price = strings.TrimSpace(price)
		url = strings.TrimSpace(url)
		condition = strings.TrimSpace(condition)

		// Create a Product struct from the extracted data and append it to
		// the slice of products to be returned.
		product := model.Product{
			Title:     title,
			Price:     price,
			URL:       url,
			Condition: condition,
		}

		products = append(products, product)
	})

	// Return the slice of extracted Product structs.
	return products, nil
}
