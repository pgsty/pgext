## Usage

Sources:

- [pg_tviews--0.1.0-beta.30--0.1.0-beta.31.sql](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/sql/pg_tviews--0.1.0-beta.30--0.1.0-beta.31.sql)
- [pg_tviews--0.1.0-beta.31--0.1.0-beta.32.sql](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/sql/pg_tviews--0.1.0-beta.31--0.1.0-beta.32.sql)

- [README.md](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/README.md)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/CHANGELOG.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/Cargo.toml)
- [sql/pg_tviews--0.1.0-beta.29--0.1.0-beta.30.sql](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/sql/pg_tviews--0.1.0-beta.29--0.1.0-beta.30.sql)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/docs/user-guides/operators.md)
- [docs/DEPRECATION_WARNINGS.md](https://github.com/fraiseql/pg_tviews/blob/778346fb5df06a38f470901974d609d2b9dc26a6/docs/DEPRECATION_WARNINGS.md)

`pg_tviews` 0.1.0-beta.32 maintains derived TVIEW tables transactionally from analyzed queries and base-table triggers. Objects live in the fixed `tviews` schema. This source supports PostgreSQL 16–18.

### Core Workflow

Basic SQL functions can load the library lazily, but the upstream quick start preloads it so hooks are active in every session. Automatic rebuild workers require preload and a server restart. Preserve other libraries in the existing configuration.

```ini
shared_preload_libraries = 'pg_tviews'
```

```sql
CREATE EXTENSION pg_tviews;
SET search_path = public, tviews;
CREATE TABLE tb_post (pk_post bigint PRIMARY KEY, title text);
INSERT INTO tb_post VALUES (1, 'Example');
SELECT tviews.pg_tviews_create('post',
  $$ SELECT pk_post, jsonb_build_object('title', title) AS data FROM tb_post $$);
SELECT * FROM tv_post;
SELECT * FROM tviews.registry;
SELECT * FROM tviews.pg_tviews_health_check();
```

Definitions must be exactly one SELECT and expose the expected integer entity key and JSONB data. Callers need schema/table/trigger privileges and must own TVIEWs they maintain. `tviews.registry`, `tviews.contract_version()` and `tviews.registry.managed_indexes` describe registered objects; use these interfaces instead of editing internal plans. `tviews.pg_tviews_create_or_replace` changes a definition, while `tviews.pg_tviews_refresh` performs a refresh as the TVIEW owner.

### Upgrade to beta.30

beta.30 includes the separately unreleased beta.28 and beta.29 changes. Back up and rehearse, install the matching library and restart PostgreSQL, update each database, then re-register all TVIEWs as an authorized maintenance role:

```sql
ALTER EXTENSION pg_tviews UPDATE;
SELECT * FROM tviews.pg_tviews_reregister_all();
```

The upgrade re-derives propagation plans, replaces several interfaces and creates missing reference-column indexes; allow time for index builds on large TVIEWs. `tviews.contract_version()` is now 2. Tools read options through `tviews.registry.options`, and named arguments use `tview =>`. `tviews.registry.references` describes relationships proven by foreign keys; `tviews.registry.typename` supplies the type name.

Options passed to `tviews.pg_tviews_create_or_replace` are a complete declaration: omitted options revert to defaults, so explicitly preserve the options you need. New TVIEWs default to LOGGED storage; existing tables retain their persistence. Changing storage options can rewrite all rows. Remove deleted GUCs from configuration, notably the old `pg_tviews.uncascaded_policy`; policy now belongs to each TVIEW's options. When crossing earlier releases, also review removed APIs, CASCADE removal of `tviews.tviewschema` and maintenance-function privilege changes.

### Refresh and Concurrency Boundaries

Pending refresh work at commit now causes SQLSTATE `55000`; malformed plans and missing mappings also fail writes instead of committing stale results. Concurrent writes take value/key locks and can produce `40P01` or `40001`; retry both, and prefer READ COMMITTED. `pg_tviews.lock_escalation_threshold` defaults to 64; `tviews.pg_tviews_queue_stats()` reports lock waits and escalations.

Time-dependent definitions need an external refresh schedule through `tviews.pg_tviews_refresh_time_dependent()`. Function table reads and materialized-view reads require the documented maintenance policies. Refresh functions run with a restricted search path, so called functions should qualify their relations. UNLOGGED TVIEWs lose contents after crash recovery and cannot be read on standbys; plan rebuilding or use LOGGED storage. The per-TVIEW uncascaded policy defaults to error. Removed cascade/insert/delete and conversion APIs must not be used by existing integration code.

### Statistics, Notifications and Write Limits

`tviews.stats` exposes refresh counters across sessions and requires preload. `pg_tviews.auto_rebuild_databases` now defaults to all databases to refill UNLOGGED TVIEWs reset by crash recovery or promotion. `pg_tviews.max_queue_size` defaults to 1000000. A superuser can bound a statement's refresh work with `pg_tviews.max_refresh_rows`; exceeding it rejects the write with `54000`.

`pg_tviews.publish_changes` defaults to off. A superuser can select types or rows to publish committed changes on `tviews_changes`. beta.30 messages use `v: 2` and add `c` for affected parent collections; listeners that do not understand the version should invalidate all cached data. Above 90% NOTIFY queue usage, a transaction may commit without publishing and log a warning; this is not a durable change log. PREPARE TRANSACTION fails with `25000` while changes await publication.

### Upgrading from beta.30 to beta.32

Install the matching library, restart through the preload-library procedure, and run `ALTER EXTENSION pg_tviews UPDATE` to reach catalog revision 9. Upgrading from beta.30 or beta.31 does not require TVIEW re-registration. The earlier re-registration requirements still apply when crossing older releases.

`tviews.pg_tviews_feed_gaps()` and queue-statistics field `feed_gaps` count committed transactions that skipped publication since cluster startup and require preload. Listeners should poll the counter and invalidate their entire cache when it changes. beta.32 removes per-session `publish_skipped`; use the new counter. `pg_tviews_flush_and_report` adds collections when publishing rows; invalidated_types covers an entire unlistable type and its collections, so stale row lists are not complete results.

These releases also fix writes depending on privileges on unrelated TVIEW schemas and read the catalog as entities and plans to reduce unrelated TVIEW work. After a logical restore that omitted TVIEW table data, a later write reaching an empty TVIEW refills it when its source view has rows; this does not guarantee an immediately complete restored table.
