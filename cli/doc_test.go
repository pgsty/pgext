package cli

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSafeStubName(t *testing.T) {
	for _, ok := range []string{"vector", "uuid-ossp", "pg_stat_statements", "a.b"} {
		if err := SafeStubName(ok); err != nil {
			t.Errorf("SafeStubName(%q) unexpected error: %v", ok, err)
		}
	}
	for _, bad := range []string{"", ".", "..", "../bad", "bad/name", `bad\name`, "bad\x00name"} {
		if err := SafeStubName(bad); err == nil {
			t.Errorf("SafeStubName(%q) expected error, got nil", bad)
		}
	}
}

func TestCollectStubDocs(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "vector.md"), []byte("## Usage\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "ignore.txt"), []byte("nope"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "subdir.md"), 0755); err != nil {
		t.Fatal(err)
	}

	docs, err := CollectStubDocs(dir)
	if err != nil {
		t.Fatalf("CollectStubDocs: %v", err)
	}
	want := map[string]string{"vector": "## Usage\n"}
	if !reflect.DeepEqual(docs, want) {
		t.Fatalf("docs = %#v, want %#v", docs, want)
	}
}

func TestPrepareStubStage(t *testing.T) {
	universe := []string{"alpha", "beta"}
	en := map[string]string{"alpha": "en-a", "beta": "en-b", "orphan": "x"}
	zh := map[string]string{"alpha": "zh-a", "beta": "zh-b", "extra_zh": "y"}

	stats, rows, err := prepareStubStage(universe, en, zh, true)
	if err != nil {
		t.Fatalf("prepareStubStage: %v", err)
	}
	if stats.UniverseCount != 2 || len(rows) != 2 || !stats.DryRun {
		t.Fatalf("unexpected stats/rows: %+v, len=%d", stats, len(rows))
	}
	if !reflect.DeepEqual(stats.ExtraEN, []string{"orphan"}) || !reflect.DeepEqual(stats.ExtraZH, []string{"extra_zh"}) {
		t.Fatalf("unexpected extras: EN=%v ZH=%v", stats.ExtraEN, stats.ExtraZH)
	}

	// Missing or blank ZH doc must fail cleanly before any mutation.
	_, _, err = prepareStubStage(universe, en, map[string]string{"alpha": "zh-a", "beta": "   "}, false)
	if err == nil || !strings.Contains(err.Error(), "ZH=1 (beta)") {
		t.Fatalf("expected missing ZH error, got: %v", err)
	}
}

func TestDocLoadLegacyForeignKeysIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	tx, err := DB.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	if _, err := tx.Exec(ctx, `
		ALTER TABLE pgext.doc ADD CONSTRAINT doc_ext_fkey
			FOREIGN KEY (ext) REFERENCES pgext.extension (name) ON DELETE CASCADE;
		ALTER TABLE pgext.doc ADD CONSTRAINT custom_doc_id_fkey
			FOREIGN KEY (id) REFERENCES pgext.universe (id);
	`); err != nil {
		t.Fatal(err)
	}
	if err := dropLegacyDocForeignKeys(ctx, tx); err == nil || !strings.Contains(err.Error(), "custom_doc_id_fkey") {
		t.Fatalf("expected custom foreign key to block the migration, got %v", err)
	}
	countForeignKeys := func(want int) {
		t.Helper()
		var count int
		if err := tx.QueryRow(ctx, `SELECT count(*) FROM pg_constraint
			WHERE conrelid = 'pgext.doc'::regclass AND contype = 'f'`).Scan(&count); err != nil || count != want {
			t.Fatalf("pgext.doc has %d foreign keys, want %d: %v", count, want, err)
		}
	}
	countForeignKeys(2)
	if _, err := tx.Exec(ctx, "ALTER TABLE pgext.doc DROP CONSTRAINT custom_doc_id_fkey"); err != nil {
		t.Fatal(err)
	}
	if err := dropLegacyDocForeignKeys(ctx, tx); err != nil {
		t.Fatal(err)
	}
	countForeignKeys(0)
}
