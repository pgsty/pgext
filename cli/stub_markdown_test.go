package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateStubMarkdown(t *testing.T) {
	for _, tt := range []struct {
		name, content, want string
	}{
		{"SQL", "## Usage\n\n```sql\nSELECT 1;\n```\n", ""},
		{"Chinese", "## 用法\n\n```text\n查询结果\n```\n", ""},
		{"hard break and code whitespace", "## Usage\n\nProse  \nnext line\n\n```text\noutput \n```\n", ""},
		{"long fence", "## Usage\n\n````text\n```sql\nSELECT 1;\n```\n````\n", ""},
		{"tilde fence", "## Usage\n\n~~~ini\nsetting = on\n~~~\n", ""},
		{"comment example", "<!--\n## Literal heading\n```\n-->\n\n## Usage\n\nText\n", ""},
		{"missing language", "## Usage\n\n```\nresult\n```\n", ":3: MD110"},
		{"heading after", "## Usage\nText\n", ":1: MD118"},
		{"heading before", "Text\n### Example\n\n", ":2: MD117"},
		{"heading space", "##  Usage\n\nText\n", ":1: MD116"},
		{"trailing space", "## Usage\n\nText \n", ":3: MD109"},
		{"trailing tab", "Text\t\n", ":1: MD109"},
		{"unclosed fence", "```sql\nSELECT 1;\n", ":1: MD112"},
		{"wrong closing marker", "```sql\nSELECT 1;\n~~~\n", ":1: MD112"},
		{"closing info is not a close", "```text\n```sql\n", ":1: MD112"},
		{"CRLF", "## Usage\r\n\r\n```sql\r\nSELECT 1;\r\n```\r\n", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStubMarkdown("stub/example.md", tt.content)
			if tt.want == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), "stub/example.md"+tt.want) {
				t.Fatalf("error = %v, want source path and %s", err, tt.want)
			}
		})
	}
}

func TestReadValidatedStubReportsIOErrors(t *testing.T) {
	dir := t.TempDir()
	if content, err := ReadValidatedStub(filepath.Join(dir, "absent.md")); err != nil || content != "" {
		t.Fatalf("absent optional stub = %q, %v", content, err)
	}
	if _, err := ReadValidatedStub(dir); err == nil {
		t.Fatal("directory read error was swallowed")
	}
}

func TestPageGeneratorsRejectInvalidStubsBeforeOverwrite(t *testing.T) {
	for _, kind := range []string{"io", "cc", "catalog"} {
		t.Run(kind, func(t *testing.T) {
			base := t.TempDir()
			stubDir := filepath.Join(base, "stub")
			output := filepath.Join(t.TempDir(), "content", "e")
			for _, dir := range []string{stubDir, output} {
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatal(err)
				}
			}
			stubPath := filepath.Join(stubDir, "demo.md")
			if err := os.WriteFile(stubPath, []byte("## Usage\n\n```\nSELECT 1;\n```\n"), 0644); err != nil {
				t.Fatal(err)
			}
			pagePath := filepath.Join(output, "demo.md")
			if err := os.WriteFile(pagePath, []byte("previous valid page\n"), 0644); err != nil {
				t.Fatal(err)
			}
			ext := &Extension{Name: "demo", Pkg: "demo"}
			var err error
			switch kind {
			case "io":
				err = NewIOPageGenerator(&ExtensionCache{}, filepath.Dir(output), stubDir).GenerateExtensionPage(context.Background(), ext)
			case "cc":
				err = NewCCPageGenerator(&ExtensionCache{}, filepath.Dir(output), stubDir).GenerateExtensionPage(context.Background(), ext)
			case "catalog":
				g := NewExtensionGenerator(&ExtensionCache{}, output)
				g.StubDir = stubDir
				err = g.GenerateExtensionPage(context.Background(), ext)
			}
			if err == nil || !strings.Contains(err.Error(), stubPath+":3: MD110") {
				t.Fatalf("generation error = %v, want invalid stub diagnostic before database access", err)
			}
			body, err := os.ReadFile(pagePath)
			if err != nil || string(body) != "previous valid page\n" {
				t.Fatalf("existing page was changed: %q, %v", body, err)
			}
		})
	}
}

func TestValidateExtensionStubsChecksCompleteBatch(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"first", "second"} {
		if err := os.WriteFile(filepath.Join(dir, name+".md"), []byte("```\nresult\n```\n"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	err := ValidateExtensionStubs(dir, []*Extension{{Name: "first"}, {Name: "second"}})
	if err == nil || !strings.Contains(err.Error(), "first.md:1: MD110") || !strings.Contains(err.Error(), "second.md:1: MD110") {
		t.Fatalf("incomplete batch diagnostics: %v", err)
	}
}

func TestLintStubPathsRejectsMissingPaths(t *testing.T) {
	if _, err := LintStubPaths([]string{filepath.Join(t.TempDir(), "missing")}); err == nil {
		t.Fatal("a nonexistent source directory passed lint")
	}
}

func TestLintStubPairContract(t *testing.T) {
	dir := t.TempDir()
	enDir := filepath.Join(dir, "stub")
	zhDir := filepath.Join(dir, "stub-zh")
	if err := os.MkdirAll(enDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(zhDir, 0755); err != nil {
		t.Fatal(err)
	}

	validEN := "## Usage\n\nSources:\n\n- [README](https://github.com/example/demo/blob/v1.0.0/README.md)\n\nThe `demo` extension runs `demo_fn()`.\n\n```sql\nCREATE EXTENSION demo;\nSELECT demo_fn();\n```\n"
	validZH := "## 用法\n\n来源：\n\n- [README](https://github.com/example/demo/blob/v1.0.0/README.md)\n\n`demo` 扩展提供 `demo_fn()` 函数用于演示。\n\n```sql\nCREATE EXTENSION demo;\nSELECT demo_fn();\n```\n"

	writePair := func(name, en, zh string) (string, string) {
		enP := filepath.Join(enDir, name+".md")
		zhP := filepath.Join(zhDir, name+".md")
		if err := os.WriteFile(enP, []byte(en), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(zhP, []byte(zh), 0644); err != nil {
			t.Fatal(err)
		}
		return enP, zhP
	}

	enP, zhP := writePair("demo", validEN, validZH)
	if err := LintStubPair(enP, zhP, "demo", nil); err != nil {
		t.Fatalf("valid pair rejected: %v", err)
	}

	// need_ddl=false must forbid CREATE EXTENSION in fenced code
	needDDL := false
	if err := LintStubPair(enP, zhP, "demo", &needDDL); err == nil || !strings.Contains(err.Error(), "CREATE EXTENSION") {
		t.Fatalf("expected need_ddl=false error, got: %v", err)
	}

	// Mismatched inline identifier, URL, and untranslated English paragraph in ZH
	badZH := strings.ReplaceAll(validZH, "`demo_fn()`", "`other_fn()`")
	badZH += "\nThis extension provides a complete workflow for running queries with options.\n"
	enP2, zhP2 := writePair("bad", validEN, badZH)
	err := LintStubPair(enP2, zhP2, "demo", nil)
	if err == nil || !strings.Contains(err.Error(), "inline code identifiers differ") || !strings.Contains(err.Error(), "untranslated English paragraph") {
		t.Fatalf("expected pair & language errors, got: %v", err)
	}

	// Untranslated Han prose in EN
	badEN := validEN + "\n这是一段没有翻译成英文的中文段落说明。\n"
	enP3, zhP3 := writePair("han", badEN, validZH)
	if err := LintStubPair(enP3, zhP3, "demo", nil); err == nil || !strings.Contains(err.Error(), "untranslated Han text") {
		t.Fatalf("expected EN Han prose error, got: %v", err)
	}
}
