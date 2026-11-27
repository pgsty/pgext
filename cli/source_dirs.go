package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExpandDir resolves a CLI filesystem path, including a leading home shorthand.
func ExpandDir(dir string) (string, error) {
	if dir == "~" || strings.HasPrefix(dir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(home, strings.TrimPrefix(strings.TrimPrefix(dir, "~"), "/"))
	}
	return filepath.Abs(dir)
}

// SourceDir locates the editable pgext catalog and usage sources.
func SourceDir() (string, error) {
	if dir := os.Getenv("PGEXT_SOURCE_DIR"); dir != "" {
		return ExpandDir(dir)
	}
	// Support a source checkout anywhere, including CI and Git worktrees.
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "db", "schema.sql")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "cmd", "gen.go")); err == nil {
				return dir, nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ExpandDir("~/pgsty/pgext")
}

// SourceStubDir locates a language directory and detects a missing checkout before generation.
func SourceStubDir(explicit, name string) (string, error) {
	var dir string
	var err error
	if explicit != "" {
		dir, err = ExpandDir(explicit)
	} else {
		dir, err = SourceDir()
		dir = filepath.Join(dir, name)
	}
	if err != nil {
		return "", err
	}
	info, err := os.Stat(dir)
	if err != nil {
		return "", fmt.Errorf("locate usage sources (set PGEXT_SOURCE_DIR or --stub-dir): %w", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("usage source %s is not a directory", dir)
	}
	return dir, nil
}
