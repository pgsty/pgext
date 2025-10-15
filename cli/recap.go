/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"pgext/db"
	"strings"
	"time"

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

	// Call the reload_pkg procedure
	logrus.Info("reload pgext.pkg")

	// Apply reload.sql
	if err := applyReloadPkg(ctx); err != nil {
		return fmt.Errorf("failed to apply pkg fixes: %w", err)
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

// RecapPackages is the main entry point for the recap operation
func RecapPackages() error {
	return RecapMatrix()
}

// applyReloadPkg executes db/reload.sql to generate pgext.pkg table
func applyReloadPkg(ctx context.Context) error {
	logrus.Info("run db/reload.sql...")
	sqlBytes, err := db.ReadFile("reload.sql")
	if err != nil {
		return fmt.Errorf("failed to read reload.sql: %w", err)
	}
	sqlContent := string(sqlBytes)

	// Skip if file is empty or only contains comments/whitespace
	trimmed := strings.TrimSpace(sqlContent)
	if trimmed == "" || !containsSQL(trimmed) {
		logrus.Debug("reload.sql is empty, skip")
		return nil
	}

	// Execute the SQL
	if _, err := ExecSQLContext(ctx, sqlContent); err != nil {
		return fmt.Errorf("failed to execute reload.sql: %w", err)
	}

	logrus.Info("pkg fixes applied successfully")
	return nil
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
