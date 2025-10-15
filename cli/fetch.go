/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"bytes"
	"compress/bzip2"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// RepoMetadata represents a package repository with cached metadata
type RepoMetadata struct {
	ID           string
	Type         string // rpm or deb
	MetadataURL  string
	CachedETag   sql.NullString
	CachedTime   sql.NullTime
	CachedSize   sql.NullInt64
}

// FetchResult represents the result of fetching a repository
type FetchResult struct {
	Repository *RepoMetadata
	Data       []byte
	ETag       string
	LastMod    time.Time
	Updated    bool
	Error      error
}

// FetchOptions contains options for fetch operation
type FetchOptions struct {
	Force    bool   // Force re-download all repos
	Region   string // Region for mirror selection (default|china)
	Parallel int    // Number of parallel workers
	Retry    int    // Number of retry attempts
}

// Fetcher handles repository metadata fetching
type Fetcher struct {
	httpClient *http.Client
	region     Region
	force      bool
	maxWorkers int
	retries    int
}

// NewFetcher creates a new repository fetcher
func NewFetcher(opts FetchOptions) *Fetcher {
	// Set defaults
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
			Timeout: 2 * time.Minute,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		region:     region,
		force:      opts.Force,
		maxWorkers: opts.Parallel,
		retries:    opts.Retry,
	}
}

// FetchAll fetches metadata for all repositories
func (f *Fetcher) FetchAll(ctx context.Context) error {
	// Load repositories from database
	repos, err := f.loadRepositories(ctx)
	if err != nil {
		return fmt.Errorf("load repositories: %w", err)
	}

	if len(repos) == 0 {
		logrus.Info("No repositories to fetch")
		return nil
	}

	logrus.Infof("Fetching %d repositories (workers=%d, force=%v)",
		len(repos), f.maxWorkers, f.force)

	// Fetch repositories concurrently
	results := f.fetchConcurrent(ctx, repos)

	// Process results
	var updated, skipped, failed int
	var errors []string

	for _, result := range results {
		if result.Error != nil {
			failed++
			errors = append(errors, fmt.Sprintf("%s: %v", result.Repository.ID, result.Error))
			continue
		}

		if result.Updated {
			if err := f.saveMetadata(ctx, result); err != nil {
				failed++
				errors = append(errors, fmt.Sprintf("%s: save failed: %v", result.Repository.ID, err))
			} else {
				updated++
				logrus.Debugf("Updated %s (%s)", result.Repository.ID, formatBytes(int64(len(result.Data))))
			}
		} else {
			skipped++
			logrus.Debugf("Skipped %s (not modified)", result.Repository.ID)
		}
	}

	// Update status timestamp
	if updated > 0 {
		if err := f.updateFetchTime(ctx); err != nil {
			logrus.Warnf("Failed to update fetch time: %v", err)
		}
	}

	// Log summary
	logrus.Infof("Fetch complete: %d updated, %d skipped, %d failed", updated, skipped, failed)

	// Report errors
	for _, errMsg := range errors {
		logrus.Error(errMsg)
	}

	// Fail if all repositories failed
	if failed == len(repos) {
		return fmt.Errorf("all repositories failed to fetch")
	}

	return nil
}

// loadRepositories loads repository configurations from database
func (f *Fetcher) loadRepositories(ctx context.Context) ([]*RepoMetadata, error) {
	urlColumn := "default_meta"
	if f.region == RegionChina {
		urlColumn = "COALESCE(mirror_meta, default_meta)"
	}

	query := fmt.Sprintf(`
		SELECT
			r.id,
			r.type,
			%s as url,
			d.etag,
			d.last_modified,
			d.size
		FROM pgext.repository r
		LEFT JOIN pgext.repo_data d ON r.id = d.id
		WHERE %s IS NOT NULL
		ORDER BY r.id
	`, urlColumn, urlColumn)

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []*RepoMetadata
	for rows.Next() {
		repo := &RepoMetadata{}
		err := rows.Scan(
			&repo.ID,
			&repo.Type,
			&repo.MetadataURL,
			&repo.CachedETag,
			&repo.CachedTime,
			&repo.CachedSize,
		)
		if err != nil {
			logrus.Warnf("Failed to scan repository: %v", err)
			continue
		}
		repos = append(repos, repo)
	}

	return repos, rows.Err()
}

// fetchConcurrent fetches multiple repositories concurrently
func (f *Fetcher) fetchConcurrent(ctx context.Context, repos []*RepoMetadata) []*FetchResult {
	results := make([]*FetchResult, len(repos))

	// Use semaphore to limit concurrency
	sem := semaphore.NewWeighted(int64(f.maxWorkers))

	// Use errgroup for better error handling
	g, ctx := errgroup.WithContext(ctx)

	for i, repo := range repos {
		i, repo := i, repo // capture loop variables

		g.Go(func() error {
			// Acquire semaphore
			if err := sem.Acquire(ctx, 1); err != nil {
				results[i] = &FetchResult{Repository: repo, Error: err}
				return nil // Don't fail the whole group
			}
			defer sem.Release(1)

			// Fetch with retries
			result := f.fetchWithRetry(ctx, repo)
			results[i] = result

			return nil
		})
	}

	// Wait for all fetches to complete
	_ = g.Wait() // We handle errors in results

	return results
}

// fetchWithRetry fetches a repository with retry logic
func (f *Fetcher) fetchWithRetry(ctx context.Context, repo *RepoMetadata) *FetchResult {
	var lastErr error

	for attempt := 0; attempt <= f.retries; attempt++ {
		// Add backoff for retries
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
		logrus.Debugf("Fetch attempt %d/%d failed for %s: %v",
			attempt+1, f.retries+1, repo.ID, lastErr)
	}

	return &FetchResult{
		Repository: repo,
		Error:      fmt.Errorf("failed after %d attempts: %w", f.retries+1, lastErr),
	}
}

// fetchOne fetches a single repository
func (f *Fetcher) fetchOne(ctx context.Context, repo *RepoMetadata) *FetchResult {
	// Check if we should skip (not forced and have cache)
	if !f.force && repo.CachedETag.Valid {
		// Check if modified using conditional request
		modified, err := f.checkModified(ctx, repo)
		if err != nil {
			return &FetchResult{Repository: repo, Error: err}
		}

		if !modified {
			return &FetchResult{Repository: repo, Updated: false}
		}
	}

	// Fetch based on repository type
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

// fetchRPM fetches RPM repository metadata
func (f *Fetcher) fetchRPM(ctx context.Context, repo *RepoMetadata) *FetchResult {
	// Parse repomd.xml to find primary.sqlite.bz2
	repomd, err := f.fetchRepoMD(ctx, repo.MetadataURL)
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("fetch repomd.xml: %w", err)}
	}

	// Find primary_db URL
	primaryDB := repomd.FindPrimaryDB()
	if primaryDB == nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("primary_db not found in repomd.xml")}
	}

	// Construct full URL for primary.sqlite.bz2
	baseURL := strings.TrimSuffix(repo.MetadataURL, "repodata/repomd.xml")
	primaryURL := baseURL + primaryDB.Location.Href

	// Download primary.sqlite.bz2
	resp, err := f.download(ctx, primaryURL, repo)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	defer resp.Body.Close()

	// Handle 304 Not Modified
	if resp.StatusCode == http.StatusNotModified {
		return &FetchResult{Repository: repo, Updated: false}
	}

	// Read and decompress
	compressed, err := io.ReadAll(io.LimitReader(resp.Body, MaxFileSize))
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("read response: %w", err)}
	}

	// Decompress bzip2
	decompressed, err := io.ReadAll(bzip2.NewReader(bytes.NewReader(compressed)))
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("decompress: %w", err)}
	}

	// Verify checksum if provided
	if primaryDB.OpenChecksum.Text != "" {
		if !f.verifyChecksum(decompressed, primaryDB.OpenChecksum.Text) {
			return &FetchResult{Repository: repo, Error: fmt.Errorf("checksum verification failed")}
		}
	}

	// Parse response headers
	etag := normalizeETag(resp.Header.Get("ETag"))
	lastMod, _ := http.ParseTime(resp.Header.Get("Last-Modified"))

	return &FetchResult{
		Repository: repo,
		Data:       decompressed,
		ETag:       etag,
		LastMod:    lastMod,
		Updated:    true,
	}
}

// fetchDEB fetches DEB repository metadata
func (f *Fetcher) fetchDEB(ctx context.Context, repo *RepoMetadata) *FetchResult {
	// Download Packages file directly
	resp, err := f.download(ctx, repo.MetadataURL, repo)
	if err != nil {
		return &FetchResult{Repository: repo, Error: err}
	}
	defer resp.Body.Close()

	// Handle 304 Not Modified
	if resp.StatusCode == http.StatusNotModified {
		return &FetchResult{Repository: repo, Updated: false}
	}

	// Read response
	data, err := io.ReadAll(io.LimitReader(resp.Body, MaxFileSize))
	if err != nil {
		return &FetchResult{Repository: repo, Error: fmt.Errorf("read response: %w", err)}
	}

	// Parse response headers
	etag := normalizeETag(resp.Header.Get("ETag"))
	lastMod, _ := http.ParseTime(resp.Header.Get("Last-Modified"))

	return &FetchResult{
		Repository: repo,
		Data:       data,
		ETag:       etag,
		LastMod:    lastMod,
		Updated:    true,
	}
}

// checkModified checks if remote content has been modified
func (f *Fetcher) checkModified(ctx context.Context, repo *RepoMetadata) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, "HEAD", repo.MetadataURL, nil)
	if err != nil {
		return false, err
	}

	// Add conditional headers
	if repo.CachedETag.Valid && repo.CachedETag.String != "" {
		req.Header.Set("If-None-Match", repo.CachedETag.String)
	}
	if repo.CachedTime.Valid {
		req.Header.Set("If-Modified-Since", repo.CachedTime.Time.Format(http.TimeFormat))
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return true, nil // Assume modified on error
	}
	defer resp.Body.Close()

	return resp.StatusCode != http.StatusNotModified, nil
}

// download performs HTTP GET with caching headers
func (f *Fetcher) download(ctx context.Context, url string, repo *RepoMetadata) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add conditional headers for caching
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

// fetchRepoMD fetches and parses repomd.xml
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

	var repomd RepoMD
	if err := xml.NewDecoder(resp.Body).Decode(&repomd); err != nil {
		return nil, fmt.Errorf("parse XML: %w", err)
	}

	return &repomd, nil
}

// verifyChecksum verifies SHA256 checksum
func (f *Fetcher) verifyChecksum(data []byte, expected string) bool {
	hash := sha256.Sum256(data)
	actual := hex.EncodeToString(hash[:])
	return actual == expected
}

// saveMetadata saves repository metadata to database
func (f *Fetcher) saveMetadata(ctx context.Context, result *FetchResult) error {
	query := `
		INSERT INTO pgext.repo_data (id, data, size, etag, last_modified, update_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
		ON CONFLICT (id) DO UPDATE SET
			data = EXCLUDED.data,
			size = EXCLUDED.size,
			etag = EXCLUDED.etag,
			last_modified = EXCLUDED.last_modified,
			update_at = CURRENT_TIMESTAMP
	`

	var lastMod *time.Time
	if !result.LastMod.IsZero() {
		lastMod = &result.LastMod
	}

	_, err := ExecSQLContext(ctx, query,
		result.Repository.ID,
		result.Data,
		len(result.Data),
		result.ETag,
		lastMod,
	)

	return err
}

// updateFetchTime updates the fetch timestamp in status table
func (f *Fetcher) updateFetchTime(ctx context.Context) error {
	_, err := ExecSQLContext(ctx, "UPDATE pgext.status SET fetch_time = CURRENT_TIMESTAMP")
	return err
}

// RepoMD represents YUM repository metadata structure
type RepoMD struct {
	XMLName xml.Name     `xml:"repomd"`
	Data    []RepoMDData `xml:"data"`
}

// RepoMDData represents a data entry in repomd.xml
type RepoMDData struct {
	Type     string `xml:"type,attr"`
	Location struct {
		Href string `xml:"href,attr"`
	} `xml:"location"`
	Checksum struct {
		Type string `xml:"type,attr"`
		Text string `xml:",chardata"`
	} `xml:"checksum"`
	OpenChecksum struct {
		Type string `xml:"type,attr"`
		Text string `xml:",chardata"`
	} `xml:"open-checksum"`
}

// FindPrimaryDB finds the primary_db entry in repomd
func (r *RepoMD) FindPrimaryDB() *RepoMDData {
	for i := range r.Data {
		if r.Data[i].Type == "primary_db" {
			return &r.Data[i]
		}
	}
	return nil
}



// normalizeETag normalizes an ETag value for consistent storage and comparison
// It keeps the quotes but ensures consistent format
func normalizeETag(etag string) string {
	if etag == "" {
		return ""
	}

	// Trim whitespace
	etag = strings.TrimSpace(etag)

	// Already properly quoted
	if strings.HasPrefix(etag, `"`) && strings.HasSuffix(etag, `"`) {
		return etag
	}

	// Handle weak ETags (W/"...")
	if strings.HasPrefix(etag, `W/"`) && strings.HasSuffix(etag, `"`) {
		return etag
	}

	// Add quotes if missing
	return `"` + etag + `"`
}

// compareETags compares two ETag values, handling different formats
func compareETags(etag1, etag2 string) bool {
	if etag1 == "" || etag2 == "" {
		return false
	}

	// Direct comparison first (most common case)
	if etag1 == etag2 {
		return true
	}

	// Strip W/ prefix for weak ETags
	e1 := strings.TrimPrefix(etag1, "W/")
	e2 := strings.TrimPrefix(etag2, "W/")

	// Compare after stripping weak prefix
	if e1 == e2 {
		return true
	}

	// Try comparing without quotes
	e1 = strings.Trim(e1, `"`)
	e2 = strings.Trim(e2, `"`)

	return e1 == e2
}