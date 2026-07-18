## Usage

Sources:

- [Upstream README](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/README.md)
- [Extension control file](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/pgtelemetry.control)
- [Version 1.7 SQL](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/extension/pgtelemetry--1.7.sql)
- [Generated object documentation](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/doc/pgtelemetry.html)

`pgtelemetry` version `1.7` is a pure-SQL monitoring bundle. Views and functions summarize relation/database/tablespace size, connections and waits, locks, table access and autovacuum, `pg_stat_statements` timing and buffers, WAL rate, replication slots, and standby lag.

### Example

Preload and install `pg_stat_statements` first, then install this extension in its fixed `pgtelemetry` schema:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgtelemetry;
SELECT * FROM pgtelemetry.database_size;
SELECT * FROM pgtelemetry.connections_by_state;
SELECT * FROM pgtelemetry.longest_running_active_queries LIMIT 10;
```

Installation is superuser-only. Several objects expose cluster-wide query text, client addresses, sizes, locks, replication state, and other operational metadata; `wal_telemetry()` also maintains a history table. Grant only the specific views/functions a monitoring role needs, protect exported metrics, set retention for WAL samples, and account for polling cost and PostgreSQL-version changes in statistics catalogs.
