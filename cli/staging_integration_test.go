package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestPackageStagingIntegration validates CREATE LIKE column parity, INSERT
// SELECT compatibility, and the real reload.sql rewrite without publishing a
// staged catalog. PGEXT_TEST_PGURL must name an explicitly disposable database,
// and PGEXT_TEST_ALLOW_DESTRUCTIVE must be set to yes.
func TestPackageStagingIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	seedStagingIntegrationRows(t, ctx)
	stage, err := newPackageStaging(ctx, true, true)
	if err != nil {
		t.Fatalf("create staging tables: %v", err)
	}
	t.Cleanup(stage.Drop)

	for _, table := range []struct {
		name  string
		live  string
		stage string
	}{
		{"apt", liveTable("apt"), stage.apt},
		{"dnf", liveTable("dnf"), stage.dnf},
		{"bin", liveTable("bin"), stage.bin},
		{"pkg", liveTable("pkg"), stage.pkg},
	} {
		liveColumns := integrationTableColumns(t, ctx, table.live)
		stageColumns := integrationTableColumns(t, ctx, table.stage)
		if !reflect.DeepEqual(liveColumns, stageColumns) {
			t.Fatalf("%s staging columns differ:\nlive:  %v\nstage: %v", table.name, liveColumns, stageColumns)
		}
		if table.name != "pkg" {
			query := fmt.Sprintf("INSERT INTO %s SELECT * FROM %s LIMIT 1", table.stage, table.live)
			if _, err := ExecSQLContext(ctx, query); err != nil {
				t.Fatalf("%s INSERT SELECT compatibility: %v", table.name, err)
			}
		}
	}

	var before, after sql.NullString
	if err := QueryRowContext(ctx, "SELECT recap_time::text FROM pgext.status WHERE id = 0").Scan(&before); err != nil {
		t.Fatalf("read recap timestamp before staged build: %v", err)
	}
	if _, err := stage.buildPkg(ctx, stage.bin); err != nil {
		t.Fatalf("build staged pkg with real reload.sql: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT recap_time::text FROM pgext.status WHERE id = 0").Scan(&after); err != nil {
		t.Fatalf("read recap timestamp after staged build: %v", err)
	}
	if before != after {
		t.Fatalf("staged pkg build changed live recap timestamp: before=%v after=%v", before, after)
	}
}

func TestStrictParserDoesNotPublishPartialCatalogIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	seedStagingIntegrationRows(t, ctx)

	var goodRepo, badRepo string
	var pg int
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'deb' ORDER BY id LIMIT 1").Scan(&goodRepo); err != nil {
		t.Fatalf("select valid repository: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'deb' AND id <> $1 ORDER BY id LIMIT 1", goodRepo).Scan(&badRepo); err != nil {
		t.Fatalf("select failing repository: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg LIMIT 1").Scan(&pg); err != nil {
		t.Fatalf("select active PG: %v", err)
	}
	packageData := []byte(fmt.Sprintf(`Package: postgresql-%d-stage-test
Version: 1.0.0
Architecture: amd64
Size: 100
Installed-Size: 1
Filename: pool/stage-test.deb
SHA256: stage-test

`, pg))
	if _, err := ExecSQLContext(ctx, "INSERT INTO pgext.repo_data (id, data) VALUES ($1, $2), ($3, $4)",
		goodRepo, packageData, badRepo, []byte("invalid")); err != nil {
		t.Fatalf("seed repository metadata: %v", err)
	}
	parser := NewParserContext(ctx, ParseOptions{Parallel: 2})
	if err := parser.ParseAndRecap(); err == nil {
		t.Fatal("strict parser accepted malformed non-empty APT metadata")
	}
	var aptCount int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.apt").Scan(&aptCount); err != nil {
		t.Fatalf("count live APT rows after failed parse: %v", err)
	}
	if aptCount != 1 {
		t.Fatalf("failed strict parse changed live APT table: got %d rows, want sentinel row", aptCount)
	}

	if _, err := ExecSQLContext(ctx, "DELETE FROM pgext.repo_data WHERE id = $1", badRepo); err != nil {
		t.Fatalf("remove failing repository metadata: %v", err)
	}
	if err := NewParserContext(ctx, ParseOptions{Parallel: 2}).ParseAndRecap(); err != nil {
		t.Fatalf("publish complete staged catalog: %v", err)
	}

	var binCount, pkgCount int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.apt").Scan(&aptCount); err != nil {
		t.Fatalf("count published APT rows: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.bin").Scan(&binCount); err != nil {
		t.Fatalf("count published bin rows: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.pkg").Scan(&pkgCount); err != nil {
		t.Fatalf("count published pkg rows: %v", err)
	}
	if aptCount != 1 || binCount == 0 || pkgCount == 0 {
		t.Fatalf("published catalog counts = apt:%d bin:%d pkg:%d", aptCount, binCount, pkgCount)
	}

	var leaked int
	if err := QueryRowContext(ctx, `SELECT count(*) FROM pg_tables
		WHERE schemaname = 'pgext' AND tablename LIKE '\_stage\_%' ESCAPE '\'`).Scan(&leaked); err != nil {
		t.Fatalf("count leaked staging tables: %v", err)
	}
	if leaked != 0 {
		t.Fatalf("found %d leaked staging tables", leaked)
	}
}

func TestBinGeneratorRejectsPartialSourcesIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	stage, err := newPackageStaging(ctx, true, false)
	if err != nil {
		t.Fatalf("create binary-generator staging tables: %v", err)
	}
	t.Cleanup(stage.Drop)

	var repoID string
	var pg int
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'deb' ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
		t.Fatalf("select binary-generator repository: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg LIMIT 1").Scan(&pg); err != nil {
		t.Fatalf("select binary-generator PostgreSQL version: %v", err)
	}
	if _, err := ExecSQLContext(ctx, fmt.Sprintf(`
		INSERT INTO %s (repo, package, version, architecture, filename)
		VALUES ($1, $2, '1.0.0', 'amd64', 'pool/partial-source.deb')
	`, stage.apt), repoID, fmt.Sprintf("postgresql-%d-partial-source", pg)); err != nil {
		t.Fatalf("seed valid APT source: %v", err)
	}

	generator := newBinGenerator(ctx, `"pgext"."missing_dnf_source"`, stage.apt, stage.bin)
	if _, err := generator.Generate(); err == nil {
		t.Fatal("binary generator accepted an APT-only result after the DNF source query failed")
	}
}

func TestRecapRejectsDestructiveEmptyMatrixIntegration(t *testing.T) {
	for _, tt := range []struct {
		name   string
		mutate string
	}{
		{name: "empty binary package source", mutate: "TRUNCATE pgext.bin"},
		{name: "no active operating systems", mutate: "UPDATE pgext.os SET active = false"},
		{name: "no lead extensions", mutate: "UPDATE pgext.extension SET lead = false"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			setupDisposableIntegrationDB(t)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()

			fixture := selectRecapFixture(t, ctx)
			if _, err := ExecSQLContext(ctx, `
				INSERT INTO pgext.bin (pg, os, name, repo, ver, version, file)
				VALUES ($1, $2, $3, $4, '1.0.0', '1.0.0', 'empty-guard.pkg')
			`, fixture.pg, fixture.os, fixture.packageName, fixture.repo); err != nil {
				t.Fatalf("seed empty-matrix guard: %v", err)
			}
			if err := RecapMatrixContext(ctx); err != nil {
				t.Fatalf("build initial package matrix: %v", err)
			}

			var before int
			if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.pkg").Scan(&before); err != nil {
				t.Fatalf("count initial package matrix: %v", err)
			}
			if before == 0 {
				t.Fatal("initial package matrix is empty")
			}
			if _, err := ExecSQLContext(ctx, tt.mutate); err != nil {
				t.Fatalf("apply empty-matrix fixture: %v", err)
			}
			if err := RecapMatrixContext(ctx); err == nil {
				t.Fatal("recap accepted a destructive empty package matrix")
			}

			var after int
			if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.pkg").Scan(&after); err != nil {
				t.Fatalf("count package matrix after rejected recap: %v", err)
			}
			if after != before {
				t.Fatalf("rejected recap changed live pkg rows: before=%d after=%d", before, after)
			}
		})
	}
}

func TestParseAllPublishesWithoutPkgIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	var repoID, osName, extName, pkgName string
	var pg int
	if err := QueryRowContext(ctx, `
		SELECT r.id, o.os, p.pg, e.name, e.pkg
		FROM pgext.repository r
		JOIN pgext.os o ON o.os = r.os AND o.active
		CROSS JOIN LATERAL (SELECT pg FROM pgext.pg WHERE active ORDER BY pg LIMIT 1) p
		CROSS JOIN LATERAL (
			SELECT name, pkg FROM pgext.extension
			WHERE lead AND NOT contrib ORDER BY id LIMIT 1
		) e
		WHERE r.type = 'deb'
		ORDER BY r.id
		LIMIT 1
	`).Scan(&repoID, &osName, &pg, &extName, &pkgName); err != nil {
		t.Fatalf("select no-pkg fixture: %v", err)
	}

	const recapMarker = "2001-02-03 04:05:06+00"
	if _, err := ExecSQLContext(ctx, `
		INSERT INTO pgext.pkg (pg, os, name, pkg, ext, state, version, count)
		VALUES ($1, $2, 'sentinel-package', $3, $4, 'AVAIL', 'sentinel', 1)
	`, pg, osName, pkgName, extName); err != nil {
		t.Fatalf("insert pkg sentinel: %v", err)
	}
	if _, err := ExecSQLContext(ctx, `
		UPDATE pgext.status SET parse_time = NULL, recap_time = $1::timestamptz WHERE id = 0
	`, recapMarker); err != nil {
		t.Fatalf("set status sentinel: %v", err)
	}
	packageData := []byte(fmt.Sprintf(`Package: postgresql-%d-no-pkg-test
Version: 1.0.0
Architecture: amd64
Size: 100
Installed-Size: 1
Filename: pool/no-pkg-test.deb
SHA256: no-pkg-test

`, pg))
	if _, err := ExecSQLContext(ctx, "INSERT INTO pgext.repo_data (id, data) VALUES ($1, $2)", repoID, packageData); err != nil {
		t.Fatalf("seed no-pkg repository data: %v", err)
	}

	if err := NewParserContext(ctx, ParseOptions{Parallel: 1}).ParseAll(); err != nil {
		t.Fatalf("publish no-pkg parse: %v", err)
	}

	var aptCount, binCount, sentinelCount int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.apt").Scan(&aptCount); err != nil {
		t.Fatalf("count no-pkg apt rows: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.bin").Scan(&binCount); err != nil {
		t.Fatalf("count no-pkg bin rows: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.pkg WHERE pkg = $1 AND version = 'sentinel'", pkgName).Scan(&sentinelCount); err != nil {
		t.Fatalf("count pkg sentinel: %v", err)
	}
	if aptCount == 0 || binCount == 0 || sentinelCount != 1 {
		t.Fatalf("no-pkg counts = apt:%d bin:%d sentinel:%d", aptCount, binCount, sentinelCount)
	}

	var parseTime sql.NullTime
	var recapTime string
	if err := QueryRowContext(ctx, `
		SELECT parse_time, to_char(recap_time AT TIME ZONE 'UTC', 'YYYY-MM-DD HH24:MI:SS')
		FROM pgext.status WHERE id = 0
	`).Scan(&parseTime, &recapTime); err != nil {
		t.Fatalf("read no-pkg timestamps: %v", err)
	}
	if !parseTime.Valid {
		t.Fatal("no-pkg parse did not update parse_time")
	}
	if recapTime != "2001-02-03 04:05:06" {
		t.Fatalf("no-pkg parse changed recap_time to %q", recapTime)
	}
}

func TestParserWaitsForPipelineLockBeforeReadingRepoDataIntegration(t *testing.T) {
	for _, tt := range []struct {
		name string
		run  func(*Parser) error
	}{
		{name: "core", run: func(parser *Parser) error { return parser.ParseAll() }},
		{name: "full", run: func(parser *Parser) error { return parser.ParseAndRecap() }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			setupDisposableIntegrationDB(t)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()

			var repoID string
			var pg int
			if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'deb' ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
				t.Fatalf("select pipeline-lock repository: %v", err)
			}
			if err := QueryRowContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg LIMIT 1").Scan(&pg); err != nil {
				t.Fatalf("select pipeline-lock PostgreSQL version: %v", err)
			}

			v1Name := fmt.Sprintf("postgresql-%d-pipeline-lock-v1", pg)
			v2Name := fmt.Sprintf("postgresql-%d-pipeline-lock-v2", pg)
			v1Data := aptIntegrationPackage(v1Name, "1.0.0", "pipeline-lock-v1")
			v2Data := aptIntegrationPackage(v2Name, "2.0.0", "pipeline-lock-v2")
			if _, err := ExecSQLContext(ctx, `
				INSERT INTO pgext.repo_data (id, data) VALUES ($1, $2)
				ON CONFLICT (id) DO UPDATE SET data = EXCLUDED.data
			`, repoID, v1Data); err != nil {
				t.Fatalf("seed pipeline-lock v1 metadata: %v", err)
			}

			gate, err := Begin(ctx)
			if err != nil {
				t.Fatalf("begin pipeline-lock gate: %v", err)
			}
			defer rollbackPackageTransaction(gate)
			if err := lockPackagePipeline(ctx, gate); err != nil {
				t.Fatalf("lock parse pipeline gate: %v", err)
			}

			parseDone := make(chan error, 1)
			go func() {
				parseDone <- tt.run(NewParserContext(ctx, ParseOptions{Parallel: 1}))
			}()
			if err := waitForAdvisoryLockWaiter(ctx, 5*time.Second); err != nil {
				_ = gate.Rollback(context.Background())
				<-parseDone
				t.Fatalf("parser did not wait for the pipeline lock before reading repo_data: %v", err)
			}

			if _, err := gate.Exec(ctx, "UPDATE pgext.repo_data SET data = $2 WHERE id = $1", repoID, v2Data); err != nil {
				t.Fatalf("replace pipeline-lock metadata while parser waits: %v", err)
			}
			if err := gate.Commit(ctx); err != nil {
				t.Fatalf("commit pipeline-lock metadata update: %v", err)
			}
			if err := <-parseDone; err != nil {
				t.Fatalf("parse after pipeline-lock gate: %v", err)
			}

			var v1Count, v2Count int
			if err := QueryRowContext(ctx, `
				SELECT count(*) FILTER (WHERE package = $1),
				       count(*) FILTER (WHERE package = $2)
				FROM pgext.apt WHERE repo = $3
			`, v1Name, v2Name, repoID).Scan(&v1Count, &v2Count); err != nil {
				t.Fatalf("read pipeline-lock parse result: %v", err)
			}
			if v1Count != 0 || v2Count != 1 {
				t.Fatalf("pipeline-lock parse used wrong repo_data snapshot: v1=%d v2=%d", v1Count, v2Count)
			}
		})
	}
}

func TestFullPublishUsesSharedMonotonicTimestampIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	seedStagingIntegrationRows(t, ctx)

	stage, err := newPackageStaging(ctx, true, true)
	if err != nil {
		t.Fatalf("create timestamp staging tables: %v", err)
	}
	t.Cleanup(stage.Drop)
	for _, table := range []struct {
		stage string
		live  string
	}{
		{stage.apt, liveTable("apt")},
		{stage.dnf, liveTable("dnf")},
		{stage.bin, liveTable("bin")},
	} {
		if _, err := ExecSQLContext(ctx, fmt.Sprintf("INSERT INTO %s SELECT * FROM %s", table.stage, table.live)); err != nil {
			t.Fatalf("copy timestamp fixture into %s: %v", table.stage, err)
		}
	}
	if _, err := stage.buildPkg(ctx, stage.bin); err != nil {
		t.Fatalf("build timestamp pkg fixture: %v", err)
	}

	marker := time.Date(2099, time.January, 2, 3, 4, 5, 0, time.UTC)
	if _, err := ExecSQLContext(ctx, `
		UPDATE pgext.status SET parse_time = $1, recap_time = $1 WHERE id = 0
	`, marker); err != nil {
		t.Fatalf("set future timestamp marker: %v", err)
	}
	if err := stage.publish(ctx, true, true); err != nil {
		t.Fatalf("publish catalog with monotonic timestamp: %v", err)
	}

	var parseTime, recapTime time.Time
	if err := QueryRowContext(ctx, `
		SELECT parse_time, recap_time FROM pgext.status WHERE id = 0
	`).Scan(&parseTime, &recapTime); err != nil {
		t.Fatalf("read full publish timestamps: %v", err)
	}
	if !parseTime.Equal(recapTime) {
		t.Fatalf("full publish timestamps differ: parse=%s recap=%s", parseTime, recapTime)
	}
	if !parseTime.After(marker) {
		t.Fatalf("full publish timestamp regressed: got %s, previous %s", parseTime, marker)
	}
}

func aptIntegrationPackage(name, version, checksum string) []byte {
	return []byte(fmt.Sprintf(`Package: %s
Version: %s
Architecture: amd64
Size: 100
Installed-Size: 1
Filename: pool/%s.deb
SHA256: %s

`, name, version, checksum, checksum))
}

func TestStandaloneRecapBuildsAfterCatalogLockIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	fixture := selectRecapFixture(t, ctx)
	if _, err := ExecSQLContext(ctx, `
		INSERT INTO pgext.bin (pg, os, name, repo, ver, version, file)
		VALUES ($1, $2, $3, $4, '1.0.0', '1.0.0', 'recap-race-1.pkg')
	`, fixture.pg, fixture.os, fixture.packageName, fixture.repo); err != nil {
		t.Fatalf("seed recap v1 bin: %v", err)
	}
	if err := RecapMatrixContext(ctx); err != nil {
		t.Fatalf("build initial recap: %v", err)
	}
	assertPkgVersion(t, ctx, fixture, "1.0.0")

	gate, err := Begin(ctx)
	if err != nil {
		t.Fatalf("begin recap gate: %v", err)
	}
	defer rollbackPackageTransaction(gate)
	if err := lockPackageCatalog(ctx, gate); err != nil {
		t.Fatalf("lock recap gate: %v", err)
	}

	recapDone := make(chan error, 1)
	go func() {
		recapDone <- RecapMatrixContext(ctx)
	}()
	if err := waitForAdvisoryLockWaiter(ctx, 5*time.Second); err != nil {
		_ = gate.Rollback(context.Background())
		<-recapDone
		t.Fatalf("standalone recap did not wait for catalog lock: %v", err)
	}

	if _, err := gate.Exec(ctx, `
		UPDATE pgext.bin SET ver = '2.0.0', version = '2.0.0', file = 'recap-race-2.pkg'
		WHERE pg = $1 AND os = $2 AND name = $3 AND repo = $4
	`, fixture.pg, fixture.os, fixture.packageName, fixture.repo); err != nil {
		t.Fatalf("publish recap v2 bin: %v", err)
	}
	if err := gate.Commit(ctx); err != nil {
		t.Fatalf("commit recap gate: %v", err)
	}
	if err := <-recapDone; err != nil {
		t.Fatalf("standalone recap after gate: %v", err)
	}
	assertPkgVersion(t, ctx, fixture, "2.0.0")
}

type recapFixture struct {
	pg          int
	os          string
	pkg         string
	packageName string
	repo        string
}

func selectRecapFixture(t *testing.T, ctx context.Context) recapFixture {
	t.Helper()
	var fixture recapFixture
	err := QueryRowContext(ctx, `
		SELECT p.pg, o.os, e.pkg,
		       CASE o.os_type
		           WHEN 'rpm' THEN replace(replace(e.rpm_pkg, '*', ''), '$v', p.pg::text)
		           WHEN 'deb' THEN replace(replace(e.deb_pkg, '*', ''), '$v', p.pg::text)
		       END AS package_name,
		       r.id
		FROM pgext.extension e
		CROSS JOIN LATERAL (SELECT pg FROM pgext.pg WHERE active ORDER BY pg LIMIT 1) p
		CROSS JOIN pgext.os o
		JOIN pgext.repository r ON r.os = o.os AND r.type = o.os_type
		WHERE e.lead AND NOT e.contrib AND o.active
		  AND e.pkg NOT IN ('postgis', 'pgrouting')
		  AND CASE o.os_type
		          WHEN 'rpm' THEN e.rpm_pkg IS NOT NULL AND position(' ' IN e.rpm_pkg) = 0
		          WHEN 'deb' THEN e.deb_pkg IS NOT NULL AND position(' ' IN e.deb_pkg) = 0
		      END
		ORDER BY e.id, o.os, r.id
		LIMIT 1
	`).Scan(&fixture.pg, &fixture.os, &fixture.pkg, &fixture.packageName, &fixture.repo)
	if err != nil {
		t.Fatalf("select recap fixture: %v", err)
	}
	return fixture
}

func waitForAdvisoryLockWaiter(ctx context.Context, timeout time.Duration) error {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		var waiting bool
		if err := QueryRowContext(ctx, `
			SELECT EXISTS (
				SELECT 1 FROM pg_locks
				WHERE locktype = 'advisory' AND database = (
					SELECT oid FROM pg_database WHERE datname = current_database()
				) AND NOT granted
			)
		`).Scan(&waiting); err != nil {
			return fmt.Errorf("query advisory lock waiter: %w", err)
		}
		if waiting {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-deadline.C:
			return fmt.Errorf("timed out after %s", timeout)
		case <-ticker.C:
		}
	}
}

func assertPkgVersion(t *testing.T, ctx context.Context, fixture recapFixture, want string) {
	t.Helper()
	var got string
	if err := QueryRowContext(ctx, `
		SELECT version FROM pgext.pkg WHERE pg = $1 AND os = $2 AND pkg = $3
	`, fixture.pg, fixture.os, fixture.pkg).Scan(&got); err != nil {
		t.Fatalf("read pkg version: %v", err)
	}
	if got != want {
		t.Fatalf("pkg version = %q, want %q", got, want)
	}
}

func setupDisposableIntegrationDB(t *testing.T) {
	t.Helper()
	pgurl := os.Getenv("PGEXT_TEST_PGURL")
	if pgurl == "" {
		t.Skip("set PGEXT_TEST_PGURL to an explicitly disposable database")
	}
	openDisposableIntegrationDB(t, pgurl)
	if err := InitSchema(true); err != nil {
		t.Fatalf("initialize fresh schema: %v", err)
	}
}

func seedStagingIntegrationRows(t *testing.T, ctx context.Context) {
	t.Helper()
	var debRepo, rpmRepo, osName string
	var pg int
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'deb' ORDER BY id LIMIT 1").Scan(&debRepo); err != nil {
		t.Fatalf("select DEB repository: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type = 'rpm' ORDER BY id LIMIT 1").Scan(&rpmRepo); err != nil {
		t.Fatalf("select RPM repository: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT os FROM pgext.active_os ORDER BY os LIMIT 1").Scan(&osName); err != nil {
		t.Fatalf("select active OS: %v", err)
	}
	if err := QueryRowContext(ctx, "SELECT pg FROM pgext.active_pg ORDER BY pg LIMIT 1").Scan(&pg); err != nil {
		t.Fatalf("select active PG: %v", err)
	}

	statements := []struct {
		sql  string
		args []any
	}{
		{"INSERT INTO pgext.apt (repo, package, version, architecture) VALUES ($1, 'test-deb', '1.0', 'amd64')", []any{debRepo}},
		{"INSERT INTO pgext.dnf (repo, pkg_key, pkg_id, name) VALUES ($1, 1, 'test-rpm-id', 'test-rpm')", []any{rpmRepo}},
		{"INSERT INTO pgext.bin (pg, os, name, repo, ver, version) VALUES ($1, $2, 'test-bin', $3, '1.0', '1.0')", []any{pg, osName, debRepo}},
	}
	for _, statement := range statements {
		if _, err := ExecSQLContext(ctx, statement.sql, statement.args...); err != nil {
			t.Fatalf("seed staging integration row: %v", err)
		}
	}
}

func integrationTableColumns(t *testing.T, ctx context.Context, table string) []string {
	t.Helper()
	rows, err := QueryContext(ctx, `
		SELECT attname
		FROM pg_attribute
		WHERE attrelid = $1::regclass AND attnum > 0 AND NOT attisdropped
		ORDER BY attnum`, table)
	if err != nil {
		t.Fatalf("query columns for %s: %v", table, err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			t.Fatalf("scan column for %s: %v", table, err)
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("iterate columns for %s: %v", table, err)
	}
	return columns
}
