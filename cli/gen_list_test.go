package cli

import (
	"bytes"
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSortLicenseItemsUsesStableNameTieBreak(t *testing.T) {
	oneExtension := []*Extension{{}}
	items := []licenseItem{
		{name: "Zulu", exts: oneExtension, order: 999},
		{name: "beta", exts: oneExtension, order: 999},
		{name: "Alpha", exts: oneExtension, order: 999},
		{name: "alpha", exts: oneExtension, order: 999},
	}

	sortLicenseItems(items)

	got := make([]string, len(items))
	for i, item := range items {
		got[i] = item.name
	}
	if want := "Alpha,alpha,beta,Zulu"; strings.Join(got, ",") != want {
		t.Fatalf("license order = %q, want %q", strings.Join(got, ","), want)
	}
}

func TestGenerateLicenseListIsDeterministicForEqualRankLicenses(t *testing.T) {
	extensions := []*Extension{
		{ID: 4, Name: "zulu_ext", Pkg: "zulu_ext", License: sql.NullString{String: "Zulu", Valid: true}},
		{ID: 3, Name: "beta_ext", Pkg: "beta_ext", License: sql.NullString{String: "beta", Valid: true}},
		{ID: 2, Name: "alpha_upper_ext", Pkg: "alpha_upper_ext", License: sql.NullString{String: "Alpha", Valid: true}},
		{ID: 1, Name: "alpha_lower_ext", Pkg: "alpha_lower_ext", License: sql.NullString{String: "alpha", Valid: true}},
	}

	generate := func(name string, exts []*Extension) []byte {
		t.Helper()
		path := filepath.Join(t.TempDir(), name+".md")
		generator := NewListGenerator(&ExtensionCache{Extensions: exts}, filepath.Dir(path))
		if err := generator.GenerateLicenseList("en", path); err != nil {
			t.Fatalf("generate %s: %v", name, err)
		}
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		return content
	}

	forward := generate("forward", extensions)
	reverse := generate("reverse", []*Extension{extensions[3], extensions[2], extensions[1], extensions[0]})
	if !bytes.Equal(forward, reverse) {
		t.Fatal("license output changed when extension input order changed")
	}

	previous := -1
	for _, section := range []string{"alpha", "Alpha", "beta", "Zulu"} {
		index := bytes.Index(forward, []byte("\n## "+section+"\n"))
		if index < 0 {
			t.Fatalf("generated output is missing %s section", section)
		}
		if index <= previous {
			t.Fatalf("%s section appears out of deterministic extension ID order", section)
		}
		previous = index
	}
}

func TestGenerateCatalogContentDistinguishesUniverseFromPackaged(t *testing.T) {
	stats := &CatalogStats{
		ExtensionStats: []StatsRow{{Title: "ALL", Total: 572}},
		PackageStats:   []StatsRow{{Title: "ALL", Total: 405}},
	}
	generator := &ExtensionGenerator{}

	en := generator.generateCatalogContent(stats, nil, 2230, "en")
	for _, want := range []string{
		"packaged PostgreSQL extension catalog contains **572** extensions",
		"full PGEXT.CLOUD directory contains **2230** extensions",
	} {
		if !strings.Contains(en, want) {
			t.Fatalf("English catalog output is missing %q", want)
		}
	}

	zh := generator.generateCatalogContent(stats, nil, 2230, "zh")
	for _, want := range []string{
		"已打包扩展目录包含 **572** 个扩展",
		"总目录收录 **2230** 个扩展",
	} {
		if !strings.Contains(zh, want) {
			t.Fatalf("Chinese catalog output is missing %q", want)
		}
	}
}
