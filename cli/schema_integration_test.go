package cli

import (
	"bytes"
	"context"
	"encoding/csv"
	"os"
	"testing"

	"pgext/db"
)

// TestInitSchemaAtomicIntegration requires an explicitly disposable database.
// Example:
//
//	PGEXT_TEST_PGURL='postgresql://postgres@/pgext_test?host=/tmp' go test ./cli -run TestInitSchemaAtomicIntegration -v
func TestInitSchemaAtomicIntegration(t *testing.T) {
	pgurl := os.Getenv("PGEXT_TEST_PGURL")
	if pgurl == "" {
		t.Skip("set PGEXT_TEST_PGURL to an explicitly disposable database")
	}
	if err := InitDB(pgurl); err != nil {
		t.Fatalf("InitDB: %v", err)
	}
	t.Cleanup(func() {
		_, _ = ExecContext(context.Background(), "DROP SCHEMA IF EXISTS pgext CASCADE")
		CloseDB()
		DB = nil
	})

	if err := InitSchema(true); err != nil {
		t.Fatalf("fresh InitSchema: %v", err)
	}
	extensionRows := embeddedCSVRowCount(t, "extension")
	universeRows := embeddedCSVRowCount(t, "universe")
	assertCatalogCounts(t, extensionRows, universeRows, 0)

	ctx := context.Background()
	if _, err := ExecContext(ctx, `UPDATE pgext.status SET extra = '{"atomic_test": true}' WHERE id = 0`); err != nil {
		t.Fatalf("set rollback sentinel: %v", err)
	}

	schemaSQL, err := db.GetSchema()
	if err != nil {
		t.Fatal(err)
	}
	files, err := db.GetCSVFiles()
	if err != nil {
		t.Fatal(err)
	}
	for i := range files {
		if files[i].Name == "universe" {
			files[i].Content = []byte("id,name\ninvalid,invalid\n")
		}
	}
	if _, err := initializeSchema(ctx, true, schemaSQL, files); err == nil {
		t.Fatal("expected invalid universe CSV to fail initialization")
	}

	var sentinel bool
	if err := QueryRowContext(ctx, "SELECT (extra->>'atomic_test')::boolean FROM pgext.status WHERE id = 0").Scan(&sentinel); err != nil {
		t.Fatalf("read rollback sentinel: %v", err)
	}
	if !sentinel {
		t.Fatal("failed force initialization did not restore the previous schema")
	}
	assertCatalogCounts(t, extensionRows, universeRows, 0)

	if err := LoadTable("universe", "", RegionChina); err != nil {
		t.Fatalf("LoadTable(universe) from embedded CSV: %v", err)
	}
	assertCatalogCounts(t, extensionRows, universeRows, 0)
}

// TestEnsureUniverseMigrationIntegration requires an explicitly disposable database.
// It verifies the non-destructive upgrade path for schemas created before the
// universe catalog was embedded.
func TestEnsureUniverseMigrationIntegration(t *testing.T) {
	pgurl := os.Getenv("PGEXT_TEST_PGURL")
	if pgurl == "" {
		t.Skip("set PGEXT_TEST_PGURL to an explicitly disposable database")
	}
	if err := InitDB(pgurl); err != nil {
		t.Fatalf("InitDB: %v", err)
	}
	t.Cleanup(func() {
		_, _ = ExecContext(context.Background(), "DROP SCHEMA IF EXISTS pgext CASCADE")
		CloseDB()
		DB = nil
	})

	if err := InitSchema(true); err != nil {
		t.Fatalf("fresh InitSchema: %v", err)
	}
	extensionRows := embeddedCSVRowCount(t, "extension")
	universeRows := embeddedCSVRowCount(t, "universe")
	ctx := context.Background()

	if _, err := ExecContext(ctx, "DROP TABLE pgext.universe"); err != nil {
		t.Fatalf("drop universe to simulate old schema: %v", err)
	}
	universeDDL, err := db.GetUniverseSchema()
	if err != nil {
		t.Fatal(err)
	}
	invalidCSV := db.CSVFile{Name: "universe", Content: []byte("id,name\ninvalid,invalid\n")}
	if _, _, err := ensureUniverse(ctx, universeDDL, invalidCSV); err == nil {
		t.Fatal("expected invalid universe CSV to fail migration")
	}

	var exists bool
	if err := QueryRowContext(ctx, "SELECT to_regclass('pgext.universe') IS NOT NULL").Scan(&exists); err != nil {
		t.Fatalf("check failed migration: %v", err)
	}
	if exists {
		t.Fatal("failed universe migration left a partial table behind")
	}

	if err := InitSchema(false); err != nil {
		t.Fatalf("upgrade existing schema: %v", err)
	}
	assertCatalogCounts(t, extensionRows, universeRows, 0)

	const sentinel = "universe migration no-op sentinel"
	if _, err := ExecContext(ctx, `
UPDATE pgext.universe SET comment = $1
WHERE id = (SELECT min(id) FROM pgext.universe)`, sentinel); err != nil {
		t.Fatalf("set universe sentinel: %v", err)
	}
	if err := InitSchema(false); err != nil {
		t.Fatalf("repeat existing schema initialization: %v", err)
	}
	var sentinelCount int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.universe WHERE comment = $1", sentinel).Scan(&sentinelCount); err != nil {
		t.Fatalf("read universe sentinel: %v", err)
	}
	if sentinelCount != 1 {
		t.Fatalf("existing universe was reloaded: sentinel count = %d, want 1", sentinelCount)
	}
}

func embeddedCSVRowCount(t *testing.T, table string) int {
	t.Helper()
	file, err := db.GetCSVFile(table)
	if err != nil {
		t.Fatal(err)
	}
	records, err := csv.NewReader(bytes.NewReader(file.Content)).ReadAll()
	if err != nil {
		t.Fatalf("read embedded %s.csv: %v", table, err)
	}
	return len(records) - 1
}

func assertCatalogCounts(t *testing.T, wantExtension, wantUniverse, wantMissing int) {
	t.Helper()
	var extension, universe, missing int
	err := QueryRowContext(context.Background(), `
SELECT
    (SELECT count(*) FROM pgext.extension),
    (SELECT count(*) FROM pgext.universe),
    (SELECT count(*) FROM pgext.extension e LEFT JOIN pgext.universe u USING (name) WHERE u.name IS NULL)`).
		Scan(&extension, &universe, &missing)
	if err != nil {
		t.Fatalf("query catalog counts: %v", err)
	}
	if extension != wantExtension || universe != wantUniverse || missing != wantMissing {
		t.Fatalf("catalog counts = (%d, %d, %d), want (%d, %d, %d)",
			extension, universe, missing, wantExtension, wantUniverse, wantMissing)
	}
}
