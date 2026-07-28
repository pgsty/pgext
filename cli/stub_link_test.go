package cli

import (
	"bufio"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var markdownDestinationRE = regexp.MustCompile(`!?\[[^\]\n]*\]\(\s*<?([^)\s>]+)`)

func TestStubLinksDoNotDependOnSourcePath(t *testing.T) {
	for _, dir := range []string{"../stub", "../stub-zh"} {
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatalf("read %s: %v", dir, err)
		}
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			path := filepath.Join(dir, entry.Name())
			file, err := os.Open(path)
			if err != nil {
				t.Errorf("open %s: %v", path, err)
				continue
			}

			scanner := bufio.NewScanner(file)
			inFence := false
			lineNumber := 0
			for scanner.Scan() {
				lineNumber++
				line := scanner.Text()
				trimmed := strings.TrimSpace(line)
				if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
					inFence = !inFence
					continue
				}
				if inFence {
					continue
				}
				for _, match := range markdownDestinationRE.FindAllStringSubmatch(line, -1) {
					destination := strings.Trim(match[1], "<>")
					parsed, err := url.Parse(destination)
					if err == nil && (parsed.IsAbs() ||
						strings.HasPrefix(destination, "/") ||
						strings.HasPrefix(destination, "#")) {
						continue
					}
					t.Errorf(
						"%s:%d: path-relative link %q will be resolved from the generated page; use an absolute upstream URL, a site-root URL, or a fragment",
						path, lineNumber, destination,
					)
				}
			}
			if err := scanner.Err(); err != nil {
				t.Errorf("scan %s: %v", path, err)
			}
			if err := file.Close(); err != nil {
				t.Errorf("close %s: %v", path, err)
			}
		}
	}
}
