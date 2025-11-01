/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"bytes"
	"compress/bzip2"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// ScanOptions contains options for scan operation
type ScanOptions struct {
	RepoDir  string // Local repository directory (default: ~/pgsty/repo)
	Parallel int    // Number of parallel workers
}

// Scanner handles local repository metadata scanning
type Scanner struct {
	repoDir    string
	maxWorkers int
}

// ScanResult represents the result of scanning a local repository
type ScanResult struct {
	Repository *RepoMetadata
	Data       []byte
	Updated    bool
	Error      error
}

// NewScanner creates a new repository scanner
func NewScanner(opts ScanOptions) (*Scanner, error) {
	// Set defaults
	if opts.Parallel <= 0 {
		opts.Parallel = 8
	}

	// Expand ~ in repoDir
	repoDir := opts.RepoDir
	if repoDir == "" {
		repoDir = "~/pgsty/repo"
	}
	if strings.HasPrefix(repoDir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("get home directory: %w", err)
		}
		repoDir = filepath.Join(home, repoDir[2:])
	}

	// Verify directory exists
	info, err := os.Stat(repoDir)
	if err != nil {
		return nil, fmt.Errorf("access repo directory %s: %w", repoDir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("not a directory: %s", repoDir)
	}

	return &Scanner{
		repoDir:    repoDir,
		maxWorkers: opts.Parallel,
	}, nil
}

// ScanAll scans metadata for all Pigsty repositories from local directory
func (s *Scanner) ScanAll(ctx context.Context) error {
	// Load Pigsty repositories from database
	repos, err := s.loadPigstyRepositories(ctx)
	if err != nil {
		return fmt.Errorf("load repositories: %w", err)
	}

	if len(repos) == 0 {
		logrus.Info("No Pigsty repositories to scan")
		return nil
	}

	logrus.Infof("Scanning %d Pigsty repositories from %s (workers=%d)",
		len(repos), s.repoDir, s.maxWorkers)

	// Scan repositories concurrently
	results := s.scanConcurrent(ctx, repos)

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
			if err := s.saveMetadata(ctx, result); err != nil {
				failed++
				errors = append(errors, fmt.Sprintf("%s: save failed: %v", result.Repository.ID, err))
			} else {
				updated++
				logrus.Debugf("Updated %s (%s)", result.Repository.ID, formatBytes(int64(len(result.Data))))
			}
		} else {
			skipped++
			logrus.Debugf("Skipped %s (no change)", result.Repository.ID)
		}
	}

	// Update status timestamp
	if updated > 0 {
		if err := s.updateFetchTime(ctx); err != nil {
			logrus.Warnf("Failed to update fetch time: %v", err)
		}
	}

	// Log summary
	logrus.Infof("Scan complete: %d updated, %d skipped, %d failed", updated, skipped, failed)

	// Report errors
	for _, errMsg := range errors {
		logrus.Error(errMsg)
	}

	// Fail if all repositories failed
	if failed == len(repos) {
		return fmt.Errorf("all repositories failed to scan")
	}

	return nil
}

// loadPigstyRepositories loads Pigsty repository configurations from database
func (s *Scanner) loadPigstyRepositories(ctx context.Context) ([]*RepoMetadata, error) {
	query := `SELECT r.id, r.type, r.default_meta as url, d.etag, d.last_modified, d.size
		FROM pgext.repository r LEFT JOIN pgext.repo_data d ON r.id = d.id
		WHERE r.org = 'pigsty' AND r.default_meta IS NOT NULL
		ORDER BY r.id`

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

// scanConcurrent scans multiple repositories concurrently
func (s *Scanner) scanConcurrent(ctx context.Context, repos []*RepoMetadata) []*ScanResult {
	results := make([]*ScanResult, len(repos))
	total := len(repos)

	// Use semaphore to limit concurrency
	sem := semaphore.NewWeighted(int64(s.maxWorkers))

	// Progress tracking
	completed := make(chan int, total)
	done := make(chan bool)

	// Start progress display goroutine
	go s.displayProgress(completed, total, done)

	// Use errgroup for better error handling
	g, ctx := errgroup.WithContext(ctx)

	for i, repo := range repos {
		i, repo := i, repo // capture loop variables

		g.Go(func() error {
			// Acquire semaphore
			if err := sem.Acquire(ctx, 1); err != nil {
				results[i] = &ScanResult{Repository: repo, Error: err}
				completed <- 1
				return nil // Don't fail the whole group
			}
			defer sem.Release(1)

			// Scan repository
			result := s.scanOne(ctx, repo)
			results[i] = result

			// Signal completion
			completed <- 1

			return nil
		})
	}

	// Wait for all scans to complete
	_ = g.Wait() // We handle errors in results

	// Stop progress display
	close(completed)
	<-done

	return results
}

// displayProgress shows a progress bar for repository scanning
func (s *Scanner) displayProgress(completed chan int, total int, done chan bool) {
	count := 0
	barWidth := 40

	// Hide cursor
	fmt.Print("\033[?25l")

	for range completed {
		count++

		// Calculate progress
		percent := float64(count) / float64(total) * 100
		filled := int(float64(count) / float64(total) * float64(barWidth))

		// Build progress bar
		bar := "["
		for i := 0; i < barWidth; i++ {
			if i < filled {
				bar += "="
			} else if i == filled {
				bar += ">"
			} else {
				bar += " "
			}
		}
		bar += "]"

		// Print progress (carriage return to overwrite)
		fmt.Printf("\r%s %d/%d (%.1f%%)", bar, count, total, percent)
	}

	// Show cursor and newline
	fmt.Print("\033[?25h\n")
	done <- true
}

// scanOne scans a single repository from local filesystem
func (s *Scanner) scanOne(ctx context.Context, repo *RepoMetadata) *ScanResult {
	// Convert URL to local file path
	localPath, err := s.urlToLocalPath(repo.MetadataURL)
	if err != nil {
		return &ScanResult{Repository: repo, Error: err}
	}

	// Scan based on repository type
	switch repo.Type {
	case "rpm":
		return s.scanRPM(ctx, repo, localPath)
	case "deb":
		return s.scanDEB(ctx, repo, localPath)
	default:
		return &ScanResult{
			Repository: repo,
			Error:      fmt.Errorf("unsupported repository type: %s", repo.Type),
		}
	}
}

// urlToLocalPath converts a repository URL to local file path
func (s *Scanner) urlToLocalPath(url string) (string, error) {
	// Replace https://repo.pigsty.io with local repo directory
	const baseURL = "https://repo.pigsty.io"
	if !strings.HasPrefix(url, baseURL) {
		return "", fmt.Errorf("URL does not start with %s: %s", baseURL, url)
	}

	// Get relative path
	relPath := strings.TrimPrefix(url, baseURL)
	relPath = strings.TrimPrefix(relPath, "/")

	// Construct local path
	localPath := filepath.Join(s.repoDir, relPath)

	return localPath, nil
}

// scanRPM scans RPM repository metadata from local filesystem
func (s *Scanner) scanRPM(ctx context.Context, repo *RepoMetadata, repomdPath string) *ScanResult {
	// Check if file exists
	if _, err := os.Stat(repomdPath); err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("repomd.xml not found: %w", err)}
	}

	// Read and parse repomd.xml
	repomd, err := s.parseRepoMDFile(repomdPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("parse repomd.xml: %w", err)}
	}

	// Find primary_db entry
	primaryDB := repomd.FindPrimaryDB()
	if primaryDB == nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("primary_db not found in repomd.xml")}
	}

	// Construct path to primary.sqlite.bz2
	repoDir := filepath.Dir(repomdPath)
	primaryPath := filepath.Join(repoDir, "..", primaryDB.Location.Href)

	// Check if primary.sqlite.bz2 exists
	if _, err := os.Stat(primaryPath); err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("primary.sqlite.bz2 not found: %w", err)}
	}

	// Read compressed file
	compressed, err := os.ReadFile(primaryPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read primary.sqlite.bz2: %w", err)}
	}

	// Decompress bzip2
	decompressed, err := io.ReadAll(bzip2.NewReader(bytes.NewReader(compressed)))
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("decompress: %w", err)}
	}

	// Verify checksum if provided
	if primaryDB.OpenChecksum.Text != "" {
		if !s.verifyChecksum(decompressed, primaryDB.OpenChecksum.Text) {
			return &ScanResult{Repository: repo, Error: fmt.Errorf("checksum verification failed")}
		}
	}

	// Check if content has changed
	if s.hasChanged(repo, decompressed) {
		return &ScanResult{
			Repository: repo,
			Data:       decompressed,
			Updated:    true,
		}
	}

	return &ScanResult{Repository: repo, Updated: false}
}

// scanDEB scans DEB repository metadata from local filesystem
func (s *Scanner) scanDEB(ctx context.Context, repo *RepoMetadata, packagesPath string) *ScanResult {
	// Check if file exists
	if _, err := os.Stat(packagesPath); err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("Packages file not found: %w", err)}
	}

	// Read Packages file
	data, err := os.ReadFile(packagesPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read Packages file: %w", err)}
	}

	// Check if content has changed
	if s.hasChanged(repo, data) {
		return &ScanResult{
			Repository: repo,
			Data:       data,
			Updated:    true,
		}
	}

	return &ScanResult{Repository: repo, Updated: false}
}

// parseRepoMDFile reads and parses repomd.xml from local file
func (s *Scanner) parseRepoMDFile(path string) (*RepoMD, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var repomd RepoMD
	if err := xml.Unmarshal(data, &repomd); err != nil {
		return nil, fmt.Errorf("parse XML: %w", err)
	}

	return &repomd, nil
}

// verifyChecksum verifies SHA256 checksum
func (s *Scanner) verifyChecksum(data []byte, expected string) bool {
	hash := sha256.Sum256(data)
	actual := hex.EncodeToString(hash[:])
	return actual == expected
}

// hasChanged checks if repository data has changed compared to cached version
func (s *Scanner) hasChanged(repo *RepoMetadata, data []byte) bool {
	// If no cache, consider it changed
	if !repo.CachedSize.Valid {
		return true
	}

	// Check size first (quick check)
	if int64(len(data)) != repo.CachedSize.Int64 {
		return true
	}

	// If we have cached data in DB, we could compare checksums
	// For now, just consider size change as the indicator
	return false
}

// saveMetadata saves repository metadata to database
func (s *Scanner) saveMetadata(ctx context.Context, result *ScanResult) error {
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

	// Use file modification time as last_modified
	lastMod := time.Now()

	_, err := ExecSQLContext(ctx, query,
		result.Repository.ID,
		result.Data,
		len(result.Data),
		"", // No ETag for local files
		&lastMod,
	)

	return err
}

// updateFetchTime updates the fetch timestamp in status table
func (s *Scanner) updateFetchTime(ctx context.Context) error {
	_, err := ExecSQLContext(ctx, "UPDATE pgext.status SET fetch_time = CURRENT_TIMESTAMP")
	return err
}
