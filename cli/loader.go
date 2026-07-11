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

	"pgext/db"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// Default data URLs for different regions
const (
	DefaultExtensionURL  = "https://repo.pigsty.io/ext/data/extension.csv"
	DefaultRepositoryURL = "https://repo.pigsty.io/ext/data/repository.csv"
	DefaultSourceURL     = "https://repo.pigsty.io/ext/src/"
	ChinaExtensionURL    = "https://repo.pigsty.cc/ext/data/extension.csv"
	ChinaRepositoryURL   = "https://repo.pigsty.cc/ext/data/repository.csv"
	ChinaSourceURL       = "https://repo.pigsty.cc/ext/src/"
)

// Region represents the deployment region
type Region string

const (
	RegionDefault Region = "default"
	RegionChina   Region = "china"
)

// LoadTable replaces one registered pgext table from CSV data.
// An empty source uses the CSV embedded in the pgext binary. An explicit source
// may be a URL, a file path, or "-" for stdin. Region is retained for API
// compatibility; embedded catalog data is region-independent.
func LoadTable(tableName string, source string, region Region) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	_ = region
	embedded, err := db.GetCSVFile(tableName)
	if err != nil {
		return err
	}

	data := embedded.Content
	if source == "" {
		logrus.Infof("loading %s from embedded %s.csv", embedded.Name, embedded.Name)
	} else {
		logrus.Infof("loading %s from: %s", embedded.Name, source)
		data, err = fetchContent(source)
		if err != nil {
			return fmt.Errorf("failed to fetch data: %w", err)
		}
	}

	// Load into database using COPY
	return loadCSVData(embedded.Name, data)
}

// EmbeddedTableNames returns canonical tables supported by LoadTable.
func EmbeddedTableNames() []string { return db.CSVTableNames() }

// IsEmbeddedTable reports whether a canonical table name or alias is supported.
func IsEmbeddedTable(name string) bool { return db.IsCSVTable(name) }

// loadCSVData loads CSV data into a pgext table using COPY protocol
func loadCSVData(table string, data []byte) error {
	// Validate table name against whitelist to prevent SQL injection
	if !db.IsCSVTable(table) {
		return fmt.Errorf("invalid or unauthorized table name: %s", table)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// This replacement API intentionally cascades. The CLI help warns callers
	// that dependent catalog tables may be cleared as a result. Package inputs
	// also clear pkg explicitly: repository is not a direct pkg foreign key, so
	// PostgreSQL's cascade alone would otherwise leave a stale availability map.
	logrus.Warnf("replacing pgext.%s with TRUNCATE ... CASCADE", table)
	targets := []string{fmt.Sprintf("pgext.%s", pgx.Identifier{table}.Sanitize())}
	if loadInvalidatesPackageCatalog(table) {
		targets = append(targets, liveTable("pkg"))
	}
	truncateSQL := "TRUNCATE TABLE " + strings.Join(targets, ", ") + " CASCADE"
	if _, err := tx.Exec(ctx, truncateSQL); err != nil {
		return fmt.Errorf("failed to truncate table %s: %w", table, err)
	}

	rows, err := copyCSVToTable(ctx, tx, table, data)
	if err != nil {
		return fmt.Errorf("failed to COPY data to %s: %w", table, err)
	}
	if err := updateStatusAfterLoad(ctx, tx, table); err != nil {
		return fmt.Errorf("failed to invalidate status after loading %s: %w", table, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	logrus.Infof("loaded %d records into pgext.%s", rows, table)

	return nil
}

func loadInvalidatesPackageCatalog(table string) bool {
	switch table {
	case "pg", "os", "repository", "extension":
		return true
	default:
		return false
	}
}

func updateStatusAfterLoad(ctx context.Context, tx pgx.Tx, table string) error {
	var statement string
	switch table {
	case "pg":
		statement = "UPDATE pgext.status SET recap_time = NULL, pkg_time = NULL WHERE id = 0"
	case "os":
		statement = "UPDATE pgext.status SET fetch_time = NULL, recap_time = NULL, pkg_time = NULL WHERE id = 0"
	case "repository":
		statement = "UPDATE pgext.status SET repo_time = clock_timestamp(), fetch_time = NULL, recap_time = NULL, pkg_time = NULL WHERE id = 0"
	case "extension":
		statement = "UPDATE pgext.status SET ext_time = clock_timestamp(), recap_time = NULL, pkg_time = NULL WHERE id = 0"
	case "category", "universe":
		statement = "UPDATE pgext.status SET ext_time = clock_timestamp() WHERE id = 0"
	default:
		return fmt.Errorf("unsupported embedded table %q", table)
	}
	_, err := tx.Exec(ctx, statement)
	return err
}

// copyCSVToTable copies CSV data using the caller's transaction.
func copyCSVToTable(ctx context.Context, tx pgx.Tx, table string, data []byte) (int64, error) {
	copySQL := fmt.Sprintf("COPY pgext.%s FROM STDIN WITH (FORMAT CSV, HEADER true)",
		pgx.Identifier{table}.Sanitize())
	result, err := tx.Conn().PgConn().CopyFrom(ctx, bytes.NewReader(data), copySQL)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

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
