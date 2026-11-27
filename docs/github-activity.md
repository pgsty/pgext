# GitHub Activity Refresh

`pgext sync github` refreshes GitHub repository metadata for the GitHub
repositories referenced by the current `pgext.universe`. It stores raw API
evidence and parsed activity in `pgext.gh_repo`, and optionally backfills
`pgext.universe` and `pgext.extension.extra->'star'` inside a single transaction.

Database selection follows `-d/--database`, `PGURL`, or `postgres:///data`.
Authentication is read from `--token`, `GITHUB_TOKEN`, `GH_TOKEN`, or
`gh auth token`. Cross-origin redirects off `api.github.com` are refused so
tokens are never forwarded off-origin.

## Targeted refresh

Preview selection without calling GitHub or writing to PostgreSQL:

```bash
pgext sync github --ext vector --dry-run
```

Refresh one extension or repository URL:

```bash
pgext sync github --ext vector
pgext sync github --url https://github.com/pgvector/pgvector
```

## Incremental refresh and backfill

Skip repositories already fetched within the last 7 days and stop gracefully
before rate-limit exhaustion (`--reserve` defaults to 200):

```bash
pgext sync github --stale-days 7 --reserve 200
```

Retry previously errored or rate-limited repositories:

```bash
pgext sync github --status error,rate_limited
```

The normal refresh writes the GitHub cache. To also update Universe and the
packaged catalog, use `--backfill`. This includes only repositories successfully
fetched by this run.

```bash
pgext sync github --ext vector --backfill
```

To backfill from the cache, use `--backfill-only` with the `refresh_after` value
logged at the start of the accepted refresh:

```bash
pgext sync github --backfill-only --refresh-after 2026-11-27T00:00:00Z
```

## Backfill mapping

- `stars` = GitHub `stargazers_count`
- `watchers` = GitHub `subscribers_count` (explicit Watch subscriptions)
- `forks` = GitHub `forks_count`
- repository, commit, release/tag, activity, and checked timestamps are
  converted to UTC calendar dates
- backfill accepts complete metrics that match the raw API response and meet
  the fetch-time cutoff
- the UTC catalog check date advances or stays the same
- `last_release` and `last_active` preserve the newer of existing catalog
  evidence and GitHub activity
- `pgext.extension.extra.star` stays synchronized with non-null Universe stars,
  updating `mtime` only for rows whose values actually changed
