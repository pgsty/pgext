package cli

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
)

func compressPrimaryFixture(t *testing.T, codec string) []byte {
	t.Helper()
	var b bytes.Buffer
	switch codec {
	case "none":
		return []byte(primaryXMLFixture)
	case "gzip":
		w := gzip.NewWriter(&b)
		_, _ = w.Write([]byte(primaryXMLFixture))
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
	case "bzip2":
		data, err := os.ReadFile("testdata/primary.xml.bz2")
		if err != nil {
			t.Fatal(err)
		}
		return data
	case "xz":
		w, err := xz.NewWriter(&b)
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write([]byte(primaryXMLFixture))
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
	case "zstd":
		w, err := zstd.NewWriter(&b, zstd.WithEncoderConcurrency(1))
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write([]byte(primaryXMLFixture))
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
	default:
		t.Fatalf("unknown test codec %q", codec)
	}
	return b.Bytes()
}

func primaryManifest(data []byte, href string) string {
	return fmt.Sprintf(`<repomd xmlns="http://linux.duke.edu/metadata/repo"><data type="primary"><location href="%s"/><checksum type="sha256">%x</checksum><open-checksum type="sha256">%x</open-checksum><size>%d</size><open-size>%d</open-size></data></repomd>`, href, sha256.Sum256(data), sha256.Sum256([]byte(primaryXMLFixture)), len(data), len(primaryXMLFixture))
}

func TestRPMPrimaryFetchAndScanCompression(t *testing.T) {
	for _, tc := range []struct{ codec, suffix string }{
		{"none", ".xml"}, {"gzip", ".xml.gz"}, {"bzip2", ".xml.bz2"},
		{"xz", ".xml.xz"}, {"zstd", ".xml.zst"},
	} {
		t.Run(tc.codec, func(t *testing.T) {
			data := compressPrimaryFixture(t, tc.codec)
			href := "repodata/primary" + tc.suffix
			manifest := primaryManifest(data, href)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/repo/repodata/repomd.xml":
					_, _ = w.Write([]byte(manifest))
				case "/repo/" + href:
					_, _ = w.Write(data)
				default:
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			repo := &RepoMetadata{ID: "rpm-test", Type: "rpm", MetadataURL: server.URL + "/repo/repodata/repomd.xml"}
			fetched := NewFetcher(FetchOptions{}).fetchOne(context.Background(), repo)
			if fetched.Error != nil || string(fetched.Data) != primaryXMLFixture {
				t.Fatalf("fetch primary XML: %v", fetched.Error)
			}
			if fetched.Extra.Format != rpmPrimaryFormat || fetched.Extra.Compression != tc.codec || fetched.Extra.PackageCount != 1 {
				t.Fatalf("wrong fetch descriptor: %+v", fetched.Extra)
			}
			root := t.TempDir()
			if err := os.Mkdir(filepath.Join(root, "repodata"), 0755); err != nil {
				t.Fatal(err)
			}
			manifestPath := filepath.Join(root, "repodata", "repomd.xml")
			if err := os.WriteFile(manifestPath, []byte(manifest), 0644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(root, href), data, 0644); err != nil {
				t.Fatal(err)
			}
			scanned := (&Scanner{}).scanRPM(context.Background(), repo, manifestPath)
			if scanned.Error != nil || !bytes.Equal(scanned.Data, fetched.Data) || scanned.Extra.Compression != tc.codec {
				t.Fatalf("scan/fetch mismatch: %v", scanned.Error)
			}
		})
	}
}

func TestRPMPrimaryRejectsDatabaseOnlyAndAmbiguousManifests(t *testing.T) {
	for _, tc := range []struct {
		name     string
		entries  string
		wantFail bool
	}{
		{"database only", `<data type="primary_db"><location href="legacy.db"/></data>`, true},
		{"empty location", `<data type="primary"/>`, true},
		{"duplicate primary", `<data type="primary"><location href="a.xml"/></data><data type="primary"><location href="b.xml"/></data>`, true},
		{"both formats", `<data type="primary_db"><location href="legacy.db"/></data><data type="primary"><location href="primary.xml"/></data>`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "repomd.xml") {
					_, _ = fmt.Fprintf(w, "<repomd>%s</repomd>", tc.entries)
				} else if strings.HasSuffix(r.URL.Path, "primary.xml") {
					_, _ = w.Write([]byte(primaryXMLFixture))
				} else {
					t.Errorf("requested obsolete metadata: %s", r.URL.Path)
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			result := NewFetcher(FetchOptions{}).fetchRPM(context.Background(), &RepoMetadata{Type: "rpm", MetadataURL: server.URL + "/repodata/repomd.xml"})
			if (result.Error != nil) != tc.wantFail {
				t.Fatalf("manifest result = %v, want error %t", result.Error, tc.wantFail)
			}
		})
	}
}

func TestRPMPrimaryCacheIsBoundToContentAndURL(t *testing.T) {
	data := compressPrimaryFixture(t, "gzip")
	var downloads, manifests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.Header.Get("If-None-Match") != "" || r.Header.Get("If-Modified-Since") != "" {
			t.Errorf("RPM request reused unrelated validators: %s %v", r.Method, r.Header)
		}
		if strings.HasSuffix(r.URL.Path, "repomd.xml") {
			manifests.Add(1)
			_, _ = w.Write([]byte(primaryManifest(data, "repodata/primary.xml.gz")))
		} else {
			downloads.Add(1)
			_, _ = w.Write(data)
		}
	}))
	defer server.Close()
	repo := &RepoMetadata{Type: "rpm", MetadataURL: server.URL + "/repodata/repomd.xml", CachedETag: sql.NullString{String: `"legacy"`, Valid: true}}
	fetcher := NewFetcher(FetchOptions{})
	first := fetcher.fetchOne(context.Background(), repo)
	if first.Error != nil || !first.Updated {
		t.Fatalf("refresh old cache: %v", first.Error)
	}
	repo.CachedPrimary = first.Extra
	repo.CachedDataSHA256 = sql.NullString{String: first.Extra.ContentSHA256, Valid: true}
	second := fetcher.fetchOne(context.Background(), repo)
	if second.Error != nil || second.Updated || manifests.Load() != 2 || downloads.Load() != 1 {
		t.Fatalf("unchanged cache: updated=%t, manifests=%d, downloads=%d, error=%v", second.Updated, manifests.Load(), downloads.Load(), second.Error)
	}
	repo.CachedDataSHA256.String = "corrupt"
	third := fetcher.fetchOne(context.Background(), repo)
	if third.Error != nil || !third.Updated || downloads.Load() != 2 {
		t.Fatalf("corrupt cache was reused: %+v", third)
	}
	repo.CachedDataSHA256.String = first.Extra.ContentSHA256
	repo.CachedPrimary.SourceURL = "https://other.example/primary.xml.gz"
	fourth := fetcher.fetchOne(context.Background(), repo)
	if fourth.Error != nil || !fourth.Updated || downloads.Load() != 3 {
		t.Fatalf("another source's cache was reused: %+v", fourth)
	}
	repo.CachedPrimary = fourth.Extra
	forced := NewFetcher(FetchOptions{Force: true}).fetchOne(context.Background(), repo)
	if forced.Error != nil || !forced.Updated || downloads.Load() != 4 {
		t.Fatalf("forced refresh skipped primary XML: %+v", forced)
	}
}

func TestPrimaryXMLRequiresCompleteDocument(t *testing.T) {
	for _, data := range []string{
		primaryXMLFixture + `<metadata packages="0"/>`, primaryXMLFixture + "trailing text",
		strings.Replace(primaryXMLFixture, `packages="1"`, `packages="2"`, 1),
		strings.Replace(primaryXMLFixture, `packages="1"`, `packages="0"`, 1),
		strings.Replace(primaryXMLFixture, "metadata/common", "metadata/filelists", 1),
		strings.Replace(primaryXMLFixture, "</metadata>", "", 1),
	} {
		if _, err := readPrimaryXML(context.Background(), []byte(data), nil); err == nil {
			t.Fatalf("accepted an incomplete or wrong XML document: %q", data)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := readPrimaryXML(ctx, []byte(primaryXMLFixture), nil); err != context.Canceled {
		t.Fatalf("canceled parse returned %v", err)
	}
}

func TestPrimaryMetadataBoundsAndLocations(t *testing.T) {
	if _, err := readMetadata(strings.NewReader("abcd"), 3); err == nil {
		t.Fatal("silently truncated an oversized payload")
	}
	if b, err := readMetadata(strings.NewReader("abc"), 3); err != nil || string(b) != "abc" {
		t.Fatalf("exact-size payload: %q, %v", b, err)
	}
	for _, href := range []string{"../outside.xml", "repodata/../../outside.xml", "/absolute.xml", "https://other.example/a.xml", "%2e%2e/outside.xml"} {
		if _, err := rpmPrimaryPath("/repo/repodata/repomd.xml", href); err == nil {
			t.Errorf("accepted unsafe local location %q", href)
		}
	}
	for _, tc := range []struct{ href, want string }{
		{"repodata/primary.xml.gz", "https://example.org/repo/repodata/primary.xml.gz"},
		{"https://cdn.example/primary.xml.zst", "https://cdn.example/primary.xml.zst"},
	} {
		got, err := rpmPrimaryURL("https://example.org/repo/repodata/repomd.xml?old=1", tc.href)
		if err != nil || got != tc.want {
			t.Errorf("resolve %q = %q, %v", tc.href, got, err)
		}
	}
}

func TestRPMPrimaryRejectsCorruptMetadata(t *testing.T) {
	good := compressPrimaryFixture(t, "gzip")
	for _, tc := range []struct {
		name   string
		mutate func(*RepoMDData) []byte
	}{
		{"compressed checksum", func(p *RepoMDData) []byte { p.Checksum = MetadataChecksum{Type: "sha256", Text: "bad"}; return good }},
		{"empty compressed checksum", func(p *RepoMDData) []byte {
			p.Checksum = MetadataChecksum{Type: "sha256", Text: " "}
			return good
		}},
		{"empty open checksum", func(p *RepoMDData) []byte {
			p.OpenChecksum = MetadataChecksum{Type: "sha256"}
			return good
		}},
		{"open checksum", func(p *RepoMDData) []byte {
			p.OpenChecksum = MetadataChecksum{Type: "sha256", Text: "bad"}
			return good
		}},
		{"unknown checksum", func(p *RepoMDData) []byte { p.Checksum = MetadataChecksum{Type: "unknown", Text: "bad"}; return good }},
		{"compressed size", func(p *RepoMDData) []byte { n := int64(len(good) + 1); p.Size = &n; return good }},
		{"open size", func(p *RepoMDData) []byte { n := int64(len(primaryXMLFixture) + 1); p.OpenSize = &n; return good }},
		{"truncated compression", func(p *RepoMDData) []byte { return good[:len(good)-8] }},
		{"truncated XML", func(p *RepoMDData) []byte { return []byte(strings.TrimSuffix(primaryXMLFixture, "</metadata>")) }},
		{"wrong format", func(p *RepoMDData) []byte { return []byte("<html/>") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			primary := &RepoMDData{}
			primary.Location.Href = "repodata/primary.xml.gz"
			data, descriptor, err := decodeRPMPrimary(context.Background(), tc.mutate(primary), primary, "https://example.org/primary.xml.gz")
			if err == nil || data != nil || descriptor != nil {
				t.Fatalf("corrupt metadata accepted: %v", err)
			}
		})
	}
}

func TestMetadataChecksums(t *testing.T) {
	for _, tc := range []struct{ algorithm, checksum string }{
		{"md5", "900150983cd24fb0d6963f7d28e17f72"},
		{"sha", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"sha1", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"sha224", "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"},
		{"sha256", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"sha384", "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"},
		{"sha512", "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
	} {
		t.Run(tc.algorithm, func(t *testing.T) {
			c := MetadataChecksum{Type: tc.algorithm, Text: tc.checksum}
			if err := verifyMetadataChecksum([]byte("abc"), c); err != nil {
				t.Fatal(err)
			}
			if err := verifyMetadataChecksum([]byte("abd"), c); err == nil {
				t.Fatal("accepted checksum mismatch")
			}
		})
	}
}

func TestRPMManifestRequiresCompleteDocument(t *testing.T) {
	for _, data := range []string{"<repomd>", "<repomd/><repomd/>", "<repomd/>garbage", "<metadata/>"} {
		if _, err := parseRPMManifest(strings.NewReader(data)); err == nil {
			t.Fatalf("accepted manifest %q", data)
		}
	}
}
