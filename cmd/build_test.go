package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMakeAliasesAndQueryMigration(t *testing.T) {
	for _, name := range []string{"make", "mak", "mk", "m"} {
		c, _, err := rootCmd.Find([]string{name})
		if err != nil || c.Name() != "make" {
			t.Fatalf("%s: %v %v", name, c, err)
		}
	}
	for _, pair := range [][2]string{{"pkg", "pkg"}, {"matrix", "matrix"}, {"ext", "info"}, {"info", "info"}} {
		c, _, err := rootCmd.Find([]string{pair[0]})
		if err != nil || c.Name() != pair[1] {
			t.Fatal(pair, err)
		}
	}
	if matrixCmd.PreRunE == nil {
		t.Fatal("matrix lost DB initialization")
	}
	c, _, _ := rootCmd.Find([]string{"pkg"})
	if c.PreRunE != nil {
		t.Fatal("builder depends on database")
	}
}

func TestBuildPlanWithoutDatabaseOrPig(t *testing.T) {
	t.Setenv("PGURL", "postgresql://invalid.invalid/unreachable")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	dir := t.TempDir()
	recipe := filepath.Join(dir, "pgvector")
	os.MkdirAll(recipe, 0755)
	os.WriteFile(filepath.Join(recipe, "Makefile"), []byte("PG_VERSIONS := 17 18\nTARBALL := pgvector-0.8.7.tar.gz\n"), 0644)
	for _, action := range []string{"get", "dep", "make", "pkg"} {
		c := newBuildCommand(action)
		var out, diag bytes.Buffer
		c.SetOut(&out)
		c.SetErr(&diag)
		c.SetArgs([]string{"vector", "pgvector", "--plan", "--platform", "u24.arm64", "--spec-dir", dir, "--pg", "18", "--json"})
		if err := c.Execute(); err != nil {
			t.Fatal(err)
		}
		var result map[string]any
		if err := json.Unmarshal(out.Bytes(), &result); err != nil {
			t.Fatalf("dirty stdout %q: %v", out.String(), err)
		}
		if result["action"] != action {
			t.Fatal(result)
		}
		if _, err := os.Stat(filepath.Join(dir, "SOURCES")); !os.IsNotExist(err) {
			t.Fatal("plan mutated recipes")
		}
		if strings.Contains(c.Example, "pgext make") {
			t.Fatal("examples must use mak")
		}
	}
}
