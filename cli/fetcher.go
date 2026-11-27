/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// MaxFileSize defines the maximum allowed file size for downloads (500MB).
const MaxFileSize = 500 * 1024 * 1024

// RepoMetadata represents a package repository with cached metadata.
type RepoMetadata struct {
	ID               string
	Type             string // rpm or deb
	MetadataURL      string
	MirrorURL        string
	CachedETag       sql.NullString
	CachedTime       sql.NullTime
	CachedSize       sql.NullInt64
	CachedPrimary    *RPMMetadataCache
	CachedDataSHA256 sql.NullString
}

// FetchResult represents the result of fetching a repository.
type FetchResult struct {
	Repository *RepoMetadata
	Data       []byte
	ETag       string
	LastMod    time.Time
	Updated    bool
	Error      error
	Extra      *RPMMetadataCache
}

// FetchOptions contains options for fetch operation.
type FetchOptions struct {
	Force      bool   // Force re-download all repos
	Region     string // Region for mirror selection (default|china)
	Parallel   int    // Number of parallel workers
	Retry      int    // Number of retry attempts
	BestEffort bool   // Allow a partial result when at least one repository succeeds
}

// Fetcher handles repository metadata fetching.
type Fetcher struct {
	httpClient *http.Client
	region     Region
	force      bool
	maxWorkers int
	retries    int
	bestEffort bool
}

// NewFetcher creates a new repository fetcher.
func NewFetcher(opts FetchOptions) *Fetcher {
	if opts.Parallel <= 0 {
		opts.Parallel = 8
	}
	if opts.Retry < 0 {
		opts.Retry = 1
	}

	region := RegionDefault
	if opts.Region == "china" || opts.Region == "mirror" {
		region = RegionChina
	}

	return &Fetcher{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy:               http.ProxyFromEnvironment,
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		region:     region,
		force:      opts.Force,
		maxWorkers: opts.Parallel,
		retries:    opts.Retry,
		bestEffort: opts.BestEffort,
	}
}

// FetchAll fetches metadata for all repositories.
func (f *Fetcher) FetchAll(ctx context.Context) error {
	repos, err := f.loadRepositories(ctx)
	if err != nil {
		return fmt.Errorf("load repositories: %w", err)
	}
	if len(repos) == 0 {
		logrus.Info("No repositories to fetch")
		return nil
	}

	logrus.Infof("Fetching %d repositories (workers=%d, force=%v)", len(repos), f.maxWorkers, f.force)
	results := runParallelRepos(ctx, repos, f.maxWorkers,
		func(repo *RepoMetadata, err error) *FetchResult { return &FetchResult{Repository: repo, Error: err} },
		f.fetchWithRetry,
	)

	report := RunReport{Operation: "fetch", Total: len(repos)}
	for _, result := range results {
		if result.Error != nil {
			report.AddFailure(result.Repository.ID, result.Error)
			continue
		}
		if result.Updated {
			if err := f.saveMetadata(ctx, result); err != nil {
				report.AddFailure(result.Repository.ID, fmt.Errorf("save metadata: %w", err))
			} else {
				report.Succeeded++
				logrus.Debugf("Updated %s (%s)", result.Repository.ID, FormatSize(int64(len(result.Data))))
			}
		} else {
			report.Skipped++
			logrus.Debugf("Skipped %s (not modified)", result.Repository.ID)
		}
	}

	logrus.Infof("Fetch complete: %d updated, %d skipped, %d failed",
		report.Succeeded, report.Skipped, report.Failed())
	for _, failure := range report.Failures {
		logrus.Errorf("%s: %v", failure.Item, failure.Err)
	}

	if err := report.Err(f.bestEffort); err != nil {
		return err
	}
	if report.Failed() > 0 {
		logrus.Warnf("best-effort fetch accepted %d failed repositories", report.Failed())
	}
	if report.Succeeded > 0 {
		if err := updateFetchTime(ctx); err != nil {
			logrus.Warnf("Failed to update fetch time: %v", err)
		}
	}
	return nil
}

// runParallelRepos runs fn across repos with bounded concurrency and an optional TTY progress bar.
func runParallelRepos[T any](
	ctx context.Context,
	repos []*RepoMetadata,
	workers int,
	makeErr func(*RepoMetadata, error) T,
	fn func(context.Context, *RepoMetadata) T,
) []T {
	results := make([]T, len(repos))
	sem := semaphore.NewWeighted(int64(workers))
	completed := make(chan struct{}, len(repos))
	done := make(chan struct{})
	go displayProgress(completed, len(repos), done)

	g, gctx := errgroup.WithContext(ctx)
	for i, repo := range repos {
		i, repo := i, repo
		g.Go(func() error {
			if err := sem.Acquire(gctx, 1); err != nil {
				results[i] = makeErr(repo, err)
				completed <- struct{}{}
				return nil
			}
			defer sem.Release(1)
			results[i] = fn(gctx, repo)
			completed <- struct{}{}
			return nil
		})
	}
	_ = g.Wait()
	close(completed)
	<-done
	return results
}

func isStdoutTTY() bool {
	fi, err := os.Stdout.Stat()
	return err == nil && (fi.Mode()&os.ModeCharDevice) != 0
}

func displayProgress(completed <-chan struct{}, total int, done chan<- struct{}) {
	defer close(done)
	if !isStdoutTTY() || total == 0 {
		for range completed {
		}
		return
	}
	const barWidth = 40
	count := 0
	fmt.Print("\033[?25l")
	for range completed {
		count++
		percent := float64(count) / float64(total) * 100
		filled := int(float64(count) / float64(total) * float64(barWidth))
		var bar strings.Builder
		bar.Grow(barWidth + 2)
		bar.WriteByte('[')
		for i := 0; i < barWidth; i++ {
			switch {
			case i < filled:
				bar.WriteByte('=')
			case i == filled:
				bar.WriteByte('>')
			default:
				bar.WriteByte(' ')
			}
		}
		bar.WriteByte(']')
		fmt.Printf("\r%s %d/%d (%.1f%%)", bar.String(), count, total, percent)
	}
	fmt.Print("\033[?25h\n")
}

func (f *Fetcher) loadRepositories(ctx context.Context) ([]*RepoMetadata, error) {
	urlColumn := "r.default_meta"
	fallbackColumn := "r.mirror_meta"
	if f.region == RegionChina {
		urlColumn = "COALESCE(r.mirror_meta, r.default_meta)"
		fallbackColumn = "r.default_meta"
	}

	query := fmt.Sprintf(`SELECT r.id,r.type,%s as url,COALESCE(%s, '') as mirror_url,d.etag,d.last_modified,d.size,d.extra,encode(sha256(d.data), 'hex')
		FROM pgext.repository r LEFT JOIN pgext.repo_data d ON r.id = d.id
		WHERE %s IS NOT NULL AND `+activeRepositoryPredicate+` ORDER BY r.id `, urlColumn, fallbackColumn, urlColumn)

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []*RepoMetadata
	for rows.Next() {
		repo := &RepoMetadata{}
		var extra []byte
		if err := rows.Scan(
			&repo.ID, &repo.Type, &repo.MetadataURL, &repo.MirrorURL,
			&repo.CachedETag, &repo.CachedTime, &repo.CachedSize,
			&extra, &repo.CachedDataSHA256,
		); err != nil {
			return nil, fmt.Errorf("scan repository metadata: %w", err)
		}
		if len(extra) > 0 {
			var cached RPMMetadataCache
			if json.Unmarshal(extra, &cached) == nil && cached.Format == rpmPrimaryFormat {
				repo.CachedPrimary = &cached
			}
		}
		repos = append(repos, repo)
	}
	return repos, rows.Err()
}

func (f *Fetcher) fetchWithRetry(ctx context.Context, repo *RepoMetadata) *FetchResult {
	var lastErr error
	for attempt := 0; attempt <= f.retries; attempt++ {
		if attempt > 0 {
			select {
			case <-time.After(time.Duration(attempt) * 2 * time.Second):
			case <-ctx.Done():
				return &FetchResult{Repository: repo, Error: ctx.Err()}
			}
		}
		result := f.fetchOne(ctx, repo)
		if result.Error == nil {
			return result
		}
		lastErr = result.Error
		logrus.Debugf("Fetch attempt %d/%d failed for %s: %v", attempt+1, f.retries+1, repo.ID, lastErr)
	}
	return &FetchResult{
		Repository: repo,
		Error:      fmt.Errorf("failed after %d attempts: %w", f.retries+1, lastErr),
	}
}

func (f *Fetcher) fetchOne(ctx context.Context, repo *RepoMetadata) *FetchResult {
	switch repo.Type {
	case "rpm":
		return f.fetchRPM(ctx, repo)
	case "deb":
		return f.fetchDEB(ctx, repo)
	default:
		return &FetchResult{
			Repository: repo,
			Error:      fmt.Errorf("unsupported repository type: %s", repo.Type),
		}
	}
}

func (f *Fetcher) fetchRPM(ctx context.Context, repo *RepoMetadata) *FetchResult {
	var errs []string
	for _, metadataURL := range repo.metadataURLs() {
		result := f.fetchRPMFromMetadataURL(ctx, repo, metadataURL)
		if result.Error == nil {
			return result
		}
		errs = append(errs, fmt.Sprintf("%s: %v", metadataURL, result.Error))
	}
	return &FetchResult{Repository: repo, Error: fmt.Errorf("%s", strings.Join(errs, "; "))}
}

func (f *Fetcher) fetchRPMFromMetadataURL(ctx context.Context, repo *RepoMetadata, metadataURL string) *FetchResult {
	repomd, err := f.fetchRepoMD(ctx, metadataURL)
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("fetch repomd.xml: %w", err)}
	}
	primary, err := repomd.FindPrimary()
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	primaryURL, err := rpmPrimaryURL(repomd.SourceURL, primary.Location.Href)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	if !f.force && repo.hasCurrentPrimary(primary, primaryURL) {
		return &FetchResult{Repository: repo, Updated: false}
	}
	resp, err := f.download(ctx, primaryURL, nil)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("primary XML HTTP %d without a reusable cache", resp.StatusCode)}
	}
	compressed, err := readMetadata(resp.Body, MaxFileSize)
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("read primary XML: %w", err)}
	}
	data, extra, err := decodeRPMPrimary(ctx, compressed, primary, primaryURL)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	lastMod, _ := http.ParseTime(resp.Header.Get("Last-Modified"))
	logrus.Debugf("Validated primary XML for %s: %d packages (%s)", repo.ID, extra.PackageCount, extra.Compression)
	return &FetchResult{
		Repository: repo, Data: data, Extra: extra,
		ETag: normalizeETag(resp.Header.Get("ETag")), LastMod: lastMod, Updated: true,
	}
}

func (r *RepoMetadata) metadataURLs() []string {
	urls := []string{r.MetadataURL}
	if r.MirrorURL != "" && r.MirrorURL != r.MetadataURL {
		urls = append(urls, r.MirrorURL)
	}
	return urls
}

// fetchDEB fetches DEB repository metadata using conditional HTTP GET (304 Not Modified).
func (f *Fetcher) fetchDEB(ctx context.Context, repo *RepoMetadata) *FetchResult {
	resp, err := f.download(ctx, repo.MetadataURL, repo)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return &FetchResult{Repository: repo, Updated: false}
	}

	data, err := io.ReadAll(io.LimitReader(resp.Body, MaxFileSize))
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("read response: %w", err)}
	}

	etag := normalizeETag(resp.Header.Get("ETag"))
	lastMod, _ := http.ParseTime(resp.Header.Get("Last-Modified"))
	return &FetchResult{Repository: repo, Data: data, ETag: etag, LastMod: lastMod, Updated: true}
}

func (f *Fetcher) download(ctx context.Context, url string, repo *RepoMetadata) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if !f.force && repo != nil {
		if repo.CachedETag.Valid && repo.CachedETag.String != "" {
			req.Header.Set("If-None-Match", repo.CachedETag.String)
		}
		if repo.CachedTime.Valid {
			req.Header.Set("If-Modified-Since", repo.CachedTime.Time.Format(http.TimeFormat))
		}
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return resp, nil
}

func (f *Fetcher) fetchRepoMD(ctx context.Context, url string) (*RepoMD, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	repomd, err := parseRPMManifest(resp.Body)
	if err != nil {
		return nil, err
	}
	repomd.SourceURL = resp.Request.URL.String()
	return repomd, nil
}

func (f *Fetcher) saveMetadata(ctx context.Context, result *FetchResult) error {
	var lastMod *time.Time
	if !result.LastMod.IsZero() {
		lastMod = &result.LastMod
	}
	return saveRepositoryMetadata(ctx, result.Repository.ID, result.Data, result.ETag, lastMod, result.Extra)
}

func saveRepositoryMetadata(ctx context.Context, id string, data []byte, etag string, lastMod *time.Time, extra *RPMMetadataCache) error {
	var descriptor any
	if extra != nil {
		raw, err := json.Marshal(extra)
		if err != nil {
			return fmt.Errorf("encode primary XML cache descriptor: %w", err)
		}
		descriptor = string(raw)
	}
	_, err := ExecSQLContext(ctx, `
        INSERT INTO pgext.repo_data (id, data, size, etag, last_modified, extra, update_at)
        VALUES ($1, $2, $3, $4, $5, $6::jsonb, CURRENT_TIMESTAMP)
        ON CONFLICT (id) DO UPDATE SET
            data = EXCLUDED.data,
            size = EXCLUDED.size,
            etag = EXCLUDED.etag,
            last_modified = EXCLUDED.last_modified,
            extra = CASE WHEN EXCLUDED.extra IS NULL THEN repo_data.extra
                         ELSE (COALESCE(repo_data.extra, '{}'::jsonb) - 'compressed_size' - 'open_size') || EXCLUDED.extra END,
            update_at = CURRENT_TIMESTAMP
    `, id, data, len(data), etag, lastMod, descriptor)
	return err
}

func updateFetchTime(ctx context.Context) error {
	_, err := ExecSQLContext(ctx, "UPDATE pgext.status SET fetch_time = CURRENT_TIMESTAMP")
	return err
}

func normalizeETag(etag string) string {
	etag = strings.TrimSpace(etag)
	if etag == "" {
		return ""
	}
	if strings.HasPrefix(etag, `"`) && strings.HasSuffix(etag, `"`) {
		return etag
	}
	if strings.HasPrefix(etag, `W/"`) && strings.HasSuffix(etag, `"`) {
		return etag
	}
	return `"` + etag + `"`
}
