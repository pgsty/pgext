package cli

import (
	"database/sql"
	"strings"
	"testing"
)

func TestInstallRepoIncludesBothPackageSuppliers(t *testing.T) {
	nullString := func(s string) sql.NullString { return sql.NullString{String: s, Valid: s != ""} }
	cases := []struct {
		name, repo, rpm, deb, command string
	}{
		{"pgdg", "PGDG", "PGDG", "PGDG", "pig repo add pgdg -u"},
		{"pgexporter_ext", "PGDG", "PGDG", "PIGSTY", "pig repo add pgsql -u"},
		{"pigsty_rpm", "PGDG", "PIGSTY", "PGDG", "pig repo add pgsql -u"},
		{"legacy_pgdg", "PGDG", "", "", "pig repo add pgdg -u"},
		{"mixed", "MIXED", "PGDG", "PIGSTY", "pig repo add pgsql -u"},
		{"pigsty", "PIGSTY", "PIGSTY", "PIGSTY", "pig repo add pgsql -u"},
		{"unknown", "", "", "", "pig repo add pgsql -u"},
	}
	cache := &ExtensionCache{PGVersions: []int{17}}
	generators := map[string]func(*Extension) string{
		"catalog": NewExtensionGenerator(cache, "").generateInstallSection,
		"io":      NewIOPageGenerator(cache, "", "").generateInstall,
		"cc":      NewCCPageGenerator(cache, "", "").generateInstall,
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ext := &Extension{Name: tc.name, Pkg: tc.name, Repo: nullString(tc.repo), RpmRepo: nullString(tc.rpm), DebRepo: nullString(tc.deb)}
			for name, generate := range generators {
				t.Run(name, func(t *testing.T) {
					got := generate(ext)
					if !strings.Contains(got, tc.command) {
						t.Fatalf("missing %q in install instructions: %s", tc.command, got)
					}
					if ext.Repo != nullString(tc.repo) {
						t.Fatal("generation changed curated supplier")
					}
				})
			}
		})
	}
}
