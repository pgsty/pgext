package cli

import (
	"context"
	"strings"
	"testing"
)

func TestDocsyPageGeneratorsReturnDatabaseErrors(t *testing.T) {
	originalDB := DB
	DB = nil
	t.Cleanup(func() { DB = originalDB })

	ext := &Extension{Name: "demo", Pkg: "demo"}
	cache := &ExtensionCache{}
	tests := []struct {
		name     string
		generate func() error
	}{
		{
			name: "io",
			generate: func() error {
				return NewIOPageGenerator(cache, t.TempDir(), t.TempDir()).GenerateExtensionPage(context.Background(), ext)
			},
		},
		{
			name: "cc",
			generate: func() error {
				return NewCCPageGenerator(cache, t.TempDir(), t.TempDir()).GenerateExtensionPage(context.Background(), ext)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.generate()
			if err == nil {
				t.Fatal("page generation swallowed a database error")
			}
			if !strings.Contains(err.Error(), "database not initialized") {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
