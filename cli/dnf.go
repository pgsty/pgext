/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// DNFParser handles parsing of DNF/YUM repository metadata
type DNFParser struct {
	ctx context.Context
	tx  pgx.Tx
}

// NewDNFParser creates a new DNF parser
func NewDNFParser(ctx context.Context) *DNFParser {
	return &DNFParser{ctx: ctx}
}

// ParseRepository parses a single DNF/YUM repository
func (p *DNFParser) ParseRepository(repoID string, data []byte) (int, error) {
	// Create temporary SQLite file
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("dnf-%s-*.db", repoID))
	if err != nil {
		return 0, fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	// Write SQLite data to temp file
	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return 0, fmt.Errorf("write temp file: %w", err)
	}
	tmpFile.Close()

	// Open SQLite database
	sqliteDB, err := sql.Open("sqlite3", tmpPath+"?mode=ro")
	if err != nil {
		return 0, fmt.Errorf("open sqlite: %w", err)
	}
	defer sqliteDB.Close()

	// Extract packages from SQLite
	packages, err := p.extractPackages(sqliteDB)
	if err != nil {
		return 0, fmt.Errorf("extract packages: %w", err)
	}

	if len(packages) == 0 {
		return 0, nil
	}

	// Begin transaction
	tx, err := Begin(p.ctx)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	// Store transaction for use in insertPackages
	p.tx = tx

	// Insert packages
	count := p.insertPackages(repoID, packages)

	// Commit transaction
	if err := tx.Commit(p.ctx); err != nil {
		// Transaction failed, rollback is automatic
		p.tx = nil
		return count, fmt.Errorf("commit transaction: %w", err)
	}

	p.tx = nil

	return count, nil
}

// extractPackages reads package data from SQLite database
func (p *DNFParser) extractPackages(db *sql.DB) ([]DNFPackage, error) {
	query := `
		SELECT
			pkgKey, pkgId, name, arch, version, epoch, release,
			summary, description, url, time_file, time_build,
			rpm_license, rpm_vendor, rpm_group, rpm_buildhost,
			rpm_sourcerpm, rpm_header_start, rpm_header_end,
			rpm_packager, size_package, size_installed, size_archive,
			location_href, location_base, checksum_type
		FROM packages
		ORDER BY name, version DESC
	`

	rows, err := db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query packages: %w", err)
	}
	defer rows.Close()

	var packages []DNFPackage

	for rows.Next() {
		var pkg DNFPackage
		err := rows.Scan(
			&pkg.PkgKey, &pkg.PkgId, &pkg.Name, &pkg.Arch, &pkg.Version,
			&pkg.Epoch, &pkg.Release, &pkg.Summary, &pkg.Description, &pkg.URL,
			&pkg.TimeFile, &pkg.TimeBuild, &pkg.RPMLicense, &pkg.RPMVendor,
			&pkg.RPMGroup, &pkg.RPMBuildHost, &pkg.RPMSourceRPM, &pkg.RPMHeaderStart,
			&pkg.RPMHeaderEnd, &pkg.RPMPackager, &pkg.SizePackage, &pkg.SizeInstalled,
			&pkg.SizeArchive, &pkg.LocationHref, &pkg.LocationBase, &pkg.ChecksumType,
		)
		if err != nil {
			logrus.Debugf("failed to scan package: %v", err)
			continue
		}

		packages = append(packages, pkg)
	}

	return packages, rows.Err()
}

// insertPackages inserts DNF packages into the database
func (p *DNFParser) insertPackages(repoID string, packages []DNFPackage) int {
	if len(packages) == 0 {
		return 0
	}

	batch := &pgx.Batch{}
	count := 0

	for _, pkg := range packages {
		// Convert SQL null types to interface{} for pgx
		batch.Queue(`
			INSERT INTO pgext.dnf (
				repo, pkg_key, pkg_id, name, arch, version, epoch, release,
				summary, description, url, time_file, time_build,
				rpm_license, rpm_vendor, rpm_group, rpm_buildhost, rpm_sourcerpm,
				rpm_header_start, rpm_header_end, rpm_packager,
				size_package, size_installed, size_archive,
				location_href, location_base, checksum_type
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
				$21, $22, $23, $24, $25, $26, $27
			)`,
			repoID, pkg.PkgKey, pkg.PkgId, pkg.Name,
			nullStr(pkg.Arch), nullStr(pkg.Version), nullStr(pkg.Epoch), nullStr(pkg.Release),
			nullStr(pkg.Summary), nullStr(pkg.Description), nullStr(pkg.URL),
			nullInt(pkg.TimeFile), nullInt(pkg.TimeBuild),
			nullStr(pkg.RPMLicense), nullStr(pkg.RPMVendor), nullStr(pkg.RPMGroup),
			nullStr(pkg.RPMBuildHost), nullStr(pkg.RPMSourceRPM),
			nullInt(pkg.RPMHeaderStart), nullInt(pkg.RPMHeaderEnd), nullStr(pkg.RPMPackager),
			nullInt(pkg.SizePackage), nullInt(pkg.SizeInstalled), nullInt(pkg.SizeArchive),
			nullStr(pkg.LocationHref), nullStr(pkg.LocationBase), nullStr(pkg.ChecksumType))

		count++

		// Send batch every 1000 records
		if batch.Len() >= 1000 {
			if p.tx != nil {
				br := p.tx.SendBatch(p.ctx, batch)
				if err := br.Close(); err != nil {
					logrus.Warnf("failed to send batch: %v", err)
				}
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 && p.tx != nil {
		br := p.tx.SendBatch(p.ctx, batch)
		if err := br.Close(); err != nil {
			logrus.Warnf("failed to send final batch: %v", err)
		}
	}

	return count
}

// DNFPackage represents a package from DNF/YUM repository
type DNFPackage struct {
	PkgKey         int
	PkgId          string
	Name           string
	Arch           sql.NullString
	Version        sql.NullString
	Epoch          sql.NullString
	Release        sql.NullString
	Summary        sql.NullString
	Description    sql.NullString
	URL            sql.NullString
	TimeFile       sql.NullInt64
	TimeBuild      sql.NullInt64
	RPMLicense     sql.NullString
	RPMVendor      sql.NullString
	RPMGroup       sql.NullString
	RPMBuildHost   sql.NullString
	RPMSourceRPM   sql.NullString
	RPMHeaderStart sql.NullInt64
	RPMHeaderEnd   sql.NullInt64
	RPMPackager    sql.NullString
	SizePackage    sql.NullInt64
	SizeInstalled  sql.NullInt64
	SizeArchive    sql.NullInt64
	LocationHref   sql.NullString
	LocationBase   sql.NullString
	ChecksumType   sql.NullString
}

// ExtractDNFPackageInfo extracts package information for a specific PostgreSQL version
func ExtractDNFPackageInfo(pkg *DNFPackage, pgVer int) *PackageInfo {
	// Check if this package is for the specified PostgreSQL version
	pgSuffix := fmt.Sprintf("_%d", pgVer)
	if !strings.Contains(pkg.Name, pgSuffix) {
		return nil
	}

	// Skip development and LLVM JIT packages
	if strings.Contains(pkg.Name, "-llvmjit") || strings.Contains(pkg.Name, "-devel") {
		return nil
	}

	info := &PackageInfo{
		Name:    pkg.Name,
		PkgId:   pkg.PkgId,
		Version: getString2(pkg.Version),
		Arch:    getString2(pkg.Arch),
		Release: getString2(pkg.Release),
	}

	// Extract pname (package name without PG version suffix)
	if idx := strings.Index(pkg.Name, pgSuffix); idx > 0 {
		info.PName = pkg.Name[:idx+len(pgSuffix)]
	} else {
		info.PName = pkg.Name
	}

	// Build full version string
	if pkg.Release.Valid && pkg.Release.String != "" {
		info.FullVersion = info.Version + "-" + info.Release
	} else {
		info.FullVersion = info.Version
	}

	// Extract base version (without release)
	info.BaseVersion = info.Version

	// Extract size information
	if pkg.SizePackage.Valid {
		info.Size = pkg.SizePackage.Int64
	}
	if pkg.SizeInstalled.Valid {
		info.SizeInstall = pkg.SizeInstalled.Int64
	}

	// Extract file information
	if pkg.LocationHref.Valid {
		info.Filename = pkg.LocationHref.String
		// Extract just the filename from the path
		if idx := strings.LastIndex(info.Filename, "/"); idx >= 0 {
			info.File = info.Filename[idx+1:]
		} else {
			info.File = info.Filename
		}
	}

	info.SHA256 = pkg.PkgId

	return info
}

// Helper functions to convert sql.Null types
func nullStr(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

func nullInt(ni sql.NullInt64) interface{} {
	if ni.Valid {
		return ni.Int64
	}
	return nil
}

func getString2(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// PackageInfo represents common package information
type PackageInfo struct {
	Name        string
	PName       string
	Version     string
	BaseVersion string
	FullVersion string
	Release     string
	Arch        string
	PkgId       string
	SHA256      string
	Filename    string
	File        string
	Size        int64
	SizeInstall int64
}