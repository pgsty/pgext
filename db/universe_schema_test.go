package db

import (
	"encoding/csv"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestUniverseCatalogContract(t *testing.T) {
	schema, err := GetSchema()
	if err != nil {
		t.Fatal(err)
	}
	universeCSV, err := ReadFile("universe.csv")
	if err != nil {
		t.Fatal(err)
	}
	extensionCSV, err := ReadFile("extension.csv")
	if err != nil {
		t.Fatal(err)
	}

	header := universeCSVHeader(t, universeCSV)
	columns := universeSchemaColumns(t, schema)
	if !reflect.DeepEqual(columns, header) {
		t.Fatalf("pgext.universe schema columns do not match universe.csv\nschema: %v\nCSV:    %v", columns, header)
	}

	legacyCatalog := "extension" + "_" + "all"
	if strings.Contains(schema, legacyCatalog) || strings.Contains(string(universeCSV), legacyCatalog) {
		t.Fatalf("legacy catalog name %q must not appear in schema or universe.csv", legacyCatalog)
	}

	packaged := universeCSVColumn(t, extensionCSV, "name")
	universe := universeCSVColumn(t, universeCSV, "name")
	for name := range packaged {
		if !universe[name] {
			t.Errorf("packaged extension %q is missing from universe.csv", name)
		}
	}
	if len(universe) <= len(packaged) {
		t.Fatalf("universe must be a strict superset: extension=%d universe=%d", len(packaged), len(universe))
	}

	files, err := GetCSVFiles()
	if err != nil {
		t.Fatal(err)
	}
	extensionPos, universePos := -1, -1
	for i, file := range files {
		switch file.Name {
		case "extension":
			extensionPos = i
		case "universe":
			universePos = i
		}
	}
	if extensionPos < 0 || universePos < 0 || extensionPos >= universePos {
		t.Fatalf("embedded load order must place extension before universe: extension=%d universe=%d", extensionPos, universePos)
	}
}

func universeCSVHeader(t *testing.T, data []byte) []string {
	t.Helper()
	record, err := csv.NewReader(strings.NewReader(string(data))).Read()
	if err != nil {
		t.Fatalf("read CSV header: %v", err)
	}
	return record
}

func universeCSVColumn(t *testing.T, data []byte, column string) map[string]bool {
	t.Helper()
	reader := csv.NewReader(strings.NewReader(string(data)))
	header, err := reader.Read()
	if err != nil {
		t.Fatalf("read CSV header: %v", err)
	}
	columnIndex := -1
	for i, name := range header {
		if name == column {
			columnIndex = i
			break
		}
	}
	if columnIndex < 0 {
		t.Fatalf("CSV column %q not found", column)
	}

	values := make(map[string]bool)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("read CSV record: %v", err)
		}
		values[record[columnIndex]] = true
	}
	return values
}

func universeSchemaColumns(t *testing.T, schema string) []string {
	t.Helper()
	marker := "CREATE TABLE IF NOT EXISTS pgext.universe"
	start := strings.Index(schema, marker)
	if start < 0 {
		t.Fatal("pgext.universe table is missing from schema.sql")
	}
	start = strings.Index(schema[start:], "(") + start + 1
	end := strings.Index(schema[start:], "\n);")
	if end < 0 {
		t.Fatal("pgext.universe table terminator is missing from schema.sql")
	}

	var columns []string
	for _, line := range strings.Split(schema[start:start+end], "\n") {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) == 0 || strings.HasPrefix(fields[0], "--") {
			continue
		}
		columns = append(columns, strings.TrimSuffix(fields[0], ","))
	}
	return columns
}
