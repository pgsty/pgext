package cli

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

// dnfPrimarySchema mirrors the columns extractPackages reads from repodata primary.sqlite.
const dnfPrimarySchema = `
CREATE TABLE packages (
	pkgKey INTEGER PRIMARY KEY,
	pkgId TEXT,
	name TEXT,
	arch TEXT,
	version TEXT,
	epoch TEXT,
	release TEXT,
	summary TEXT,
	description TEXT,
	url TEXT,
	time_file INTEGER,
	time_build INTEGER,
	rpm_license TEXT,
	rpm_vendor TEXT,
	rpm_group TEXT,
	rpm_buildhost TEXT,
	rpm_sourcerpm TEXT,
	rpm_header_start INTEGER,
	rpm_header_end INTEGER,
	rpm_packager TEXT,
	size_package INTEGER,
	size_installed INTEGER,
	size_archive INTEGER,
	location_href TEXT,
	location_base TEXT,
	checksum_type TEXT
)`

// makeDNFPrimary builds a primary.sqlite file and returns its raw bytes.
func makeDNFPrimary(t *testing.T, inserts ...string) []byte {
	t.Helper()
	path := filepath.Join(t.TempDir(), "primary.sqlite")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if _, err := db.Exec(dnfPrimarySchema); err != nil {
		t.Fatalf("create schema: %v", err)
	}
	for _, stmt := range inserts {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatalf("insert: %v", err)
		}
	}
	if err := db.Close(); err != nil {
		t.Fatalf("close sqlite: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read sqlite file: %v", err)
	}
	return data
}

func TestDNFParserExtractPackages(t *testing.T) {
	data := makeDNFPrimary(t,
		`INSERT INTO packages (pkgKey, pkgId, name, arch, version, epoch, release,
			summary, url, time_build, rpm_license, rpm_sourcerpm,
			size_package, location_href, checksum_type)
		VALUES (1, 'abc123', 'pg_duckdb_18', 'x86_64', '1.1.2', '0', '1PGDG.rhel9',
			'DuckDB embedded in Postgres', 'https://github.com/duckdb/pg_duckdb', 1750000000,
			'MIT', 'pg_duckdb_18-1.1.2-1PGDG.rhel9.src.rpm',
			12345678, 'pg_duckdb_18-1.1.2-1PGDG.rhel9.x86_64.rpm', 'sha256')`,
		`INSERT INTO packages (pkgKey, pkgId, name) VALUES (2, 'def456', 'sparse-pkg')`,
	)

	tmp := filepath.Join(t.TempDir(), "extract.sqlite")
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		t.Fatalf("write temp sqlite: %v", err)
	}
	db, err := sql.Open("sqlite", "file:"+tmp+"?mode=ro")
	if err != nil {
		t.Fatalf("open sqlite ro: %v", err)
	}
	defer db.Close()

	parser := newDNFParser(context.Background(), liveTable("dnf"), false)
	packages, err := parser.extractPackages(db)
	if err != nil {
		t.Fatalf("extractPackages: %v", err)
	}
	if len(packages) != 2 {
		t.Fatalf("got %d packages, want 2", len(packages))
	}

	full := packages[0]
	if full.Name != "pg_duckdb_18" || full.PkgKey != 1 || full.PkgId != "abc123" {
		t.Fatalf("unexpected identity: %+v", full)
	}
	if !full.Version.Valid || full.Version.String != "1.1.2" {
		t.Fatalf("version = %+v, want 1.1.2", full.Version)
	}
	if !full.TimeBuild.Valid || full.TimeBuild.Int64 != 1750000000 {
		t.Fatalf("time_build = %+v, want 1750000000", full.TimeBuild)
	}
	if !full.SizePackage.Valid || full.SizePackage.Int64 != 12345678 {
		t.Fatalf("size_package = %+v, want 12345678", full.SizePackage)
	}

	sparse := packages[1]
	if sparse.Name != "sparse-pkg" {
		t.Fatalf("sparse name = %q", sparse.Name)
	}
	if sparse.Arch.Valid || sparse.Summary.Valid || sparse.TimeBuild.Valid || sparse.SizePackage.Valid {
		t.Fatalf("sparse row should scan as NULLs: %+v", sparse)
	}

	// mode=ro must reject writes
	if _, err := db.Exec(`INSERT INTO packages (pkgKey, pkgId, name) VALUES (3, 'x', 'y')`); err == nil {
		t.Fatal("write on mode=ro connection should fail")
	}
}

func TestDNFParserParseRepositoryEmpty(t *testing.T) {
	// An empty packages table returns before any PostgreSQL access,
	// exercising the temp-file write and read-only DSN open path.
	data := makeDNFPrimary(t)
	parser := newDNFParser(context.Background(), liveTable("dnf"), false)
	count, err := parser.ParseRepository("test-repo", data)
	if err != nil {
		t.Fatalf("ParseRepository: %v", err)
	}
	if count != 0 {
		t.Fatalf("count = %d, want 0", count)
	}
}
