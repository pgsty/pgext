package db

import (
	"encoding/csv"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestCSVTableRegistry(t *testing.T) {
	want := []string{"pg", "os", "category", "repository", "extension", "universe"}
	if got := CSVTableNames(); !reflect.DeepEqual(got, want) {
		t.Fatalf("CSVTableNames() = %v, want %v", got, want)
	}

	if !IsCSVTable("ext") || !IsCSVTable(" repo ") || !IsCSVTable("CAT") {
		t.Fatal("expected documented aliases to resolve")
	}
	if IsCSVTable("all_extensions") || IsCSVTable("apt") {
		t.Fatal("unregistered tables must not be accepted as embedded CSV tables")
	}

	extension, err := GetCSVFile("ext")
	if err != nil {
		t.Fatalf("GetCSVFile(ext): %v", err)
	}
	if extension.Name != "extension" || len(extension.Content) == 0 {
		t.Fatalf("GetCSVFile(ext) = name %q, %d bytes", extension.Name, len(extension.Content))
	}
}

func TestCSVTableRegistryCoversEmbeddedAssets(t *testing.T) {
	entries, err := embeddedFS.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	assets := make(map[string]bool)
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".csv") {
			assets[entry.Name()] = true
		}
	}
	for _, table := range csvTables {
		if !assets[table.Asset] {
			t.Errorf("registered CSV asset %q is not embedded", table.Asset)
		}
		delete(assets, table.Asset)
	}
	for asset := range assets {
		t.Errorf("embedded CSV asset %q is missing from the table registry", asset)
	}
}

func TestUniverseSchemaMatchesCSV(t *testing.T) {
	schema, err := GetSchema()
	if err != nil {
		t.Fatal(err)
	}
	universeDDL, err := GetUniverseSchema()
	if err != nil {
		t.Fatal(err)
	}
	universe, err := GetCSVFile("universe")
	if err != nil {
		t.Fatal(err)
	}

	header := csvHeader(t, universe.Content)
	columns := schemaTableColumns(t, schema, "universe")
	if !reflect.DeepEqual(columns, header) {
		t.Fatalf("pgext.universe schema columns do not match universe.csv\nschema: %v\nCSV:    %v", columns, header)
	}
	migrationColumns := schemaTableColumns(t, universeDDL, "universe")
	if !reflect.DeepEqual(migrationColumns, header) {
		t.Fatalf("pgext.universe migration columns do not match universe.csv\nDDL: %v\nCSV: %v", migrationColumns, header)
	}
	if strings.Contains(universeDDL, "CREATE TABLE IF NOT EXISTS pgext.doc") {
		t.Fatal("universe migration DDL must not include the following doc table")
	}
	legacyCatalog := "extension" + "_" + "all"
	if strings.Contains(schema, legacyCatalog) || strings.Contains(string(universe.Content), legacyCatalog) {
		t.Fatalf("legacy catalog name %q must not appear in the schema or universe.csv", legacyCatalog)
	}

	extensionPos := strings.Index(schema, "CREATE TABLE IF NOT EXISTS pgext.extension")
	universePos := strings.Index(schema, "CREATE TABLE pgext.universe")
	if extensionPos < 0 || universePos < 0 || extensionPos >= universePos {
		t.Fatalf("extension DDL must precede universe DDL: extension=%d universe=%d", extensionPos, universePos)
	}
	for _, fragment := range []string{
		"CREATE INDEX universe_pkg_idx ON pgext.universe (pkg)",
		"COMMENT ON TABLE pgext.universe",
		"COMMENT ON COLUMN pgext.universe.kind",
		"COMMENT ON COLUMN pgext.universe.required_by",
		"CREATE OR REPLACE VIEW pgext.ext",
		"COMMENT ON COLUMN pgext.universe.mtime",
	} {
		if !strings.Contains(schema, fragment) {
			t.Errorf("schema is missing %q", fragment)
		}
	}
}

func TestMatrixSchemaIsStandaloneCompactMaterialization(t *testing.T) {
	matrixDDL, err := GetMatrixSchema()
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		"CREATE MATERIALIZED VIEW pgext.matrix",
		"string_agg(c.code, '' ORDER BY c.oi, c.pi)",
		"'n', to_jsonb(names), 'v', to_jsonb(versions), 'd', details",
		"CREATE UNIQUE INDEX matrix_pkg_idx",
	} {
		if !strings.Contains(matrixDDL, fragment) {
			t.Errorf("matrix migration DDL is missing %q", fragment)
		}
	}
	if strings.Contains(matrixDDL, "CREATE DOMAIN pgext.version") {
		t.Fatal("matrix migration DDL must not include the following version domain")
	}
}

func TestUniverseContainsPackagedExtensionCatalog(t *testing.T) {
	extension, err := GetCSVFile("extension")
	if err != nil {
		t.Fatal(err)
	}
	universe, err := GetCSVFile("universe")
	if err != nil {
		t.Fatal(err)
	}

	packaged := csvColumn(t, extension.Content, "name")
	all := csvColumn(t, universe.Content, "name")
	for name := range packaged {
		if !all[name] {
			t.Errorf("packaged extension %q is missing from universe.csv", name)
		}
	}
	if len(all) <= len(packaged) {
		t.Fatalf("universe must be a strict superset: extension=%d universe=%d", len(packaged), len(all))
	}
}

func csvHeader(t *testing.T, data []byte) []string {
	t.Helper()
	record, err := csv.NewReader(strings.NewReader(string(data))).Read()
	if err != nil {
		t.Fatalf("read CSV header: %v", err)
	}
	return record
}

func csvColumn(t *testing.T, data []byte, column string) map[string]bool {
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

func schemaTableColumns(t *testing.T, schema, table string) []string {
	t.Helper()
	marker := "CREATE TABLE IF NOT EXISTS pgext." + table
	start := strings.Index(schema, marker)
	if start < 0 {
		marker = "CREATE TABLE pgext." + table
		start = strings.Index(schema, marker)
	}
	if start < 0 {
		t.Fatalf("schema table %s not found", table)
	}
	start = strings.Index(schema[start:], "(") + start + 1
	end := strings.Index(schema[start:], "\n);")
	if end < 0 {
		t.Fatalf("schema table %s terminator not found", table)
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
