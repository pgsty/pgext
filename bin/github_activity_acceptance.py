#!/usr/bin/env python3
"""Read-only Phase A acceptance evidence for the universe GitHub refresh.

This utility never calls GitHub and never writes to PostgreSQL.  Both database
commands are pinned to postgres on the local /tmp socket, port 5432, database
data.  ``snapshot`` captures the pre-run baseline; ``report`` compares a fresh
read-only capture with that baseline and writes complete JSON and Markdown
evidence.

Exit status for ``report`` is 0 only when every acceptance gate passes, 1 when
the report was written but at least one gate failed, and 2 for refusal, invalid
input, or another execution error.
"""

from __future__ import annotations

import argparse
import csv
import hashlib
import json
import os
import pathlib
import subprocess
import sys
import tempfile
import unittest
import urllib.parse
from datetime import datetime, timezone
from typing import Any, Iterable

SCHEMA_VERSION = 1
TOOL_NAME = "github_activity_acceptance.py"
REST_API = "https://api.github.com"
GRAPHQL_API = "https://api.github.com/graphql"
EXPECTED_BACKUP_ROWS = 1342
EXPECTED_BACKUP_FINGERPRINT = "9354ad62777856b2d60dc7d68f110e2b"
EXPECTED_BACKUP_OWNER = "postgres"
RATE_LIMIT_RESERVE = 200
REQUIRED_RELATIONS = {
    "gh_repo": "pgext.gh_repo",
    "gh_repo_20260507": "pgext.gh_repo_20260507",
    "universe": "pgext.universe",
    "extension": "pgext.extension",
}
REQUIRED_MANIFEST_COLUMNS = {
    "url_norm",
    "repo_owner",
    "repo_name",
    "api_url",
    "extension_ids",
    "extension_names",
    "extension_count",
}
TERMINAL_STATUSES = {"fetched", "blocked"}
TRANSIENT_STATUSES = {"pending", "collecting", "error", "rate_limited"}
METADATA_COVERAGE_FIELDS = (
    "api_json",
    "etag",
    "rate_limit_remaining",
    "rate_limit_reset",
    "stargazers_count",
    "forks_count",
    "watchers_count",
    "subscribers_count",
    "pushed_at",
    "updated_at_api",
    "default_branch",
    "archived",
    "fetched_at",
    "activity_json",
)
STATUS_LIST_ORDER = (
    "blocked",
    "error",
    "rate_limited",
    "pending",
    "collecting",
    "missing",
)


class AcceptanceError(RuntimeError):
    """Base class for fail-closed acceptance errors."""


class ManifestError(AcceptanceError):
    """Raised when the mandatory target manifest is invalid."""


class PreconditionError(AcceptanceError):
    """Raised when the fixed database or required relations are unavailable."""


class SnapshotError(AcceptanceError):
    """Raised when a snapshot is invalid or does not belong to this run."""


def utc_now() -> str:
    return datetime.now(timezone.utc).isoformat().replace("+00:00", "Z")


def canonical_json_bytes(value: Any) -> bytes:
    return json.dumps(
        value,
        ensure_ascii=False,
        sort_keys=True,
        separators=(",", ":"),
    ).encode("utf-8")


def digest_json(value: Any) -> str:
    return hashlib.sha256(canonical_json_bytes(value)).hexdigest()


def sha256_file(path: pathlib.Path) -> str:
    hasher = hashlib.sha256()
    with path.open("rb") as handle:
        for chunk in iter(lambda: handle.read(1024 * 1024), b""):
            hasher.update(chunk)
    return hasher.hexdigest()


def parse_pg_array(value: str | None) -> list[str | None]:
    """Parse the one-dimensional PostgreSQL array literals used by the CSV."""

    text = (value or "").strip()
    if len(text) < 2 or text[0] != "{" or text[-1] != "}":
        raise ManifestError(f"invalid PostgreSQL array literal: {value!r}")
    body = text[1:-1]
    if not body:
        return []

    result: list[str | None] = []
    token: list[str] = []
    quoted = False
    escaped = False
    token_was_quoted = False
    for char in body:
        if escaped:
            token.append(char)
            escaped = False
            continue
        if char == "\\":
            escaped = True
            continue
        if quoted:
            if char == '"':
                quoted = False
                token_was_quoted = True
            else:
                token.append(char)
            continue
        if char == '"':
            if token:
                raise ManifestError(
                    f"invalid quote placement in array literal: {value!r}"
                )
            quoted = True
            continue
        if char == ",":
            item = "".join(token)
            result.append(None if not token_was_quoted and item == "NULL" else item)
            token = []
            token_was_quoted = False
            continue
        token.append(char)
    if escaped or quoted:
        raise ManifestError(f"unterminated escape or quote in array literal: {value!r}")
    item = "".join(token)
    result.append(None if not token_was_quoted and item == "NULL" else item)
    return result


def normalize_github_repo(raw_url: str) -> tuple[str, str, str]:
    value = (raw_url or "").strip()
    parsed = urllib.parse.urlparse(value)
    host = (parsed.hostname or "").lower()
    if parsed.scheme.lower() not in {"http", "https"} or host not in {
        "github.com",
        "www.github.com",
    }:
        raise ManifestError(f"not an HTTP GitHub repository URL: {raw_url!r}")
    parts = [urllib.parse.unquote(part) for part in parsed.path.split("/") if part]
    if len(parts) < 2:
        raise ManifestError(f"not a GitHub owner/repository URL: {raw_url!r}")
    owner, repo = parts[:2]
    if repo.lower().endswith(".git"):
        repo = repo[:-4]
    if not owner or not repo:
        raise ManifestError(f"empty GitHub owner/repository: {raw_url!r}")
    return owner, repo, f"https://github.com/{owner.lower()}/{repo.lower()}"


def validate_manifest_row(raw: dict[str, str], line_number: int) -> dict[str, Any]:
    prefix = f"manifest line {line_number}"
    url_norm = (raw.get("url_norm") or "").strip()
    try:
        _, _, normalized = normalize_github_repo(url_norm)
    except ManifestError as exc:
        raise ManifestError(f"{prefix}: {exc}") from exc
    if url_norm != normalized:
        raise ManifestError(
            f"{prefix}: noncanonical url_norm {url_norm!r}; expected {normalized!r}"
        )

    url_owner, url_repo = url_norm.removeprefix("https://github.com/").split("/", 1)
    owner = (raw.get("repo_owner") or "").strip()
    repo = (raw.get("repo_name") or "").strip()
    if owner != url_owner or repo != url_repo:
        raise ManifestError(
            f"{prefix}: lowercase owner/repo columns must match url_norm"
        )
    api_url = (raw.get("api_url") or "").strip()
    expected_api_url = f"{REST_API}/repos/{url_owner}/{url_repo}"
    if api_url != expected_api_url:
        raise ManifestError(f"{prefix}: api_url must be {expected_api_url!r}")

    ids_raw = parse_pg_array(raw.get("extension_ids"))
    names_raw = parse_pg_array(raw.get("extension_names"))
    if any(item in (None, "") for item in ids_raw + names_raw):
        raise ManifestError(f"{prefix}: extension arrays contain NULL/empty values")
    try:
        extension_ids = [int(item) for item in ids_raw]
        extension_count = int((raw.get("extension_count") or "").strip())
    except (TypeError, ValueError) as exc:
        raise ManifestError(f"{prefix}: invalid extension id/count") from exc
    extension_names = [str(item) for item in names_raw]
    if extension_count <= 0:
        raise ManifestError(f"{prefix}: extension_count must be positive")
    if extension_count != len(extension_ids) or extension_count != len(extension_names):
        raise ManifestError(f"{prefix}: extension_count disagrees with arrays")
    if len(set(extension_ids)) != len(extension_ids):
        raise ManifestError(f"{prefix}: duplicate extension id")
    if len(set(extension_names)) != len(extension_names):
        raise ManifestError(f"{prefix}: duplicate extension name")
    if extension_ids != sorted(extension_ids):
        raise ManifestError(f"{prefix}: extension_ids are not sorted")
    if extension_names != sorted(
        extension_names, key=lambda item: item.encode("utf-8")
    ):
        raise ManifestError(f"{prefix}: extension_names are not in C order")

    return {
        "url_norm": url_norm,
        "repo_owner": owner,
        "repo_name": repo,
        "api_url": api_url,
        "extension_ids": extension_ids,
        "extension_names": extension_names,
        "extension_count": extension_count,
        "prior_url_candidates": (raw.get("prior_url_candidates") or "").strip(),
        "action": (raw.get("action") or "").strip(),
    }


def load_manifest(
    path_value: str,
) -> tuple[pathlib.Path, list[dict[str, Any]], dict[str, Any]]:
    path = pathlib.Path(path_value).expanduser().resolve()
    if not path.is_file():
        raise ManifestError(f"manifest does not exist: {path}")
    with path.open(newline="", encoding="utf-8") as handle:
        reader = csv.DictReader(handle)
        fieldnames = set(reader.fieldnames or [])
        missing = REQUIRED_MANIFEST_COLUMNS - fieldnames
        if missing:
            raise ManifestError(
                f"manifest missing columns: {', '.join(sorted(missing))}"
            )
        rows = [
            validate_manifest_row(raw, line_number)
            for line_number, raw in enumerate(reader, start=2)
        ]
    if not rows:
        raise ManifestError("manifest has no target rows")
    urls = [row["url_norm"] for row in rows]
    if len(urls) != len(set(urls)):
        duplicates = sorted({url for url in urls if urls.count(url) > 1})
        raise ManifestError(f"manifest has duplicate url_norm values: {duplicates[:5]}")
    if urls != sorted(urls, key=lambda item: item.encode("utf-8")):
        raise ManifestError("manifest rows are not in C order by url_norm")

    signature_rows = [
        {
            key: row[key]
            for key in (
                "url_norm",
                "repo_owner",
                "repo_name",
                "api_url",
                "extension_ids",
                "extension_names",
                "extension_count",
            )
        }
        for row in rows
    ]
    metadata = {
        "path": str(path),
        "sha256": sha256_file(path),
        "target_count": len(rows),
        "mapping_sha256": digest_json(signature_rows),
    }
    return path, rows, metadata


def psql_environment() -> dict[str, str]:
    env = os.environ.copy()
    for name in (
        "PGHOSTADDR",
        "PGSERVICE",
        "PGSERVICEFILE",
        "PGTARGETSESSIONATTRS",
        "PGOPTIONS",
        "GITHUB_TOKEN",
        "GH_TOKEN",
    ):
        env.pop(name, None)
    env.update(
        {
            "PGUSER": "postgres",
            "PGHOST": "/tmp",
            "PGPORT": "5432",
            "PGDATABASE": "data",
            "PGOPTIONS": (
                "-c default_transaction_read_only=on "
                "-c statement_timeout=120000 "
                "-c lock_timeout=5000"
            ),
        }
    )
    return env


def run_psql_json(sql: str) -> Any:
    command = [
        "psql",
        "-X",
        "-q",
        "-t",
        "-A",
        "-v",
        "ON_ERROR_STOP=1",
        "-U",
        "postgres",
        "-h",
        "/tmp",
        "-p",
        "5432",
        "-d",
        "data",
        "-c",
        sql,
    ]
    process = subprocess.run(
        command,
        text=True,
        capture_output=True,
        env=psql_environment(),
    )
    if process.returncode != 0:
        message = process.stderr.strip() or process.stdout.strip()
        raise PreconditionError(f"read-only psql query failed: {message}")
    lines = [line for line in process.stdout.splitlines() if line.strip()]
    if len(lines) != 1:
        raise PreconditionError(
            f"read-only psql query returned {len(lines)} nonempty lines; expected one"
        )
    try:
        return json.loads(lines[0])
    except json.JSONDecodeError as exc:
        raise PreconditionError(
            "read-only psql query did not return valid JSON"
        ) from exc


IDENTITY_SQL = r"""
SELECT jsonb_build_object(
    'database', current_database(),
    'session_user', session_user,
    'current_user', current_user,
    'server_addr', inet_server_addr(),
    'port', current_setting('port'),
    'server_version', current_setting('server_version'),
    'cluster_system_identifier', (SELECT system_identifier::text FROM pg_control_system()),
    'transaction_read_only', current_setting('transaction_read_only'),
    'relations', jsonb_build_object(
        'gh_repo', to_regclass('pgext.gh_repo')::text,
        'gh_repo_20260507', to_regclass('pgext.gh_repo_20260507')::text,
        'universe', to_regclass('pgext.universe')::text,
        'extension', to_regclass('pgext.extension')::text
    )
)::text;
"""


LIVE_TARGET_CTE = r"""
WITH normalized AS (
    SELECT id, name, pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\.)?github\.com/'
      AND pgext.repo_url_norm(repo_url) IS NOT NULL
), live_target AS (
    SELECT url_norm,
           split_part(url_norm, '/', 4) AS repo_owner,
           split_part(url_norm, '/', 5) AS repo_name,
           'https://api.github.com/repos/' || split_part(url_norm, '/', 4) || '/' || split_part(url_norm, '/', 5) AS api_url,
           array_agg(id ORDER BY id) AS extension_ids,
           array_agg(name ORDER BY name COLLATE "C") AS extension_names,
           count(*)::integer AS extension_count
    FROM normalized
    GROUP BY url_norm
)
"""


STATE_SQL = LIVE_TARGET_CTE + r"""
, target_state AS (
    SELECT
        t.url_norm,
        t.repo_owner AS target_repo_owner,
        t.repo_name AS target_repo_name,
        t.api_url AS target_api_url,
        t.extension_ids AS target_extension_ids,
        t.extension_names AS target_extension_names,
        t.extension_count AS target_extension_count,
        r.url_norm IS NOT NULL AS row_exists,
        r.repo_host,
        r.repo_owner,
        r.repo_name,
        r.api_url,
        r.extension_ids,
        r.extension_names,
        r.extension_count,
        r.status,
        r.attempts,
        r.http_status,
        r.error,
        r.etag,
        r.rate_limit_remaining,
        r.rate_limit_reset,
        r.stargazers_count,
        r.forks_count,
        r.watchers_count,
        r.subscribers_count,
        r.pushed_at,
        r.updated_at_api,
        r.default_branch,
        r.archived,
        r.fetched_at,
        r.created_at,
        r.updated_at,
        r.last_commit_at,
        r.last_commit_date,
        r.last_commit_sha,
        r.last_commit_html_url,
        r.latest_release_tag,
        r.latest_release_published_at,
        r.latest_release_created_at,
        r.last_release_at,
        r.last_release_date,
        r.latest_tag_name,
        r.latest_tag_at,
        r.last_tag_date,
        r.latest_tag_date_source,
        r.latest_tag_target_type,
        r.latest_tag_commit_sha,
        r.latest_tag_commit_at,
        r.last_release_or_tag_at,
        r.last_release_or_tag_date,
        r.last_release_or_tag_source,
        r.last_update_at,
        r.last_update_date,
        r.last_update_source,
        r.api_json IS NOT NULL AS api_json_present,
        r.activity_json IS NOT NULL AS activity_json_present,
        r.api_json->'stargazers_count' AS api_stargazers_count,
        r.api_json->'forks_count' AS api_forks_count,
        r.api_json->'watchers_count' AS api_watchers_count,
        r.api_json->'subscribers_count' AS api_subscribers_count,
        r.api_json->>'pushed_at' AS api_pushed_at,
        r.api_json->>'updated_at' AS api_updated_at,
        r.api_json->>'default_branch' AS api_default_branch,
        r.api_json->'archived' AS api_archived,
        r.api_json->>'full_name' AS api_full_name,
        r.api_json->>'html_url' AS api_html_url,
        r.activity_json->>'collected_at' AS activity_collected_at,
        r.activity_json->'requested' AS requested_evidence,
        (r.activity_json->'repo') - 'data' AS repo_evidence,
        (r.activity_json->'repo'->'data' IS NOT DISTINCT FROM r.api_json) AS repo_evidence_matches_api_json,
        r.activity_json->'latest_commit' AS latest_commit_evidence,
        r.activity_json->'release_tag' AS release_tag_evidence,
        r.activity_json->'redirect'->>'detected' = 'true' AS redirect_detected,
        r.activity_json->'redirect'->>'canonical_url' AS redirect_canonical_url,
        r.activity_json->'redirect'->>'full_name' AS redirect_full_name,
        r.activity_json->'redirect'->>'html_url' AS redirect_html_url,
        r.activity_json->'redirect'->>'final_api_url' AS redirect_final_api_url,
        CASE
            WHEN jsonb_typeof(r.activity_json->'warnings') = 'array'
            THEN r.activity_json->'warnings'
            ELSE '[]'::jsonb
        END AS activity_warnings,
        CASE WHEN r.url_norm IS NULL THEN NULL ELSE md5((
            to_jsonb(r) - ARRAY[
                'extension_ids', 'extension_names', 'extension_count',
                'status', 'attempts', 'http_status', 'rate_limit_remaining',
                'rate_limit_reset', 'error', 'updated_at'
            ]
        )::text) END AS success_data_fingerprint
    FROM live_target AS t
    LEFT JOIN pgext.gh_repo AS r USING (url_norm)
), backup_summary AS (
    SELECT count(*) AS rows,
           md5(coalesce(string_agg(md5(to_jsonb(b)::text), '' ORDER BY url_norm), '')) AS fingerprint,
           (SELECT pg_get_userbyid(c.relowner)
            FROM pg_class AS c
            WHERE c.oid = 'pgext.gh_repo_20260507'::regclass) AS owner
    FROM pgext.gh_repo_20260507 AS b
), non_target_summary AS (
    SELECT count(*) AS rows,
           md5(coalesce(string_agg(md5(to_jsonb(r)::text), '' ORDER BY url_norm), '')) AS fingerprint
    FROM pgext.gh_repo AS r
    WHERE NOT EXISTS (SELECT 1 FROM live_target AS t WHERE t.url_norm = r.url_norm)
), universe_rows AS (
    SELECT to_jsonb(u)::text AS row_text, md5(to_jsonb(u)::text) AS row_hash
    FROM pgext.universe AS u
), universe_summary AS (
    SELECT count(*) AS rows,
           md5(coalesce(string_agg(row_hash, '' ORDER BY row_hash, row_text), '')) AS fingerprint,
           (SELECT pg_get_userbyid(c.relowner)
            FROM pg_class AS c
            WHERE c.oid = 'pgext.universe'::regclass) AS owner
    FROM universe_rows
), extension_rows AS (
    SELECT to_jsonb(e)::text AS row_text, md5(to_jsonb(e)::text) AS row_hash
    FROM pgext.extension AS e
), extension_summary AS (
    SELECT count(*) AS rows,
           md5(coalesce(string_agg(row_hash, '' ORDER BY row_hash, row_text), '')) AS fingerprint,
           (SELECT pg_get_userbyid(c.relowner)
            FROM pg_class AS c
            WHERE c.oid = 'pgext.extension'::regclass) AS owner
    FROM extension_rows
), duplicate_keys AS (
    SELECT coalesce(jsonb_agg(to_jsonb(d) ORDER BY url_norm), '[]'::jsonb) AS rows
    FROM (
        SELECT url_norm, count(*) AS row_count
        FROM pgext.gh_repo
        GROUP BY url_norm
        HAVING count(*) <> 1
    ) AS d
), invalid_url_keys AS (
    SELECT coalesce(jsonb_agg(to_jsonb(k) ORDER BY url_norm), '[]'::jsonb) AS rows
    FROM (
        SELECT r.url_norm, r.repo_host
        FROM pgext.gh_repo AS r
        WHERE EXISTS (SELECT 1 FROM live_target AS t WHERE t.url_norm = r.url_norm)
          AND (
              r.url_norm !~ '^https://github\.com/[^/]+/[^/]+$'
              OR r.url_norm <> lower(r.url_norm)
              OR r.repo_host <> 'github.com'
          )
    ) AS k
), array_anomalies AS (
    SELECT coalesce(jsonb_agg(to_jsonb(a) ORDER BY url_norm), '[]'::jsonb) AS rows
    FROM (
        SELECT url_norm, extension_count,
               cardinality(extension_ids) AS id_count,
               cardinality(extension_names) AS name_count
        FROM pgext.gh_repo AS r
        WHERE extension_ids IS NULL
           OR extension_names IS NULL
           OR extension_count IS NULL
           OR extension_count <> cardinality(extension_ids)
           OR extension_count <> cardinality(extension_names)
           OR array_position(extension_ids, NULL) IS NOT NULL
           OR array_position(extension_names, NULL) IS NOT NULL
           OR extension_ids IS DISTINCT FROM ARRAY(
               SELECT value FROM unnest(r.extension_ids) AS value ORDER BY value
           )
           OR extension_names IS DISTINCT FROM ARRAY(
               SELECT value FROM unnest(r.extension_names) AS value ORDER BY value COLLATE "C"
           )
           OR cardinality(extension_ids) <> (
               SELECT count(*) FROM (SELECT DISTINCT value FROM unnest(r.extension_ids) AS value) AS x
           )
           OR cardinality(extension_names) <> (
               SELECT count(*) FROM (SELECT DISTINCT value FROM unnest(r.extension_names) AS value) AS x
           )
    ) AS a
)
SELECT jsonb_build_object(
    'captured_at', clock_timestamp(),
    'protected', jsonb_build_object(
        'non_target', (SELECT to_jsonb(n) FROM non_target_summary AS n),
        'backup', (SELECT to_jsonb(b) FROM backup_summary AS b),
        'universe', (SELECT to_jsonb(u) FROM universe_summary AS u),
        'extension', (SELECT to_jsonb(e) FROM extension_summary AS e)
    ),
    'live', jsonb_build_object(
        'rows', (SELECT count(*) FROM pgext.gh_repo),
        'status', (
            SELECT coalesce(jsonb_object_agg(status, row_count), '{}'::jsonb)
            FROM (SELECT status, count(*) AS row_count FROM pgext.gh_repo GROUP BY status) AS s
        ),
        'target_count', (SELECT count(*) FROM live_target),
        'target_existing', (SELECT count(*) FROM target_state WHERE row_exists),
        'target_missing', (SELECT count(*) FROM target_state WHERE NOT row_exists),
        'target_status', (
            SELECT coalesce(jsonb_object_agg(status_key, row_count), '{}'::jsonb)
            FROM (
                SELECT coalesce(status, 'missing') AS status_key, count(*) AS row_count
                FROM target_state GROUP BY coalesce(status, 'missing')
            ) AS s
        ),
        'targets', (
            SELECT coalesce(jsonb_agg(to_jsonb(t) ORDER BY url_norm), '[]'::jsonb)
            FROM target_state AS t
        ),
        'integrity', jsonb_build_object(
            'duplicate_url_norms', (SELECT rows FROM duplicate_keys),
            'invalid_url_keys', (SELECT rows FROM invalid_url_keys),
            'array_anomalies', (SELECT rows FROM array_anomalies)
        )
    )
)::text;
"""


def database_identity() -> dict[str, Any]:
    identity = run_psql_json(IDENTITY_SQL)
    if not isinstance(identity, dict):
        raise PreconditionError("database identity query returned the wrong shape")
    if (
        identity.get("database") != "data"
        or identity.get("session_user") != "postgres"
        or identity.get("current_user") != "postgres"
        or identity.get("server_addr") is not None
        or identity.get("port") != "5432"
        or identity.get("transaction_read_only") != "on"
    ):
        raise PreconditionError(
            "refusing database endpoint other than read-only postgres@/tmp:5432/data: "
            + json.dumps(identity, sort_keys=True)
        )
    relations = identity.get("relations") or {}
    absent = [
        name
        for name, expected in REQUIRED_RELATIONS.items()
        if relations.get(name) != expected
    ]
    if absent:
        raise PreconditionError(
            "required pgext relations absent: " + ", ".join(sorted(absent))
        )
    return identity


def collect_state(identity: dict[str, Any]) -> dict[str, Any]:
    state = run_psql_json(STATE_SQL)
    if not isinstance(state, dict):
        raise PreconditionError("database state query returned the wrong shape")
    state["database"] = identity
    return state


def manifest_mapping_signature(row: dict[str, Any]) -> tuple[Any, ...]:
    return (
        row["repo_owner"],
        row["repo_name"],
        row["api_url"],
        tuple(row["extension_ids"]),
        tuple(row["extension_names"]),
        int(row["extension_count"]),
    )


def live_mapping_signature(row: dict[str, Any]) -> tuple[Any, ...]:
    return (
        row["target_repo_owner"],
        row["target_repo_name"],
        row["target_api_url"],
        tuple(row["target_extension_ids"]),
        tuple(row["target_extension_names"]),
        int(row["target_extension_count"]),
    )


def verify_manifest_matches_live(
    manifest_rows: list[dict[str, Any]], state: dict[str, Any]
) -> None:
    manifest = {
        row["url_norm"]: manifest_mapping_signature(row) for row in manifest_rows
    }
    target_rows = state.get("live", {}).get("targets") or []
    live = {row["url_norm"]: live_mapping_signature(row) for row in target_rows}
    missing = sorted(set(live) - set(manifest))
    extra = sorted(set(manifest) - set(live))
    mismatched = sorted(
        url for url in set(manifest) & set(live) if manifest[url] != live[url]
    )
    if missing or extra or mismatched:
        raise PreconditionError(
            "manifest does not exactly match live universe targets: "
            f"missing={missing[:5]}, extra={extra[:5]}, "
            f"mapping_mismatch={mismatched[:5]}"
        )


def verify_snapshot_preflight(state: dict[str, Any]) -> None:
    backup = state.get("protected", {}).get("backup") or {}
    if (
        backup.get("rows") != EXPECTED_BACKUP_ROWS
        or backup.get("fingerprint") != EXPECTED_BACKUP_FINGERPRINT
        or backup.get("owner") != EXPECTED_BACKUP_OWNER
    ):
        raise PreconditionError(
            "immutable backup does not match the audited baseline: "
            + json.dumps(backup, sort_keys=True)
        )


def tool_metadata() -> dict[str, Any]:
    path = pathlib.Path(__file__).resolve()
    return {"name": TOOL_NAME, "path": str(path), "sha256": sha256_file(path)}


def make_snapshot(
    manifest_rows: list[dict[str, Any]], manifest_meta: dict[str, Any]
) -> dict[str, Any]:
    identity = database_identity()
    state = collect_state(identity)
    verify_manifest_matches_live(manifest_rows, state)
    verify_snapshot_preflight(state)
    snapshot = {
        "schema_version": SCHEMA_VERSION,
        "kind": "pgext-gh-repo-phase-a-before",
        "created_at": utc_now(),
        "tool": tool_metadata(),
        "manifest": manifest_meta,
        "state": state,
    }
    snapshot["integrity_sha256"] = digest_json(snapshot)
    return snapshot


def parse_timestamp(value: Any) -> datetime | None:
    if not value or not isinstance(value, str):
        return None
    try:
        return datetime.fromisoformat(value.replace("Z", "+00:00")).astimezone(
            timezone.utc
        )
    except ValueError:
        return None


def timestamp_advanced(before: Any, after: Any) -> bool:
    after_dt = parse_timestamp(after)
    if after_dt is None:
        return False
    before_dt = parse_timestamp(before)
    return before_dt is None or after_dt > before_dt


def same_timestamp(left: Any, right: Any) -> bool:
    if left is None and right is None:
        return True
    left_dt = parse_timestamp(left)
    right_dt = parse_timestamp(right)
    return left_dt is not None and right_dt is not None and left_dt == right_dt


def warning_rows(row: dict[str, Any]) -> list[dict[str, Any]]:
    warnings = row.get("activity_warnings")
    if not isinstance(warnings, list):
        return []
    return [warning for warning in warnings if isinstance(warning, dict)]


def is_empty_repository(row: dict[str, Any]) -> bool:
    return any(
        warning.get("kind") in {"empty_or_unborn_repository", "missing_default_branch"}
        for warning in warning_rows(row)
    )


def row_status_line(row: dict[str, Any]) -> dict[str, Any]:
    return {
        "url_norm": row.get("url_norm"),
        "status": row.get("status") or "missing",
        "http_status": row.get("http_status"),
        "attempts": row.get("attempts"),
        "error": row.get("error"),
    }


def metadata_missing_fields(row: dict[str, Any]) -> list[str]:
    missing: list[str] = []
    always_required = {
        "api_json": row.get("api_json_present"),
        "activity_json": row.get("activity_json_present"),
        "etag": row.get("etag"),
        "rate_limit_remaining": row.get("rate_limit_remaining"),
        "rate_limit_reset": row.get("rate_limit_reset"),
        "stargazers_count": row.get("stargazers_count"),
        "forks_count": row.get("forks_count"),
        "watchers_count": row.get("watchers_count"),
        "subscribers_count": row.get("subscribers_count"),
        "updated_at_api": row.get("updated_at_api"),
        "archived": row.get("archived"),
        "fetched_at": row.get("fetched_at"),
    }
    for field, value in always_required.items():
        if (
            value is None
            or value == ""
            or value is False
            and field in {"api_json", "activity_json"}
        ):
            missing.append(field)
    remaining = row.get("rate_limit_remaining")
    if type(remaining) is not int or remaining < RATE_LIMIT_RESERVE:
        if "rate_limit_remaining" not in missing:
            missing.append("rate_limit_remaining")
    if parse_timestamp(row.get("rate_limit_reset")) is None:
        if "rate_limit_reset" not in missing:
            missing.append("rate_limit_reset")
    if not is_empty_repository(row):
        for field in ("pushed_at", "default_branch"):
            if row.get(field) in (None, ""):
                missing.append(field)
    return missing


def scalar_equal(left: Any, right: Any) -> bool:
    return left == right and type(left) is type(right)


def metadata_api_mismatches(row: dict[str, Any]) -> list[str]:
    if not row.get("api_json_present"):
        return []
    mismatches: list[str] = []
    pairs = (
        ("stargazers_count", "api_stargazers_count"),
        ("forks_count", "api_forks_count"),
        ("watchers_count", "api_watchers_count"),
        ("subscribers_count", "api_subscribers_count"),
        ("default_branch", "api_default_branch"),
        ("archived", "api_archived"),
    )
    for stored, raw in pairs:
        if not scalar_equal(row.get(stored), row.get(raw)):
            mismatches.append(stored)
    for stored, raw in (
        ("pushed_at", "api_pushed_at"),
        ("updated_at_api", "api_updated_at"),
    ):
        if not same_timestamp(row.get(stored), row.get(raw)):
            mismatches.append(stored)
    return mismatches


def expected_latest(items: Iterable[tuple[str, Any]]) -> tuple[str | None, Any]:
    latest_source: str | None = None
    latest_value: Any = None
    latest_dt: datetime | None = None
    for source, value in items:
        parsed = parse_timestamp(value)
        if parsed is not None and (latest_dt is None or parsed > latest_dt):
            latest_source, latest_value, latest_dt = source, value, parsed
    return latest_source, latest_value


def activity_coherence_errors(row: dict[str, Any]) -> list[str]:
    errors: list[str] = []
    aliases = (
        ("last_commit_at", "last_commit_date"),
        ("last_release_at", "last_release_date"),
        ("latest_tag_at", "last_tag_date"),
        ("last_release_or_tag_at", "last_release_or_tag_date"),
        ("last_update_at", "last_update_date"),
    )
    for source, alias in aliases:
        if not same_timestamp(row.get(source), row.get(alias)):
            errors.append(f"{alias}!={source}")

    if row.get("last_commit_at") is not None:
        if not row.get("last_commit_sha"):
            errors.append("last_commit_sha missing")
        if not row.get("last_commit_html_url"):
            errors.append("last_commit_html_url missing")
    elif (
        row.get("last_commit_sha") is not None
        or row.get("last_commit_html_url") is not None
    ):
        errors.append("commit detail exists without last_commit_at")

    if row.get("last_release_at") is not None:
        expected_release = row.get("latest_release_published_at") or row.get(
            "latest_release_created_at"
        )
        if not row.get("latest_release_tag"):
            errors.append("latest_release_tag missing")
        if not same_timestamp(expected_release, row.get("last_release_at")):
            errors.append("last_release_at is not published_at/created_at")
    elif any(
        row.get(field) is not None
        for field in (
            "latest_release_tag",
            "latest_release_published_at",
            "latest_release_created_at",
        )
    ):
        errors.append("release detail exists without last_release_at")

    if row.get("latest_tag_at") is not None:
        if not row.get("latest_tag_name"):
            errors.append("latest_tag_name missing")
        if row.get("latest_tag_date_source") not in {
            "tagger.date",
            "commit.committedDate",
            "underlying_commit.committedDate",
        }:
            errors.append("latest_tag_date_source invalid")
        if row.get("latest_tag_target_type") not in {"Tag", "Commit"}:
            errors.append("latest_tag_target_type invalid")
        if row.get("latest_tag_target_type") == "Commit":
            if not row.get("latest_tag_commit_sha"):
                errors.append("lightweight tag commit SHA missing")
            if not same_timestamp(
                row.get("latest_tag_at"), row.get("latest_tag_commit_at")
            ):
                errors.append("lightweight tag/commit date mismatch")
    elif any(
        row.get(field) is not None
        for field in (
            "latest_tag_name",
            "latest_tag_date_source",
            "latest_tag_target_type",
            "latest_tag_commit_sha",
            "latest_tag_commit_at",
        )
    ):
        errors.append("tag detail exists without latest_tag_at")

    release_source, release_value = expected_latest(
        (("release", row.get("last_release_at")), ("tag", row.get("latest_tag_at")))
    )
    if not same_timestamp(release_value, row.get("last_release_or_tag_at")):
        errors.append("last_release_or_tag_at is not max(release,tag)")
    if release_source != row.get("last_release_or_tag_source"):
        errors.append("last_release_or_tag_source mismatch")

    update_source, update_value = expected_latest(
        (
            ("commit", row.get("last_commit_at")),
            ("release", row.get("last_release_at")),
            ("tag", row.get("latest_tag_at")),
        )
    )
    if not same_timestamp(update_value, row.get("last_update_at")):
        errors.append("last_update_at is not max(commit,release,tag)")
    if update_source != row.get("last_update_source"):
        errors.append("last_update_source mismatch")
    return errors


def response_envelope_errors(
    evidence: Any,
    label: str,
    allowed_statuses: set[int],
) -> list[str]:
    if not isinstance(evidence, dict):
        return [f"{label} evidence missing"]
    errors: list[str] = []
    status = evidence.get("status")
    attempts = evidence.get("request_attempts")
    if type(status) is not int or status not in allowed_statuses:
        errors.append(f"{label} status invalid")
    if type(attempts) is not int or attempts < 1:
        errors.append(f"{label} request_attempts invalid")
    for field in ("request_url", "final_url"):
        if not isinstance(evidence.get(field), str) or not evidence[field]:
            errors.append(f"{label} {field} missing")
    return errors


def same_github_api_origin(value: Any) -> bool:
    if not isinstance(value, str):
        return False
    parsed = urllib.parse.urlparse(value)
    return (
        parsed.scheme == "https"
        and parsed.hostname == "api.github.com"
        and parsed.port is None
        and parsed.username is None
        and parsed.password is None
        and not parsed.fragment
    )


def commit_request_matches(value: Any, owner: str, repo: str, branch: str) -> bool:
    if not same_github_api_origin(value):
        return False
    parsed = urllib.parse.urlparse(value)
    query = urllib.parse.parse_qs(parsed.query, keep_blank_values=True)
    return (
        parsed.path.lower() == f"/repos/{owner}/{repo}/commits".lower()
        and query.get("per_page") == ["1"]
        and query.get("sha") == [branch]
    )


def activity_evidence_errors(row: dict[str, Any]) -> list[str]:
    """Validate the raw request evidence retained for every fetched row."""

    errors: list[str] = []
    url = row["url_norm"]
    target_owner = row["target_repo_owner"]
    target_repo = row["target_repo_name"]
    target_api = row["target_api_url"]
    stored_owner = str(row.get("repo_owner") or "")
    stored_repo = str(row.get("repo_name") or "")
    stored_url = f"https://github.com/{stored_owner.lower()}/{stored_repo.lower()}"
    stored_api = str(row.get("api_url") or "")

    collected_at = parse_timestamp(row.get("activity_collected_at"))
    fetched_at = parse_timestamp(row.get("fetched_at"))
    if collected_at is None:
        errors.append("activity collected_at missing or invalid")
    elif fetched_at is None or collected_at > fetched_at:
        errors.append("activity collected_at is after fetched_at")

    requested = row.get("requested_evidence")
    expected_requested = {
        "url_norm": url,
        "repo_owner": target_owner,
        "repo_name": target_repo,
        "api_url": target_api,
    }
    if not isinstance(requested, dict):
        errors.append("requested evidence missing")
    else:
        for field, expected in expected_requested.items():
            if requested.get(field) != expected:
                errors.append(f"requested.{field} mismatch")

    repo_evidence = row.get("repo_evidence")
    errors.extend(response_envelope_errors(repo_evidence, "repo", {200}))
    if isinstance(repo_evidence, dict):
        if repo_evidence.get("request_url") != target_api:
            errors.append("repo request_url mismatch")
        if not same_github_api_origin(repo_evidence.get("final_url")):
            errors.append("repo final_url is not same-origin GitHub API")
    if row.get("repo_evidence_matches_api_json") is not True:
        errors.append("repo evidence data differs from api_json")

    api_full_name = row.get("api_full_name")
    api_html_url = row.get("api_html_url")
    if not isinstance(api_full_name, str) or not api_full_name:
        errors.append("api_json.full_name missing")
    if not isinstance(api_html_url, str) or not api_html_url:
        errors.append("api_json.html_url missing")

    redirect_detected = row.get("redirect_detected")
    if type(redirect_detected) is not bool:
        errors.append("redirect.detected missing or invalid")
    redirect_fields = {
        "canonical_url": row.get("redirect_canonical_url"),
        "full_name": row.get("redirect_full_name"),
        "html_url": row.get("redirect_html_url"),
        "final_api_url": row.get("redirect_final_api_url"),
    }
    for field, value in redirect_fields.items():
        if not isinstance(value, str) or not value:
            errors.append(f"redirect.{field} missing")
    if row.get("redirect_canonical_url") != stored_url:
        errors.append("redirect.canonical_url mismatch")
    if isinstance(api_full_name, str) and row.get("redirect_full_name") != api_full_name:
        errors.append("redirect.full_name differs from api_json")
    if isinstance(api_html_url, str) and str(row.get("redirect_html_url") or "").lower().rstrip("/") != api_html_url.lower().rstrip("/"):
        errors.append("redirect.html_url differs from api_json")
    if isinstance(repo_evidence, dict) and row.get(
        "redirect_final_api_url"
    ) != repo_evidence.get("final_url"):
        errors.append("redirect.final_api_url differs from repo response")
    if not same_github_api_origin(row.get("redirect_final_api_url")):
        errors.append("redirect.final_api_url is not same-origin GitHub API")
    if redirect_detected is False and row.get("redirect_final_api_url") != stored_api:
        errors.append("non-redirect repo final URL differs from canonical api_url")

    warning_kinds = {
        warning.get("kind") for warning in warning_rows(row) if warning.get("kind")
    }
    commit_evidence = row.get("latest_commit_evidence")
    default_branch = row.get("default_branch")
    if default_branch:
        errors.extend(
            response_envelope_errors(commit_evidence, "latest_commit", {200, 404, 409})
        )
        if isinstance(commit_evidence, dict):
            if not commit_request_matches(
                commit_evidence.get("request_url"),
                stored_owner,
                stored_repo,
                default_branch,
            ):
                errors.append("latest_commit request_url mismatch")
            final_commit_url = commit_evidence.get("final_url")
            if not commit_request_matches(
                final_commit_url,
                stored_owner,
                stored_repo,
                default_branch,
            ):
                errors.append("latest_commit final_url mismatch")
            status = commit_evidence.get("status")
            payload = commit_evidence.get("data")
            if status == 200:
                if not isinstance(payload, list):
                    errors.append("latest_commit 200 data is not an array")
                elif not payload:
                    errors.append("latest_commit 200 array is empty without empty-repo warning")
                elif not isinstance(payload[0], dict):
                    errors.append("latest_commit first row is invalid")
                else:
                    raw_commit = payload[0]
                    nested = raw_commit.get("commit")
                    nested = nested if isinstance(nested, dict) else {}
                    committer = nested.get("committer")
                    author = nested.get("author")
                    committer = committer if isinstance(committer, dict) else {}
                    author = author if isinstance(author, dict) else {}
                    raw_date = committer.get("date") or author.get("date")
                    if not same_timestamp(raw_date, row.get("last_commit_at")):
                        errors.append("last_commit_at differs from raw evidence")
                    if raw_commit.get("sha") != row.get("last_commit_sha"):
                        errors.append("last_commit_sha differs from raw evidence")
                    if raw_commit.get("html_url") != row.get("last_commit_html_url"):
                        errors.append("last_commit_html_url differs from raw evidence")
            elif status in {404, 409}:
                if "empty_or_unborn_repository" not in warning_kinds:
                    errors.append("empty commit response lacks empty-repo warning")
                if row.get("last_commit_at") is not None:
                    errors.append("empty commit response has stored commit activity")
    else:
        if commit_evidence is not None:
            errors.append("latest_commit evidence exists without default_branch")
        if "missing_default_branch" not in warning_kinds:
            errors.append("missing default_branch warning absent")

    release_evidence = row.get("release_tag_evidence")
    errors.extend(response_envelope_errors(release_evidence, "release_tag", {200}))
    if isinstance(release_evidence, dict):
        if release_evidence.get("request_url") != GRAPHQL_API:
            errors.append("release_tag request_url mismatch")
        if release_evidence.get("final_url") != GRAPHQL_API:
            errors.append("release_tag final_url mismatch")
        payload = release_evidence.get("data")
        if isinstance(payload, dict) and payload.get("errors"):
            errors.append("release_tag GraphQL errors present")
        graphql_data = payload.get("data") if isinstance(payload, dict) else None
        repository = graphql_data.get("repository") if isinstance(graphql_data, dict) else None
        releases = repository.get("releases") if isinstance(repository, dict) else None
        refs = repository.get("refs") if isinstance(repository, dict) else None
        release_nodes = releases.get("nodes") if isinstance(releases, dict) else None
        tag_nodes = refs.get("nodes") if isinstance(refs, dict) else None
        if not isinstance(release_nodes, list) or not isinstance(tag_nodes, list):
            errors.append("release_tag GraphQL payload shape invalid")
        else:
            if not isinstance(repository.get("nameWithOwner"), str) or repository[
                "nameWithOwner"
            ].lower() != str(api_full_name or "").lower():
                errors.append("release_tag repository identity mismatch")
            if repository.get("isArchived") is not row.get("archived"):
                errors.append("release_tag archived state mismatch")
            release = release_nodes[0] if release_nodes else None
            if release is not None and not isinstance(release, dict):
                errors.append("release_tag release node invalid")
            else:
                raw_release_tag = release.get("tagName") if release else None
                raw_published = release.get("publishedAt") if release else None
                raw_created = release.get("createdAt") if release else None
                if release is not None and (
                    not isinstance(raw_release_tag, str)
                    or not isinstance(raw_published or raw_created, str)
                ):
                    errors.append("release_tag release node lacks name/date")
                if raw_release_tag != row.get("latest_release_tag"):
                    errors.append("latest_release_tag differs from raw evidence")
                if not same_timestamp(raw_published, row.get("latest_release_published_at")):
                    errors.append("latest_release_published_at differs from raw evidence")
                if not same_timestamp(raw_created, row.get("latest_release_created_at")):
                    errors.append("latest_release_created_at differs from raw evidence")

            tag_ref = tag_nodes[0] if tag_nodes else None
            raw_name = None
            raw_type = None
            raw_tag_at = None
            raw_source = None
            raw_commit_sha = None
            raw_commit_at = None
            if tag_ref is not None and not isinstance(tag_ref, dict):
                errors.append("release_tag tag node invalid")
            elif tag_ref is not None:
                raw_name = tag_ref.get("name")
                target = tag_ref.get("target")
                if not isinstance(raw_name, str) or not isinstance(target, dict):
                    errors.append("release_tag tag node lacks name/target")
                target = target if isinstance(target, dict) else {}
                raw_type = target.get("__typename")
                if raw_type == "Tag":
                    tagger = target.get("tagger")
                    underlying = target.get("target")
                    tagger = tagger if isinstance(tagger, dict) else {}
                    underlying = underlying if isinstance(underlying, dict) else {}
                    raw_tag_at = tagger.get("date")
                    raw_source = "tagger.date" if raw_tag_at else None
                    if underlying.get("__typename") == "Commit":
                        raw_commit_sha = underlying.get("oid")
                        raw_commit_at = underlying.get("committedDate")
                        if not raw_tag_at:
                            raw_tag_at = raw_commit_at
                            raw_source = "underlying_commit.committedDate"
                elif raw_type == "Commit":
                    raw_tag_at = target.get("committedDate")
                    raw_source = "commit.committedDate" if raw_tag_at else None
                    raw_commit_sha = target.get("oid")
                    raw_commit_at = target.get("committedDate")
            raw_pairs = (
                ("latest_tag_name", raw_name),
                ("latest_tag_target_type", raw_type),
                ("latest_tag_date_source", raw_source),
                ("latest_tag_commit_sha", raw_commit_sha),
            )
            for field, expected in raw_pairs:
                if row.get(field) != expected:
                    errors.append(f"{field} differs from raw evidence")
            if not same_timestamp(raw_tag_at, row.get("latest_tag_at")):
                errors.append("latest_tag_at differs from raw evidence")
            if not same_timestamp(raw_commit_at, row.get("latest_tag_commit_at")):
                errors.append("latest_tag_commit_at differs from raw evidence")
    return errors


def target_identity_errors(row: dict[str, Any]) -> list[str]:
    if not row.get("row_exists"):
        return []
    errors: list[str] = []
    url = row["url_norm"]
    expected_owner, expected_repo = url.removeprefix("https://github.com/").split(
        "/", 1
    )
    stored_owner = str(row.get("repo_owner") or "").lower()
    stored_repo = str(row.get("repo_name") or "").lower()
    stored_canonical = f"https://github.com/{stored_owner}/{stored_repo}"
    api_url = str(row.get("api_url") or "").lower().rstrip("/")
    expected_api = f"https://api.github.com/repos/{stored_owner}/{stored_repo}"
    if row.get("repo_host") != "github.com":
        errors.append("repo_host")
    if row.get("redirect_detected"):
        if row.get("redirect_canonical_url") != stored_canonical:
            errors.append("redirect canonical/repo identity mismatch")
        if stored_canonical == url:
            errors.append("redirect marked but canonical key is unchanged")
        redirect_full_name = row.get("redirect_full_name")
        if (
            redirect_full_name
            and redirect_full_name.lower() != f"{stored_owner}/{stored_repo}"
        ):
            errors.append("redirect full_name/repo identity mismatch")
        redirect_html_url = str(row.get("redirect_html_url") or "").lower().rstrip("/")
        if redirect_html_url and redirect_html_url != stored_canonical:
            errors.append("redirect html_url/repo identity mismatch")
    elif stored_owner != expected_owner or stored_repo != expected_repo:
        errors.append("repo identity/key mismatch without redirect evidence")
    if api_url != expected_api:
        errors.append("api_url/repo identity mismatch")
    full_name = row.get("api_full_name")
    if full_name and full_name.lower() != f"{stored_owner}/{stored_repo}":
        errors.append("api_json.full_name/repo identity mismatch")
    html_url = str(row.get("api_html_url") or "").lower().rstrip("/")
    if html_url and html_url != stored_canonical:
        errors.append("api_json.html_url/repo identity mismatch")
    return errors


def add_gate(gates: list[dict[str, Any]], name: str, passed: bool, detail: str) -> None:
    gates.append({"name": name, "passed": bool(passed), "detail": detail})


def protected_equal(before: dict[str, Any], after: dict[str, Any], key: str) -> bool:
    return (before.get("protected", {}).get(key) or {}) == (
        after.get("protected", {}).get(key) or {}
    )


def evaluate_report(
    snapshot: dict[str, Any],
    after: dict[str, Any],
    manifest_rows: list[dict[str, Any]],
    manifest_meta: dict[str, Any],
) -> dict[str, Any]:
    before = snapshot["state"]
    before_rows = {
        row["url_norm"]: row for row in before.get("live", {}).get("targets", [])
    }
    after_rows = {
        row["url_norm"]: row for row in after.get("live", {}).get("targets", [])
    }
    manifest_by_url = {row["url_norm"]: row for row in manifest_rows}
    urls = sorted(manifest_by_url, key=lambda item: item.encode("utf-8"))

    inserted: list[str] = []
    attempted: list[str] = []
    success_refreshed: list[str] = []
    unattempted: list[str] = []
    missing_after: list[str] = []
    updated_at_not_advanced: list[str] = []
    preexisting_created_at_changed: list[str] = []
    failure_payload_changed: list[str] = []
    failure_fetched_at_changed: list[str] = []
    attempt_deltas: list[dict[str, Any]] = []

    for url in urls:
        pre = before_rows[url]
        post = after_rows[url]
        pre_exists = bool(pre.get("row_exists"))
        post_exists = bool(post.get("row_exists"))
        if not pre_exists and post_exists:
            inserted.append(url)
        if not post_exists:
            missing_after.append(url)
        if pre_exists and post_exists and not same_timestamp(
            pre.get("created_at"), post.get("created_at")
        ):
            preexisting_created_at_changed.append(url)
        pre_attempts = pre.get("attempts") if pre_exists else 0
        post_attempts = post.get("attempts") if post_exists else 0
        pre_attempts = pre_attempts if isinstance(pre_attempts, int) else 0
        post_attempts = post_attempts if isinstance(post_attempts, int) else 0
        delta = post_attempts - pre_attempts
        attempt_deltas.append(
            {
                "url_norm": url,
                "before": pre_attempts,
                "after": post_attempts,
                "delta": delta,
            }
        )
        was_attempted = post_exists and delta > 0
        if was_attempted:
            attempted.append(url)
            if not timestamp_advanced(pre.get("updated_at"), post.get("updated_at")):
                updated_at_not_advanced.append(url)
        else:
            unattempted.append(url)
        fetched_advanced = timestamp_advanced(
            pre.get("fetched_at"), post.get("fetched_at")
        )
        if (
            was_attempted
            and post.get("status") == "fetched"
            and post.get("http_status") == 200
            and post.get("error") is None
            and fetched_advanced
        ):
            success_refreshed.append(url)
        if post_exists and post.get("status") != "fetched":
            if pre_exists and pre.get("success_data_fingerprint") != post.get(
                "success_data_fingerprint"
            ):
                failure_payload_changed.append(url)
            if pre_exists and not same_timestamp(
                pre.get("fetched_at"), post.get("fetched_at")
            ):
                failure_fetched_at_changed.append(url)

    target_status: dict[str, int] = {}
    status_lists: dict[str, list[dict[str, Any]]] = {
        key: [] for key in STATUS_LIST_ORDER
    }
    for url in urls:
        row = after_rows[url]
        status = row.get("status") if row.get("row_exists") else "missing"
        status = status or "missing"
        target_status[status] = target_status.get(status, 0) + 1
        if status in status_lists:
            status_lists[status].append(row_status_line(row))

    success_refreshed_set = set(success_refreshed)
    redirects = [
        {
            "url_norm": row["url_norm"],
            "canonical_url": row.get("redirect_canonical_url"),
            "full_name": row.get("redirect_full_name"),
            "html_url": row.get("redirect_html_url"),
            "final_api_url": row.get("redirect_final_api_url"),
        }
        for row in after_rows.values()
        if row.get("url_norm") in success_refreshed_set and row.get("redirect_detected")
    ]
    redirects.sort(key=lambda row: row["url_norm"].encode("utf-8"))
    empty_repositories = [
        {
            "url_norm": row["url_norm"],
            "warnings": warning_rows(row),
        }
        for row in after_rows.values()
        if row.get("url_norm") in success_refreshed_set and is_empty_repository(row)
    ]
    empty_repositories.sort(key=lambda row: row["url_norm"].encode("utf-8"))

    fetched_rows = [
        row for row in after_rows.values() if row.get("status") == "fetched"
    ]
    metadata_coverage: dict[str, dict[str, int]] = {}
    for field in METADATA_COVERAGE_FIELDS:
        if field == "api_json":
            getter = lambda row: row.get("api_json_present") is True
        elif field == "activity_json":
            getter = lambda row: row.get("activity_json_present") is True
        else:
            getter = lambda row, field=field: row.get(field) is not None
        metadata_coverage[field] = {
            "target_count": sum(1 for row in after_rows.values() if getter(row)),
            "fetched_count": sum(1 for row in fetched_rows if getter(row)),
            "fetched_denominator": len(fetched_rows),
        }

    metadata_incomplete = []
    metadata_api_mismatch = []
    activity_incoherent = []
    activity_evidence_invalid = []
    activity_evidence_outside_run = []
    commit_missing_not_empty = []
    update_missing_not_empty = []
    for row in fetched_rows:
        missing_fields = metadata_missing_fields(row)
        if missing_fields:
            metadata_incomplete.append(
                {"url_norm": row["url_norm"], "fields": missing_fields}
            )
        mismatch_fields = metadata_api_mismatches(row)
        if mismatch_fields:
            metadata_api_mismatch.append(
                {"url_norm": row["url_norm"], "fields": mismatch_fields}
            )
        coherence = activity_coherence_errors(row)
        if coherence:
            activity_incoherent.append(
                {"url_norm": row["url_norm"], "errors": coherence}
            )
        evidence_errors = activity_evidence_errors(row)
        if evidence_errors:
            activity_evidence_invalid.append(
                {"url_norm": row["url_norm"], "errors": evidence_errors}
            )
        if row.get("last_commit_at") is None and not is_empty_repository(row):
            commit_missing_not_empty.append(row["url_norm"])
        if row.get("last_update_at") is None and not is_empty_repository(row):
            update_missing_not_empty.append(row["url_norm"])

    activity_coverage = {
        "commit": sum(row.get("last_commit_at") is not None for row in fetched_rows),
        "release": sum(row.get("last_release_at") is not None for row in fetched_rows),
        "tag": sum(row.get("latest_tag_at") is not None for row in fetched_rows),
        "release_or_tag": sum(
            row.get("last_release_or_tag_at") is not None for row in fetched_rows
        ),
        "update": sum(row.get("last_update_at") is not None for row in fetched_rows),
        "fetched_denominator": len(fetched_rows),
    }

    mapping_mismatches = []
    target_identity_anomalies = []
    for url in urls:
        row = after_rows[url]
        manifest = manifest_by_url[url]
        if row.get("row_exists") and (
            row.get("extension_ids") != manifest["extension_ids"]
            or row.get("extension_names") != manifest["extension_names"]
            or row.get("extension_count") != manifest["extension_count"]
        ):
            mapping_mismatches.append(
                {
                    "url_norm": url,
                    "expected_ids": manifest["extension_ids"],
                    "actual_ids": row.get("extension_ids"),
                    "expected_names": manifest["extension_names"],
                    "actual_names": row.get("extension_names"),
                    "expected_count": manifest["extension_count"],
                    "actual_count": row.get("extension_count"),
                }
            )
        identity_errors = target_identity_errors(row)
        if identity_errors:
            target_identity_anomalies.append(
                {"url_norm": url, "errors": identity_errors}
            )

    prior_candidates = []
    for url in urls:
        manifest = manifest_by_url[url]
        if not before_rows[url].get("row_exists") and manifest.get(
            "prior_url_candidates"
        ):
            prior_candidates.append(
                {
                    "url_norm": url,
                    "extension_names": manifest["extension_names"],
                    "prior_url_candidates": manifest["prior_url_candidates"],
                    "final_status": after_rows[url].get("status") or "missing",
                    "final_http_status": after_rows[url].get("http_status"),
                    "redirect_canonical_url": after_rows[url].get(
                        "redirect_canonical_url"
                    ),
                }
            )

    before_live = before["live"]
    after_live = after["live"]
    initial_missing = sorted(
        (url for url in urls if not before_rows[url].get("row_exists")),
        key=lambda item: item.encode("utf-8"),
    )
    fetched_invalid = [
        row_status_line(row)
        for row in fetched_rows
        if row.get("http_status") != 200 or row.get("error") is not None
    ]
    blocked_invalid = [
        row_status_line(row)
        for row in after_rows.values()
        if row.get("status") == "blocked"
        and (row.get("http_status") != 404 or not row.get("error"))
    ]
    rate_limit_invalid = [
        {
            **row_status_line(row),
            "rate_limit_remaining": row.get("rate_limit_remaining"),
            "rate_limit_reset": row.get("rate_limit_reset"),
        }
        for row in after_rows.values()
        if row.get("row_exists")
        and (
            type(row.get("rate_limit_remaining")) is not int
            or row.get("rate_limit_remaining") < RATE_LIMIT_RESERVE
            or parse_timestamp(row.get("rate_limit_reset")) is None
        )
    ]

    before_identity = before.get("database") or {}
    after_identity = after.get("database") or {}
    identity_keys = (
        "database",
        "session_user",
        "current_user",
        "server_addr",
        "port",
        "cluster_system_identifier",
    )
    same_identity = all(
        before_identity.get(key) == after_identity.get(key) for key in identity_keys
    )
    before_capture = parse_timestamp(before.get("captured_at"))
    after_capture = parse_timestamp(after.get("captured_at"))
    chronological = (
        before_capture is not None
        and after_capture is not None
        and after_capture > before_capture
    )
    if before_capture is not None and after_capture is not None:
        for row in fetched_rows:
            collected = parse_timestamp(row.get("activity_collected_at"))
            if (
                collected is None
                or collected <= before_capture
                or collected > after_capture
            ):
                activity_evidence_outside_run.append(row["url_norm"])

    gates: list[dict[str, Any]] = []
    add_gate(
        gates,
        "same_database_cluster",
        same_identity,
        "fixed endpoint and cluster system identifier match snapshot",
    )
    add_gate(
        gates,
        "post_capture_after_snapshot",
        chronological,
        f"before={before.get('captured_at')} after={after.get('captured_at')}",
    )
    add_gate(
        gates,
        "target_count",
        len(after_rows) == len(urls),
        f"live={len(after_rows)} manifest={len(urls)}",
    )
    add_gate(
        gates, "target_coverage", not missing_after, f"missing={len(missing_after)}"
    )
    add_gate(
        gates,
        "initial_missing_inserted",
        set(inserted) == set(initial_missing),
        f"initial_missing={len(initial_missing)} inserted={len(inserted)}",
    )
    add_gate(
        gates,
        "all_targets_attempted",
        not unattempted,
        f"attempted={len(attempted)} target={len(urls)}",
    )
    add_gate(
        gates,
        "attempt_updates_updated_at",
        not updated_at_not_advanced,
        f"anomalies={len(updated_at_not_advanced)}",
    )
    add_gate(
        gates,
        "preexisting_created_at_unchanged",
        not preexisting_created_at_changed,
        f"anomalies={len(preexisting_created_at_changed)}",
    )
    transient_count = sum(
        target_status.get(status, 0) for status in TRANSIENT_STATUSES
    ) + target_status.get("missing", 0)
    add_gate(
        gates,
        "terminal_statuses",
        transient_count == 0 and set(target_status) <= TERMINAL_STATUSES,
        json.dumps(target_status, sort_keys=True),
    )
    add_gate(
        gates,
        "fetched_status_fields",
        not fetched_invalid,
        f"anomalies={len(fetched_invalid)}",
    )
    add_gate(
        gates,
        "blocked_status_fields",
        not blocked_invalid,
        f"anomalies={len(blocked_invalid)}",
    )
    add_gate(
        gates,
        "all_targets_preserve_rate_limit_reserve",
        not rate_limit_invalid,
        f"reserve={RATE_LIMIT_RESERVE} anomalies={len(rate_limit_invalid)}",
    )
    fetched_urls = {row["url_norm"] for row in fetched_rows}
    add_gate(
        gates,
        "successful_refresh_evidence",
        fetched_urls == set(success_refreshed),
        f"fetched={len(fetched_urls)} success_refreshed={len(success_refreshed)}",
    )
    add_gate(
        gates,
        "failure_preserves_success_payload",
        not failure_payload_changed,
        f"anomalies={len(failure_payload_changed)}",
    )
    add_gate(
        gates,
        "failure_does_not_advance_fetched_at",
        not failure_fetched_at_changed,
        f"anomalies={len(failure_fetched_at_changed)}",
    )
    add_gate(
        gates,
        "metadata_complete",
        not metadata_incomplete,
        f"fetched_metadata_anomalies={len(metadata_incomplete)}",
    )
    add_gate(
        gates,
        "metadata_matches_api_json",
        not metadata_api_mismatch,
        f"mismatches={len(metadata_api_mismatch)}",
    )
    add_gate(
        gates,
        "commit_or_empty_repository",
        not commit_missing_not_empty,
        f"unexplained_missing_commit={len(commit_missing_not_empty)}",
    )
    add_gate(
        gates,
        "update_or_empty_repository",
        not update_missing_not_empty,
        f"unexplained_missing_update={len(update_missing_not_empty)}",
    )
    add_gate(
        gates,
        "activity_field_coherence",
        not activity_incoherent,
        f"anomalies={len(activity_incoherent)}",
    )
    add_gate(
        gates,
        "activity_raw_evidence",
        not activity_evidence_invalid,
        f"anomalies={len(activity_evidence_invalid)}",
    )
    add_gate(
        gates,
        "activity_evidence_within_run",
        not activity_evidence_outside_run,
        f"anomalies={len(activity_evidence_outside_run)}",
    )
    add_gate(
        gates,
        "target_mapping",
        not mapping_mismatches,
        f"mismatches={len(mapping_mismatches)}",
    )
    add_gate(
        gates,
        "target_key_identity",
        not target_identity_anomalies,
        f"anomalies={len(target_identity_anomalies)}",
    )
    integrity = after_live.get("integrity") or {}
    add_gate(
        gates,
        "no_duplicate_url_norm",
        not integrity.get("duplicate_url_norms"),
        f"anomalies={len(integrity.get('duplicate_url_norms') or [])}",
    )
    add_gate(
        gates,
        "canonical_url_keys",
        not integrity.get("invalid_url_keys"),
        f"anomalies={len(integrity.get('invalid_url_keys') or [])}",
    )
    add_gate(
        gates,
        "array_integrity",
        not integrity.get("array_anomalies"),
        f"anomalies={len(integrity.get('array_anomalies') or [])}",
    )
    expected_live_rows = before_live.get("rows", 0) + len(inserted)
    add_gate(
        gates,
        "append_only_live_row_count",
        after_live.get("rows") == expected_live_rows,
        f"before={before_live.get('rows')} inserted={len(inserted)} after={after_live.get('rows')}",
    )
    for protected_key, gate_name in (
        ("non_target", "non_target_full_content_unchanged"),
        ("backup", "backup_unchanged"),
        ("universe", "universe_unchanged"),
        ("extension", "extension_unchanged"),
    ):
        equal = protected_equal(before, after, protected_key)
        add_gate(
            gates,
            gate_name,
            equal,
            f"before={before['protected'][protected_key]} after={after['protected'][protected_key]}",
        )
    backup_after = after["protected"]["backup"]
    backup_expected = (
        backup_after.get("rows") == EXPECTED_BACKUP_ROWS
        and backup_after.get("fingerprint") == EXPECTED_BACKUP_FINGERPRINT
        and backup_after.get("owner") == EXPECTED_BACKUP_OWNER
    )
    add_gate(
        gates,
        "backup_matches_audited_baseline",
        backup_expected,
        json.dumps(backup_after, sort_keys=True),
    )

    passed = all(gate["passed"] for gate in gates)
    return {
        "schema_version": SCHEMA_VERSION,
        "kind": "pgext-gh-repo-phase-a-report",
        "generated_at": utc_now(),
        "outcome": "PASS" if passed else "FAIL",
        "manifest": manifest_meta,
        "snapshot": {
            "created_at": snapshot.get("created_at"),
            "captured_at": before.get("captured_at"),
            "integrity_sha256": snapshot.get("integrity_sha256"),
        },
        "after_captured_at": after.get("captured_at"),
        "gates": gates,
        "counts": {
            "target": len(urls),
            "initial_existing": len(urls) - len(initial_missing),
            "initial_missing": len(initial_missing),
            "inserted": len(inserted),
            "attempted": len(attempted),
            "success_refreshed": len(success_refreshed),
            "redirect_or_rename": len(redirects),
            "empty_or_unborn": len(empty_repositories),
            "missing_after": len(missing_after),
        },
        "status": target_status,
        "metadata_coverage": metadata_coverage,
        "activity_coverage": activity_coverage,
        "protected": {
            key: {"before": before["protected"][key], "after": after["protected"][key]}
            for key in ("non_target", "backup", "universe", "extension")
        },
        "live_rows": {
            "before": before_live.get("rows"),
            "after": after_live.get("rows"),
        },
        "lists": {
            "initial_missing": initial_missing,
            "inserted": inserted,
            "attempted": attempted,
            "unattempted": unattempted,
            "success_refreshed": success_refreshed,
            "missing_after": missing_after,
            "redirects": redirects,
            "empty_repositories": empty_repositories,
            "prior_url_candidates_for_initial_missing": prior_candidates,
            "status_anomalies": status_lists,
            "metadata_incomplete": metadata_incomplete,
            "metadata_api_mismatch": metadata_api_mismatch,
            "activity_incoherent": activity_incoherent,
            "activity_evidence_invalid": activity_evidence_invalid,
            "activity_evidence_outside_run": activity_evidence_outside_run,
            "commit_missing_not_empty": commit_missing_not_empty,
            "update_missing_not_empty": update_missing_not_empty,
            "mapping_mismatches": mapping_mismatches,
            "target_identity_anomalies": target_identity_anomalies,
            "updated_at_not_advanced": updated_at_not_advanced,
            "preexisting_created_at_changed": preexisting_created_at_changed,
            "failure_payload_changed": failure_payload_changed,
            "failure_fetched_at_changed": failure_fetched_at_changed,
            "fetched_status_invalid": fetched_invalid,
            "blocked_status_invalid": blocked_invalid,
            "rate_limit_invalid": rate_limit_invalid,
            "duplicate_url_norms": integrity.get("duplicate_url_norms") or [],
            "invalid_url_keys": integrity.get("invalid_url_keys") or [],
            "array_anomalies": integrity.get("array_anomalies") or [],
            "attempt_deltas": attempt_deltas,
        },
    }


def load_snapshot(
    path_value: str,
    manifest_rows: list[dict[str, Any]],
    manifest_meta: dict[str, Any],
) -> tuple[pathlib.Path, dict[str, Any]]:
    path = pathlib.Path(path_value).expanduser().resolve()
    if not path.is_file():
        raise SnapshotError(f"snapshot does not exist: {path}")
    try:
        snapshot = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as exc:
        raise SnapshotError(f"invalid snapshot JSON: {path}") from exc
    if not isinstance(snapshot, dict):
        raise SnapshotError("snapshot root must be an object")
    supplied_digest = snapshot.get("integrity_sha256")
    unsigned = dict(snapshot)
    unsigned.pop("integrity_sha256", None)
    if supplied_digest != digest_json(unsigned):
        raise SnapshotError("snapshot integrity_sha256 mismatch")
    if (
        snapshot.get("schema_version") != SCHEMA_VERSION
        or snapshot.get("kind") != "pgext-gh-repo-phase-a-before"
    ):
        raise SnapshotError("snapshot schema/kind is not supported")
    snapshot_manifest = snapshot.get("manifest") or {}
    for key in ("sha256", "mapping_sha256", "target_count"):
        if snapshot_manifest.get(key) != manifest_meta.get(key):
            raise SnapshotError(f"snapshot belongs to a different manifest ({key})")
    state = snapshot.get("state")
    if not isinstance(state, dict):
        raise SnapshotError("snapshot lacks a database state object")
    target_rows = state.get("live", {}).get("targets") or []
    target_urls = [row.get("url_norm") for row in target_rows if isinstance(row, dict)]
    if (
        len(target_rows) != manifest_meta["target_count"]
        or len(target_urls) != len(target_rows)
        or len(set(target_urls)) != len(target_urls)
    ):
        raise SnapshotError("snapshot target states are incomplete or duplicated")
    try:
        verify_manifest_matches_live(manifest_rows, state)
    except PreconditionError as exc:
        raise SnapshotError(f"snapshot target mapping is invalid: {exc}") from exc
    verify_snapshot_preflight(state)
    return path, snapshot


def markdown_escape(value: Any) -> str:
    if value is None:
        return ""
    return str(value).replace("|", "\\|").replace("\n", " ")


def percent(count: int, denominator: int) -> str:
    return f"{100.0 * count / denominator:.1f}%" if denominator else "0.0%"


def render_simple_url_list(title: str, rows: list[Any]) -> list[str]:
    lines = [f"### {title}", "", f"Count: {len(rows)}", ""]
    if not rows:
        lines.append("None.")
        lines.append("")
        return lines
    lines.append("```text")
    for row in rows:
        lines.append(str(row if isinstance(row, str) else row.get("url_norm", row)))
    lines.extend(["```", ""])
    return lines


def render_markdown(report: dict[str, Any], snapshot_path: pathlib.Path) -> str:
    lines = [
        "# Phase A GitHub refresh acceptance report",
        "",
        f"**Outcome: {report['outcome']}**",
        "",
        f"- Generated: `{report['generated_at']}`",
        f"- Snapshot: `{snapshot_path}`",
        f"- Manifest: `{report['manifest']['path']}`",
        f"- Manifest SHA-256: `{report['manifest']['sha256']}`",
        f"- Target count: {report['counts']['target']}",
        "- Database: `postgres@/tmp:5432/data` (read-only acceptance queries)",
        "",
        "## Acceptance gates",
        "",
        "| Gate | Result | Detail |",
        "|---|---:|---|",
    ]
    for gate in report["gates"]:
        lines.append(
            f"| `{markdown_escape(gate['name'])}` | {'PASS' if gate['passed'] else 'FAIL'} | {markdown_escape(gate['detail'])} |"
        )

    counts = report["counts"]
    lines.extend(
        [
            "",
            "## Run accounting",
            "",
            "| Metric | Count |",
            "|---|---:|",
        ]
    )
    for key in (
        "target",
        "initial_existing",
        "initial_missing",
        "inserted",
        "attempted",
        "success_refreshed",
        "redirect_or_rename",
        "empty_or_unborn",
        "missing_after",
    ):
        lines.append(f"| `{key}` | {counts[key]} |")

    lines.extend(["", "## Target status", "", "| Status | Count |", "|---|---:|"])
    for status, count in sorted(report["status"].items()):
        lines.append(f"| `{markdown_escape(status)}` | {count} |")

    lines.extend(
        [
            "",
            "## Metadata coverage",
            "",
            "Release and tag fields are allowed to be empty. Metadata completeness gates apply only to fetched rows and allow `pushed_at`/`default_branch` to be empty when the activity evidence marks an empty or unborn repository.",
            "",
            "| Field | All targets | Fetched targets | Fetched coverage |",
            "|---|---:|---:|---:|",
        ]
    )
    for field, values in report["metadata_coverage"].items():
        lines.append(
            f"| `{field}` | {values['target_count']} | {values['fetched_count']} / {values['fetched_denominator']} | {percent(values['fetched_count'], values['fetched_denominator'])} |"
        )

    activity = report["activity_coverage"]
    lines.extend(
        [
            "",
            "## Activity coverage",
            "",
            "| Field group | Count | Coverage of fetched |",
            "|---|---:|---:|",
        ]
    )
    for key in ("commit", "release", "tag", "release_or_tag", "update"):
        lines.append(
            f"| `{key}` | {activity[key]} | {percent(activity[key], activity['fetched_denominator'])} |"
        )

    lines.extend(
        [
            "",
            "## Protected fingerprints",
            "",
            "| Object | Before rows | After rows | Before fingerprint | After fingerprint |",
            "|---|---:|---:|---|---|",
        ]
    )
    for key, values in report["protected"].items():
        before, after = values["before"], values["after"]
        lines.append(
            f"| `{key}` | {before.get('rows')} | {after.get('rows')} | `{before.get('fingerprint')}` | `{after.get('fingerprint')}` |"
        )

    lists = report["lists"]
    lines.extend(["", "## Status anomaly lines", ""])
    for status in STATUS_LIST_ORDER:
        rows = lists["status_anomalies"].get(status) or []
        lines.extend(
            [
                f"### {status}",
                "",
                f"Count: {len(rows)}",
                "",
                "| URL | HTTP | Attempts | Error |",
                "|---|---:|---:|---|",
            ]
        )
        if rows:
            for row in rows:
                lines.append(
                    f"| `{markdown_escape(row['url_norm'])}` | {markdown_escape(row.get('http_status'))} | {markdown_escape(row.get('attempts'))} | {markdown_escape(row.get('error'))} |"
                )
        else:
            lines.append("| _none_ |  |  |  |")
        lines.append("")

    lines.extend(["## Redirect / rename evidence", ""])
    redirects = lists["redirects"]
    lines.extend(
        [
            f"Count: {len(redirects)}",
            "",
            "| Requested key | Canonical URL | API full name | Final API URL |",
            "|---|---|---|---|",
        ]
    )
    if redirects:
        for row in redirects:
            lines.append(
                f"| `{markdown_escape(row['url_norm'])}` | `{markdown_escape(row.get('canonical_url'))}` | `{markdown_escape(row.get('full_name'))}` | `{markdown_escape(row.get('final_api_url'))}` |"
            )
    else:
        lines.append("| _none_ |  |  |  |")

    lines.extend(["", "## Empty / unborn repository evidence", ""])
    empty_rows = lists["empty_repositories"]
    lines.extend([f"Count: {len(empty_rows)}", "", "| URL | Warnings |", "|---|---|"])
    if empty_rows:
        for row in empty_rows:
            lines.append(
                f"| `{markdown_escape(row['url_norm'])}` | `{markdown_escape(json.dumps(row['warnings'], ensure_ascii=False, sort_keys=True))}` |"
            )
    else:
        lines.append("| _none_ |  |")

    lines.extend(["", "## Prior URL candidates for targets missing before the run", ""])
    candidates = lists["prior_url_candidates_for_initial_missing"]
    lines.extend(
        [
            f"Count: {len(candidates)}",
            "",
            "| Target | Extensions | Prior candidates | Final state | Redirect canonical |",
            "|---|---|---|---|---|",
        ]
    )
    if candidates:
        for row in candidates:
            final = f"{row.get('final_status')}/{row.get('final_http_status')}"
            lines.append(
                f"| `{markdown_escape(row['url_norm'])}` | `{markdown_escape(', '.join(row['extension_names']))}` | {markdown_escape(row['prior_url_candidates'])} | `{markdown_escape(final)}` | `{markdown_escape(row.get('redirect_canonical_url'))}` |"
            )
    else:
        lines.append("| _none_ |  |  |  |  |")

    lines.extend(["", "## Complete accounting lists", ""])
    for key, title in (
        ("initial_missing", "Initial missing targets"),
        ("inserted", "Inserted targets"),
        ("attempted", "Attempted targets"),
        ("success_refreshed", "Successfully refreshed targets"),
        ("unattempted", "Unattempted targets"),
        ("missing_after", "Targets missing after run"),
    ):
        lines.extend(render_simple_url_list(title, lists[key]))

    lines.extend(["## Additional QA anomaly lists", ""])
    anomaly_keys = (
        "metadata_incomplete",
        "metadata_api_mismatch",
        "activity_incoherent",
        "activity_evidence_invalid",
        "activity_evidence_outside_run",
        "commit_missing_not_empty",
        "update_missing_not_empty",
        "mapping_mismatches",
        "target_identity_anomalies",
        "updated_at_not_advanced",
        "preexisting_created_at_changed",
        "failure_payload_changed",
        "failure_fetched_at_changed",
        "fetched_status_invalid",
        "blocked_status_invalid",
        "rate_limit_invalid",
        "duplicate_url_norms",
        "invalid_url_keys",
        "array_anomalies",
    )
    for key in anomaly_keys:
        rows = lists[key]
        lines.extend([f"### {key}", "", f"Count: {len(rows)}", ""])
        if rows:
            lines.append("```json")
            lines.append(json.dumps(rows, ensure_ascii=False, sort_keys=True, indent=2))
            lines.append("```")
        else:
            lines.append("None.")
        lines.append("")
    return "\n".join(lines).rstrip() + "\n"


def atomic_write(path: pathlib.Path, content: str, overwrite: bool) -> None:
    path = path.expanduser().resolve()
    path.parent.mkdir(parents=True, exist_ok=True)
    if path.exists() and not overwrite:
        raise AcceptanceError(f"refusing to overwrite existing evidence file: {path}")
    descriptor, temporary_name = tempfile.mkstemp(
        prefix=f".{path.name}.", dir=path.parent
    )
    try:
        with os.fdopen(descriptor, "w", encoding="utf-8") as handle:
            handle.write(content)
            handle.flush()
            os.fsync(handle.fileno())
        os.replace(temporary_name, path)
    except BaseException:
        try:
            os.unlink(temporary_name)
        except FileNotFoundError:
            pass
        raise


def validate_output_paths(
    outputs: Iterable[pathlib.Path],
    protected_inputs: Iterable[pathlib.Path],
) -> list[pathlib.Path]:
    resolved_outputs = [path.expanduser().resolve() for path in outputs]
    resolved_inputs = {path.expanduser().resolve() for path in protected_inputs}
    if len(resolved_outputs) != len(set(resolved_outputs)):
        raise AcceptanceError("output paths must not alias each other")
    if set(resolved_outputs) & resolved_inputs:
        raise AcceptanceError("output paths must not alias an input artifact")
    return resolved_outputs


def write_json(path: pathlib.Path, payload: dict[str, Any], overwrite: bool) -> None:
    atomic_write(
        path,
        json.dumps(payload, ensure_ascii=False, sort_keys=True, indent=2) + "\n",
        overwrite,
    )


def command_snapshot(args: argparse.Namespace) -> int:
    manifest_path, rows, manifest_meta = load_manifest(args.manifest)
    output = validate_output_paths(
        [pathlib.Path(args.output)], [manifest_path]
    )[0]
    snapshot = make_snapshot(rows, manifest_meta)
    write_json(output, snapshot, args.overwrite)
    print(f"snapshot written: {output}")
    print(
        f"targets={manifest_meta['target_count']} initial_missing={snapshot['state']['live']['target_missing']}"
    )
    return 0


def command_report(args: argparse.Namespace) -> int:
    manifest_path, rows, manifest_meta = load_manifest(args.manifest)
    snapshot_path, snapshot = load_snapshot(args.snapshot, rows, manifest_meta)
    json_output, markdown_output = validate_output_paths(
        [pathlib.Path(args.json_output), pathlib.Path(args.markdown_output)],
        [manifest_path, snapshot_path],
    )
    identity = database_identity()
    after = collect_state(identity)
    verify_manifest_matches_live(rows, after)
    report = evaluate_report(snapshot, after, rows, manifest_meta)
    for path in (json_output, markdown_output):
        if path.exists() and not args.overwrite:
            raise AcceptanceError(
                f"refusing to overwrite existing evidence file: {path}"
            )
    write_json(json_output, report, args.overwrite)
    atomic_write(
        markdown_output, render_markdown(report, snapshot_path), args.overwrite
    )
    print(f"report outcome: {report['outcome']}")
    print(f"JSON report: {json_output}")
    print(f"Markdown report: {markdown_output}")
    failed = [gate["name"] for gate in report["gates"] if not gate["passed"]]
    if failed:
        print("failed gates: " + ", ".join(failed), file=sys.stderr)
        return 1
    return 0


def fixture_target(
    url: str,
    exists: bool,
    status: str | None,
    attempts: int | None,
    fetched_at: str | None,
) -> dict[str, Any]:
    owner, repo = url.removeprefix("https://github.com/").split("/", 1)
    timestamp = "2026-07-15T01:00:00+00:00"
    collected_at = "2026-07-15T02:30:00+00:00"
    api_url = f"https://api.github.com/repos/{owner}/{repo}"
    commit_url = f"{api_url}/commits?per_page=1&sha=main"
    row: dict[str, Any] = {
        "url_norm": url,
        "target_repo_owner": owner,
        "target_repo_name": repo,
        "target_api_url": api_url,
        "target_extension_ids": [1],
        "target_extension_names": [repo],
        "target_extension_count": 1,
        "row_exists": exists,
        "repo_host": "github.com" if exists else None,
        "repo_owner": owner if exists else None,
        "repo_name": repo if exists else None,
        "api_url": api_url if exists else None,
        "extension_ids": [1] if exists else None,
        "extension_names": [repo] if exists else None,
        "extension_count": 1 if exists else None,
        "status": status,
        "attempts": attempts,
        "http_status": (
            200 if status == "fetched" else (404 if status == "blocked" else None)
        ),
        "error": "repository not found" if status == "blocked" else None,
        "etag": '"etag"' if status == "fetched" else None,
        "rate_limit_remaining": (
            4998 if status in {"fetched", "blocked"} else None
        ),
        "rate_limit_reset": (
            "2026-07-15T05:00:00+00:00"
            if status in {"fetched", "blocked"}
            else None
        ),
        "stargazers_count": 1 if status == "fetched" else None,
        "forks_count": 2 if status == "fetched" else None,
        "watchers_count": 1 if status == "fetched" else None,
        "subscribers_count": 3 if status == "fetched" else None,
        "pushed_at": timestamp if status == "fetched" else None,
        "updated_at_api": timestamp if status == "fetched" else None,
        "default_branch": "main" if status == "fetched" else None,
        "archived": False if status == "fetched" else None,
        "fetched_at": fetched_at,
        "created_at": "2026-07-14T00:00:00+00:00" if exists else None,
        "updated_at": fetched_at,
        "last_commit_at": timestamp if status == "fetched" else None,
        "last_commit_date": timestamp if status == "fetched" else None,
        "last_commit_sha": "abc" if status == "fetched" else None,
        "last_commit_html_url": f"{url}/commit/abc" if status == "fetched" else None,
        "latest_release_tag": None,
        "latest_release_published_at": None,
        "latest_release_created_at": None,
        "last_release_at": None,
        "last_release_date": None,
        "latest_tag_name": None,
        "latest_tag_at": None,
        "last_tag_date": None,
        "latest_tag_date_source": None,
        "latest_tag_target_type": None,
        "latest_tag_commit_sha": None,
        "latest_tag_commit_at": None,
        "last_release_or_tag_at": None,
        "last_release_or_tag_date": None,
        "last_release_or_tag_source": None,
        "last_update_at": timestamp if status == "fetched" else None,
        "last_update_date": timestamp if status == "fetched" else None,
        "last_update_source": "commit" if status == "fetched" else None,
        "api_json_present": status == "fetched",
        "activity_json_present": status == "fetched",
        "api_stargazers_count": 1 if status == "fetched" else None,
        "api_forks_count": 2 if status == "fetched" else None,
        "api_watchers_count": 1 if status == "fetched" else None,
        "api_subscribers_count": 3 if status == "fetched" else None,
        "api_pushed_at": timestamp if status == "fetched" else None,
        "api_updated_at": timestamp if status == "fetched" else None,
        "api_default_branch": "main" if status == "fetched" else None,
        "api_archived": False if status == "fetched" else None,
        "api_full_name": f"{owner}/{repo}" if status == "fetched" else None,
        "api_html_url": url if status == "fetched" else None,
        "activity_collected_at": collected_at if status == "fetched" else None,
        "requested_evidence": (
            {
                "url_norm": url,
                "repo_owner": owner,
                "repo_name": repo,
                "api_url": api_url,
            }
            if status == "fetched"
            else None
        ),
        "repo_evidence": (
            {
                "status": 200,
                "request_url": api_url,
                "final_url": api_url,
                "request_attempts": 1,
            }
            if status == "fetched"
            else None
        ),
        "repo_evidence_matches_api_json": status == "fetched",
        "latest_commit_evidence": (
            {
                "status": 200,
                "request_url": commit_url,
                "final_url": commit_url,
                "request_attempts": 1,
                "data": [
                    {
                        "sha": "abc",
                        "html_url": f"{url}/commit/abc",
                        "commit": {"committer": {"date": timestamp}},
                    }
                ],
            }
            if status == "fetched"
            else None
        ),
        "release_tag_evidence": (
            {
                "status": 200,
                "request_url": GRAPHQL_API,
                "final_url": GRAPHQL_API,
                "request_attempts": 1,
                "data": {
                    "data": {
                        "repository": {
                            "nameWithOwner": f"{owner}/{repo}",
                            "isArchived": False,
                            "releases": {"nodes": []},
                            "refs": {"nodes": []},
                        }
                    }
                },
            }
            if status == "fetched"
            else None
        ),
        "redirect_detected": False,
        "redirect_canonical_url": url if status == "fetched" else None,
        "redirect_full_name": f"{owner}/{repo}" if status == "fetched" else None,
        "redirect_html_url": url if status == "fetched" else None,
        "redirect_final_api_url": (
            api_url
            if status == "fetched"
            else None
        ),
        "activity_warnings": [],
        "success_data_fingerprint": f"fp-{url}" if exists else None,
    }
    return row


def fixture_state(
    targets: list[dict[str, Any]],
    live_rows: int,
    non_target_fp: str = "non-target",
    captured_at: str = "2026-07-15T02:00:00+00:00",
) -> dict[str, Any]:
    status: dict[str, int] = {}
    for row in targets:
        key = row.get("status") if row.get("row_exists") else "missing"
        status[key] = status.get(key, 0) + 1
    protected = {
        "non_target": {
            "rows": live_rows - sum(bool(row.get("row_exists")) for row in targets),
            "fingerprint": non_target_fp,
        },
        "backup": {
            "rows": EXPECTED_BACKUP_ROWS,
            "fingerprint": EXPECTED_BACKUP_FINGERPRINT,
            "owner": EXPECTED_BACKUP_OWNER,
        },
        "universe": {"rows": 2, "fingerprint": "u", "owner": "postgres"},
        "extension": {"rows": 1, "fingerprint": "e", "owner": "postgres"},
    }
    return {
        "captured_at": captured_at,
        "database": {
            "database": "data",
            "session_user": "postgres",
            "current_user": "postgres",
            "server_addr": None,
            "port": "5432",
            "cluster_system_identifier": "123",
        },
        "protected": protected,
        "live": {
            "rows": live_rows,
            "targets": targets,
            "target_count": len(targets),
            "target_existing": sum(bool(row.get("row_exists")) for row in targets),
            "target_missing": sum(not bool(row.get("row_exists")) for row in targets),
            "target_status": status,
            "status": status,
            "integrity": {
                "duplicate_url_norms": [],
                "invalid_url_keys": [],
                "array_anomalies": [],
            },
        },
    }


class OfflineAcceptanceTests(unittest.TestCase):
    def setUp(self) -> None:
        self.url_a = "https://github.com/example/a"
        self.url_b = "https://github.com/example/b"
        self.manifest = []
        for index, url in enumerate((self.url_a, self.url_b), start=1):
            owner, repo = url.removeprefix("https://github.com/").split("/", 1)
            self.manifest.append(
                {
                    "url_norm": url,
                    "repo_owner": owner,
                    "repo_name": repo,
                    "api_url": f"https://api.github.com/repos/{owner}/{repo}",
                    "extension_ids": [1],
                    "extension_names": [repo],
                    "extension_count": 1,
                    "prior_url_candidates": (
                        "https://github.com/old/b" if repo == "b" else ""
                    ),
                }
            )
        self.meta = {
            "path": "/fixture.csv",
            "sha256": "x",
            "mapping_sha256": "y",
            "target_count": 2,
        }

    def make_passing_states(self) -> tuple[dict[str, Any], dict[str, Any]]:
        pre_a = fixture_target(
            self.url_a, True, "fetched", 1, "2026-07-15T00:00:00+00:00"
        )
        pre_b = fixture_target(self.url_b, False, None, None, None)
        before = fixture_state([pre_a, pre_b], live_rows=2)
        post_a = fixture_target(
            self.url_a, True, "fetched", 2, "2026-07-15T03:00:00+00:00"
        )
        post_b = fixture_target(self.url_b, True, "blocked", 1, None)
        post_b["updated_at"] = "2026-07-15T03:00:00+00:00"
        after = fixture_state(
            [post_a, post_b],
            live_rows=3,
            captured_at="2026-07-15T04:00:00+00:00",
        )
        snapshot = {"state": before, "created_at": "now", "integrity_sha256": "fixture"}
        return snapshot, after

    def make_passing_report(self) -> dict[str, Any]:
        snapshot, after = self.make_passing_states()
        return evaluate_report(snapshot, after, self.manifest, self.meta)

    def test_passing_fixture_covers_insert_fetch_and_block(self) -> None:
        report = self.make_passing_report()
        self.assertEqual(report["outcome"], "PASS")
        self.assertEqual(report["counts"]["inserted"], 1)
        self.assertEqual(report["counts"]["attempted"], 2)
        self.assertEqual(report["counts"]["success_refreshed"], 1)
        self.assertEqual(
            len(report["lists"]["prior_url_candidates_for_initial_missing"]), 1
        )

    def test_passing_fixture_renders_complete_markdown_and_json(self) -> None:
        report = self.make_passing_report()
        markdown = render_markdown(report, pathlib.Path("/fixture/snapshot.json"))
        self.assertIn("**Outcome: PASS**", markdown)
        self.assertIn("## Redirect / rename evidence", markdown)
        self.assertIn("## Prior URL candidates", markdown)
        self.assertIn("Successfully refreshed targets", markdown)
        self.assertTrue(json.dumps(report, sort_keys=True))

    def test_non_target_mutation_fails_gate(self) -> None:
        pre_a = fixture_target(
            self.url_a, True, "fetched", 1, "2026-07-15T00:00:00+00:00"
        )
        pre_b = fixture_target(self.url_b, False, None, None, None)
        before = fixture_state([pre_a, pre_b], live_rows=2)
        post_a = fixture_target(
            self.url_a, True, "fetched", 2, "2026-07-15T03:00:00+00:00"
        )
        post_b = fixture_target(self.url_b, True, "blocked", 1, None)
        post_b["updated_at"] = "2026-07-15T03:00:00+00:00"
        after = fixture_state(
            [post_a, post_b],
            live_rows=3,
            non_target_fp="changed",
            captured_at="2026-07-15T04:00:00+00:00",
        )
        report = evaluate_report({"state": before}, after, self.manifest, self.meta)
        gate = next(
            gate
            for gate in report["gates"]
            if gate["name"] == "non_target_full_content_unchanged"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(report["outcome"], "FAIL")

    def test_unattempted_target_fails_gate(self) -> None:
        pre_a = fixture_target(
            self.url_a, True, "fetched", 1, "2026-07-15T00:00:00+00:00"
        )
        pre_b = fixture_target(self.url_b, False, None, None, None)
        before = fixture_state([pre_a, pre_b], live_rows=2)
        post_a = fixture_target(
            self.url_a, True, "fetched", 1, "2026-07-15T00:00:00+00:00"
        )
        post_b = fixture_target(self.url_b, False, None, None, None)
        after = fixture_state(
            [post_a, post_b],
            live_rows=2,
            captured_at="2026-07-15T04:00:00+00:00",
        )
        report = evaluate_report({"state": before}, after, self.manifest, self.meta)
        gate = next(
            gate for gate in report["gates"] if gate["name"] == "all_targets_attempted"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(report["outcome"], "FAIL")

    def test_missing_raw_activity_evidence_fails_gate(self) -> None:
        snapshot, after = self.make_passing_states()
        fetched = after["live"]["targets"][0]
        for field in (
            "requested_evidence",
            "repo_evidence",
            "latest_commit_evidence",
            "release_tag_evidence",
            "api_full_name",
            "api_html_url",
            "redirect_canonical_url",
            "redirect_full_name",
            "redirect_html_url",
            "redirect_final_api_url",
        ):
            fetched[field] = None
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        gate = next(
            gate for gate in report["gates"] if gate["name"] == "activity_raw_evidence"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(report["outcome"], "FAIL")

    def test_rate_limit_below_reserve_fails_gate(self) -> None:
        snapshot, after = self.make_passing_states()
        fetched = after["live"]["targets"][0]
        fetched["rate_limit_remaining"] = RATE_LIMIT_RESERVE - 1
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        gate = next(
            gate for gate in report["gates"] if gate["name"] == "metadata_complete"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(
            report["metadata_coverage"]["rate_limit_remaining"]["fetched_count"],
            1,
        )
        self.assertEqual(report["outcome"], "FAIL")

    def test_blocked_target_requires_rate_limit_evidence(self) -> None:
        snapshot, after = self.make_passing_states()
        blocked = after["live"]["targets"][1]
        blocked["rate_limit_remaining"] = None
        blocked["rate_limit_reset"] = None
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        gate = next(
            gate
            for gate in report["gates"]
            if gate["name"] == "all_targets_preserve_rate_limit_reserve"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(report["outcome"], "FAIL")

    def test_preexisting_target_replacement_fails_gate(self) -> None:
        snapshot, after = self.make_passing_states()
        after["live"]["targets"][0]["created_at"] = "2026-07-15T02:15:00+00:00"
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        gate = next(
            gate
            for gate in report["gates"]
            if gate["name"] == "preexisting_created_at_unchanged"
        )
        self.assertFalse(gate["passed"])
        self.assertEqual(report["outcome"], "FAIL")

    def test_redirected_repo_uses_canonical_commit_endpoint(self) -> None:
        snapshot, after = self.make_passing_states()
        fetched = after["live"]["targets"][0]
        canonical_url = "https://github.com/canonical/renamed"
        canonical_api = "https://api.github.com/repos/canonical/renamed"
        transport_api = "https://api.github.com/repositories/212439157"
        commit_url = f"{canonical_api}/commits?per_page=1&sha=main"
        commit_html = f"{canonical_url}/commit/abc"
        fetched.update(
            {
                "repo_owner": "canonical",
                "repo_name": "renamed",
                "api_url": canonical_api,
                "api_full_name": "canonical/renamed",
                "api_html_url": canonical_url,
                "redirect_detected": True,
                "redirect_canonical_url": canonical_url,
                "redirect_full_name": "canonical/renamed",
                "redirect_html_url": canonical_url,
                "redirect_final_api_url": transport_api,
                "last_commit_html_url": commit_html,
            }
        )
        fetched["repo_evidence"]["final_url"] = transport_api
        fetched["latest_commit_evidence"]["request_url"] = commit_url
        fetched["latest_commit_evidence"]["final_url"] = commit_url
        fetched["latest_commit_evidence"]["data"][0]["html_url"] = commit_html
        fetched["release_tag_evidence"]["data"]["data"]["repository"][
            "nameWithOwner"
        ] = "canonical/renamed"
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        self.assertEqual(report["outcome"], "PASS")
        self.assertEqual(report["counts"]["redirect_or_rename"], 1)

    def test_mixed_case_api_identity_is_not_a_redirect(self) -> None:
        snapshot, after = self.make_passing_states()
        fetched = after["live"]["targets"][0]
        canonical_api = "https://api.github.com/repos/example/a"
        commit_url = "https://api.github.com/repos/Example/A/commits?per_page=1&sha=main"
        mixed_html = "https://github.com/Example/A"
        commit_html = f"{mixed_html}/commit/abc"
        fetched.update(
            {
                "repo_owner": "Example",
                "repo_name": "A",
                "api_url": canonical_api,
                "api_full_name": "Example/A",
                "api_html_url": mixed_html,
                "redirect_detected": False,
                "redirect_canonical_url": self.url_a,
                "redirect_full_name": "Example/A",
                "redirect_html_url": mixed_html,
                "redirect_final_api_url": canonical_api,
                "last_commit_html_url": commit_html,
            }
        )
        fetched["latest_commit_evidence"]["request_url"] = commit_url
        fetched["latest_commit_evidence"]["final_url"] = commit_url
        fetched["latest_commit_evidence"]["data"][0]["html_url"] = commit_html
        fetched["release_tag_evidence"]["data"]["data"]["repository"][
            "nameWithOwner"
        ] = "Example/A"
        report = evaluate_report(snapshot, after, self.manifest, self.meta)
        self.assertEqual(report["outcome"], "PASS")
        self.assertEqual(report["counts"]["redirect_or_rename"], 0)

    def test_output_paths_cannot_alias_inputs_or_each_other(self) -> None:
        with tempfile.TemporaryDirectory() as tmpdir:
            root = pathlib.Path(tmpdir)
            manifest = root / "manifest.csv"
            manifest.write_text("fixture", encoding="utf-8")
            manifest_alias = root / "manifest-alias.csv"
            manifest_alias.symlink_to(manifest)
            with self.assertRaises(AcceptanceError):
                validate_output_paths([manifest_alias], [manifest])
            output = root / "report.json"
            with self.assertRaises(AcceptanceError):
                validate_output_paths([output, output], [manifest])


def command_self_test(_: argparse.Namespace) -> int:
    suite = unittest.defaultTestLoader.loadTestsFromTestCase(OfflineAcceptanceTests)
    result = unittest.TextTestRunner(verbosity=2).run(suite)
    return 0 if result.wasSuccessful() else 1


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    subparsers = parser.add_subparsers(dest="command", required=True)

    snapshot_parser = subparsers.add_parser(
        "snapshot", help="capture immutable pre-run JSON evidence"
    )
    snapshot_parser.add_argument(
        "--manifest", required=True, help="audited universe GitHub target CSV"
    )
    snapshot_parser.add_argument(
        "--output", required=True, help="pre-run snapshot JSON path"
    )
    snapshot_parser.add_argument(
        "--overwrite", action="store_true", help="explicitly replace an existing output"
    )
    snapshot_parser.set_defaults(handler=command_snapshot)

    report_parser = subparsers.add_parser(
        "report", help="compare post-run state and write JSON plus Markdown"
    )
    report_parser.add_argument(
        "--manifest", required=True, help="the exact manifest used for snapshot/worker"
    )
    report_parser.add_argument(
        "--snapshot", required=True, help="pre-run snapshot JSON"
    )
    report_parser.add_argument(
        "--json-output", required=True, help="post-run JSON report path"
    )
    report_parser.add_argument(
        "--markdown-output", required=True, help="post-run Markdown report path"
    )
    report_parser.add_argument(
        "--overwrite", action="store_true", help="explicitly replace existing outputs"
    )
    report_parser.set_defaults(handler=command_report)

    test_parser = subparsers.add_parser(
        "self-test", help="run deterministic offline acceptance fixtures"
    )
    test_parser.set_defaults(handler=command_self_test)
    return parser.parse_args(argv)


def main(argv: list[str] | None = None) -> int:
    try:
        args = parse_args(argv or sys.argv[1:])
        return int(args.handler(args))
    except (AcceptanceError, OSError, subprocess.SubprocessError) as exc:
        print(f"{TOOL_NAME}: {exc}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
