## Usage

Sources:

- [0.3.4 setup and upgrades](https://github.com/pgsty/pg_mooncake/blob/v0.3.4/README.md)
- [0.3.4 SQL procedures](https://github.com/pgsty/pg_mooncake/blob/v0.3.4/src/functions.rs)
- [0.3.4 dependency guard](https://github.com/pgsty/pg_mooncake/blob/v0.3.4/src/sql/bootstrap.sql)
- [0.3.4 release](https://github.com/pgsty/pg_mooncake/releases/tag/v0.3.4)

`pg_mooncake` 0.3.4 maintains read-only columnstore mirrors of PostgreSQL tables through logical replication into Arrow and Parquet/Iceberg storage. This PGSTY release supports PostgreSQL 15–19 and accepts `pg_duckdb` 1.2.1 or later. The reference build uses 1.2.2 with DuckDB 2.0.1 and Mooncake scanner 0.0.5.

### Configure and Enable

An administrator must retain existing preload entries, add both libraries in this order and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
wal_level = logical
duckdb.allow_unsigned_extensions = true
mooncake.extension_repository = '/usr/pgsql-17/share/extension/pg_mooncake'
```

Use the absolute directory from the target server's `pg_config --sharedir`, followed by `/extension/pg_mooncake`; the example is an EL PostgreSQL 17 path. The local scanner works with community extensions disabled. Install using a superuser in a fresh session:

```sql
CREATE EXTENSION pg_mooncake CASCADE;
```

The dependency version is checked. Configure enough replication slots, WAL senders and workers. Source connections need authentication, table access and replication/publication privileges.

### Mirror and Inspect

Run mirror creation as a standalone command in a writable session, after source DDL commits:

```sql
CREATE TABLE trades (id bigint PRIMARY KEY, symbol text, price real);
CALL mooncake.create_table('trades_analytics', 'trades');
INSERT INTO trades VALUES (1, 'AMZN', 207), (2, 'AMZN', 210);
SELECT symbol, avg(price) FROM trades_analytics GROUP BY symbol;
SELECT * FROM mooncake.list_tables();
CALL mooncake.create_snapshot('trades_analytics');
CALL mooncake.optimize_table('trades_analytics', 'full');
```

Write to the source table; the mirror is not writable. Use the `src_uri` argument for a remote source. Fresh reads and snapshots use the source server's WAL position. External Iceberg readers see published snapshots, while local queries can also see buffered data.

### Durability and Upgrades

Version 0.3.3 fixes snapshots that could acknowledge a source commit before Iceberg persisted it. From 0.3.1 or 0.3.2, install matching libraries and scanner, restart, refresh a cached older scanner from the configured repository, then reconnect before updating:

```sql
SELECT duckdb.raw_query(format('FORCE INSTALL mooncake FROM %L', current_setting('mooncake.extension_repository')));
ALTER EXTENSION pg_mooncake UPDATE TO '0.3.4';
```

Take a fresh snapshot of every mirror before using it for backup or migration. From the WAL-based 0.3.0, first stop source writers, compare every mirror against its source, and verify snapshots under the old binaries. Stop PostgreSQL cleanly and preserve complete storage backups and the old `_wal` offline before starting the new worker. Follow the tagged per-database procedure; a new snapshot cannot recover old WAL already lost. Rebuild unverifiable mirrors from authoritative sources.

### Storage and Lifecycle

For physical backups or streaming standbys, set `mooncake.data_directory` outside `PGDATA` before creating mirrors, give each instance private storage and back it up separately. Logical dumps do not preserve Moonlink data or replication state. Protect source connection strings and metadata files.

Before renaming or dropping a database with mirrors, drop its mirrors, commit and run `SELECT mooncake._reconcile_drops()` while the worker is available. Existing leaf partitions receive full replica identity during setup; later partition changes require separate handling. Monitor slot lag and retained WAL, and verify cleanup before discarding source data.

### Dependency Compatibility in 0.3.4

The version guard now accepts `pg_duckdb` 1.2.1 and later; 1.2.1 and 1.2.2 are tested with DuckDB 2.0.1 and scanner 0.0.5. For a newer dependency, verify its PostgreSQL interfaces and provide scanner 0.0.5 built for its DuckDB engine and platform. Update the dependency's SQL extension to match its library. Direct SQL upgrades are provided from 0.3.0–0.3.3; the earlier WAL, snapshot and scanner-cache maintenance requirements still apply.
