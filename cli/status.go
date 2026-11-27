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

var statusTables = []string{
	"pg", "os", "category", "repository", "extension", "universe", "doc",
	"apt", "dnf", "bin", "pkg",
}

// StatusInfo represents the complete status of pgext metadata.
type StatusInfo struct {
	SchemaExists bool
	ActivePG     []int
	ActiveOS     []string
	TableCounts  map[string]int64
	UpdateTimes  map[string]*time.Time
}

// ShowStatus displays the metadata status.
func ShowStatus() error {
	exists, err := SchemaExists()
	if err != nil {
		return fmt.Errorf("failed to check schema: %w", err)
	}
	if !exists {
		printSchemaNotFound()
		return nil
	}

	status, err := gatherStatus()
	if err != nil {
		return fmt.Errorf("failed to gather status: %w", err)
	}
	displayStatus(status)
	return nil
}

func gatherStatus() (*StatusInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	status := &StatusInfo{
		SchemaExists: true,
		TableCounts:  make(map[string]int64, len(statusTables)),
		UpdateTimes:  make(map[string]*time.Time),
	}

	pgVersions, err := ActivePGVersions(ctx)
	if err != nil {
		logrus.Warnf("failed to get active PG versions: %v", err)
	}
	status.ActivePG = pgVersions

	osList, err := GetActiveOS(ctx)
	if err != nil {
		logrus.Warnf("failed to get active OS list: %v", err)
	}
	status.ActiveOS = osList

	for _, table := range statusTables {
		var count int64
		if err := QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM pgext.%s", table)).Scan(&count); err != nil {
			logrus.Debugf("failed to count %s: %v", table, err)
			count = 0
		}
		status.TableCounts[table] = count
	}

	times, err := getUpdateTimes(ctx)
	if err != nil {
		logrus.Warnf("failed to get update times: %v", err)
	} else {
		status.UpdateTimes = times
	}
	return status, nil
}

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

func packageCatalogStale(times map[string]*time.Time) bool {
	return packageCatalogRefreshCommand(times) != ""
}

func packageCatalogRefreshCommand(times map[string]*time.Time) string {
	fetchTime := times["fetch"]
	parseTime := times["parse"]
	if fetchTime == nil && parseTime != nil {
		return "pgext fetch"
	}
	if fetchTime != nil && (parseTime == nil || fetchTime.After(*parseTime)) {
		return "pgext parse"
	}
	recapTime := times["recap"]
	if parseTime != nil && (recapTime == nil || parseTime.After(*recapTime)) {
		return "pgext recap"
	}
	return ""
}

func displayStatus(status *StatusInfo) {
	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║            pgext Status                   ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Printf("\nDatabase: %s\n", SanitizeURL(PGURL))

	fmt.Println("\n📌 Active PostgreSQL Versions:")
	if len(status.ActivePG) > 0 {
		parts := make([]string, len(status.ActivePG))
		for i, pg := range status.ActivePG {
			parts[i] = fmt.Sprintf("%d", pg)
		}
		fmt.Printf("  %s\n", strings.Join(parts, ", "))
	} else {
		fmt.Println("  (none)")
	}

	fmt.Println("\n🐧 Active Linux Distributions:")
	if len(status.ActiveOS) > 0 {
		fmt.Printf("  %s\n", strings.Join(status.ActiveOS, ", "))
	} else {
		fmt.Println("  (none)")
	}

	fmt.Println("\n📊 Metadata Tables:")
	fmt.Println("┌─────────────────┬──────────┐")
	for _, table := range statusTables {
		fmt.Printf("│ %-15s │ %8d │\n", table, status.TableCounts[table])
	}
	fmt.Println("└─────────────────┴──────────┘")

	fmt.Println("\n🕐 Last Update:")
	printUpdateTime("Fetch", status.UpdateTimes["fetch"])
	printUpdateTime("Parse", status.UpdateTimes["parse"])
	printUpdateTime("Recap", status.UpdateTimes["recap"])
	if refreshCommand := packageCatalogRefreshCommand(status.UpdateTimes); refreshCommand != "" {
		fmt.Printf("\n⚠️  Package catalog is stale; run '%s' before serving or generating content.\n", refreshCommand)
	}
	fmt.Println()
}

func printUpdateTime(label string, t *time.Time) {
	if t != nil {
		fmt.Printf("  %-6s %s\n", label+":", t.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("  %-6s %s\n", label+":", "never")
	}
}

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

// ShowRepositories displays repository summary.
func ShowRepositories() error {
	ctx := context.Background()
	rows, err := QueryContext(ctx, `
		SELECT r.id, r.type, r.os, r.org, COALESCE(d.size, 0) AS size, d.data IS NOT NULL AS has_data
		FROM pgext.repository r
		LEFT JOIN pgext.repo_data d ON r.id = d.id
		ORDER BY r.id
	`)
	if err != nil {
		return fmt.Errorf("failed to query repositories: %w", err)
	}
	defer rows.Close()

	sep := strings.Repeat("─", 85)
	fmt.Println("\n📦 Repository Summary:")
	fmt.Println(sep)
	fmt.Printf("%-25s %-5s %-15s %-10s %-12s %-10s\n", "Repository", "Type", "OS", "Org", "Size", "Status")
	fmt.Println(sep)

	count, hasData := 0, 0
	var totalSize int64
	for rows.Next() {
		var id, repoType, osName, org string
		var size int64
		var dataExists bool
		if err := rows.Scan(&id, &repoType, &osName, &org, &size, &dataExists); err != nil {
			logrus.Warnf("scan repository row: %v", err)
			continue
		}
		status := "❌ Missing"
		if dataExists {
			status = "✅ OK"
			hasData++
		}
		sizeStr := "-"
		if size > 0 {
			sizeStr = FormatSize(size)
		}
		fmt.Printf("%-25s %-5s %-15s %-10s %-12s %-10s\n", id, repoType, osName, org, sizeStr, status)
		count++
		totalSize += size
	}

	fmt.Println(sep)
	fmt.Printf("Total: %d repositories, %d with data, Total size: %s\n\n", count, hasData, FormatSize(totalSize))
	return rows.Err()
}
