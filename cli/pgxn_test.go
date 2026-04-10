package cli

import "testing"

func TestPGXNIndexRegex(t *testing.T) {
	html := `
<tr><td class='name'><a href='/src/acl/'>acl/</a></td></tr>
<tr><td class='name'><a href='/src/e-maj/'>e-maj/</a></td></tr>
<tr><td class='name'><a href='/src/vector/'>vector/</a></td></tr>
`

	matches := pgxnIndexHrefRE.FindAllStringSubmatch(html, -1)
	if len(matches) != 3 {
		t.Fatalf("expected 3 matches, got %d", len(matches))
	}
	if matches[1][1] != "e-maj" {
		t.Fatalf("expected e-maj, got %q", matches[1][1])
	}
}

func TestNormalizeURL(t *testing.T) {
	cases := map[string]string{
		"https://github.com/pgvector/pgvector":        "github.com/pgvector/pgvector",
		"http://github.com/pgvector/pgvector.git":     "github.com/pgvector/pgvector",
		"https://github.com/pgvector/pgvector/":       "github.com/pgvector/pgvector",
		"https://github.com/pgvector/pgvector/issues": "github.com/pgvector/pgvector",
	}

	for input, expected := range cases {
		if actual := normalizeURL(input); actual != expected {
			t.Fatalf("normalizeURL(%q) expected %q, got %q", input, expected, actual)
		}
	}
}

func TestMatchPGXNRecordByProvides(t *testing.T) {
	idx := &catalogIndex{
		byPkg:       make(map[string][]CatalogRecord),
		byName:      make(map[string][]CatalogRecord),
		byCanonical: make(map[string][]CatalogRecord),
		byURL:       make(map[string][]CatalogRecord),
	}
	row := CatalogRecord{ID: 1050, Name: "emaj", Pkg: "emaj", Lead: true}
	idx.byPkg["emaj"] = []CatalogRecord{row}
	idx.byName["emaj"] = []CatalogRecord{row}
	idx.byCanonical[canonicalName("emaj")] = []CatalogRecord{row}

	record := &PgxnRecord{
		Dist: "e-maj",
		RawDist: map[string]interface{}{
			"name": "E-Maj",
			"provides": map[string]interface{}{
				"emaj": map[string]interface{}{"version": "4.7.1"},
			},
		},
	}

	match := matchPGXNRecord(record, idx)
	if !match.Matched {
		t.Fatal("expected matched record")
	}
	if match.Row.Pkg != "emaj" {
		t.Fatalf("expected pkg emaj, got %q", match.Row.Pkg)
	}
	if match.Reason != "provide.name" {
		t.Fatalf("expected provide.name, got %q", match.Reason)
	}
}

func TestMatchPGXNRecordByURL(t *testing.T) {
	idx := &catalogIndex{
		byPkg:       make(map[string][]CatalogRecord),
		byName:      make(map[string][]CatalogRecord),
		byCanonical: make(map[string][]CatalogRecord),
		byURL:       make(map[string][]CatalogRecord),
	}
	row := CatalogRecord{ID: 2580, Name: "vector", Pkg: "vector", URL: "https://github.com/pgvector/pgvector", Lead: true}
	idx.byURL[normalizeURL(row.URL)] = []CatalogRecord{row}

	record := &PgxnRecord{
		Dist: "vector",
		RawDist: map[string]interface{}{
			"resources": map[string]interface{}{
				"repository": map[string]interface{}{
					"url": "https://github.com/pgvector/pgvector.git",
				},
			},
		},
	}

	match := matchPGXNRecord(record, idx)
	if !match.Matched {
		t.Fatal("expected matched record")
	}
	if match.Reason != "url" && match.Reason != "pkg" && match.Reason != "name" {
		t.Fatalf("unexpected reason %q", match.Reason)
	}
}
