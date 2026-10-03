package db

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

// Run the real embedded SQL in an isolated schema rolled back after testing.
// This never replaces functions or package data in the application schema.
func TestVersionCompareIntegration(t *testing.T) {
	url := os.Getenv("PGEXT_VERSION_TEST_PGURL")
	if url == "" {
		t.Skip("set PGEXT_VERSION_TEST_PGURL to exercise the SQL comparator")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(context.Background())
	namespace := fmt.Sprintf("pgext_version_test_%d", time.Now().UnixNano())
	if _, err := tx.Exec(ctx, "CREATE SCHEMA "+namespace); err != nil {
		t.Fatal(err)
	}
	schema, err := GetSchema()
	if err != nil {
		t.Fatal(err)
	}
	pattern := regexp.MustCompile(`(?s)CREATE OR REPLACE FUNCTION pgext\.version_(?:segment_)?compare\(v1 TEXT, v2 TEXT\).*?\$\$ LANGUAGE plpgsql IMMUTABLE STRICT PARALLEL SAFE;`)
	functions := pattern.FindAllString(schema, -1)
	if len(functions) != 2 {
		t.Fatalf("expected two comparator functions, got %d", len(functions))
	}
	for _, function := range functions {
		if _, err := tx.Exec(ctx, strings.ReplaceAll(function, "pgext.", namespace+".")); err != nil {
			t.Fatal(err)
		}
	}
	cases := []struct {
		left, right string
		want        int
	}{
		{"1.2.1-1PGSTY.el10", "1.2-2PGDG.rhel10.2", 1},
		{"1.2.1-1PGSTY.el8", "1.2-2PGDG.rhel8.10", 1},
		{"4.5.0-1PGSTY~noble", "4.2.3-9.pgdg24.04+1", 1},
		{"1.2-2", "1.2-1", 1},
		{"1.2.1", "1.2-9", 1},
		{"1.2.1-1", "1.2", 1},
		{"1:1.0-1", "2.0-9", 1},
		{"0:2.0-1", "2.0-1", 0},
		{"1.2~rc1-1", "1.2-1", -1},
		{"1.2-1~beta", "1.2-1", -1},
		{"1.0-a-2", "1.0-a-1", 1},
		{"1.2.1-1", "1.2.1-1", 0},
	}
	for _, tc := range cases {
		t.Run(tc.left+"_vs_"+tc.right, func(t *testing.T) {
			for _, pair := range []struct {
				left, right string
				want        int
			}{{tc.left, tc.right, tc.want}, {tc.right, tc.left, -tc.want}} {
				var got int
				if err := tx.QueryRow(ctx, "SELECT "+namespace+".version_compare($1,$2)", pair.left, pair.right).Scan(&got); err != nil {
					t.Fatal(err)
				}
				if got != pair.want {
					t.Errorf("compare(%q,%q)=%d, want %d", pair.left, pair.right, got, pair.want)
				}
			}
		})
	}
}
