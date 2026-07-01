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


DEFAULT_PGURL = os.environ.get("PGURL", "postgres:///data")

DOC_TABLE_SQL = """
CREATE SCHEMA IF NOT EXISTS pgext;

CREATE TABLE IF NOT EXISTS pgext.doc
(
    id          INTEGER PRIMARY KEY REFERENCES pgext.extension (id) ON DELETE CASCADE,
    ext         TEXT NOT NULL UNIQUE REFERENCES pgext.extension (name) ON DELETE CASCADE,
    pkg         TEXT NOT NULL,
    repo_url    TEXT,
    home_url    TEXT,
    license_url TEXT,
    control_url TEXT,
    author_url  TEXT,
    en_doc      TEXT,
    zh_doc      TEXT
);

CREATE INDEX IF NOT EXISTS doc_ext_pkg_idx ON pgext.doc (ext, pkg);
"""

SYNC_EXTENSION_ROWS_SQL = """
INSERT INTO pgext.doc (id, ext, pkg, repo_url)
SELECT id, name, pkg, url
FROM pgext.extension
ON CONFLICT (id) DO UPDATE
SET ext = EXCLUDED.ext,
    pkg = EXCLUDED.pkg,
    repo_url = EXCLUDED.repo_url;

DELETE FROM pgext.doc d
WHERE NOT EXISTS (
    SELECT 1
    FROM pgext.extension e
    WHERE e.id = d.id
);
"""


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

    def load_docs_from_csv(self, column: str, docs: Mapping[str, str], known_exts: set) -> int:
        if column not in {"en_doc", "zh_doc"}:
            raise ValueError(f"unsupported doc column: {column}")
        updates = [(ext, doc) for ext, doc in sorted(docs.items()) if ext in known_exts]
        if not updates:
            return 0

        temp_path = None
        try:
            with tempfile.NamedTemporaryFile(
                "w",
                encoding="utf-8",
                newline="",
                prefix=f"pgext-{column}-",
                suffix=".csv",
                delete=False,
            ) as temp_file:
                temp_path = Path(temp_file.name)
                writer = csv.writer(temp_file, quoting=csv.QUOTE_ALL)
                writer.writerow(["ext", "doc"])
                writer.writerows(updates)

            quoted_path = str(temp_path).replace("'", "''")
            self._run(
                f"""
CREATE TEMP TABLE stub_doc_load (
    ext TEXT PRIMARY KEY,
    doc TEXT
);
\\copy stub_doc_load (ext, doc) FROM '{quoted_path}' WITH (FORMAT csv, HEADER true, NULL '\\N')
UPDATE pgext.doc d
SET {column} = s.doc
FROM stub_doc_load s
WHERE d.ext = s.ext;
"""
            )
            return len(updates)
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
    return {row["ext"] for row in conn.query_csv("SELECT ext FROM pgext.doc")}


def _update_doc_column(conn, column: str, docs: Mapping[str, str], known_exts: set) -> int:
    return conn.load_docs_from_csv(column, docs, known_exts)


def load_stub_docs(conn, en_docs: Mapping[str, str], zh_docs: Mapping[str, str]) -> Dict[str, Any]:
    """Sync extension rows and load local Markdown docs into pgext.doc."""
    ensure_doc_table(conn)
    conn.execute(f"{SYNC_EXTENSION_ROWS_SQL}\nUPDATE pgext.doc SET en_doc = NULL, zh_doc = NULL;")

    known_exts = _known_extensions(conn)
    en_count = _update_doc_column(conn, "en_doc", en_docs, known_exts)
    zh_count = _update_doc_column(conn, "zh_doc", zh_docs, known_exts)

    return {
        "base_rows": len(known_exts),
        "en_doc": en_count,
        "zh_doc": zh_count,
        "extra_en": sorted(set(en_docs) - known_exts),
        "extra_zh": sorted(set(zh_docs) - known_exts),
        "missing_en": sorted(known_exts - set(en_docs)),
        "missing_zh": sorted(known_exts - set(zh_docs)),
    }


def fetch_doc_rows(conn) -> Sequence[Dict[str, Any]]:
    """Fetch pgext.doc rows for dumping to local Markdown files."""
    ensure_doc_table(conn)
    return conn.query_csv("SELECT ext, en_doc, zh_doc FROM pgext.doc ORDER BY id")
