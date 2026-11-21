package cli

import (
	"context"
	"strings"
	"testing"

	"pgext/db"
)

// Uses only the explicitly disposable integration database, never data.pgext.
func TestPG19PackageCatalogIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	if _, err := ExecSQLContext(ctx, `
		-- Retired repositories need not be present in the bundled catalog.
		INSERT INTO pgext.repository (id, os, org, type)
		VALUES ('el7.x86_64.pigsty', 'el7.x86_64', 'pigsty', 'rpm');
		INSERT INTO pgext.apt (repo,package,version,architecture,filename,sha256)
		VALUES ('d13.aarch64.pgdg','postgresql-19-pgvector','0.8.7-1.pgdg13+1','arm64','pgvector19.deb',repeat('a',64)),
		       ('d13.aarch64.pgdg','postgresql-14-pgvector','0.8.7-1.pgdg13+1','arm64','pgvector14.deb',repeat('b',64));
		INSERT INTO pgext.dnf (repo,pkg_key,pkg_id,name,version,release,location_href)
		VALUES ('el9.x86_64.pigsty',1,'test-pg19','pgvector_19','0.8.7','1PGSTY.el9','pgvector19.rpm'),
		       ('el8.x86_64.pigsty',2,'el8-pg18','pgvector_18','0.8.7','1PGSTY.el8','pgvector18.rpm'),
		       ('el7.x86_64.pigsty',3,'retired-os','pgvector_15','0.8.7','1PGSTY.el7','retired.rpm');
	`); err != nil {
		t.Fatal(err)
	}
	g := NewBinGenerator(ctx)
	if _, err := g.Generate(); err != nil {
		t.Fatal(err)
	}
	if err := RecapMatrixContext(ctx); err != nil {
		t.Fatal(err)
	}
	for _, target := range []struct {
		pg   int
		os   string
		want string
	}{
		{19, "d13.aarch64", "AVAIL"},
		{19, "el9.x86_64", "AVAIL"},
		{18, "el8.x86_64", "AVAIL"},
		{19, "el8.x86_64", "N/A"},
	} {
		var state string
		if err := QueryRowContext(ctx, `SELECT state FROM pgext.pkg WHERE pkg='pgvector' AND pg=$1 AND os=$2`, target.pg, target.os).Scan(&state); err != nil {
			t.Fatal(err)
		}
		if state != target.want {
			t.Fatalf("PG%d %s: %s, want %s", target.pg, target.os, state, target.want)
		}
	}
	var count int
	for _, table := range []string{"bin", "pkg"} {
		if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext."+table+" WHERE pg=14 OR os LIKE 'el7.%'").Scan(&count); err != nil || count != 0 {
			t.Fatalf("retired targets in %s: %d rows, %v", table, count, err)
		}
	}
	var majors, platforms int
	if err := QueryRowContext(ctx, "SELECT (SELECT count(*) FROM pgext.active_pg), (SELECT count(*) FROM pgext.active_os)").Scan(&majors, &platforms); err != nil || majors != 5 || platforms != 16 {
		t.Fatalf("active matrix dimensions: %d x %d, %v", majors, platforms, err)
	}
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.summary WHERE pg19 IS NOT NULL AND pg14 IS NOT NULL").Scan(&count); err != nil || count != 6 {
		t.Fatalf("compatibility summary: %d rows, %v", count, err)
	}
	// Historical repository rows stay present but must not be fetched/scanned.
	f := &Fetcher{}
	repos, err := f.loadRepositories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(repos) == 0 {
		t.Fatal("no active repositories")
	}
	for _, r := range repos {
		if strings.HasPrefix(r.ID, "el7.") || strings.HasSuffix(r.ID, "pgdg14") || strings.HasSuffix(r.ID, "pgnf14") {
			t.Fatalf("retired repository scheduled: %s", r.ID)
		}
	}
	if _, err := ExecSQLContext(ctx, `INSERT INTO pgext.repo_data(id,data)
		VALUES ('el9.x86_64.pgdg15','active'), ('el8.x86_64.pigsty','active'), ('el9.x86_64.pgdg14','retired'), ('el7.x86_64.pigsty','retired')`); err != nil {
		t.Fatal(err)
	}
	parser := &Parser{ctx: ctx}
	cached, err := parser.loadRepositories()
	if err != nil || len(cached) != 2 || cached[0].ID != "el8.x86_64.pigsty" || cached[1].ID != "el9.x86_64.pgdg15" {
		t.Fatalf("parser selected retired caches: %+v, %v", cached, err)
	}
	scanner := &Scanner{}
	local, err := scanner.loadPigstyRepositories(ctx)
	if err != nil || len(local) == 0 {
		t.Fatalf("local repository selection: %v", err)
	}
	for _, r := range local {
		if strings.HasPrefix(r.ID, "el7.") {
			t.Fatalf("scanner selected retired platform: %s", r.ID)
		}
	}
}

func TestBuildTargetMigrationPreservesHistoryIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	if _, err := ExecSQLContext(ctx, `
		UPDATE pgext.pg SET active=true WHERE pg=14;
		UPDATE pgext.pg SET active=false WHERE pg=19;
		UPDATE pgext.os SET active=false WHERE os_code='el8';
		INSERT INTO pgext.pkg (pg,os,name,pkg,ext,state,version)
		VALUES (14,'el8.x86_64','pgvector_14','pgvector','vector','AVAIL','0.8.7');
	`); err != nil {
		t.Fatal(err)
	}
	var before string
	query := `SELECT md5(jsonb_agg(to_jsonb(p) ORDER BY pg,os,name)::text) FROM pgext.pkg p`
	if err := QueryRowContext(ctx, query).Scan(&before); err != nil {
		t.Fatal(err)
	}
	migration, err := db.ReadFile("build-targets.sql")
	if err != nil {
		t.Fatal(err)
	}
	var firstTime string
	for run := 0; run < 2; run++ {
		if _, err := ExecSQLContext(ctx, string(migration)); err != nil {
			t.Fatal(err)
		}
		var after, recapTime string
		if err := QueryRowContext(ctx, query).Scan(&after); err != nil || after != before {
			t.Fatalf("historical packages changed: %s != %s, %v", after, before, err)
		}
		if err := QueryRowContext(ctx, `SELECT recap_time::text FROM pgext.status WHERE id=0`).Scan(&recapTime); err != nil {
			t.Fatal(err)
		}
		if run == 0 {
			firstTime = recapTime
		} else if recapTime != firstTime {
			t.Fatal("no-op migration invalidated caches again")
		}
	}
	var cells int
	if err := QueryRowContext(ctx, `SELECT length(codes) FROM pgext.matrix WHERE pkg='pgvector'`).Scan(&cells); err != nil || cells != 80 {
		t.Fatalf("materialized matrix must contain 80 cells, got %d: %v", cells, err)
	}
}
