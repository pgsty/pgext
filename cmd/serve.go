/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pgext/cli"
	"pgext/server"

	"github.com/spf13/cobra"
)

var (
	serveDB       string
	serveListen   string
	serveCacheTTL time.Duration
)

// serveCmd runs the pgext.cloud web server
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the pgext.cloud web server",
	Long: `Serve the PGEXT.CLOUD web application and its query API from this binary.

The server connects to a PostgreSQL database that already has the pgext
schema loaded (pgext init && pgext reload), keeps an in-memory snapshot of
the catalog as its cache, and serves both the single-page app and the JSON
query API under /api/v1. Web assets are embedded — the binary is all you need.`,
	Example: `
  pgext serve                                    # local dev: postgres:///data (or $PGURL)
  pgext serve --db postgres://user:pw@host/data  # explicit connection string
  pgext serve --db mydb --listen :8080           # database name shorthand, custom port
  curl localhost:8432/api/v1/ext?q=vector        # query the API
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if serveDB != "" {
			dbURL = serveDB // --db wins over the global -d/--database flag
		}
		return initDatabase(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		if err := server.Serve(ctx, server.Options{
			Listen:   serveListen,
			CacheTTL: serveCacheTTL,
			Pool:     cli.DB,
		}); err != nil {
			return fmt.Errorf("server failed: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

func init() {
	serveCmd.Flags().StringVar(&serveDB, "db", "", "database connection URL or name (default: $PGURL or postgres:///data)")
	serveCmd.Flags().StringVar(&serveListen, "listen", ":8432", "listen address")
	serveCmd.Flags().DurationVar(&serveCacheTTL, "cache-ttl", 5*time.Minute, "catalog snapshot refresh interval")
	rootCmd.AddCommand(serveCmd)
}
