/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// BinaryPackage represents a binary package record.
type BinaryPackage struct {
	PG       int
	OS       string
	Name     string
	Repo     string
	Ver      string
	Version  string
	File     string
	SHA256   string
	Href     string
	Size     int64
	SizeFull int64
}

// cleanVersion extracts clean semantic version from version strings.
// Examples: "1:2.0.12-3.el8" -> "2.0.12", "1.2.3.citus-1" -> "1.2.3"
func cleanVersion(ver string) string {
	// Remove epoch prefix (e.g., "1:2.0.12" → "2.0.12")
	if idx := strings.Index(ver, ":"); idx > 0 && idx <= 2 {
		ver = ver[idx+1:]
	}

	// Remove special suffixes (.citus, .pgdg, .el8, etc.)
	suffixes := regexp.MustCompile(`\.(citus|pgdg|PGDG|el\d+|rhel\d+)`)
	ver = suffixes.ReplaceAllString(ver, "")

	// Find first delimiter (-, ~, +, _) and truncate
	for _, delim := range []string{"-", "~", "+", "_"} {
		if idx := strings.Index(ver, delim); idx > 0 {
			return ver[:idx]
		}
	}

	return ver
}

// BinGenerator generates pgext.bin records from apt and dnf tables.
type BinGenerator struct {
	ctx context.Context
}

// NewBinGenerator creates a new binary package generator.
func NewBinGenerator(ctx context.Context) *BinGenerator {
	return &BinGenerator{ctx: ctx}
}

// Generate creates pgext.bin records from apt and dnf tables.
func (g *BinGenerator) Generate() (int, error) {
	logrus.Info("generating binary package table...")

	// Clear existing data
	_, err := ExecSQLContext(g.ctx, "TRUNCATE TABLE pgext.bin RESTART IDENTITY")
	if err != nil {
		return 0, fmt.Errorf("truncate bin table: %w", err)
	}

	// Get active PostgreSQL versions
	pgVersions, err := g.getActivePGVersions()
	if err != nil {
		return 0, fmt.Errorf("get active PG versions: %w", err)
	}

	// Collect and insert packages for each PG version
	total := 0
	for _, pg := range pgVersions {
		count, err := g.processVersion(pg)
		if err != nil {
			logrus.Warnf("failed to process PG %d: %v", pg, err)
			continue
		}
		logrus.Infof("  PG %d: %d packages", pg, count)
		total += count
	}

	logrus.Infof("generated %d binary package records", total)
	return total, nil
}

// getActivePGVersions retrieves active PostgreSQL versions in descending order.
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

// processVersion collects and inserts packages for a specific PG version.
func (g *BinGenerator) processVersion(pg int) (int, error) {
	var packages []BinaryPackage

	// Collect DNF packages
	dnf, err := g.collectDNF(pg)
	if err != nil {
		logrus.Warnf("DNF collection failed for PG %d: %v", pg, err)
	} else {
		packages = append(packages, dnf...)
	}

	// Collect APT packages
	apt, err := g.collectAPT(pg)
	if err != nil {
		logrus.Warnf("APT collection failed for PG %d: %v", pg, err)
	} else {
		packages = append(packages, apt...)
	}

	// Insert packages
	return g.insertPackages(packages)
}

// collectDNF collects DNF packages for a specific PG version.
func (g *BinGenerator) collectDNF(pg int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT r.os, d.name, d.repo,
		       d.version || COALESCE('-' || d.release, ''),
		       d.location_href, d.pkg_id,
		       d.size_package, d.size_installed
		FROM pgext.dnf d
		JOIN pgext.repository r ON d.repo = r.id
		WHERE d.name LIKE '%%_%d%%'
		  AND d.name NOT LIKE '%%-llvmjit%%'
		  AND d.name NOT LIKE '%%-devel%%'
	`, pg)

	rows, err := QueryContext(g.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []BinaryPackage
	for rows.Next() {
		var pkg BinaryPackage
		var href string
		var sizePackage, sizeInstalled interface{}

		err := rows.Scan(&pkg.OS, &pkg.Name, &pkg.Repo, &pkg.Ver,
			&href, &pkg.SHA256, &sizePackage, &sizeInstalled)
		if err != nil {
			continue
		}

		pkg.PG = pg
		pkg.Version = cleanVersion(pkg.Ver)
		pkg.Href = href
		pkg.File = extractFilename(href)
		pkg.Size = toInt64(sizePackage)
		pkg.SizeFull = toInt64(sizeInstalled)

		packages = append(packages, pkg)
	}
	return packages, nil
}

// collectAPT collects APT packages for a specific PG version.
func (g *BinGenerator) collectAPT(pg int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT r.os, a.package, a.repo, a.version,
		       a.filename, a.sha256, a.size, a.size_install
		FROM pgext.apt a
		JOIN pgext.repository r ON a.repo = r.id
		WHERE a.package LIKE '%%-%d%%'
		  AND a.package NOT LIKE '%%-dbgsym%%'
	`, pg)

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

		err := rows.Scan(&pkg.OS, &pkg.Name, &pkg.Repo, &pkg.Ver,
			&filename, &pkg.SHA256, &size, &sizeInstall)
		if err != nil {
			continue
		}

		pkg.PG = pg
		pkg.Version = cleanVersion(pkg.Ver)
		pkg.File = extractFilename(filename)
		pkg.Href = filename
		pkg.Size = toInt64(size)
		pkg.SizeFull = toInt64(sizeInstall)

		packages = append(packages, pkg)
	}
	return packages, nil
}

// insertPackages inserts packages into pgext.bin using batch operations.
func (g *BinGenerator) insertPackages(packages []BinaryPackage) (int, error) {
	if len(packages) == 0 {
		return 0, nil
	}

	tx, err := Begin(g.ctx)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(g.ctx)

	batch := &pgx.Batch{}
	for _, pkg := range packages {
		batch.Queue(`
			INSERT INTO pgext.bin (
				pg, os, name, repo, ver, version,
				file, sha256, href, size, size_full
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`, pkg.PG, pkg.OS, pkg.Name, pkg.Repo, pkg.Ver, pkg.Version,
			pkg.File, pkg.SHA256, pkg.Href, pkg.Size, pkg.SizeFull)

		// Send batch every 1000 records
		if batch.Len() >= 1000 {
			br := tx.SendBatch(g.ctx, batch)
			br.Close()
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 {
		br := tx.SendBatch(g.ctx, batch)
		br.Close()
	}

	if err := tx.Commit(g.ctx); err != nil {
		return 0, fmt.Errorf("commit: %w", err)
	}

	return len(packages), nil
}

// extractFilename extracts filename from path/URL.
func extractFilename(path string) string {
	if idx := strings.LastIndex(path, "/"); idx >= 0 {
		return path[idx+1:]
	}
	return path
}

// toInt64 converts interface{} to int64 (handles int32/int/int64).
func toInt64(v interface{}) int64 {
	switch val := v.(type) {
	case int64:
		return val
	case int32:
		return int64(val)
	case int:
		return int64(val)
	default:
		return 0
	}
}
