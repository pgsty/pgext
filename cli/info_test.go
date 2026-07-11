package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestExtensionInfoJSONContract(t *testing.T) {
	trusted := false
	info := &ExtensionInfo{
		ID:       1800,
		Name:     "vector",
		Pkg:      "pgvector",
		Category: "RAG",
		Version:  "0.8.5",
		Trusted:  &trusted,
		PGVer:    []string{"18", "17"},
		Requires: []string{"demo"},
		Extra:    JsonMap{"type": "standard"},
		EnDesc:   "vector data type",
		MTime:    "2026-07-10",
	}

	var out bytes.Buffer
	if err := writeExtensionInfoJSON(&out, info); err != nil {
		t.Fatalf("write JSON: %v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &decoded); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	gotKeys := make([]string, 0, len(decoded))
	for key := range decoded {
		gotKeys = append(gotKeys, key)
	}
	sort.Strings(gotKeys)
	wantKeys := []string{
		"category", "comment", "contrib", "deb_deps", "deb_pg", "deb_pkg", "deb_repo", "deb_ver",
		"en_desc", "extra", "has_bin", "has_lib", "id", "lang", "lead", "lead_ext", "license", "mtime",
		"name", "need_ddl", "need_load", "pg_ver", "pkg", "relocatable", "repo", "require_by", "requires",
		"rpm_deps", "rpm_pg", "rpm_pkg", "rpm_repo", "rpm_ver", "schemas", "see_also", "source", "state",
		"tags", "trusted", "url", "version", "zh_desc",
	}
	sort.Strings(wantKeys)
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Fatalf("JSON keys = %v, want %v", gotKeys, wantKeys)
	}
	if decoded["name"] != "vector" || decoded["pkg"] != "pgvector" || decoded["trusted"] != false {
		t.Fatalf("unexpected JSON values: %#v", decoded)
	}
	if decoded["relocatable"] != nil || decoded["tags"] != nil || decoded["contrib"] != nil {
		t.Fatalf("SQL NULL/absent values should remain explicit nulls: %#v", decoded)
	}
	if !strings.HasSuffix(out.String(), "}\n") {
		t.Fatalf("JSON output is not newline terminated: %q", out.String())
	}
}

func TestLoadExtensionInfoPreservesNullableValuesIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()

	var name string
	if err := QueryRowContext(ctx, "SELECT name FROM pgext.extension ORDER BY id LIMIT 1").Scan(&name); err != nil {
		t.Fatalf("select nullable extension fixture: %v", err)
	}
	if _, err := ExecSQLContext(ctx, `
		UPDATE pgext.extension
		SET contrib = NULL, lead = NULL, has_bin = NULL, has_lib = NULL,
		    need_ddl = NULL, need_load = NULL, tags = NULL, pg_ver = '{}'::text[]
		WHERE name = $1
	`, name); err != nil {
		t.Fatalf("set nullable extension fixture: %v", err)
	}

	info, err := LoadExtensionInfo(ctx, name)
	if err != nil {
		t.Fatalf("load nullable extension: %v", err)
	}
	for field, value := range map[string]*bool{
		"contrib": info.Contrib, "lead": info.Lead, "has_bin": info.HasBin,
		"has_lib": info.HasLib, "need_ddl": info.NeedDDL, "need_load": info.NeedLoad,
	} {
		if value != nil {
			t.Errorf("%s = %v, want nil", field, *value)
		}
	}
	if info.Tags != nil {
		t.Fatalf("NULL tags = %#v, want nil", info.Tags)
	}
	if info.PGVer == nil || len(info.PGVer) != 0 {
		t.Fatalf("empty pg_ver = %#v, want non-nil empty slice", info.PGVer)
	}
}

func TestWriteExtensionInfoKeepsLegacyTextLabels(t *testing.T) {
	info := &ExtensionInfo{
		Name:     "vector",
		Pkg:      "pgvector",
		Category: "RAG",
		Repo:     "PGDG",
		Lang:     "C",
		License:  "PostgreSQL",
		Version:  "0.8.5",
		EnDesc:   "vector data type",
		Requires: []string{"demo", "other"},
		URL:      "https://example.com/vector",
		Source:   "vector-0.8.5.tar.gz",
	}

	var out bytes.Buffer
	if err := writeExtensionInfo(&out, info); err != nil {
		t.Fatalf("write text: %v", err)
	}
	for _, fragment := range []string{
		"Extension Information",
		"Name:        vector",
		"Package:     pgvector",
		"Repository:  PGDG",
		"Description:\n  vector data type",
		"Requires:    demo, other",
		"Website:     https://example.com/vector",
		"Source:      vector-0.8.5.tar.gz",
	} {
		if !strings.Contains(out.String(), fragment) {
			t.Fatalf("text output missing %q:\n%s", fragment, out.String())
		}
	}
}

type failingExtensionWriter struct{}

func (failingExtensionWriter) Write([]byte) (int, error) {
	return 0, errors.New("broken writer")
}

func TestWriteExtensionInfoPropagatesWriterError(t *testing.T) {
	err := writeExtensionInfo(failingExtensionWriter{}, &ExtensionInfo{Name: "vector", Pkg: "pgvector"})
	if err == nil || !strings.Contains(err.Error(), "broken writer") {
		t.Fatalf("write error = %v, want broken writer", err)
	}
}

func TestExtensionInfoQueryUsesCurrentSchema(t *testing.T) {
	for _, column := range []string{"version", "url", "en_desc", "zh_desc", "mtime::text"} {
		if !strings.Contains(extensionInfoQuery, column) {
			t.Fatalf("extension query missing current column %q", column)
		}
	}
	for _, obsolete := range []string{" doc_url", " src_url", " created_at", " updated_at"} {
		if strings.Contains(extensionInfoQuery, obsolete) {
			t.Fatalf("extension query still uses obsolete column %q", obsolete)
		}
	}
}
