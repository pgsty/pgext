# GitHub Activity Refresh

`bin/github_activity.py` refreshes GitHub repository metadata for the GitHub
repositories referenced by the current `pgext.universe`. It stores raw API
evidence and parsed activity in `pgext.gh_repo`; it does not change Universe
activity fields by itself.

The audited database target is the local `data` database on `/tmp:5432` as
`postgres`. Authentication is read from `GH_TOKEN`, `GITHUB_TOKEN`, or
`gh auth token`. Never put a token in a manifest or ledger.

## Targeted refresh

Dry-run one canonical extension:

```bash
python3 bin/github_activity.py --extension pgvector --dry-run
```

Write one canonical extension using an automatically generated complete
manifest and append-only ledger:

```bash
python3 bin/github_activity.py \
  --run-dir tmp/github_activity_runs/pgvector \
  --extension pgvector
```

The same operation can be selected by any normalizable repository URL:

```bash
python3 bin/github_activity.py \
  --run-dir tmp/github_activity_runs/pgvector-url \
  --url https://github.com/pgvector/pgvector
```

The manifest always contains the complete current Universe GitHub scope, even
for a single selected extension. Production writes are refused if that
manifest differs from live Universe mappings. Only the selected repository
mapping and activity row are written.

## Full refresh and resume

Create the acceptance baseline before the first database write:

```bash
python3 bin/github_activity.py \
  --write-manifest tmp/github_activity_run/manifest.csv

python3 bin/github_activity_acceptance.py snapshot \
  --manifest tmp/github_activity_run/manifest.csv \
  --output tmp/github_activity_run/baseline.json
```

Run every repository explicitly:

```bash
python3 bin/github_activity.py \
  --manifest tmp/github_activity_run/manifest.csv \
  --all \
  --ledger tmp/github_activity_run/ledger.jsonl
```

Each batch is transactional and is followed by a durable ledger checkpoint.
Resume the exact saved plan without repeating selection flags:

```bash
python3 bin/github_activity.py \
  --manifest tmp/github_activity_run/manifest.csv \
  --ledger tmp/github_activity_run/ledger.jsonl \
  --resume
```

Use a new ledger for a bounded retry pass:

```bash
python3 bin/github_activity.py \
  --manifest tmp/github_activity_run/manifest.csv \
  --status error --status rate_limited \
  --ledger tmp/github_activity_run/retry.jsonl
```

Generate the read-only acceptance report after collection:

```bash
python3 bin/github_activity_acceptance.py report \
  --manifest tmp/github_activity_run/manifest.csv \
  --snapshot tmp/github_activity_run/baseline.json \
  --json-output tmp/github_activity_run/report.json \
  --markdown-output tmp/github_activity_run/report.md
```

## Analyze and backfill Universe

Run the read-only SQL analysis first:

```bash
PGUSER=postgres psql -X -h /tmp -p 5432 -d data \
  -f db/github_activity_analysis.sql
```

Only after acceptance and analysis pass, backfill successful repository rows.
`refresh_after` is the UTC start time of the accepted full run and prevents an
old cache row from being treated as freshly collected:

```bash
PGUSER=postgres psql -X -h /tmp -p 5432 -d data \
  -v refresh_after='2026-08-10T08:24:00Z' \
  -f db/github_universe_backfill.sql
```

The mapping is deliberate:

- `stars` = GitHub `stargazers_count`
- `watchers` = GitHub `subscribers_count` (explicit Watch subscriptions)
- `forks` = GitHub `forks_count`
- repository, commit, release/tag, activity, and checked timestamps are
  converted to UTC calendar dates
- a failed or blocked repository does not overwrite prior successful Universe
  data
- packaged compatibility rows copy each non-null Universe Star value into
  `pgext.extension.extra.star` while preserving every other `extra` key

`last_release` keeps the newer of existing catalog evidence and GitHub's latest
release-or-tag date; `last_active` does the same for the newest credible GitHub
activity. The backfill updates `mtime` only for rows whose activity fields
actually changed.
