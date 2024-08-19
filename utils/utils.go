package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Metadata struct {
	Filename  string `json:"filename"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	LowerName string `json:"lower_name"`
	Extension string `json:"extension"`
}

// GenerateMetadata recursively scans the given directory and returns a slice of FileMetadata objects.
func GenerateMetadata(rootDir string) ([]Metadata, error) {
	var metadata []Metadata

	// Define the root folder if null
	if rootDir == "" {
		rootDir = "gitignores"

	}

	// Check if the root folder exists
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("The folder '%s' does not exist.\n", rootDir)
		return metadata, err
	}

	// Walk through the directory recursively
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Extract name and extension
		name := strings.SplitN(info.Name(), ".", 2)[0]
		extension := filepath.Ext(info.Name())

		// Add metadata for this file
		metadata = append(metadata, Metadata{
			Name:      name,
			LowerName: strings.ToLower(name), // For case-insensitive search
			Location:  path,
			Extension: extension,
			Filename:  info.Name(),
		})

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return metadata, err
	}

	// Check if any metadata was collected
	if len(metadata) == 0 {
		fmt.Println("No files were found in the directory.")
		return metadata, errors.New("no files were found in the directory")
	}

	// Save regular metadata to a JSON file
	saveToFile("gitignore-metadata.json", metadata)

	fmt.Println("metadata is generated successfully.")

	return metadata, nil
}

func saveToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating JSON file '%s': %v\n", filename, err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Error encoding JSON for '%s': %v\n", filename, err)
	}
	return err
}

func GetMetadata(path string) ([]Metadata, error) {
	if path == "" {
		path = "./gitignore-metadata.json"
	}

	metadata := []Metadata{}

	// Read the file content
	file, err := os.ReadFile(path)
	if err != nil {
		return metadata, err
	}

	// Unmarshal the JSON content into the metadata slice
	err = json.Unmarshal(file, &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}
