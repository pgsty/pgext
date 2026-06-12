/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"net/http"
	"os"
	"os/exec"
	"testing"
)

const testEnvProxyURL = "http://127.0.0.1:3128"

func TestNewFetcherUsesEnvironmentProxy(t *testing.T) {
	runProxyHelper(t, "fetcher")
}

func TestPGXNExporterUsesEnvironmentProxy(t *testing.T) {
	runProxyHelper(t, "pgxn")
}

func TestProxyHelper(t *testing.T) {
	switch os.Getenv("PGEXT_PROXY_HELPER") {
	case "":
		t.Skip("helper process only")
	case "fetcher":
		assertFetcherProxy(t)
	case "pgxn":
		assertPGXNProxy(t)
	default:
		t.Fatalf("unknown proxy helper %q", os.Getenv("PGEXT_PROXY_HELPER"))
	}
}

func runProxyHelper(t *testing.T, helper string) {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestProxyHelper", "-test.v")
	cmd.Env = append(os.Environ(),
		"PGEXT_PROXY_HELPER="+helper,
		"HTTP_PROXY="+testEnvProxyURL,
		"NO_PROXY=",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("proxy helper failed: %v\n%s", err, output)
	}
}

func assertFetcherProxy(t *testing.T) {
	fetcher := NewFetcher(FetchOptions{})
	transport, ok := fetcher.httpClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("expected *http.Transport, got %T", fetcher.httpClient.Transport)
	}
	if transport.Proxy == nil {
		t.Fatal("expected fetcher transport to honor proxy environment variables")
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com/repo", nil)
	if err != nil {
		t.Fatal(err)
	}
	proxyURL, err := transport.Proxy(req)
	if err != nil {
		t.Fatal(err)
	}
	if proxyURL == nil || proxyURL.String() != testEnvProxyURL {
		t.Fatalf("expected HTTP_PROXY URL, got %v", proxyURL)
	}
}

func assertPGXNProxy(t *testing.T) {
	exporter := newPGXNExporter(PgxnExportOptions{})
	transport, ok := exporter.client.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("expected *http.Transport, got %T", exporter.client.Transport)
	}
	if transport.Proxy == nil {
		t.Fatal("expected PGXN exporter transport to honor proxy environment variables")
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com/src", nil)
	if err != nil {
		t.Fatal(err)
	}
	proxyURL, err := transport.Proxy(req)
	if err != nil {
		t.Fatal(err)
	}
	if proxyURL == nil || proxyURL.String() != testEnvProxyURL {
		t.Fatalf("expected HTTP_PROXY URL, got %v", proxyURL)
	}
}
