// Package model defines the data structures used by the eBay crawler application.
// These structures represent the core data that is extracted from the eBay listings
// and serialized into JSON format for storage.
package model

// Product represents the essential information about an eBay listing that is
// extracted by the crawler. This includes the title, price, URL, and condition
// of the product. Each field is tagged with `json:"name"` to specify how it should
// be serialized into JSON format, allowing for easy storage and retrieval of data.
type Product struct {
	Title     string `json:"title"`     // Title of the product listing
	Price     string `json:"price"`     // Price of the product, as a string to accommodate various formats
	URL       string `json:"url"`       // Direct URL to the product listing on eBay
	Condition string `json:"condition"` // Condition of the product (e.g., "New", "Pre-Owned")
}
