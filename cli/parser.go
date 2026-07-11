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
	KeepTemp   bool // Keep temporary DNF SQLite files for debugging
	Keep       bool // Deprecated: use KeepTemp. Retained for source compatibility.
	Parallel   int  // Number of parallel workers
	BestEffort bool // Allow a partial result when at least one repository succeeds
}

// Parser orchestrates the parsing of repository metadata
type Parser struct {
	ctx  context.Context
	opts ParseOptions
}

// NewParser creates a new parser instance
func NewParser(opts ParseOptions) *Parser {
	return NewParserContext(context.Background(), opts)
}

// NewParserContext creates a parser that honors caller cancellation.
func NewParserContext(ctx context.Context, opts ParseOptions) *Parser {
	if opts.Keep {
		opts.KeepTemp = true
	}
	// Apply defaults
	if opts.Parallel <= 0 {
		opts.Parallel = 8
	}

	return &Parser{
		ctx:  ctx,
		opts: opts,
	}
}

// ParseRepoData parses repository data and publishes a complete package catalog.
func ParseRepoData(opts ParseOptions) error {
	parser := NewParser(opts)
	return parser.ParseAndRecap()
}

// ParseAll parses and publishes only the core apt, dnf, and bin tables.
// Callers are responsible for rebuilding pgext.pkg afterward.
func (p *Parser) ParseAll() error {
	return p.withPipelineLock(p.parseAllLocked)
}

func (p *Parser) parseAllLocked() error {
	stage, _, err := p.build(false)
	if err != nil {
		return err
	}
	defer stage.Drop()

	logrus.Info("Publishing parsed package data atomically...")
	if err := stage.publish(p.ctx, true, false); err != nil {
		return fmt.Errorf("publish parsed package data: %w", err)
	}
	logrus.Warn("published apt, dnf, and bin without rebuilding pgext.pkg; run 'pgext recap' before serving the catalog")
	return nil
}

// ParseAndRecap builds apt, dnf, bin, and pkg in run-scoped staging tables and
// publishes all four live tables in one transaction.
func (p *Parser) ParseAndRecap() error {
	return p.withPipelineLock(p.parseAndRecapLocked)
}

func (p *Parser) parseAndRecapLocked() error {
	stage, _, err := p.build(true)
	if err != nil {
		return err
	}
	defer stage.Drop()

	logrus.Info("Generating staged package availability matrix...")
	pkgCount, err := stage.buildPkg(p.ctx, stage.bin)
	if err != nil {
		return fmt.Errorf("generate package matrix: %w", err)
	}
	logrus.Infof("Generated %d staged package matrix records", pkgCount)

	logrus.Info("Publishing package catalog atomically...")
	if err := stage.publish(p.ctx, true, true); err != nil {
		return fmt.Errorf("publish package catalog: %w", err)
	}
	return nil
}

// withPipelineLock prevents overlapping parse runs from publishing in reverse
// source-snapshot order. The lock is acquired before build reads repo_data and
// held until the selected core/full catalog has been published.
func (p *Parser) withPipelineLock(run func() error) error {
	lock, err := acquirePackagePipelineLock(p.ctx)
	if err != nil {
		return err
	}
	defer lock.Release()
	return run()
}

func (p *Parser) build(includePkg bool) (_ *packageStaging, _ ParseStats, retErr error) {
	startTime := time.Now()
	stage, err := newPackageStaging(p.ctx, true, includePkg)
	if err != nil {
		return nil, ParseStats{}, fmt.Errorf("create package staging tables: %w", err)
	}
	defer func() {
		if retErr != nil {
			stage.Drop()
		}
	}()

	repos, err := p.loadRepositories()
	if err != nil {
		return nil, ParseStats{}, fmt.Errorf("load repositories: %w", err)
	}
	if len(repos) == 0 {
		return nil, ParseStats{}, fmt.Errorf("no repositories with data to parse")
	}

	logrus.Infof("Starting staged parse operation (workers=%d, best-effort=%v)",
		p.opts.Parallel, p.opts.BestEffort)
	logrus.Infof("Found %d repositories to process", len(repos))

	results := p.parseRepositories(repos, stage)
	stats := p.processResults(results)
	p.showSummary(stats, time.Since(startTime))

	report := RunReport{Operation: "parse", Total: stats.TotalRepos, Succeeded: stats.TotalSuccess}
	for _, result := range results {
		if result.Error != nil {
			report.AddFailure(result.RepoID, result.Error)
		}
	}
	if err := report.Err(p.opts.BestEffort); err != nil {
		return nil, stats, err
	}
	if report.Failed() > 0 {
		logrus.Warnf("best-effort parse accepted %d failed repositories", report.Failed())
	}

	logrus.Info("Generating staged binary package table...")
	binGen := newBinGenerator(p.ctx, stage.dnf, stage.apt, stage.bin)
	binCount, err := binGen.Generate()
	if err != nil {
		return nil, stats, fmt.Errorf("generate binary packages: %w", err)
	}
	if binCount == 0 {
		return nil, stats, fmt.Errorf("generated no binary package records; refusing to replace the live catalog")
	}
	logrus.Infof("Generated %d staged binary package records", binCount)

	return stage, stats, nil
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
			return nil, fmt.Errorf("scan repository metadata: %w", err)
		}
		repos = append(repos, repo)
	}

	return repos, rows.Err()
}

// parseRepositories parses multiple repositories in parallel
func (p *Parser) parseRepositories(repos []RepositoryData, stage *packageStaging) []ParseResult {
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
			result := p.parseOneRepository(repo, stage)

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
func (p *Parser) parseOneRepository(repo RepositoryData, stage *packageStaging) ParseResult {
	result := ParseResult{
		RepoID: repo.ID,
		Type:   repo.Type,
	}

	var count int
	var err error

	// Create new parser instance for each repository to avoid concurrent access issues
	switch repo.Type {
	case "rpm":
		dnfParser := newDNFParser(p.ctx, stage.dnf, p.opts.KeepTemp)
		count, err = dnfParser.ParseRepository(repo.ID, repo.Data)
	case "deb":
		aptParser := newAPTParser(p.ctx, stage.apt)
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
