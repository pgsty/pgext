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
	primaryBzip2 := []byte{
		0x42, 0x5a, 0x68, 0x39, 0x31, 0x41, 0x59, 0x26, 0x53, 0x59, 0x6a, 0xa1,
		0xe6, 0xb3, 0x00, 0x00, 0x03, 0x91, 0x80, 0x00, 0x02, 0x24, 0x22, 0x54,
		0x20, 0x20, 0x00, 0x22, 0x01, 0xa6, 0xd4, 0x20, 0xc9, 0x88, 0x0e, 0x5a,
		0x36, 0x54, 0x67, 0x8b, 0xb9, 0x22, 0x9c, 0x28, 0x48, 0x35, 0x50, 0xf3,
		0x59, 0x80,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/default/repodata/repomd.xml":
			http.NotFound(w, r)
		case "/mirror/repodata/repomd.xml":
			w.Header().Set("Content-Type", "application/xml")
			_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo">
  <data type="primary_db">
    <location href="repodata/primary.sqlite.bz2"/>
  </data>
</repomd>`))
		case "/mirror/repodata/primary.sqlite.bz2":
			_, _ = w.Write(primaryBzip2)
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
	if string(result.Data) != "primary-data" {
		t.Fatalf("expected primary DB payload from mirror, got %q", result.Data)
	}
}
