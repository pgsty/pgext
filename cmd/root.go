/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
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
	Short: "PostgreSQL Extension Metadata Manager",
	Long:  `pgext - PostgreSQL Extension Metadata Manager using PostgreSQL`,
	Example: `
  pgext init                    # setup everything (schema + reload)
  pgext schema [-f]             # initialize pgext schema
  pgext reload                  # reload data: fetch + parse + recap
  pgext fetch                   # get repo metadata from upstream
  pgext parse                   # populate apt, dnf, bin tables
  pgext recap                   # generate pkg table from bin info
  pgext status                  # show metadata status
  pgext purge                   # drop pgext schema
  pgext pkg <name>              # show package availability matrix
  pgext bin <name> -p 17 -o el9 # show binary packages with URLs
  pgext ext <name>              # show extension information
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setupLogging()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

// setupLogging configures the logging system
func setupLogging() error {
	if debug {
		logLevel = "debug"
	}

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if logPath != "" {
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		logFile = file
		logrus.SetOutput(file)
	}

	return nil
}

// initDatabase initializes the database connection
func initDatabase(cmd *cobra.Command, args []string) error {
	if dbURL == "" {
		dbURL = os.Getenv("PGURL")
	}
	if dbURL == "" {
		dbURL = "postgres:///"
	}

	if err := cli.InitDB(dbURL); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	logrus.Debugf("Connected to database: %s", cli.SanitizeURL(dbURL))
	return nil
}

// closeDatabase closes the database connection
func closeDatabase(cmd *cobra.Command, args []string) {
	cli.CloseDB()
	if logFile != nil {
		logFile.Close()
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&dbURL, "database", "d", "", "database connection URL or name")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "log level (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringVar(&logPath, "log-file", "", "log file path (default: stderr)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode (shorthand for --log-level=debug)")

	// Add all subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(schemaCmd)
	rootCmd.AddCommand(reloadCmd)

	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(recapCmd)

	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(pkgCmd)
	rootCmd.AddCommand(binCmd)
	rootCmd.AddCommand(extCmd)

	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(purgeCmd)
}
