/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
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
		universeDDL, err := db.GetUniverseSchema()
		if err != nil {
			return fmt.Errorf("failed to get embedded universe schema: %w", err)
		}
		universeCSV, err := db.GetCSVFile("universe")
		if err != nil {
			return fmt.Errorf("failed to get embedded universe CSV: %w", err)
		}

		universeCreated, universeRows, err := ensureUniverse(ctx, universeDDL, universeCSV)
		if err != nil {
			return fmt.Errorf("failed to ensure pgext.universe: %w", err)
		}
		if !universeCreated {
			logrus.Warn("pgext schema already exists, use --force to recreate")
			return nil
		}

		logrus.Infof("added pgext.universe to existing schema with %d rows", universeRows)
		return nil
	}

	// Read assets before opening the transaction so asset failures cannot leave
	// a transaction (or, with --force, a dropped schema) behind.
	schemaSQL, err := db.GetSchema()
	if err != nil {
		return fmt.Errorf("failed to get embedded schema: %w", err)
	}
	csvFiles, err := db.GetCSVFiles()
	if err != nil {
		return fmt.Errorf("failed to get embedded CSV files: %w", err)
	}

	// DROP is transactional in PostgreSQL. Keeping it in the same transaction
	// as DDL and COPY restores the old schema if any initialization step fails.
	if exists && force {
		logrus.Warn("dropping existing pgext schema...")
	}

	logrus.Info("creating pgext schema...")
	results, err := initializeSchema(ctx, exists && force, schemaSQL, csvFiles)
	if err != nil {
		return err
	}

	if exists && force {
		logrus.Info("existing pgext schema replaced")
	} else {
		logrus.Info("pgext schema created")
	}
	for _, result := range results {
		logrus.Infof("  loaded %d rows into pgext.%s", result.rows, result.name)
	}
	logrus.Info("pgext schema initialized successfully")
	return nil
}

type csvLoadResult struct {
	name string
	rows int64
}

// ensureSchemaRelation atomically applies an idempotent schema addition under
// the pgext schema upgrade advisory lock.
func ensureSchemaRelation(ctx context.Context, relation, ddl string, populate func(pgx.Tx) (int64, error)) (bool, int64, error) {
	tx, err := Begin(ctx)
	if err != nil {
		return false, 0, fmt.Errorf("failed to begin %s transaction: %w", relation, err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, "SELECT pg_advisory_xact_lock(hashtext('pgext'), hashtext('schema'))"); err != nil {
		return false, 0, fmt.Errorf("failed to acquire schema lock: %w", err)
	}

	var exists bool
	if err := tx.QueryRow(ctx, "SELECT to_regclass($1) IS NOT NULL", relation).Scan(&exists); err != nil {
		return false, 0, fmt.Errorf("failed to check %s: %w", relation, err)
	}
	if exists {
		if err := tx.Commit(ctx); err != nil {
			return false, 0, fmt.Errorf("failed to commit %s check: %w", relation, err)
		}
		return false, 0, nil
	}

	if _, err := tx.Exec(ctx, ddl); err != nil {
		return false, 0, fmt.Errorf("failed to create %s: %w", relation, err)
	}
	rows, err := populate(tx)
	if err != nil {
		return false, 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return false, 0, fmt.Errorf("failed to commit %s transaction: %w", relation, err)
	}
	return true, rows, nil
}

// ensureUniverse upgrades schemas created before pgext.universe was part of schema.sql.
func ensureUniverse(ctx context.Context, universeDDL string, universeCSV db.CSVFile) (bool, int64, error) {
	return ensureSchemaRelation(ctx, "pgext.universe", universeDDL, func(tx pgx.Tx) (int64, error) {
		rows, err := copyCSVToTable(ctx, tx, universeCSV.Name, universeCSV.Content)
		if err != nil {
			return 0, fmt.Errorf("failed to load pgext.universe: %w", err)
		}
		return rows, nil
	})
}

func initializeSchema(ctx context.Context, replace bool, schemaSQL string, csvFiles []db.CSVFile) ([]csvLoadResult, error) {
	tx, err := Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin schema transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if replace {
		if _, err := tx.Exec(ctx, "DROP SCHEMA pgext CASCADE"); err != nil {
			return nil, fmt.Errorf("failed to drop schema: %w", err)
		}
	}
	if _, err := tx.Exec(ctx, schemaSQL); err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	results, err := loadEmbeddedCSVData(ctx, tx, csvFiles)
	if err != nil {
		return nil, fmt.Errorf("failed to load CSV data: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit schema transaction: %w", err)
	}
	return results, nil
}

// loadEmbeddedCSVData loads all embedded CSV files using the caller's transaction.
func loadEmbeddedCSVData(ctx context.Context, tx pgx.Tx, csvFiles []db.CSVFile) ([]csvLoadResult, error) {
	logrus.Infof("loading %d CSV files...", len(csvFiles))
	results := make([]csvLoadResult, 0, len(csvFiles))
	for _, csvFile := range csvFiles {
		rowCount, err := copyCSVToTable(ctx, tx, csvFile.Name, csvFile.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to load pgext.%s: %w", csvFile.Name, err)
		}
		results = append(results, csvLoadResult{name: csvFile.Name, rows: rowCount})
	}
	return results, nil
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
