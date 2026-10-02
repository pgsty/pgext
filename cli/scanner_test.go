package cli

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

func TestScanDEBAlwaysUpdatesEvenWhenCachedSizeMatches(t *testing.T) {
	packagesPath := filepath.Join(t.TempDir(), "Packages")
	data := []byte("Package: demo\nVersion: 1.0.0\n")
	if err := os.WriteFile(packagesPath, data, 0644); err != nil {
		t.Fatalf("write Packages file: %v", err)
	}

	scanner := &Scanner{}
	repo := &RepoMetadata{
		ID:         "u24.x86_64.pigsty",
		Type:       "deb",
		CachedSize: sql.NullInt64{Int64: int64(len(data)), Valid: true},
	}

	result := scanner.scanDEB(context.Background(), repo, packagesPath)
	if result.Error != nil {
		t.Fatalf("scanDEB() error = %v", result.Error)
	}
	if !result.Updated {
		t.Fatal("scanDEB() must update local metadata even when cached size matches")
	}
	if string(result.Data) != string(data) {
		t.Fatalf("scanDEB() data = %q, want %q", result.Data, data)
	}
}

func TestScanRPMAlwaysUpdatesEvenWhenCachedSizeMatches(t *testing.T) {
	primaryGzip := compressPrimaryFixture(t, "gzip")

	repoDir := t.TempDir()
	repodataDir := filepath.Join(repoDir, "repodata")
	if err := os.MkdirAll(repodataDir, 0755); err != nil {
		t.Fatalf("create repodata dir: %v", err)
	}
	repomdPath := filepath.Join(repodataDir, "repomd.xml")
	repomd := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo">
  <data type="primary">
    <location href="repodata/primary.xml.gz"/>
  </data>
</repomd>`)
	if err := os.WriteFile(repomdPath, repomd, 0644); err != nil {
		t.Fatalf("write repomd.xml: %v", err)
	}
	if err := os.WriteFile(filepath.Join(repodataDir, "primary.xml.gz"), primaryGzip, 0644); err != nil {
		t.Fatalf("write primary.xml.gz: %v", err)
	}

	scanner := &Scanner{}
	repo := &RepoMetadata{
		ID:         "el8.aarch64.pigsty",
		Type:       "rpm",
		CachedSize: sql.NullInt64{Int64: int64(len(primaryXMLFixture)), Valid: true},
	}

	result := scanner.scanRPM(context.Background(), repo, repomdPath)
	if result.Error != nil {
		t.Fatalf("scanRPM() error = %v", result.Error)
	}
	if !result.Updated {
		t.Fatal("scanRPM() must update local metadata even when cached size matches")
	}
	if string(result.Data) != primaryXMLFixture {
		t.Fatalf("scanRPM() data = %q, want %q", result.Data, primaryXMLFixture)
	}
}
