package cli

import (
	"context"
	"strings"
	"testing"
)

func TestDNFParserExtractPrimaryXMLSparseFields(t *testing.T) {
	parser := NewDNFParser(context.Background())
	data := []byte(`<metadata packages="1"><package type="rpm"><name>demo</name><arch>noarch</arch><version epoch="0" ver="1.0" rel="1"/><checksum type="sha256">abc</checksum><location href="demo.rpm"/></package></metadata>`)
	packages, err := parser.extractPrimaryXML(data)
	if err != nil || len(packages) != 1 {
		t.Fatalf("parse sparse package: %+v, %v", packages, err)
	}
	p := packages[0]
	if p.PkgKey != 1 || p.PkgId != "abc" || p.Name != "demo" || p.Summary.Valid || p.TimeBuild.Valid || p.SizePackage.Valid {
		t.Fatalf("incorrect sparse XML fields: %+v", p)
	}
}

func TestDNFParserRejectsLegacyCache(t *testing.T) {
	parser := NewDNFParser(context.Background())
	for _, data := range [][]byte{[]byte("SQLite format 3\x00"), []byte("not metadata")} {
		if _, err := parser.ParseRepository("legacy", data); err == nil || !strings.Contains(err.Error(), "pgext fetch") {
			t.Fatalf("legacy cache error = %v", err)
		}
	}
}

func TestDNFParserParseRepositoryEmpty(t *testing.T) {
	parser := NewDNFParser(context.Background())
	count, err := parser.ParseRepository("empty", []byte(`<metadata packages="0"/>`))
	if err != nil || count != 0 {
		t.Fatalf("empty primary XML: %d, %v", count, err)
	}
}
