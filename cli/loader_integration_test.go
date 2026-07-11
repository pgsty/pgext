package cli

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

func TestLoadRepositoryInvalidatesDependentCatalogIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var pg int
	var osName, extName, pkgName, repoID string
	if err := QueryRowContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg LIMIT 1").Scan(&pg); err != nil {
		t.Fatalf("select PostgreSQL fixture: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT os FROM pgext.active_os ORDER BY os LIMIT 1").Scan(&osName); err != nil {
		t.Fatalf("select OS fixture: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT name, pkg FROM pgext.extension WHERE lead ORDER BY id LIMIT 1").Scan(&extName, &pkgName); err != nil {
		t.Fatalf("select extension fixture: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
		t.Fatalf("select repository fixture: %v", err)
	}
	if _, err := ExecSQLContext(ctx, "INSERT INTO pgext.repo_data (id, data) VALUES ($1, 'fixture')", repoID); err != nil {
		t.Fatalf("seed repository data: %v", err)
	}
	if _, err := ExecSQLContext(ctx, `
		INSERT INTO pgext.pkg (pg, os, name, pkg, ext, state)
		VALUES ($1, $2, 'load-sentinel', $3, $4, 'AVAIL')
	`, pg, osName, pkgName, extName); err != nil {
		t.Fatalf("seed package sentinel: %v", err)
	}
	if _, err := ExecSQLContext(ctx, `
		UPDATE pgext.status
		SET fetch_time = clock_timestamp(), parse_time = clock_timestamp(), recap_time = clock_timestamp()
		WHERE id = 0
	`); err != nil {
		t.Fatalf("seed status timestamps: %v", err)
	}

	if err := LoadTable("repository", "", RegionDefault); err != nil {
		t.Fatalf("load embedded repository table: %v", err)
	}

	var repoData, pkgRows int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.repo_data").Scan(&repoData); err != nil {
		t.Fatalf("count repo_data after load: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.pkg").Scan(&pkgRows); err != nil {
		t.Fatalf("count pkg after load: %v", err)
	}
	if repoData != 0 || pkgRows != 0 {
		t.Fatalf("dependent rows after repository load = repo_data:%d pkg:%d, want 0/0", repoData, pkgRows)
	}

	var fetchTime, parseTime, recapTime sql.NullTime
	if err := QueryRowContext(ctx, "SELECT fetch_time, parse_time, recap_time FROM pgext.status WHERE id = 0").Scan(&fetchTime, &parseTime, &recapTime); err != nil {
		t.Fatalf("read status after repository load: %v", err)
	}
	if fetchTime.Valid || !parseTime.Valid || recapTime.Valid {
		t.Fatalf("status after repository load = fetch:%v parse:%v recap:%v", fetchTime, parseTime, recapTime)
	}
	if !packageCatalogStale(map[string]*time.Time{"parse": &parseTime.Time}) {
		t.Fatal("repository load did not mark the package catalog stale")
	}
	if got := packageCatalogRefreshCommand(map[string]*time.Time{"parse": &parseTime.Time}); got != "pgext fetch" {
		t.Fatalf("repository load refresh command = %q, want pgext fetch", got)
	}
}
