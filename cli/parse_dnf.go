/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// DNFParser parses RPM primary XML metadata into PostgreSQL.
type DNFParser struct {
	ctx   context.Context
	tx    pgx.Tx
	table string
}

// NewDNFParser creates a new DNF parser
func NewDNFParser(ctx context.Context) *DNFParser {
	return newDNFParser(ctx, liveTable("dnf"))
}

func newDNFParser(ctx context.Context, table string) *DNFParser {
	return &DNFParser{ctx: ctx, table: table}
}

// ParseRepository parses a single DNF/YUM repository
func (p *DNFParser) ParseRepository(repoID string, data []byte) (int, error) {
	packages, err := p.extractPrimaryXML(data)
	if err != nil {
		return 0, fmt.Errorf("parse RPM primary XML: %w; refresh cached metadata with pgext fetch or pgext scan", err)
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
	defer func() {
		rollbackPackageTransaction(tx)
		p.tx = nil
	}()

	// Insert packages
	count, err := p.insertPackages(repoID, packages)
	if err != nil {
		return count, err
	}

	// Commit transaction
	if err := tx.Commit(p.ctx); err != nil {
		// Transaction failed, rollback is automatic
		return count, fmt.Errorf("commit transaction: %w", err)
	}

	return count, nil
}

// insertPackages inserts DNF packages into the database
func (p *DNFParser) insertPackages(repoID string, packages []DNFPackage) (int, error) {
	if len(packages) == 0 {
		return 0, nil
	}

	batch := &pgx.Batch{}
	count := 0

	for _, pkg := range packages {
		// Convert SQL null types to interface{} for pgx
		batch.Queue(fmt.Sprintf(`
			INSERT INTO %s (
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
			)`, p.table),
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
					return count, fmt.Errorf("insert DNF batch: %w", err)
				}
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 && p.tx != nil {
		br := p.tx.SendBatch(p.ctx, batch)
		if err := br.Close(); err != nil {
			return count, fmt.Errorf("insert final DNF batch: %w", err)
		}
	}

	return count, nil
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
