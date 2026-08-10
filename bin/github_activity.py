#!/usr/bin/env python3
"""Refresh Universe-scoped GitHub metadata and activity in pgext.gh_repo.

The worker is deliberately manifest scoped.  Production writes are refused unless
the manifest exactly matches the current pgext.universe GitHub aggregation.  When
the legacy immutable pgext.gh_repo_20260507 backup exists, its audited fingerprint
is also verified before any write.
"""

import argparse
import csv
import hashlib
import io
import json
import os
import pathlib
import socket
import subprocess
import sys
import tempfile
import time
import urllib.error
import urllib.parse
import urllib.request
import uuid
from datetime import datetime, timezone


REST_API = "https://api.github.com"
GRAPHQL_API = "https://api.github.com/graphql"
USER_AGENT = "pgext-github-activity-worker"
RATE_LIMIT_RESERVE = 200
MAX_REDIRECT_REQUESTS = 11
EXPECTED_BACKUP_ROWS = 1342
EXPECTED_BACKUP_FINGERPRINT = "9354ad62777856b2d60dc7d68f110e2b"
REQUIRED_MANIFEST_COLUMNS = {
    "url_norm",
    "repo_owner",
    "repo_name",
    "api_url",
    "extension_ids",
    "extension_names",
    "extension_count",
}

RESULT_COLUMNS = [
    "url_norm",
    "success",
    "status",
    "http_status",
    "rate_limit_remaining",
    "rate_limit_reset",
    "error",
    "etag",
    "repo_owner",
    "repo_name",
    "api_url",
    "api_json",
    "stargazers_count",
    "forks_count",
    "watchers_count",
    "subscribers_count",
    "pushed_at",
    "updated_at_api",
    "default_branch",
    "archived",
    "last_commit_at",
    "last_commit_date",
    "last_commit_sha",
    "last_commit_html_url",
    "latest_release_tag",
    "latest_release_published_at",
    "latest_release_created_at",
    "last_release_at",
    "last_release_date",
    "latest_tag_name",
    "latest_tag_at",
    "last_tag_date",
    "latest_tag_date_source",
    "latest_tag_target_type",
    "latest_tag_commit_sha",
    "latest_tag_commit_at",
    "last_release_or_tag_at",
    "last_release_or_tag_date",
    "last_release_or_tag_source",
    "last_update_at",
    "last_update_date",
    "last_update_source",
    "activity_json",
]

RELEASE_TAG_QUERY = """
query RepoReleaseTagActivity($owner: String!, $name: String!) {
  repository(owner: $owner, name: $name, followRenames: true) {
    nameWithOwner
    isArchived
    releases(first: 1, orderBy: {field: CREATED_AT, direction: DESC}) {
      nodes {
        tagName
        publishedAt
        createdAt
        url
      }
    }
    refs(refPrefix: "refs/tags/", first: 1, orderBy: {field: TAG_COMMIT_DATE, direction: DESC}) {
      nodes {
        name
        target {
          __typename
          oid
          ... on Commit {
            committedDate
          }
          ... on Tag {
            tagger {
              date
            }
            target {
              __typename
              oid
              ... on Commit {
                committedDate
              }
            }
          }
        }
      }
    }
  }
  rateLimit {
    cost
    remaining
    resetAt
  }
}
"""

RATE_LIMIT_QUERY = """
query WorkerRateLimit {
  rateLimit {
    cost
    remaining
    resetAt
  }
}
"""


class GitHubRequestError(Exception):
    def __init__(self, status, message, headers=None, payload=None):
        super().__init__(message)
        self.status = status
        self.headers = headers or {}
        self.payload = payload


class RateLimitError(GitHubRequestError):
    pass


class SafeStop(RuntimeError):
    """Raised when the configured GitHub rate-limit reserve must be preserved."""


class ManifestError(ValueError):
    """Raised when a target manifest is not safe to use."""


class DatabasePreflightError(RuntimeError):
    """Raised when production database prerequisites fail closed."""


class RateBudget:
    def __init__(self, reserve=RATE_LIMIT_RESERVE, request_retries=0):
        self.reserve = reserve
        self.request_retries = request_retries
        self.core_remaining = None
        self.core_reset = None
        self.graphql_remaining = None
        self.graphql_reset = None

    def update_rest(self, response):
        headers = response.get("headers") or {}
        remaining = rate_limit_remaining(headers)
        if remaining is not None:
            self.core_remaining = remaining
        elif self.core_remaining is not None:
            attempts = max(1, int(response.get("request_attempts") or 1))
            self.core_remaining = max(
                0,
                self.core_remaining - attempts * MAX_REDIRECT_REQUESTS,
            )
        reset = rate_limit_reset(headers)
        if reset:
            self.core_reset = reset

    def update_graphql(self, response):
        payload = response.get("data") if isinstance(response.get("data"), dict) else None
        headers = response.get("headers") or {}
        remaining = rate_limit_remaining(headers, payload)
        if remaining is not None:
            self.graphql_remaining = remaining
        elif self.graphql_remaining is not None:
            attempts = max(1, int(response.get("request_attempts") or 1))
            self.graphql_remaining = max(
                0,
                self.graphql_remaining - attempts * MAX_REDIRECT_REQUESTS,
            )
        reset = rate_limit_reset(headers, payload)
        if reset:
            self.graphql_reset = reset

    def ensure_repo_capacity(self):
        # Reserve for the worst urllib request envelope: every endpoint may use
        # all bounded retries and every open may follow urllib's full redirect
        # chain.  Successful response headers then replace this pessimistic
        # estimate with GitHub's authoritative remaining count.
        attempts = self.request_retries + 1
        core_needed = 2 * attempts * MAX_REDIRECT_REQUESTS
        graphql_needed = attempts * MAX_REDIRECT_REQUESTS
        if self.core_remaining is None or self.graphql_remaining is None:
            raise SafeStop("rate-limit preflight did not return both core and GraphQL budgets")
        if self.core_remaining - core_needed < self.reserve:
            raise SafeStop(
                f"REST/core reserve reached: remaining={self.core_remaining}, "
                f"reset={self.core_reset or 'unknown'}"
            )
        if self.graphql_remaining - graphql_needed < self.reserve:
            raise SafeStop(
                f"GraphQL reserve reached: remaining={self.graphql_remaining}, "
                f"reset={self.graphql_reset or 'unknown'}"
            )

    def conservative_values(self):
        choices = []
        if self.core_remaining is not None:
            choices.append((self.core_remaining, self.core_reset))
        if self.graphql_remaining is not None:
            choices.append((self.graphql_remaining, self.graphql_reset))
        if not choices:
            return None, None
        return min(choices, key=lambda item: item[0])

    def as_dict(self):
        return {
            "reserve": self.reserve,
            "request_retries": self.request_retries,
            "core_remaining": self.core_remaining,
            "core_reset": self.core_reset,
            "graphql_remaining": self.graphql_remaining,
            "graphql_reset": self.graphql_reset,
        }


def parse_iso(value):
    if not value:
        return None
    return datetime.fromisoformat(value.replace("Z", "+00:00")).astimezone(timezone.utc)


def latest_date(*items):
    latest = None
    for source, value in items:
        parsed = parse_iso(value)
        if parsed is None:
            continue
        if latest is None or parsed > latest[2]:
            latest = (source, value, parsed)
    if latest is None:
        return None, None
    return latest[0], latest[1]


def normalize_github_repo(raw_url):
    value = (raw_url or "").strip()
    if not value:
        raise ValueError("empty GitHub repository URL")

    if value.startswith("git@github.com:"):
        path = value.split(":", 1)[1]
    else:
        if value.startswith("github.com/"):
            value = "https://" + value
        parsed = urllib.parse.urlparse(value)
        host = parsed.netloc.lower()
        if host == "www.github.com":
            host = "github.com"
        if parsed.scheme == "ssh" and host == "github.com" and parsed.path.startswith("/"):
            path = parsed.path.lstrip("/")
        elif host == "github.com":
            path = parsed.path.lstrip("/")
        else:
            raise ValueError(f"not a github.com repository URL: {raw_url}")

    parts = [urllib.parse.unquote(p) for p in path.split("/") if p]
    if len(parts) < 2:
        raise ValueError(f"not a GitHub owner/repo URL: {raw_url}")
    owner = parts[0]
    repo = parts[1]
    if repo.lower().endswith(".git"):
        repo = repo[:-4]
    if not owner or not repo:
        raise ValueError(f"not a GitHub owner/repo URL: {raw_url}")
    url_norm = f"https://github.com/{owner.lower()}/{repo.lower()}"
    return owner, repo, url_norm


def parse_pg_array(value):
    """Parse the one-dimensional PostgreSQL array literals used by the manifest."""
    text = (value or "").strip()
    if len(text) < 2 or text[0] != "{" or text[-1] != "}":
        raise ManifestError(f"invalid PostgreSQL array literal: {value!r}")
    body = text[1:-1]
    if not body:
        return []

    result = []
    token = []
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
                raise ManifestError(f"invalid quote placement in array literal: {value!r}")
            quoted = True
            continue
        if char == ",":
            item = "".join(token)
            if not token_was_quoted and item == "NULL":
                item = None
            result.append(item)
            token = []
            token_was_quoted = False
            continue
        token.append(char)
    if escaped or quoted:
        raise ManifestError(f"unterminated escape or quote in array literal: {value!r}")
    item = "".join(token)
    if not token_was_quoted and item == "NULL":
        item = None
    result.append(item)
    return result


def validate_manifest_row(raw, line_number):
    prefix = f"manifest line {line_number}"
    url_norm = (raw.get("url_norm") or "").strip()
    try:
        _, _, normalized = normalize_github_repo(url_norm)
    except ValueError as exc:
        raise ManifestError(f"{prefix}: {exc}") from exc
    if url_norm != normalized:
        raise ManifestError(f"{prefix}: noncanonical url_norm {url_norm!r}; expected {normalized!r}")

    owner = (raw.get("repo_owner") or "").strip()
    repo = (raw.get("repo_name") or "").strip()
    url_owner, url_repo = url_norm.removeprefix("https://github.com/").split("/", 1)
    if owner.lower() != url_owner or repo.lower() != url_repo:
        raise ManifestError(f"{prefix}: owner/repo columns do not match url_norm")
    expected_api_url = f"{REST_API}/repos/{url_owner}/{url_repo}"
    api_url = (raw.get("api_url") or "").strip()
    if api_url != expected_api_url:
        raise ManifestError(f"{prefix}: api_url must be {expected_api_url!r}")

    ids_raw = parse_pg_array(raw.get("extension_ids"))
    names = parse_pg_array(raw.get("extension_names"))
    if any(value in (None, "") for value in ids_raw + names):
        raise ManifestError(f"{prefix}: extension arrays may not contain NULL/empty values")
    try:
        extension_ids = [int(value) for value in ids_raw]
        extension_count = int(raw.get("extension_count") or "")
    except ValueError as exc:
        raise ManifestError(f"{prefix}: invalid extension id/count") from exc
    if extension_count <= 0:
        raise ManifestError(f"{prefix}: extension_count must be positive")
    if extension_count != len(extension_ids) or extension_count != len(names):
        raise ManifestError(f"{prefix}: extension_count disagrees with extension arrays")
    if len(set(extension_ids)) != len(extension_ids):
        raise ManifestError(f"{prefix}: duplicate extension id")
    if len(set(names)) != len(names):
        raise ManifestError(f"{prefix}: duplicate extension name")
    if extension_ids != sorted(extension_ids):
        raise ManifestError(f"{prefix}: extension_ids must be sorted")
    if names != sorted(names, key=lambda value: value.encode("utf-8")):
        raise ManifestError(f"{prefix}: extension_names must use C-order sorting")

    return {
        **raw,
        "url_norm": url_norm,
        "repo_owner": owner,
        "repo_name": repo,
        "api_url": api_url,
        "extension_ids_list": extension_ids,
        "extension_names_list": names,
        "extension_count_int": extension_count,
    }


def load_manifest(path):
    manifest_path = pathlib.Path(path).expanduser().resolve()
    if not manifest_path.is_file():
        raise ManifestError(f"manifest does not exist: {manifest_path}")
    with manifest_path.open(newline="", encoding="utf-8") as handle:
        reader = csv.DictReader(handle)
        missing_columns = REQUIRED_MANIFEST_COLUMNS - set(reader.fieldnames or [])
        if missing_columns:
            raise ManifestError(f"manifest missing columns: {', '.join(sorted(missing_columns))}")
        rows = [validate_manifest_row(row, line_number) for line_number, row in enumerate(reader, start=2)]
    if not rows:
        raise ManifestError("manifest has no target rows")
    urls = [row["url_norm"] for row in rows]
    if len(set(urls)) != len(urls):
        duplicates = sorted({url for url in urls if urls.count(url) > 1})
        raise ManifestError(f"manifest has duplicate url_norm values: {duplicates[:5]}")
    if urls != sorted(urls, key=lambda value: value.encode("utf-8")):
        raise ManifestError("manifest rows must be sorted by url_norm in C order")
    return manifest_path, rows


def pg_array_literal(values):
    """Render a one-dimensional PostgreSQL array literal for the manifest."""
    escaped = []
    for value in values:
        item = str(value).replace("\\", "\\\\").replace('"', '\\"')
        escaped.append(f'"{item}"')
    return "{" + ",".join(escaped) + "}"


def manifest_rows_from_live(live_rows):
    rows = []
    for index, live in enumerate(live_rows, start=2):
        raw = {
            "url_norm": live["url_norm"],
            "repo_owner": live["repo_owner"],
            "repo_name": live["repo_name"],
            "api_url": live["api_url"],
            "extension_ids": pg_array_literal(live.get("extension_ids") or []),
            "extension_names": pg_array_literal(live.get("extension_names") or []),
            "extension_count": str(live.get("extension_count") or 0),
        }
        rows.append(validate_manifest_row(raw, index))
    urls = [row["url_norm"] for row in rows]
    if not rows or urls != sorted(urls, key=lambda value: value.encode("utf-8")):
        raise ManifestError("live Universe target rows are empty or not C-order sorted")
    if len(urls) != len(set(urls)):
        raise ManifestError("live Universe target rows contain duplicate URLs")
    return rows


def write_manifest(path, rows, overwrite=False):
    """Atomically write the audited worker columns for a live target set."""
    manifest_path = pathlib.Path(path).expanduser().resolve()
    if manifest_path.exists() and not overwrite:
        raise ManifestError(f"refusing to overwrite existing manifest: {manifest_path}")
    manifest_path.parent.mkdir(parents=True, exist_ok=True)
    fields = [
        "url_norm",
        "repo_owner",
        "repo_name",
        "api_url",
        "extension_ids",
        "extension_names",
        "extension_count",
    ]
    fd, temporary = tempfile.mkstemp(
        prefix=f".{manifest_path.name}.", dir=manifest_path.parent
    )
    try:
        with os.fdopen(fd, "w", encoding="utf-8", newline="") as handle:
            writer = csv.DictWriter(handle, fieldnames=fields, lineterminator="\n")
            writer.writeheader()
            for row in rows:
                writer.writerow({field: row[field] for field in fields})
            handle.flush()
            os.fsync(handle.fileno())
        os.replace(temporary, manifest_path)
    finally:
        try:
            os.unlink(temporary)
        except FileNotFoundError:
            pass
    return manifest_path


def resolve_extension_urls(manifest_rows, requested_extensions):
    """Resolve canonical Universe extension names to their GitHub repositories."""
    exact = {}
    folded = {}
    for row in manifest_rows:
        for name in row.get("extension_names_list") or []:
            exact.setdefault(name, set()).add(row["url_norm"])
            folded.setdefault(name.lower(), set()).add(name)

    resolved = set()
    missing = []
    ambiguous = []
    for requested in requested_extensions or []:
        name = requested.strip()
        if name in exact:
            resolved.update(exact[name])
            continue
        candidates = sorted(folded.get(name.lower()) or [])
        if len(candidates) == 1:
            resolved.update(exact[candidates[0]])
        elif len(candidates) > 1:
            ambiguous.append((name, candidates))
        else:
            missing.append(name)
    if ambiguous:
        raise ManifestError(f"ambiguous extension names: {ambiguous[:5]}")
    if missing:
        raise ManifestError(
            "extensions are absent from the current GitHub target set: "
            + ", ".join(missing[:10])
        )
    return resolved


def parse_commit_response(payload):
    row = {
        "last_commit_at": None,
        "last_commit_sha": None,
        "last_commit_html_url": None,
    }
    if not isinstance(payload, list):
        raise ValueError("commit response must be a JSON array")
    if not payload:
        return row
    commit = payload[0]
    if not isinstance(commit, dict):
        raise ValueError("commit response first element must be an object")
    commit_data = commit.get("commit")
    if not isinstance(commit_data, dict):
        raise ValueError("commit response lacks the nested commit object")
    committer = commit_data.get("committer")
    author = commit_data.get("author")
    if committer is not None and not isinstance(committer, dict):
        raise ValueError("commit.committer must be an object or null")
    if author is not None and not isinstance(author, dict):
        raise ValueError("commit.author must be an object or null")
    commit_at = (committer or {}).get("date") or (author or {}).get("date")
    sha = commit.get("sha")
    html_url = commit.get("html_url")
    if not isinstance(commit_at, str) or not isinstance(sha, str) or not isinstance(html_url, str):
        raise ValueError("commit response lacks date, sha, or html_url")
    parse_iso(commit_at)
    row["last_commit_at"] = commit_at
    row["last_commit_sha"] = sha
    row["last_commit_html_url"] = html_url
    return row


def parse_release_tag_response(payload):
    row = {
        "latest_release_tag": None,
        "latest_release_published_at": None,
        "latest_release_created_at": None,
        "last_release_at": None,
        "latest_tag_name": None,
        "latest_tag_at": None,
        "latest_tag_date_source": None,
        "latest_tag_target_type": None,
        "latest_tag_commit_sha": None,
        "latest_tag_commit_at": None,
        "last_release_or_tag_at": None,
        "last_release_or_tag_source": None,
    }
    if not isinstance(payload, dict) or not isinstance(payload.get("data"), dict):
        raise ValueError("GraphQL response lacks a data object")
    repo = payload["data"].get("repository")
    if not isinstance(repo, dict):
        raise ValueError("GraphQL response lacks the repository object")
    releases = repo.get("releases")
    refs = repo.get("refs")
    if not isinstance(releases, dict) or not isinstance(releases.get("nodes"), list):
        raise ValueError("GraphQL releases.nodes must be an array")
    if not isinstance(refs, dict) or not isinstance(refs.get("nodes"), list):
        raise ValueError("GraphQL refs.nodes must be an array")

    release_nodes = releases["nodes"]
    if release_nodes:
        release = release_nodes[0]
        if not isinstance(release, dict):
            raise ValueError("GraphQL release node must be an object")
        row["latest_release_tag"] = release.get("tagName")
        row["latest_release_published_at"] = release.get("publishedAt")
        row["latest_release_created_at"] = release.get("createdAt")
        row["last_release_at"] = release.get("publishedAt") or release.get("createdAt")
        if not isinstance(row["latest_release_tag"], str) or not isinstance(row["last_release_at"], str):
            raise ValueError("GraphQL release node lacks tagName or date")
        for field in ("latest_release_published_at", "latest_release_created_at"):
            value = row[field]
            if value is not None:
                if not isinstance(value, str):
                    raise ValueError(f"GraphQL {field} must be a timestamp string or null")
                parse_iso(value)

    tag_nodes = refs["nodes"]
    if tag_nodes:
        tag_ref = tag_nodes[0]
        if not isinstance(tag_ref, dict) or not isinstance(tag_ref.get("name"), str):
            raise ValueError("GraphQL tag ref must be an object with a name")
        target = tag_ref.get("target")
        if not isinstance(target, dict):
            raise ValueError("GraphQL tag ref lacks a target object")
        target_type = target.get("__typename")
        if not isinstance(target_type, str) or not target_type:
            raise ValueError("GraphQL tag target lacks __typename")
        row["latest_tag_name"] = tag_ref.get("name")
        row["latest_tag_target_type"] = target_type

        if target_type == "Tag":
            tagger = target.get("tagger")
            underlying = target.get("target")
            if tagger is not None and not isinstance(tagger, dict):
                raise ValueError("GraphQL tagger must be an object or null")
            if underlying is not None and not isinstance(underlying, dict):
                raise ValueError("GraphQL annotated-tag target must be an object or null")
            tagger = tagger or {}
            underlying = underlying or {}
            row["latest_tag_at"] = tagger.get("date")
            row["latest_tag_date_source"] = "tagger.date" if row["latest_tag_at"] else None
            if underlying.get("__typename") == "Commit":
                row["latest_tag_commit_sha"] = underlying.get("oid")
                row["latest_tag_commit_at"] = underlying.get("committedDate")
                if not row["latest_tag_at"]:
                    row["latest_tag_at"] = row["latest_tag_commit_at"]
                    row["latest_tag_date_source"] = "underlying_commit.committedDate"
        elif target_type == "Commit":
            row["latest_tag_at"] = target.get("committedDate")
            row["latest_tag_date_source"] = "commit.committedDate" if row["latest_tag_at"] else None
            row["latest_tag_commit_sha"] = target.get("oid")
            row["latest_tag_commit_at"] = target.get("committedDate")

        if row["latest_tag_at"] is not None:
            if not isinstance(row["latest_tag_at"], str):
                raise ValueError("GraphQL tag date must be a string")
            parse_iso(row["latest_tag_at"])
        if row["latest_tag_commit_at"] is not None:
            if not isinstance(row["latest_tag_commit_at"], str):
                raise ValueError("GraphQL tag commit date must be a string")
            parse_iso(row["latest_tag_commit_at"])
        if row["latest_tag_commit_sha"] is not None and not isinstance(row["latest_tag_commit_sha"], str):
            raise ValueError("GraphQL tag commit oid must be a string")

    source, value = latest_date(
        ("release", row["last_release_at"]),
        ("tag", row["latest_tag_at"]),
    )
    row["last_release_or_tag_at"] = value
    row["last_release_or_tag_source"] = source
    return row


def validate_repo_response(payload):
    if not isinstance(payload, dict):
        raise ValueError("repository response must be an object")
    full_name = payload.get("full_name")
    html_url = payload.get("html_url")
    if not isinstance(full_name, str) or full_name.count("/") != 1:
        raise ValueError("repository response lacks a valid full_name")
    if not isinstance(html_url, str) or not html_url.startswith("https://github.com/"):
        raise ValueError("repository response lacks a valid html_url")
    owner, repo = full_name.split("/", 1)
    if not owner or not repo:
        raise ValueError("repository full_name has an empty owner or name")
    for field in (
        "stargazers_count",
        "forks_count",
        "watchers_count",
        "subscribers_count",
    ):
        value = payload.get(field)
        if isinstance(value, bool) or not isinstance(value, int):
            raise ValueError(f"repository {field} must be an integer")
    if not isinstance(payload.get("archived"), bool):
        raise ValueError("repository archived must be boolean")
    default_branch = payload.get("default_branch")
    if default_branch is not None and not isinstance(default_branch, str):
        raise ValueError("repository default_branch must be a string or null")
    for field in ("pushed_at", "updated_at"):
        value = payload.get(field)
        if value is not None:
            if not isinstance(value, str):
                raise ValueError(f"repository {field} must be a timestamp string or null")
            parse_iso(value)
    return owner, repo


def merge_activity_dates(row):
    merged = dict(row)
    merged["last_commit_date"] = merged.get("last_commit_at")
    merged["last_release_date"] = merged.get("last_release_at")
    merged["last_tag_date"] = merged.get("latest_tag_at")
    merged["last_release_or_tag_date"] = merged.get("last_release_or_tag_at")
    source, value = latest_date(
        ("commit", merged.get("last_commit_at")),
        ("release", merged.get("last_release_at")),
        ("tag", merged.get("latest_tag_at")),
    )
    merged["last_update_at"] = value
    merged["last_update_date"] = value
    merged["last_update_source"] = source
    return merged


def headers_to_dict(headers):
    return {k.lower(): v for k, v in headers.items()}


def decode_json(text):
    if not text:
        return None
    try:
        return json.loads(text)
    except json.JSONDecodeError:
        return None


class GitHubSameOriginRedirectHandler(urllib.request.HTTPRedirectHandler):
    """Allow GitHub API path redirects without forwarding auth off-origin."""

    def redirect_request(self, req, fp, code, msg, headers, newurl):
        resolved = urllib.parse.urljoin(req.full_url, newurl)
        parsed = urllib.parse.urlparse(resolved)
        if (
            parsed.scheme.lower() != "https"
            or (parsed.hostname or "").lower() != "api.github.com"
            or parsed.port not in (None, 443)
            or parsed.username is not None
            or parsed.password is not None
        ):
            raise urllib.error.HTTPError(
                req.full_url,
                code,
                f"refusing cross-origin GitHub API redirect to {resolved}",
                headers,
                fp,
            )
        return super().redirect_request(req, fp, code, msg, headers, resolved)


class GitHubClient:
    def __init__(self, token, proxy=None, timeout=45, max_retries=2, retry_delay=0.75):
        self.token = token
        self.timeout = timeout
        self.max_retries = max_retries
        self.retry_delay = retry_delay
        handlers = [GitHubSameOriginRedirectHandler()]
        if proxy:
            handlers.append(urllib.request.ProxyHandler({"http": proxy, "https": proxy}))
        self.opener = urllib.request.build_opener(*handlers)

    def request_json(self, url, method="GET", payload=None, accept="application/vnd.github+json"):
        body = None
        headers = {
            "Accept": accept,
            "User-Agent": USER_AGENT,
            "X-GitHub-Api-Version": "2022-11-28",
        }
        if self.token:
            headers["Authorization"] = f"Bearer {self.token}"
        if payload is not None:
            body = json.dumps(payload).encode("utf-8")
            headers["Content-Type"] = "application/json"

        last_response = None
        for attempt in range(self.max_retries + 1):
            req = urllib.request.Request(url, data=body, headers=headers, method=method)
            try:
                with self.opener.open(req, timeout=self.timeout) as resp:
                    response_text = resp.read().decode("utf-8", errors="replace")
                    response = {
                        "status": resp.status,
                        "headers": headers_to_dict(resp.headers),
                        "text": response_text,
                        "data": decode_json(response_text),
                        "url": resp.url,
                        "request_url": url,
                        "request_attempts": attempt + 1,
                    }
            except urllib.error.HTTPError as exc:
                response_text = exc.read().decode("utf-8", errors="replace")
                response = {
                    "status": exc.code,
                    "headers": headers_to_dict(exc.headers),
                    "text": response_text,
                    "data": decode_json(response_text),
                    "url": exc.geturl() or url,
                    "request_url": url,
                    "request_attempts": attempt + 1,
                }
            except (urllib.error.URLError, TimeoutError, socket.timeout, OSError) as exc:
                response = {
                    "status": None,
                    "headers": {},
                    "text": str(exc)[:2000],
                    "data": None,
                    "url": url,
                    "request_url": url,
                    "request_attempts": attempt + 1,
                    "network_error": True,
                }

            last_response = response
            transient = response.get("network_error") or (
                isinstance(response.get("status"), int) and 500 <= response["status"] <= 599
            )
            if not transient or attempt >= self.max_retries:
                return response
            time.sleep(self.retry_delay * (2**attempt))
        return last_response

    def rest_repo(self, owner, repo):
        return self.request_json(f"{REST_API}/repos/{quote_path(owner)}/{quote_path(repo)}")

    def rest_latest_commit(self, owner, repo, default_branch):
        query = urllib.parse.urlencode({"per_page": 1, "sha": default_branch})
        url = f"{REST_API}/repos/{quote_path(owner)}/{quote_path(repo)}/commits?{query}"
        return self.request_json(url)

    def release_tag_activity(self, owner, repo):
        payload = {"query": RELEASE_TAG_QUERY, "variables": {"owner": owner, "name": repo}}
        return self.request_json(GRAPHQL_API, method="POST", payload=payload, accept="application/json")

    def rest_rate_limit(self):
        return self.request_json(f"{REST_API}/rate_limit")

    def graphql_rate_limit(self):
        return self.request_json(
            GRAPHQL_API,
            method="POST",
            payload={"query": RATE_LIMIT_QUERY},
            accept="application/json",
        )


def quote_path(value):
    return urllib.parse.quote(value or "", safe="")


def is_rate_limited(response):
    headers = response.get("headers") or {}
    data = response.get("data") or {}
    message = ""
    if isinstance(data, dict):
        message = str(data.get("message") or "")
    message = message.lower()
    graphql_rate = ((data.get("data") or {}).get("rateLimit") or {}) if isinstance(data, dict) else {}
    graphql_errors = data.get("errors") or [] if isinstance(data, dict) else []
    graphql_limited = graphql_rate.get("remaining") == 0 or any(
        str(error.get("type") or "").upper() == "RATE_LIMITED"
        or "rate limit" in str(error.get("message") or "").lower()
        for error in graphql_errors
        if isinstance(error, dict)
    )
    status = response.get("status")
    if status in (403, 429) and not message:
        message = str(response.get("text") or "").lower()
    return (
        graphql_limited
        or status == 429
        or (
            status == 403
            and (
                headers.get("x-ratelimit-remaining") == "0"
                or headers.get("retry-after") not in (None, "")
                or "rate limit" in message
            )
        )
    )


def rate_limit_reset(headers, graphql_payload=None):
    if graphql_payload:
        reset_at = ((graphql_payload.get("data") or {}).get("rateLimit") or {}).get("resetAt")
        if reset_at:
            return reset_at
    reset_epoch = (headers or {}).get("x-ratelimit-reset")
    if reset_epoch:
        try:
            return datetime.fromtimestamp(float(reset_epoch), tz=timezone.utc).isoformat().replace("+00:00", "Z")
        except ValueError:
            return None
    return None


def rate_limit_remaining(headers, graphql_payload=None):
    if graphql_payload:
        remaining = ((graphql_payload.get("data") or {}).get("rateLimit") or {}).get("remaining")
        if remaining is not None:
            return remaining
    remaining = (headers or {}).get("x-ratelimit-remaining")
    if remaining in (None, ""):
        return None
    try:
        return int(remaining)
    except ValueError:
        return None


def api_error_message(response):
    data = response.get("data")
    if isinstance(data, dict):
        if data.get("message"):
            return str(data["message"])
        if data.get("errors"):
            return json.dumps(data["errors"], ensure_ascii=False)[:2000]
    return (response.get("text") or "").strip()[:2000]


def preflight_rate_limits(client, reserve=RATE_LIMIT_RESERVE):
    rest_response = client.rest_rate_limit()
    if rest_response.get("status") != 200 or not isinstance(rest_response.get("data"), dict):
        raise SafeStop(f"REST rate-limit preflight failed: {api_error_message(rest_response) or rest_response.get('status')}")
    resources = rest_response["data"].get("resources") or {}
    core = resources.get("core") or {}
    graphql_hint = resources.get("graphql") or {}
    graphql_hint_remaining = graphql_hint.get("remaining")
    probe_envelope = (client.max_retries + 1) * MAX_REDIRECT_REQUESTS
    if graphql_hint_remaining is None:
        raise SafeStop("REST rate-limit response did not include the GraphQL budget")
    if graphql_hint_remaining - probe_envelope < reserve:
        raise SafeStop(
            "GraphQL reserve would be crossed by the preflight probe: "
            f"remaining={graphql_hint_remaining}, "
            f"reset={epoch_to_iso(graphql_hint.get('reset')) or 'unknown'}"
        )

    graphql_response = client.graphql_rate_limit()
    if graphql_response.get("status") != 200 or not isinstance(graphql_response.get("data"), dict):
        raise SafeStop(
            f"GraphQL rate-limit preflight failed: "
            f"{api_error_message(graphql_response) or graphql_response.get('status')}"
        )
    if graphql_response["data"].get("errors"):
        raise SafeStop(f"GraphQL rate-limit preflight errors: {api_error_message(graphql_response)}")
    graphql = ((graphql_response["data"].get("data") or {}).get("rateLimit") or {})

    budget = RateBudget(reserve=reserve, request_retries=client.max_retries)
    budget.core_remaining = core.get("remaining")
    budget.core_reset = epoch_to_iso(core.get("reset"))
    budget.graphql_remaining = graphql.get("remaining")
    budget.graphql_reset = graphql.get("resetAt")
    budget.ensure_repo_capacity()
    return budget


def epoch_to_iso(value):
    if value in (None, ""):
        return None
    try:
        return datetime.fromtimestamp(float(value), tz=timezone.utc).isoformat().replace("+00:00", "Z")
    except (TypeError, ValueError, OverflowError):
        return None


def discover_token():
    token = os.environ.get("GITHUB_TOKEN") or os.environ.get("GH_TOKEN")
    if token:
        return token, "env"
    gh = subprocess.run(
        ["sh", "-lc", "command -v gh >/dev/null 2>&1 && gh auth token"],
        text=True,
        capture_output=True,
    )
    if gh.returncode == 0 and gh.stdout.strip():
        return gh.stdout.strip(), "gh"
    return None, None


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


def psql_env():
    env = os.environ.copy()
    # The objective explicitly identifies this local database.  Do not inherit a
    # stray PGDATABASE/PGHOST and silently operate on another cluster.
    for variable in ("PGHOSTADDR", "PGSERVICE", "PGSERVICEFILE"):
        env.pop(variable, None)
    env.update({"PGUSER": "postgres", "PGHOST": "/tmp", "PGPORT": "5432", "PGDATABASE": "data"})
    return env


def run_psql(sql, dsn=None, input_text=None, capture=True):
    cmd = ["psql", "-X", "-v", "ON_ERROR_STOP=1"]
    if dsn:
        cmd.extend(["-d", dsn])
    proc = subprocess.run(
        cmd,
        input=input_text if input_text is not None else sql,
        text=True,
        capture_output=capture,
        env=psql_env(),
    )
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.strip() or proc.stdout.strip())
    return proc.stdout


def run_psql_json(sql, dsn=None):
    cmd = ["psql", "-X", "-q", "-t", "-A", "-v", "ON_ERROR_STOP=1"]
    if dsn:
        cmd.extend(["-d", dsn])
    cmd.extend(["-c", sql])
    proc = subprocess.run(cmd, text=True, capture_output=True, env=psql_env())
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.strip() or proc.stdout.strip())
    output = proc.stdout.strip()
    if not output:
        return None
    return json.loads(output)


def run_psql_script_json(script, dsn=None):
    cmd = ["psql", "-X", "-q", "-t", "-A", "-v", "ON_ERROR_STOP=1"]
    if dsn:
        cmd.extend(["-d", dsn])
    proc = subprocess.run(cmd, input=script, text=True, capture_output=True, env=psql_env())
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.strip() or proc.stdout.strip())
    lines = [line for line in proc.stdout.splitlines() if line.strip()]
    if not lines:
        return None
    return json.loads(lines[-1])


def live_universe_targets(dsn=None):
    sql = LIVE_TARGET_CTE + """
SELECT coalesce(jsonb_agg(to_jsonb(live_target) ORDER BY url_norm), '[]'::jsonb)::text
FROM live_target;
"""
    return run_psql_json(sql, dsn=dsn) or []


def mapping_signature(row):
    return (
        row["repo_owner"],
        row["repo_name"],
        row["api_url"],
        tuple(row.get("extension_ids_list", row.get("extension_ids") or [])),
        tuple(row.get("extension_names_list", row.get("extension_names") or [])),
        int(row.get("extension_count_int", row.get("extension_count") or 0)),
    )


def verify_manifest_matches_live(manifest_rows, live_rows):
    manifest = {row["url_norm"]: mapping_signature(row) for row in manifest_rows}
    live = {row["url_norm"]: mapping_signature(row) for row in live_rows}
    missing = sorted(set(live) - set(manifest))
    extra = sorted(set(manifest) - set(live))
    mismatched = sorted(url for url in set(manifest) & set(live) if manifest[url] != live[url])
    if missing or extra or mismatched:
        raise DatabasePreflightError(
            "manifest does not exactly match live Universe targets: "
            f"missing={missing[:5]}, extra={extra[:5]}, mapping_mismatch={mismatched[:5]}"
        )


def database_preflight(manifest_rows, dsn=None):
    objects = run_psql_json(
        """
SELECT jsonb_build_object(
    'database', current_database(),
    'user', current_user,
    'server_addr', inet_server_addr(),
    'port', current_setting('port'),
    'live', to_regclass('pgext.gh_repo'),
    'backup', to_regclass('pgext.gh_repo_20260507'),
    'universe', to_regclass('pgext.universe')
)::text;
""",
        dsn=dsn,
    )
    if (
        objects.get("database") != "data"
        or objects.get("user") != "postgres"
        or objects.get("server_addr") is not None
        or objects.get("port") != "5432"
    ):
        raise DatabasePreflightError(
            "refusing non-local database endpoint: "
            f"database={objects.get('database')}, user={objects.get('user')}, "
            f"server_addr={objects.get('server_addr')}, port={objects.get('port')}"
        )
    required = {"live": "pgext.gh_repo", "universe": "pgext.universe"}
    absent = [name for name, expected in required.items() if objects.get(name) != expected]
    if absent:
        raise DatabasePreflightError(f"required pgext relations absent: {', '.join(absent)}")

    backup = {"present": False, "rows": None, "fingerprint": None, "owner": None}
    if objects.get("backup") == "pgext.gh_repo_20260507":
        backup = run_psql_json(
            """
SELECT jsonb_build_object(
    'present', true,
    'rows', count(*),
    'fingerprint', md5(string_agg(md5(to_jsonb(b)::text), '' ORDER BY url_norm)),
    'owner', (SELECT pg_get_userbyid(c.relowner) FROM pg_class c WHERE c.oid='pgext.gh_repo_20260507'::regclass)
)::text
FROM pgext.gh_repo_20260507 b;
""",
            dsn=dsn,
        )
        if (
            backup.get("rows") != EXPECTED_BACKUP_ROWS
            or backup.get("fingerprint") != EXPECTED_BACKUP_FINGERPRINT
            or backup.get("owner") != "postgres"
        ):
            raise DatabasePreflightError(f"immutable backup check failed: {backup}")

    live_rows = live_universe_targets(dsn=dsn)
    verify_manifest_matches_live(manifest_rows, live_rows)
    states = run_psql_json(
        """
SELECT coalesce(jsonb_agg(jsonb_build_object(
    'url_norm', url_norm,
    'status', status,
    'default_branch', default_branch,
    'api_default_branch', api_json->>'default_branch'
) ORDER BY url_norm), '[]'::jsonb)::text
FROM pgext.gh_repo;
""",
        dsn=dsn,
    ) or []
    return {"objects": objects, "backup": backup, "target_count": len(live_rows)}, {
        row["url_norm"]: row for row in states
    }


def upsert_target_mappings(manifest_rows, selected_urls=None, dsn=None):
    manifest_urls = {row["url_norm"] for row in manifest_rows}
    selected_urls = sorted(
        manifest_urls if selected_urls is None else set(selected_urls)
    )
    unknown_urls = sorted(set(selected_urls) - manifest_urls)
    if not selected_urls:
        raise ValueError("selected URL set is empty")
    if unknown_urls:
        raise ValueError(f"selected URLs are absent from manifest: {unknown_urls[:5]}")

    buf = io.StringIO()
    fields = ["url_norm", "repo_owner", "repo_name", "api_url", "extension_ids", "extension_names", "extension_count"]
    writer = csv.DictWriter(buf, fieldnames=fields, lineterminator="\n")
    writer.writeheader()
    for row in manifest_rows:
        writer.writerow({field: row[field] for field in fields})

    selected_buf = io.StringIO()
    selected_writer = csv.DictWriter(
        selected_buf, fieldnames=["url_norm"], lineterminator="\n"
    )
    selected_writer.writeheader()
    for url_norm in selected_urls:
        selected_writer.writerow({"url_norm": url_norm})

    sql = """
BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
CREATE TEMP TABLE q_gh_repo_target(
    url_norm text PRIMARY KEY,
    repo_owner text NOT NULL,
    repo_name text NOT NULL,
    api_url text NOT NULL,
    extension_ids integer[] NOT NULL,
    extension_names text[] NOT NULL,
    extension_count integer NOT NULL
) ON COMMIT DROP;
COPY q_gh_repo_target FROM STDIN WITH (FORMAT csv, HEADER true);
""" + buf.getvalue() + """\\.
CREATE TEMP TABLE q_selected_url(
    url_norm text PRIMARY KEY
) ON COMMIT DROP;
COPY q_selected_url FROM STDIN WITH (FORMAT csv, HEADER true);
""" + selected_buf.getvalue() + """\\.
CREATE TEMP TABLE q_live_universe_target ON COMMIT DROP AS
WITH normalized AS (
    SELECT id, name, pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\\.)?github\\.com/'
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
SELECT * FROM live_target;
DO $guard$
BEGIN
    IF EXISTS (
        SELECT 1 FROM q_gh_repo_target
        WHERE url_norm !~ '^https://github\\.com/[^/]+/[^/]+$'
           OR url_norm <> lower(url_norm)
           OR repo_owner <> lower(repo_owner)
           OR repo_name <> lower(repo_name)
           OR url_norm <> 'https://github.com/' || repo_owner || '/' || repo_name
           OR api_url <> 'https://api.github.com/repos/' || repo_owner || '/' || repo_name
           OR extension_count <> cardinality(extension_ids)
           OR extension_count <> cardinality(extension_names)
    ) THEN
        RAISE EXCEPTION 'invalid target staging row';
    END IF;
    IF EXISTS (
        SELECT url_norm, repo_owner, repo_name, api_url,
               extension_ids, extension_names, extension_count
        FROM q_gh_repo_target
        EXCEPT
        SELECT url_norm, repo_owner, repo_name, api_url,
               extension_ids, extension_names, extension_count
        FROM q_live_universe_target
    ) OR EXISTS (
        SELECT url_norm, repo_owner, repo_name, api_url,
               extension_ids, extension_names, extension_count
        FROM q_live_universe_target
        EXCEPT
        SELECT url_norm, repo_owner, repo_name, api_url,
               extension_ids, extension_names, extension_count
        FROM q_gh_repo_target
    ) THEN
        RAISE EXCEPTION 'staged target manifest differs from live Universe';
    END IF;
END
$guard$;
WITH missing_before AS (
    SELECT count(*) AS n
    FROM q_gh_repo_target q
    JOIN q_selected_url s USING (url_norm)
    LEFT JOIN pgext.gh_repo r USING (url_norm)
    WHERE r.url_norm IS NULL
), changed AS (
    INSERT INTO pgext.gh_repo(
        url_norm, repo_host, repo_owner, repo_name, api_url,
        extension_ids, extension_names, extension_count
    )
    SELECT url_norm, 'github.com', repo_owner, repo_name, api_url,
           extension_ids, extension_names, extension_count
    FROM q_gh_repo_target
    JOIN q_selected_url USING (url_norm)
    ON CONFLICT (url_norm) DO UPDATE SET
        extension_ids = EXCLUDED.extension_ids,
        extension_names = EXCLUDED.extension_names,
        extension_count = EXCLUDED.extension_count,
        updated_at = now()
    WHERE (pgext.gh_repo.extension_ids, pgext.gh_repo.extension_names, pgext.gh_repo.extension_count)
          IS DISTINCT FROM
          (EXCLUDED.extension_ids, EXCLUDED.extension_names, EXCLUDED.extension_count)
    RETURNING url_norm, status, default_branch,
              api_json->>'default_branch' AS api_default_branch
), states AS (
    SELECT q.url_norm,
           coalesce(c.status, r.status) AS status,
           coalesce(c.default_branch, r.default_branch) AS default_branch,
           coalesce(c.api_default_branch, r.api_json->>'default_branch') AS api_default_branch
    FROM q_gh_repo_target q
    JOIN q_selected_url s USING (url_norm)
    LEFT JOIN pgext.gh_repo r USING (url_norm)
    LEFT JOIN changed c USING (url_norm)
)
SELECT jsonb_build_object(
    'staged', (SELECT count(*) FROM q_gh_repo_target),
    'selected', (SELECT count(*) FROM q_selected_url),
    'missing_before', (SELECT n FROM missing_before),
    'mapping_changed', (SELECT count(*) FROM changed),
    'upserted', (SELECT count(*) FROM states),
    'states', (SELECT jsonb_agg(jsonb_build_object(
        'url_norm', u.url_norm, 'status', u.status,
        'default_branch', u.default_branch,
        'api_default_branch', u.api_default_branch
    ) ORDER BY u.url_norm) FROM states u)
)::text;
COMMIT;
"""
    result = run_psql_script_json(sql, dsn=dsn)
    if (
        result.get("staged") != len(manifest_rows)
        or result.get("selected") != len(selected_urls)
        or result.get("upserted") != len(selected_urls)
    ):
        raise RuntimeError(f"target upsert count mismatch: {result}")
    return result


def sql_literal(value):
    return "'" + str(value).replace("'", "''") + "'"


def response_evidence(response, include_data=True):
    if not response:
        return None
    evidence = {
        "status": response.get("status"),
        "request_url": response.get("request_url"),
        "final_url": response.get("url"),
        "request_attempts": response.get("request_attempts"),
        "rate_limit_remaining": (response.get("headers") or {}).get("x-ratelimit-remaining"),
        "rate_limit_reset": (response.get("headers") or {}).get("x-ratelimit-reset"),
    }
    if include_data:
        evidence["data"] = response.get("data")
    return evidence


def build_error_record(target, status, http_status, error, response=None, budget=None, audit=None):
    headers = (response or {}).get("headers") or {}
    data = (response or {}).get("data") if response else None
    record = empty_activity_record(target)
    remaining = rate_limit_remaining(headers, data if isinstance(data, dict) else None)
    reset = rate_limit_reset(headers, data if isinstance(data, dict) else None)
    if budget:
        budget_remaining, budget_reset = budget.conservative_values()
        remaining = budget_remaining if budget_remaining is not None else remaining
        reset = budget_reset or reset
    record.update(
        {
            "success": False,
            "status": status,
            "http_status": http_status,
            "rate_limit_remaining": remaining,
            "rate_limit_reset": reset,
            "error": (error or "unknown GitHub request error")[:2000],
            # Failure staging deliberately ignores activity_json, but retaining it
            # in the returned record makes dry runs and JSONL ledgers auditable.
            "activity_json": json.dumps(audit or {}, separators=(",", ":"), ensure_ascii=False),
        }
    )
    return record


def empty_activity_record(target):
    record = {column: None for column in RESULT_COLUMNS}
    record["url_norm"] = target.get("url_norm")
    record["repo_owner"] = target.get("repo_owner")
    record["repo_name"] = target.get("repo_name")
    record["api_url"] = target.get("api_url")
    return record


def fetch_activity(target, client, budget):
    budget.ensure_repo_capacity()
    try:
        owner = target.get("repo_owner")
        repo = target.get("repo_name")
        if target.get("url_norm"):
            owner, repo, _ = normalize_github_repo(target["url_norm"])
    except ValueError:
        owner = target.get("repo_owner")
        repo = target.get("repo_name")

    audit = {
        "requested": {
            "url_norm": target.get("url_norm"),
            "repo_owner": owner,
            "repo_name": repo,
            "api_url": target.get("api_url"),
        },
        "warnings": [],
    }
    repo_response = client.rest_repo(owner, repo)
    budget.update_rest(repo_response)
    audit["repo"] = response_evidence(repo_response)
    if is_rate_limited(repo_response):
        return build_error_record(target, "rate_limited", repo_response.get("status"), api_error_message(repo_response), repo_response, budget, audit)
    if repo_response.get("status") == 404:
        return build_error_record(target, "blocked", 404, api_error_message(repo_response) or "repository not found", repo_response, budget, audit)
    if repo_response.get("status") != 200 or not isinstance(repo_response.get("data"), dict):
        return build_error_record(target, "error", repo_response.get("status"), api_error_message(repo_response), repo_response, budget, audit)

    repo_json = repo_response["data"]
    try:
        owner, repo = validate_repo_response(repo_json)
    except (TypeError, ValueError, AttributeError, KeyError) as exc:
        return build_error_record(
            target,
            "error",
            200,
            f"repository metadata parse error: {exc}",
            repo_response,
            budget,
            audit,
        )
    # The successful REST payload is authoritative for this refresh.  Falling
    # back to a cached branch would retain stale state when a repo becomes
    # empty or unborn and GitHub now returns null.
    default_branch = repo_json.get("default_branch")
    canonical_url = f"https://github.com/{owner.lower()}/{repo.lower()}"
    redirect_detected = canonical_url != target.get("url_norm")
    audit["redirect"] = {
        "detected": redirect_detected,
        "canonical_url": canonical_url,
        "full_name": repo_json.get("full_name"),
        "html_url": repo_json.get("html_url"),
        "final_api_url": repo_response.get("url"),
    }

    commit_payload = None
    if default_branch:
        commit_response = client.rest_latest_commit(owner, repo, default_branch)
        budget.update_rest(commit_response)
        audit["latest_commit"] = response_evidence(commit_response)
        if is_rate_limited(commit_response):
            return build_error_record(target, "rate_limited", commit_response.get("status"), api_error_message(commit_response), commit_response, budget, audit)
        if commit_response.get("status") == 200:
            commit_payload = commit_response.get("data")
            try:
                commit_fields = parse_commit_response(commit_payload)
            except (TypeError, ValueError, AttributeError, KeyError) as exc:
                return build_error_record(
                    target,
                    "error",
                    200,
                    f"commit parse error: {exc}",
                    commit_response,
                    budget,
                    audit,
                )
        elif commit_response.get("status") in (404, 409):
            commit_fields = parse_commit_response([])
            warning = api_error_message(commit_response) or f"commit fetch HTTP {commit_response.get('status')}"
            audit["warnings"].append({"kind": "empty_or_unborn_repository", "message": warning})
        else:
            return build_error_record(target, "error", commit_response.get("status"), api_error_message(commit_response), commit_response, budget, audit)
    else:
        commit_fields = parse_commit_response([])
        audit["latest_commit"] = None
        audit["warnings"].append({"kind": "missing_default_branch", "message": "commit request skipped"})

    release_tag_response = client.release_tag_activity(owner, repo)
    budget.update_graphql(release_tag_response)
    audit["release_tag"] = response_evidence(release_tag_response)
    if is_rate_limited(release_tag_response):
        return build_error_record(
            target,
            "rate_limited",
            release_tag_response.get("status"),
            api_error_message(release_tag_response),
            release_tag_response,
            budget,
            audit,
        )
    if release_tag_response.get("status") != 200 or not isinstance(release_tag_response.get("data"), dict):
        return build_error_record(
            target,
            "error",
            release_tag_response.get("status"),
            api_error_message(release_tag_response),
            release_tag_response,
            budget,
            audit,
        )
    if release_tag_response["data"].get("errors"):
        return build_error_record(
            target,
            "error",
            200,
            json.dumps(release_tag_response["data"]["errors"], ensure_ascii=False),
            release_tag_response,
            budget,
            audit,
        )

    try:
        release_tag_fields = parse_release_tag_response(release_tag_response["data"])
        merged = merge_activity_dates({**commit_fields, **release_tag_fields})
    except (TypeError, ValueError, AttributeError, KeyError) as exc:
        return build_error_record(target, "error", 200, f"activity parse error: {exc}", release_tag_response, budget, audit)

    audit["collected_at"] = datetime.now(timezone.utc).isoformat()
    record = empty_activity_record(target)
    record.update(merged)
    record.update(
        {
            "success": True,
            "status": "fetched",
            "http_status": 200,
            "rate_limit_remaining": budget.conservative_values()[0],
            "rate_limit_reset": budget.conservative_values()[1],
            "error": None,
            "etag": (repo_response.get("headers") or {}).get("etag"),
            "repo_owner": owner,
            "repo_name": repo,
            # Keep the stored API identity canonical and human-auditable.  A
            # rename may finish at GitHub's opaque /repositories/<id> URL;
            # that transport URL remains in activity_json.redirect.final_api_url.
            "api_url": (
                f"{REST_API}/repos/{quote_path(owner.lower())}/"
                f"{quote_path(repo.lower())}"
            ),
            "api_json": json.dumps(repo_json, separators=(",", ":"), ensure_ascii=False),
            "stargazers_count": repo_json.get("stargazers_count"),
            "forks_count": repo_json.get("forks_count"),
            "watchers_count": repo_json.get("watchers_count"),
            "subscribers_count": repo_json.get("subscribers_count"),
            "pushed_at": repo_json.get("pushed_at"),
            "updated_at_api": repo_json.get("updated_at"),
            "default_branch": default_branch,
            "archived": repo_json.get("archived"),
            "activity_json": json.dumps(audit, separators=(",", ":"), ensure_ascii=False),
        }
    )
    return record


def apply_activity_updates(records, dsn=None):
    if not records:
        return 0
    buf = io.StringIO()
    writer = csv.DictWriter(buf, fieldnames=RESULT_COLUMNS, lineterminator="\n")
    writer.writeheader()
    for record in records:
        writer.writerow({column: csv_value(record.get(column)) for column in RESULT_COLUMNS})

    sql = """
BEGIN;
CREATE TEMP TABLE q_gh_repo_activity(
    url_norm text PRIMARY KEY,
    success boolean NOT NULL,
    status text,
    http_status integer,
    rate_limit_remaining integer,
    rate_limit_reset timestamptz,
    error text,
    etag text,
    repo_owner text,
    repo_name text,
    api_url text,
    api_json jsonb,
    stargazers_count integer,
    forks_count integer,
    watchers_count integer,
    subscribers_count integer,
    pushed_at timestamptz,
    updated_at_api timestamptz,
    default_branch text,
    archived boolean,
    last_commit_at timestamptz,
    last_commit_date timestamptz,
    last_commit_sha text,
    last_commit_html_url text,
    latest_release_tag text,
    latest_release_published_at timestamptz,
    latest_release_created_at timestamptz,
    last_release_at timestamptz,
    last_release_date timestamptz,
    latest_tag_name text,
    latest_tag_at timestamptz,
    last_tag_date timestamptz,
    latest_tag_date_source text,
    latest_tag_target_type text,
    latest_tag_commit_sha text,
    latest_tag_commit_at timestamptz,
    last_release_or_tag_at timestamptz,
    last_release_or_tag_date timestamptz,
    last_release_or_tag_source text,
    last_update_at timestamptz,
    last_update_date timestamptz,
    last_update_source text,
    activity_json jsonb
) ON COMMIT DROP;
COPY q_gh_repo_activity FROM STDIN WITH (FORMAT csv, HEADER true);
"""
    sql += buf.getvalue()
    sql += """\\.
DO $apply$
DECLARE
    failed_rows integer;
    success_rows integer;
BEGIN
    UPDATE pgext.gh_repo r
    SET status = q.status,
        http_status = q.http_status,
        rate_limit_remaining = q.rate_limit_remaining,
        rate_limit_reset = q.rate_limit_reset,
        error = q.error,
        attempts = r.attempts + 1,
        updated_at = now()
    FROM q_gh_repo_activity q
    WHERE NOT q.success
      AND r.url_norm = q.url_norm
      AND EXISTS (
          SELECT 1 FROM pgext.universe u
          WHERE u.repo_url ~* '^https?://(www\\.)?github\\.com/'
            AND pgext.repo_url_norm(u.repo_url) = q.url_norm
      );
    GET DIAGNOSTICS failed_rows = ROW_COUNT;

    UPDATE pgext.gh_repo r
    SET status = 'fetched',
        http_status = 200,
        rate_limit_remaining = q.rate_limit_remaining,
        rate_limit_reset = q.rate_limit_reset,
        error = NULL,
        etag = q.etag,
        repo_owner = q.repo_owner,
        repo_name = q.repo_name,
        api_url = q.api_url,
        api_json = q.api_json,
        stargazers_count = q.stargazers_count,
        forks_count = q.forks_count,
        watchers_count = q.watchers_count,
        subscribers_count = q.subscribers_count,
        pushed_at = q.pushed_at,
        updated_at_api = q.updated_at_api,
        default_branch = q.default_branch,
        archived = q.archived,
        last_commit_at = q.last_commit_at,
        last_commit_date = q.last_commit_date,
        last_commit_sha = q.last_commit_sha,
        last_commit_html_url = q.last_commit_html_url,
        latest_release_tag = q.latest_release_tag,
        latest_release_published_at = q.latest_release_published_at,
        latest_release_created_at = q.latest_release_created_at,
        last_release_at = q.last_release_at,
        last_release_date = q.last_release_date,
        latest_tag_name = q.latest_tag_name,
        latest_tag_at = q.latest_tag_at,
        last_tag_date = q.last_tag_date,
        latest_tag_date_source = q.latest_tag_date_source,
        latest_tag_target_type = q.latest_tag_target_type,
        latest_tag_commit_sha = q.latest_tag_commit_sha,
        latest_tag_commit_at = q.latest_tag_commit_at,
        last_release_or_tag_at = q.last_release_or_tag_at,
        last_release_or_tag_date = q.last_release_or_tag_date,
        last_release_or_tag_source = q.last_release_or_tag_source,
        last_update_at = q.last_update_at,
        last_update_date = q.last_update_date,
        last_update_source = q.last_update_source,
        activity_json = q.activity_json,
        fetched_at = now(),
        attempts = r.attempts + 1,
        updated_at = now()
    FROM q_gh_repo_activity q
    WHERE q.success
      AND r.url_norm = q.url_norm
      AND EXISTS (
          SELECT 1 FROM pgext.universe u
          WHERE u.repo_url ~* '^https?://(www\\.)?github\\.com/'
            AND pgext.repo_url_norm(u.repo_url) = q.url_norm
      );
    GET DIAGNOSTICS success_rows = ROW_COUNT;

    IF failed_rows <> (SELECT count(*) FROM q_gh_repo_activity WHERE NOT success)
       OR success_rows <> (SELECT count(*) FROM q_gh_repo_activity WHERE success) THEN
        RAISE EXCEPTION 'manifest-scoped result count mismatch: failed %, success %', failed_rows, success_rows;
    END IF;
END
$apply$;
SELECT jsonb_build_object(
    'staged', count(*),
    'success', count(*) FILTER (WHERE success),
    'failed', count(*) FILTER (WHERE NOT success)
)::text
FROM q_gh_repo_activity;
COMMIT;
"""
    result = run_psql_script_json(sql, dsn=dsn)
    if result.get("staged") != len(records):
        raise RuntimeError(f"result staging count mismatch: {result}")
    return result


def csv_value(value):
    if value is None:
        return ""
    if isinstance(value, bool):
        return "true" if value else "false"
    return value


def fetch_summary(dsn=None, anomaly_limit=25):
    sql = LIVE_TARGET_CTE + f"""
, scoped AS (
    SELECT t.url_norm AS target_url_norm, r.*
    FROM live_target t
    LEFT JOIN pgext.gh_repo r USING (url_norm)
), status_counts AS (
    SELECT status, count(*) AS count
    FROM scoped
    WHERE status IS NOT NULL
    GROUP BY status
), summary AS (
    SELECT jsonb_build_object(
        'total', count(*),
        'missing', count(*) FILTER (WHERE status IS NULL),
        'with_api_json', count(api_json),
        'with_stars', count(stargazers_count),
        'with_subscribers', count(subscribers_count),
        'with_commit', count(last_commit_at),
        'with_release', count(last_release_at),
        'with_tag', count(latest_tag_at),
        'with_release_or_tag', count(last_release_or_tag_at),
        'with_update', count(last_update_date),
        'status', coalesce((SELECT jsonb_object_agg(status, count) FROM status_counts), '{{}}'::jsonb)
    ) AS payload
    FROM scoped
), anomalies AS (
    SELECT coalesce(jsonb_agg(to_jsonb(a)), '[]'::jsonb) AS payload
    FROM (
        SELECT target_url_norm AS url_norm, status, http_status, error, last_commit_at, last_release_at, latest_tag_at, last_update_date
        FROM scoped
        WHERE status IS NULL OR status <> 'fetched' OR error IS NOT NULL OR last_update_date IS NULL
        ORDER BY status, url_norm
        LIMIT {int(anomaly_limit)}
    ) a
)
SELECT jsonb_build_object(
    'summary', (SELECT payload FROM summary),
    'anomalies', (SELECT payload FROM anomalies)
)::text;
"""
    return run_psql_json(sql, dsn=dsn)


def print_summary(summary):
    if not summary:
        return
    stats = summary.get("summary") or {}
    total = stats.get("total") or 0
    def pct(value):
        return f"{(value or 0) * 100.0 / total:.1f}%" if total else "0.0%"

    print("Coverage:")
    print(f"  Universe target repos: {total}; missing={stats.get('missing', 0)}")
    print(f"  api_json: {stats.get('with_api_json', 0)} ({pct(stats.get('with_api_json'))})")
    print(f"  stargazers_count: {stats.get('with_stars', 0)} ({pct(stats.get('with_stars'))})")
    print(f"  subscribers_count: {stats.get('with_subscribers', 0)} ({pct(stats.get('with_subscribers'))})")
    print(f"  last_commit_date: {stats.get('with_commit', 0)} ({pct(stats.get('with_commit'))})")
    print(f"  last_release_date: {stats.get('with_release', 0)} ({pct(stats.get('with_release'))})")
    print(f"  last_tag_date: {stats.get('with_tag', 0)} ({pct(stats.get('with_tag'))})")
    print(f"  last_release_or_tag_at: {stats.get('with_release_or_tag', 0)} ({pct(stats.get('with_release_or_tag'))})")
    print(f"  last_update_date: {stats.get('with_update', 0)} ({pct(stats.get('with_update'))})")
    print(f"  status: {json.dumps(stats.get('status', {}), sort_keys=True)}")

    anomalies = summary.get("anomalies") or []
    if anomalies:
        print("Anomaly repos:")
        for row in anomalies:
            msg = row.get("error") or "missing activity date"
            print(f"  {row.get('url_norm')} [{row.get('status')}/{row.get('http_status')}]: {msg}")
    else:
        print("Anomaly repos: none")


def parse_args(argv):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--manifest", help="audited current-Universe GitHub target CSV")
    parser.add_argument(
        "--write-manifest",
        metavar="PATH",
        help="write the current complete Universe GitHub target manifest and exit",
    )
    parser.add_argument(
        "--run-dir",
        help="directory for an automatically generated manifest and production ledger",
    )
    parser.add_argument(
        "--dsn",
        default=None,
        help="reserved; this audited run requires postgres@/tmp:5432/data",
    )
    parser.add_argument("--limit", type=int, default=0, help="maximum repos to process; 0 means all selected repos")
    parser.add_argument("--status", action="append", help="filter by existing gh_repo status; repeatable")
    parser.add_argument("--url", action="append", help="process a manifest GitHub URL; repeatable")
    parser.add_argument(
        "--extension",
        action="append",
        help="process the GitHub repository for a canonical Universe extension; repeatable",
    )
    parser.add_argument(
        "--all",
        action="store_true",
        help="explicitly process every repository in the current manifest",
    )
    parser.add_argument("--proxy", default=os.environ.get("HTTPS_PROXY") or os.environ.get("HTTP_PROXY"))
    parser.add_argument("--dry-run", action="store_true", help="fetch without any database write")
    parser.add_argument(
        "--dry-run-without-db-preflight",
        action="store_true",
        help="dry-run-only escape hatch for absent DB prerequisites; requires --limit 1..5 and performs no DB call",
    )
    parser.add_argument("--summary-only", action="store_true", help="only print coverage and anomaly summary")
    parser.add_argument("--anomaly-limit", type=int, default=25)
    parser.add_argument("--batch-size", type=int, default=25, help="atomic database batch size (25-50)")
    parser.add_argument("--min-delay", type=float, default=0.05, help="seconds to sleep between repos")
    parser.add_argument("--timeout", type=float, default=45.0, help="per-request timeout seconds")
    parser.add_argument("--request-retries", type=int, default=2, help="bounded retries for network/5xx responses (0-3)")
    parser.add_argument(
        "--ledger",
        help="append-only JSONL ledger (required for production; use one file per logical pass)",
    )
    parser.add_argument(
        "--resume",
        action="store_true",
        help="resume the saved run plan; selection flags must be omitted",
    )
    return parser.parse_args(argv)


def append_ledger(path, payload, durable=False):
    if not path:
        return
    ledger_path = pathlib.Path(path).expanduser().resolve()
    ledger_path.parent.mkdir(parents=True, exist_ok=True)
    if ledger_path.exists() and not ledger_path.is_file():
        raise ManifestError(f"ledger is not a regular file: {ledger_path}")
    if ledger_path.is_file() and ledger_path.stat().st_size:
        with ledger_path.open("rb") as existing:
            existing.seek(-1, os.SEEK_END)
            if existing.read(1) != b"\n":
                raise ManifestError(
                    "ledger has an incomplete final record; resume the saved "
                    "run to repair it before appending"
                )
    with ledger_path.open("a", encoding="utf-8") as handle:
        handle.write(json.dumps(payload, ensure_ascii=False, separators=(",", ":")) + "\n")
        if durable:
            handle.flush()
            os.fsync(handle.fileno())


def read_ledger_events(path, repair_torn_tail=False):
    ledger_path = pathlib.Path(path).expanduser().resolve()
    if not ledger_path.is_file():
        return []
    mode = "r+b" if repair_torn_tail else "rb"
    with ledger_path.open(mode) as handle:
        data = handle.read()
        chunks = data.splitlines(keepends=True)
        events = []
        offset = 0
        truncated = False
        for index, raw_line in enumerate(chunks, start=1):
            line_start = offset
            offset += len(raw_line)
            content = raw_line.rstrip(b"\r\n")
            if not content.strip():
                continue
            try:
                event = json.loads(content.decode("utf-8"))
            except (UnicodeDecodeError, json.JSONDecodeError) as exc:
                is_torn_tail = index == len(chunks) and not raw_line.endswith(b"\n")
                if repair_torn_tail and is_torn_tail:
                    handle.seek(line_start)
                    handle.truncate()
                    handle.flush()
                    os.fsync(handle.fileno())
                    truncated = True
                    print(
                        f"Repaired torn final ledger record at line {index}; "
                        "the uncheckpointed batch will be replayed.",
                        file=sys.stderr,
                        flush=True,
                    )
                    break
                raise ManifestError(f"invalid JSONL ledger line {index}: {exc}") from exc
            if not isinstance(event, dict):
                raise ManifestError(f"JSONL ledger line {index} is not an object")
            events.append((index, event))
        if repair_torn_tail and data and not truncated and not data.endswith(b"\n"):
            handle.seek(0, os.SEEK_END)
            handle.write(b"\n")
            handle.flush()
            os.fsync(handle.fileno())
    return events


def plan_sha256(plan):
    canonical = json.dumps(
        plan,
        ensure_ascii=False,
        sort_keys=True,
        separators=(",", ":"),
    ).encode("utf-8")
    return hashlib.sha256(canonical).hexdigest()


def build_run_plan(manifest_sha256, statuses, requested_urls, limit, planned_urls):
    statuses = normalize_statuses(statuses)
    requested_urls = sorted(set(requested_urls), key=lambda value: value.encode("utf-8"))
    planned_urls = list(planned_urls)
    if not planned_urls or len(planned_urls) != len(set(planned_urls)):
        raise ManifestError("run plan requires nonempty unique planned URLs")
    return {
        "schema_version": 1,
        "run_id": str(uuid.uuid4()),
        "manifest_sha256": manifest_sha256,
        "statuses": statuses,
        "requested_urls": requested_urls,
        "limit": limit,
        "planned_count": len(planned_urls),
        "planned_urls": list(planned_urls),
    }


def load_resume_state(path, manifest_sha256, manifest_urls):
    ledger_path = pathlib.Path(path).expanduser().resolve()
    if not ledger_path.is_file() or ledger_path.stat().st_size == 0:
        raise ManifestError(f"--resume requires an existing nonempty ledger: {ledger_path}")
    events = read_ledger_events(ledger_path, repair_torn_tail=True)
    plans = [(line, event) for line, event in events if event.get("event") == "run_plan"]
    if len(plans) != 1:
        raise ManifestError(
            f"resume ledger must contain exactly one run_plan event; found {len(plans)}"
        )
    plan_line, plan_event = plans[0]
    if plan_event.get("manifest_sha256") != manifest_sha256:
        raise ManifestError(f"run_plan at ledger line {plan_line} belongs to a different manifest")
    plan = plan_event.get("plan")
    if not isinstance(plan, dict):
        raise ManifestError(f"run_plan at ledger line {plan_line} has no plan object")
    required_plan_keys = {
        "schema_version",
        "run_id",
        "manifest_sha256",
        "statuses",
        "requested_urls",
        "limit",
        "planned_count",
        "planned_urls",
    }
    if set(plan) != required_plan_keys:
        raise ManifestError(f"run_plan at ledger line {plan_line} has an invalid schema")
    if plan.get("schema_version") != 1 or plan.get("manifest_sha256") != manifest_sha256:
        raise ManifestError(f"run_plan at ledger line {plan_line} has invalid identity fields")
    if not isinstance(plan.get("run_id"), str) or not plan["run_id"]:
        raise ManifestError(f"run_plan at ledger line {plan_line} has an invalid run_id")
    statuses = plan.get("statuses")
    requested_urls = plan.get("requested_urls")
    planned_urls = plan.get("planned_urls")
    if (
        not isinstance(statuses, list)
        or any(not isinstance(value, str) for value in statuses)
        or statuses != normalize_statuses(statuses)
    ):
        raise ManifestError(f"run_plan at ledger line {plan_line} has invalid statuses")
    if (
        not isinstance(requested_urls, list)
        or any(not isinstance(value, str) for value in requested_urls)
        or requested_urls != sorted(set(requested_urls), key=lambda value: value.encode("utf-8"))
    ):
        raise ManifestError(f"run_plan at ledger line {plan_line} has invalid requested_urls")
    if type(plan.get("limit")) is not int or plan["limit"] < 0:
        raise ManifestError(f"run_plan at ledger line {plan_line} has an invalid limit")
    if (
        not isinstance(planned_urls, list)
        or not planned_urls
        or any(not isinstance(value, str) for value in planned_urls)
        or len(planned_urls) != len(set(planned_urls))
        or type(plan.get("planned_count")) is not int
        or plan["planned_count"] != len(planned_urls)
    ):
        raise ManifestError(f"run_plan at ledger line {plan_line} has invalid planned URLs")
    unknown_plan_urls = sorted(set(planned_urls) - manifest_urls)
    if unknown_plan_urls:
        raise ManifestError(
            f"run_plan at ledger line {plan_line} contains non-manifest URLs: "
            f"{unknown_plan_urls[:5]}"
        )
    unknown_requested_urls = sorted(set(requested_urls) - manifest_urls)
    if unknown_requested_urls:
        raise ManifestError(
            f"run_plan at ledger line {plan_line} requests non-manifest URLs: "
            f"{unknown_requested_urls[:5]}"
        )
    if requested_urls and not set(planned_urls) <= set(requested_urls):
        raise ManifestError(f"run_plan at ledger line {plan_line} exceeds requested_urls")
    if plan["limit"] and len(planned_urls) > plan["limit"]:
        raise ManifestError(f"run_plan at ledger line {plan_line} exceeds its limit")
    expected_plan_sha256 = plan_sha256(plan)
    if plan_event.get("plan_sha256") != expected_plan_sha256:
        raise ManifestError(f"run_plan checksum mismatch at ledger line {plan_line}")

    applied = set()
    for index, event in events:
        if event.get("event") != "batch_applied":
            continue
        if event.get("manifest_sha256") != manifest_sha256:
            raise ManifestError(
                f"ledger batch at line {index} belongs to a different manifest"
            )
        if event.get("plan_sha256") != expected_plan_sha256:
            raise ManifestError(f"ledger batch at line {index} belongs to a different run plan")
        urls = event.get("urls")
        if (
            not isinstance(urls, list)
            or not urls
            or any(not isinstance(url, str) for url in urls)
            or len(urls) != len(set(urls))
        ):
            raise ManifestError(f"invalid batch_applied URLs at ledger line {index}")
        unknown = sorted(set(urls) - set(planned_urls))
        if unknown:
            raise ManifestError(
                f"ledger batch at line {index} contains non-plan URLs: {unknown[:5]}"
            )
        overlap = sorted(set(urls) & applied)
        if overlap:
            raise ManifestError(
                f"ledger batch at line {index} repeats checkpointed URLs: {overlap[:5]}"
            )
        result = event.get("result")
        if not isinstance(result, dict):
            raise ManifestError(f"ledger batch at line {index} has no result object")
        staged = result.get("staged")
        success = result.get("success")
        failed = result.get("failed")
        if (
            type(staged) is not int
            or type(success) is not int
            or type(failed) is not int
            or staged != len(urls)
            or success < 0
            or failed < 0
            or success + failed != staged
        ):
            raise ManifestError(f"ledger batch result mismatch at line {index}")
        applied.update(urls)
    return plan, expected_plan_sha256, applied


def apply_and_checkpoint_batch(
    records,
    dsn,
    ledger,
    manifest_path,
    manifest_sha256,
    current_plan_sha256,
    planned_urls,
):
    urls = [record["url_norm"] for record in records]
    if not urls or len(urls) != len(set(urls)) or not set(urls) <= set(planned_urls):
        raise ManifestError("refusing to apply a duplicate, empty, or non-plan batch")
    result = apply_activity_updates(records, dsn=dsn)
    if (
        not isinstance(result, dict)
        or type(result.get("staged")) is not int
        or type(result.get("success")) is not int
        or type(result.get("failed")) is not int
        or result["staged"] != len(urls)
        or result["success"] < 0
        or result["failed"] < 0
        or result["success"] + result["failed"] != result["staged"]
    ):
        raise RuntimeError(f"invalid activity update result: {result}")
    append_ledger(
        ledger,
        {
            "event": "batch_applied",
            "recorded_at": datetime.now(timezone.utc).isoformat(),
            "manifest": str(manifest_path),
            "manifest_sha256": manifest_sha256,
            "plan_sha256": current_plan_sha256,
            "urls": urls,
            "result": result,
        },
        durable=True,
    )
    return result


def selected_targets(
    manifest_rows,
    states,
    statuses=None,
    url_norms=None,
    limit=0,
    exclude_url_norms=None,
):
    statuses = normalize_statuses(statuses)
    requested = set(url_norms or [])
    excluded = set(exclude_url_norms or [])
    manifest_urls = {row["url_norm"] for row in manifest_rows}
    absent = sorted(requested - manifest_urls)
    if absent:
        raise ManifestError(f"requested URLs are not in the manifest: {absent[:5]}")
    targets = []
    for manifest_row in manifest_rows:
        if requested and manifest_row["url_norm"] not in requested:
            continue
        if manifest_row["url_norm"] in excluded:
            continue
        state = states.get(manifest_row["url_norm"]) or {}
        if statuses and state.get("status") not in statuses:
            continue
        target = dict(manifest_row)
        target.update(
            {
                "current_status": state.get("status"),
                "default_branch": state.get("default_branch"),
                "api_default_branch": state.get("api_default_branch"),
            }
        )
        targets.append(target)
    if limit > 0:
        targets = targets[:limit]
    return targets


def normalize_statuses(statuses):
    allowed = {"pending", "collecting", "fetched", "rate_limited", "blocked", "error"}
    normalized = sorted(set(statuses or []))
    bad = set(normalized) - allowed
    if bad:
        raise ManifestError(f"invalid gh_repo status filters: {', '.join(sorted(bad))}")
    return normalized


def main(argv=None):
    try:
        return main_checked(argv or sys.argv[1:])
    except (ManifestError, DatabasePreflightError, SafeStop, RuntimeError, OSError) as exc:
        print(f"Refusing to continue: {exc}", file=sys.stderr)
        return 2


def main_checked(argv):
    args = parse_args(argv)
    manifest_was_supplied = bool(args.manifest)
    if args.dsn:
        raise ManifestError(
            "explicit --dsn is disabled for this audited run; use the fixed "
            "postgres@/tmp:5432/data connection"
        )
    if args.limit < 0:
        raise ManifestError("--limit may not be negative")
    if not 25 <= args.batch_size <= 50:
        raise ManifestError("--batch-size must be between 25 and 50")
    if not 0 <= args.request_retries <= 3:
        raise ManifestError("--request-retries must be between 0 and 3")

    selection_flags = bool(args.status or args.url or args.extension or args.limit)
    if args.all and selection_flags:
        raise ManifestError("--all cannot be combined with --status, --url, --extension, or --limit")

    if args.write_manifest:
        if (
            args.manifest
            or args.run_dir
            or args.resume
            or args.dry_run
            or args.summary_only
            or args.all
            or selection_flags
        ):
            raise ManifestError("--write-manifest is a standalone operation")
        live_rows = live_universe_targets(dsn=args.dsn)
        manifest_rows = manifest_rows_from_live(live_rows)
        manifest_path = write_manifest(args.write_manifest, manifest_rows)
        digest = hashlib.sha256(manifest_path.read_bytes()).hexdigest()
        print(f"Wrote {len(manifest_rows)} current Universe GitHub targets to {manifest_path}")
        print(f"Manifest SHA-256: {digest}")
        return 0

    run_dir = pathlib.Path(args.run_dir).expanduser().resolve() if args.run_dir else None
    if args.resume and run_dir:
        if not args.manifest:
            args.manifest = str(run_dir / "manifest.csv")
        if not args.ledger:
            args.ledger = str(run_dir / "ledger.jsonl")

    if not args.manifest:
        if args.dry_run_without_db_preflight:
            raise ManifestError("DB-bypass dry runs require an explicit --manifest")
        if not manifest_was_supplied and not (
            args.all or selection_flags or args.summary_only or args.resume
        ):
            raise ManifestError(
                "an automatic manifest requires --all, --extension, --url, --status, or --limit"
            )
        if run_dir is None:
            stamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ")
            run_dir = (
                pathlib.Path.cwd()
                / "tmp"
                / "github_activity_runs"
                / f"{stamp}-{uuid.uuid4().hex[:8]}"
            ).resolve()
        live_rows = live_universe_targets(dsn=args.dsn)
        auto_rows = manifest_rows_from_live(live_rows)
        args.manifest = str(write_manifest(run_dir / "manifest.csv", auto_rows))
        print(f"Generated current Universe manifest: {args.manifest}", flush=True)

    production = not args.dry_run and not args.summary_only
    if production and not args.ledger and run_dir:
        args.ledger = str(run_dir / "ledger.jsonl")
    if production and not args.ledger:
        raise ManifestError("production runs require --ledger or --run-dir")

    if args.resume and (not args.ledger or args.dry_run or args.summary_only):
        raise ManifestError(
            "--resume requires a production run with --ledger and cannot be "
            "combined with --dry-run or --summary-only"
        )
    if args.resume and (args.status or args.url or args.extension or args.limit or args.all):
        raise ManifestError(
            "--resume restores the saved run plan; omit selection flags"
        )
    if args.dry_run_without_db_preflight:
        if not args.dry_run or not 1 <= args.limit <= 5:
            raise ManifestError("--dry-run-without-db-preflight requires --dry-run and --limit 1..5")
        if args.summary_only or args.status:
            raise ManifestError("DB-bypass dry run cannot use --summary-only or --status")

    manifest_path, manifest_rows = load_manifest(args.manifest)
    manifest_sha256 = hashlib.sha256(manifest_path.read_bytes()).hexdigest()
    manifest_urls = {row["url_norm"] for row in manifest_rows}
    statuses = normalize_statuses(args.status)
    url_norms = set()
    for requested_url in args.url or []:
        _, _, url_norm = normalize_github_repo(requested_url)
        url_norms.add(url_norm)
    url_norms.update(resolve_extension_urls(manifest_rows, args.extension))
    requested_urls = sorted(url_norms, key=lambda value: value.encode("utf-8"))

    plan = None
    current_plan_sha256 = None
    applied_urls = set()
    if production:
        ledger_path = pathlib.Path(args.ledger).expanduser().resolve()
        if args.resume:
            plan, current_plan_sha256, applied_urls = load_resume_state(
                ledger_path,
                manifest_sha256,
                manifest_urls,
            )
            print(
                f"Resume plan loaded: planned={plan['planned_count']}, "
                f"committed={len(applied_urls)}",
                flush=True,
            )
        else:
            if ledger_path.exists() and not ledger_path.is_file():
                raise ManifestError(f"ledger is not a regular file: {ledger_path}")
            if ledger_path.is_file() and ledger_path.stat().st_size:
                raise ManifestError(
                    "new production runs require a new or empty ledger; use --resume "
                    "for the existing logical pass"
                )

    states = {}
    preflight = None
    if not args.dry_run_without_db_preflight:
        preflight, states = database_preflight(manifest_rows, dsn=args.dsn)
        backup_label = "absent"
        if preflight["backup"].get("present"):
            backup_label = (
                f"{preflight['backup']['rows']}/"
                f"{preflight['backup']['fingerprint']}"
            )
        print(
            f"Database preflight passed: targets={preflight['target_count']}, "
            f"backup={backup_label}",
            flush=True,
        )

    if args.summary_only:
        print_summary(fetch_summary(dsn=args.dsn, anomaly_limit=args.anomaly_limit))
        return 0

    if args.dry_run:
        targets = selected_targets(
            manifest_rows,
            states,
            statuses,
            url_norms,
            args.limit,
        )
        planned_urls = [target["url_norm"] for target in targets]
    else:
        if not args.resume:
            planned_targets = selected_targets(
                manifest_rows,
                states,
                statuses,
                url_norms,
                args.limit,
            )
            if not planned_targets:
                print("No manifest-scoped GitHub repos selected; no database writes made.")
                print_summary(fetch_summary(dsn=args.dsn, anomaly_limit=args.anomaly_limit))
                return 0
            planned_urls = [target["url_norm"] for target in planned_targets]
            plan = build_run_plan(
                manifest_sha256,
                statuses,
                requested_urls,
                args.limit,
                planned_urls,
            )
            current_plan_sha256 = plan_sha256(plan)
            append_ledger(
                args.ledger,
                {
                    "event": "run_plan",
                    "recorded_at": datetime.now(timezone.utc).isoformat(),
                    "manifest": str(manifest_path),
                    "manifest_sha256": manifest_sha256,
                    "plan_sha256": current_plan_sha256,
                    "plan": plan,
                },
                durable=True,
            )
        else:
            planned_urls = plan["planned_urls"]

        upsert = upsert_target_mappings(
            manifest_rows, selected_urls=planned_urls, dsn=args.dsn
        )
        states = {row["url_norm"]: row for row in upsert.get("states") or []}
        print(
            f"Target mapping upsert passed: staged={upsert['staged']}, "
            f"new={upsert['missing_before']}, "
            f"mapping_changed={upsert.get('mapping_changed', 0)}",
            flush=True,
        )
        targets = selected_targets(
            manifest_rows,
            states,
            url_norms=planned_urls,
            exclude_url_norms=applied_urls,
        )
    if not targets:
        print("No manifest-scoped GitHub repos selected.")
        if not args.dry_run_without_db_preflight:
            print_summary(fetch_summary(dsn=args.dsn, anomaly_limit=args.anomaly_limit))
        if production:
            append_ledger(
                args.ledger,
                {
                    "event": "run_complete",
                    "recorded_at": datetime.now(timezone.utc).isoformat(),
                    "manifest": str(manifest_path),
                    "manifest_sha256": manifest_sha256,
                    "plan_sha256": current_plan_sha256,
                    "processed": 0,
                    "already_applied": len(applied_urls),
                },
                durable=True,
            )
        return 0

    token, _ = discover_token()
    if not token:
        print("No GitHub token available. Set GITHUB_TOKEN/GH_TOKEN or run `gh auth login`.", file=sys.stderr)
        return 2

    print(f"GitHub authentication ready; selected manifest repos: {len(targets)}", flush=True)
    client = GitHubClient(
        token=token,
        proxy=args.proxy,
        timeout=args.timeout,
        max_retries=args.request_retries,
    )
    try:
        budget = preflight_rate_limits(client)
    except SafeStop as exc:
        append_ledger(
            args.ledger,
            {
                "event": "safe_stop",
                "recorded_at": datetime.now(timezone.utc).isoformat(),
                "manifest": str(manifest_path),
                "manifest_sha256": manifest_sha256,
                "plan_sha256": current_plan_sha256,
                "stage": "rate_limit_preflight",
                "reason": str(exc),
                "processed": 0,
                "budget": None,
            },
            durable=not args.dry_run,
        )
        print(f"Safe stop during rate-limit preflight: {exc}", file=sys.stderr)
        return 3
    print(
        f"Rate-limit preflight: core={budget.core_remaining}, "
        f"graphql={budget.graphql_remaining}, reserve={budget.reserve}",
        flush=True,
    )
    append_ledger(
        args.ledger,
        {
            "event": "run_start",
            "recorded_at": datetime.now(timezone.utc).isoformat(),
            "manifest": str(manifest_path),
            "manifest_sha256": manifest_sha256,
            "plan_sha256": current_plan_sha256,
            "selected": len(targets),
            "already_applied": len(applied_urls),
            "resume": args.resume,
            "statuses": plan["statuses"] if production else statuses,
            "budget": budget.as_dict(),
        },
        durable=not args.dry_run,
    )
    updates = []
    processed = 0
    stopped = None
    for idx, target in enumerate(targets, start=1):
        try:
            record = fetch_activity(target, client, budget)
        except SafeStop as exc:
            stopped = str(exc)
            break
        updates.append(record)
        processed += 1
        append_ledger(
            args.ledger,
            {
                "event": "attempt",
                "recorded_at": datetime.now(timezone.utc).isoformat(),
                "manifest": str(manifest_path),
                "manifest_sha256": manifest_sha256,
                "plan_sha256": current_plan_sha256,
                "index": idx,
                "total": len(targets),
                "record": record,
                "budget": budget.as_dict(),
            },
        )
        print(
            f"[{idx}/{len(targets)}] {target['url_norm']} -> {record['status']}"
            f" commit={record.get('last_commit_date') or '-'}"
            f" release={record.get('last_release_date') or '-'}"
            f" tag={record.get('last_tag_date') or '-'}",
            flush=True,
        )
        if not args.dry_run and len(updates) >= args.batch_size:
            result = apply_and_checkpoint_batch(
                updates,
                args.dsn,
                args.ledger,
                manifest_path,
                manifest_sha256,
                current_plan_sha256,
                planned_urls,
            )
            applied_urls.update(record["url_norm"] for record in updates)
            print(f"Wrote batch of {len(updates)} rows.", flush=True)
            updates.clear()
        if record["status"] == "rate_limited":
            stopped = record.get("error") or "GitHub rate limit response"
            break
        if args.min_delay:
            time.sleep(args.min_delay)

    if args.dry_run:
        print(f"Dry run complete: fetched={processed}; no database writes.", flush=True)
    else:
        if updates:
            apply_and_checkpoint_batch(
                updates,
                args.dsn,
                args.ledger,
                manifest_path,
                manifest_sha256,
                current_plan_sha256,
                planned_urls,
            )
            applied_urls.update(record["url_norm"] for record in updates)
            print(f"Wrote final batch of {len(updates)} rows.", flush=True)
        print(f"Processed {processed} manifest-scoped pgext.gh_repo rows.", flush=True)

    if stopped:
        append_ledger(
            args.ledger,
            {
                "event": "safe_stop",
                "recorded_at": datetime.now(timezone.utc).isoformat(),
                "manifest": str(manifest_path),
                "manifest_sha256": manifest_sha256,
                "plan_sha256": current_plan_sha256,
                "reason": stopped,
                "processed": processed,
                "budget": budget.as_dict(),
            },
            durable=not args.dry_run,
        )
    if stopped:
        if not args.dry_run_without_db_preflight:
            try:
                print_summary(
                    fetch_summary(dsn=args.dsn, anomaly_limit=args.anomaly_limit)
                )
            except Exception as exc:  # The durable stop must remain resumable.
                print(f"Summary after safe stop failed: {exc}", file=sys.stderr)
        print(f"Safe stop: {stopped}", file=sys.stderr)
        return 3
    if not args.dry_run_without_db_preflight:
        print_summary(fetch_summary(dsn=args.dsn, anomaly_limit=args.anomaly_limit))
    append_ledger(
        args.ledger,
        {
            "event": "run_complete",
            "recorded_at": datetime.now(timezone.utc).isoformat(),
            "manifest": str(manifest_path),
            "manifest_sha256": manifest_sha256,
            "plan_sha256": current_plan_sha256,
            "processed": processed,
            "already_applied": len(applied_urls),
        },
        durable=not args.dry_run,
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
