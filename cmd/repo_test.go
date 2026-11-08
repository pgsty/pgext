package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"pgext/cli/repo"
	"strings"
	"testing"
)

func runRepoTest(args ...string) (string, error) {
	c := newRepoCommand()
	var out, diag bytes.Buffer
	c.SetOut(&out)
	c.SetErr(&diag)
	c.SetArgs(args)
	err := c.Execute()
	return out.String(), err
}

func TestRepoCommandsNeedNoDatabase(t *testing.T) {
	t.Setenv("PGURL", "postgresql://invalid.invalid/unreachable")
	for _, args := range [][]string{
		{}, {"list", "all", "--json"},
		{"list", "pgdg", "--platform", "el9.arm64", "--json"},
		{"info", "pigsty", "--platform", "u24.arm64"},
		{"show", "pgdg", "--platform", "el9.arm64", "--os-version", "9.6", "--mirror"},
		{"add", "pgdg,pigsty", "--platform", "u24.arm64", "--dry-run", "--root", t.TempDir(), "--json"},
		{"status", "--platform", "u24.arm64", "--root", t.TempDir(), "--json"},
	} {
		out, err := runRepoTest(args...)
		if err != nil || out == "" {
			t.Fatalf("%v: %q %v", args, out, err)
		}
		if len(args) > 0 && args[len(args)-1] == "--json" {
			if !json.Valid([]byte(out)) {
				t.Fatalf("dirty JSON output: %q", out)
			}
		}
	}
	c := newRepoCommand()
	summary, _, err := c.Find([]string{"summary"})
	if err != nil || summary.PreRunE == nil || summary.PostRun == nil {
		t.Fatal("database summary lost its connection lifecycle")
	}
}

func TestRepoCLIStagingRoundTrip(t *testing.T) {
	root := t.TempDir()
	base := []string{"--platform", "u24.arm64", "--root", root, "--json"}
	run := func(args ...string) repo.Result {
		t.Helper()
		out, err := runRepoTest(append(args, base...)...)
		if err != nil {
			t.Fatalf("%v: %v %s", args, err, out)
		}
		var r repo.Result
		if err := json.Unmarshal([]byte(out), &r); err != nil || !r.Success {
			t.Fatal(out, err)
		}
		return r
	}
	plan := run("add", "pigsty", "--plan", "--mirror")
	if !plan.DryRun || len(plan.Changes) != 2 || !strings.Contains(plan.Changes[0].After.Content, "repo.pigsty.cc") {
		t.Fatal(plan)
	}
	r := run("add", "pigsty", "--mirror")
	if r.Backup == "" {
		t.Fatal("missing backup")
	}
	run("restore", r.Backup)
	files, err := filepath.Glob(filepath.Join(root, "etc/apt/sources.list.d/*.list"))
	if err != nil || len(files) != 0 {
		t.Fatal("restore failed", files, err)
	}
	if entries, err := os.ReadDir(filepath.Join(root, "var/lib/pgext/repo/backups")); err != nil || len(entries) != 2 {
		t.Fatal("missing restore backup", entries, err)
	}
}

func TestRepoCLIRejectsInvalidOptions(t *testing.T) {
	for _, args := range [][]string{
		{"show", "pgdg", "--platform", "u24.arm64", "--mirror", "--region", "europe"},
		{"rm"}, {"update", "unexpected"}, {"restore"},
		{"add", "pigsty", "--platform", "u24.arm64", "--update", "--root", t.TempDir()},
		{"set", "pigsty", "--platform", "u24.arm64", "--root", t.TempDir()},
	} {
		if out, err := runRepoTest(args...); err == nil {
			t.Fatalf("accepted %v: %s", args, out)
		}
	}
}
