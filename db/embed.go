/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package db

import (
	"embed"
	"fmt"
	"strings"
)

const (
	universeDDLBegin = "-- BEGIN PGEXT UNIVERSE DDL"
	universeDDLEnd   = "-- END PGEXT UNIVERSE DDL"
	matrixDDLBegin   = "-- BEGIN PGEXT MATRIX DDL"
	matrixDDLEnd     = "-- END PGEXT MATRIX DDL"
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

// GetUniverseSchema returns the independently executable universe DDL block.
// It is used to add the universe catalog to databases created by older pgext
// versions without maintaining a second copy of the table definition.
func GetUniverseSchema() (string, error) {
	schema, err := GetSchema()
	if err != nil {
		return "", err
	}
	return extractSchemaBlock(schema, universeDDLBegin, universeDDLEnd)
}

// GetMatrixSchema returns the independently executable global matrix DDL.
// It upgrades existing catalogs without duplicating the materialized-view SQL.
func GetMatrixSchema() (string, error) {
	schema, err := GetSchema()
	if err != nil {
		return "", err
	}
	return extractSchemaBlock(schema, matrixDDLBegin, matrixDDLEnd)
}

func extractSchemaBlock(schema, begin, end string) (string, error) {
	start := strings.Index(schema, begin)
	if start < 0 {
		return "", fmt.Errorf("embedded schema.sql is missing %q", begin)
	}
	start += len(begin)
	endPos := strings.Index(schema[start:], end)
	if endPos < 0 {
		return "", fmt.Errorf("embedded schema.sql is missing %q", end)
	}

	ddl := strings.TrimSpace(schema[start : start+endPos])
	if ddl == "" {
		return "", fmt.Errorf("embedded DDL block %q is empty", begin)
	}
	return ddl + "\n", nil
}

// CSVFile represents an embedded CSV file with its metadata
type CSVFile struct {
	Name    string // File name without extension
	Content []byte // Raw CSV content
}

// csvTable describes one table that may be populated from an embedded CSV.
// The registry order is the load order; extension must precede its universe
// superset and dimension tables must precede tables that reference them.
type csvTable struct {
	Name    string
	Asset   string
	Aliases []string
}

var csvTables = []csvTable{
	{Name: "pg", Asset: "pg.csv"},
	{Name: "os", Asset: "os.csv"},
	{Name: "category", Asset: "category.csv", Aliases: []string{"cat", "c"}},
	{Name: "repository", Asset: "repository.csv", Aliases: []string{"repositories", "repo", "r"}},
	{Name: "extension", Asset: "extension.csv", Aliases: []string{"extensions", "ext", "e"}},
	{Name: "universe", Asset: "universe.csv"},
}

// CSVTableNames returns canonical embedded table names in load order.
func CSVTableNames() []string {
	names := make([]string, 0, len(csvTables))
	for _, table := range csvTables {
		names = append(names, table.Name)
	}
	return names
}

// IsCSVTable reports whether name identifies an embedded CSV table or alias.
func IsCSVTable(name string) bool {
	_, ok := lookupCSVTable(name)
	return ok
}

// GetCSVFile returns the registered embedded CSV for a canonical name or alias.
func GetCSVFile(name string) (CSVFile, error) {
	table, ok := lookupCSVTable(name)
	if !ok {
		return CSVFile{}, fmt.Errorf("unknown embedded CSV table %q", name)
	}

	content, err := embeddedFS.ReadFile(table.Asset)
	if err != nil {
		return CSVFile{}, fmt.Errorf("failed to read %s: %w", table.Asset, err)
	}
	return CSVFile{Name: table.Name, Content: content}, nil
}

// GetCSVFiles returns all registered embedded CSV files in dependency order.
func GetCSVFiles() ([]CSVFile, error) {
	files := make([]CSVFile, 0, len(csvTables))
	for _, table := range csvTables {
		file, err := GetCSVFile(table.Name)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func lookupCSVTable(name string) (csvTable, bool) {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, table := range csvTables {
		if name == table.Name {
			return table, true
		}
		for _, alias := range table.Aliases {
			if name == alias {
				return table, true
			}
		}
	}
	return csvTable{}, false
}

// ReadFile returns content of any embedded file by name
func ReadFile(name string) ([]byte, error) {
	data, err := embeddedFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded file %s: %w", name, err)
	}
	return data, nil
}
