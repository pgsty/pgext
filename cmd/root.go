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
	force   bool
	region  string
	workers int
	retry   int
	keep    bool

	// Query filter flags
	pgVer    int
	osFilter string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "pgext",
	Short: "PG Extension Catalog CLI",
	Long:  `pgext - PostgreSQL Extension Catalog CLI`,
	Example: `
  pgext init                    # setup everything (schema + reload)

  pgext schema [-f]             # initialize pgext schema
  pgext reload                  # reload data: fetch + parse + recap
  pgext status                  # show metadata status

  pgext fetch                   # get repo metadata from upstream
  pgext parse                   # populate apt, dnf, bin tables 
  pgext recap                   # generate pkg table from bin info 

  pgext pkg <name>              # show package availability matrix
  pgext bin <name> -p 17 -o el9 # show binary packages with URLs
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setupLogging()
	},
}

// schemaCmd initializes the pgext schema
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "initialize pgext schema",
	Long: `Initialize the pgext schema in the target PostgreSQL database.

This command will:
1. Create the pgext schema
2. Load initial CSV data (pg, os, category, repository, extension)

If the schema already exists, this command will skip initialization.
Use --force to drop and recreate the schema.
Requires the semver extension to be available in the database.`,
	Example: `
  pgext schema                  # initialize schema
  pgext schema -f               # force drop and recreate
  pgext schema -d vonng         # use specific database
  PGURL=meta pgext schema       # use env variable
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

// initCmd performs complete initialization: schema + reload
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "complete setup: schema + reload",
	Long: `Perform complete initialization from scratch:
1. Initialize pgext schema with force
2. Fetch repository metadata
3. Parse repository data into pgext.bin
4. Generate availability info pgext.pkg 

This is equivalent to: pgext schema && pgext reload`,
	Example: `
  pgext init                    # complete setup
  pgext init -p 16              # use 16 parallel workers
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize schema
		logrus.Info("Step 1/4: Initializing schema...")
		if err := cli.InitSchema(true); err != nil {
			return fmt.Errorf("schema initialization failed: %w", err)
		}

		// Fetch
		logrus.Info("Step 2/4: Fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    true, // Always fetch for init
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})

		ctx := context.Background()
		if err := fetcher.FetchAll(ctx); err != nil {
			return fmt.Errorf("fetch failed: %w", err)
		}

		// Parse
		logrus.Info("Step 3/4: Parsing repository data...")
		if err := cli.ParseRepoData(cli.ParseOptions{
			Keep:     false,
			Parallel: workers,
		}); err != nil {
			return fmt.Errorf("parse failed: %w", err)
		}

		// Recap
		logrus.Info("Step 4/4: Generating availability matrix...")
		if err := cli.RecapPackages(); err != nil {
			return fmt.Errorf("recap failed: %w", err)
		}

		logrus.Info("Initialization completed successfully")
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
	Long:  `Display the current status of the pgext metadata catalog including table counts and update timestamps.`,
	Example: `
  pgext status                  # show status
  pgext status -d vonng         # use specific database
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
	Short: "fetch repository metadata",
	Long:  `Fetch repository metadata from upstream package repositories.`,
	Example: `
  pgext fetch                   # fetch missing repos
  pgext fetch -f                # force re-download
  pgext fetch -r china          # use China mirrors
  pgext fetch -p 16             # use 16 workers
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
	Short: "parse repository data",
	Long: `Parse repository metadata into package tables.

Processes repository data into:
  - pgext.dnf  (YUM/DNF packages)
  - pgext.apt  (APT packages)
  - pgext.bin  (Combined binary packages)`,
	Example: `
  pgext parse                   # parse all repos
  pgext parse -k                # keep existing data
  pgext parse -p 16             # use 16 workers
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
	Short: "generate availability matrix",
	Long: `Generate the pgext.pkg availability matrix from parsed package data.

Shows which extensions are available for which PostgreSQL versions and operating systems.`,
	Example: `
  pgext recap                   # generate matrix
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
	Long: `Perform complete metadata refresh:
1. Fetch repository metadata
2. Parse repository data
3. Generate availability matrix

Equivalent to: pgext fetch && pgext parse && pgext recap`,
	Example: `
  pgext reload                  # complete refresh
  pgext reload -f               # force re-download
  pgext reload -p 16            # use 16 workers
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
	Long:  `Display summary of all configured repositories and their status.`,
	Example: `
  pgext repo                    # show repositories
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
	Short: "drop pgext schema",
	Long: `Drop the pgext schema and all its objects.

WARNING: This is irreversible and deletes all metadata.`,
	Example: `
  pgext purge                   # drop schema
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

// pkgCmd shows package availability matrix
var pkgCmd = &cobra.Command{
	Use:   "pkg <name>",
	Short: "show package availability matrix",
	Long:  `Display the availability matrix for a specific extension package across PostgreSQL versions and operating systems.`,
	Args:  cobra.ExactArgs(1),
	Example: `
  pgext pkg pgvector            # show pgvector availability
  pgext pkg postgis             # show postgis availability
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgName := args[0]
		if err := cli.ShowPackage(pkgName); err != nil {
			return fmt.Errorf("failed to show package: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// binCmd shows binary package information
var binCmd = &cobra.Command{
	Use:   "bin <name>",
	Short: "show binary package information",
	Long:  `Display binary packages from pgext.bin table with download URLs.`,
	Args:  cobra.ExactArgs(1),
	Example: `
  pgext bin pgvector            # show all pgvector packages
  pgext bin pgvector -p 17      # filter by PostgreSQL version
  pgext bin pgvector -o el9     # filter by OS
  pgext bin vector -p 17 -o u24 -r china
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := cli.ShowBin(name, pgVer, osFilter, region); err != nil {
			return fmt.Errorf("failed to show bin: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

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
	rootCmd.PersistentFlags().StringVarP(&dbURL, "db", "d", "", "database connection (URL/name, env: PGURL)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logging")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (debug|info|warn|error)")
	rootCmd.PersistentFlags().StringVar(&logPath, "log-path", "", "log file path (default: stderr)")

	// Command-specific flags
	schemaCmd.Flags().BoolVarP(&force, "force", "f", false, "force recreate schema")

	initCmd.Flags().StringVarP(&region, "region", "r", "", "region (china/mirror)")
	initCmd.Flags().IntVarP(&workers, "parallel", "p", 16, "concurrent workers")
	initCmd.Flags().IntVar(&retry, "retry", 1, "retry attempts")

	loadCmd.Flags().StringVarP(&region, "region", "r", "", "region (china/mirror)")

	fetchCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-download")
	fetchCmd.Flags().StringVarP(&region, "region", "r", "", "region (china/mirror)")
	fetchCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "concurrent workers")
	fetchCmd.Flags().IntVar(&retry, "retry", 1, "retry attempts")

	parseCmd.Flags().BoolVarP(&keep, "keep", "k", false, "keep existing data")
	parseCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "concurrent workers")

	reloadCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-download")
	reloadCmd.Flags().StringVarP(&region, "region", "r", "", "region (china/mirror)")
	reloadCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "concurrent workers")
	reloadCmd.Flags().IntVar(&retry, "retry", 1, "retry attempts")

	// Query filter flags
	binCmd.Flags().IntVarP(&pgVer, "pg", "p", 0, "PostgreSQL major version")
	binCmd.Flags().StringVarP(&osFilter, "os", "o", "", "OS filter (e.g., el9, u24)")
	binCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")

	// Add commands to root
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(reloadCmd)

	rootCmd.AddCommand(schemaCmd)
	rootCmd.AddCommand(purgeCmd)
	rootCmd.AddCommand(statusCmd)

	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(recapCmd)

	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(pkgCmd)
	rootCmd.AddCommand(binCmd)

}
