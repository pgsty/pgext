## Usage

Sources:

- [Upstream README](https://github.com/citusdata/cstore_fdw/blob/90e22b62fbee6852529104fdd463f532cf7a3311/README.md)
- [Extension control file](https://github.com/citusdata/cstore_fdw/blob/90e22b62fbee6852529104fdd463f532cf7a3311/cstore_fdw.control)

`cstore_fdw` stores PostgreSQL foreign tables in a columnar format intended for batch-loaded analytical data. It supports column projection, min/max skip indexes, and optional `pglz` compression.

The library must be preloaded before the extension is created:

```conf
shared_preload_libraries = 'cstore_fdw'
```

Restart PostgreSQL after changing that setting, then create the FDW objects:

```sql
CREATE EXTENSION cstore_fdw;
CREATE SERVER cstore_server FOREIGN DATA WRAPPER cstore_fdw;

CREATE FOREIGN TABLE measurements (
    measured_at timestamptz,
    sensor_id bigint,
    reading double precision
)
SERVER cstore_server
OPTIONS (compression 'pglz', stripe_row_count '150000');
```

Load data in batches with `COPY` or `INSERT INTO ... SELECT ...`, then run `ANALYZE`. Upstream version `1.7` supports PostgreSQL 9.3 through 12 and does not support `UPDATE`, `DELETE`, or single-row inserts.

### Deprecated status

Upstream has moved columnar storage into Citus using PostgreSQL's table-access-method API. It recommends migrating existing `cstore_fdw` data to Citus columnar tables, which add features such as streaming replication, rollback, and easier `pg_upgrade`. Use `cstore_fdw` mainly to operate or migrate legacy installations.
