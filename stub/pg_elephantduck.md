## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/README.md)
- [Extension control file](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/pg_elephantduck.control)
- [Table-access-method implementation](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/src/tam.rs)
- [Cargo metadata](https://github.com/sonesuke/pg_elephantduck/blob/0427d744ab3e3b8e2fbb95328888e3551942eae7/Cargo.toml)

`pg_elephantduck` is an experimental PostgreSQL table access method backed by embedded DuckDB and Parquet files. It offers columnar storage and limited `WHERE` pushdown.

```sql
CREATE EXTENSION pg_elephantduck;
SET elephantduck.path = '/srv/postgresql/elephantduck';
CREATE TABLE sample USING elephantduck
AS SELECT generate_series(1, 10000) AS number;
SELECT number FROM sample WHERE number < 10;
```

The storage directory defaults to a location under `PGDATA` or `/tmp`; configure a durable, PostgreSQL-owned path explicitly. Include that path in backup, restore, replication, access-control, and disk-capacity planning rather than assuming heap-table WAL semantics.

Version 0.0.0 is active prototype code for PostgreSQL 16–17. Upstream directs users to open issues for limitations and does not claim complete DML, crash recovery, replication, vacuum, locking, or transactional parity with heap tables. Use only with disposable data until those behaviors are verified on the exact build.
