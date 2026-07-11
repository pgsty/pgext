package cli

import (
	"reflect"
	"testing"
)

func TestEmbeddedTableRegistry(t *testing.T) {
	want := []string{"pg", "os", "category", "repository", "extension", "universe"}
	if got := EmbeddedTableNames(); !reflect.DeepEqual(got, want) {
		t.Fatalf("EmbeddedTableNames() = %v, want %v", got, want)
	}
	if !IsEmbeddedTable("extension") || !IsEmbeddedTable("ext") || !IsEmbeddedTable("universe") {
		t.Fatal("expected extension, ext, and universe to be supported")
	}
	if IsEmbeddedTable("all_extensions") || IsEmbeddedTable("apt") {
		t.Fatal("unsupported tables must be rejected")
	}
}

func TestLoadPackageCatalogInvalidationScope(t *testing.T) {
	for _, table := range []string{"pg", "os", "repository", "extension"} {
		if !loadInvalidatesPackageCatalog(table) {
			t.Errorf("loading %s should invalidate pgext.pkg", table)
		}
	}
	for _, table := range []string{"category", "universe"} {
		if loadInvalidatesPackageCatalog(table) {
			t.Errorf("loading %s unexpectedly invalidates pgext.pkg", table)
		}
	}
}
