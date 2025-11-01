/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// PkgAvailability represents package availability information.
type PkgAvailability struct {
	PG      int
	OS      string
	State   string
	Org     string
	Version string
	Count   int64
}

// ShowPackage displays package availability matrix for a given extension.
func ShowPackage(pkgName string) error {
	ctx := context.Background()

	// Get active OS and PG versions
	osList, err := GetActiveOS(ctx)
	if err != nil {
		return fmt.Errorf("get active OS: %w", err)
	}

	pgList, err := getActivePG(ctx)
	if err != nil {
		return fmt.Errorf("get active PG: %w", err)
	}

	// Get package data
	pkgData, err := getPackageData(ctx, pkgName)
	if err != nil {
		return fmt.Errorf("get package data: %w", err)
	}

	if len(pkgData) == 0 {
		fmt.Printf("\nPackage '%s' not found.\n\n", pkgName)
		return nil
	}

	// Build matrix
	matrix := buildMatrix(pkgData, osList, pgList)

	// Display table
	displayMatrix(pkgName, matrix, osList, pgList)

	return nil
}

// GetActiveOS retrieves active OS list from database.
func GetActiveOS(ctx context.Context) ([]string, error) {
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

// getActivePG retrieves active PostgreSQL versions from database.
func getActivePG(ctx context.Context) ([]int, error) {
	query := "SELECT pg FROM pgext.active_pg ORDER BY pg DESC"

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pgList []int
	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err != nil {
			return nil, err
		}
		pgList = append(pgList, pg)
	}
	return pgList, rows.Err()
}

// getPackageData retrieves package availability data from database.
func getPackageData(ctx context.Context, pkgName string) (map[string]PkgAvailability, error) {
	query := `SELECT pg, os, state, org, version, count FROM pgext.pkg WHERE pkg = $1 OR ext = $1 ORDER BY pg DESC, os`

	rows, err := QueryContext(ctx, query, pkgName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]PkgAvailability)
	for rows.Next() {
		var info PkgAvailability
		var state, org, version sql.NullString
		var count sql.NullInt64

		err := rows.Scan(&info.PG, &info.OS, &state, &org, &version, &count)
		if err != nil {
			logrus.Debugf("scan error: %v", err)
			continue
		}

		// Handle NULL values
		if state.Valid {
			info.State = state.String
		}
		if org.Valid {
			info.Org = org.String
		}
		if version.Valid {
			info.Version = version.String
		}
		if count.Valid {
			info.Count = count.Int64
		}

		key := fmt.Sprintf("%d-%s", info.PG, info.OS)
		data[key] = info
	}

	return data, rows.Err()
}

// buildMatrix creates a 2D matrix from package data.
func buildMatrix(data map[string]PkgAvailability, osList []string, pgList []int) map[string]map[int]PkgAvailability {
	matrix := make(map[string]map[int]PkgAvailability)

	for _, os := range osList {
		matrix[os] = make(map[int]PkgAvailability)
		for _, pg := range pgList {
			key := fmt.Sprintf("%d-%s", pg, os)
			if info, exists := data[key]; exists {
				matrix[os][pg] = info
			}
		}
	}

	return matrix
}

// displayMatrix renders the package availability matrix.
func displayMatrix(pkgName string, matrix map[string]map[int]PkgAvailability, osList []string, pgList []int) {
	// Calculate cell width based on content
	cellWidth := calculateCellWidth(matrix, osList, pgList)
	osWidth := calculateOSWidth(osList)

	// Print header
	fmt.Printf("\n╔══════════════════════════════════════════╗\n")
	fmt.Printf("║  Package: %-31s║\n", pkgName)
	fmt.Printf("╚══════════════════════════════════════════╝\n\n")

	// Print column headers
	fmt.Printf("%-*s", osWidth, "OS \\ PG")
	for _, pg := range pgList {
		fmt.Printf("│ %*s ", cellWidth, fmt.Sprintf("PG %d", pg))
	}
	fmt.Println()

	// Print header separator
	fmt.Print(strings.Repeat("═", osWidth))
	for range pgList {
		fmt.Print("╪" + strings.Repeat("═", cellWidth+2))
	}
	fmt.Println()

	// Print matrix rows
	for i, os := range osList {
		// First line: state and count
		fmt.Printf("%-*s", osWidth, os)
		for _, pg := range pgList {
			line1 := formatCellLine1(matrix[os][pg], cellWidth)
			fmt.Printf("│ %s ", line1)
		}
		fmt.Println()

		// Second line: org and version
		fmt.Printf("%-*s", osWidth, "")
		for _, pg := range pgList {
			line2 := formatCellLine2(matrix[os][pg], cellWidth)
			fmt.Printf("│ %s ", line2)
		}
		fmt.Println()

		// Separator between rows (except last)
		if i < len(osList)-1 {
			fmt.Print(strings.Repeat("─", osWidth))
			for range pgList {
				fmt.Print("┼" + strings.Repeat("─", cellWidth+2))
			}
			fmt.Println()
		}
	}

	// Bottom border
	fmt.Print(strings.Repeat("═", osWidth))
	for range pgList {
		fmt.Print("╧" + strings.Repeat("═", cellWidth+2))
	}
	fmt.Println("\n")
}

// calculateCellWidth calculates the minimum cell width based on content.
func calculateCellWidth(matrix map[string]map[int]PkgAvailability, osList []string, pgList []int) int {
	maxWidth := 12 // Minimum width

	for _, os := range osList {
		for _, pg := range pgList {
			info := matrix[os][pg]
			if info.State == "" {
				continue
			}

			// Calculate width needed for second line (org + version)
			width := 6 // org field width
			if info.Version != "" {
				width += 1 + len(info.Version) // space + version
			}

			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	return maxWidth
}

// calculateOSWidth calculates the OS column width.
func calculateOSWidth(osList []string) int {
	maxWidth := 8 // Minimum for "OS \ PG"

	for _, os := range osList {
		if len(os) > maxWidth {
			maxWidth = len(os)
		}
	}

	return maxWidth + 2 // Add padding
}

// formatCellLine1 formats the first line: state (6 chars left) + count (right).
func formatCellLine1(info PkgAvailability, cellWidth int) string {
	if info.State == "" {
		return strings.Repeat(" ", cellWidth)
	}

	state := info.State
	if len(state) > 6 {
		state = state[:6]
	}

	// Format: "STATE " + right-aligned count
	countStr := ""
	if info.Count > 0 {
		countStr = fmt.Sprintf("%d", info.Count)
	}

	// Left-align state (6 chars), right-align count
	stateField := fmt.Sprintf("%-6s", state)
	remaining := cellWidth - 6

	if remaining <= 0 {
		return stateField[:cellWidth]
	}

	if countStr == "" {
		return stateField + strings.Repeat(" ", remaining)
	}

	padding := remaining - len(countStr)
	if padding < 0 {
		padding = 0
	}

	return stateField + strings.Repeat(" ", padding) + countStr
}

// formatCellLine2 formats the second line: org (6 chars left) + version (right).
func formatCellLine2(info PkgAvailability, cellWidth int) string {
	if info.Version == "" {
		return strings.Repeat(" ", cellWidth)
	}

	org := info.Org
	if len(org) > 6 {
		org = org[:6]
	}

	version := info.Version

	// Format: "ORG   " + right-aligned version
	orgField := fmt.Sprintf("%-6s", org)
	remaining := cellWidth - 6

	if remaining <= 0 {
		return orgField[:cellWidth]
	}

	if len(version) > remaining {
		version = version[:remaining]
	}

	padding := remaining - len(version)
	if padding < 0 {
		padding = 0
	}

	return orgField + strings.Repeat(" ", padding) + version
}
