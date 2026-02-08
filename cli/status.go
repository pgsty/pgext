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

// StatusInfo represents the complete status of pgext metadata
type StatusInfo struct {
	SchemaExists bool
	ActivePG     []int
	ActiveOS     []string
	TableCounts  map[string]int64
	UpdateTimes  map[string]*time.Time
}

// ShowStatus displays the metadata status
func ShowStatus() error {
	// Check schema existence first
	exists, err := SchemaExists()
	if err != nil {
		return fmt.Errorf("failed to check schema: %w", err)
	}

	if !exists {
		printSchemaNotFound()
		return nil
	}

	// Gather status information
	status, err := gatherStatus()
	if err != nil {
		return fmt.Errorf("failed to gather status: %w", err)
	}

	// Display status
	displayStatus(status)

	return nil
}

// gatherStatus collects all status information from database
func gatherStatus() (*StatusInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	status := &StatusInfo{
		SchemaExists: true,
		TableCounts:  make(map[string]int64),
		UpdateTimes:  make(map[string]*time.Time),
	}

	// Get active PG versions (descending)
	pgVersions, err := getActivePGVersions(ctx)
	if err != nil {
		logrus.Warnf("failed to get active PG versions: %v", err)
	}
	status.ActivePG = pgVersions

	// Get active OS list (ordered by major)
	osList, err := getActiveOSVersions(ctx)
	if err != nil {
		logrus.Warnf("failed to get active OS list: %v", err)
	}
	status.ActiveOS = osList

	// Get table counts
	tables := []string{"pg", "os", "category", "repository", "extension",
		"apt", "dnf", "bin", "pkg"}
	for _, table := range tables {
		count, err := getTableCount(ctx, table)
		if err != nil {
			logrus.Debugf("failed to count %s: %v", table, err)
			count = 0
		}
		status.TableCounts[table] = count
	}

	// Get update timestamps
	times, err := getUpdateTimes(ctx)
	if err != nil {
		logrus.Warnf("failed to get update times: %v", err)
	} else {
		status.UpdateTimes = times
	}

	return status, nil
}

// getActivePGVersions retrieves active PostgreSQL versions in descending order
func getActivePGVersions(ctx context.Context) ([]int, error) {
	query := "SELECT pg FROM pgext.pg WHERE active ORDER BY pg DESC"

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []int
	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err != nil {
			return nil, err
		}
		versions = append(versions, pg)
	}
	return versions, rows.Err()
}

// getActiveOSVersions retrieves active OS list ordered by major version
func getActiveOSVersions(ctx context.Context) ([]string, error) {
	query := "SELECT os FROM pgext.active_os ORDER BY os"

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var osList []string
	for rows.Next() {
		var os string
		if err := rows.Scan(&os); err != nil {
			return nil, err
		}
		osList = append(osList, os)
	}
	return osList, rows.Err()
}

// getTableCount returns the count for a specific table
func getTableCount(ctx context.Context, table string) (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM pgext.%s", table)
	err := QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

// getUpdateTimes retrieves update timestamps from status table
func getUpdateTimes(ctx context.Context) (map[string]*time.Time, error) {
	times := make(map[string]*time.Time)

	var fetchTime, parseTime, recapTime sql.NullTime
	err := QueryRowContext(ctx, `
		SELECT fetch_time, parse_time, recap_time
		FROM pgext.status
		WHERE id = 0
	`).Scan(&fetchTime, &parseTime, &recapTime)

	if err != nil {
		return times, err
	}

	if fetchTime.Valid {
		times["fetch"] = &fetchTime.Time
	}
	if parseTime.Valid {
		times["parse"] = &parseTime.Time
	}
	if recapTime.Valid {
		times["recap"] = &recapTime.Time
	}

	return times, nil
}

// displayStatus prints the status information in a beautiful format
func displayStatus(status *StatusInfo) {
	// Header
	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║            pgext Status                   ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Printf("\nDatabase: %s\n", SanitizeURL(PGURL))

	// Active PG Versions
	fmt.Println("\n📌 Active PostgreSQL Versions:")
	if len(status.ActivePG) > 0 {
		fmt.Print("  ")
		for i, pg := range status.ActivePG {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", pg)
		}
		fmt.Println()
	} else {
		fmt.Println("  (none)")
	}

	// Active Linux Distributions
	fmt.Println("\n🐧 Active Linux Distributions:")
	if len(status.ActiveOS) > 0 {
		fmt.Print("  ")
		for i, os := range status.ActiveOS {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(os)
		}
		fmt.Println()
	} else {
		fmt.Println("  (none)")
	}

	// Metadata Tables (merged)
	fmt.Println("\n📊 Metadata Tables:")
	fmt.Println("┌─────────────────┬──────────┐")
	fmt.Printf("│ %-15s │ %8d │\n", "pg", status.TableCounts["pg"])
	fmt.Printf("│ %-15s │ %8d │\n", "os", status.TableCounts["os"])
	fmt.Printf("│ %-15s │ %8d │\n", "category", status.TableCounts["category"])
	fmt.Printf("│ %-15s │ %8d │\n", "repository", status.TableCounts["repository"])
	fmt.Printf("│ %-15s │ %8d │\n", "extension", status.TableCounts["extension"])
	fmt.Printf("│ %-15s │ %8d │\n", "apt", status.TableCounts["apt"])
	fmt.Printf("│ %-15s │ %8d │\n", "dnf", status.TableCounts["dnf"])
	fmt.Printf("│ %-15s │ %8d │\n", "bin", status.TableCounts["bin"])
	fmt.Printf("│ %-15s │ %8d │\n", "pkg", status.TableCounts["pkg"])
	fmt.Println("└─────────────────┴──────────┘")

	// Update Times
	fmt.Println("\n🕐 Last Update:")
	printUpdateTime("Fetch", status.UpdateTimes["fetch"])
	printUpdateTime("Parse", status.UpdateTimes["parse"])
	printUpdateTime("Recap", status.UpdateTimes["recap"])

	fmt.Println()
}

// printUpdateTime formats and prints an update timestamp
func printUpdateTime(label string, t *time.Time) {
	if t != nil {
		fmt.Printf("  %-6s %s\n", label+":", t.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  %-6s %s\n", label+":", "never")
	}
}

// printSchemaNotFound displays message when schema doesn't exist
func printSchemaNotFound() {
	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║         pgext Metadata Status             ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Printf("\nDatabase: %s\n", SanitizeURL(PGURL))
	fmt.Println("Schema:   not found")
	fmt.Println()
	fmt.Println("⚠️  Schema not initialized. Run 'pgext init' to create it.")
	fmt.Println()
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
