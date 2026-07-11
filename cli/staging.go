/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"pgext/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

const (
	recapStatusStatement                    = "UPDATE pgext.status SET recap_time = now();"
	packagePipelineLockStatement            = "SELECT pg_advisory_lock(hashtext('pgext.package_catalog_pipeline'))"
	packagePipelineUnlockStatement          = "SELECT pg_advisory_unlock(hashtext('pgext.package_catalog_pipeline'))"
	packagePipelineTransactionLockStatement = "SELECT pg_advisory_xact_lock(hashtext('pgext.package_catalog_pipeline'))"
	packageCatalogLockStatement             = "SELECT pg_advisory_xact_lock(hashtext('pgext.package_catalog_publish'))"
)

var stageSequence atomic.Uint64

// packageStaging owns uniquely named, run-scoped tables used to build a new
// package catalog without exposing empty or partially populated live tables.
type packageStaging struct {
	apt     string
	dnf     string
	bin     string
	pkg     string
	hasCore bool
	hasPkg  bool
}

type packageSQLExecutor interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

func newPackageStaging(ctx context.Context, core, pkg bool) (*packageStaging, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	if !core && !pkg {
		return nil, fmt.Errorf("staging requires at least one table group")
	}

	stage := &packageStaging{hasCore: core, hasPkg: pkg}
	if core {
		stage.apt = nextStageTable("apt")
		stage.dnf = nextStageTable("dnf")
		stage.bin = nextStageTable("bin")
	}
	if pkg {
		stage.pkg = nextStageTable("pkg")
	}

	tables := []struct {
		stage string
		live  string
	}{
		{stage.apt, liveTable("apt")},
		{stage.dnf, liveTable("dnf")},
		{stage.bin, liveTable("bin")},
		{stage.pkg, liveTable("pkg")},
	}

	for _, table := range tables {
		if table.stage == "" {
			continue
		}
		query := fmt.Sprintf("CREATE UNLOGGED TABLE %s (LIKE %s INCLUDING ALL)", table.stage, table.live)
		if _, err := ExecSQLContext(ctx, query); err != nil {
			stage.Drop()
			return nil, fmt.Errorf("create staging table for %s: %w", table.live, err)
		}
	}

	return stage, nil
}

func nextStageTable(kind string) string {
	name := fmt.Sprintf("_stage_%s_%d_%d_%d", kind, os.Getpid(), time.Now().UnixNano(), stageSequence.Add(1))
	return pgx.Identifier{"pgext", name}.Sanitize()
}

func liveTable(name string) string {
	return pgx.Identifier{"pgext", name}.Sanitize()
}

// Drop removes every staging table owned by this run. It deliberately uses a
// fresh bounded context so cancellation of the pipeline does not leak tables.
func (s *packageStaging) Drop() {
	if s == nil || DB == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, table := range []string{s.pkg, s.bin, s.apt, s.dnf} {
		if table == "" {
			continue
		}
		if _, err := ExecSQLContext(ctx, "DROP TABLE IF EXISTS "+table); err != nil {
			logrus.Warnf("failed to drop staging table %s: %v", table, err)
		}
	}
}

// buildPkg generates a complete package matrix in the staging pkg table using
// the supplied (live or staged) binary package table.
func (s *packageStaging) buildPkg(ctx context.Context, binTable string) (int64, error) {
	if DB == nil {
		return 0, fmt.Errorf("database not initialized")
	}
	return s.buildPkgWith(ctx, DB, binTable)
}

func (s *packageStaging) buildPkgWith(ctx context.Context, executor packageSQLExecutor, binTable string) (int64, error) {
	if s == nil || !s.hasPkg || s.pkg == "" {
		return 0, fmt.Errorf("pkg staging table is not initialized")
	}
	if executor == nil {
		return 0, fmt.Errorf("package SQL executor is not initialized")
	}
	if binTable == "" {
		return 0, fmt.Errorf("binary package source table is empty")
	}
	var binCount int64
	if err := executor.QueryRow(ctx, "SELECT count(*) FROM "+binTable).Scan(&binCount); err != nil {
		return 0, fmt.Errorf("count binary package source rows: %w", err)
	}
	if binCount == 0 {
		return 0, fmt.Errorf("binary package source contains no rows; refusing to replace pgext.pkg")
	}
	for _, dimension := range []struct {
		name  string
		query string
	}{
		{name: "active PostgreSQL versions", query: "SELECT count(*) FROM pgext.active_pg"},
		{name: "active operating systems", query: "SELECT count(*) FROM pgext.active_os"},
	} {
		var count int64
		if err := executor.QueryRow(ctx, dimension.query).Scan(&count); err != nil {
			return 0, fmt.Errorf("count %s: %w", dimension.name, err)
		}
		if count == 0 {
			return 0, fmt.Errorf("no %s; refusing to build an empty package matrix", dimension.name)
		}
	}

	sqlBytes, err := db.ReadFile("reload.sql")
	if err != nil {
		return 0, fmt.Errorf("read reload.sql: %w", err)
	}
	sqlContent := string(sqlBytes)
	if !containsSQL(strings.TrimSpace(sqlContent)) {
		return 0, fmt.Errorf("reload.sql contains no executable SQL")
	}
	if !strings.Contains(sqlContent, recapStatusStatement) {
		return 0, fmt.Errorf("reload.sql recap timestamp statement changed; refusing non-atomic status update")
	}

	// Build only in run-scoped tables. The live status timestamp is updated in
	// the same transaction that publishes the new catalog.
	sqlContent, err = stagedReloadSQL(sqlContent, s.pkg, binTable)
	if err != nil {
		return 0, err
	}
	if _, err := executor.Exec(ctx, sqlContent); err != nil {
		return 0, fmt.Errorf("generate staged pkg table: %w", err)
	}

	var count int64
	if err := executor.QueryRow(ctx, "SELECT COUNT(*) FROM "+s.pkg).Scan(&count); err != nil {
		return 0, fmt.Errorf("count staged pkg records: %w", err)
	}
	if count == 0 {
		return 0, fmt.Errorf("staged package matrix is empty; refusing to replace pgext.pkg")
	}
	return count, nil
}

func stagedReloadSQL(sqlContent, pkgTable, binTable string) (string, error) {
	if !strings.Contains(sqlContent, "TRUNCATE pgext.pkg") ||
		!strings.Contains(sqlContent, "INSERT INTO pgext.pkg") ||
		!strings.Contains(sqlContent, "UPDATE pgext.pkg SET") {
		return "", fmt.Errorf("reload.sql package table statements changed")
	}
	if !strings.Contains(sqlContent, recapStatusStatement) {
		return "", fmt.Errorf("reload.sql recap timestamp statement changed")
	}

	// UPDATE statements must retain the logical `pkg` relation name because
	// correlated subqueries in reload.sql refer to pkg.pg/pkg.os/pkg.name.
	sqlContent = strings.ReplaceAll(sqlContent, "TRUNCATE pgext.pkg", "TRUNCATE "+pkgTable)
	sqlContent = strings.ReplaceAll(sqlContent, "INSERT INTO pgext.pkg", "INSERT INTO "+pkgTable)
	sqlContent = strings.ReplaceAll(sqlContent, "UPDATE pgext.pkg SET", "UPDATE "+pkgTable+" AS pkg SET")
	sqlContent = strings.ReplaceAll(sqlContent, "pgext.bin", binTable)
	sqlContent = strings.ReplaceAll(sqlContent, recapStatusStatement, "")
	return sqlContent, nil
}

// publish atomically replaces the requested live table groups. PostgreSQL's
// transactional TRUNCATE and AccessExclusive locks ensure concurrent readers
// see either the complete old catalog or the complete new catalog, never an
// empty or half-populated intermediate state.
func (s *packageStaging) publish(ctx context.Context, core, pkg bool) error {
	tx, err := beginPackageCatalogTransaction(ctx)
	if err != nil {
		return fmt.Errorf("begin atomic publish: %w", err)
	}
	defer rollbackPackageTransaction(tx)

	if err := lockPackageCatalog(ctx, tx); err != nil {
		return fmt.Errorf("lock package catalog publish: %w", err)
	}
	liveTables, err := s.publishLocked(ctx, tx, core, pkg)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit atomic publish: %w", err)
	}
	analyzePackageTables(ctx, liveTables)
	return nil
}

// rebuildAndPublishPkgFromLive holds the package catalog advisory lock while it
// reads live pgext.bin, builds the staged matrix, and replaces live pgext.pkg.
// This prevents an older standalone recap from publishing after a newer full
// parse has committed.
func (s *packageStaging) rebuildAndPublishPkgFromLive(ctx context.Context) (int64, error) {
	tx, err := beginPackageCatalogTransaction(ctx)
	if err != nil {
		return 0, fmt.Errorf("begin atomic recap: %w", err)
	}
	defer rollbackPackageTransaction(tx)

	if err := lockPackageCatalog(ctx, tx); err != nil {
		return 0, fmt.Errorf("lock package catalog recap: %w", err)
	}
	pkgCount, err := s.buildPkgWith(ctx, tx, liveTable("bin"))
	if err != nil {
		return 0, err
	}
	liveTables, err := s.publishLocked(ctx, tx, false, true)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("commit atomic recap: %w", err)
	}
	analyzePackageTables(ctx, liveTables)
	return pkgCount, nil
}

func lockPackageCatalog(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, packageCatalogLockStatement)
	return err
}

// lockPackagePipeline acquires the pipeline key transactionally for test gates
// and other short coordination transactions. Production parse runs use the
// dedicated session lock below so they do not hold an idle transaction.
func lockPackagePipeline(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, packagePipelineTransactionLockStatement)
	return err
}

type packagePipelineLock struct {
	conn *pgx.Conn
}

// acquirePackagePipelineLock serializes complete parse runs from their first
// repo_data read through publication. A dedicated connection prevents the lock
// from consuming a pool slot (and deadlocking a pool configured with one slot),
// while a session advisory lock avoids a long idle-in-transaction interval.
func acquirePackagePipelineLock(ctx context.Context) (*packagePipelineLock, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	conn, err := pgx.ConnectConfig(ctx, DB.Config().ConnConfig.Copy())
	if err != nil {
		return nil, fmt.Errorf("connect for package pipeline lock: %w", err)
	}
	if _, err := conn.Exec(ctx, packagePipelineLockStatement); err != nil {
		closePackagePipelineConnection(conn)
		return nil, fmt.Errorf("acquire package pipeline lock: %w", err)
	}
	return &packagePipelineLock{conn: conn}, nil
}

func (l *packagePipelineLock) Release() {
	if l == nil || l.conn == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := l.conn.Exec(ctx, packagePipelineUnlockStatement); err != nil {
		logrus.Warnf("failed to release package pipeline lock explicitly: %v", err)
	}
	if err := l.conn.Close(ctx); err != nil {
		logrus.Warnf("failed to close package pipeline lock connection: %v", err)
	}
	l.conn = nil
}

func closePackagePipelineConnection(conn *pgx.Conn) {
	if conn == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = conn.Close(ctx)
}

// beginPackageCatalogTransaction fixes the isolation contract used by the
// advisory-lock protocol. READ COMMITTED gives statements that run after a
// lock wait a fresh snapshot even when the database default is stricter.
func beginPackageCatalogTransaction(ctx context.Context) (pgx.Tx, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.BeginTx(ctx, packageCatalogTxOptions())
}

func packageCatalogTxOptions() pgx.TxOptions {
	return pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
}

func rollbackPackageTransaction(tx pgx.Tx) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = tx.Rollback(ctx)
}

// publishLocked replaces the requested live tables using a transaction that
// already holds the package catalog advisory lock.
func (s *packageStaging) publishLocked(ctx context.Context, tx pgx.Tx, core, pkg bool) ([]string, error) {
	if s == nil {
		return nil, fmt.Errorf("staging tables are not initialized")
	}
	if core && !s.hasCore {
		return nil, fmt.Errorf("core staging tables are not initialized")
	}
	if pkg && !s.hasPkg {
		return nil, fmt.Errorf("pkg staging table is not initialized")
	}

	var liveTables []string
	if pkg {
		liveTables = append(liveTables, liveTable("pkg"))
	}
	if core {
		liveTables = append(liveTables, liveTable("bin"), liveTable("apt"), liveTable("dnf"))
	}
	if len(liveTables) == 0 {
		return nil, fmt.Errorf("nothing selected for publish")
	}
	if _, err := tx.Exec(ctx, "TRUNCATE TABLE "+strings.Join(liveTables, ", ")); err != nil {
		return nil, fmt.Errorf("truncate live package tables: %w", err)
	}

	if core {
		for _, table := range []struct {
			live  string
			stage string
		}{
			{liveTable("apt"), s.apt},
			{liveTable("dnf"), s.dnf},
			{liveTable("bin"), s.bin},
		} {
			if _, err := tx.Exec(ctx, fmt.Sprintf("INSERT INTO %s SELECT * FROM %s", table.live, table.stage)); err != nil {
				return nil, fmt.Errorf("publish %s: %w", table.live, err)
			}
		}
	}

	if pkg {
		if _, err := tx.Exec(ctx, fmt.Sprintf("INSERT INTO %s SELECT * FROM %s", liveTable("pkg"), s.pkg)); err != nil {
			return nil, fmt.Errorf("publish %s: %w", liveTable("pkg"), err)
		}
	}

	// Take one wall-clock timestamp after the publish lock has been acquired and
	// all staged rows have been copied. GREATEST keeps the logical publication
	// clock monotonic even if the system clock moves backwards. A full publish
	// writes this same value to parse_time and recap_time.
	var publishedAt time.Time
	if err := tx.QueryRow(ctx, `
		SELECT GREATEST(
			clock_timestamp(),
			COALESCE(parse_time, '-infinity'::timestamptz) + interval '1 microsecond',
			COALESCE(recap_time, '-infinity'::timestamptz) + interval '1 microsecond'
		)
		FROM pgext.status
		WHERE id = 0
	`).Scan(&publishedAt); err != nil {
		return nil, fmt.Errorf("determine package catalog publish timestamp: %w", err)
	}
	if core {
		if _, err := tx.Exec(ctx, "UPDATE pgext.status SET parse_time = $1 WHERE id = 0", publishedAt); err != nil {
			return nil, fmt.Errorf("update parse timestamp: %w", err)
		}
	}
	if pkg {
		if _, err := tx.Exec(ctx, "UPDATE pgext.status SET recap_time = $1 WHERE id = 0", publishedAt); err != nil {
			return nil, fmt.Errorf("update recap timestamp: %w", err)
		}
	}
	return liveTables, nil
}

func analyzePackageTables(ctx context.Context, liveTables []string) {
	// Statistics are advisory and intentionally outside the publish transaction.
	for _, table := range liveTables {
		if _, err := ExecSQLContext(ctx, "ANALYZE "+table); err != nil {
			logrus.Warnf("failed to analyze %s after publish: %v", table, err)
		}
	}
}
