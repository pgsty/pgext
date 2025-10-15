/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// BinPackage represents a binary package record for display.
type BinPackage struct {
	Name    string
	OS      string
	Org     string
	Size    string
	Version string
	URL     string
}

// ShowBin displays binary packages from pgext.bin table.
func ShowBin(name string, pgVer int, osFilter string, region string) error {
	ctx := context.Background()
	if name == "" {
		return fmt.Errorf("package or extension name is required")
	}

	// Choose URL field based on region
	urlField := "r.default_url"
	if region == "china" || region == "mirror" {
		urlField = "COALESCE(r.mirror_url, r.default_url)"
	}

	// Build query
	query := fmt.Sprintf(`
		SELECT b.name, p.os, r.org, b.size, b.version, format('%%s/%%s', %s, b.href) AS url
		FROM pgext.pkg p
		JOIN pgext.bin b USING (pg, os, name)
		JOIN pgext.repository r ON b.repo = r.id
		WHERE (p.pkg = $1 OR p.ext = $1)
	`, urlField)

	args := []interface{}{name}
	argCount := 2

	if pgVer > 0 {
		query += fmt.Sprintf(" AND p.pg = $%d", argCount)
		args = append(args, pgVer)
		argCount++
	}
	if osFilter != "" {
		query += fmt.Sprintf(" AND p.os LIKE $%d", argCount)
		args = append(args, "%"+osFilter+"%")
		argCount++
	}

	query += " ORDER BY p.pg DESC, p.os, b.ver::pgext.version DESC"

	// Execute query
	rows, err := QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Collect results
	var packages []BinPackage
	for rows.Next() {
		var pkg BinPackage
		var size int64
		err := rows.Scan(&pkg.Name, &pkg.OS, &pkg.Org, &size, &pkg.Version, &pkg.URL)
		if err != nil {
			logrus.Debugf("scan error: %v", err)
			continue
		}
		pkg.Size = formatSize(size)
		packages = append(packages, pkg)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration failed: %w", err)
	}

	// Display results
	displayBinTable(packages)

	return nil
}

// formatSize formats byte size to human-readable format with 2 significant digits.
func formatSize(bytes int64) string {
	if bytes < 0 {
		return "0"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	units := []string{"KiB", "MiB", "GiB", "TiB"}
	size := float64(bytes) / unit
	unitIdx := 0

	for size >= unit && unitIdx < len(units)-1 {
		size /= unit
		unitIdx++
	}

	// Format: 100+ = integer, 1-99.9 = one decimal
	if size >= 100 {
		return fmt.Sprintf("%.0f %s", size, units[unitIdx])
	}
	return fmt.Sprintf("%.1f %s", size, units[unitIdx])
}

// displayBinTable displays binary packages in a table format.
func displayBinTable(packages []BinPackage) {
	if len(packages) == 0 {
		fmt.Println("\nNo packages found.\n")
		return
	}

	// Calculate column widths
	widths := map[string]int{
		"name":    4,
		"os":      2,
		"org":     3,
		"size":    4,
		"version": 7,
	}

	for _, pkg := range packages {
		if len(pkg.Name) > widths["name"] {
			widths["name"] = len(pkg.Name)
		}
		if len(pkg.OS) > widths["os"] {
			widths["os"] = len(pkg.OS)
		}
		if len(pkg.Org) > widths["org"] {
			widths["org"] = len(pkg.Org)
		}
		if len(pkg.Size) > widths["size"] {
			widths["size"] = len(pkg.Size)
		}
		if len(pkg.Version) > widths["version"] {
			widths["version"] = len(pkg.Version)
		}
	}

	// Print header
	fmt.Println()
	fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n",
		widths["name"], "NAME",
		widths["os"], "OS",
		widths["org"], "ORG",
		widths["size"], "SIZE",
		widths["version"], "VER", "URL")

	// Print separator
	totalWidth := widths["name"] + widths["os"] + widths["org"] + widths["size"] + widths["version"] + 12
	fmt.Println(strings.Repeat("─", totalWidth))

	// Print rows
	for _, pkg := range packages {
		fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n",
			widths["name"], pkg.Name,
			widths["os"], pkg.OS,
			widths["org"], pkg.Org,
			widths["size"], pkg.Size,
			widths["version"], pkg.Version,
			pkg.URL)
	}

	// Print footer
	fmt.Println(strings.Repeat("─", totalWidth))
	fmt.Printf("Total: %d packages\n\n", len(packages))
}


// ShowExt displays extension information from pgext.extension table.
func ShowExt(extName string) error {
	ctx := context.Background()

	query := `
		SELECT name, pkg, category, repo, lang, license, ver,
		       description, requires, website, doc_url, src_url,
		       created_at, updated_at
		FROM pgext.extension
		WHERE name = $1
	`

	var name, pkg, category, repo, lang, license, ver, description string
	var requires, website, docURL, srcURL string
	var createdAt, updatedAt interface{}

	err := QueryRowContext(ctx, query, extName).Scan(
		&name, &pkg, &category, &repo, &lang, &license, &ver,
		&description, &requires, &website, &docURL, &srcURL,
		&createdAt, &updatedAt,
	)
	if err != nil {
		return fmt.Errorf("extension not found: %w", err)
	}

	// Print extension info
	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║        Extension Information              ║")
	fmt.Println("╚═══════════════════════════════════════════╝\n")

	fmt.Printf("Name:        %s\n", name)
	fmt.Printf("Package:     %s\n", pkg)
	fmt.Printf("Category:    %s\n", category)
	fmt.Printf("Repository:  %s\n", repo)
	fmt.Printf("Language:    %s\n", lang)
	fmt.Printf("License:     %s\n", license)
	fmt.Printf("Version:     %s\n", ver)
	fmt.Printf("\nDescription:\n%s\n", wrapText(description, 60))

	if requires != "" {
		fmt.Printf("\nRequires:    %s\n", requires)
	}
	if website != "" {
		fmt.Printf("Website:     %s\n", website)
	}
	if docURL != "" {
		fmt.Printf("Docs:        %s\n", docURL)
	}
	if srcURL != "" {
		fmt.Printf("Source:      %s\n", srcURL)
	}

	fmt.Println()

	return nil
}

// truncate truncates a string to specified length.
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// wrapText wraps text to specified width.
func wrapText(text string, width int) string {
	if len(text) <= width {
		return "  " + text
	}

	var result strings.Builder
	words := strings.Fields(text)
	line := "  "

	for _, word := range words {
		if len(line)+len(word)+1 > width {
			result.WriteString(line)
			result.WriteString("\n")
			line = "  " + word
		} else {
			if line != "  " {
				line += " "
			}
			line += word
		}
	}

	if line != "  " {
		result.WriteString(line)
	}

	return result.String()
}
