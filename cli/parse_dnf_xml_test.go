package cli

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const primaryXMLFixture = `<?xml version="1.0"?>
<metadata xmlns="http://linux.duke.edu/metadata/common" xmlns:rpm="http://linux.duke.edu/metadata/rpm" packages="1">
<package type="rpm"><name>orafce_18</name><arch>x86_64</arch>
<version epoch="0" ver="4.16.12" rel="1.el9"/><checksum type="sha256">abc123</checksum>
<summary>Oracle &amp; Postgres</summary><description>Functions</description><packager>Pigsty</packager><url>https://example.org</url>
<time file="0" build="1790865631"/><size package="1234" installed="4321" archive="4500"/>
<location href="orafce_18.rpm" xml:base="https://example.org/"/>
<format><rpm:license>BSD</rpm:license><rpm:vendor>Pigsty</rpm:vendor><rpm:group>Database</rpm:group>
<rpm:buildhost>builder</rpm:buildhost><rpm:sourcerpm>orafce.src.rpm</rpm:sourcerpm>
<rpm:header-range start="10" end="20"/></format></package></metadata>`

func TestDNFPrimaryXML(t *testing.T) {
	p := NewDNFParser(context.Background())
	pkgs, err := p.extractPrimaryXML([]byte(primaryXMLFixture))
	if err != nil {
		t.Fatal(err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d", len(pkgs))
	}
	x := pkgs[0]
	if x.Name != "orafce_18" || x.Version.String != "4.16.12" || x.Release.String != "1.el9" || x.Arch.String != "x86_64" || x.Epoch.String != "0" || x.PkgId != "abc123" || x.ChecksumType.String != "sha256" {
		t.Fatalf("incorrect package identity: %+v", x)
	}
	if x.Summary.String != "Oracle & Postgres" || !x.TimeFile.Valid || x.TimeFile.Int64 != 0 || x.TimeBuild.Int64 != 1790865631 || x.SizePackage.Int64 != 1234 || x.SizeInstalled.Int64 != 4321 || x.RPMLicense.String != "BSD" || x.RPMHeaderEnd.Int64 != 20 || x.LocationBase.String != "https://example.org/" || x.LocationHref.String != "orafce_18.rpm" || x.RPMSourceRPM.String != "orafce.src.rpm" || x.RPMPackager.String != "Pigsty" {
		t.Fatalf("incorrect package details: %+v", x)
	}
	for _, input := range []string{"", "not sqlite or XML", `<html/>`, `<metadata/>`, `<metadata packages="1"/>`, strings.Replace(primaryXMLFixture, "</metadata>", "", 1), strings.Replace(primaryXMLFixture, "<name>orafce_18</name>", "", 1)} {
		if _, err := p.ParseRepository("invalid", []byte(input)); err == nil {
			t.Errorf("accepted invalid metadata: %q", input)
		}
	}
	if n, err := p.ParseRepository("empty", []byte(`<metadata packages="0"/>`)); err != nil || n != 0 {
		t.Fatalf("empty XML: %d, %v", n, err)
	}
}

func TestFetchRPMPrimaryXML(t *testing.T) {
	for _, badChecksum := range []bool{false, true} {
		t.Run(fmt.Sprintf("badChecksum=%t", badChecksum), func(t *testing.T) {
			var compressed bytes.Buffer
			gz := gzip.NewWriter(&compressed)
			_, _ = gz.Write([]byte(primaryXMLFixture))
			if err := gz.Close(); err != nil {
				t.Fatal(err)
			}
			checksum := fmt.Sprintf("%x", sha256.Sum256([]byte(primaryXMLFixture)))
			if badChecksum {
				checksum = strings.Repeat("0", 64)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/repodata/repomd.xml":
					fmt.Fprintf(w, `<repomd><data type="primary"><location href="repodata/primary.xml.gz"/><open-checksum type="sha256">%s</open-checksum></data></repomd>`, checksum)
				case "/repodata/primary.xml.gz":
					_, _ = w.Write(compressed.Bytes())
				default:
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			result := NewFetcher(FetchOptions{}).fetchRPM(context.Background(), &RepoMetadata{ID: "test", Type: "rpm", MetadataURL: server.URL + "/repodata/repomd.xml"})
			if badChecksum {
				if result.Error == nil {
					t.Fatal("accepted corrupt checksum")
				}
				return
			}
			if result.Error != nil {
				t.Fatal(result.Error)
			}
			pkgs, err := NewDNFParser(context.Background()).extractPrimaryXML(result.Data)
			if err != nil || len(pkgs) != 1 || pkgs[0].Name != "orafce_18" {
				t.Fatalf("fetch/parse failed: %+v, %v", pkgs, err)
			}
		})
	}
}
