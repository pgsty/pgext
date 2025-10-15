/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package db

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
)

// Embed all SQL and CSV files from the db directory
//
//go:embed *.sql *.csv
var embeddedFS embed.FS

// GetSchema returns the embedded schema.sql content
func GetSchema() (string, error) {
	data, err := embeddedFS.ReadFile("schema.sql")
	if err != nil {
		return "", fmt.Errorf("failed to read embedded schema.sql: %w", err)
	}
	return string(data), nil
}

// CSVFile represents an embedded CSV file with its metadata
type CSVFile struct {
	Name    string // File name without extension
	Content []byte // Raw CSV content
}

// GetCSVFiles returns all embedded CSV files
func GetCSVFiles() ([]CSVFile, error) {
	entries, err := embeddedFS.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded directory: %w", err)
	}

	var csvFiles []CSVFile
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasSuffix(name, ".csv") {
			content, err := embeddedFS.ReadFile(name)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %w", name, err)
			}

			// Remove .csv extension for table name
			tableName := strings.TrimSuffix(name, filepath.Ext(name))
			csvFiles = append(csvFiles, CSVFile{
				Name:    tableName,
				Content: content,
			})
		}
	}

	// Define the loading order for CSV files
	// Some tables have foreign key constraints and must be loaded in order
	loadOrder := []string{"pg", "os", "category", "repository", "extension"}

	// Sort CSV files according to the defined order
	orderedFiles := make([]CSVFile, 0, len(csvFiles))
	fileMap := make(map[string]CSVFile)

	for _, f := range csvFiles {
		fileMap[f.Name] = f
	}

	// Add files in the specified order first
	for _, name := range loadOrder {
		if f, exists := fileMap[name]; exists {
			orderedFiles = append(orderedFiles, f)
			delete(fileMap, name)
		}
	}

	// Add any remaining files that weren't in the order list
	for _, f := range fileMap {
		orderedFiles = append(orderedFiles, f)
	}

	return orderedFiles, nil
}

// GetCSVFile returns a specific CSV file by name
func GetCSVFile(name string) ([]byte, error) {
	// Try with .csv extension
	filename := name + ".csv"
	data, err := embeddedFS.ReadFile(filename)
	if err != nil {
		// Also try without adding extension (in case it's already included)
		data, err = embeddedFS.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV file %s: %w", name, err)
		}
	}
	return data, nil
}

// ReadFile returns content of any embedded file by name
func ReadFile(name string) ([]byte, error) {
	data, err := embeddedFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded file %s: %w", name, err)
	}
	return data, nil
}