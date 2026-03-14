package cli

import (
	"database/sql"
	"strings"
	"testing"
)

func TestExtensionGetSourceURL(t *testing.T) {
	ext := &Extension{
		Source: sql.NullString{Valid: true, String: "aggs_for_arrays-1.3.3.tar.gz"},
	}

	if got := ext.GetSourceURL(RegionDefault); got != "https://repo.pigsty.io/ext/src/aggs_for_arrays-1.3.3.tar.gz" {
		t.Fatalf("default source URL = %q", got)
	}
	if got := ext.GetSourceURL(RegionChina); got != "https://repo.pigsty.cc/ext/src/aggs_for_arrays-1.3.3.tar.gz" {
		t.Fatalf("china source URL = %q", got)
	}
}

func TestExtensionGetSourceURLKeepsAbsoluteURL(t *testing.T) {
	ext := &Extension{
		Source: sql.NullString{Valid: true, String: "https://example.com/ext/src/demo.tar.gz"},
	}

	if got := ext.GetSourceURL(RegionChina); got != "https://example.com/ext/src/demo.tar.gz" {
		t.Fatalf("absolute source URL = %q", got)
	}
	if got := ext.GetSourceFilename(); got != "demo.tar.gz" {
		t.Fatalf("source filename = %q", got)
	}
}

func TestBinaryGetDownloadURL(t *testing.T) {
	bin := &Binary{
		Href:       "yum/el9.x86_64/demo.rpm",
		DefaultURL: "https://repo.pigsty.io/ext",
		MirrorURL:  sql.NullString{Valid: true, String: "https://repo.pigsty.cc/ext"},
	}

	if got := bin.GetDownloadURL(RegionDefault); got != "https://repo.pigsty.io/ext/yum/el9.x86_64/demo.rpm" {
		t.Fatalf("default binary URL = %q", got)
	}
	if got := bin.GetDownloadURL(RegionChina); got != "https://repo.pigsty.cc/ext/yum/el9.x86_64/demo.rpm" {
		t.Fatalf("china binary URL = %q", got)
	}
}

func TestBinaryGetDownloadURLFallsBackWithoutMirror(t *testing.T) {
	bin := &Binary{
		Href:       "apt/u24/demo.deb",
		DefaultURL: "https://repo.pigsty.io/ext",
	}

	if got := bin.GetDownloadURL(RegionChina); got != "https://repo.pigsty.io/ext/apt/u24/demo.deb" {
		t.Fatalf("china fallback binary URL = %q", got)
	}
}

func TestPageGeneratorsRenderSiteSpecificSourceURL(t *testing.T) {
	ext := &Extension{
		Source: sql.NullString{Valid: true, String: "aggs_for_arrays-1.3.3.tar.gz"},
	}

	ioCards := NewIOPageGenerator(nil, "", "").generateCards(ext)
	if !strings.Contains(ioCards, `href="https://repo.pigsty.io/ext/src/aggs_for_arrays-1.3.3.tar.gz"`) {
		t.Fatalf("io cards missing full source URL: %s", ioCards)
	}

	ccCards := NewCCPageGenerator(nil, "", "").generateCards(ext)
	if !strings.Contains(ccCards, `href="https://repo.pigsty.cc/ext/src/aggs_for_arrays-1.3.3.tar.gz"`) {
		t.Fatalf("cc cards missing full source URL: %s", ccCards)
	}
}

func TestPageGeneratorsRenderSiteSpecificBinaryURL(t *testing.T) {
	cache := &ExtensionCache{
		PGVersions: []int{17},
		OSVersions: []OSVersion{{OS: "el9.x86_64", Active: true}},
	}
	ext := &Extension{
		Name: "demo",
		Pkg:  "demo",
	}
	packages := []*PkgInfo{
		{
			PG:      17,
			OS:      "el9.x86_64",
			Pkg:     "demo",
			Ext:     "demo",
			State:   sql.NullString{Valid: true, String: "AVAIL"},
			Org:     sql.NullString{Valid: true, String: "PIGSTY"},
			Version: sql.NullString{Valid: true, String: "1.0.0"},
			Count:   sql.NullInt64{Valid: true, Int64: 1},
		},
	}
	binaries := []*Binary{
		{
			PG:         17,
			OS:         "el9.x86_64",
			Name:       "demo_17",
			Org:        sql.NullString{Valid: true, String: "PIGSTY"},
			Version:    "1.0.0",
			File:       "demo_17-1.0.0.rpm",
			Href:       "yum/el9.x86_64/demo_17-1.0.0.rpm",
			DefaultURL: "https://repo.pigsty.io/ext",
			MirrorURL:  sql.NullString{Valid: true, String: "https://repo.pigsty.cc/ext"},
			Size:       1024,
		},
	}

	ioAvailability := NewIOPageGenerator(cache, "", "").generateAvailability(ext, packages, binaries)
	if !strings.Contains(ioAvailability, "https://repo.pigsty.io/ext/yum/el9.x86_64/demo_17-1.0.0.rpm") {
		t.Fatalf("io availability missing default binary URL: %s", ioAvailability)
	}

	ccAvailability := NewCCPageGenerator(cache, "", "").generateAvailability(ext, packages, binaries)
	if !strings.Contains(ccAvailability, "https://repo.pigsty.cc/ext/yum/el9.x86_64/demo_17-1.0.0.rpm") {
		t.Fatalf("cc availability missing mirror binary URL: %s", ccAvailability)
	}
}
