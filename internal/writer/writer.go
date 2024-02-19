// Package writer provides functionality to write product data to JSON files.
// It handles the creation of files and ensures data is correctly serialized into JSON format.
package writer

import (
	"encoding/json"
	"fmt"
	"jeovahfialho/internal/model"
	"os"
	"path/filepath"
	"strings"
)

// WriteProductToFile writes the product data to a JSON file in the specified base directory.
// It creates the directory if it doesn't exist, and names the file using the product's unique ID extracted from its URL.
func WriteProductToFile(product model.Product, basePath string) error {
	// Check if the base directory exists, and create it if it does not.
	_, err := os.Stat(basePath)
	if os.IsNotExist(err) {
		fmt.Printf("The directory %s does not exist; creating it now...\n", basePath)
		os.Mkdir(basePath, os.ModePerm) // Create the directory with default permissions.
	}

	// Extract the item ID from the product URL to use as the filename.
	parts := strings.Split(product.URL, "/")
	itemID := strings.Split(parts[len(parts)-1], "?")[0] // Assumes the item ID is the last part of the URL, before any query parameters.
	fileName := itemID + ".json"                         // Append ".json" to create the filename.
	filePath := filepath.Join(basePath, fileName)        // Create the full file path.

	// Attempt to create the file at the specified path.
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create the file: %v", err) // Return an error if file creation fails.
	}
	defer file.Close() // Ensure the file is closed after writing is complete.
	fmt.Printf("File %s created successfully.\n", filePath)

	// Encode the product data as JSON and write it to the file.
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(product); err != nil {
		return err // Return any errors encountered during JSON encoding.
	}

	return nil // Return nil to indicate success.
}
