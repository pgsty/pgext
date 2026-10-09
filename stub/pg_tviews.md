## Usage

Sources:

- [beta.25 migration](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/sql/pg_tviews--0.1.0-beta.24--0.1.0-beta.25.sql)
- [sql/pg_tviews--0.1.0-beta.23--0.1.0-beta.24.sql](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/sql/pg_tviews--0.1.0-beta.23--0.1.0-beta.24.sql)
- [sql/pg_tviews--0.1.0-beta.21--0.1.0-beta.22.sql](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/sql/pg_tviews--0.1.0-beta.21--0.1.0-beta.22.sql)
- [sql/pg_tviews--0.1.0-beta.22--0.1.0-beta.23.sql](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/sql/pg_tviews--0.1.0-beta.22--0.1.0-beta.23.sql)
- [README.md](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/README.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/CHANGELOG.md)
- [scripts/migrate-from-0.1.0.sql](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/scripts/migrate-from-0.1.0.sql)
- [docs/reference/read-contract.md](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/docs/reference/read-contract.md)
- [docs/operations/replication.md](https://github.com/fraiseql/pg_tviews/blob/ef096bc3515652969bffb69bac69db4e86325193/docs/operations/replication.md)

`pg_tviews` 0.1.0-beta.25 maintains derived TVIEW tables transactionally through analyzed queries and base-table triggers. All extension objects now live in the fixed `tviews` schema, and the SQL version matches the release.

### Core Workflow

```sql
CREATE EXTENSION pg_tviews;
CREATE TABLE tb_post (pk_post bigint PRIMARY KEY, title text);
INSERT INTO tb_post VALUES (1, 'Example');
SELECT tviews.pg_tviews_create('post',
  $$ SELECT pk_post, jsonb_build_object('title', title) AS data FROM tb_post $$);
SELECT * FROM tv_post;
SELECT * FROM tviews.registry;
SELECT * FROM tviews.pg_tviews_health_check();
```

### Operational Boundaries

Basic creation/refresh can load lazily without shared preload in beta.23. Preload and restart remain necessary for automatic rebuild workers and their postmaster settings. The control is not superuser-only; callers need schema creation/trigger rights and must own TVIEWs they replace or drop. Cascades run as each TVIEW owner. The query must expose the entity key and JSONB data expected by the TVIEW contract. UNLOGGED TVIEWs remain unreadable on standbys and are emptied by crash/promotion; select LOGGED storage or plan rebuilding.

Use `tviews.registry` and `tviews.contract_version()` for tool discovery. `tviews.pg_tviews_create_or_replace` applies compatible definition/options changes, while `tviews.pg_tviews_reregister_all` re-derives metadata and triggers without replacing rows. A library/catalog mismatch blocks writes until migration is complete. Existing SQL version 0.1.0 (through beta.19) requires the linked transactional migration script, which preserves TVIEW rows but refuses external extension dependencies; ordinary ALTER EXTENSION alone is not that migration. Back up and rehearse first.

### beta.21–22 Upgrade

`pg_tviews.uncascaded_policy` is captured at creation: `warn` (the default before beta.25), `error`, or `full_refresh` controls tables that cannot map writes to a TVIEW key. beta.22 fixes lost updates between READ COMMITTED refreshes and several outer-join, CTE and DISTINCT ON dependency paths. After the normal beta-version upgrade, re-register all TVIEWs; legacy SQL version 0.1.0 still needs the separate migration above.

```sql
ALTER EXTENSION pg_tviews UPDATE;
SELECT * FROM tviews.pg_tviews_reregister_all();
```

### beta.23 Row Identity

beta.23 gives every TVIEW one recorded identity in `tviews.registry.identity`. A `DISTINCT ON` view uses its single column key; composite or expression keys are rejected. It fixes multi-group writes, key changes and propagation to parent TVIEWs. Re-registration also adjusts primary/unique indexes, so allow for its table changes.

### beta.24 Generated Columns

beta.24 repairs dependency tracking and refresh-key mapping for PostgreSQL 18 virtual generated columns. After installing the matching library, run ALTER EXTENSION UPDATE and `tviews.pg_tviews_reregister_all()` so existing TVIEW metadata and triggers are rebuilt. This is an extension update; it does not add virtual generated columns to older PostgreSQL releases.

### beta.25 Backing Views and Policy

Backing views move from the application schema into `tviews.<schema>__tv_<entity>`. Tools should resolve them through `tviews.registry.view`. The default `pg_tviews.uncascaded_policy` becomes `error`; existing TVIEWs retain their recorded policy, and new definitions must explicitly permit an uncascaded source when needed. Upgrade as a superuser, or as the extension owner that owns every backing view, then run `tviews.pg_tviews_reregister_all()`. Moved views inherit SELECT grants from their TVIEW tables; view-only grants can be lost. Review custom ACLs and grant SELECT on the corresponding TVIEW table where appropriate.
