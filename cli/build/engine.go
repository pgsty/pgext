package build

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type Options struct {
	CatalogFile, PG, SpecDir, Root, Platform, SourceURL, Checksums string
	Mirror, Force, NoDebug, Plan                                   bool
	Jobs                                                           int
	ReplaceRepos                                                   bool
}

type Job struct {
	Package          string     `json:"package"`
	Requested        []string   `json:"requested"`
	PG               []int      `json:"pg,omitempty"`
	Recipe           string     `json:"recipe,omitempty"`
	RecipeSHA256     string     `json:"recipe_sha256,omitempty"`
	RecipeTreeSHA256 string     `json:"recipe_tree_sha256,omitempty"`
	Sources          []string   `json:"sources"`
	Stages           []string   `json:"stages"`
	Extension        *Extension `json:"-"`
}
type Plan struct {
	Action   string   `json:"action"`
	Platform Platform `json:"platform"`
	Catalog  *Catalog `json:"catalog"`
	Root     string   `json:"root"`
	Jobs     []*Job   `json:"jobs"`
}
type StageResult struct {
	Package  string  `json:"package"`
	Stage    string  `json:"stage"`
	PG       int     `json:"pg,omitempty"`
	Success  bool    `json:"success"`
	Error    string  `json:"error,omitempty"`
	Duration float64 `json:"seconds"`
	Log      string  `json:"log,omitempty"`
}
type File struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
	Size   int64  `json:"size"`
}
type Artifact struct {
	File
	Package string `json:"package"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}
type Result struct {
	Plan
	Success      bool              `json:"success"`
	Error        string            `json:"error,omitempty"`
	Manifest     string            `json:"manifest,omitempty"`
	Stages       []StageResult     `json:"stages"`
	Sources      []File            `json:"sources"`
	Artifacts    []Artifact        `json:"artifacts"`
	ToolVersions map[string]string `json:"tools,omitempty"`
	Image        string            `json:"image,omitempty"`
}

type Engine struct {
	Options    Options
	Platform   Platform
	Catalog    *Catalog
	Exec       Executor
	Err        io.Writer
	Version    string
	aptUpdated bool
}

func New(opts Options, diagnostics io.Writer) (*Engine, error) {
	e, err := NewEnvironment(opts, diagnostics)
	if err != nil {
		return nil, err
	}
	e.Catalog, err = LoadCatalog(opts.CatalogFile)
	return e, err
}

// NewEnvironment does not load extension metadata: repairing a builder must
// remain possible when its local catalog is absent or invalid.
func NewEnvironment(opts Options, diagnostics io.Writer) (*Engine, error) {
	p, err := DetectPlatform()
	if err != nil {
		return nil, err
	}
	if opts.Platform != "" {
		q, e := PlanPlatform(opts.Platform)
		if e != nil {
			return nil, e
		}
		if !opts.Plan && (q.OS != p.OS || q.Arch != p.Arch) {
			return nil, fmt.Errorf("--platform can select a different system only with --plan")
		}
		if opts.Plan && (q.OS != p.OS || q.Arch != p.Arch) {
			p = q
		}
	}
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	if opts.Root == "" {
		opts.Root = filepath.Join(h, "ext")
	}
	opts.Root, err = filepath.Abs(opts.Root)
	if err != nil {
		return nil, err
	}
	if opts.SpecDir == "" {
		if p.Family == "rpm" {
			opts.SpecDir = filepath.Join(h, "rpmbuild")
		} else {
			opts.SpecDir = filepath.Join(h, "debbuild")
		}
	}
	opts.SpecDir, err = filepath.Abs(opts.SpecDir)
	if err != nil {
		return nil, err
	}
	if opts.Jobs < 0 {
		return nil, fmt.Errorf("jobs cannot be negative")
	}
	if _, err := ParsePG(opts.PG); err != nil {
		return nil, err
	}
	if diagnostics == nil {
		diagnostics = io.Discard
	}
	return &Engine{Options: opts, Platform: p, Exec: NativeExecutor{}, Err: diagnostics}, nil
}

func (e *Engine) Prepare(action string, names []string) (*Plan, error) {
	if _, err := loadChecksums(e.Options.Checksums); err != nil {
		return nil, err
	}
	stages := map[string][]string{"get": {"get"}, "dep": {"dep"}, "make": {"make"}, "pkg": {"get", "dep", "make"}}[action]
	if stages == nil {
		return nil, fmt.Errorf("unknown action %s", action)
	}
	if len(names) == 0 {
		return nil, fmt.Errorf("specify at least one extension or package")
	}
	if action != "get" && e.Platform.Family == "" {
		return nil, fmt.Errorf("package operations require a supported Linux builder; use --plan --platform for a preview")
	}
	p := &Plan{Action: action, Platform: e.Platform, Catalog: e.Catalog, Root: e.Options.Root}
	seen := map[string]*Job{}
	for _, name := range names {
		if !safeName(name) {
			return nil, fmt.Errorf("invalid extension/package/source name %q", name)
		}
		j := &Job{Package: name, Requested: []string{name}, Stages: stages, Extension: e.Catalog.Resolve(name)}
		if j.Extension != nil {
			j.Package = j.Extension.Package
		}
		if previous := seen[j.Package]; previous != nil {
			previous.Requested = append(previous.Requested, name)
			continue
		}
		if e.Platform.Family == "rpm" {
			j.Recipe = filepath.Join(e.Options.SpecDir, "SPECS", j.Package+".spec")
		} else {
			j.Recipe = filepath.Join(e.Options.SpecDir, j.Package, "Makefile")
		}
		raw, err := os.ReadFile(j.Recipe)
		if err != nil {
			if action != "get" {
				return nil, fmt.Errorf("recipe for %s: %w; run pgext env spec", name, err)
			}
			j.Recipe = ""
		} else {
			j.RecipeSHA256 = Hash(raw)
			j.RecipeTreeSHA256, err = recipeTreeHash(filepath.Dir(j.Recipe))
			if err != nil {
				return nil, err
			}
		}
		if action != "get" || e.Options.PG != "" || (j.Recipe != "" && e.Platform.Family != "") {
			selector := *e
			if action == "get" {
				// Source preparation must also work before development tools are
				// installed, and fetch every declared per-major source variant.
				selector.Options.Plan = true
			}
			j.PG, err = selector.jobPG(j, raw)
			if err != nil {
				return nil, err
			}
		}
		j.Sources = e.sourceNames(j, raw)
		if action == "get" && len(j.Sources) == 0 && j.Recipe == "" {
			return nil, fmt.Errorf("no source declared for %s", name)
		}
		for _, s := range j.Sources {
			if !safeName(s) {
				return nil, fmt.Errorf("unsafe source filename %q", s)
			}
		}
		seen[j.Package] = j
		p.Jobs = append(p.Jobs, j)
	}
	return p, nil
}

func (e *Engine) jobPG(j *Job, recipe []byte) ([]int, error) {
	requested, err := ParsePG(e.Options.PG)
	if err != nil {
		return nil, err
	}
	supported := recipePG(e.Platform.Family, e.Options.SpecDir, j.Package, recipe)
	// A recipe's declared targets are authoritative. The source compatibility
	// list is only a fallback; rpm_pg/deb_pg describe published packages.
	if len(supported) == 0 && j.Extension != nil {
		supported = j.Extension.PG
	}
	perPG := strings.Contains(string(recipe), "pgmajorversion") || strings.Contains(string(recipe), "PG_VERSIONS") || strings.Contains(string(recipe), "PG_MAJOR")
	if e.Platform.Family == "deb" {
		_, err := os.Stat(filepath.Join(filepath.Dir(j.Recipe), "debian", "pgversions"))
		perPG = perPG || err == nil
	}
	if j.Extension != nil {
		template := j.Extension.DEB
		if e.Platform.Family == "rpm" {
			template = j.Extension.RPM
		}
		perPG = perPG || strings.Contains(template, "$v")
	}
	if !perPG {
		if len(requested) > 0 {
			return nil, fmt.Errorf("%s is not a PostgreSQL-major-specific recipe; omit --pg", j.Package)
		}
		return nil, nil
	}
	if len(requested) > 0 {
		for _, pg := range requested {
			if len(supported) > 0 && !slices.Contains(supported, pg) {
				return nil, fmt.Errorf("%s recipe does not support PG%d (declared: %v)", j.Package, pg, supported)
			}
		}
		return requested, nil
	}
	if len(supported) == 0 {
		supported = []int{14, 15, 16, 17, 18}
	}
	var out []int
	for _, pg := range supported {
		if pg < 14 || pg > 18 {
			continue
		}
		if e.Options.Plan {
			out = append(out, pg)
		} else if _, err := os.Stat(e.Platform.PGConfig(pg)); err == nil {
			out = append(out, pg)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no supported PostgreSQL development installation for %s; run pgext env tool or specify --pg", j.Package)
	}
	return out, nil
}

// Run emits a result even on failure. No stage invokes another stage implicitly.
func (e *Engine) Run(ctx context.Context, p *Plan) (r *Result, runErr error) {
	r = &Result{Plan: *p, Stages: []StageResult{}, Sources: []File{}, Artifacts: []Artifact{}, ToolVersions: map[string]string{"pgext": e.Version}, Image: os.Getenv("PGEXT_BUILD_IMAGE")}
	defer func() {
		if runErr != nil {
			r.Success, r.Error = false, runErr.Error()
		}
		if r.Manifest != "" {
			if err := writeJSON(r.Manifest, r); err != nil {
				runErr = errors.Join(runErr, fmt.Errorf("save result: %w", err))
				r.Success, r.Error = false, runErr.Error()
			}
		}
	}()
	if e.Options.Plan {
		r.Success = true
		return r, nil
	}
	if err := os.MkdirAll(filepath.Join(e.Options.Root, "log"), 0755); err != nil {
		return r, err
	}
	run, err := os.MkdirTemp(filepath.Join(e.Options.Root, "log"), "run-")
	if err != nil {
		return r, err
	}
	r.Manifest = filepath.Join(run, "result.json")
	for _, j := range p.Jobs {
		for _, stage := range j.Stages {
			if j.Recipe != "" {
				b, err := os.ReadFile(j.Recipe)
				if err != nil {
					return r, err
				}
				tree, err := recipeTreeHash(filepath.Dir(j.Recipe))
				if err != nil {
					return r, err
				}
				if Hash(b) != j.RecipeSHA256 || tree != j.RecipeTreeSHA256 {
					return r, fmt.Errorf("recipe changed after planning: %s", j.Recipe)
				}
			}
			start := time.Now()
			s := StageResult{Package: j.Package, Stage: stage, Log: filepath.Join(run, j.Package+"-"+stage+".log")}
			log, logErr := os.Create(s.Log)
			if logErr != nil {
				return r, logErr
			}
			w := io.MultiWriter(e.Err, log)
			fmt.Fprintf(w, "pgext %s %s: %s PG%v\n", p.Action, j.Package, stage, j.PG)
			switch stage {
			case "get":
				err = e.get(ctx, j, r, w)
			case "dep":
				err = e.dep(ctx, j, w)
			case "make":
				err = e.make(ctx, j, r, w)
			}
			_ = log.Close()
			s.Success = err == nil
			s.Duration = time.Since(start).Seconds()
			if err != nil {
				s.Error = err.Error()
				r.Error = err.Error()
			}
			r.Stages = append(r.Stages, s)
			if err != nil {
				return r, err
			}
		}
	}
	r.Success = true
	return r, nil
}

func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return AtomicWrite(path, append(b, '\n'), 0644)
}
