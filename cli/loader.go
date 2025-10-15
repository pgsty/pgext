/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// Default data URLs for different regions
const (
	DefaultExtensionURL  = "https://repo.pigsty.io/ext/data/extension.csv"
	DefaultRepositoryURL = "https://repo.pigsty.io/ext/data/repository.csv"
	ChinaExtensionURL    = "https://repo.pigsty.cc/ext/data/extension.csv"
	ChinaRepositoryURL   = "https://repo.pigsty.cc/ext/data/repository.csv"
)

// Region represents the deployment region
type Region string

const (
	RegionDefault Region = "default"
	RegionChina   Region = "china"
)

// LoadTable loads CSV data into a pgext table
// tableName can be: extension/ext/e, repository/repo/r, category/cat/c, or any table name
// source can be: URL, file path, "-" for stdin, or empty for default URL
func LoadTable(tableName string, source string, region Region) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	// Normalize table name and get default source if needed
	table, defaultSource := normalizeTable(tableName, region)
	if table == "" {
		return fmt.Errorf("invalid table name: %s", tableName)
	}

	// Use default source if not specified
	if source == "" {
		source = defaultSource
		if source == "" {
			return fmt.Errorf("no source specified for table %s", table)
		}
		logrus.Infof("loading %s from default: %s", table, source)
	} else {
		logrus.Infof("loading %s from: %s", table, source)
	}

	// Fetch data
	var data []byte
	var err error

	// Fetch the data
	data, err = fetchContent(source)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	// Load into database using COPY
	return loadCSVData(table, data)
}

// normalizeTable normalizes table name and returns default source URL
func normalizeTable(name string, region Region) (table string, defaultURL string) {
	name = strings.ToLower(strings.TrimSpace(name))

	// Map aliases to canonical names
	switch name {
	case "extension", "extensions", "ext", "e":
		table = "extension"
		if region == RegionChina {
			defaultURL = ChinaExtensionURL
		} else {
			defaultURL = DefaultExtensionURL
		}
	case "repository", "repositories", "repo", "r":
		table = "repository"
		if region == RegionChina {
			defaultURL = ChinaRepositoryURL
		} else {
			defaultURL = DefaultRepositoryURL
		}
	default:
		// For other tables, use the name as-is
		table = name
		defaultURL = ""
	}

	return table, defaultURL
}

// validTables defines the allowed table names for COPY operations
var validTables = map[string]bool{
	"extension":  true,
	"repository": true,
	"category":   true,
	"yum":        true,
	"apt":        true,
	"package":    true,
	"matrix":     true,
	"repo_data":  true,
}

// loadCSVData loads CSV data into a pgext table using COPY protocol
func loadCSVData(table string, data []byte) error {
	// Validate table name against whitelist to prevent SQL injection
	if !isValidTableName(table) {
		return fmt.Errorf("invalid or unauthorized table name: %s", table)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Truncate existing data using parameterized identifier
	truncateSQL := fmt.Sprintf("TRUNCATE TABLE pgext.%s CASCADE", pgx.Identifier{table}.Sanitize())
	if _, err := tx.Exec(ctx, truncateSQL); err != nil {
		return fmt.Errorf("failed to truncate table %s: %w", table, err)
	}

	// Use COPY protocol with sanitized identifier
	reader := bytes.NewReader(data)
	copySQL := fmt.Sprintf("COPY pgext.%s FROM STDIN WITH (FORMAT CSV, HEADER true)",
		pgx.Identifier{table}.Sanitize())

	copyResult, err := tx.Conn().PgConn().CopyFrom(ctx, reader, copySQL)
	if err != nil {
		return fmt.Errorf("failed to COPY data to %s: %w", table, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	logrus.Infof("loaded %d records into pgext.%s", copyResult.RowsAffected(), table)

	return nil
}

// isValidTableName checks if a table name is in the whitelist
func isValidTableName(name string) bool {
	return validTables[name]
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return defaultValue
}

// Use the shared embedded extension CSV data
// fetchContent fetches content from various sources
func fetchContent(source string) ([]byte, error) {
	// Empty source or stdin
	if source == "" || source == "-" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to read from stdin: %w", err)
		}
		return data, nil
	}

	// Check if it's a URL
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		resp, err := http.Get(source)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch from URL: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return data, nil
	}

	// Otherwise, treat as file path
	data, err := os.ReadFile(source)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", source, err)
	}
	return data, nil
}
