/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// MetaStatus represents the status of pgext metadata
type MetaStatus struct {
	SchemaExists    bool
	CategoryCount   int64
	RepositoryCount int64
	ExtensionCount  int64
	RepoDataCount   int64
	YumCount        int64
	AptCount        int64
	PackageCount    int64
	MatrixCount     int64
	FetchTime       *time.Time
	ParseTime       *time.Time
	RecapTime       *time.Time
}

// GetStatus returns the current status of pgext metadata
func GetStatus() (*MetaStatus, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	status := &MetaStatus{}

	// Check schema existence
	var err error
	status.SchemaExists, err = SchemaExists()
	if err != nil {
		return nil, err
	}

	if !status.SchemaExists {
		return status, nil
	}

	// Get table counts
	type tableCount struct {
		table string
		count *int64
	}

	tables := []tableCount{
		{"pgext.category", &status.CategoryCount},
		{"pgext.repository", &status.RepositoryCount},
		{"pgext.extension", &status.ExtensionCount},
		{"pgext.repo_data", &status.RepoDataCount},
		{"pgext.dnf", &status.YumCount},
		{"pgext.apt", &status.AptCount},
		{"pgext.bin", &status.PackageCount},
		{"pgext.pkg", &status.MatrixCount},
	}

	for _, tc := range tables {
		err := QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s", tc.table)).Scan(tc.count)
		if err != nil {
			logrus.Debugf("failed to count %s: %v", tc.table, err)
			*tc.count = 0
		}
	}

	// Get timestamps from pgext.status
	var fetchTime, parseTime, recapTime sql.NullTime
	err = QueryRowContext(ctx, `
		SELECT fetch_time, parse_time, recap_time
		FROM pgext.status
		WHERE id = 0
	`).Scan(&fetchTime, &parseTime, &recapTime)

	if err == nil {
		if fetchTime.Valid {
			status.FetchTime = &fetchTime.Time
		}
		if parseTime.Valid {
			status.ParseTime = &parseTime.Time
		}
		if recapTime.Valid {
			status.RecapTime = &recapTime.Time
		}
	}

	return status, nil
}

// PrintStatus prints the metadata status in a formatted way
func PrintStatus() error {
	status, err := GetStatus()
	if err != nil {
		return err
	}

	fmt.Println("=== pgext Metadata Status ===")
	fmt.Printf("Database URL:     %s\n", sanitizeURL(PGURL))
	fmt.Printf("Schema Exists:    %v\n", status.SchemaExists)

	if !status.SchemaExists {
		fmt.Println("\nSchema not initialized. Run 'pig meta init' to create it.")
		return nil
	}

	fmt.Println("\nTable Counts:")
	fmt.Printf("  Category:       %d\n", status.CategoryCount)
	fmt.Printf("  Repository:     %d\n", status.RepositoryCount)
	fmt.Printf("  Extension:      %d\n", status.ExtensionCount)
	fmt.Printf("  Repo Data:      %d\n", status.RepoDataCount)
	fmt.Printf("  YUM Packages:   %d\n", status.YumCount)
	fmt.Printf("  APT Packages:   %d\n", status.AptCount)
	fmt.Printf("  Package:        %d\n", status.PackageCount)
	fmt.Printf("  Matrix:         %d\n", status.MatrixCount)

	fmt.Println("\nTimestamps:")
	if status.FetchTime != nil {
		fmt.Printf("  Last Fetch:     %s\n", status.FetchTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Last Fetch:     never\n")
	}
	if status.ParseTime != nil {
		fmt.Printf("  Last Parse:     %s\n", status.ParseTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Last Parse:     never\n")
	}
	if status.RecapTime != nil {
		fmt.Printf("  Last Recap:     %s\n", status.RecapTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Last Recap:     never\n")
	}

	return nil
}

// PrintRepoSummary prints a summary of repositories
func PrintRepoSummary() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := QueryContext(ctx, `
		SELECT r.id, r.type, r.os, r.org,
		       COALESCE(rd.size, 0) as size,
		       COALESCE(rd.update_at::text, 'never') as last_update
		FROM pgext.repository r
		LEFT JOIN pgext.repo_data rd ON r.id = rd.id
		ORDER BY r.id
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("=== Repository Summary ===")
	fmt.Printf("%-30s %-6s %-12s %-8s %-12s %s\n",
		"ID", "Type", "OS", "Org", "Size", "Last Update")
	fmt.Println(strings.Repeat("-", 100))

	for rows.Next() {
		var id, repoType, os, org, lastUpdate string
		var size int64

		if err := rows.Scan(&id, &repoType, &os, &org, &size, &lastUpdate); err != nil {
			logrus.Warnf("failed to scan row: %v", err)
			continue
		}

		sizeStr := "-"
		if size > 0 {
			sizeStr = formatBytes(size)
		}

		updateStr := lastUpdate
		if lastUpdate == "never" {
			updateStr = "-"
		} else if len(lastUpdate) > 19 {
			updateStr = lastUpdate[:19]
		}

		fmt.Printf("%-30s %-6s %-12s %-8s %-12s %s\n",
			id, repoType, os, org, sizeStr, updateStr)
	}

	return rows.Err()
}

// ShowStatus displays the metadata status
func ShowStatus() error {
	status, err := GetStatus()
	if err != nil {
		return err
	}

	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║         pgext Metadata Status             ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Printf("\nDatabase: %s\n", sanitizeURL(PGURL))
	fmt.Printf("Schema:   %v\n", status.SchemaExists)

	if !status.SchemaExists {
		fmt.Println("\n⚠️  Schema not initialized. Run 'pgext init' to create it.")
		return nil
	}

	fmt.Println("\n📊 Table Statistics:")
	fmt.Println("┌─────────────────┬──────────┐")
	fmt.Printf("│ %-15s │ %8d │\n", "Categories", status.CategoryCount)
	fmt.Printf("│ %-15s │ %8d │\n", "Repositories", status.RepositoryCount)
	fmt.Printf("│ %-15s │ %8d │\n", "Extensions", status.ExtensionCount)
	fmt.Printf("│ %-15s │ %8d │\n", "Repo Data", status.RepoDataCount)
	fmt.Printf("│ %-15s │ %8d │\n", "DNF Packages", status.YumCount)
	fmt.Printf("│ %-15s │ %8d │\n", "APT Packages", status.AptCount)
	fmt.Printf("│ %-15s │ %8d │\n", "Binary Packages", status.PackageCount)
	fmt.Printf("│ %-15s │ %8d │\n", "Availability", status.MatrixCount)
	fmt.Println("└─────────────────┴──────────┘")

	fmt.Println("\n🕐 Last Update Times:")
	if status.FetchTime != nil {
		fmt.Printf("  Fetch: %s\n", status.FetchTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Fetch: never\n")
	}
	if status.ParseTime != nil {
		fmt.Printf("  Parse: %s\n", status.ParseTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Parse: never\n")
	}
	if status.RecapTime != nil {
		fmt.Printf("  Recap: %s\n", status.RecapTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  Recap: never\n")
	}
	fmt.Println()

	return nil
}

// ShowRepositories displays repository summary
func ShowRepositories() error {
	ctx := context.Background()

	query := `
		SELECT
			r.id,
			r.type,
			r.os,
			r.org,
			COALESCE(d.size, 0) as size,
			d.last_modified,
			d.data IS NOT NULL as has_data
		FROM pgext.repository r
		LEFT JOIN pgext.repo_data d ON r.id = d.id
		ORDER BY r.id
	`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to query repositories: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n📦 Repository Summary:")
	fmt.Println(strings.Repeat("─", 85))
	fmt.Printf("%-25s %-5s %-15s %-10s %-12s %-10s\n",
		"Repository", "Type", "OS", "Org", "Size", "Status")
	fmt.Println(strings.Repeat("─", 85))

	count := 0
	totalSize := int64(0)
	hasData := 0

	for rows.Next() {
		var id, repoType, os, org string
		var size int64
		var lastMod sql.NullTime
		var dataExists bool

		err := rows.Scan(&id, &repoType, &os, &org, &size, &lastMod, &dataExists)
		if err != nil {
			continue
		}

		status := "❌ Missing"
		if dataExists {
			status = "✅ OK"
			hasData++
		}

		sizeStr := "-"
		if size > 0 {
			sizeStr = formatBytes(size)
		}

		fmt.Printf("%-25s %-5s %-15s %-10s %-12s %-10s\n",
			id, repoType, os, org, sizeStr, status)

		count++
		totalSize += size
	}

	fmt.Println(strings.Repeat("─", 85))
	fmt.Printf("Total: %d repositories, %d with data, Total size: %s\n\n",
		count, hasData, formatBytes(totalSize))

	return rows.Err()
}

// formatBytes formats byte size to human-readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
