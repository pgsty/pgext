/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

// This file combines binary package generation (bin.go) and recap (recap.go)
// for package availability matrix generation.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

//============================================================================
// Binary Package Generation (from bin.go)
//============================================================================

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
// Only keeps digits and dots, stops at first non-digit/non-dot character.
// Examples:
//   - "0.2.0+git20211101.d7d10f2" -> "0.2.0"
//   - "5.5.0+debpgdg" -> "5.5.0"
//   - "1.2_20240606" -> "1.2"
//   - "2.22.0+dfsg" -> "2.22.0"
//   - "1.3.0~alpha" -> "1.3.0"
//   - "1:2.0.12-3.el8" -> "2.0.12"
func cleanVersion(ver string) string {
	// Remove epoch prefix (e.g., "1:2.0.12" → "2.0.12")
	if idx := strings.Index(ver, ":"); idx > 0 && idx <= 2 {
		ver = ver[idx+1:]
	}

	// Extract version: keep only digits and dots
	// Stop at first character that's not a digit or dot
	var result strings.Builder
	for _, ch := range ver {
		if (ch >= '0' && ch <= '9') || ch == '.' {
			result.WriteRune(ch)
		} else {
			// Found non-version character, stop here
			break
		}
	}

	return strings.TrimRight(result.String(), ".")
}

// BinGenerator generates pgext.bin records from apt and dnf tables.
type BinGenerator struct {
	ctx      context.Context
	dnfTable string
	aptTable string
	binTable string
}

// NewBinGenerator creates a new binary package generator.
func NewBinGenerator(ctx context.Context) *BinGenerator {
	return newBinGenerator(ctx, liveTable("dnf"), liveTable("apt"), liveTable("bin"))
}

func newBinGenerator(ctx context.Context, dnfTable, aptTable, binTable string) *BinGenerator {
	return &BinGenerator{
		ctx:      ctx,
		dnfTable: dnfTable,
		aptTable: aptTable,
		binTable: binTable,
	}
}

// Generate creates pgext.bin records from apt and dnf tables.
func (g *BinGenerator) Generate() (int, error) {
	logrus.Info("generating binary package table...")

	// Clear existing data
	_, err := ExecSQLContext(g.ctx, "TRUNCATE TABLE "+g.binTable+" RESTART IDENTITY")
	if err != nil {
		return 0, fmt.Errorf("truncate bin table: %w", err)
	}

	// Get active PostgreSQL versions
	pgVersions, err := g.getActivePGVersions()
	if err != nil {
		return 0, fmt.Errorf("get active PG versions: %w", err)
	}
	if len(pgVersions) == 0 {
		return 0, fmt.Errorf("no active PostgreSQL versions")
	}

	// Collect and insert packages for each PG version
	total := 0
	report := RunReport{Operation: "binary package generation", Total: len(pgVersions) * 2}
	for _, pg := range pgVersions {
		count, succeeded, failures, err := g.processVersion(pg)
		if err != nil {
			return total, fmt.Errorf("process PG %d: %w", pg, err)
		}
		report.Succeeded += succeeded
		for _, failure := range failures {
			report.AddFailure(failure.Item, failure.Err)
		}
		logrus.Infof("  PG %d: %d packages", pg, count)
		total += count
	}
	// Repository-level best-effort policy is resolved before this phase. A
	// database query or scan failure while normalizing accepted repository data
	// is always fatal because publishing one package-manager half would corrupt
	// the catalog.
	if err := report.Err(false); err != nil {
		return total, err
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
func (g *BinGenerator) processVersion(pg int) (int, int, []RunFailure, error) {
	var packages []BinaryPackage
	var failures []RunFailure
	succeeded := 0

	// Collect DNF packages
	dnf, err := g.collectDNF(pg)
	if err != nil {
		logrus.Warnf("DNF collection failed for PG %d: %v", pg, err)
		failures = append(failures, RunFailure{Item: fmt.Sprintf("PG %d DNF", pg), Err: err})
	} else {
		succeeded++
		packages = append(packages, dnf...)
	}

	// Collect APT packages
	apt, err := g.collectAPT(pg)
	if err != nil {
		logrus.Warnf("APT collection failed for PG %d: %v", pg, err)
		failures = append(failures, RunFailure{Item: fmt.Sprintf("PG %d APT", pg), Err: err})
	} else {
		succeeded++
		packages = append(packages, apt...)
	}

	// Insert packages
	count, err := g.insertPackages(packages)
	return count, succeeded, failures, err
}

// collectDNF collects DNF packages for a specific PG version.
func (g *BinGenerator) collectDNF(pg int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT r.os, d.name, d.repo,
		       d.version || COALESCE('-' || d.release, ''),
		       d.location_href, d.pkg_id,
		       d.size_package, d.size_installed
		FROM %s d
		JOIN pgext.repository r ON d.repo = r.id
		WHERE d.name LIKE '%%_%d%%'
		  AND d.name NOT LIKE '%%-llvmjit%%'
		  AND d.name NOT LIKE '%%-devel%%'
	`, g.dnfTable, pg)

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
			return nil, fmt.Errorf("scan DNF package: %w", err)
		}

		pkg.PG = pg
		pkg.Version = cleanVersion(pkg.Ver)
		pkg.Href = href
		pkg.File = extractFilename(href)
		pkg.Size = toInt64(sizePackage)
		pkg.SizeFull = toInt64(sizeInstalled)

		packages = append(packages, pkg)
	}
	return packages, rows.Err()
}

// collectAPT collects APT packages for a specific PG version.
func (g *BinGenerator) collectAPT(pg int) ([]BinaryPackage, error) {
	query := fmt.Sprintf(`
		SELECT r.os, a.package, a.repo, a.version,
		       a.filename, a.sha256, a.size, a.size_install
		FROM %s a
		JOIN pgext.repository r ON a.repo = r.id
		WHERE (a.package LIKE '%%-%d%%' OR a.package = 'pgagent')
		  AND a.package NOT LIKE '%%-dbgsym%%'
	`, g.aptTable, pg)

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
			return nil, fmt.Errorf("scan APT package: %w", err)
		}

		pkg.PG = pg
		pkg.Version = cleanVersion(pkg.Ver)
		pkg.File = extractFilename(filename)
		pkg.Href = filename
		pkg.Size = toInt64(size)
		pkg.SizeFull = toInt64(sizeInstall)

		packages = append(packages, pkg)
	}
	return packages, rows.Err()
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
		batch.Queue(fmt.Sprintf(`
			INSERT INTO %s (
				pg, os, name, repo, ver, version,
				file, sha256, href, size, size_full
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`, g.binTable), pkg.PG, pkg.OS, pkg.Name, pkg.Repo, pkg.Ver, pkg.Version,
			pkg.File, pkg.SHA256, pkg.Href, pkg.Size, pkg.SizeFull)

		// Send batch every 1000 records
		if batch.Len() >= 1000 {
			br := tx.SendBatch(g.ctx, batch)
			if err := br.Close(); err != nil {
				return 0, fmt.Errorf("insert binary package batch: %w", err)
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 {
		br := tx.SendBatch(g.ctx, batch)
		if err := br.Close(); err != nil {
			return 0, fmt.Errorf("insert final binary package batch: %w", err)
		}
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

//============================================================================
// Package Availability Matrix Generation (from recap.go)
//============================================================================

// Repository represents a repository with its metadata
type Repository struct {
	ID         string
	Type       string // rpm or deb
	OS         string
	Org        string
	OSCode     string
	OSArch     string
	DefaultURL string
	MirrorURL  string
}

// Package represents a package entry
type Package struct {
	PG        int
	OS        string
	PName     string // normalized package name
	Org       string
	Type      string
	OSCode    string
	OSArch    string
	Repo      string
	Name      string // original package name
	Ver       string // full version string
	Version   string // semantic version
	Release   string
	File      string
	SHA256    string
	URL       string
	MirrorURL string
	Size      int64
	SizeFull  int64
}

// RecapMatrix generates pgext.pkg availability matrix using SQL procedures
func RecapMatrix() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	return RecapMatrixContext(ctx)
}

// RecapMatrixContext builds pgext.pkg in a run-scoped staging table and
// atomically publishes the finished matrix.
func RecapMatrixContext(ctx context.Context) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	logrus.Info("generate pgext.pkg from repo package data")

	// Get active PostgreSQL versions
	var activePG []int
	rows, err := QueryContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg DESC")
	if err != nil {
		return fmt.Errorf("failed to query active pg versions: %w", err)
	}
	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err != nil {
			rows.Close()
			return fmt.Errorf("scan active pg version: %w", err)
		}
		activePG = append(activePG, pg)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return fmt.Errorf("iterate active pg versions: %w", err)
	}
	rows.Close()

	// Get active OS versions
	var activeOS []string
	rows, err = QueryContext(ctx, "SELECT os FROM pgext.active_os ORDER BY os")
	if err != nil {
		return fmt.Errorf("failed to query active os versions: %w", err)
	}
	for rows.Next() {
		var os string
		if err := rows.Scan(&os); err != nil {
			rows.Close()
			return fmt.Errorf("scan active os: %w", err)
		}
		activeOS = append(activeOS, os)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return fmt.Errorf("iterate active os versions: %w", err)
	}
	rows.Close()

	// Format and log active versions
	pgVersions := make([]string, len(activePG))
	for i, pg := range activePG {
		pgVersions[i] = fmt.Sprintf("%d", pg)
	}
	logrus.Infof("active pg: %s", strings.Join(pgVersions, ", "))
	logrus.Infof("active os: %s", strings.Join(activeOS, ", "))
	if len(activePG) == 0 {
		return fmt.Errorf("no active PostgreSQL versions; refusing to replace pgext.pkg")
	}
	if len(activeOS) == 0 {
		return fmt.Errorf("no active operating systems; refusing to replace pgext.pkg")
	}

	stage, err := newPackageStaging(ctx, false, true)
	if err != nil {
		return fmt.Errorf("create pkg staging table: %w", err)
	}
	defer stage.Drop()

	pkgCount, err := stage.rebuildAndPublishPkgFromLive(ctx)
	if err != nil {
		return fmt.Errorf("rebuild package matrix: %w", err)
	}
	logrus.Infof("recap completed: %d package records in pgext.pkg", pkgCount)
	return nil
}

// RecapPackages is the main entry point for the recap operation
func RecapPackages() error {
	return RecapMatrix()
}

// containsSQL checks if the content has actual SQL statements (not just comments)
func containsSQL(content string) bool {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Skip empty lines and comment lines
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			continue
		}
		// Found a non-comment line
		return true
	}
	return false
}
