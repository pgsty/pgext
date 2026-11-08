package build

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var tarballPattern = regexp.MustCompile(`[A-Za-z0-9][A-Za-z0-9_.+~-]*\.(?:tar\.(?:gz|bz2|xz|zst)|tgz|zip)`)
var auxiliarySourcePattern = regexp.MustCompile(`[A-Za-z0-9][A-Za-z0-9_.+~-]*\.(?:whl|dsc|jar(?:\.sha256)?)`)
var sourcePathPattern = regexp.MustCompile(`(?:\.\./SOURCES/|/ext/src/)([^\s"'();|&<>]+)`)

func (e *Engine) sourceNames(j *Job, recipe []byte) []string {
	var sources []string
	// Literal recipe tarballs include per-major inputs, e.g. Citus PG15.
	if len(recipe) > 0 && e.Platform.Family == "deb" {
		sources = recipeSources(string(recipe))
	}
	if len(sources) == 0 && j.Extension != nil {
		sources = strings.Fields(j.Extension.Source)
	}
	if len(sources) == 0 {
		sources = append(sources, SpecialSourceMapping[j.Package]...)
	}
	if len(sources) == 0 && tarballPattern.MatchString(j.Package) {
		sources = []string{j.Package}
	}
	if j.Package == "documentdb" && e.Platform.Family != "rpm" {
		for _, source := range documentDBSources {
			sources = append(sources, source.name)
		}
	}
	return unique(sources)
}

// Expand simple recipe constants without executing Makefile shell functions.
// Conditional recipes retain all literal source variants for offline reuse.
func recipeSources(content string) []string {
	// Ignore examples and historical filenames in whole-line comments.
	content = regexp.MustCompile(`(?m)^\s*#[^\n]*`).ReplaceAllString(content, "")
	vars := map[string]string{}
	assignment := regexp.MustCompile(`(?m)^([A-Za-z_][A-Za-z0-9_]*)\s*[:?]?=\s*([^\n]+)`)
	for _, m := range assignment.FindAllStringSubmatch(content, -1) {
		vars[m[1]] = strings.TrimSpace(strings.Split(m[2], "#")[0])
	}
	ref := regexp.MustCompile(`\$\(([A-Za-z_][A-Za-z0-9_]*)\)`)
	for i := 0; i < 8; i++ {
		next := ref.ReplaceAllStringFunc(content, func(s string) string {
			k := s[2 : len(s)-1]
			if v, ok := vars[k]; ok {
				return v
			}
			return s
		})
		if next == content {
			break
		}
		content = next
	}
	// Download URLs can use a different basename from the cached input (Valkey
	// tags are <version>.tar.gz). Do not treat those basenames as extra inputs.
	content = regexp.MustCompile(`https?://[^\s"'()]+`).ReplaceAllString(content, "")
	sources := tarballPattern.FindAllString(content, -1)
	// Wheels, Debian source descriptors and wrapper JARs are also build inputs.
	// Other literal SOURCES paths (e.g. RDKit's enum.h) need no suffix whitelist.
	for _, assignment := range assignment.FindAllStringSubmatch(content, -1) {
		sources = append(sources, auxiliarySourcePattern.FindAllString(assignment[2], -1)...)
	}
	for _, match := range sourcePathPattern.FindAllStringSubmatch(content, -1) {
		if safeName(match[1]) && strings.Contains(match[1], ".") {
			sources = append(sources, match[1])
		}
	}
	return unique(sources)
}

// RPM supports a URL fragment to rename a source (archive.tar.gz#/local.tar.gz).
// get and mak must resolve exactly the same local filename.
func rpmSourceName(raw string) (string, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", err
	}
	path := u.Path
	if u.Fragment != "" {
		path = u.Fragment
	}
	name := filepath.Base(path)
	if !safeName(name) {
		return "", fmt.Errorf("invalid expanded Source/Patch: %s", raw)
	}
	return name, nil
}

func unique(in []string) []string {
	out := []string{}
	seen := map[string]bool{}
	for _, s := range in {
		if s != "" && !seen[s] {
			out = append(out, s)
			seen[s] = true
		}
	}
	return out
}

func (e *Engine) sourceBases() []string {
	if e.Options.SourceURL != "" {
		return []string{strings.TrimRight(e.Options.SourceURL, "/")}
	}
	urls := []string{"https://repo.pigsty.io/ext/src", "https://repo.pigsty.cc/ext/src"}
	if e.Options.Mirror {
		urls[0], urls[1] = urls[1], urls[0]
	}
	return urls
}

func loadChecksums(path string) (map[string]string, error) {
	m := map[string]string{}
	if path == "" {
		return m, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range m {
		if !safeName(k) || !regexp.MustCompile(`^[a-fA-F0-9]{64}$`).MatchString(v) {
			return nil, fmt.Errorf("invalid source checksum for %q", k)
		}
		m[k] = strings.ToLower(v)
	}
	return m, nil
}

func (e *Engine) get(ctx context.Context, j *Job, r *Result, w io.Writer) error {
	sums, err := loadChecksums(e.Options.Checksums)
	if err != nil {
		return err
	}
	if e.Platform.Family == "rpm" && j.Recipe != "" && !e.Options.Plan {
		pgs := j.PG
		if len(pgs) == 0 {
			pgs = []int{18}
		}
		var actual []string
		for _, pg := range pgs {
			out, err := e.Exec.Capture(ctx, Command{Args: []string{"rpmspec", "-P", "--define", fmt.Sprintf("pgmajorversion %d", pg), "--define", "_topdir " + e.Options.SpecDir, j.Recipe}})
			if err != nil {
				return err
			}
			for _, line := range strings.Split(out, "\n") {
				if regexp.MustCompile(`^Source[0-9]*\s*:`).MatchString(line) {
					_, s, _ := strings.Cut(line, ":")
					name, err := rpmSourceName(s)
					if err != nil {
						return err
					}
					actual = append(actual, name)
				}
			}
		}
		if len(actual) > 0 {
			j.Sources = unique(actual)
		}
	}
	if len(j.Sources) == 0 {
		return fmt.Errorf("no source declared for %s", j.Package)
	}
	for _, name := range j.Sources {
		if e.Options.Checksums != "" && sums[name] == "" {
			return fmt.Errorf("checksum manifest lacks %s", name)
		}
		f, err := Download(ctx, e.sourceBases(), name, filepath.Join(e.Options.Root, "src", name), sums[name], e.Options.Force, w)
		if err != nil {
			return err
		}
		r.Sources = append(r.Sources, f)
	}
	return nil
}

// Download validates cached bytes whenever an expected digest is supplied,
// stages downloads beside the destination, and replaces it only after success.
func Download(ctx context.Context, bases []string, name, path, expected string, force bool, w io.Writer) (File, error) {
	if !safeName(name) {
		return File{}, fmt.Errorf("invalid download filename")
	}
	if info, err := os.Lstat(path); err == nil {
		if !info.Mode().IsRegular() {
			return File{}, fmt.Errorf("source is not a regular file: %s", path)
		}
		if !force {
			f, err := fileInfo(path)
			if err != nil {
				return f, err
			}
			if f.Size == 0 {
				return f, fmt.Errorf("empty cached source: %s", path)
			}
			if expected != "" && f.SHA256 != expected {
				return f, fmt.Errorf("cached source checksum mismatch: %s", name)
			}
			return f, nil
		}
	} else if !os.IsNotExist(err) {
		return File{}, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return File{}, err
	}
	client := &http.Client{Timeout: 20 * time.Minute}
	var last error
	for _, base := range bases {
		u, err := url.Parse(strings.TrimRight(base, "/") + "/" + url.PathEscape(name))
		if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil || u.RawQuery != "" {
			return File{}, fmt.Errorf("invalid source base URL")
		}
		for attempt := 0; attempt < 3; attempt++ {
			if err := ctx.Err(); err != nil {
				return File{}, err
			}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
			if err != nil {
				return File{}, err
			}
			resp, err := client.Do(req)
			if err != nil {
				last = err
				continue
			}
			if resp.StatusCode != 200 {
				resp.Body.Close()
				last = fmt.Errorf("download %s: HTTP %d", name, resp.StatusCode)
				if resp.StatusCode < 500 {
					break
				}
				continue
			}
			f, err := os.CreateTemp(filepath.Dir(path), ".download-*")
			if err != nil {
				resp.Body.Close()
				return File{}, err
			}
			_, err = io.Copy(f, resp.Body)
			resp.Body.Close()
			closeErr := f.Close()
			if err == nil {
				err = closeErr
			}
			var result File
			if err == nil {
				result, err = fileInfo(f.Name())
			}
			if err == nil && result.Size == 0 {
				err = fmt.Errorf("empty downloaded source %s", name)
			}
			if err == nil && expected != "" && result.SHA256 != expected {
				err = fmt.Errorf("download checksum mismatch: %s", name)
			}
			if err == nil {
				err = os.Chmod(f.Name(), 0644)
			}
			if err == nil {
				err = os.Rename(f.Name(), path)
			}
			if err != nil {
				os.Remove(f.Name())
				last = err
				continue
			}
			result.Path = path
			fmt.Fprintf(w, "downloaded %s sha256=%s\n", name, result.SHA256)
			return result, nil
		}
	}
	return File{}, fmt.Errorf("fetch %s: %w", name, last)
}
