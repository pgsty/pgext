/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"pgext/cli"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Global flags
var (
	// Database connection
	dbURL string

	// Logging
	logLevel string
	logPath  string
	debug    bool
	logFile  *os.File

	// Command-specific flags
	force    bool
	region   string
	workers  int
	retry    int
	keep     bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "pgext",
	Short: "PostgreSQL Extension Metadata Manager",
	Long:  `pgext - PostgreSQL Extension Metadata Manager using PostgreSQL`,
	Example: `
  pgext init                    # initialize pgext schema
  pgext load ext                # load extension catalog
  pgext fetch                   # fetch repo metadata
  pgext parse                   # parse repo data
  pgext recap                   # generate matrix
  pgext reload                  # fetch + parse + recap
  pgext status                  # show metadata status
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setupLogging()
	},
}

// initCmd initializes the pgext schema
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize pgext schema and load initial data",
	Long: `Initialize the pgext schema in the target PostgreSQL database and load initial data.

This command will:
1. Create the pgext schema
2. Load initial CSV data (pg, os, category, repository, extension) from embedded files

If the schema already exists, this command will warn and skip initialization.
Use --force to drop and recreate the schema.
Requires the semver extension to be available in the database.`,
	Example: `
  pgext init                              # use default database
  pgext init --force                      # force drop and recreate
  pgext init -d postgres:///              # use local PostgreSQL
  pgext init -d postgres://localhost/vonng
  PGURL=meta pgext init                   # use env variable
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.InitSchema(force); err != nil {
			return fmt.Errorf("failed to initialize schema: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// loadCmd loads CSV data into pgext tables
var loadCmd = &cobra.Command{
	Use:   "load <table> [source]",
	Short: "load CSV data into pgext table",
	Args:  cobra.RangeArgs(1, 2),
	Long: `Load CSV data into a pgext schema table using PostgreSQL COPY protocol.

The table parameter can be:
  - extension/ext/e     : Extension catalog
  - repository/repo/r   : Repository definitions
  - category/cat/c      : Extension categories
  - Any other pgext table name

The source parameter can be:
  - HTTP/HTTPS URL      : Download from web
  - File path           : Read from local file
  - "-"                 : Read from stdin
  - (empty)             : Use default URL for known tables`,
	Example: `
  pgext load extension          # load from default URL
  pgext load ext                # short alias
  pgext load repo data.csv      # load from local file
  pgext load category https://example.com/cat.csv
  cat data.csv | pgext load extension -
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		tableName := args[0]
		source := ""
		if len(args) > 1 {
			source = args[1]
		}

		var r cli.Region
		if region != "" {
			if region == "china" || region == "mirror" {
				r = cli.RegionChina
			} else {
				r = cli.RegionDefault
			}
		} else {
			r = cli.GetRegion()
		}

		if err := cli.LoadTable(tableName, source, r); err != nil {
			return fmt.Errorf("failed to load table: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// statusCmd shows metadata status
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show metadata status",
	Long:  `Display the current status of the pgext metadata catalog including update timestamps and statistics.`,
	Example: `
  pgext status          # show current status
  pgext status -d vonng # use specific database
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.ShowStatus(); err != nil {
			return fmt.Errorf("failed to show status: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// fetchCmd fetches repository metadata
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch repository metadata from upstream",
	Example: `
  pgext fetch                   # fetch only missing repos
  pgext fetch -f                # force re-download all
  pgext fetch -r china          # use China mirrors
  pgext fetch -p 12             # use 12 concurrent workers
  pgext fetch --retry 3         # retry 3 times on failure
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    force,
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})

		ctx := context.Background()
		if err := fetcher.FetchAll(ctx); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// parseCmd parses repository data
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse repo data and generate package tables",
	Long: `Parse repository metadata and populate package tables.

Processes:
  - YUM SQLite databases → pgext.dnf
  - APT Packages files → pgext.apt
  - Combined data → pgext.bin`,
	Example: `
  pgext parse           # parse all repos
  pgext parse -k        # keep existing data (skip truncate)
  pgext parse -p 16     # use 16 parallel workers
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := cli.ParseOptions{
			Keep:     keep,
			Parallel: workers,
		}
		if err := cli.ParseRepoData(opts); err != nil {
			return fmt.Errorf("failed to parse: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// recapCmd generates availability matrix
var recapCmd = &cobra.Command{
	Use:   "recap",
	Short: "generate availability matrix from packages",
	Long: `Recapitulate extension package availability from parsed package data.
Generates the pgext.pkg availability matrix showing which extensions are available
for which PostgreSQL versions and operating systems.`,
	Example: `
  pgext recap           # generate availability matrix
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.RecapPackages(); err != nil {
			return fmt.Errorf("failed to recap: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// reloadCmd performs a complete reload
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "complete reload: fetch + parse + recap",
	Long: `Perform a complete metadata refresh cycle:
1. Fetch repository metadata from upstream
2. Parse repository data into package tables
3. Recapitulate availability matrix

This is equivalent to running: pgext fetch && pgext parse && pgext recap`,
	Example: `
  pgext reload          # complete refresh
  pgext reload -f       # force re-download all
  pgext reload -p 16    # use 16 parallel workers
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch
		logrus.Info("Step 1/3: Fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    force,
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})

		ctx := context.Background()
		if err := fetcher.FetchAll(ctx); err != nil {
			return fmt.Errorf("fetch failed: %w", err)
		}

		// Parse
		logrus.Info("Step 2/3: Parsing repository data...")
		if err := cli.ParseRepoData(cli.ParseOptions{
			Keep:     false, // Always truncate for reload
			Parallel: workers,
		}); err != nil {
			return fmt.Errorf("parse failed: %w", err)
		}

		// Recap
		logrus.Info("Step 3/3: Generating availability matrix...")
		if err := cli.RecapPackages(); err != nil {
			return fmt.Errorf("recap failed: %w", err)
		}

		logrus.Info("Reload completed successfully")
		return nil
	},
	PostRun: closeDatabase,
}

// repoCmd shows repository summary
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "show repository summary",
	Long:  `Display a summary of all configured repositories and their status.`,
	Example: `
  pgext repo            # show repository summary
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.ShowRepositories(); err != nil {
			return fmt.Errorf("failed to show repositories: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// purgeCmd drops the pgext schema
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "drop pgext schema and all its objects",
	Long: `Completely remove the pgext schema and all its objects from the database.
WARNING: This action is irreversible and will delete all metadata.`,
	Example: `
  pgext purge           # drop pgext schema
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ask for confirmation
		logrus.Warn("This will permanently delete the pgext schema and all data.")
		fmt.Print("Type 'yes' to confirm: ")

		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "yes" {
			logrus.Info("Purge cancelled")
			return nil
		}

		if err := cli.PurgeSchema(); err != nil {
			return fmt.Errorf("failed to purge schema: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// Helper functions

func setupLogging() error {
	if debug {
		logLevel = "debug"
	}

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		lvl = logrus.InfoLevel
		logrus.Warnf("invalid log level %q, using INFO", logLevel)
	}
	logrus.SetLevel(lvl)

	if logPath != "" {
		if logFile != nil {
			logFile.Close()
		}

		f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file %s: %w", logPath, err)
		}
		logFile = f
		logrus.SetOutput(f)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			FullTimestamp:   true,
		})
	}

	logrus.Debugf("logger initialized at level %s", lvl.String())
	return nil
}

func initDatabase(cmd *cobra.Command, args []string) error {
	if err := setupLogging(); err != nil {
		return err
	}
	return cli.InitDB(dbURL)
}

func closeDatabase(cmd *cobra.Command, args []string) {
	cli.CloseDB()
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("command execution failed")
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&dbURL, "db", "d", "", "PostgreSQL connection (URL, database name, or key=value format, env: PGURL)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level: debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().StringVar(&logPath, "log-path", "", "log file path, terminal by default")

	// Command-specific flags
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "force drop and recreate schema")

	loadCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")

	fetchCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-download all repos")
	fetchCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	fetchCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of concurrent download workers")
	fetchCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts on failure")

	parseCmd.Flags().BoolVarP(&keep, "keep", "k", false, "keep existing data (skip truncate)")
	parseCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")

	reloadCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-download all repos")
	reloadCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	reloadCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of concurrent workers")
	reloadCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts on failure")

	// Add commands to root
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(recapCmd)
	rootCmd.AddCommand(reloadCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(purgeCmd)
}