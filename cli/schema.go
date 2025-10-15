/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"bytes"
	"context"
	"fmt"
	"pgext/db"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// InitSchema initializes the pgext schema and loads initial data
func InitSchema(force bool) error {
	ctx := context.Background()

	// Check if schema exists
	var exists bool
	err := QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'pgext')").Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check schema existence: %w", err)
	}

	if exists && !force {
		logrus.Warn("pgext schema already exists, use --force to recreate")
		return nil
	}

	if exists && force {
		logrus.Warn("dropping existing pgext schema...")
		if _, err := ExecContext(ctx, "DROP SCHEMA pgext CASCADE"); err != nil {
			return fmt.Errorf("failed to drop schema: %w", err)
		}
		logrus.Info("existing pgext schema dropped")
	}

	// Check semver extension
	var semverExists bool
	err = QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM pg_extension
			WHERE extname = 'semver'
		)
	`).Scan(&semverExists)
	if err != nil {
		return fmt.Errorf("failed to check semver extension: %w", err)
	}

	if !semverExists {
		logrus.Warn("semver extension not found, attempting to create...")
		if _, err := ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS semver"); err != nil {
			logrus.Errorf("failed to create semver extension: %v", err)
			logrus.Warn("please install semver extension first: CREATE EXTENSION semver;")
			return fmt.Errorf("semver extension is required but not available")
		}
		logrus.Info("semver extension created")
	}

	// Get embedded schema SQL
	schemaSQL, err := db.GetSchema()
	if err != nil {
		return fmt.Errorf("failed to get embedded schema: %w", err)
	}

	logrus.Info("creating pgext schema...")
	if _, err := ExecContext(ctx, schemaSQL); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	logrus.Info("pgext schema created successfully")

	// Create pgext.repo_status table if not exists
	_, err = ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS pgext.repo_status (
			id          TEXT PRIMARY KEY,
			downloaded  TIMESTAMPTZ,
			parsed      TIMESTAMPTZ,
			error       TEXT,
			file_size   BIGINT,
			package_count INTEGER
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create repo_status table: %w", err)
	}

	// Load CSV data
	if err := loadEmbeddedCSVData(ctx); err != nil {
		return fmt.Errorf("failed to load CSV data: %w", err)
	}

	// Initialize repos in repo_status
	_, err = ExecContext(ctx, `
		INSERT INTO pgext.repo_status (id)
		SELECT id FROM pgext.repository
		ON CONFLICT (id) DO NOTHING
	`)
	if err != nil {
		logrus.Warnf("failed to initialize repo_status: %v", err)
	}

	logrus.Info("pgext schema initialized successfully with data")

	return nil
}

// loadEmbeddedCSVData loads all embedded CSV files into the database
func loadEmbeddedCSVData(ctx context.Context) error {
	// Get all embedded CSV files
	csvFiles, err := db.GetCSVFiles()
	if err != nil {
		return fmt.Errorf("failed to get CSV files: %w", err)
	}

	logrus.Infof("loading %d CSV files into database...", len(csvFiles))

	// Load each CSV file
	for _, csvFile := range csvFiles {
		rowCount, err := loadCSVToTable(ctx, csvFile.Name, csvFile.Content)
		if err != nil {
			logrus.Errorf("failed to load %s.csv: %v", csvFile.Name, err)
			// Continue with other files even if one fails
			continue
		}
		logrus.Infof("  loaded %d rows into pgext.%s", rowCount, csvFile.Name)
	}

	return nil
}

// loadCSVToTable loads CSV data into a specific table using COPY protocol
func loadCSVToTable(ctx context.Context, tableName string, csvData []byte) (int64, error) {
	if DB == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	// Acquire a connection from the pool
	conn, err := DB.Acquire(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// Use COPY protocol to load CSV data
	// Properly escape the table name to prevent SQL injection
	copyQuery := fmt.Sprintf("COPY pgext.%s FROM STDIN WITH (FORMAT csv, HEADER true)",
		pgx.Identifier{tableName}.Sanitize())

	// Start the COPY operation using CopyFrom
	reader := bytes.NewReader(csvData)
	tag, err := conn.Conn().PgConn().CopyFrom(ctx, reader, copyQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data to %s: %w", tableName, err)
	}

	return tag.RowsAffected(), nil
}

// PurgeSchema drops the pgext schema and all its objects
func PurgeSchema() error {
	ctx := context.Background()

	logrus.Warn("dropping pgext schema and all its objects...")
	_, err := ExecContext(ctx, "DROP SCHEMA IF EXISTS pgext CASCADE")
	if err != nil {
		return fmt.Errorf("failed to drop schema: %w", err)
	}

	logrus.Info("pgext schema dropped successfully")
	return nil
}