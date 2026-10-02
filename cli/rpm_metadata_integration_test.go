package cli

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRPMPrimaryCacheAndParseIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	var repoID string
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type='rpm' ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
		t.Fatal(err)
	}
	if _, err := ExecSQLContext(ctx, `INSERT INTO pgext.repo_data (id,data,etag,extra) VALUES ($1,$2,'legacy','{"retained":"yes"}')`, repoID, []byte("old payload")); err != nil {
		t.Fatal(err)
	}
	compressed := compressPrimaryFixture(t, "zstd")
	primary := &RepoMDData{}
	primary.Location.Href = "repodata/primary.xml.zst"
	data, extra, err := decodeRPMPrimary(ctx, compressed, primary, "https://example.org/primary.xml.zst")
	if err != nil {
		t.Fatal(err)
	}
	fetcher := NewFetcher(FetchOptions{})
	if err := fetcher.saveMetadata(ctx, &FetchResult{Repository: &RepoMetadata{ID: repoID}, Data: data, Extra: extra}); err != nil {
		t.Fatal(err)
	}
	var stored []byte
	var raw []byte
	var lastMod sql.NullTime
	if err := QueryRowContext(ctx, "SELECT data,extra,last_modified FROM pgext.repo_data WHERE id=$1", repoID).Scan(&stored, &raw, &lastMod); err != nil {
		t.Fatal(err)
	}
	var descriptor map[string]any
	if err := json.Unmarshal(raw, &descriptor); err != nil {
		t.Fatal(err)
	}
	if string(stored) != primaryXMLFixture || descriptor["format"] != rpmPrimaryFormat || descriptor["retained"] != "yes" || lastMod.Valid {
		t.Fatalf("incorrect cache replacement: %v, lastMod=%v", descriptor, lastMod)
	}
	repos, err := fetcher.loadRepositories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, repo := range repos {
		if repo.ID == repoID {
			found = true
			if repo.CachedPrimary == nil || repo.CachedDataSHA256.String != extra.ContentSHA256 {
				t.Fatalf("cached descriptor did not round-trip: %+v", repo)
			}
		}
	}
	if !found {
		t.Fatal("refreshed RPM repository missing from cache query")
	}
	stage, err := newPackageStaging(ctx, true, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(stage.Drop)
	parser := newDNFParser(ctx, stage.dnf)
	if n, err := parser.ParseRepository(repoID, stored); err != nil || n != 1 {
		t.Fatalf("parse cached primary XML: %d, %v", n, err)
	}
	var key int
	var name, epoch, version, release, checksumType string
	var fileTime int64
	if err := QueryRowContext(ctx, fmt.Sprintf("SELECT pkg_key,name,epoch,version,release,checksum_type,time_file FROM %s WHERE repo=$1", stage.dnf), repoID).Scan(&key, &name, &epoch, &version, &release, &checksumType, &fileTime); err != nil {
		t.Fatal(err)
	}
	if key != 1 || name != "orafce_18" || epoch != "0" || version != "4.16.12" || release != "1.el9" || checksumType != "sha256" || fileTime != 0 {
		t.Fatalf("incorrect staged XML package: %d %s %s:%s-%s %s time=%d", key, name, epoch, version, release, checksumType, fileTime)
	}
	var live int
	if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext.dnf").Scan(&live); err != nil || live != 0 {
		t.Fatalf("staged XML parsing changed the live catalog: %d, %v", live, err)
	}
}

func TestInvalidPrimaryXMLPreservesCacheIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	var repoID string
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type='rpm' ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-None-Match") != "" || r.Header.Get("If-Modified-Since") != "" {
			t.Error("old cache validators were sent for primary XML")
		}
		if r.URL.Path == "/repodata/repomd.xml" {
			_, _ = w.Write([]byte(`<repomd><data type="primary"><location href="repodata/primary.xml"/></data></repomd>`))
		} else {
			_, _ = w.Write([]byte(`<metadata packages="1">`))
		}
	}))
	defer server.Close()
	if _, err := ExecSQLContext(ctx, "UPDATE pgext.repository SET default_meta=NULL,mirror_meta=NULL"); err != nil {
		t.Fatal(err)
	}
	if _, err := ExecSQLContext(ctx, "UPDATE pgext.repository SET default_meta=$2 WHERE id=$1", repoID, server.URL+"/repodata/repomd.xml"); err != nil {
		t.Fatal(err)
	}
	old := []byte("legacy payload")
	if _, err := ExecSQLContext(ctx, `INSERT INTO pgext.repo_data (id,data,etag) VALUES ($1,$2,'legacy')`, repoID, old); err != nil {
		t.Fatal(err)
	}
	if err := NewFetcher(FetchOptions{}).FetchAll(ctx); err == nil {
		t.Fatal("accepted truncated primary XML")
	}
	var stored []byte
	var etag string
	if err := QueryRowContext(ctx, "SELECT data,etag FROM pgext.repo_data WHERE id=$1", repoID).Scan(&stored, &etag); err != nil {
		t.Fatal(err)
	}
	if string(stored) != string(old) || etag != "legacy" {
		t.Fatal("failed XML refresh replaced the old cache")
	}
	var fetched sql.NullTime
	if err := QueryRowContext(ctx, "SELECT fetch_time FROM pgext.status WHERE id=0").Scan(&fetched); err != nil || fetched.Valid {
		t.Fatalf("failed refresh advanced the accepted timestamp: %v, %v", fetched, err)
	}
}

func TestLegacyRPMCacheDoesNotPublishIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	seedStagingIntegrationRows(t, ctx)
	var repoID string
	if err := QueryRowContext(ctx, "SELECT id FROM pgext.repository WHERE type='rpm' ORDER BY id LIMIT 1").Scan(&repoID); err != nil {
		t.Fatal(err)
	}
	if _, err := ExecSQLContext(ctx, "INSERT INTO pgext.repo_data (id,data) VALUES ($1,$2)", repoID, []byte("legacy RPM cache")); err != nil {
		t.Fatal(err)
	}
	if err := NewParserContext(ctx, ParseOptions{}).ParseAndRecap(); err == nil {
		t.Fatal("published a catalog from an unrefreshed RPM cache")
	}
	for _, table := range []string{"apt", "dnf", "bin"} {
		var count int
		if err := QueryRowContext(ctx, "SELECT count(*) FROM pgext."+table).Scan(&count); err != nil || count != 1 {
			t.Fatalf("failed RPM parse changed live %s: %d, %v", table, count, err)
		}
	}
}
