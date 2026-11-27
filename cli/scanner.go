/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// ScanOptions contains options for scan operation.
type ScanOptions struct {
	RepoDir    string // Local repository directory (default: ~/pgsty/repo)
	Parallel   int    // Number of parallel workers
	BestEffort bool   // Allow a partial result when at least one repository succeeds
}

// Scanner handles local repository metadata scanning.
type Scanner struct {
	repoDir    string
	maxWorkers int
	bestEffort bool
}

// ScanResult represents the result of scanning a local repository.
type ScanResult struct {
	Repository *RepoMetadata
	Data       []byte
	Updated    bool
	Error      error
	Extra      *RPMMetadataCache
}

// NewScanner creates a new repository scanner.
func NewScanner(opts ScanOptions) (*Scanner, error) {
	if opts.Parallel <= 0 {
		opts.Parallel = 8
	}
	repoDir := opts.RepoDir
	if repoDir == "" {
		repoDir = "~/pgsty/repo"
	}
	repoDir, err := ExpandDir(repoDir)
	if err != nil {
		return nil, err
	}
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
		bestEffort: opts.BestEffort,
	}, nil
}

// ScanAll scans metadata for all Pigsty repositories from local directory.
func (s *Scanner) ScanAll(ctx context.Context) error {
	repos, err := s.loadPigstyRepositories(ctx)
	if err != nil {
		return fmt.Errorf("load repositories: %w", err)
	}
	if len(repos) == 0 {
		logrus.Info("No Pigsty repositories to scan")
		return nil
	}

	logrus.Infof("Scanning %d Pigsty repositories from %s (workers=%d)", len(repos), s.repoDir, s.maxWorkers)
	results := runParallelRepos(ctx, repos, s.maxWorkers,
		func(repo *RepoMetadata, err error) *ScanResult { return &ScanResult{Repository: repo, Error: err} },
		s.scanOne,
	)

	report := RunReport{Operation: "scan", Total: len(repos)}
	for _, result := range results {
		if result.Error != nil {
			report.AddFailure(result.Repository.ID, result.Error)
			continue
		}
		if result.Updated {
			if err := saveRepositoryMetadata(ctx, result.Repository.ID, result.Data, "", nil, result.Extra); err != nil {
				report.AddFailure(result.Repository.ID, fmt.Errorf("save metadata: %w", err))
			} else {
				report.Succeeded++
				logrus.Debugf("Updated %s (%s)", result.Repository.ID, FormatSize(int64(len(result.Data))))
			}
		} else {
			report.Skipped++
			logrus.Debugf("Skipped %s (no change)", result.Repository.ID)
		}
	}

	logrus.Infof("Scan complete: %d updated, %d skipped, %d failed",
		report.Succeeded, report.Skipped, report.Failed())
	for _, failure := range report.Failures {
		logrus.Errorf("%s: %v", failure.Item, failure.Err)
	}

	if err := report.Err(s.bestEffort); err != nil {
		return err
	}
	if report.Failed() > 0 {
		logrus.Warnf("best-effort scan accepted %d failed repositories", report.Failed())
	}
	if report.Succeeded > 0 {
		if err := updateFetchTime(ctx); err != nil {
			logrus.Warnf("Failed to update fetch time: %v", err)
		}
	}
	return nil
}

func (s *Scanner) loadPigstyRepositories(ctx context.Context) ([]*RepoMetadata, error) {
	query := `SELECT r.id, r.type, r.default_meta as url, d.etag, d.last_modified, d.size
		FROM pgext.repository r LEFT JOIN pgext.repo_data d ON r.id = d.id
		WHERE r.org = 'pigsty' AND r.default_meta IS NOT NULL AND ` + activeRepositoryPredicate + `
		ORDER BY r.id`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []*RepoMetadata
	for rows.Next() {
		repo := &RepoMetadata{}
		if err := rows.Scan(&repo.ID, &repo.Type, &repo.MetadataURL, &repo.CachedETag, &repo.CachedTime, &repo.CachedSize); err != nil {
			return nil, fmt.Errorf("scan repository metadata: %w", err)
		}
		repos = append(repos, repo)
	}
	return repos, rows.Err()
}

func (s *Scanner) scanOne(ctx context.Context, repo *RepoMetadata) *ScanResult {
	localPath, err := s.urlToLocalPath(repo.MetadataURL)
	if err != nil {
		return &ScanResult{Repository: repo, Error: err}
	}
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

func (s *Scanner) urlToLocalPath(url string) (string, error) {
	const baseURL = "https://repo.pigsty.io"
	if !strings.HasPrefix(url, baseURL) {
		return "", fmt.Errorf("URL does not start with %s: %s", baseURL, url)
	}
	relPath := strings.TrimPrefix(strings.TrimPrefix(url, baseURL), "/")
	return filepath.Join(s.repoDir, relPath), nil
}

func (s *Scanner) scanRPM(ctx context.Context, repo *RepoMetadata, repomdPath string) *ScanResult {
	file, err := os.Open(repomdPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read repomd.xml: %w", err)}
	}
	repomd, err := parseRPMManifest(file)
	file.Close()
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read repomd.xml: %w", err)}
	}
	primary, err := repomd.FindPrimary()
	if err != nil {
		return &ScanResult{Repository: repo, Error: err}
	}
	primaryPath, err := rpmPrimaryPath(repomdPath, primary.Location.Href)
	if err != nil {
		return &ScanResult{Repository: repo, Error: err}
	}
	pf, err := os.Open(primaryPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("open primary XML: %w", err)}
	}
	defer pf.Close()
	compressed, err := readMetadata(pf, MaxFileSize)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read primary XML: %w", err)}
	}
	data, extra, err := decodeRPMPrimary(ctx, compressed, primary, primaryPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: err}
	}
	return &ScanResult{Repository: repo, Data: data, Extra: extra, Updated: true}
}

func (s *Scanner) scanDEB(_ context.Context, repo *RepoMetadata, packagesPath string) *ScanResult {
	data, err := os.ReadFile(packagesPath)
	if err != nil {
		return &ScanResult{Repository: repo, Error: fmt.Errorf("read packages file: %w", err)}
	}
	return &ScanResult{Repository: repo, Data: data, Updated: true}
}
