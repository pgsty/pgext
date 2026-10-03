package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStubLintDoesNotRequireDatabase(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "example.md")
	if err := os.WriteFile(path, []byte("## 用法\n\n```sql\nSELECT 1;\n```\n"), 0644); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	previous := stubLintCmd.OutOrStdout()
	stubLintCmd.SetOut(&output)
	t.Cleanup(func() { stubLintCmd.SetOut(previous) })
	if err := stubLintCmd.RunE(stubLintCmd, []string{path}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "passed: 1 files") {
		t.Fatalf("unexpected lint output: %s", output.String())
	}
	if err := os.WriteFile(path, []byte("```\nresult\n```\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := stubLintCmd.RunE(stubLintCmd, []string{path}); err == nil || !strings.Contains(err.Error(), "MD110") {
		t.Fatalf("invalid stub passed CLI lint: %v", err)
	}
}
