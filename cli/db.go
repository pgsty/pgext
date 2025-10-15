/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	// DB is the global database connection pool
	DB *pgxpool.Pool
	// PGURL stores the connection string (for display purposes)
	PGURL string
)

// InitDB initializes PostgreSQL connection pool
func InitDB(pgurl string) error {
	ctx := context.Background()

	if pgurl == "" {
		// Check environment variable
		if envURL := os.Getenv("PGURL"); envURL != "" {
			pgurl = envURL
		} else {
			pgurl = "postgres:///" // Default to local socket
		}
	}

	// If pgurl is a simple database name (no postgres:// prefix), construct connection string
	if pgurl != "" && pgurl != "postgres:///" && !hasScheme(pgurl) {
		pgurl = "postgres:///" + pgurl
	}

	PGURL = pgurl

	config, err := pgxpool.ParseConfig(pgurl)
	if err != nil {
		return fmt.Errorf("failed to parse PostgreSQL URL: %w", err)
	}

	// Configure connection pool
	config.MaxConns = int32(getEnvInt("PGPOOL_MAX_CONNS", 20))
	config.MinConns = int32(getEnvInt("PGPOOL_MIN_CONNS", 4))
	config.MaxConnIdleTime = 30 * time.Minute
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection with timeout
	testCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(testCtx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = pool

	// Print connection summary
	dbSummary := getDBSummary(config.ConnConfig)
	logrus.Infof("connected to %s (pool: min=%d, max=%d)",
		dbSummary, config.MinConns, config.MaxConns)

	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		logrus.Debug("PostgreSQL connection pool closed")
	}
}

// ExecSQLContext executes SQL statement with context and returns affected rows
func ExecSQLContext(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	if DB == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	result, err := DB.Exec(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

// ExecContext is an alias for compatibility
func ExecContext(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if DB == nil {
		return pgconn.CommandTag{}, fmt.Errorf("database not initialized")
	}
	return DB.Exec(ctx, sql, args...)
}

// QueryRowContext executes a query with context that returns a single row
func QueryRowContext(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if DB == nil {
		// Return a row that will error on Scan
		return &emptyRow{err: fmt.Errorf("database not initialized")}
	}
	return DB.QueryRow(ctx, sql, args...)
}

// QueryContext executes a query with context that returns multiple rows
func QueryContext(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.Query(ctx, sql, args...)
}

// Begin starts a new transaction
func Begin(ctx context.Context) (pgx.Tx, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.Begin(ctx)
}

// hasScheme checks if a string has a URL scheme (postgres://, postgresql://)
func hasScheme(s string) bool {
	return strings.HasPrefix(s, "postgres://") || strings.HasPrefix(s, "postgresql://")
}

// getDBSummary returns a human-readable connection summary (database@host)
func getDBSummary(config *pgx.ConnConfig) string {
	db := config.Database
	if db == "" {
		db = "postgres"
	}

	host := config.Host
	if host == "" || host == "/var/run/postgresql" || host == "/tmp" || (len(host) > 0 && host[0] == '/') {
		host = "localhost"
	}

	// Include port if not default
	if config.Port != 5432 && config.Port != 0 {
		return fmt.Sprintf("%s@%s:%d", db, host, config.Port)
	}

	return fmt.Sprintf("%s@%s", db, host)
}

// sanitizeURL removes password from connection string for logging
func sanitizeURL(url string) string {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return "***"
	}
	if config.ConnConfig.Password != "" {
		config.ConnConfig.Password = "***"
	}
	return config.ConnString()
}

// getEnvInt gets an integer environment variable with default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// emptyRow is a placeholder row that returns an error on Scan
type emptyRow struct {
	err error
}

func (r *emptyRow) Scan(dest ...interface{}) error {
	return r.err
}

// SchemaExists checks if the pgext schema exists
func SchemaExists() (bool, error) {
	ctx := context.Background()
	var exists bool
	err := QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'pgext')").Scan(&exists)
	return exists, err
}