package db

import (
	"strings"
	"testing"
)

func TestPGNFHideExemptsTimescaleDB(t *testing.T) {
	reload, err := embeddedFS.ReadFile("reload.sql")
	if err != nil {
		t.Fatal(err)
	}
	reloadSQL := string(reload)

	if !strings.Contains(reloadSQL, "position('pgnf' in id) > 0 AS hide") {
		t.Fatal("package recap no longer derives the default hide flag from PGDG non-free repositories")
	}
	if !strings.Contains(reloadSQL, "hide = sub.hide AND p.pkg <> 'timescaledb'") {
		t.Fatal("TimescaleDB must remain visible when the selected RPM comes from PGDG non-free")
	}
}
