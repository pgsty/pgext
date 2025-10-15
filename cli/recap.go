/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

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

var (
	// Repository cache for package processing
	repositoryCache      map[string]*Repository
	repositoryCacheMutex sync.RWMutex

	// PostgreSQL version patterns
	pgVersionPatterns = []struct {
		version int
		pattern *regexp.Regexp
	}{
		{18, regexp.MustCompile(`[_-]18`)},
		{17, regexp.MustCompile(`[_-]17`)},
		{16, regexp.MustCompile(`[_-]16`)},
		{15, regexp.MustCompile(`[_-]15`)},
		{14, regexp.MustCompile(`[_-]14`)},
		{13, regexp.MustCompile(`[_-]13`)},
	}

	// Package name cleanup patterns for normalization
	packageCleanupPatterns = []*regexp.Regexp{
		regexp.MustCompile(`-scripts$`),
		regexp.MustCompile(`-doc$`),
		regexp.MustCompile(`-dbgsym$`),
		regexp.MustCompile(`-pgq-node$`),
	}
)

// RecapMatrix generates pgext.pkg availability matrix using SQL procedures
func RecapMatrix() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	logrus.Info("generate pgext.pkg from repo package data")

	// Get active PostgreSQL versions
	var activePG []int
	rows, err := QueryContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg DESC")
	if err != nil {
		return fmt.Errorf("failed to query active pg versions: %w", err)
	}
	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err == nil {
			activePG = append(activePG, pg)
		}
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
		if err := rows.Scan(&os); err == nil {
			activeOS = append(activeOS, os)
		}
	}
	rows.Close()

	// Format and log active versions
	pgVersions := make([]string, len(activePG))
	for i, pg := range activePG {
		pgVersions[i] = fmt.Sprintf("%d", pg)
	}
	logrus.Infof("active pg: %s", strings.Join(pgVersions, ", "))
	logrus.Infof("active os: %s", strings.Join(activeOS, ", "))

	// Call the refresh_pkg procedure
	logrus.Info("refreshing package availability matrix...")
	if _, err := ExecSQLContext(ctx, "CALL pgext.refresh_pkg()"); err != nil {
		return fmt.Errorf("failed to refresh pkg: %w", err)
	}

	// Vacuum full the pkg table
	logrus.Info("vacuuming pgext.pkg table...")
	if _, err := ExecSQLContext(ctx, "VACUUM FULL pgext.pkg"); err != nil {
		return fmt.Errorf("failed to vacuum pkg: %w", err)
	}

	// Cluster the pkg table
	logrus.Info("clustering pgext.pkg table...")
	if _, err := ExecSQLContext(ctx, "CLUSTER pgext.pkg USING pkg_pkg_os_pg_idx"); err != nil {
		return fmt.Errorf("failed to cluster pkg: %w", err)
	}

	// Count records in pgext.pkg
	var pkgCount int64
	if err := QueryRowContext(ctx, "SELECT COUNT(*) FROM pgext.pkg").Scan(&pkgCount); err != nil {
		return fmt.Errorf("failed to count pkg records: %w", err)
	}

	logrus.Infof("recap completed: %d package records in pgext.pkg", pkgCount)
	return nil
}

// loadRepositoryCache loads repository metadata into cache
func loadRepositoryCache(ctx context.Context) error {
	query := `
		SELECT id, type, os, org, os_code, os_arch,
		       COALESCE(default_url, ''), COALESCE(mirror_url, '')
		FROM pgext.repository`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("query repositories: %w", err)
	}
	defer rows.Close()

	repositoryCacheMutex.Lock()
	repositoryCache = make(map[string]*Repository)
	repositoryCacheMutex.Unlock()

	count := 0
	for rows.Next() {
		var repo Repository
		if err := rows.Scan(&repo.ID, &repo.Type, &repo.OS, &repo.Org,
			&repo.OSCode, &repo.OSArch, &repo.DefaultURL, &repo.MirrorURL); err != nil {
			logrus.Warnf("failed to scan repository: %v", err)
			continue
		}

		repositoryCacheMutex.Lock()
		repositoryCache[repo.ID] = &repo
		repositoryCacheMutex.Unlock()
		count++
	}

	logrus.Debugf("loaded %d repositories into cache", count)
	return rows.Err()
}

// getRepository gets repository from cache
func getRepository(id string) *Repository {
	repositoryCacheMutex.RLock()
	defer repositoryCacheMutex.RUnlock()
	return repositoryCache[id]
}

// buildPackageTable builds the package table from yum and apt tables
func buildPackageTable(ctx context.Context) (int64, error) {
	// Clear package table
	if _, err := ExecSQLContext(ctx, "TRUNCATE TABLE pgext.package"); err != nil {
		return 0, fmt.Errorf("truncate package table: %w", err)
	}

	// Process YUM packages
	yumCount, err := processYumPackages(ctx)
	if err != nil {
		return 0, fmt.Errorf("process yum packages: %w", err)
	}

	// Process APT packages
	aptCount, err := processAptPackages(ctx)
	if err != nil {
		return 0, fmt.Errorf("process apt packages: %w", err)
	}

	return yumCount + aptCount, nil
}

// processYumPackages processes YUM packages and inserts into package table
func processYumPackages(ctx context.Context) (int64, error) {
	query := `
		SELECT repo_id, name, version, release, location_href,
		       pkg_id, size_package, size_installed
		FROM pgext.yum
		ORDER BY repo_id, name`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	tx, err := Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}
	count := int64(0)

	for rows.Next() {
		var (
			repoID        string
			name          string
			version       string
			release       string
			locationHref  string
			pkgID         sql.NullString
			sizePackage   sql.NullInt64
			sizeInstalled sql.NullInt64
		)

		if err := rows.Scan(&repoID, &name, &version, &release, &locationHref,
			&pkgID, &sizePackage, &sizeInstalled); err != nil {
			continue
		}

		// Skip devel and debug packages
		if strings.Contains(name, "-devel") || strings.Contains(name, "-debuginfo") {
			continue
		}

		// Get repository info
		repo := getRepository(repoID)
		if repo == nil {
			continue
		}

		// Extract PostgreSQL version
		pgVersion := extractPGVersion(name)
		if pgVersion == 0 {
			continue // Skip non-PostgreSQL packages
		}

		// Normalize package name
		pname := normalizePackageName(name, pgVersion, "rpm")

		// Build package entry
		pkg := Package{
			PG:        pgVersion,
			OS:        repo.OS,
			PName:     pname,
			Org:       repo.Org,
			Type:      repo.Type,
			OSCode:    repo.OSCode,
			OSArch:    repo.OSArch,
			Repo:      repoID,
			Name:      name,
			Ver:       fmt.Sprintf("%s-%s", version, release),
			Version:   version,
			Release:   release,
			File:      extractFileName(locationHref),
			SHA256:    pkgID.String,
			URL:       buildURL(repo.DefaultURL, locationHref),
			MirrorURL: buildURL(repo.MirrorURL, locationHref),
			Size:      sizePackage.Int64,
			SizeFull:  sizeInstalled.Int64,
		}

		// Queue insert
		batch.Queue(`
			INSERT INTO pgext.package (
				pg, os, pname, org, type, os_code, os_arch, repo, name, ver,
				version, release, file, sha256, url, mirror_url, size, size_full
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, $18
			)`, pkg.PG, pkg.OS, pkg.PName, pkg.Org, pkg.Type, pkg.OSCode, pkg.OSArch,
			pkg.Repo, pkg.Name, pkg.Ver, pkg.Version, pkg.Release, pkg.File,
			pkg.SHA256, pkg.URL, pkg.MirrorURL, pkg.Size, pkg.SizeFull)

		count++

		// Send batch every 1000 rows
		if batch.Len() >= 1000 {
			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return count, fmt.Errorf("send batch: %w", err)
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 {
		br := tx.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			return count, fmt.Errorf("send final batch: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return count, err
	}

	logrus.Debugf("processed %d YUM packages into package table", count)
	return count, rows.Err()
}

// processAptPackages processes APT packages and inserts into package table
func processAptPackages(ctx context.Context) (int64, error) {
	query := `
		SELECT repo_id, package, version, filename, sha256, size, size_install
		FROM pgext.apt
		ORDER BY repo_id, package`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	tx, err := Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}
	count := int64(0)

	for rows.Next() {
		var (
			repoID      string
			packageName string
			version     string
			filename    string
			sha256      sql.NullString
			size        sql.NullInt64
			sizeInstall sql.NullInt64
		)

		if err := rows.Scan(&repoID, &packageName, &version, &filename,
			&sha256, &size, &sizeInstall); err != nil {
			continue
		}

		// Skip devel and debug packages
		if strings.Contains(packageName, "-dev") || strings.Contains(packageName, "-dbgsym") {
			continue
		}

		// Get repository info
		repo := getRepository(repoID)
		if repo == nil {
			continue
		}

		// Extract PostgreSQL version
		pgVersion := extractPGVersion(packageName)
		if pgVersion == 0 {
			continue // Skip non-PostgreSQL packages
		}

		// Normalize package name
		pname := normalizePackageName(packageName, pgVersion, "deb")

		// Extract version and release from version string
		semVer, release := extractVersionRelease(version)

		// Build package entry
		pkg := Package{
			PG:        pgVersion,
			OS:        repo.OS,
			PName:     pname,
			Org:       repo.Org,
			Type:      repo.Type,
			OSCode:    repo.OSCode,
			OSArch:    repo.OSArch,
			Repo:      repoID,
			Name:      packageName,
			Ver:       version,
			Version:   semVer,
			Release:   release,
			File:      extractFileName(filename),
			SHA256:    sha256.String,
			URL:       buildURL(repo.DefaultURL, filename),
			MirrorURL: buildURL(repo.MirrorURL, filename),
			Size:      size.Int64,
			SizeFull:  sizeInstall.Int64,
		}

		// Queue insert
		batch.Queue(`
			INSERT INTO pgext.package (
				pg, os, pname, org, type, os_code, os_arch, repo, name, ver,
				version, release, file, sha256, url, mirror_url, size, size_full
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, $18
			)`, pkg.PG, pkg.OS, pkg.PName, pkg.Org, pkg.Type, pkg.OSCode, pkg.OSArch,
			pkg.Repo, pkg.Name, pkg.Ver, pkg.Version, pkg.Release, pkg.File,
			pkg.SHA256, pkg.URL, pkg.MirrorURL, pkg.Size, pkg.SizeFull)

		count++

		// Send batch every 1000 rows
		if batch.Len() >= 1000 {
			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return count, fmt.Errorf("send batch: %w", err)
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 {
		br := tx.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			return count, fmt.Errorf("send final batch: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return count, err
	}

	logrus.Debugf("processed %d APT packages into package table", count)
	return count, rows.Err()
}

// extractPGVersion extracts PostgreSQL version from package name
func extractPGVersion(name string) int {
	for _, vp := range pgVersionPatterns {
		if vp.pattern.MatchString(name) {
			// Additional check to ensure it's really a PostgreSQL package
			// Look for common PostgreSQL-related patterns
			lowerName := strings.ToLower(name)
			if strings.Contains(lowerName, "postgresql") ||
				strings.Contains(lowerName, "pg_") ||
				strings.Contains(lowerName, "pgaudit") ||
				strings.Contains(lowerName, "pgrouting") ||
				strings.Contains(lowerName, "postgis") ||
				strings.Contains(lowerName, "pgtap") ||
				strings.Contains(lowerName, "pglogical") ||
				strings.Contains(lowerName, "repmgr") ||
				strings.Contains(lowerName, "pgbouncer") ||
				strings.Contains(lowerName, "pgpool") ||
				strings.Contains(lowerName, "barman") ||
				strings.Contains(lowerName, "pgbackrest") ||
				strings.Contains(lowerName, "wal2json") ||
				strings.Contains(lowerName, "timescaledb") ||
				strings.Contains(lowerName, "citus") ||
				strings.Contains(lowerName, "hydra") ||
				strings.Contains(lowerName, "age") ||
				strings.Contains(lowerName, "plpgsql") ||
				strings.Contains(lowerName, "plproxy") ||
				strings.Contains(lowerName, "multicorn") ||
				strings.Contains(lowerName, "cstore") ||
				strings.Contains(lowerName, "orafce") ||
				strings.Contains(lowerName, "pgvector") {
				return vp.version
			}
			// For packages with version suffix directly (e.g., package_17)
			if regexp.MustCompile(fmt.Sprintf(`[_-]%d($|[_-])`, vp.version)).MatchString(name) {
				return vp.version
			}
		}
	}
	return 0
}

// normalizePackageName normalizes package name by removing suffixes and version
func normalizePackageName(name string, pgVersion int, pkgType string) string {
	result := name

	// Remove PostgreSQL version suffix
	if pkgType == "rpm" {
		// For RPM: postgresql_17 -> postgresql
		pattern := fmt.Sprintf(`_%d.*$`, pgVersion)
		result = regexp.MustCompile(pattern).ReplaceAllString(result, "")
		result = fmt.Sprintf("%s_%d", result, pgVersion)
	} else {
		// For DEB: postgresql-17 -> postgresql
		pattern := fmt.Sprintf(`-%d.*$`, pgVersion)
		result = regexp.MustCompile(pattern).ReplaceAllString(result, "")
		result = fmt.Sprintf("%s-%d", result, pgVersion)
	}

	// Clean up known suffixes
	for _, pattern := range packageCleanupPatterns {
		result = pattern.ReplaceAllString(result, "")
	}

	// Special case for pgq-node -> pgq
	result = strings.ReplaceAll(result, "-pgq-node", "-pgq")

	return result
}

// extractVersionRelease extracts version and release from version string
func extractVersionRelease(fullVersion string) (version, release string) {
	// Try to match pattern: version-release
	pattern := regexp.MustCompile(`^(.*?)-([^-]+)$`)
	if matches := pattern.FindStringSubmatch(fullVersion); len(matches) == 3 {
		return matches[1], matches[2]
	}
	// If no release part, return version as is
	return fullVersion, ""
}

// extractFileName extracts filename from path
func extractFileName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}

// buildURL builds complete URL from base and path
func buildURL(base, path string) string {
	if base == "" || path == "" {
		return ""
	}
	// Remove trailing slash from base and leading slash from path
	base = strings.TrimSuffix(base, "/")
	path = strings.TrimPrefix(path, "/")
	return fmt.Sprintf("%s/%s", base, path)
}

// RecapPackages is the main entry point for the recap operation
func RecapPackages() error {
	return RecapMatrix()
}

// updateRecapTimestamp updates the recap_mtime in pgext.status table
func updateRecapTimestamp(ctx context.Context) error {
	_, err := ExecSQLContext(ctx, `
		INSERT INTO pgext.status (id, recap_mtime, matrix_mtime)
		VALUES (0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (id) DO UPDATE
		SET recap_mtime = CURRENT_TIMESTAMP,
		    matrix_mtime = CURRENT_TIMESTAMP`)

	if err != nil {
		return fmt.Errorf("update recap_mtime: %w", err)
	}

	logrus.Debug("Updated recap_mtime and matrix_mtime in pgext.status")
	return nil
}