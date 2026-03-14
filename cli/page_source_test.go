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
