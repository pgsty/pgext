## Usage

Sources:

- [Official README](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/README.md)
- [Version 2.0.0 extension SQL](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/sql/pgtt_rsl--2.0.0.sql)
- [Parser-hook implementation](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_rsl.c)
- [Background-worker implementation](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_bgw.c)

`pgtt_rsl` implements Oracle/DB2-style global temporary tables as permanent unlogged backing tables, row-level-security policies, and security-barrier views. Table definitions persist and may live in ordinary schemas, while rows are isolated by session or transaction. Its main tradeoff is avoiding repeated catalog churn at lower data-path performance than native temporary tables.

### Core Workflow

Preload both libraries and restart PostgreSQL before creating or using the extension:

```conf
shared_preload_libraries = 'pgtt_bgw,pgtt_rsl'
pgtt_bgw.naptime = 5
pgtt_bgw.analyze = off
pgtt_bgw.chunk_size = 250000
```

```sql
CREATE EXTENSION pgtt_rsl;

CREATE GLOBAL TEMPORARY TABLE work_items (
    id integer,
    label text
) ON COMMIT PRESERVE ROWS;

INSERT INTO work_items VALUES (1, 'session private');
SELECT * FROM work_items;

DROP TABLE work_items;
```

The explicit helper is useful when generated SQL cannot use the intercepted syntax:

```sql
SELECT pgtt_schema.pgtt_create_table(
    'txn_items',
    'id integer, label text',
    false,
    'public'
);
```

`false` gives transaction-scoped rows (`ON COMMIT DELETE ROWS` behavior); `true` preserves rows for the session.

### Object Index

- `pgtt_schema.pgtt_create_table(name, definition, preserved DEFAULT false, schema DEFAULT 'public')` creates a global temporary table facade.
- `pgtt_schema.pgtt_drop_table(name, schema DEFAULT 'public')` removes it.
- `lsid` and `get_session_id()` identify the owning session internally.
- `pgtt_bgw.naptime`, `pgtt_bgw.analyze`, and `pgtt_bgw.chunk_size` control obsolete-row cleanup.
- `pgtt_schema.pgtt_maintenance(...)` is internal and must be called only by the background worker.

### Operational Notes

Version `2.0.0` is non-relocatable and requires PostgreSQL restart-level preloading. The background worker scans every database containing the extension and removes rows for ended sessions or transactions. Monitor worker health, dead tuples, autovacuum, and cleanup lag; row-level security, views, indexes, and deletion make large workloads slower than native temporary tables.

Backing tables are `UNLOGGED`: their contents are not crash-durable and are not suitable for replicated or recoverable state. The creation helper grants DML on backing tables to `PUBLIC` and relies on forced RLS plus the facade view for isolation; review grants and never access `pgtt_schema.pgtt_*` backing tables directly. Test session disconnect, transaction rollback, restart, failover, pooler behavior, and extension upgrades before migration workloads.
