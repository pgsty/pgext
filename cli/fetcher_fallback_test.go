/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRPMFallsBackToMirrorMetadata(t *testing.T) {
	primaryGzip := compressPrimaryFixture(t, "gzip")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/default/repodata/repomd.xml":
			http.NotFound(w, r)
		case "/mirror/repodata/repomd.xml":
			w.Header().Set("Content-Type", "application/xml")
			_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo">
  <data type="primary">
    <location href="repodata/primary.xml.gz"/>
  </data>
</repomd>`))
		case "/mirror/repodata/primary.xml.gz":
			_, _ = w.Write(primaryGzip)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	fetcher := NewFetcher(FetchOptions{})
	repo := &RepoMetadata{
		ID:          "el8.x86_64.pgnf15",
		Type:        "rpm",
		MetadataURL: server.URL + "/default/repodata/repomd.xml",
		MirrorURL:   server.URL + "/mirror/repodata/repomd.xml",
	}

	result := fetcher.fetchRPM(context.Background(), repo)
	if result.Error != nil {
		t.Fatalf("expected mirror fallback to succeed, got %v", result.Error)
	}
	if string(result.Data) != primaryXMLFixture {
		t.Fatalf("expected primary XML payload from mirror, got %q", result.Data)
	}
}
