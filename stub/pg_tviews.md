## Usage

Sources:

- [Official README](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/README.md)
- [Official extension SQL](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/sql/pg_tviews--0.1.0.sql)
- [Official creation API source](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/src/ddl/mod.rs)
- [Official runtime configuration source](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/src/config/mod.rs)

`pg_tviews` maintains derived `tv_*` tables incrementally and transactionally by analyzing a SELECT query, tracking dependencies, and installing refresh triggers on base tables. The catalog/control version is `0.1.0`, while the reviewed README identifies the current code as `0.1.0-beta.11` and explicitly limits it to evaluation rather than mission-critical systems.

### Core Workflow

Create the extension, then use the explicit creation API with a query that returns the required entity key and JSONB data column:

```sql
CREATE EXTENSION pg_tviews;

SELECT pg_tviews_create(
  'post',
  $query$
    SELECT
      p.pk_post,
      p.fk_user,
      jsonb_build_object(
        'id', p.id,
        'title', p.title,
        'author', u.name
      ) AS data
    FROM tb_post AS p
    JOIN tb_user AS u ON u.pk_user = p.fk_user
  $query$
);

SELECT data FROM tv_post WHERE pk_post = 42;
```

The API creates `tv_post`, records dependency metadata, and installs triggers so changes to source rows refresh affected derived rows in the same transaction. The README also documents intercepted `CREATE TABLE tv_* AS SELECT ...` syntax; the explicit function makes extension loading and error reporting clearer.

### Important Operations

- `pg_tviews_create(name, select_sql)` creates a TVIEW; `pg_tviews_drop(name, if_exists, cascade)` removes it.
- `pg_tviews_health_check()` and `pg_tviews_queue_realtime` expose health and queue information.
- `pg_tviews_suspend_triggers()`, `pg_tviews_resume_triggers()`, and `pg_tviews_refresh_all()` support ordered refresh around bulk loads.
- `pg_tviews_recover_after_crash(name)` checks and rebuilds a truncated UNLOGGED TVIEW.
- `pg_tviews.unlogged_by_default` defaults to true in the reviewed source; other GUCs control queue size, propagation depth, caches, metrics, audit logging, and trigger suspension.

### Bulk Load Pattern

```sql
BEGIN;
SELECT pg_tviews_suspend_triggers();
INSERT INTO tb_post SELECT * FROM staging_post;
SELECT pg_tviews_resume_triggers();
COMMIT;

SELECT pg_tviews_refresh_all();
```

### Operational Caveats

TVIEW definitions follow FraiseQL conventions: the query must expose the expected `pk_<entity>` lineage key and a `data` JSONB column, with foreign-key columns used for cascade propagation. The default UNLOGGED storage improves writes but is truncated after a PostgreSQL crash and must be reconstructed from base tables. Treat benchmark claims as upstream measurements, not guarantees; measure trigger latency, cascades, bulk loads, WAL/storage tradeoffs, crash recovery, backup/restore, and schema changes on representative data. The API may change before `1.0.0`; pin the exact build, rehearse upgrades, and avoid mission-critical use until the beta warning is removed and your workload is validated.
