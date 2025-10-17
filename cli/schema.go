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

// InitSchema initializes the pgext schema and loads initial data.
// If force is true, it will drop and recreate the schema.
func InitSchema(force bool) error {
	ctx := context.Background()

	// Check if schema already exists
	var exists bool
	err := QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'pgext')").Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check schema existence: %w", err)
	}

	if exists && !force {
		logrus.Warn("pgext schema already exists, use --force to recreate")
		return nil
	}

	// Drop existing schema if force flag is set
	if exists && force {
		logrus.Warn("dropping existing pgext schema...")
		if _, err := ExecContext(ctx, "DROP SCHEMA pgext CASCADE"); err != nil {
			return fmt.Errorf("failed to drop schema: %w", err)
		}
		logrus.Info("existing pgext schema dropped")
	}

	// Create pgext schema from embedded SQL
	schemaSQL, err := db.GetSchema()
	if err != nil {
		return fmt.Errorf("failed to get embedded schema: %w", err)
	}

	logrus.Info("creating pgext schema...")
	if _, err := ExecContext(ctx, schemaSQL); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	logrus.Info("pgext schema created")

	// Create repo_status tracking table
	_, err = ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS pgext.repo_status (
			id            TEXT PRIMARY KEY,
			downloaded    TIMESTAMPTZ,
			parsed        TIMESTAMPTZ,
			error         TEXT,
			file_size     BIGINT,
			package_count INTEGER
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create repo_status table: %w", err)
	}

	// Load initial CSV data (pg, os, category, repository, extension)
	if err := loadEmbeddedCSVData(ctx); err != nil {
		return fmt.Errorf("failed to load CSV data: %w", err)
	}

	// Initialize repository tracking
	_, err = ExecContext(ctx, `
		INSERT INTO pgext.repo_status (id)
		SELECT id FROM pgext.repository
		ON CONFLICT (id) DO NOTHING
	`)
	if err != nil {
		logrus.Warnf("failed to initialize repo_status: %v", err)
	}

	logrus.Info("pgext schema initialized successfully")
	return nil
}

// loadEmbeddedCSVData loads all embedded CSV files into the database.
func loadEmbeddedCSVData(ctx context.Context) error {
	csvFiles, err := db.GetCSVFiles()
	if err != nil {
		return fmt.Errorf("failed to get CSV files: %w", err)
	}

	logrus.Infof("loading %d CSV files...", len(csvFiles))

	for _, csvFile := range csvFiles {
		rowCount, err := loadCSVToTable(ctx, csvFile.Name, csvFile.Content)
		if err != nil {
			logrus.Errorf("failed to load %s: %v", csvFile.Name, err)
			continue // Continue with other files
		}
		logrus.Infof("  loaded %d rows into pgext.%s", rowCount, csvFile.Name)
	}

	return nil
}

// loadCSVToTable loads CSV data into a table using COPY protocol.
func loadCSVToTable(ctx context.Context, tableName string, csvData []byte) (int64, error) {
	if DB == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	conn, err := DB.Acquire(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// Build COPY query with sanitized table name
	copyQuery := fmt.Sprintf("COPY pgext.%s FROM STDIN WITH (FORMAT csv, HEADER true)",
		pgx.Identifier{tableName}.Sanitize())

	reader := bytes.NewReader(csvData)
	tag, err := conn.Conn().PgConn().CopyFrom(ctx, reader, copyQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data: %w", err)
	}

	return tag.RowsAffected(), nil
}

// PurgeSchema drops the pgext schema and all its objects.
func PurgeSchema() error {
	ctx := context.Background()

	logrus.Warn("dropping pgext schema...")
	_, err := ExecContext(ctx, "DROP SCHEMA IF EXISTS pgext CASCADE")
	if err != nil {
		return fmt.Errorf("failed to drop schema: %w", err)
	}

	logrus.Info("pgext schema dropped")
	return nil
}
