/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	DefaultPGXNIndexURL = "https://api.pgxn.org/src/"
	DefaultPGXNAPIBase  = "https://api.pgxn.org"
	DefaultPGXNSiteBase = "https://pgxn.org"
)

var pgxnIndexHrefRE = regexp.MustCompile(`<a href=['"]/src/([^/'"]+)/['"]>`)

// PgxnExportOptions controls PGXN catalog export.
type PgxnExportOptions struct {
	OutputDir   string
	CatalogPath string
	Workers     int
	Retry       int
	IndexURL    string
	APIBase     string
	SiteBase    string
}

// PgxnExportResult summarizes the generated PGXN export.
type PgxnExportResult struct {
	GeneratedAt     string            `json:"generated_at"`
	DistCount       int               `json:"dist_count"`
	MatchedCount    int               `json:"matched_count"`
	MissingCount    int               `json:"missing_count"`
	CatalogRows     int               `json:"catalog_rows"`
	CatalogPackages int               `json:"catalog_packages"`
	FailedCount     int               `json:"failed_count"`
	FailedDists     []string          `json:"failed_dists,omitempty"`
	Files           map[string]string `json:"files"`
}

// CatalogRecord is a lightweight view of db/extension.csv.
type CatalogRecord struct {
	ID       int
	Name     string
	Pkg      string
	LeadExt  string
	Category string
	State    string
	URL      string
	License  string
	Version  string
	Repo     string
	Lang     string
	Lead     bool
	Contrib  bool
}

type catalogIndex struct {
	rows          []CatalogRecord
	byPkg         map[string][]CatalogRecord
	byName        map[string][]CatalogRecord
	byCanonical   map[string][]CatalogRecord
	byURL         map[string][]CatalogRecord
	uniquePkgKeys map[string]struct{}
}

// PgxnMatch describes how a PGXN distribution maps to the current catalog.
type PgxnMatch struct {
	Matched bool
	Reason  string
	Score   int
	Row     CatalogRecord
}

// PgxnRecord contains normalized PGXN metadata plus raw JSON payloads.
type PgxnRecord struct {
	Dist    string
	DistURL string
	APIURL  string
	MetaURL string
	RawDist map[string]interface{}
	RawMeta map[string]interface{}
	Match   PgxnMatch
}

type pgxnExporter struct {
	opts   PgxnExportOptions
	client *http.Client
}

type pgxnFetchResult struct {
	record *PgxnRecord
	err    error
	dist   string
}

type scoredCatalogMatch struct {
	row    CatalogRecord
	score  int
	reason string
}

// ExportPGXNCatalog crawls PGXN and writes standalone export artifacts.
func ExportPGXNCatalog(ctx context.Context, opts PgxnExportOptions) (*PgxnExportResult, error) {
	exporter := newPGXNExporter(opts)
	return exporter.Run(ctx)
}

func newPGXNExporter(opts PgxnExportOptions) *pgxnExporter {
	if opts.OutputDir == "" {
		opts.OutputDir = filepath.Join("research", "pgxn")
	}
	if opts.CatalogPath == "" {
		opts.CatalogPath = filepath.Join("db", "extension.csv")
	}
	if opts.Workers <= 0 {
		opts.Workers = 16
	}
	if opts.Retry <= 0 {
		opts.Retry = 2
	}
	if opts.IndexURL == "" {
		opts.IndexURL = DefaultPGXNIndexURL
	}
	if opts.APIBase == "" {
		opts.APIBase = DefaultPGXNAPIBase
	}
	if opts.SiteBase == "" {
		opts.SiteBase = DefaultPGXNSiteBase
	}

	return &pgxnExporter{
		opts: opts,
		client: &http.Client{
			Timeout: 45 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        128,
				MaxIdleConnsPerHost: 32,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

func (e *pgxnExporter) Run(ctx context.Context) (*PgxnExportResult, error) {
	catalogRows, idx, err := loadCatalogIndex(e.opts.CatalogPath)
	if err != nil {
		return nil, fmt.Errorf("load catalog index: %w", err)
	}

	dists, err := e.discoverDists(ctx)
	if err != nil {
		return nil, fmt.Errorf("discover PGXN dists: %w", err)
	}
	logrus.Infof("discovered %d PGXN distributions from %s", len(dists), e.opts.IndexURL)

	records, failed := e.fetchAll(ctx, dists)
	if len(records) == 0 {
		return nil, fmt.Errorf("no PGXN records fetched successfully")
	}

	sort.Slice(records, func(i, j int) bool { return records[i].Dist < records[j].Dist })
	for _, record := range records {
		record.Match = matchPGXNRecord(record, idx)
	}

	matchedCount := 0
	var missing []*PgxnRecord
	for _, record := range records {
		if record.Match.Matched {
			matchedCount++
			continue
		}
		missing = append(missing, record)
	}
	sort.Slice(missing, func(i, j int) bool {
		return missing[i].dateValue().After(missing[j].dateValue())
	})

	if err := os.MkdirAll(e.opts.OutputDir, 0o755); err != nil {
		return nil, fmt.Errorf("create output dir: %w", err)
	}

	pkgMatches := groupMatchesByPackage(records)

	files := map[string]string{
		"full_csv":     filepath.Join(e.opts.OutputDir, "pgxn_dist.csv"),
		"full_jsonl":   filepath.Join(e.opts.OutputDir, "pgxn_dist.jsonl"),
		"missing_csv":  filepath.Join(e.opts.OutputDir, "pgxn_missing.csv"),
		"missing_md":   filepath.Join(e.opts.OutputDir, "pgxn_missing.md"),
		"catalog_csv":  filepath.Join(e.opts.OutputDir, "pgext_enriched.csv"),
		"summary_json": filepath.Join(e.opts.OutputDir, "summary.json"),
		"readme_md":    filepath.Join(e.opts.OutputDir, "README.md"),
	}

	if err := writePGXNCSV(files["full_csv"], records); err != nil {
		return nil, fmt.Errorf("write full csv: %w", err)
	}
	if err := writePGXNJSONL(files["full_jsonl"], records); err != nil {
		return nil, fmt.Errorf("write full jsonl: %w", err)
	}
	if err := writePGXNCSV(files["missing_csv"], missing); err != nil {
		return nil, fmt.Errorf("write missing csv: %w", err)
	}
	if err := writeMissingMarkdown(files["missing_md"], missing, len(records), matchedCount); err != nil {
		return nil, fmt.Errorf("write missing markdown: %w", err)
	}
	if err := writeCatalogEnrichedCSV(files["catalog_csv"], catalogRows, pkgMatches); err != nil {
		return nil, fmt.Errorf("write catalog csv: %w", err)
	}

	result := &PgxnExportResult{
		GeneratedAt:     time.Now().UTC().Format(time.RFC3339),
		DistCount:       len(records),
		MatchedCount:    matchedCount,
		MissingCount:    len(missing),
		CatalogRows:     len(catalogRows),
		CatalogPackages: len(idx.uniquePkgKeys),
		FailedCount:     len(failed),
		FailedDists:     failed,
		Files:           files,
	}

	if err := writeSummaryJSON(files["summary_json"], result); err != nil {
		return nil, fmt.Errorf("write summary json: %w", err)
	}
	if err := writePGXNReadme(files["readme_md"], result, e.opts); err != nil {
		return nil, fmt.Errorf("write readme: %w", err)
	}

	return result, nil
}

func (e *pgxnExporter) discoverDists(ctx context.Context) ([]string, error) {
	body, err := e.fetchURL(ctx, e.opts.IndexURL)
	if err != nil {
		return nil, err
	}

	matches := pgxnIndexHrefRE.FindAllStringSubmatch(string(body), -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no distributions found in %s", e.opts.IndexURL)
	}

	seen := make(map[string]struct{}, len(matches))
	dists := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		dist := strings.TrimSpace(match[1])
		if dist == "" {
			continue
		}
		if _, ok := seen[dist]; ok {
			continue
		}
		seen[dist] = struct{}{}
		dists = append(dists, dist)
	}
	sort.Strings(dists)
	return dists, nil
}

func (e *pgxnExporter) fetchAll(ctx context.Context, dists []string) ([]*PgxnRecord, []string) {
	results := make([]*PgxnRecord, len(dists))
	var failed []string
	var failedMu sync.Mutex

	jobs := make(chan int)
	var wg sync.WaitGroup
	var completed atomic.Int64

	for worker := 0; worker < e.opts.Workers; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				dist := dists[idx]
				record, err := e.fetchDist(ctx, dist)
				if err != nil {
					failedMu.Lock()
					failed = append(failed, fmt.Sprintf("%s: %v", dist, err))
					failedMu.Unlock()
				} else {
					results[idx] = record
				}

				done := completed.Add(1)
				if done%25 == 0 || done == int64(len(dists)) {
					logrus.Infof("PGXN crawl progress: %d/%d", done, len(dists))
				}
			}
		}()
	}

	for idx := range dists {
		jobs <- idx
	}
	close(jobs)
	wg.Wait()

	records := make([]*PgxnRecord, 0, len(dists))
	for _, result := range results {
		if result != nil {
			records = append(records, result)
		}
	}
	sort.Strings(failed)
	return records, failed
}

func (e *pgxnExporter) fetchDist(ctx context.Context, dist string) (*PgxnRecord, error) {
	apiURL := strings.TrimRight(e.opts.APIBase, "/") + "/dist/" + dist + ".json"
	distBody, err := e.fetchURL(ctx, apiURL)
	if err != nil {
		return nil, fmt.Errorf("fetch dist json: %w", err)
	}

	rawDist := make(map[string]interface{})
	if err := json.Unmarshal(distBody, &rawDist); err != nil {
		return nil, fmt.Errorf("decode dist json: %w", err)
	}

	version := getString(rawDist["version"])
	record := &PgxnRecord{
		Dist:    dist,
		DistURL: strings.TrimRight(e.opts.SiteBase, "/") + "/dist/" + dist + "/",
		APIURL:  apiURL,
		RawDist: rawDist,
	}

	if version == "" {
		return record, nil
	}

	metaURL := strings.TrimRight(e.opts.APIBase, "/") + "/dist/" + dist + "/" + version + "/META.json"
	record.MetaURL = metaURL

	metaBody, err := e.fetchURL(ctx, metaURL)
	if err != nil {
		logrus.Warnf("META.json unavailable for %s %s: %v", dist, version, err)
		return record, nil
	}

	rawMeta := make(map[string]interface{})
	if err := json.Unmarshal(metaBody, &rawMeta); err != nil {
		logrus.Warnf("failed to decode META.json for %s %s: %v", dist, version, err)
		return record, nil
	}
	record.RawMeta = rawMeta
	return record, nil
}

func (e *pgxnExporter) fetchURL(ctx context.Context, target string) ([]byte, error) {
	var lastErr error

	for attempt := 1; attempt <= e.opts.Retry; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "pgext-pgxn-crawler/1.0")

		resp, err := e.client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("unexpected status %s", resp.Status)
			continue
		}
		if readErr != nil {
			lastErr = readErr
			continue
		}
		return body, nil
	}

	return nil, lastErr
}

func loadCatalogIndex(path string) ([]CatalogRecord, *catalogIndex, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.ReuseRecord = false
	reader.FieldsPerRecord = -1

	header, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}

	col := make(map[string]int, len(header))
	for idx, name := range header {
		col[name] = idx
	}

	required := []string{"id", "name", "pkg", "lead_ext", "category", "state", "url", "license", "version", "repo", "lang", "lead", "contrib"}
	for _, name := range required {
		if _, ok := col[name]; !ok {
			return nil, nil, fmt.Errorf("missing column %q in %s", name, path)
		}
	}

	rows := make([]CatalogRecord, 0, 512)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		id, _ := strconv.Atoi(getCSVField(record, col, "id"))
		row := CatalogRecord{
			ID:       id,
			Name:     getCSVField(record, col, "name"),
			Pkg:      getCSVField(record, col, "pkg"),
			LeadExt:  getCSVField(record, col, "lead_ext"),
			Category: getCSVField(record, col, "category"),
			State:    getCSVField(record, col, "state"),
			URL:      getCSVField(record, col, "url"),
			License:  getCSVField(record, col, "license"),
			Version:  getCSVField(record, col, "version"),
			Repo:     getCSVField(record, col, "repo"),
			Lang:     getCSVField(record, col, "lang"),
			Lead:     parseCSVBool(getCSVField(record, col, "lead")),
			Contrib:  parseCSVBool(getCSVField(record, col, "contrib")),
		}
		rows = append(rows, row)
	}

	idx := &catalogIndex{
		rows:          rows,
		byPkg:         make(map[string][]CatalogRecord),
		byName:        make(map[string][]CatalogRecord),
		byCanonical:   make(map[string][]CatalogRecord),
		byURL:         make(map[string][]CatalogRecord),
		uniquePkgKeys: make(map[string]struct{}),
	}

	for _, row := range rows {
		pkgKey := strings.ToLower(strings.TrimSpace(row.Pkg))
		nameKey := strings.ToLower(strings.TrimSpace(row.Name))

		if pkgKey != "" {
			idx.byPkg[pkgKey] = append(idx.byPkg[pkgKey], row)
			idx.uniquePkgKeys[pkgKey] = struct{}{}
		}
		if nameKey != "" {
			idx.byName[nameKey] = append(idx.byName[nameKey], row)
		}

		for _, key := range uniqueNonEmpty(
			canonicalName(row.Pkg),
			canonicalName(row.Name),
			canonicalName(row.LeadExt),
		) {
			idx.byCanonical[key] = append(idx.byCanonical[key], row)
		}

		if urlKey := normalizeURL(row.URL); urlKey != "" {
			idx.byURL[urlKey] = append(idx.byURL[urlKey], row)
		}
	}

	return rows, idx, nil
}

func matchPGXNRecord(record *PgxnRecord, idx *catalogIndex) PgxnMatch {
	seen := make(map[string]scoredCatalogMatch)
	addCandidates := func(rows []CatalogRecord, score int, reason string) {
		for _, row := range rows {
			key := fmt.Sprintf("%d/%s/%s", row.ID, row.Name, row.Pkg)
			current, ok := seen[key]
			next := scoredCatalogMatch{row: row, score: score, reason: reason}
			if !ok || score > current.score || (score == current.score && betterCatalogTieBreak(next.row, current.row)) {
				seen[key] = next
			}
		}
	}

	distKey := strings.ToLower(record.Dist)
	addCandidates(idx.byPkg[distKey], 100, "pkg")
	addCandidates(idx.byName[distKey], 95, "name")

	for _, provided := range record.provides() {
		key := strings.ToLower(provided)
		addCandidates(idx.byName[key], 93, "provide.name")
		addCandidates(idx.byPkg[key], 90, "provide.pkg")
	}

	for _, rawName := range uniqueNonEmpty(record.displayName(), record.Dist) {
		if canonical := canonicalName(rawName); canonical != "" {
			addCandidates(idx.byCanonical[canonical], 70, "canonical")
		}
	}
	for _, provided := range record.provides() {
		if canonical := canonicalName(provided); canonical != "" {
			addCandidates(idx.byCanonical[canonical], 72, "provide.canonical")
		}
	}

	for _, rawURL := range record.resourceURLs() {
		if key := normalizeURL(rawURL); key != "" {
			addCandidates(idx.byURL[key], 88, "url")
		}
	}

	var best *scoredCatalogMatch
	for _, candidate := range seen {
		if best == nil || candidate.score > best.score || (candidate.score == best.score && betterCatalogTieBreak(candidate.row, best.row)) {
			next := candidate
			best = &next
		}
	}

	if best == nil {
		return PgxnMatch{}
	}

	return PgxnMatch{
		Matched: true,
		Reason:  best.reason,
		Score:   best.score,
		Row:     best.row,
	}
}

func betterCatalogTieBreak(a, b CatalogRecord) bool {
	if a.Lead != b.Lead {
		return a.Lead
	}
	if a.Contrib != b.Contrib {
		return !a.Contrib
	}
	if a.Pkg != b.Pkg {
		return a.Pkg < b.Pkg
	}
	return a.ID < b.ID
}

func groupMatchesByPackage(records []*PgxnRecord) map[string][]*PgxnRecord {
	grouped := make(map[string][]*PgxnRecord)
	for _, record := range records {
		if !record.Match.Matched {
			continue
		}
		pkg := record.Match.Row.Pkg
		if pkg == "" {
			continue
		}
		grouped[pkg] = append(grouped[pkg], record)
	}
	for _, recs := range grouped {
		sort.Slice(recs, func(i, j int) bool { return recs[i].Dist < recs[j].Dist })
	}
	return grouped
}

func writePGXNCSV(path string, records []*PgxnRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"dist", "display_name", "url", "api_url", "meta_url", "version", "release_status", "date",
		"user", "abstract", "description", "maintainers", "license_keys", "tags", "provides",
		"homepage", "repository_web", "repository_url", "repository_type", "bugtracker",
		"docs_json", "prereqs_json", "resources_json", "license_json", "provides_json",
		"releases_json", "special_files_json", "sha1",
		"matched", "match_reason", "match_score", "matched_id", "matched_name", "matched_pkg",
		"matched_url", "matched_version", "matched_repo", "matched_category",
		"raw_meta_json", "raw_dist_json",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, record := range records {
		row := []string{
			record.Dist,
			record.displayName(),
			record.DistURL,
			record.APIURL,
			record.MetaURL,
			record.version(),
			record.releaseStatus(),
			record.date(),
			record.user(),
			record.abstract(),
			record.description(),
			strings.Join(record.maintainers(), " | "),
			strings.Join(record.licenseKeys(), " | "),
			strings.Join(record.tags(), " | "),
			strings.Join(record.provides(), " | "),
			record.homepage(),
			record.repositoryWeb(),
			record.repositoryURL(),
			record.repositoryType(),
			record.bugtracker(),
			jsonCompact(record.rawValue("docs")),
			jsonCompact(record.rawValue("prereqs")),
			jsonCompact(record.rawValue("resources")),
			jsonCompact(record.rawValue("license")),
			jsonCompact(record.rawProvides()),
			jsonCompact(record.rawValue("releases")),
			jsonCompact(record.rawValue("special_files")),
			getString(record.rawValue("sha1")),
			strconv.FormatBool(record.Match.Matched),
			record.Match.Reason,
			strconv.Itoa(record.Match.Score),
			strconv.Itoa(record.Match.Row.ID),
			record.Match.Row.Name,
			record.Match.Row.Pkg,
			record.Match.Row.URL,
			record.Match.Row.Version,
			record.Match.Row.Repo,
			record.Match.Row.Category,
			jsonCompact(record.RawMeta),
			jsonCompact(record.RawDist),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return writer.Error()
}

func writePGXNJSONL(path string, records []*PgxnRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, record := range records {
		payload := map[string]interface{}{
			"dist":         record.Dist,
			"url":          record.DistURL,
			"api_url":      record.APIURL,
			"meta_url":     record.MetaURL,
			"version":      record.version(),
			"release_date": record.date(),
			"match": map[string]interface{}{
				"matched":  record.Match.Matched,
				"reason":   record.Match.Reason,
				"score":    record.Match.Score,
				"id":       record.Match.Row.ID,
				"name":     record.Match.Row.Name,
				"pkg":      record.Match.Row.Pkg,
				"url":      record.Match.Row.URL,
				"version":  record.Match.Row.Version,
				"repo":     record.Match.Row.Repo,
				"category": record.Match.Row.Category,
			},
			"meta":      record.RawMeta,
			"dist_json": record.RawDist,
		}
		if err := encoder.Encode(payload); err != nil {
			return err
		}
	}
	return nil
}

func writeMissingMarkdown(path string, records []*PgxnRecord, total, matched int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var b strings.Builder
	b.WriteString("# PGXN Missing Extensions\n\n")
	b.WriteString(fmt.Sprintf("- Generated at: `%s`\n", time.Now().UTC().Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("- PGXN distributions crawled: `%d`\n", total))
	b.WriteString(fmt.Sprintf("- Already matched in pgext: `%d`\n", matched))
	b.WriteString(fmt.Sprintf("- Missing from pgext: `%d`\n\n", len(records)))

	b.WriteString("| Name | URL | Version | Release | Date | Provides | Homepage | Tags | Maintainers | Abstract |\n")
	b.WriteString("|:-----|:----|:--------|:--------|:-----|:---------|:---------|:-----|:------------|:---------|\n")
	for _, record := range records {
		b.WriteString(fmt.Sprintf("| `%s` | [%s](%s) | `%s` | `%s` | `%s` | `%s` | %s | `%s` | `%s` | %s |\n",
			record.Dist,
			record.Dist,
			record.DistURL,
			nullDash(record.version()),
			nullDash(record.releaseStatus()),
			nullDash(record.date()),
			nullDash(strings.Join(record.provides(), ", ")),
			markdownLinkOrDash(record.homepage()),
			nullDash(strings.Join(record.tags(), ", ")),
			nullDash(strings.Join(record.maintainers(), ", ")),
			escapeMarkdownTable(record.abstract()),
		))
	}

	_, err = file.WriteString(b.String())
	return err
}

func writeCatalogEnrichedCSV(path string, rows []CatalogRecord, pkgMatches map[string][]*PgxnRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"id", "name", "pkg", "lead_ext", "category", "state", "url", "license", "version", "repo", "lang", "lead", "contrib",
		"pgxn_matched", "pgxn_dists", "pgxn_urls", "pgxn_meta_urls", "pgxn_versions", "pgxn_release_statuses", "pgxn_dates",
		"pgxn_homepages", "pgxn_repository_webs", "pgxn_repository_urls", "pgxn_repository_types", "pgxn_bugtrackers",
		"pgxn_maintainers", "pgxn_tags", "pgxn_license_keys", "pgxn_provides", "pgxn_match_reasons",
		"pgxn_meta_json", "pgxn_dist_json",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	sort.Slice(rows, func(i, j int) bool { return rows[i].ID < rows[j].ID })
	for _, row := range rows {
		matches := pkgMatches[row.Pkg]
		payload := aggregatedMatchPayload(matches)
		record := []string{
			strconv.Itoa(row.ID),
			row.Name,
			row.Pkg,
			row.LeadExt,
			row.Category,
			row.State,
			row.URL,
			row.License,
			row.Version,
			row.Repo,
			row.Lang,
			strconv.FormatBool(row.Lead),
			strconv.FormatBool(row.Contrib),
			strconv.FormatBool(len(matches) > 0),
			strings.Join(payload["dists"], " | "),
			strings.Join(payload["urls"], " | "),
			strings.Join(payload["meta_urls"], " | "),
			strings.Join(payload["versions"], " | "),
			strings.Join(payload["release_statuses"], " | "),
			strings.Join(payload["dates"], " | "),
			strings.Join(payload["homepages"], " | "),
			strings.Join(payload["repository_webs"], " | "),
			strings.Join(payload["repository_urls"], " | "),
			strings.Join(payload["repository_types"], " | "),
			strings.Join(payload["bugtrackers"], " | "),
			strings.Join(payload["maintainers"], " | "),
			strings.Join(payload["tags"], " | "),
			strings.Join(payload["license_keys"], " | "),
			strings.Join(payload["provides"], " | "),
			strings.Join(payload["match_reasons"], " | "),
			jsonCompact(collectRawMeta(matches)),
			jsonCompact(collectRawDist(matches)),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return writer.Error()
}

func writeSummaryJSON(path string, result *PgxnExportResult) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func writePGXNReadme(path string, result *PgxnExportResult, opts PgxnExportOptions) error {
	content := fmt.Sprintf(
		"# PGXN Export\n\n"+
			"This directory contains a standalone PGXN crawl generated by `pgext pgxn`.\n\n"+
			"- Generated at: `%s`\n"+
			"- Catalog source: `%s`\n"+
			"- PGXN index: `%s`\n"+
			"- PGXN API base: `%s`\n"+
			"- PGXN site base: `%s`\n"+
			"- Distributions crawled: `%d`\n"+
			"- Matched in pgext: `%d`\n"+
			"- Missing from pgext: `%d`\n"+
			"- Failed fetches: `%d`\n\n"+
			"Files:\n\n"+
			"- `pgxn_dist.csv`: flattened full PGXN table with raw JSON columns.\n"+
			"- `pgxn_dist.jsonl`: one JSON object per PGXN distribution.\n"+
			"- `pgxn_missing.csv`: subset of distributions not currently matched in pgext.\n"+
			"- `pgxn_missing.md`: human-readable missing-extension table.\n"+
			"- `pgext_enriched.csv`: current pgext catalog rows enriched with PGXN matches when available.\n"+
			"- `summary.json`: machine-readable run summary.\n",
		result.GeneratedAt, opts.CatalogPath, opts.IndexURL, opts.APIBase, opts.SiteBase, result.DistCount, result.MatchedCount, result.MissingCount, result.FailedCount)
	return os.WriteFile(path, []byte(content), 0o644)
}

func aggregatedMatchPayload(matches []*PgxnRecord) map[string][]string {
	payload := map[string][]string{
		"dists":            []string{},
		"urls":             []string{},
		"meta_urls":        []string{},
		"versions":         []string{},
		"release_statuses": []string{},
		"dates":            []string{},
		"homepages":        []string{},
		"repository_webs":  []string{},
		"repository_urls":  []string{},
		"repository_types": []string{},
		"bugtrackers":      []string{},
		"maintainers":      []string{},
		"tags":             []string{},
		"license_keys":     []string{},
		"provides":         []string{},
		"match_reasons":    []string{},
	}
	seen := make(map[string]map[string]struct{}, len(payload))
	for key := range payload {
		seen[key] = make(map[string]struct{})
	}

	add := func(key, value string) {
		value = strings.TrimSpace(value)
		if value == "" {
			return
		}
		if _, ok := seen[key][value]; ok {
			return
		}
		seen[key][value] = struct{}{}
		payload[key] = append(payload[key], value)
	}

	for _, match := range matches {
		add("dists", match.Dist)
		add("urls", match.DistURL)
		add("meta_urls", match.MetaURL)
		add("versions", match.version())
		add("release_statuses", match.releaseStatus())
		add("dates", match.date())
		add("homepages", match.homepage())
		add("repository_webs", match.repositoryWeb())
		add("repository_urls", match.repositoryURL())
		add("repository_types", match.repositoryType())
		add("bugtrackers", match.bugtracker())
		add("match_reasons", match.Match.Reason)

		for _, value := range match.maintainers() {
			add("maintainers", value)
		}
		for _, value := range match.tags() {
			add("tags", value)
		}
		for _, value := range match.licenseKeys() {
			add("license_keys", value)
		}
		for _, value := range match.provides() {
			add("provides", value)
		}
	}

	for key := range payload {
		sort.Strings(payload[key])
	}
	return payload
}

func collectRawMeta(matches []*PgxnRecord) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(matches))
	for _, match := range matches {
		if len(match.RawMeta) == 0 {
			continue
		}
		result = append(result, match.RawMeta)
	}
	return result
}

func collectRawDist(matches []*PgxnRecord) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(matches))
	for _, match := range matches {
		if len(match.RawDist) == 0 {
			continue
		}
		result = append(result, match.RawDist)
	}
	return result
}

func (r *PgxnRecord) displayName() string {
	if value := getString(r.rawValue("name")); value != "" {
		return value
	}
	if value := getString(r.rawMetaValue("name")); value != "" {
		return value
	}
	return r.Dist
}

func (r *PgxnRecord) version() string       { return getString(r.rawValue("version")) }
func (r *PgxnRecord) releaseStatus() string { return getString(r.rawValue("release_status")) }
func (r *PgxnRecord) date() string          { return getString(r.rawValue("date")) }
func (r *PgxnRecord) user() string          { return getString(r.rawValue("user")) }
func (r *PgxnRecord) abstract() string {
	return firstNonEmpty(getString(r.rawMetaValue("abstract")), getString(r.rawValue("abstract")))
}
func (r *PgxnRecord) description() string {
	return firstNonEmpty(getString(r.rawMetaValue("description")), getString(r.rawValue("description")))
}
func (r *PgxnRecord) homepage() string { return nestedString(r.rawValue("resources"), "homepage") }
func (r *PgxnRecord) repositoryWeb() string {
	return nestedString(r.rawValue("resources"), "repository", "web")
}
func (r *PgxnRecord) repositoryURL() string {
	return nestedString(r.rawValue("resources"), "repository", "url")
}
func (r *PgxnRecord) repositoryType() string {
	return nestedString(r.rawValue("resources"), "repository", "type")
}
func (r *PgxnRecord) bugtracker() string {
	return nestedString(r.rawValue("resources"), "bugtracker", "web")
}

func (r *PgxnRecord) maintainers() []string {
	return toStringSlice(firstValue(r.rawMetaValue("maintainer"), r.rawValue("maintainer")))
}

func (r *PgxnRecord) tags() []string {
	return toStringSlice(firstValue(r.rawMetaValue("tags"), r.rawValue("tags")))
}

func (r *PgxnRecord) provides() []string {
	value := firstValue(r.rawMetaValue("provides"), r.rawValue("provides"))
	if provides, ok := value.(map[string]interface{}); ok {
		keys := make([]string, 0, len(provides))
		for key := range provides {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		return keys
	}
	return nil
}

func (r *PgxnRecord) licenseKeys() []string {
	value := firstValue(r.rawMetaValue("license"), r.rawValue("license"))
	switch typed := value.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(typed))
		for key := range typed {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		return keys
	case string:
		if typed == "" {
			return nil
		}
		return []string{typed}
	case []interface{}:
		return toStringSlice(typed)
	default:
		return nil
	}
}

func (r *PgxnRecord) rawValue(key string) interface{} {
	if r.RawDist == nil {
		return nil
	}
	return r.RawDist[key]
}

func (r *PgxnRecord) rawMetaValue(key string) interface{} {
	if r.RawMeta == nil {
		return nil
	}
	return r.RawMeta[key]
}

func (r *PgxnRecord) rawProvides() interface{} {
	return firstValue(r.rawMetaValue("provides"), r.rawValue("provides"))
}

func (r *PgxnRecord) resourceURLs() []string {
	return uniqueNonEmpty(
		r.homepage(),
		r.repositoryWeb(),
		r.repositoryURL(),
	)
}

func (r *PgxnRecord) dateValue() time.Time {
	date := r.date()
	if date == "" {
		return time.Time{}
	}
	parsed, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}
	}
	return parsed
}

func getCSVField(record []string, columns map[string]int, name string) string {
	idx, ok := columns[name]
	if !ok || idx >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[idx])
}

func parseCSVBool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "t", "true", "1", "y", "yes":
		return true
	default:
		return false
	}
}

func canonicalName(value string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(value)) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func normalizeURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return strings.TrimSuffix(strings.ToLower(raw), "/")
	}

	host := strings.ToLower(strings.TrimPrefix(parsed.Host, "www."))
	path := strings.TrimSuffix(parsed.EscapedPath(), "/")
	path = strings.TrimSuffix(path, ".git")
	path = strings.TrimSuffix(path, "/issues")

	if host == "" && path == "" {
		return ""
	}
	return host + strings.ToLower(path)
}

func getString(value interface{}) string {
	switch typed := value.(type) {
	case string:
		return strings.TrimSpace(typed)
	case fmt.Stringer:
		return strings.TrimSpace(typed.String())
	default:
		return ""
	}
}

func nestedString(value interface{}, path ...string) string {
	current := value
	for _, key := range path {
		next, ok := current.(map[string]interface{})
		if !ok {
			return ""
		}
		current, ok = next[key]
		if !ok {
			return ""
		}
	}
	return getString(current)
}

func toStringSlice(value interface{}) []string {
	switch typed := value.(type) {
	case []string:
		return uniqueNonEmpty(typed...)
	case []interface{}:
		result := make([]string, 0, len(typed))
		for _, item := range typed {
			if text := getString(item); text != "" {
				result = append(result, text)
			}
		}
		return uniqueNonEmpty(result...)
	case string:
		if typed == "" {
			return nil
		}
		return []string{typed}
	default:
		return nil
	}
}

func uniqueNonEmpty(values ...string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func firstValue(values ...interface{}) interface{} {
	for _, value := range values {
		if value == nil {
			continue
		}
		return value
	}
	return nil
}

func jsonCompact(value interface{}) string {
	if value == nil {
		return ""
	}
	data, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(data)
}

func markdownLinkOrDash(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "-"
	}
	return fmt.Sprintf("[%s](%s)", raw, raw)
}

func escapeMarkdownTable(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "-"
	}
	value = strings.ReplaceAll(value, "\n", " ")
	value = strings.ReplaceAll(value, "|", "\\|")
	return value
}

func nullDash(value string) string {
	if strings.TrimSpace(value) == "" {
		return "-"
	}
	return value
}
