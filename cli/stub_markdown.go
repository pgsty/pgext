package cli

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var stubFenceRE = regexp.MustCompile("^ {0,3}(`{3,}|~{3,})(.*)$")
var stubHeadingRE = regexp.MustCompile(`^ {0,3}#{1,6}([ \t]+)\S`)

// ValidateStubMarkdown checks the source formatting that must survive copying
// a stub into a Hugo page. Code bodies and explicit Markdown hard breaks are
// preserved; language labels, heading spacing, and prose whitespace are checked.
func ValidateStubMarkdown(path, content string) error {
	lines := strings.Split(strings.TrimSuffix(content, "\n"), "\n")
	var issues []error
	var fence string
	fenceLine := 0
	inComment := false
	for index, raw := range lines {
		line := strings.TrimSuffix(raw, "\r")
		match := stubFenceRE.FindStringSubmatch(line)
		if fence != "" {
			if match != nil && match[1][0] == fence[0] && len(match[1]) >= len(fence) && strings.TrimSpace(match[2]) == "" {
				fence = ""
			}
			continue
		}
		trimmed := strings.TrimSpace(line)
		if inComment || strings.HasPrefix(trimmed, "<!--") {
			inComment = !strings.Contains(line, "-->")
			continue
		}
		addIssue := func(code, message string) {
			issues = append(issues, fmt.Errorf("%s:%d: %s %s", path, index+1, code, message))
		}
		if match != nil {
			fence, fenceLine = match[1], index+1
			if strings.TrimSpace(match[2]) == "" {
				addIssue("MD110", "fenced code block has no language; use sql, bash, ini, or text as appropriate")
			}
			continue
		}
		if heading := stubHeadingRE.FindStringSubmatch(line); heading != nil {
			if heading[1] != " " {
				addIssue("MD116", "heading marker needs one space")
			}
			if index > 0 && strings.TrimSpace(lines[index-1]) != "" {
				addIssue("MD117", "heading needs a blank line before it")
			}
			if index+1 < len(lines) && strings.TrimSpace(lines[index+1]) != "" {
				addIssue("MD118", "heading needs a blank line after it")
			}
		}
		trailing := strings.TrimRight(line, " \t")
		if suffix := line[len(trailing):]; suffix != "" && suffix != "  " {
			addIssue("MD109", "invalid trailing whitespace")
		}
	}
	if fence != "" {
		issues = append(issues, fmt.Errorf("%s:%d: MD112 unclosed fenced code block", path, fenceLine))
	}
	return errors.Join(issues...)
}

// ReadValidatedStub keeps absent optional usage stubs compatible with existing
// generators, but never silently copies invalid Markdown or ignores I/O errors.
func ReadValidatedStub(path string) (string, error) {
	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("read stub %s: %w", path, err)
	}
	if err := ValidateStubMarkdown(path, string(content)); err != nil {
		return "", err
	}
	return string(content), nil
}

// ValidateExtensionStubs runs before a batch writes any generated site files.
func ValidateExtensionStubs(stubDir string, extensions []*Extension) error {
	var issues []error
	for _, ext := range extensions {
		if _, err := ReadValidatedStub(filepath.Join(stubDir, ext.Name+".md")); err != nil {
			issues = append(issues, err)
		}
	}
	return errors.Join(issues...)
}

// LintStubPaths validates explicit files or Markdown directories without a DB.
func LintStubPaths(paths []string) (int, error) {
	count := 0
	var issues []error
	for _, path := range paths {
		err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}
			count++
			if _, err := ReadValidatedStub(path); err != nil {
				issues = append(issues, err)
			}
			return nil
		})
		if err != nil {
			issues = append(issues, fmt.Errorf("scan stubs %s: %w", path, err))
		}
	}
	return count, errors.Join(issues...)
}
