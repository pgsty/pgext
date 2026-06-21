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
	primaryBzip2 := []byte{
		0x42, 0x5a, 0x68, 0x39, 0x31, 0x41, 0x59, 0x26, 0x53, 0x59, 0x6a, 0xa1,
		0xe6, 0xb3, 0x00, 0x00, 0x03, 0x91, 0x80, 0x00, 0x02, 0x24, 0x22, 0x54,
		0x20, 0x20, 0x00, 0x22, 0x01, 0xa6, 0xd4, 0x20, 0xc9, 0x88, 0x0e, 0x5a,
		0x36, 0x54, 0x67, 0x8b, 0xb9, 0x22, 0x9c, 0x28, 0x48, 0x35, 0x50, 0xf3,
		0x59, 0x80,
	}

	repoDir := t.TempDir()
	repodataDir := filepath.Join(repoDir, "repodata")
	if err := os.MkdirAll(repodataDir, 0755); err != nil {
		t.Fatalf("create repodata dir: %v", err)
	}
	repomdPath := filepath.Join(repodataDir, "repomd.xml")
	repomd := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo">
  <data type="primary_db">
    <location href="repodata/primary.sqlite.bz2"/>
  </data>
</repomd>`)
	if err := os.WriteFile(repomdPath, repomd, 0644); err != nil {
		t.Fatalf("write repomd.xml: %v", err)
	}
	if err := os.WriteFile(filepath.Join(repodataDir, "primary.sqlite.bz2"), primaryBzip2, 0644); err != nil {
		t.Fatalf("write primary.sqlite.bz2: %v", err)
	}

	scanner := &Scanner{}
	repo := &RepoMetadata{
		ID:         "el8.aarch64.pigsty",
		Type:       "rpm",
		CachedSize: sql.NullInt64{Int64: int64(len("primary-data")), Valid: true},
	}

	result := scanner.scanRPM(context.Background(), repo, repomdPath)
	if result.Error != nil {
		t.Fatalf("scanRPM() error = %v", result.Error)
	}
	if !result.Updated {
		t.Fatal("scanRPM() must update local metadata even when cached size matches")
	}
	if string(result.Data) != "primary-data" {
		t.Fatalf("scanRPM() data = %q, want %q", result.Data, "primary-data")
	}
}
