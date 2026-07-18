## Usage

Sources:

- [Upstream README](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/README.md)
- [SQL usage and operations guide](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/doc/USAGE.md)
- [Extension control file](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/duckpipe-pg/pg_duckpipe.control)
- [Extension Cargo manifest](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/duckpipe-pg/Cargo.toml)

`pg_duckpipe` uses logical decoding to copy existing heap-table rows and stream later WAL changes into DuckLake columnar tables. Install `pg_ducklake` and its bundled `pg_duckdb`, then configure logical WAL and preload both libraries:

```conf
wal_level = logical
shared_preload_libraries = 'pg_duckdb,pg_duckpipe'
```

Restart PostgreSQL, enable the extensions, and add a primary-key table to the default sync group:

```sql
CREATE EXTENSION pg_duckdb;
CREATE EXTENSION pg_ducklake;
CREATE EXTENSION pg_duckpipe;

CREATE TABLE orders (
  id bigint PRIMARY KEY,
  total numeric NOT NULL
);

SELECT duckpipe.add_table('public.orders');
SELECT source_table, state, rows_synced, applied_lsn
FROM duckpipe.status();
```

### Compatibility and WAL safety

Version `1.0.0` has Cargo build features for PostgreSQL 14 through 18, while the top-level README says CI tests PostgreSQL 17 and 18; validate the packaged combination in use. Default upsert mode requires a primary key, while append mode can handle tables without one. Every sync group owns a logical replication slot, so stalled or orphaned groups can retain WAL until disk exhaustion. Configure `max_slot_wal_keep_size`, monitor slot lag, and call `duckpipe.drop_group()` when retiring an empty group.
