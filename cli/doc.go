package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const ensureDocTableSQL = `
CREATE TABLE IF NOT EXISTS pgext.doc (
    id          INTEGER PRIMARY KEY,
    ext         TEXT NOT NULL UNIQUE,
    pkg         TEXT NOT NULL,
    repo_url    TEXT,
    license_url TEXT,
    control_url TEXT,
    author_url  TEXT,
    home_url    TEXT,
    cargo_url   TEXT,
    en_doc      TEXT,
    zh_doc      TEXT
);
CREATE INDEX IF NOT EXISTS doc_ext_pkg_idx ON pgext.doc (ext, pkg);
`

// DocLoadStats summarizes a bilingual pgext.doc synchronization run.
type DocLoadStats struct {
	UniverseCount int
	LoadedCount   int
	ENFiles       int
	ZHFiles       int
	ExtraEN       []string
	ExtraZH       []string
	DryRun        bool
}

// SafeStubName rejects path-traversal or empty extension names before reading/writing stub files.
func SafeStubName(name string) error {
	if name == "" || name == "." || name == ".." ||
		strings.ContainsAny(name, "/\\\x00") || filepath.Base(name) != name {
		return fmt.Errorf("unsafe stub extension name %q", name)
	}
	return nil
}

// CollectStubDocs reads all top-level *.md files from dir keyed by extension stem.
func CollectStubDocs(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read stub directory %s: %w", dir, err)
	}
	docs := make(map[string]string, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		ext := strings.TrimSuffix(entry.Name(), ".md")
		if err := SafeStubName(ext); err != nil {
			return nil, err
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", filepath.Join(dir, entry.Name()), err)
		}
		docs[ext] = string(data)
	}
	return docs, nil
}

// DumpStubDocs exports non-null en_doc and zh_doc columns from pgext.doc into enDir and zhDir.
func DumpStubDocs(ctx context.Context, pool *pgxpool.Pool, enDir, zhDir string) (enCount, zhCount int, err error) {
	if err := os.MkdirAll(enDir, 0755); err != nil {
		return 0, 0, fmt.Errorf("create %s: %w", enDir, err)
	}
	if err := os.MkdirAll(zhDir, 0755); err != nil {
		return 0, 0, fmt.Errorf("create %s: %w", zhDir, err)
	}
	rows, err := pool.Query(ctx, `SELECT ext, en_doc, zh_doc FROM pgext.doc ORDER BY id, ext`)
	if err != nil {
		return 0, 0, fmt.Errorf("query pgext.doc: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ext string
		var enDoc, zhDoc sql.NullString
		if err := rows.Scan(&ext, &enDoc, &zhDoc); err != nil {
			return enCount, zhCount, fmt.Errorf("scan pgext.doc: %w", err)
		}
		if err := SafeStubName(ext); err != nil {
			return enCount, zhCount, err
		}
		if enDoc.Valid {
			if err := os.WriteFile(filepath.Join(enDir, ext+".md"), []byte(enDoc.String), 0644); err != nil {
				return enCount, zhCount, fmt.Errorf("write EN stub %s: %w", ext, err)
			}
			enCount++
		}
		if zhDoc.Valid {
			if err := os.WriteFile(filepath.Join(zhDir, ext+".md"), []byte(zhDoc.String), 0644); err != nil {
				return enCount, zhCount, fmt.Errorf("write ZH stub %s: %w", ext, err)
			}
			zhCount++
		}
	}
	return enCount, zhCount, rows.Err()
}

// LoadStubDocs atomically replaces pgext.doc from the complete bilingual Universe stub corpus.
// When dryRun is true, all checks and staged inserts run inside a transaction that rolls back.
func LoadStubDocs(ctx context.Context, pool *pgxpool.Pool, enDir, zhDir string, dryRun bool) (*DocLoadStats, error) {
	enDocs, err := CollectStubDocs(enDir)
	if err != nil {
		return nil, err
	}
	zhDocs, err := CollectStubDocs(zhDir)
	if err != nil {
		return nil, err
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin doc load transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, ensureDocTableSQL); err != nil {
		return nil, fmt.Errorf("ensure pgext.doc table: %w", err)
	}
	if _, err := tx.Exec(ctx, `LOCK TABLE pgext.universe IN SHARE MODE; LOCK TABLE pgext.doc IN SHARE ROW EXCLUSIVE MODE;`); err != nil {
		return nil, fmt.Errorf("lock catalog tables: %w", err)
	}

	universeNames, err := loadUniverseNames(ctx, tx)
	if err != nil {
		return nil, err
	}
	stats, stageRows, err := prepareStubStage(universeNames, enDocs, zhDocs, dryRun)
	if err != nil {
		return nil, err
	}

	if err := dropLegacyDocForeignKeys(ctx, tx); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(ctx, `
		CREATE TEMP TABLE stub_doc_load (
			ext    TEXT PRIMARY KEY,
			en_doc TEXT NOT NULL,
			zh_doc TEXT NOT NULL
		) ON COMMIT DROP`); err != nil {
		return nil, fmt.Errorf("create temp staging table: %w", err)
	}
	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"stub_doc_load"}, []string{"ext", "en_doc", "zh_doc"}, pgx.CopyFromRows(stageRows)); err != nil {
		return nil, fmt.Errorf("stage stub documents: %w", err)
	}

	if _, err := tx.Exec(ctx, `DELETE FROM pgext.doc`); err != nil {
		return nil, fmt.Errorf("clear pgext.doc: %w", err)
	}
	tag, err := tx.Exec(ctx, `
		INSERT INTO pgext.doc (
			id, ext, pkg, repo_url, license_url, control_url,
			author_url, home_url, cargo_url, en_doc, zh_doc
		)
		SELECT u.id, u.name, u.pkg, u.url, u.license_url, u.control_url,
		       u.author_url, u.home_url, u.cargo_url, s.en_doc, s.zh_doc
		FROM pgext.universe AS u
		JOIN stub_doc_load AS s ON s.ext = u.name
		ORDER BY u.id`)
	if err != nil {
		return nil, fmt.Errorf("populate pgext.doc: %w", err)
	}
	stats.LoadedCount = int(tag.RowsAffected())

	if err := verifyLoadedDocs(ctx, tx, stats.UniverseCount); err != nil {
		return nil, err
	}
	if dryRun {
		return stats, nil
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit pgext.doc load: %w", err)
	}
	return stats, nil
}

func loadUniverseNames(ctx context.Context, tx pgx.Tx) ([]string, error) {
	rows, err := tx.Query(ctx, `SELECT name FROM pgext.universe ORDER BY id, name`)
	if err != nil {
		return nil, fmt.Errorf("query pgext.universe names: %w", err)
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("scan pgext.universe name: %w", err)
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, fmt.Errorf("pgext.universe is empty; refusing to overwrite pgext.doc")
	}
	return names, nil
}

func prepareStubStage(universeNames []string, enDocs, zhDocs map[string]string, dryRun bool) (*DocLoadStats, [][]any, error) {
	uSet := make(map[string]struct{}, len(universeNames))
	var missingEN, missingZH []string
	stageRows := make([][]any, 0, len(universeNames))
	for _, name := range universeNames {
		uSet[name] = struct{}{}
		en, okEN := enDocs[name]
		if !okEN || strings.TrimSpace(en) == "" {
			missingEN = append(missingEN, name)
		}
		zh, okZH := zhDocs[name]
		if !okZH || strings.TrimSpace(zh) == "" {
			missingZH = append(missingZH, name)
		}
		if okEN && okZH && strings.TrimSpace(en) != "" && strings.TrimSpace(zh) != "" {
			stageRows = append(stageRows, []any{name, en, zh})
		}
	}
	if len(missingEN) > 0 || len(missingZH) > 0 {
		return nil, nil, fmt.Errorf("incomplete bilingual stub corpus: missing/empty EN=%d (%s), ZH=%d (%s)",
			len(missingEN), previewNames(missingEN), len(missingZH), previewNames(missingZH))
	}

	var extraEN, extraZH []string
	for name := range enDocs {
		if _, ok := uSet[name]; !ok {
			extraEN = append(extraEN, name)
		}
	}
	for name := range zhDocs {
		if _, ok := uSet[name]; !ok {
			extraZH = append(extraZH, name)
		}
	}
	sort.Strings(extraEN)
	sort.Strings(extraZH)

	return &DocLoadStats{
		UniverseCount: len(universeNames),
		ENFiles:       len(enDocs),
		ZHFiles:       len(zhDocs),
		ExtraEN:       extraEN,
		ExtraZH:       extraZH,
		DryRun:        dryRun,
	}, stageRows, nil
}

func dropLegacyDocForeignKeys(ctx context.Context, tx pgx.Tx) error {
	rows, err := tx.Query(ctx, `
		SELECT c.conname, coalesce(
			c.confrelid = 'pgext.extension'::regclass AND c.confdeltype = 'c'
			AND (c.conname, src.attname, dst.attname) IN (
				('doc_id_fkey', 'id', 'id'), ('doc_ext_fkey', 'ext', 'name')
			), false)
		FROM pg_constraint c
		LEFT JOIN pg_attribute src ON src.attrelid = c.conrelid AND c.conkey = ARRAY[src.attnum]
		LEFT JOIN pg_attribute dst ON dst.attrelid = c.confrelid AND c.confkey = ARRAY[dst.attnum]
		WHERE c.conrelid = 'pgext.doc'::regclass AND c.contype = 'f'`)
	if err != nil {
		return fmt.Errorf("inspect pgext.doc foreign keys: %w", err)
	}
	var names []string
	for rows.Next() {
		var name string
		var legacy bool
		if err := rows.Scan(&name, &legacy); err != nil {
			rows.Close()
			return err
		}
		if !legacy {
			rows.Close()
			return fmt.Errorf("unexpected foreign key %q on pgext.doc", name)
		}
		names = append(names, name)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	for _, name := range names {
		if _, err := tx.Exec(ctx, "ALTER TABLE pgext.doc DROP CONSTRAINT "+pgx.Identifier{name}.Sanitize()); err != nil {
			return fmt.Errorf("drop legacy foreign key %s: %w", name, err)
		}
	}
	return nil
}

func verifyLoadedDocs(ctx context.Context, tx pgx.Tx, wantCount int) error {
	var docCount, blankCount, metaMismatch, fkCount int
	if err := tx.QueryRow(ctx, `
		SELECT
			(SELECT count(*) FROM pgext.doc),
			(SELECT count(*) FROM pgext.doc WHERE btrim(coalesce(en_doc, '')) = '' OR btrim(coalesce(zh_doc, '')) = ''),
			(SELECT count(*) FROM pgext.universe u
			 LEFT JOIN pgext.doc d ON d.ext = u.name
			 WHERE d.ext IS NULL
			    OR d.id IS DISTINCT FROM u.id
			    OR d.pkg IS DISTINCT FROM u.pkg
			    OR d.repo_url IS DISTINCT FROM u.url
			    OR d.license_url IS DISTINCT FROM u.license_url
			    OR d.control_url IS DISTINCT FROM u.control_url
			    OR d.author_url IS DISTINCT FROM u.author_url
			    OR d.home_url IS DISTINCT FROM u.home_url
			    OR d.cargo_url IS DISTINCT FROM u.cargo_url),
			(SELECT count(*) FROM pg_constraint WHERE conrelid = 'pgext.doc'::regclass AND contype = 'f')
	`).Scan(&docCount, &blankCount, &metaMismatch, &fkCount); err != nil {
		return fmt.Errorf("postflight verify pgext.doc: %w", err)
	}
	if docCount != wantCount {
		return fmt.Errorf("postflight mismatch: pgext.doc has %d rows, pgext.universe has %d", docCount, wantCount)
	}
	if blankCount > 0 {
		return fmt.Errorf("postflight failure: %d pgext.doc rows have empty en_doc or zh_doc", blankCount)
	}
	if metaMismatch > 0 {
		return fmt.Errorf("postflight failure: %d pgext.doc rows disagree with pgext.universe metadata", metaMismatch)
	}
	if fkCount > 0 {
		return fmt.Errorf("postflight failure: %d foreign key constraints remain on pgext.doc", fkCount)
	}
	return nil
}

func previewNames(names []string) string {
	if len(names) == 0 {
		return "none"
	}
	if len(names) <= 8 {
		return strings.Join(names, ", ")
	}
	return fmt.Sprintf("%s ... (+%d more)", strings.Join(names[:8], ", "), len(names)-8)
}
