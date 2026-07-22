## Usage

Sources:

- [Official README](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/README.md)
- [Version 1.0 extension SQL](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower--1.0.sql)
- [Apply worker implementation](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower_apply.c)

`pg_follower` is an educational logical-replication experiment that reconstructs a limited set of SQL changes on a downstream PostgreSQL server. Upstream explicitly says not to use it in production; it is not a replacement for built-in logical replication.

### Core Workflow

Set `wal_level = logical` and restart the upstream server, then create the extension in both databases. Start the worker from the downstream with a libpq connection string for the upstream:

```sql
-- Run in both upstream and downstream databases.
CREATE EXTENSION pg_follower;

-- Run only on the downstream.
SELECT start_follow('host=upstream.example user=replicator dbname=app');
```

The worker creates a temporary logical replication slot on the upstream, receives SQL-like messages from the extension's output plugin, and applies them through SPI on the downstream.

### Object Index

- `start_follow(connection_string text) RETURNS void` launches the downstream background worker.
- `detect_ddl() RETURNS event_trigger` captures supported table DDL.
- Event trigger `test_trigger` fires after `CREATE TABLE` and `DROP TABLE`.

### Operational Notes

Version `1.0` is relocatable and declares no preload requirement. The extension must be installed at both ends, and the upstream role needs logical-replication connectivity. Protect passwords in the connection string from SQL text, logs, and monitoring.

Only `INSERT`, `CREATE TABLE`, `DROP TABLE`, and `TRUNCATE` are documented as supported. `UPDATE` and `DELETE` are not supported; table constraints and options are ignored. `UNLOGGED`, `TEMPORARY`, partitioned/inherited tables, and typed-table clauses are rejected. The temporary slot, hand-built SQL stream, background worker, and partial DDL model leave major correctness, failover, security, and recovery gaps. Use isolated test data only.
