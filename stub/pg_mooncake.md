## Usage

Sources:

- [0.3.0 setup and tutorial](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/README.md)
- [0.3.0 SQL procedures](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/functions.rs)
- [Dependency guard and DuckDB extension loading](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/sql/bootstrap.sql)
- [Scan synchronization](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/duckdb_mooncake.rs)
- [0.3.0 validation and limitations](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/docs/pg-mooncake-0.3.0-validation.md)

`pg_mooncake` 0.3.0 creates read-only columnstore mirrors of PostgreSQL tables for analytical queries. Moonlink replicates committed row changes into Arrow buffers and Parquet/Iceberg storage; DuckDB executes scans and aggregates through `pg_duckdb`. This PGSTY release targets PostgreSQL 14–18 and requires exactly `pg_duckdb` 1.2.0 with DuckDB 2.0.0 and the matching bundled DuckDB Mooncake extension.

### Configure and Enable

Have an administrator configure logical replication and add both libraries to the existing preload list, in this order. Point the repository setting at the installed local DuckDB extension. This example uses the EL PostgreSQL 17 layout:

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
wal_level = logical
duckdb.allow_community_extensions = true
duckdb.allow_unsigned_extensions = true
mooncake.extension_repository = '/usr/pgsql-17/share/extension/pg_mooncake'
```

For another installation, use the directory under the matching `pg_config --sharedir`, followed by `/extension/pg_mooncake`. The tag falls back to the community repository when this setting is absent; it does not automatically select the bundled directory. Unsigned loading is required for this local binary. Ensure sufficient replication slots, WAL senders, and worker processes, then restart PostgreSQL.

```sql
CREATE EXTENSION pg_mooncake CASCADE;
```

Extension creation requires superuser privileges and checks the installed dependency version. Mirror creation opens separate database connections: source DDL must already be committed, and the connecting role must be able to authenticate, read the source, and perform the required replication/publication setup. Run the example in autocommit mode.

### Write to the Source, Query the Mirror

```sql
CREATE TABLE trades (
    id bigint PRIMARY KEY,
    symbol text,
    traded_at timestamp,
    price real
);
CALL mooncake.create_table('trades_analytics', 'trades');

INSERT INTO trades VALUES
    (1, 'AMZN', '2026-11-05 10:00:00', 207),
    (2, 'AMZN', '2026-11-05 10:05:00', 210),
    (3, 'AAPL', '2026-11-05 10:10:00', 203);
UPDATE trades SET price = 209 WHERE id = 1;
DELETE FROM trades WHERE id = 3;

SELECT symbol, avg(price) FROM trades_analytics GROUP BY symbol;
EXPLAIN SELECT avg(price) FROM trades_analytics WHERE symbol = 'AMZN';
```

The helper copies existing rows and starts logical replication. Send `INSERT`, `UPDATE`, and `DELETE` to the source table; the mirror uses `USING mooncake` and is not a writable replacement. Rolled-back source changes are not replicated. Queries can combine mirror data with regular heap tables through supported DuckDB query plans; inspect `EXPLAIN` for a DuckDB scan.

In this tag, scan and snapshot synchronization use the current backend's last-commit LSN. Do not treat this as a general guarantee that every other session's latest write has been ingested, or that a remote source shares the local WAL position. Monitor replication progress for cross-session or remote freshness requirements.

### Procedures and Maintenance

| Interface | Purpose |
| --- | --- |
| `mooncake.create_table(dst, src, src_uri, table_config)` | Create a mirror; the last two arguments default to NULL. Use the exact named argument `src_uri` for a remote PostgreSQL source. |
| `mooncake.list_tables()` | Return registered tables, commit/flush LSNs, and Iceberg warehouse locations. |
| `mooncake.create_snapshot(dst)` | Flush buffered data into an Iceberg snapshot at the requested synchronization point. |
| `mooncake.optimize_table(dst, mode)` | Request compaction; full mode compacts data and indexes. |
| `mooncake.load_files(dst, files)` | Unsupported in 0.3.0; raises SQLSTATE `0A000`. |

```sql
SELECT * FROM mooncake.list_tables();
CALL mooncake.create_snapshot('trades_analytics');
CALL mooncake.optimize_table('trades_analytics', 'full');
```

In-database reads can include buffered Arrow data, while external Iceberg readers see published snapshots. Monitor logical-slot lag and retained WAL as well as mirror storage. Dropping the destination table invokes Moonlink cleanup; verify that its registration and replication slot are removed. Preserve the source data until cleanup and any migration have been checked.

### Compatibility and Upgrade Boundaries

Mirror setup handles replica identity for source tables without a primary key, including existing leaf partitions. Partitions added later also need `REPLICA IDENTITY FULL` before receiving writes. `TRUNCATE` replication and automatic handling of later partition DDL are outside the supported/tested scope. Do not assume arbitrary schema changes or PostgreSQL-specific functions are supported by the mirror.

The v0.3.0 tag provides no in-place SQL upgrade path from 0.2.0. Use a separately validated fresh installation and rebuild mirrors from retained source tables. Replace the PostgreSQL library and matching DuckDB extension together: binaries built for DuckDB 1.4 cannot be reused with DuckDB 2.0. The tagged validation covers local storage and normal restart; S3, crash recovery, and stress behavior are not established by those results.
