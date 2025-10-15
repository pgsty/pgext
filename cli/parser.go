/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// ParseOptions contains options for parse operation
type ParseOptions struct {
	Keep     bool // Don't truncate existing data
	Parallel int  // Number of parallel workers
}

// Parser orchestrates the parsing of repository metadata
type Parser struct {
	ctx    context.Context
	opts   ParseOptions
	binGen *BinGenerator
}

// NewParser creates a new parser instance
func NewParser(opts ParseOptions) *Parser {
	// Apply defaults
	if opts.Parallel <= 0 {
		opts.Parallel = 8
	}

	ctx := context.Background()

	return &Parser{
		ctx:    ctx,
		opts:   opts,
		binGen: NewBinGenerator(ctx),
	}
}

// ParseRepoData is the main entry point for parsing repository data
func ParseRepoData(opts ParseOptions) error {
	parser := NewParser(opts)
	return parser.ParseAll()
}

// ParseAll orchestrates the complete parsing workflow
func (p *Parser) ParseAll() error {
	startTime := time.Now()

	// Step 1: Prepare - clear existing data if needed
	if !p.opts.Keep {
		logrus.Info("Clearing existing package data...")
		if err := p.clearExistingData(); err != nil {
			return fmt.Errorf("failed to clear data: %w", err)
		}
	}

	// Step 2: Load repositories that have data
	repos, err := p.loadRepositories()
	if err != nil {
		return fmt.Errorf("failed to load repositories: %w", err)
	}

	if len(repos) == 0 {
		logrus.Warn("No repositories with data to parse")
		return nil
	}

	logrus.Infof("Starting parse operation (workers=%d, keep=%v)", p.opts.Parallel, p.opts.Keep)
	logrus.Infof("Found %d repositories to process", len(repos))

	// Step 3: Parse repositories in parallel
	results := p.parseRepositories(repos)

	// Step 4: Process results and show summary
	stats := p.processResults(results)
	p.showSummary(stats, time.Since(startTime))

	// Step 5: Generate binary packages table
	if stats.TotalSuccess > 0 {
		logrus.Info("Generating binary package table...")
		binCount, err := p.binGen.Generate()
		if err != nil {
			logrus.Errorf("Failed to generate binary packages: %v", err)
		} else {
			logrus.Infof("Generated %d binary package records", binCount)
		}
	}

	// Step 6: Update parse timestamp
	if stats.TotalSuccess > 0 {
		if err := p.updateTimestamp(); err != nil {
			logrus.Warnf("Failed to update parse timestamp: %v", err)
		}
	}

	// Return error if all failed
	if stats.TotalFailed > 0 && stats.TotalSuccess == 0 {
		return fmt.Errorf("all repositories failed to parse")
	}

	return nil
}

// RepositoryData represents a repository with its metadata
type RepositoryData struct {
	ID   string
	Type string
	Data []byte
}

// ParseResult represents the result of parsing a repository
type ParseResult struct {
	RepoID string
	Type   string
	Count  int
	Error  error
}

// ParseStats holds parsing statistics
type ParseStats struct {
	TotalRepos   int
	TotalSuccess int
	TotalFailed  int
	DNFCount     int
	APTCount     int
	DNFPackages  int
	APTPackages  int
}

// clearExistingData truncates the package tables
func (p *Parser) clearExistingData() error {
	tables := []string{"pgext.dnf", "pgext.apt", "pgext.bin"}

	for _, table := range tables {
		if _, err := ExecSQLContext(p.ctx, fmt.Sprintf("TRUNCATE TABLE %s", table)); err != nil {
			logrus.Warnf("Failed to truncate %s: %v", table, err)
			// Continue even if one fails
		}
	}

	return nil
}

// loadRepositories loads all repositories that have data
func (p *Parser) loadRepositories() ([]RepositoryData, error) {
	query := `
		SELECT rd.id, r.type, rd.data
		FROM pgext.repo_data rd
		JOIN pgext.repository r ON rd.id = r.id
		WHERE rd.data IS NOT NULL
		ORDER BY rd.id
	`

	rows, err := QueryContext(p.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []RepositoryData
	for rows.Next() {
		var repo RepositoryData
		if err := rows.Scan(&repo.ID, &repo.Type, &repo.Data); err != nil {
			logrus.Warnf("Failed to scan repository: %v", err)
			continue
		}
		repos = append(repos, repo)
	}

	return repos, rows.Err()
}

// parseRepositories parses multiple repositories in parallel
func (p *Parser) parseRepositories(repos []RepositoryData) []ParseResult {
	results := make([]ParseResult, len(repos))
	var mu sync.Mutex // Protect results slice access

	// Use semaphore to limit concurrency
	sem := semaphore.NewWeighted(int64(p.opts.Parallel))

	// Use errgroup for better error handling
	g, ctx := errgroup.WithContext(p.ctx)

	for i, repo := range repos {
		i, repo := i, repo // Capture loop variables

		g.Go(func() error {
			// Acquire semaphore
			if err := sem.Acquire(ctx, 1); err != nil {
				mu.Lock()
				results[i] = ParseResult{
					RepoID: repo.ID,
					Type:   repo.Type,
					Error:  err,
				}
				mu.Unlock()
				return nil
			}
			defer sem.Release(1)

			// Parse repository
			result := p.parseOneRepository(repo)

			// Safely store result
			mu.Lock()
			results[i] = result
			mu.Unlock()

			// Log progress
			if result.Error != nil {
				logrus.Debugf("Failed to parse %s: %v", repo.ID, result.Error)
			} else {
				logrus.Debugf("Parsed %s: %d packages", repo.ID, result.Count)
			}

			return nil
		})
	}

	// Wait for all goroutines to complete
	_ = g.Wait()

	return results
}

// parseOneRepository parses a single repository
func (p *Parser) parseOneRepository(repo RepositoryData) ParseResult {
	result := ParseResult{
		RepoID: repo.ID,
		Type:   repo.Type,
	}

	var count int
	var err error

	// Create new parser instance for each repository to avoid concurrent access issues
	switch repo.Type {
	case "rpm":
		dnfParser := NewDNFParser(p.ctx)
		count, err = dnfParser.ParseRepository(repo.ID, repo.Data)
	case "deb":
		aptParser := NewAPTParser(p.ctx)
		count, err = aptParser.ParseRepository(repo.ID, repo.Data)
	default:
		err = fmt.Errorf("unsupported repository type: %s", repo.Type)
	}

	result.Count = count
	result.Error = err

	return result
}

// processResults processes parse results and returns statistics
func (p *Parser) processResults(results []ParseResult) ParseStats {
	stats := ParseStats{
		TotalRepos: len(results),
	}

	for _, result := range results {
		if result.Error != nil {
			stats.TotalFailed++
			logrus.Errorf("Repository %s failed: %v", result.RepoID, result.Error)
		} else {
			stats.TotalSuccess++

			if result.Type == "rpm" {
				stats.DNFCount++
				stats.DNFPackages += result.Count
			} else {
				stats.APTCount++
				stats.APTPackages += result.Count
			}
		}
	}

	return stats
}

// showSummary displays parsing summary
func (p *Parser) showSummary(stats ParseStats, duration time.Duration) {
	logrus.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logrus.Info("         Parse Operation Summary        ")
	logrus.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logrus.Infof("Duration:     %s", duration.Round(time.Millisecond))
	logrus.Infof("Repositories: %d total, %d succeeded, %d failed",
		stats.TotalRepos, stats.TotalSuccess, stats.TotalFailed)
	logrus.Infof("DNF/YUM:      %d repos, %d packages",
		stats.DNFCount, stats.DNFPackages)
	logrus.Infof("APT/DEB:      %d repos, %d packages",
		stats.APTCount, stats.APTPackages)
	logrus.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// updateTimestamp updates the parse_time in pgext.status
func (p *Parser) updateTimestamp() error {
	_, err := ExecSQLContext(p.ctx, `
		UPDATE pgext.status
		SET parse_time = CURRENT_TIMESTAMP
		WHERE id = 0
	`)

	if err == nil {
		logrus.Debug("Updated parse_time in pgext.status")
	}

	return err
}