package cli

import (
	"database/sql"
	"strings"
	"testing"
)

func TestCCExtRefLinksOnlyGeneratedCatalogPages(t *testing.T) {
	cache := &ExtensionCache{
		ExtMap: map[string]*Extension{
			"vector": {Name: "vector"},
			"draft": {
				Name:  "draft",
				State: sql.NullString{Valid: true, String: "not-ready"},
			},
		},
	}

	tests := []struct {
		name string
		want string
	}{
		{name: "vector", want: "[`vector`](/ext/e/vector)"},
		{name: "draft", want: "`draft`"},
		{name: "source_only", want: "`source_only`"},
		{name: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CCExtRef(cache, tt.name); got != tt.want {
				t.Fatalf("CCExtRef(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestPageRelationshipsKeepNonPageEntitiesAsText(t *testing.T) {
	cache := &ExtensionCache{
		ExtMap: map[string]*Extension{
			"vector": {Name: "vector"},
		},
	}
	ext := &Extension{
		Requires:  []string{"vector"},
		SeeAlso:   []string{"clickhouse_fdw"},
		RequireBy: []string{"source_only"},
	}

	for name, content := range map[string]string{
		"io": NewIOPageGenerator(cache, "", "").generateRelationships(ext, nil),
		"cc": NewCCPageGenerator(cache, "", "").generateRelationships(ext, nil),
	} {
		t.Run(name, func(t *testing.T) {
			if !strings.Contains(content, "[`vector`](/ext/e/vector)") {
				t.Fatalf("generated relationships omit packaged link: %s", content)
			}
			if !strings.Contains(content, "`clickhouse_fdw`") ||
				strings.Contains(content, "/ext/e/clickhouse_fdw") {
				t.Fatalf("generated relationships link non-page see_also target: %s", content)
			}
			if !strings.Contains(content, "`source_only`") ||
				strings.Contains(content, "/ext/e/source_only") {
				t.Fatalf("generated relationships link non-page reverse dependency: %s", content)
			}
		})
	}
}
