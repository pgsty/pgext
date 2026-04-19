package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOverviewPagesOnlyEmitHTMLOutput(t *testing.T) {
	t.Run("cc", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "_index.md")
		g := &CCListGenerator{
			Cache: &ExtensionCache{},
		}

		if err := g.GenerateOverviewPage(path); err != nil {
			t.Fatalf("GenerateOverviewPage() error = %v", err)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("os.ReadFile() error = %v", err)
		}

		text := string(content)
		if !strings.Contains(text, extOverviewOutputsFrontMatter) {
			t.Fatalf("cc overview front matter missing HTML-only outputs config:\n%s", text)
		}
		if strings.Contains(text, "print") {
			t.Fatalf("cc overview front matter must not enable print output:\n%s", text)
		}
	})

	t.Run("io", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "_index.md")
		g := &IOListGenerator{
			Cache: &ExtensionCache{},
		}

		if err := g.GenerateOverviewPage(path); err != nil {
			t.Fatalf("GenerateOverviewPage() error = %v", err)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("os.ReadFile() error = %v", err)
		}

		text := string(content)
		if !strings.Contains(text, extOverviewOutputsFrontMatter) {
			t.Fatalf("io overview front matter missing HTML-only outputs config:\n%s", text)
		}
		if strings.Contains(text, "print") {
			t.Fatalf("io overview front matter must not enable print output:\n%s", text)
		}
	})
}
