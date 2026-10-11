## Usage

Sources:

- [README.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/README.md)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/CHANGELOG.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/Cargo.toml)
- [sql/pg_tviews--0.1.0-beta.26--0.1.0-beta.27.sql](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/sql/pg_tviews--0.1.0-beta.26--0.1.0-beta.27.sql)
- [docs/user-guides/operators.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/docs/user-guides/operators.md)
- [docs/DEPRECATION_WARNINGS.md](https://github.com/fraiseql/pg_tviews/blob/4820bfb7505d8c0614d5151dfbc2d1a0f2ab0596/docs/DEPRECATION_WARNINGS.md)

`pg_tviews` 0.1.0-beta.27 maintains derived TVIEW tables transactionally from analyzed queries and base-table triggers. Objects live in the fixed `tviews` schema. This source supports PostgreSQL 16–18.

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

### Upgrade to beta.27

Back up and rehearse first. From beta.20 or later, installing the matching library and running the following re-derives every TVIEW and reinstalls its triggers; no separate re-registration pass is required. If analysis fails, repair or drop the named definition on the old version, then retry. Legacy SQL version 0.1.0 uses the upstream migration script instead.

```sql
ALTER EXTENSION pg_tviews UPDATE;
```

Review dependencies before upgrading: removed APIs and `tviews.tviewschema` are dropped, the latter with CASCADE. Dependent user objects can therefore be removed. Global maintenance functions lose PUBLIC EXECUTE; grant the needed functions explicitly to deploy/restore operators. Single-TVIEW maintenance also checks ownership. Reserved managed-index names must not be reused for different user indexes; rename conflicting indexes before a later dump/restore.

### Refresh and Concurrency Boundaries

Pending refresh work at commit now causes SQLSTATE `55000`; malformed plans and missing mappings also fail writes instead of committing stale results. Concurrent writes take value/key locks and can produce `40P01` or `40001`; retry both, and prefer READ COMMITTED. `pg_tviews.lock_escalation_threshold` defaults to 64; `tviews.pg_tviews_queue_stats()` reports lock waits and escalations.

Time-dependent definitions need an external refresh schedule through `tviews.pg_tviews_refresh_time_dependent()`. Function table reads and materialized-view reads require the documented maintenance policies. Refresh functions run with a restricted search path, so called functions should qualify their relations. UNLOGGED TVIEWs lose contents after crash recovery and cannot be read on standbys; plan rebuilding or use LOGGED storage. `pg_tviews.uncascaded_policy` defaults to error. Removed cascade/insert/delete and conversion APIs must not be used by existing integration code.
