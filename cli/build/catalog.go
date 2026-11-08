// Package build implements database-free extension builds. Published package
// availability is deliberately not an input to source-build target selection.
package build

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"pgext/db"
)

type Extension struct {
	Name, Package, Lead, Source, RPM, DEB string
	PG                                    []int
	HasLib, HasBin                        bool
	Extra                                 map[string]any
}

type Catalog struct {
	Path       string                `json:"path"`
	SHA256     string                `json:"sha256"`
	Extensions map[string]*Extension `json:"-"`
	Packages   map[string]*Extension `json:"-"`
}

func ConfigDir() (string, error) {
	if d := os.Getenv("XDG_CONFIG_HOME"); d != "" {
		if !filepath.IsAbs(d) {
			return "", fmt.Errorf("XDG_CONFIG_HOME must be absolute")
		}
		return filepath.Join(d, "pgext"), nil
	}
	h, err := os.UserHomeDir()
	return filepath.Join(h, ".config", "pgext"), err
}

// LoadCatalog never falls back after an explicitly selected or existing local
// catalog fails validation. PIG's optional extra.matrix is accepted but ignored.
func LoadCatalog(path string) (*Catalog, error) {
	var b []byte
	var err error
	if path == "" {
		d, e := ConfigDir()
		if e != nil {
			return nil, e
		}
		p := filepath.Join(d, "extension.csv")
		if _, e = os.Stat(p); e == nil {
			path = p
		} else if !os.IsNotExist(e) {
			return nil, e
		}
	}
	if path != "" {
		b, err = os.ReadFile(path)
	} else {
		b, err = db.ReadFile("extension.csv")
		path = "embedded:db/extension.csv"
	}
	if err != nil {
		return nil, err
	}
	c, err := ParseCatalog(b)
	if err != nil {
		return nil, fmt.Errorf("catalog %s: %w", path, err)
	}
	c.Path = path
	return c, nil
}

func ParseCatalog(b []byte) (*Catalog, error) {
	r := csv.NewReader(bytes.NewReader(b))
	head, err := r.Read()
	if err != nil {
		return nil, err
	}
	cols := map[string]int{}
	for i, h := range head {
		if _, ok := cols[h]; ok {
			return nil, fmt.Errorf("duplicate column %q", h)
		}
		cols[h] = i
	}
	for _, k := range []string{"name", "pkg", "source"} {
		if _, ok := cols[k]; !ok {
			return nil, fmt.Errorf("missing column %s", k)
		}
	}
	c := &Catalog{SHA256: Hash(b), Extensions: map[string]*Extension{}, Packages: map[string]*Extension{}}
	for line := 2; ; line++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", line, err)
		}
		get := func(k string) string {
			if i, ok := cols[k]; ok {
				return row[i]
			}
			return ""
		}
		e := &Extension{Name: get("name"), Package: get("pkg"), Lead: get("lead_ext"), Source: get("source"), RPM: get("rpm_pkg"), DEB: get("deb_pkg"), HasLib: get("has_lib") == "t", HasBin: get("has_bin") == "t", Extra: map[string]any{}}
		if !safeName(e.Name) || !safeName(e.Package) {
			return nil, fmt.Errorf("row %d: invalid extension/package name", line)
		}
		if _, ok := c.Extensions[e.Name]; ok {
			return nil, fmt.Errorf("duplicate extension %s", e.Name)
		}
		if s := get("extra"); s != "" {
			if err := json.Unmarshal([]byte(s), &e.Extra); err != nil {
				return nil, fmt.Errorf("%s extra: %w", e.Name, err)
			}
		}
		if s := get("pg_ver"); s != "" {
			for _, v := range strings.Split(strings.Trim(s, "{}"), ",") {
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("%s pg_ver: %w", e.Name, err)
				}
				e.PG = append(e.PG, n)
			}
		}
		c.Extensions[e.Name] = e
		if c.Packages[e.Package] == nil || e.Lead == e.Name || get("lead") == "t" {
			c.Packages[e.Package] = e
		}
	}
	if len(c.Extensions) == 0 {
		return nil, fmt.Errorf("empty extension catalog")
	}
	return c, nil
}

func (c *Catalog) Resolve(name string) *Extension {
	if e := c.Extensions[name]; e != nil {
		if lead := c.Extensions[e.Lead]; lead != nil {
			return lead
		}
		return e
	}
	return c.Packages[name]
}

var namePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_.+~-]*$`)

func safeName(s string) bool { return namePattern.MatchString(s) && s != "." && s != ".." }
func Hash(b []byte) string   { s := sha256.Sum256(b); return hex.EncodeToString(s[:]) }

func ParsePG(s string) ([]int, error) {
	if s == "" {
		return nil, nil
	}
	var out []int
	seen := map[int]bool{}
	for _, v := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil || n < 14 || n > 18 {
			return nil, fmt.Errorf("invalid PostgreSQL major %q (supported: 14–18)", v)
		}
		if !seen[n] {
			out = append(out, n)
			seen[n] = true
		}
	}
	return out, nil
}

func pgStrings(vs []int) []string {
	out := make([]string, len(vs))
	for i, v := range vs {
		out[i] = strconv.Itoa(v)
	}
	return out
}
