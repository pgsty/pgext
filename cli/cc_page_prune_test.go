package cli

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCCPageGeneratorPrunesOnlyStaleGeneratedPages(t *testing.T) {
	outputDir := t.TempDir()
	extDir := filepath.Join(outputDir, "e")
	if err := os.MkdirAll(extDir, 0o755); err != nil {
		t.Fatal(err)
	}

	files := map[string]string{
		"_index.md":   "---\ntitle: Extension Index\n---\n",
		"vector.md":   "---\ntitle: \"vector\"\nlinkTitle: \"vector\"\n---\n",
		"spat.md":     "---\ntitle: \"spat\"\nlinkTitle: \"spat\"\n---\n",
		"handbook.md": "# Hand-written notes\n",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(extDir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	generator := NewCCPageGenerator(nil, outputDir, "")
	removed, err := generator.PruneStaleExtensionPages([]*Extension{{Name: "vector"}})
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"spat"}; !reflect.DeepEqual(removed, want) {
		t.Fatalf("removed = %v, want %v", removed, want)
	}
	if _, err := os.Stat(filepath.Join(extDir, "spat.md")); !os.IsNotExist(err) {
		t.Fatalf("stale generated page still exists: %v", err)
	}
	for _, name := range []string{"_index.md", "vector.md", "handbook.md"} {
		if _, err := os.Stat(filepath.Join(extDir, name)); err != nil {
			t.Fatalf("preserved page %s: %v", name, err)
		}
	}
}

func TestCCPageGeneratorRefusesEmptyCatalogPrune(t *testing.T) {
	generator := NewCCPageGenerator(nil, t.TempDir(), "")
	if _, err := generator.PruneStaleExtensionPages(nil); err == nil {
		t.Fatal("empty catalog prune unexpectedly succeeded")
	}
}
