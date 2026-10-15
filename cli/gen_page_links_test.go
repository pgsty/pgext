package cli

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCatalogRelationshipsLinkOnlyPublishedExtensions(t *testing.T) {
	cache := &ExtensionCache{ExtMap: map[string]*Extension{
		"vector": {Name: "vector"},
		"draft":  {Name: "draft", State: sql.NullString{Valid: true, String: "not-ready"}},
		"empty":  nil,
	}}
	g := NewExtensionGenerator(cache, "")
	for _, name := range []string{"source_only", "draft", "empty"} {
		t.Run(name, func(t *testing.T) {
			ext := &Extension{Requires: []string{"vector", name}, RequireBy: []string{name}, SeeAlso: []string{name}}
			got := g.generateRelationships(ext, []*Extension{{Name: name}})
			if !strings.Contains(got, ExtShortcode("vector", "")) {
				t.Fatalf("published extension must remain linked: %s", got)
			}
			if strings.Contains(got, ExtShortcode(name, "")) || strings.Count(got, "`"+name+"`") != 4 {
				t.Fatalf("all relationship roles must preserve non-page names as text: %s", got)
			}
		})
	}
	if got := g.extensionReference(""); got != "" {
		t.Fatalf("empty relationship = %q", got)
	}
	if got := NewExtensionGenerator(nil, "").extensionReference("source_only"); got != "`source_only`" {
		t.Fatalf("uncached relationship = %q", got)
	}
}

func TestCatalogLicenseNavigationUsesCanonicalTerms(t *testing.T) {
	g := NewListGenerator(&ExtensionCache{}, "")
	for _, locale := range []string{"en", "zh"} {
		t.Run(locale, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "license.md")
			if err := g.GenerateLicenseList(locale, path); err != nil {
				t.Fatal(err)
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			for _, term := range []string{"0BSD", "Artistic-2.0"} {
				if !strings.Contains(string(data), LicenseShortcode(term)) {
					t.Errorf("missing canonical license term %s", term)
				}
			}
			for _, absent := range []string{"BSD 0-Clause", "Artistic"} {
				if strings.Contains(string(data), LicenseShortcode(absent)) {
					t.Errorf("navigation links absent taxonomy %s", absent)
				}
			}
		})
	}
}

func TestCatalogInstallLinksPointToCurrentGuides(t *testing.T) {
	g := NewExtensionGenerator(&ExtensionCache{PGVersions: []int{17}}, "")
	got := g.generateInstallSection(&Extension{Name: "sample", Pkg: "sample", NeedLoad: true, NeedDDL: true})
	for _, action := range []string{"install", "config", "create"} {
		if !strings.Contains(got, "https://pigsty.io/docs/pgsql/ext/"+action+"/") {
			t.Errorf("missing current %s guide: %s", action, got)
		}
	}
	if strings.Contains(got, "ext.pgsty.com/usage/") {
		t.Fatal("removed usage route is still generated")
	}
}

func TestProductOverviewLinksUseCanonicalLicenseAnchors(t *testing.T) {
	cache := &ExtensionCache{}
	for name, generate := range map[string]func(string) error{
		"cc": NewCCListGenerator(cache, "").GenerateOverviewPage,
		"io": NewIOListGenerator(cache, "").GenerateOverviewPage,
	} {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "_index.md")
			if err := generate(path); err != nil {
				t.Fatal(err)
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			for _, anchor := range []string{"0bsd", "artistic10", "artistic20"} {
				if !strings.Contains(string(data), `href="/ext/license#`+anchor+`"`) {
					t.Errorf("missing canonical license anchor %s", anchor)
				}
			}
			for _, absent := range []string{"bsd0clause", "artistic"} {
				if strings.Contains(string(data), `href="/ext/license#`+absent+`"`) {
					t.Errorf("overview links absent license anchor %s", absent)
				}
			}
		})
	}
}
