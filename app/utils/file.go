package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/DanteLorenzo/gotch/app/types" // Import the custom types package
)

// ReadPatterns reads a JSON file from the specified file path and decodes it into a slice of Pattern structs.
// It returns the slice of patterns or an error if something goes wrong.
func ReadPatterns(filePath string) ([]types.Pattern, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err // Return an error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after the function completes

	// Create a slice to hold the decoded patterns
	var patterns []types.Pattern

	// Create a JSON decoder and decode the file content into the slice of patterns
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&patterns)
	if err != nil {
		return nil, err // Return an error if the JSON decoding fails
	}

	// Check if the patterns slice is empty
	if len(patterns) == 0 {
		return nil, errors.New("patterns.json is empty") // Return an error if the file is empty
	}

	// Return the decoded patterns and a nil error
	return patterns, nil
}
