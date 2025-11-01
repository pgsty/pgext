/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"fmt"
	"pgext/cli"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// reloadCmd reloads all data: fetch + parse + recap
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload data: fetch + parse + recap",
	Long: `Reload all package data from repositories:
1. Fetch repository metadata (respecting cache)
2. Parse repository data into pgext.bin
3. Generate availability info pgext.pkg

This is equivalent to running:
  pgext fetch
  pgext parse
  pgext recap
`,
	Example: `
  pgext reload                  # reload all data
  pgext reload -p 8             # use 8 parallel workers
  pgext reload --force          # force re-fetch all repos
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("starting data reload...")

		// Step 1: Fetch repository metadata
		logrus.Info("step 1/3: fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    force,
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})
		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		// Step 2: Parse repository data
		logrus.Info("step 2/3: parsing repository data...")
		parser := cli.NewParser(cli.ParseOptions{
			Parallel: workers,
			Keep:     keep,
		})
		if err := parser.ParseAll(); err != nil {
			return fmt.Errorf("failed to parse: %w", err)
		}

		// Step 3: Generate package matrix
		logrus.Info("step 3/3: generating package matrix...")
		if err := cli.RecapMatrix(); err != nil {
			return fmt.Errorf("failed to recap: %w", err)
		}

		logrus.Info("reload completed successfully!")
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
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    force,
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})

		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		return nil
	},
	PostRun: closeDatabase,
}

// parseCmd parses repository data
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse repository data into pgext tables",
	Long: `Parse repository metadata into structured package tables.

Processes cached repository metadata and populates:
- pgext.apt: Debian/Ubuntu packages
- pgext.dnf: RPM packages (EL, Fedora)
- pgext.bin: Unified binary package information

Uses parallel workers for faster processing.`,
	Example: `
  pgext parse                   # parse all repositories
  pgext parse -p 8              # use 8 parallel workers
  pgext parse -k                # keep temp SQLite files
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		parser := cli.NewParser(cli.ParseOptions{
			Parallel: workers,
			Keep:     keep,
		})

		if err := parser.ParseAll(); err != nil {
			return fmt.Errorf("failed to parse: %w", err)
		}

		return nil
	},
	PostRun: closeDatabase,
}

// recapCmd generates package availability matrix
var recapCmd = &cobra.Command{
	Use:   "recap",
	Short: "generate package availability matrix (pgext.pkg)",
	Long: `Generate package availability matrix from binary package data.

Creates the pgext.pkg table which shows extension availability across
PostgreSQL versions and operating systems. This is the final step in
the data processing pipeline.`,
	Example: `
  pgext recap                   # generate availability matrix
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.RecapMatrix(); err != nil {
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
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner, err := cli.NewScanner(cli.ScanOptions{
			RepoDir:  scanDir,
			Parallel: workers,
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

	// parse command flags
	parseCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	parseCmd.Flags().BoolVarP(&keep, "keep", "k", false, "keep temporary SQLite files")

	// scan command flags
	scanCmd.Flags().StringVar(&scanDir, "dir", "~/pgsty/repo", "local repository directory")
	scanCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
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
- extension: Extension catalog

The CSV files are embedded in the binary and loaded using PostgreSQL COPY.`,
	Example: `
  pgext load extension          # load extension catalog
  pgext load repository         # load repository list
  pgext load pg                 # load PG versions
`,
	Args:    cobra.ExactArgs(1),
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		tableName := args[0]

		// Validate table name
		validTables := map[string]bool{
			"pg":         true,
			"os":         true,
			"category":   true,
			"repository": true,
			"extension":  true,
		}

		if !validTables[tableName] {
			return fmt.Errorf("invalid table name: %s (valid: pg, os, category, repository, extension)", tableName)
		}

		if err := cli.LoadTable(tableName, "", cli.RegionDefault); err != nil {
			return fmt.Errorf("failed to load %s: %w", tableName, err)
		}

		return nil
	},
	PostRun: closeDatabase,
}
