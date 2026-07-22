#!/usr/bin/env python3
"""
Helpers for syncing extension Markdown stubs with pgext.doc.
"""

import csv
import io
import os
import subprocess
import tempfile
from pathlib import Path
from typing import Any, Dict, Iterable, Mapping, Sequence, Tuple


DEFAULT_PGURL = os.environ.get(
    "PGURL", "host=/tmp port=5432 dbname=data user=postgres"
)

DOC_TABLE_SQL = """
CREATE SCHEMA IF NOT EXISTS pgext;

CREATE TABLE IF NOT EXISTS pgext.doc
(
    id          INTEGER PRIMARY KEY,
    ext         TEXT NOT NULL UNIQUE,
    pkg         TEXT NOT NULL,
    repo_url    TEXT,
    home_url    TEXT,
    license_url TEXT,
    control_url TEXT,
    author_url  TEXT,
    cargo_url   TEXT,
    en_doc      TEXT,
    zh_doc      TEXT
);

ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS repo_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS home_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS license_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS control_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS author_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS cargo_url TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS en_doc TEXT;
ALTER TABLE pgext.doc ADD COLUMN IF NOT EXISTS zh_doc TEXT;

CREATE INDEX IF NOT EXISTS doc_ext_pkg_idx ON pgext.doc (ext, pkg);
"""

DOC_BACKUP_TABLE = "pgext.doc_before_universe4_20260716"


class PsqlClient:
    """Small psql wrapper used to avoid a Python database-driver dependency."""

    def __init__(self, pgurl: str):
        self.pgurl = pgurl

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc, tb):
        return False

    def _run(self, sql: str, capture: bool = False) -> str:
        cmd = ["psql", "-X", "-q", "-v", "ON_ERROR_STOP=1", "--dbname", self.pgurl]
        result = subprocess.run(
            cmd,
            input=sql,
            text=True,
            capture_output=True,
            check=False,
        )
        if result.returncode != 0:
            detail = result.stderr.strip() or result.stdout.strip()
            raise RuntimeError(f"psql failed with exit code {result.returncode}: {detail}")
        return result.stdout if capture else ""

    def execute(self, sql: str) -> None:
        self._run(sql)

    def query_csv(self, query: str) -> Sequence[Dict[str, Any]]:
        body = query.strip().rstrip(";")
        output = self._run(f"COPY ({body}) TO STDOUT WITH CSV HEADER NULL '\\N';", capture=True)
        if not output:
            return []
        rows = []
        for row in csv.DictReader(io.StringIO(output)):
            rows.append({key: None if value == "\\N" else value for key, value in row.items()})
        return rows

    def sync_docs_from_csv(
        self,
        docs: Sequence[Tuple[str, str, str]],
        expected_count: int,
        commit: bool = True,
    ) -> None:
        """Atomically replace pgext.doc's universe with staged bilingual docs."""
        if not docs:
            raise ValueError("refusing to sync an empty documentation set")
        transaction_end = "COMMIT" if commit else "ROLLBACK"

        temp_path = None
        try:
            with tempfile.NamedTemporaryFile(
                "w",
                encoding="utf-8",
                newline="",
                prefix="pgext-doc-sync-",
                suffix=".csv",
                delete=False,
            ) as temp_file:
                temp_path = Path(temp_file.name)
                writer = csv.writer(temp_file, quoting=csv.QUOTE_ALL)
                writer.writerow(["ext", "en_doc", "zh_doc"])
                writer.writerows(docs)

            quoted_path = str(temp_path).replace("'", "''")
            self._run(
                f"""
BEGIN;
LOCK TABLE pgext.universe IN SHARE MODE;
LOCK TABLE pgext.extension IN SHARE MODE;
LOCK TABLE pgext.doc IN SHARE ROW EXCLUSIVE MODE;

CREATE TEMP TABLE stub_doc_load (
    ext TEXT PRIMARY KEY,
    en_doc TEXT NOT NULL,
    zh_doc TEXT NOT NULL
) ON COMMIT DROP;
\\copy stub_doc_load (ext, en_doc, zh_doc) FROM '{quoted_path}' WITH (FORMAT csv, HEADER true, NULL '\\N')

DO $doc_preflight$
DECLARE
    stage_count INTEGER;
    universe_count INTEGER;
BEGIN
    SELECT count(*) INTO stage_count FROM stub_doc_load;
    SELECT count(*) INTO universe_count FROM pgext.universe;

    IF stage_count <> {expected_count} OR universe_count <> {expected_count} THEN
        RAISE EXCEPTION
            'documentation cardinality drift: stage=%, universe=%, expected={expected_count}',
            stage_count, universe_count;
    END IF;

    IF EXISTS (
        SELECT ext FROM stub_doc_load
        EXCEPT
        SELECT name FROM pgext.universe
    ) OR EXISTS (
        SELECT name FROM pgext.universe
        EXCEPT
        SELECT ext FROM stub_doc_load
    ) THEN
        RAISE EXCEPTION 'staged documentation names do not exactly match pgext.universe';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM stub_doc_load
        WHERE btrim(en_doc) = '' OR btrim(zh_doc) = ''
    ) THEN
        RAISE EXCEPTION 'staged documentation contains an empty body';
    END IF;

    IF EXISTS (
        SELECT name FROM pgext.extension
        EXCEPT
        SELECT ext FROM stub_doc_load
    ) OR EXISTS (
        SELECT name FROM pgext.extension
        EXCEPT
        SELECT name FROM pgext.universe
    ) THEN
        RAISE EXCEPTION
            'pgext.extension is not covered by the staged universe documentation';
    END IF;

    IF to_regclass('{DOC_BACKUP_TABLE}') IS NULL THEN
        EXECUTE 'CREATE TABLE {DOC_BACKUP_TABLE} AS TABLE pgext.doc WITH NO DATA';
        EXECUTE $doc_backup$
            INSERT INTO {DOC_BACKUP_TABLE} (
                id, ext, pkg, repo_url, home_url, license_url, control_url,
                author_url, cargo_url, en_doc, zh_doc
            )
            SELECT
                e.id,
                e.name,
                e.pkg,
                u.url,
                u.home_url,
                u.license_url,
                u.control_url,
                u.author_url,
                u.cargo_url,
                s.en_doc,
                s.zh_doc
            FROM pgext.extension AS e
            JOIN pgext.universe AS u ON u.name = e.name
            JOIN stub_doc_load AS s ON s.ext = e.name
            ORDER BY e.id
        $doc_backup$;
        EXECUTE 'ALTER TABLE {DOC_BACKUP_TABLE} ADD PRIMARY KEY (id)';
        EXECUTE 'CREATE UNIQUE INDEX doc_before_universe4_20260716_ext_idx '
             || 'ON {DOC_BACKUP_TABLE} (ext)';
    END IF;

    IF NOT EXISTS (SELECT 1 FROM {DOC_BACKUP_TABLE})
       OR EXISTS (
            SELECT 1
            FROM {DOC_BACKUP_TABLE}
            WHERE en_doc IS NULL OR btrim(en_doc) = ''
               OR zh_doc IS NULL OR btrim(zh_doc) = ''
       ) THEN
        RAISE EXCEPTION 'documentation baseline backup is empty or incomplete';
    END IF;
END
$doc_preflight$;

ALTER TABLE pgext.doc DROP CONSTRAINT IF EXISTS doc_id_fkey;
ALTER TABLE pgext.doc DROP CONSTRAINT IF EXISTS doc_ext_fkey;

COMMENT ON TABLE pgext.doc IS
    'PostgreSQL extension documentation synchronized atomically from stub directories against pgext.universe';
COMMENT ON COLUMN pgext.doc.id IS
    'Extension identifier copied from pgext.universe.id and validated by the loader';
COMMENT ON COLUMN pgext.doc.ext IS
    'Extension name copied from pgext.universe.name and validated by the loader';
COMMENT ON COLUMN pgext.doc.pkg IS
    'Package name copied from pgext.universe.pkg';
COMMENT ON COLUMN pgext.doc.repo_url IS
    'Primary project URL seeded from pgext.universe.url';
COMMENT ON COLUMN pgext.doc.license_url IS
    'Optional upstream license URL';
COMMENT ON COLUMN pgext.doc.control_url IS
    'Optional upstream control file URL';
COMMENT ON COLUMN pgext.doc.author_url IS
    'Optional upstream author or maintainer URL';
COMMENT ON COLUMN pgext.doc.home_url IS
    'Optional upstream homepage URL';
COMMENT ON COLUMN pgext.doc.cargo_url IS
    'pgrx Cargo.toml file URL for rust extension';
COMMENT ON COLUMN pgext.doc.en_doc IS
    'English Markdown documentation from stub/<ext>.md';
COMMENT ON COLUMN pgext.doc.zh_doc IS
    'Chinese Markdown documentation from stub-zh/<ext>.md';
COMMENT ON TABLE {DOC_BACKUP_TABLE} IS
    'Historical immutable pgext.extension documentation snapshot captured before universe synchronization; live Markdown may receive later corrections';

DELETE FROM pgext.doc;

INSERT INTO pgext.doc (
    id, ext, pkg, repo_url, home_url, license_url, control_url, author_url,
    cargo_url, en_doc, zh_doc
)
SELECT
    u.id,
    u.name,
    u.pkg,
    u.url,
    u.home_url,
    u.license_url,
    u.control_url,
    u.author_url,
    u.cargo_url,
    s.en_doc,
    s.zh_doc
FROM pgext.universe AS u
JOIN stub_doc_load AS s ON s.ext = u.name
ORDER BY u.id;

DO $doc_postflight$
BEGIN
    IF (SELECT count(*) FROM pgext.doc) <> {expected_count}
       OR EXISTS (
            SELECT name FROM pgext.universe
            EXCEPT
            SELECT ext FROM pgext.doc
       )
       OR EXISTS (
            SELECT ext FROM pgext.doc
            EXCEPT
            SELECT name FROM pgext.universe
       ) THEN
        RAISE EXCEPTION 'pgext.doc does not exactly match pgext.universe after sync';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pgext.doc
        WHERE en_doc IS NULL OR btrim(en_doc) = ''
           OR zh_doc IS NULL OR btrim(zh_doc) = ''
    ) THEN
        RAISE EXCEPTION 'pgext.doc contains missing bilingual Markdown after sync';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pgext.doc AS d
        JOIN pgext.universe AS u ON u.name = d.ext
        WHERE d.id <> u.id
           OR d.pkg IS DISTINCT FROM u.pkg
           OR d.repo_url IS DISTINCT FROM u.url
           OR d.home_url IS DISTINCT FROM u.home_url
           OR d.license_url IS DISTINCT FROM u.license_url
           OR d.control_url IS DISTINCT FROM u.control_url
           OR d.author_url IS DISTINCT FROM u.author_url
           OR d.cargo_url IS DISTINCT FROM u.cargo_url
    ) THEN
        RAISE EXCEPTION 'pgext.doc metadata differs from pgext.universe after sync';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conrelid = 'pgext.doc'::regclass
          AND contype = 'f'
    ) THEN
        RAISE EXCEPTION 'pgext.doc still has a foreign key after universe migration';
    END IF;

END
$doc_postflight$;

{transaction_end};
"""
            )
        finally:
            if temp_path is not None:
                temp_path.unlink(missing_ok=True)


def connect(pgurl: str) -> PsqlClient:
    return PsqlClient(pgurl)


def safe_stub_name(name: str) -> str:
    """Return a safe stub file stem or raise ValueError."""
    if not name or "\0" in name:
        raise ValueError("stub name cannot be empty or contain NUL")
    path = Path(name)
    if path.name != name or path.is_absolute():
        raise ValueError(f"unsafe stub name: {name!r}")
    return name


def collect_stub_docs(stub_dir: Path) -> Dict[str, str]:
    """Read all Markdown stub files from a directory keyed by file stem."""
    stub_dir = Path(stub_dir)
    if not stub_dir.exists():
        return {}
    docs: Dict[str, str] = {}
    for path in sorted(stub_dir.glob("*.md")):
        if path.is_file():
            ext = safe_stub_name(path.stem)
            docs[ext] = path.read_text(encoding="utf-8")
    return docs


def _row_get(row: Any, key: str) -> Any:
    if isinstance(row, Mapping):
        return row.get(key)
    return getattr(row, key)


def dump_stub_docs(rows: Iterable[Any], en_dir: Path, zh_dir: Path) -> Tuple[int, int]:
    """Write non-null pgext.doc Markdown rows into local stub directories."""
    en_dir = Path(en_dir)
    zh_dir = Path(zh_dir)
    en_dir.mkdir(parents=True, exist_ok=True)
    zh_dir.mkdir(parents=True, exist_ok=True)

    en_count = 0
    zh_count = 0
    for row in rows:
        ext = safe_stub_name(str(_row_get(row, "ext")))
        en_doc = _row_get(row, "en_doc")
        zh_doc = _row_get(row, "zh_doc")

        if en_doc is not None:
            (en_dir / f"{ext}.md").write_text(en_doc, encoding="utf-8")
            en_count += 1
        if zh_doc is not None:
            (zh_dir / f"{ext}.md").write_text(zh_doc, encoding="utf-8")
            zh_count += 1

    return en_count, zh_count


def ensure_doc_table(conn) -> None:
    """Create pgext.doc if it does not already exist."""
    conn.execute(DOC_TABLE_SQL)


def _known_extensions(conn) -> set:
    return {
        row["ext"]
        for row in conn.query_csv("SELECT name AS ext FROM pgext.universe ORDER BY name")
    }


def load_stub_docs(
    conn,
    en_docs: Mapping[str, str],
    zh_docs: Mapping[str, str],
    dry_run: bool = False,
) -> Dict[str, Any]:
    """Atomically sync complete universe Markdown pairs into pgext.doc."""
    ensure_doc_table(conn)
    known_exts = _known_extensions(conn)
    missing_en = sorted(known_exts - set(en_docs))
    missing_zh = sorted(known_exts - set(zh_docs))
    empty_en = sorted(ext for ext in known_exts if ext in en_docs and not en_docs[ext].strip())
    empty_zh = sorted(ext for ext in known_exts if ext in zh_docs and not zh_docs[ext].strip())
    if missing_en or missing_zh or empty_en or empty_zh:
        raise ValueError(
            "refusing incomplete pgext.doc sync: "
            f"missing_en={len(missing_en)} missing_zh={len(missing_zh)} "
            f"empty_en={len(empty_en)} empty_zh={len(empty_zh)}"
        )

    rows = [(ext, en_docs[ext], zh_docs[ext]) for ext in sorted(known_exts)]
    conn.sync_docs_from_csv(rows, expected_count=len(known_exts), commit=not dry_run)

    return {
        "base_rows": len(known_exts),
        "en_doc": len(rows),
        "zh_doc": len(rows),
        "extra_en": sorted(set(en_docs) - known_exts),
        "extra_zh": sorted(set(zh_docs) - known_exts),
        "missing_en": [],
        "missing_zh": [],
        "dry_run": dry_run,
    }


def fetch_doc_rows(conn) -> Sequence[Dict[str, Any]]:
    """Fetch pgext.doc rows for dumping to local Markdown files."""
    ensure_doc_table(conn)
    return conn.query_csv("SELECT ext, en_doc, zh_doc FROM pgext.doc ORDER BY id")
