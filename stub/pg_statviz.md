## Usage

Sources:

- [pg_statviz v1.2.1 release](https://github.com/vyruss/pg_statviz/releases/tag/v1.2.1)
- [pg_statviz v1.2.1 README](https://github.com/vyruss/pg_statviz/blob/v1.2.1/README.md)
- [pg_statviz v1.2 installation SQL](https://github.com/vyruss/pg_statviz/blob/v1.2.1/pg_statviz--1.2.sql)
- [pg_statviz v1.1 to v1.2 upgrade SQL](https://github.com/vyruss/pg_statviz/blob/v1.2.1/pg_statviz--1.1--1.2.sql)
- [pg_statviz v1.2 control file](https://github.com/vyruss/pg_statviz/blob/v1.2.1/pg_statviz.control)
- [pg_statviz v1.2 metadata](https://github.com/vyruss/pg_statviz/blob/v1.2.1/META.json)
- [pg_statviz v1.2.1 Python package metadata](https://github.com/vyruss/pg_statviz/blob/v1.2.1/pyproject.toml)
- [pg_statviz v1.2.1 AI provider implementation](https://github.com/vyruss/pg_statviz/blob/v1.2.1/src/pg_statviz/libs/ai.py)
- [Official PGXN v1.2.0 distribution](https://pgxn.org/dist/pg_statviz/1.2.0/README.html)

`pg_statviz` v1.2 (distributed by PGXN as 1.2.0) is a pure SQL and PL/pgSQL statistics snapshot extension plus a separately installed Python visualization utility. The extension stores cumulative and dynamic PostgreSQL statistics in the fixed `pgstatviz` schema; the utility reads a selected time range and generates charts or optional AI-assisted HTML reports. It requires PostgreSQL 13 or later and supports PostgreSQL through version 19, needs no `shared_preload_libraries`, and does not require a restart. The utility requires Python 3.11 or later.

### Capture and Retain Snapshots

Have an administrator install the extension, then let a dedicated collection role inherit `pg_monitor` and schedule `pgstatviz.snapshot()` with cron or another external job runner.

```sql
CREATE EXTENSION pg_statviz;

GRANT pg_monitor TO stats_collector;

SELECT pgstatviz.snapshot();

DELETE FROM pgstatviz.snapshots
WHERE snapshot_tstamp < CURRENT_DATE - 90;
```

Deleting parent rows cascades to the associated samples. `pgstatviz.delete_snapshots()` instead truncates the complete history. Pick an interval and retention window based on the shortest event worth observing and the resulting table growth; raw PostgreSQL counters are cumulative and can reset independently, so analyze timestamped deltas rather than treating stored values as rates.

### Stored Data and Version Boundaries

The main relations are `pgstatviz.snapshots`, `pgstatviz.blocking`, `pgstatviz.buf`, `pgstatviz.conf`, `pgstatviz.conn`, `pgstatviz.db`, `pgstatviz.io`, `pgstatviz.lock`, `pgstatviz.repl`, `pgstatviz.slru`, `pgstatviz.wait`, and `pgstatviz.wal`. Samples include configuration values, connection user names and ages, replication application and slot names, waits, blocking and ordinary locks, I/O, database counters, and WAL counters. Protect the tables, dumps, charts, and reports as operational data.

Configuration is stored only when it changes, so `pgstatviz.conf` need not contain one row for every snapshot. `pg_stat_wal` data is collected on PostgreSQL 14 and later; `pg_stat_io` data is collected on PostgreSQL 16 and later, with PostgreSQL 18's byte-based fields handled separately. Version 1.2 also records PostgreSQL 19's `wal_fpi_bytes` and captures the new PostgreSQL 18/19 I/O-worker, effective-WAL-level, and autovacuum-scoring settings when those settings exist. On older supported versions the tables remain part of the schema, but unavailable collectors and settings are skipped.

The extension marks its snapshot tables for extension-aware dumps. This allows history to be moved with `pg_dump`, but retention and backup size still need deliberate limits.

### Visualize a Time Range

Install the utility separately and pass normal libpq connection options. The `analyze` command runs every analysis module; individual modules such as `blocking`, `conn`, `io`, `wait`, and `wal` can be selected when a narrower report is sufficient. The v1.2 `blocking` module summarizes blocked and blocking session counts by lock type using snapshots built from `pg_blocking_pids()`; it therefore includes soft blocks as well as hard lock conflicts.

```bash
pip install pg_statviz

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports
```

Restrict database credentials and report-directory access. A visualization role needs read access to the captured schema but does not need permission to collect or delete snapshots.

### Privilege Boundary

The v1.2 installation SQL grants every member of `pg_monitor` schema usage, function execution, and `SELECT`, `INSERT`, `DELETE`, and `TRUNCATE` on all `pgstatviz` tables. Consequently, membership allows both snapshot collection and complete history removal through `pgstatviz.delete_snapshots()`; it is not a read-only visualization role.

If collection, visualization, and retention administration must be separated, revise the default grants after installation and grant only the required functions and table privileges to dedicated roles. Recheck those grants after an extension update.

### Upgrade to v1.2

After installing the v1.2 extension files, upgrade each database that already has `pg_statviz` installed:

```sql
ALTER EXTENSION pg_statviz UPDATE TO '1.2';
```

The v1.1-to-v1.2 migration adds `pgstatviz.blocking`, adds `pgstatviz.wal.wal_fpi_bytes`, replaces the snapshot functions, grants `pg_monitor` access to the new table, and marks it for extension-aware dumps. Replacing package files alone does not apply these database changes. Back up operational history and account for the added table and column before upgrading any external report or restore workflow that assumes the v1.1 schema.

### Optional AI and Cloud Data Review

Normal chart generation makes no LLM request. AI mode requires the optional `pg_statviz[ai]` dependencies and an explicit `--ai` flag. Claude is the default cloud provider and reads `ANTHROPIC_API_KEY`; Gemini reads `GOOGLE_API_KEY`; OpenAI or an OpenAI-compatible endpoint reads `OPENAI_API_KEY` and can be redirected with `OPENAI_BASE_URL`; `--ai local` uses a local Ollama service. The current defaults are `claude-sonnet-5`, `gemini-3.7-flash`, `gpt-5.6-luna`, and `gemma4:e4b`; `OPENAI_MODEL` overrides the OpenAI-compatible model. These are implementation defaults, not a guarantee that a provider account or local runtime will continue to offer them.

```bash
pip install 'pg_statviz[ai]'

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports \
  --ai gemini
```

For a cloud provider, the request can include chart images and summarized series together with the captured PostgreSQL version, primary/standby role, hostname, relevant configuration values, deterministic findings, user or role names, and replication identifiers. Treat that as an explicit operational-data export: review provider retention and regional policy, minimize the selected time range, secure generated HTML and PNG files, and use an approved outbound path. The prompt's data envelopes reduce prompt-injection risk but do not provide confidentiality, authorization, or a substitute for provider governance.

### 1.2.1 Companion-Utility Fix

Source tag 1.2.1 retains SQL/control version 1.2 and Python package metadata version 1.2. Its only code change relative to 1.2 is the local Ollama client compatibility fix: the utility omits the `think` argument for ollama-python below 0.5.0 or when version metadata is unavailable. Installing the updated utility addresses that failure; an additional SQL extension update beyond 1.2 is not required.
