package repo

import (
	"context"
	"errors"
	"io"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testManager(t *testing.T, platform string) *Manager {
	t.Helper()
	version := ""
	if strings.HasPrefix(platform, "el9.") {
		version = "9.6"
	}
	m, err := NewManager(t.TempDir(), testPlatform(t, platform, version), testCatalog(t), io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	m.RunCommand = func(context.Context, []string) error { t.Fatal("unexpected host command"); return nil }
	return m
}

func execute(t *testing.T, m *Manager, action string, args []string, opts Options) *Result {
	t.Helper()
	r, err := m.Execute(context.Background(), action, args, opts)
	if err != nil {
		t.Fatalf("%s %v: %v", action, args, err)
	}
	return r
}

func snapshot(t *testing.T, m *Manager) map[string]*fileImage {
	t.Helper()
	files, err := m.scan()
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func TestManagementRoundTrip(t *testing.T) {
	for _, platform := range []string{"u24.arm64", "el9.amd64"} {
		t.Run(platform, func(t *testing.T) {
			m := testManager(t, platform)
			external := "etc/apt/sources.list.d/vendor.list"
			content := "deb https://example.org/apt stable main\n"
			if m.Platform.Family == "rpm" {
				external = "etc/yum.repos.d/vendor.repo"
				content = "[vendor]\nbaseurl=https://example.org/rpm\n"
			}
			original := &fileImage{Content: content, Mode: 0640}
			if err := m.write(external, original); err != nil {
				t.Fatal(err)
			}
			r := execute(t, m, "add", []string{"pgdg", "pigsty"}, Options{})
			if len(r.Changes) == 0 || r.Backup == "" || !r.Success {
				t.Fatal(r)
			}
			backupInfo, err := os.Stat(filepath.Join(m.Root, stateDir, "backups", r.Backup))
			if err != nil || backupInfo.Mode().Perm() != 0600 {
				t.Fatalf("backup permissions: %v %v", backupInfo, err)
			}
			first := snapshot(t, m)
			if !sameImage(first[external], original) {
				t.Fatal("changed external source")
			}
			again := execute(t, m, "add", []string{"pgdg", "pigsty"}, Options{})
			if len(again.Changes) != 0 || again.Backup != "" {
				t.Fatal("idempotent add wrote backup", again)
			}
			status := execute(t, m, "status", nil, Options{})
			if len(status.Installed) != len(first) {
				t.Fatal("incomplete status")
			}
			removed := execute(t, m, "rm", []string{"pigsty"}, Options{})
			if len(removed.Changes) != 2 {
				t.Fatal(removed)
			}
			restored := execute(t, m, "restore", []string{removed.Backup}, Options{})
			if restored.Backup == "" || !maps.EqualFunc(first, snapshot(t, m), sameImage) {
				t.Fatal("restore did not restore exact files")
			}
			execute(t, m, "set", []string{"docker"}, Options{})
			files := snapshot(t, m)
			if len(files) != 2 || !sameImage(files[external], original) {
				t.Fatal("set failed to replace only managed files", files)
			}
			execute(t, m, "rm", nil, Options{All: true})
			if files := snapshot(t, m); len(files) != 1 || !sameImage(files[external], original) {
				t.Fatal("rm --all removed external source", files)
			}
		})
	}
}

func TestReplaceAndRestoreAllAPTFormats(t *testing.T) {
	m := testManager(t, "u24.arm64")
	files := map[string]*fileImage{
		"etc/apt/sources.list":                  {Content: "deb http://ports.ubuntu.com/ubuntu-ports noble main\n", Mode: 0640},
		"etc/apt/sources.list.d/ubuntu.sources": {Content: "Types: deb\nURIs: http://ports.ubuntu.com/ubuntu-ports\nSuites: noble-updates\nComponents: main\nSigned-By: /usr/share/keyrings/ubuntu-archive-keyring.gpg\n", Mode: 0644},
		"etc/apt/sources.list.d/old.list":       {Content: "deb https://old.example.org stable main\n", Mode: 0600},
	}
	for path, image := range files {
		if err := m.write(path, image); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := m.Execute(context.Background(), "add", []string{"node"}, Options{}); err == nil || !strings.Contains(err.Error(), "conflicts") {
		t.Fatal("duplicate sources accepted", err)
	}
	r := execute(t, m, "set", []string{"all"}, Options{Replace: true})
	if len(snapshot(t, m)) != 7 {
		t.Fatal("replacement did not install all default modules")
	}
	execute(t, m, "restore", []string{r.Backup}, Options{})
	if !maps.EqualFunc(files, snapshot(t, m), sameImage) {
		t.Fatal("restore lost original content/modes")
	}
}

func TestDryRunAndInvalidSelectionLeaveNoFiles(t *testing.T) {
	m := testManager(t, "u24.arm64")
	r := execute(t, m, "add", nil, Options{DryRun: true})
	if len(r.Changes) != 7 || r.Backup != "" || !r.DryRun {
		t.Fatal(r)
	}
	for _, tt := range []struct {
		action string
		args   []string
		opts   Options
	}{
		{"add", []string{"pigsty", "typo"}, Options{}},
		{"add", []string{"pgdg18"}, Options{}},
		{"rm", nil, Options{}},
		{"rm", []string{"pgdg"}, Options{All: true}},
		{"add", []string{"local"}, Options{}},
		{"update", nil, Options{}},
		{"add", []string{"pgdg"}, Options{Update: true}},
		{"restore", []string{"../anything.json"}, Options{}},
	} {
		if _, err := m.Execute(context.Background(), tt.action, tt.args, tt.opts); err == nil {
			t.Fatal("accepted invalid operation", tt)
		}
	}
	entries, err := os.ReadDir(m.Root)
	if err != nil || len(entries) != 0 {
		t.Fatal("preview or invalid request created files", entries, err)
	}
}

func TestConflictDetectionForLegacyAndDeb822(t *testing.T) {
	for _, tt := range []struct {
		path, content string
	}{
		{"pgsql.list", "deb [signed-by=/keys/pgdg.asc] https://apt.postgresql.org/pub/repos/apt noble-pgdg main\n"},
		{"pgdg.sources", "# source\nTypes: deb deb-src\nURIs:\n https://apt.postgresql.org/pub/repos/apt/\nSuites: noble-pgdg\nComponents: main\nSigned-By: /keys/pgdg.asc\n"},
	} {
		t.Run(tt.path, func(t *testing.T) {
			m := testManager(t, "u24.amd64")
			if err := m.write("etc/apt/sources.list.d/"+tt.path, &fileImage{Content: tt.content, Mode: 0644}); err != nil {
				t.Fatal(err)
			}
			before := snapshot(t, m)
			if _, err := m.Execute(context.Background(), "add", []string{"pgdg"}, Options{}); err == nil || !strings.Contains(err.Error(), "conflicts") {
				t.Fatal("failed to detect conflicting source", err)
			}
			if !maps.EqualFunc(before, snapshot(t, m), sameImage) {
				t.Fatal("conflict changed files")
			}
		})
	}
	if keys := sourceKeys("x.sources", "Types: deb\nEnabled: no\nURIs: https://example.org\nSuites: stable\n", "deb"); len(keys) != 0 {
		t.Fatal("disabled deb822 source is active", keys)
	}
	if keys := sourceKeys("x.list", "# deb https://example.org stable main\ndeb-src https://example.org stable main", "deb"); len(keys) != 0 {
		t.Fatal("comments or source-only entries considered binary sources", keys)
	}
}

func TestRestoreRejectsChangedFiles(t *testing.T) {
	m := testManager(t, "u24.amd64")
	r := execute(t, m, "add", []string{"pigsty"}, Options{})
	path := repoPath("pigsty-pgsql", "deb")
	changed := &fileImage{Content: "# manual change\n", Mode: 0644}
	if err := m.write(path, changed); err != nil {
		t.Fatal(err)
	}
	before := snapshot(t, m)
	if _, err := m.Execute(context.Background(), "restore", []string{r.Backup}, Options{}); err == nil || !strings.Contains(err.Error(), "changed since backup") {
		t.Fatal("restore accepted drift", err)
	}
	if !maps.EqualFunc(before, snapshot(t, m), sameImage) {
		t.Fatal("restore overwrote a manual edit")
	}
}

func TestUpdateFailureRollsBackFiles(t *testing.T) {
	m := testManager(t, "u24.amd64")
	execute(t, m, "add", []string{"pigsty"}, Options{})
	before := snapshot(t, m)
	plan := execute(t, m, "set", []string{"pgdg"}, Options{DryRun: true})
	plan.Commands = [][]string{{"apt-get", "update"}}
	m.RunCommand = func(context.Context, []string) error { return errors.New("repository unavailable") }
	err := m.apply(context.Background(), plan, before)
	if err == nil || !plan.RolledBack || !strings.Contains(err.Error(), "repository unavailable") {
		t.Fatalf("update failure: %v %+v", err, plan)
	}
	if !maps.EqualFunc(before, snapshot(t, m), sameImage) {
		t.Fatal("rollback did not restore/remove changed files")
	}
}

func TestConcurrentDriftAndLock(t *testing.T) {
	m := testManager(t, "u24.amd64")
	before := snapshot(t, m)
	plan := execute(t, m, "add", []string{"pgdg"}, Options{DryRun: true})
	if err := m.write("etc/apt/sources.list", &fileImage{Content: "# concurrent source\n", Mode: 0644}); err != nil {
		t.Fatal(err)
	}
	if err := m.apply(context.Background(), plan, before); err == nil || !strings.Contains(err.Error(), "changed during planning") {
		t.Fatal("concurrent write ignored", err)
	}
	unlock, err := m.lock()
	if err != nil {
		t.Fatal(err)
	}
	defer unlock()
	if _, err := m.Execute(context.Background(), "add", []string{"pgdg"}, Options{}); err == nil || !strings.Contains(err.Error(), "operation is running") {
		t.Fatal("concurrent operation accepted", err)
	}
}

func TestCompatibleHostPlatformAllowsVendorAndMinorDifferences(t *testing.T) {
	host := Platform{OS: "el9", Arch: "x86_64", Family: "rpm", Version: "9.6", Vendor: "rhel"}
	requested := Platform{OS: "el9", Arch: "x86_64", Family: "rpm", Version: "9", Vendor: "rocky"}
	if !compatibleHostPlatform(host, requested) {
		t.Fatal("generic EL9 target should match an EL9 compatible host")
	}
	requested.Version = "10"
	if compatibleHostPlatform(host, requested) {
		t.Fatal("EL10 target should not match EL9 host")
	}
	requested = Platform{OS: "u24", Arch: "arm64", Family: "deb", Version: "24", Codename: "noble"}
	host = Platform{OS: "u24", Arch: "arm64", Family: "deb", Version: "24.04", Codename: "noble", Vendor: "ubuntu"}
	if !compatibleHostPlatform(host, requested) {
		t.Fatal("generic Ubuntu 24 target should match the 24.04 host")
	}
}

func TestSymlinksAndUnownedFilesAreNeverOverwritten(t *testing.T) {
	for _, parent := range []bool{false, true} {
		t.Run(map[bool]string{true: "parent", false: "file"}[parent], func(t *testing.T) {
			m := testManager(t, "u24.amd64")
			external := t.TempDir()
			target := filepath.Join(external, "sentinel")
			if err := os.WriteFile(target, []byte("keep"), 0600); err != nil {
				t.Fatal(err)
			}
			link := filepath.Join(m.Root, repoPath("pgdg", "deb"))
			if parent {
				link, target = filepath.Dir(link), external
			}
			if err := os.MkdirAll(filepath.Dir(link), 0755); err != nil {
				t.Fatal(err)
			}
			if err := os.Symlink(target, link); err != nil {
				t.Fatal(err)
			}
			if _, err := m.Execute(context.Background(), "add", []string{"pgdg"}, Options{Replace: true}); err == nil || !strings.Contains(err.Error(), "symlink") {
				t.Fatal("followed symlink", err)
			}
			data, _ := os.ReadFile(filepath.Join(external, "sentinel"))
			if string(data) != "keep" {
				t.Fatal("symlink target modified")
			}
		})
	}
	m := testManager(t, "u24.amd64")
	path := repoPath("pgdg", "deb")
	if err := m.write(path, &fileImage{Content: "# user owned\n", Mode: 0644}); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Execute(context.Background(), "add", []string{"pgdg"}, Options{}); err == nil {
		t.Fatal("overwrote unowned pgext-prefixed file")
	}
	execute(t, m, "rm", []string{"pgdg"}, Options{})
	if len(snapshot(t, m)) != 1 {
		t.Fatal("rm removed unowned file")
	}
}
