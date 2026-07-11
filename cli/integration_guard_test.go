package cli

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

const destructiveIntegrationOptIn = "PGEXT_TEST_ALLOW_DESTRUCTIVE"

func openDisposableIntegrationDB(t *testing.T, pgurl string) {
	t.Helper()
	if os.Getenv(destructiveIntegrationOptIn) != "yes" {
		t.Fatalf("set %s=yes to run destructive database integration tests", destructiveIntegrationOptIn)
	}
	if err := InitDB(pgurl); err != nil {
		t.Fatalf("initialize database: %v", err)
	}
	t.Cleanup(func() {
		CloseDB()
		DB = nil
	})

	var databaseName string
	if err := QueryRowContext(context.Background(), "SELECT current_database()").Scan(&databaseName); err != nil {
		t.Fatalf("identify integration database: %v", err)
	}
	if !isDisposableIntegrationDatabase(databaseName) {
		t.Fatalf("refusing destructive integration test against database %q; use pgext_test or pgext_test_*", databaseName)
	}

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = ExecContext(ctx, "DROP SCHEMA IF EXISTS pgext CASCADE")
	})
}

func isDisposableIntegrationDatabase(name string) bool {
	return name == "pgext_test" || strings.HasPrefix(name, "pgext_test_")
}

func TestDisposableIntegrationDatabaseName(t *testing.T) {
	for _, tt := range []struct {
		name string
		want bool
	}{
		{name: "pgext_test", want: true},
		{name: "pgext_test_review_123", want: true},
		{name: "data", want: false},
		{name: "postgres", want: false},
		{name: "template1", want: false},
		{name: "pgext_testing", want: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDisposableIntegrationDatabase(tt.name); got != tt.want {
				t.Fatalf("isDisposableIntegrationDatabase(%q) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
