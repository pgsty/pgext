## Usage

Sources:

- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/CHANGELOG.md)
- [docs/reference/api.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/reference/api.md)
- [docs/concurrency.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/concurrency.md)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/user-guides/operators.md)
- [docs/operations/upgrade.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/operations/upgrade.md)
- [docs/operations/upgrades.md](https://github.com/fraiseql/pg_tviews/blob/84157ed337b5b94baa60f3734c04f26a17a132b6/docs/operations/upgrades.md)

`pg_tviews` 0.1.0-beta.35 maintains derived TVIEW tables transactionally from analyzed queries and base-table triggers. It supports PostgreSQL 16–18, installs its API in schema `tviews`, and requires `jsonb_delta`.

### Core Workflow

Preload the library on every server using it, preserving any existing preload entries, then restart PostgreSQL. Preloading is mandatory; installation and subsequent functions or writes fail with SQLSTATE `55000` when it is absent.

```ini
shared_preload_libraries = 'pg_tviews'
```

Install as a superuser and create a small TVIEW:

```sql
CREATE EXTENSION jsonb_delta;
CREATE EXTENSION pg_tviews;
CREATE TABLE tb_user (
    pk_user bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text NOT NULL
);
INSERT INTO tb_user (name) VALUES ('Alice'), ('Bob');
CREATE TABLE tv_user AS
SELECT pk_user, id, jsonb_build_object('id', id, 'name', name) AS data
FROM tb_user;
SELECT * FROM tv_user;
SELECT * FROM tviews.pg_tviews_check_health();
```

A definition has one SELECT, the conventional integer entity key, a UUID id and JSONB data. Maintenance runs as the TVIEW owner, who needs SELECT on its sources; writers need base-table privileges and readers need TVIEW SELECT. `tviews.pg_tviews_create_or_replace` requires ownership of an existing TVIEW and treats supplied options as a complete declaration: omitted options revert to defaults. LOGGED is the default.

### Concurrency and Objects

Writers and flush/refresh/rebuild operations require READ COMMITTED. Stronger isolation is rejected with SQLSTATE `0A000`; this release no longer deliberately raises `40001`. Deadlocks can still return `40P01` and require transaction retry. Pending or suspended refresh at commit fails rather than leaving stale committed results.

Use `tviews.registry` to inspect registrations and `tviews.contract_version()` for the API contract, now 3. Current monitoring names include `tviews.pg_tviews_show_queue`, `tviews.pg_tviews_show_suspended`, `tviews.pg_tviews_count_feed_gaps` and `tviews.pg_tviews_show_replication`. Earlier health, queue and feed-gap names were removed without aliases. `tviews.stats` reports cross-session maintenance counters.

TVIEWs over tables written by logical subscriptions are refused. Publish the maintained TVIEW on the publisher instead. UNLOGGED TVIEWs are empty after crash recovery or promotion and are unavailable on standbys; plan rebuilds or use LOGGED storage.

### Configuration and Change Feed

The five operator settings are `pg_tviews.auto_rebuild_databases` (default all databases), `pg_tviews.publish_changes` (off), `pg_tviews.max_refresh_rows` (0, unlimited), `pg_tviews.lock_escalation_threshold` (64) and `pg_tviews.audit_enabled` (off). The former escalation value -1 is removed. Database rebuild selection needs a restart; the other settings require superuser control.

The hidden superuser diagnostic `pg_tviews.max_queue_size` defaults to 1000000. Former configurable depth/cache/batch limits are now fixed internal limits; remove obsolete settings instead of assuming they still work. Exceeding refresh limits can fail a write with `54000`.

Change notifications on `tviews_changes` are not a durable log. Queue pressure can allow a commit while skipping its notification; monitor feed gaps and invalidate consumer caches when they increase. Prepared transactions with pending change publication are rejected.

### Upgrade to beta.35

Back up and rehearse first, then pause writers. Stop the server before replacing loaded libraries, install matching files and restart. A direct SQL update is supported only from the preceding release, beta.34:

```sql
ALTER EXTENSION pg_tviews UPDATE TO '0.1.0-beta.35';
SELECT tviews.pg_tviews_version(), tviews.contract_version();
SELECT * FROM tviews.pg_tviews_check_health() WHERE status <> 'OK';
SELECT schema, name, needs_reregister FROM tviews.registry;
```

Between restart and SQL update, base-table writes fail with `55000`. The update re-derives every TVIEW's plan and triggers as its owner, leaves rows intact, and replaces per-TVIEW triggers with one set per base table. Failure to analyze any definition rolls back the entire update. Restore grants for replaced functions, including `pg_tviews_ensure_propagation_indexes_all(boolean)` and `pg_tviews_reset_stats_all()`, before resuming writers.

Older releases require the documented drop/recreate migration using saved definitions in dependency order; preserve source rows and keep writers paused until all TVIEWs are recreated. Do not use historical beta.30–33 update instructions as a direct path to beta.35. No downgrade scripts exist: rollback requires the prior matching binaries and a verified backup restored into a new database.
