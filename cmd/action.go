/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"fmt"
	"pgext/cli"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	reloadBestEffort bool
	reloadKeepTemp   bool
	rescanBestEffort bool
	rescanKeepTemp   bool
	fetchBestEffort  bool
	scanBestEffort   bool
	parseBestEffort  bool
	parseKeepTemp    bool
	parseNoPkg       bool
)

// reloadCmd reloads all data: fetch + complete parse
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload data: fetch + parse",
	Long: `Reload all package data from repositories:
1. Fetch repository metadata (respecting cache)
2. Parse repository data into pgext.apt, pgext.dnf, and pgext.bin
3. Generate the pgext.pkg availability matrix

This runs the same logical stages as:
  pgext fetch
  pgext parse

The parsed package tables and availability matrix are published atomically.
`,
	Example: `
  pgext reload                  # reload all data
  pgext reload -p 8             # use 8 parallel workers
  pgext reload --force          # force re-fetch all repos
  pgext reload --best-effort    # publish a partial catalog if some repos fail
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("starting data reload...")

		// Step 1: Fetch repository metadata
		logrus.Info("step 1/3: fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:      force,
			Region:     region,
			Parallel:   workers,
			Retry:      retry,
			BestEffort: reloadBestEffort,
		})
		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		// Steps 2-3: Parse repository data and build the package matrix
		logrus.Info("steps 2-3/3: parsing repository data and generating package matrix...")
		parser := cli.NewParserContext(cmd.Context(), cli.ParseOptions{
			Parallel:   workers,
			KeepTemp:   reloadKeepTemp,
			BestEffort: reloadBestEffort,
		})
		if err := parser.ParseAndRecap(); err != nil {
			return fmt.Errorf("failed to build package catalog: %w", err)
		}

		logrus.Info("reload completed successfully!")
		return nil
	},
	PostRun: closeDatabase,
}

// rescanCmd reloads package data from local Pigsty repository metadata: scan + complete parse
var rescanCmd = &cobra.Command{
	Use:   "rescan",
	Short: "reload data from local repo: scan + parse",
	Long: `Reload package data from local Pigsty repository metadata:
1. Scan local Pigsty repository metadata
2. Parse repository data into pgext.apt, pgext.dnf, and pgext.bin
3. Generate the pgext.pkg availability matrix

This runs the same logical stages as:
  pgext scan
  pgext parse

The parsed package tables and availability matrix are published atomically.
`,
	Example: `
  pgext rescan                     # rescan from default ~/pgsty/repo
  pgext rescan --dir /path/to/repo # rescan from custom directory
  pgext rescan -p 8                # use 8 parallel workers
  pgext rescan -k                  # keep temporary DNF SQLite files
  pgext rescan --best-effort       # publish a partial catalog if some repos fail
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("starting data rescan...")

		// Step 1: Scan local repository metadata
		logrus.Info("step 1/3: scanning local repository metadata...")
		scanner, err := cli.NewScanner(cli.ScanOptions{
			RepoDir:    scanDir,
			Parallel:   workers,
			BestEffort: rescanBestEffort,
		})
		if err != nil {
			return fmt.Errorf("failed to create scanner: %w", err)
		}
		if err := scanner.ScanAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		// Steps 2-3: Parse repository data and build the package matrix
		logrus.Info("steps 2-3/3: parsing repository data and generating package matrix...")
		parser := cli.NewParserContext(cmd.Context(), cli.ParseOptions{
			Parallel:   workers,
			KeepTemp:   rescanKeepTemp,
			BestEffort: rescanBestEffort,
		})
		if err := parser.ParseAndRecap(); err != nil {
			return fmt.Errorf("failed to build package catalog: %w", err)
		}

		logrus.Info("rescan completed successfully!")
		return nil
	},
	PostRun: closeDatabase,
}

// fetchCmd fetches repository metadata
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch repository metadata from upstream",
	Long: `Fetch repository metadata from configured package repositories.

Downloads and caches repository metadata files (repomd.xml, Packages, etc.).
Uses HTTP conditional requests (ETag, Last-Modified) to avoid unnecessary downloads.
Use --force to ignore cache and re-download all repositories.`,
	Example: `
  pgext fetch                   # fetch with cache
  pgext fetch --force           # force re-fetch all
  pgext fetch -p 16             # use 16 parallel workers
  pgext fetch -r china          # use China mirrors
  pgext fetch --best-effort     # accept a partial fetch
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:      force,
			Region:     region,
			Parallel:   workers,
			Retry:      retry,
			BestEffort: fetchBestEffort,
		})

		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		return nil
	},
	PostRun: closeDatabase,
}

type parsePipeline interface {
	ParseAll() error
	ParseAndRecap() error
}

func runParsePipeline(parser parsePipeline, noPkg bool) error {
	if noPkg {
		if err := parser.ParseAll(); err != nil {
			return fmt.Errorf("failed to parse apt, dnf, and bin package data: %w", err)
		}
		return nil
	}
	if err := parser.ParseAndRecap(); err != nil {
		return fmt.Errorf("failed to build package catalog: %w", err)
	}
	return nil
}

// parseCmd parses repository data and builds a complete package catalog.
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse repository data and build the package catalog",
	Long: `Parse cached repository metadata and build a complete package catalog.

By default this atomically publishes:
- pgext.apt: Debian/Ubuntu packages
- pgext.dnf: RPM packages (EL, Fedora)
- pgext.bin: Unified binary package information
- pgext.pkg: Extension availability across PostgreSQL and operating systems

Use --no-pkg/-N only for debugging or compatibility. It stops after publishing
pgext.apt, pgext.dnf, and pgext.bin, leaving pgext.pkg stale until pgext recap
is run. Uses parallel workers for faster processing.`,
	Example: `
  pgext parse                   # build and atomically publish the complete catalog
  pgext parse -p 8              # use 8 parallel workers
  pgext parse -k                # keep temporary DNF SQLite files
  pgext parse --best-effort     # publish a partial catalog from successful repositories
  pgext parse --no-pkg          # stop after apt, dnf, and bin; pkg becomes stale
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		parser := cli.NewParserContext(cmd.Context(), cli.ParseOptions{
			Parallel:   workers,
			KeepTemp:   parseKeepTemp,
			BestEffort: parseBestEffort,
		})

		return runParsePipeline(parser, parseNoPkg)
	},
	PostRun: closeDatabase,
}

// recapCmd generates package availability matrix
var recapCmd = &cobra.Command{
	Use:   "recap",
	Short: "rebuild package availability matrix (pgext.pkg)",
	Long: `Rebuild the package availability matrix from current catalog data.

Reads pgext.bin together with extension, repository, active PostgreSQL, and
operating-system metadata. Atomically replaces pgext.pkg without fetching or
parsing repository metadata.`,
	Example: `
  pgext recap                   # rebuild availability matrix
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.RecapMatrixContext(cmd.Context()); err != nil {
			return fmt.Errorf("failed to recap: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// scanCmd scans local Pigsty repository metadata
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan local Pigsty repository metadata",
	Long: `Scan Pigsty repository metadata from local filesystem.

Reads repository metadata files from a local Pigsty repository directory
instead of downloading from the network. Only processes repositories with
org = 'pigsty'. Converts repository URLs like:
  https://repo.pigsty.io/apt/pgsql/bookworm/dists/bookworm/main/binary-arm64/Packages
to local paths:
  ~/pgsty/repo/apt/pgsql/bookworm/dists/bookworm/main/binary-arm64/Packages`,
	Example: `
  pgext scan                        # scan from default ~/pgsty/repo
  pgext scan --dir /path/to/repo    # scan from custom directory
  pgext scan -p 16                  # use 16 parallel workers
  pgext scan --best-effort          # accept a partial scan
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner, err := cli.NewScanner(cli.ScanOptions{
			RepoDir:    scanDir,
			Parallel:   workers,
			BestEffort: scanBestEffort,
		})
		if err != nil {
			return fmt.Errorf("failed to create scanner: %w", err)
		}

		if err := scanner.ScanAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		return nil
	},
	PostRun: closeDatabase,
}

func init() {
	// fetch command flags
	fetchCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-fetch all repositories")
	fetchCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	fetchCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	fetchCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts")
	fetchCmd.Flags().BoolVar(&fetchBestEffort, "best-effort", false, "allow partial success when some repositories fail")

	// parse command flags
	parseCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	parseCmd.Flags().BoolVarP(&parseKeepTemp, "keep-temp", "k", false, "keep temporary DNF SQLite files")
	parseCmd.Flags().BoolVar(&parseKeepTemp, "keep", false, "deprecated alias for --keep-temp")
	_ = parseCmd.Flags().MarkDeprecated("keep", "use --keep-temp")
	parseCmd.Flags().BoolVar(&parseBestEffort, "best-effort", false, "publish a partial catalog from repositories that parse successfully")
	parseCmd.Flags().BoolVarP(&parseNoPkg, "no-pkg", "N", false, "stop after apt, dnf, and bin; leaves pgext.pkg stale")

	// scan command flags
	scanCmd.Flags().StringVar(&scanDir, "dir", "~/pgsty/repo", "local repository directory")
	scanCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	scanCmd.Flags().BoolVar(&scanBestEffort, "best-effort", false, "allow partial success when some repositories fail")

	// rescan command flags
	rescanCmd.Flags().StringVar(&scanDir, "dir", "~/pgsty/repo", "local repository directory")
	rescanCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	rescanCmd.Flags().BoolVarP(&rescanKeepTemp, "keep-temp", "k", false, "keep temporary DNF SQLite files")
	rescanCmd.Flags().BoolVar(&rescanKeepTemp, "keep", false, "deprecated alias for --keep-temp")
	_ = rescanCmd.Flags().MarkDeprecated("keep", "use --keep-temp")
	rescanCmd.Flags().BoolVar(&rescanBestEffort, "best-effort", false, "publish a partial catalog when some repositories fail")
}

// loadCmd loads CSV data into tables
var loadCmd = &cobra.Command{
	Use:   "load <table>",
	Short: "load CSV data into specific table",
	Long: `Load embedded CSV data into a specific table.

Available tables:
- pg: PostgreSQL versions
- os: Operating systems
- category: Extension categories
- repository: Package repositories
- extension: Packaged extension catalog
- universe: Complete extension ecosystem (superset of extension)

The CSV files are embedded in the binary and loaded using PostgreSQL COPY.

WARNING: loading replaces the selected table with TRUNCATE ... CASCADE. Tables
that reference it may also be cleared; rebuild dependent catalog data after
loading dimension, repository, or extension tables.`,
	Example: `
  pgext load extension          # load extension catalog
  pgext load universe           # load complete extension universe
  pgext load repository         # load repository list
  pgext load pg                 # load PG versions
`,
	Args:    cobra.ExactArgs(1),
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		tableName := args[0]

		if !cli.IsEmbeddedTable(tableName) {
			return fmt.Errorf("invalid table name: %s (valid: %s)", tableName, strings.Join(cli.EmbeddedTableNames(), ", "))
		}

		if err := cli.LoadTable(tableName, "", cli.RegionDefault); err != nil {
			return fmt.Errorf("failed to load %s: %w", tableName, err)
		}

		return nil
	},
	PostRun: closeDatabase,
}
