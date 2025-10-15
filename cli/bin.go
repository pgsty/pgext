/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// BinaryPackage represents a binary package record
type BinaryPackage struct {
	PG       int    // PostgreSQL major version
	OS       string // OS identifier
	Name     string // Package name
	Repo     string // Repository ID
	Ver      string // Full version string
	Version  string // Clean semantic version
	File     string // Package filename
	SHA256   string // SHA256 checksum
	Href     string // Relative URL path
	Size     int64  // Package file size
	SizeFull int64  // Installed size
}

// cleanVersion extracts clean semantic version from version strings
func cleanVersion(version string) string {
	// Step 1: Remove epoch prefix (e.g., "1:2.0.12" → "2.0.12")
	if idx := strings.Index(version, ":"); idx > 0 && idx <= 2 {
		// Check if prefix is numeric (epoch)
		if isNumeric(version[:idx]) {
			version = version[idx+1:]
		}
	}

	// Step 2: Remove special suffixes like .citus, .pgdg, etc. before other processing
	// This handles cases like "12.1.6.citus-1" → "12.1.6-1"
	suffixPatterns := regexp.MustCompile(`\.(citus|pgdg|PGDG|el\d+|rhel\d+)`)
	version = suffixPatterns.ReplaceAllString(version, "")

	// Step 3: Find the first occurrence of delimiters and truncate
	// Priority: look for -, ~, +, _ in that order
	delimiters := []string{"-", "~", "+", "_"}

	minIdx := len(version)
	for _, delim := range delimiters {
		if idx := strings.Index(version, delim); idx > 0 && idx < minIdx {
			// Special handling for underscore followed by date (e.g., "1.2_20240606")
			// We want to keep just "1.2"
			if delim == "_" {
				// Check if what follows looks like a date (8 digits)
				if idx+9 <= len(version) && isDate(version[idx+1:idx+9]) {
					minIdx = idx
					continue
				}
			}
			minIdx = idx
		}
	}

	if minIdx < len(version) {
		version = version[:minIdx]
	}

	// Step 4: Clean up any remaining special patterns
	// Remove +dfsg, +git, etc. if somehow still present
	if idx := strings.IndexAny(version, "~+"); idx > 0 {
		version = version[:idx]
	}

	return version
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

// isDate checks if a string looks like a date (YYYYMMDD format)
func isDate(s string) bool {
	if len(s) != 8 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	// Basic validation: year should be reasonable
	year := s[:4]
	if year < "2000" || year > "2099" {
		return false
	}
	return true
}

// BinGenerator generates binary package records
type BinGenerator struct {
	ctx context.Context
}

// NewBinGenerator creates a new binary package generator
func NewBinGenerator(ctx context.Context) *BinGenerator {
	return &BinGenerator{ctx: ctx}
}

// Generate creates pgext.bin records from apt and dnf tables
func (g *BinGenerator) Generate() (int, error) {
	logrus.Info("Generating binary package table...")

	// Clear existing data
	if err := g.truncateTable(); err != nil {
		return 0, fmt.Errorf("truncate bin table: %w", err)
	}

	// Get active PostgreSQL versions
	pgVersions, err := g.getActivePGVersions()
	if err != nil {
		return 0, fmt.Errorf("get PG versions: %w", err)
	}

	// Collect ALL packages from ALL PG versions first
	var allPackages []BinaryPackage

	for _, pgVer := range pgVersions {
		// Collect packages for this PG version
		packages, err := g.collectPackagesForVersion(pgVer)
		if err != nil {
			logrus.Warnf("Failed to collect packages for PG %d: %v", pgVer, err)
			continue
		}

		if len(packages) == 0 {
			logrus.Debugf("No packages found for PG %d", pgVer)
			continue
		}

		logrus.Infof("Collected %d packages for PG %d", len(packages), pgVer)
		allPackages = append(allPackages, packages...)
	}

	if len(allPackages) == 0 {
		logrus.Warn("No packages found for any PG version")
		return 0, nil
	}

	// Sort ALL packages globally by PG (desc), OS, name, and version (desc)
	// This ensures consistent global ordering
	logrus.Infof("Sorting %d packages globally...", len(allPackages))
	g.sortPackages(allPackages)

	// Insert all packages
	count, err := g.insertPackages(allPackages)
	if err != nil {
		return 0, fmt.Errorf("insert packages: %w", err)
	}

	logrus.Infof("Generated %d binary package records", count)
	return count, nil
}

// truncateTable clears the bin table
func (g *BinGenerator) truncateTable() error {
	_, err := ExecSQLContext(g.ctx, "TRUNCATE TABLE pgext.bin RESTART IDENTITY")
	return err
}

// getActivePGVersions retrieves active PostgreSQL versions
func (g *BinGenerator) getActivePGVersions() ([]int, error) {
	query := "SELECT pg FROM pgext.active_pg ORDER BY pg DESC"

	rows, err := QueryContext(g.ctx, query)
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

// collectPackagesForVersion collects all packages for a specific PG version
func (g *BinGenerator) collectPackagesForVersion(pgVer int) ([]BinaryPackage, error) {
	var packages []BinaryPackage

	// Get DNF packages
	dnfPackages, err := g.collectDNFPackages(pgVer)
	if err != nil {
		logrus.Warnf("Failed to collect DNF packages for PG %d: %v", pgVer, err)
	} else {
		packages = append(packages, dnfPackages...)
	}

	// Get APT packages
	aptPackages, err := g.collectAPTPackages(pgVer)
	if err != nil {
		logrus.Warnf("Failed to collect APT packages for PG %d: %v", pgVer, err)
	} else {
		packages = append(packages, aptPackages...)
	}

	return packages, nil
}

// collectDNFPackages collects DNF packages for a specific PG version
func (g *BinGenerator) collectDNFPackages(pgVer int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT
			r.os,
			d.name,
			d.repo,
			d.version || COALESCE('-' || d.release, ''),
			d.location_href,
			d.pkg_id,
			d.size_package,
			d.size_installed
		FROM pgext.dnf d
		JOIN pgext.repository r ON d.repo = r.id
		WHERE d.name LIKE '%%_%d%%'
			AND d.name NOT LIKE '%%-llvmjit%%'
			AND d.name NOT LIKE '%%-devel%%'
		ORDER BY r.os, d.name, d.version DESC
	`, pgVer)

	rows, err := QueryContext(g.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []BinaryPackage
	for rows.Next() {
		var pkg BinaryPackage
		var locationHref string
		var sizePackage, sizeInstalled interface{}

		err := rows.Scan(
			&pkg.OS,
			&pkg.Name,
			&pkg.Repo,
			&pkg.Ver,
			&locationHref,
			&pkg.SHA256,
			&sizePackage,
			&sizeInstalled,
		)
		if err != nil {
			logrus.Debugf("Failed to scan DNF package: %v", err)
			continue
		}

		pkg.PG = pgVer

		// Clean the version string
		pkg.Version = cleanVersion(pkg.Ver)

		// Extract filename from location href
		if idx := strings.LastIndex(locationHref, "/"); idx >= 0 {
			pkg.File = locationHref[idx+1:]
		} else {
			pkg.File = locationHref
		}
		pkg.Href = locationHref

		// Handle nullable size values (PostgreSQL returns int32 for INTEGER columns)
		switch v := sizePackage.(type) {
		case int64:
			pkg.Size = v
		case int32:
			pkg.Size = int64(v)
		case int:
			pkg.Size = int64(v)
		}

		switch v := sizeInstalled.(type) {
		case int64:
			pkg.SizeFull = v
		case int32:
			pkg.SizeFull = int64(v)
		case int:
			pkg.SizeFull = int64(v)
		}

		packages = append(packages, pkg)
	}

	return packages, nil
}

// collectAPTPackages collects APT packages for a specific PG version
func (g *BinGenerator) collectAPTPackages(pgVer int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT
			r.os,
			a.package,
			a.repo,
			a.version,
			a.filename,
			a.sha256,
			a.size,
			a.size_install
		FROM pgext.apt a
		JOIN pgext.repository r ON a.repo = r.id
		WHERE a.package LIKE '%%-%d%%'
			AND a.package NOT LIKE '%%-dbgsym%%'
		ORDER BY r.os, a.package, a.version DESC
	`, pgVer)

	rows, err := QueryContext(g.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []BinaryPackage
	for rows.Next() {
		var pkg BinaryPackage
		var filename string
		var size, sizeInstall interface{}

		err := rows.Scan(
			&pkg.OS,
			&pkg.Name,
			&pkg.Repo,
			&pkg.Ver,
			&filename,
			&pkg.SHA256,
			&size,
			&sizeInstall,
		)
		if err != nil {
			logrus.Debugf("Failed to scan APT package: %v", err)
			continue
		}

		pkg.PG = pgVer

		// Clean the version string
		pkg.Version = cleanVersion(pkg.Ver)

		// Extract filename from path
		if idx := strings.LastIndex(filename, "/"); idx >= 0 {
			pkg.File = filename[idx+1:]
			pkg.Href = filename
		} else {
			pkg.File = filename
			pkg.Href = filename
		}

		// Handle size values (PostgreSQL returns int32 for INTEGER columns)
		switch v := size.(type) {
		case int64:
			pkg.Size = v
		case int32:
			pkg.Size = int64(v)
		case int:
			pkg.Size = int64(v)
		}

		switch v := sizeInstall.(type) {
		case int64:
			pkg.SizeFull = v
		case int32:
			pkg.SizeFull = int64(v)
		case int:
			pkg.SizeFull = int64(v)
		}

		packages = append(packages, pkg)
	}

	return packages, nil
}

// sortPackages sorts packages by OS, name, and version (descending)
func (g *BinGenerator) sortPackages(packages []BinaryPackage) {
	sort.Slice(packages, func(i, j int) bool {
		// Sort by PG version (descending)
		if packages[i].PG != packages[j].PG {
			return packages[i].PG > packages[j].PG
		}

		// Sort by OS
		if packages[i].OS != packages[j].OS {
			return packages[i].OS < packages[j].OS
		}

		// Sort by package name
		if packages[i].Name != packages[j].Name {
			return packages[i].Name < packages[j].Name
		}

		// Sort by version (descending) - newer versions first
		// This ensures that for the same pg/os/name, higher versions appear first
		return g.compareVersions(packages[i], packages[j]) > 0
	})
}

// compareVersions compares two package versions
// Returns: positive if a > b (a should come first), negative if a < b, 0 if equal
func (g *BinGenerator) compareVersions(a, b BinaryPackage) int {
	// Step 1: Compare clean semantic versions first
	cmp := compareSemanticVersions(a.Version, b.Version)
	if cmp != 0 {
		return cmp
	}

	// Step 2: If semantic versions are equal, try to compare release numbers
	aRelease := extractReleaseNumber(a.Ver)
	bRelease := extractReleaseNumber(b.Ver)

	// If both have numeric release numbers, compare them
	if aRelease > 0 && bRelease > 0 {
		if aRelease > bRelease {
			return 1
		} else if aRelease < bRelease {
			return -1
		}
		// If release numbers are also equal, fall through to string comparison
	}

	// Step 3: Fall back to full string comparison
	// This ensures consistent ordering when semantic versions match
	return strings.Compare(a.Ver, b.Ver)
}

// extractReleaseNumber extracts the numeric release number from a version string
// For example: "13.2.0-9PIGSTY.el8" returns 9, "13.2.0-11PIGSTY.el9" returns 11
func extractReleaseNumber(ver string) int {
	// First, skip any special prefixes like epoch (1:version)
	startIdx := 0
	if idx := strings.Index(ver, ":"); idx > 0 && idx <= 2 {
		startIdx = idx + 1
	}

	// Look for the first dash or tilde that indicates a release
	// Skip underscores as they typically indicate dates or special builds
	searchStr := ver[startIdx:]

	// Find the position of version-release separator
	separatorIdx := -1
	for i, ch := range searchStr {
		if ch == '-' || ch == '~' {
			separatorIdx = i
			break
		} else if ch == '_' || ch == '+' {
			// These typically indicate special builds, not release numbers
			return 0
		}
	}

	if separatorIdx > 0 && separatorIdx+1 < len(searchStr) {
		releaseStr := searchStr[separatorIdx+1:]
		// Extract the numeric part at the beginning
		numStr := ""
		for _, ch := range releaseStr {
			if ch >= '0' && ch <= '9' {
				numStr += string(ch)
			} else {
				break
			}
		}

		if numStr != "" {
			if val, err := strconv.Atoi(numStr); err == nil {
				return val
			}
		}
	}
	return 0
}

// compareSemanticVersions compares two semantic version strings
// Returns: 1 if a > b, -1 if a < b, 0 if a == b
func compareSemanticVersions(a, b string) int {
	// Handle empty strings
	if a == "" && b == "" {
		return 0
	}
	if a == "" {
		return -1
	}
	if b == "" {
		return 1
	}

	// Split versions into components
	aParts := splitVersion(a)
	bParts := splitVersion(b)

	// Compare each numeric component
	maxLen := len(aParts)
	if len(bParts) > maxLen {
		maxLen = len(bParts)
	}

	for i := 0; i < maxLen; i++ {
		var aVal, bVal int
		if i < len(aParts) {
			aVal = aParts[i]
		}
		if i < len(bParts) {
			bVal = bParts[i]
		}

		if aVal > bVal {
			return 1
		} else if aVal < bVal {
			return -1
		}
	}

	// If all components are equal up to maxLen,
	// the version with more components is considered greater
	// unless the extra components are zeros
	if len(aParts) > len(bParts) {
		// Check if remaining components are non-zero
		for i := len(bParts); i < len(aParts); i++ {
			if aParts[i] > 0 {
				return 1
			}
		}
	} else if len(bParts) > len(aParts) {
		// Check if remaining components are non-zero
		for i := len(aParts); i < len(bParts); i++ {
			if bParts[i] > 0 {
				return -1
			}
		}
	}

	return 0
}

// splitVersion splits a version string into numeric components
func splitVersion(version string) []int {
	// Handle special cases like "18rc1" by extracting numeric prefix
	if matches := regexp.MustCompile(`^(\d+)([a-zA-Z])`).FindStringSubmatch(version); len(matches) > 1 {
		version = matches[1]
	}

	// Split by dots and parse each part
	parts := strings.Split(version, ".")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		// Extract numeric portion from each part
		numStr := ""
		for _, ch := range part {
			if ch >= '0' && ch <= '9' {
				numStr += string(ch)
			} else {
				break
			}
		}

		if numStr != "" {
			if val, err := strconv.Atoi(numStr); err == nil {
				result = append(result, val)
			}
		}
	}

	return result
}

// insertPackages inserts packages into the bin table
func (g *BinGenerator) insertPackages(packages []BinaryPackage) (int, error) {
	if len(packages) == 0 {
		return 0, nil
	}

	tx, err := Begin(g.ctx)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	batch := &pgx.Batch{}
	count := 0

	for _, pkg := range packages {
		batch.Queue(`
			INSERT INTO pgext.bin (
				pg, os, name, repo, ver, version,
				file, sha256, href, size, size_full
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`,
			pkg.PG, pkg.OS, pkg.Name, pkg.Repo, pkg.Ver, pkg.Version,
			pkg.File, pkg.SHA256, pkg.Href, pkg.Size, pkg.SizeFull,
		)
		count++

		// Send batch every 1000 records
		if batch.Len() >= 1000 {
			br := tx.SendBatch(g.ctx, batch)
			if err := br.Close(); err != nil {
				return count, fmt.Errorf("send batch: %w", err)
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 {
		br := tx.SendBatch(g.ctx, batch)
		if err := br.Close(); err != nil {
			return count, fmt.Errorf("send final batch: %w", err)
		}
	}

	if err := tx.Commit(g.ctx); err != nil {
		return count, fmt.Errorf("commit transaction: %w", err)
	}

	return count, nil
}