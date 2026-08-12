## Usage

Sources:

- [pg_clickhouse v0.10.0 README](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/README.md)
- [pg_clickhouse v0.10.0 reference](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/doc/pg_clickhouse.md)
- [pg_clickhouse v0.10.0 tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/doc/tutorial.md)
- [pg_clickhouse v0.10.0 changelog](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/CHANGELOG.md)
- [pg_clickhouse v0.10.0 control file](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/pg_clickhouse.control)
- [pg_clickhouse 0.3 to 0.10 upgrade SQL](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/sql/pg_clickhouse--0.3--0.10.sql)
- [Pigsty pg_clickhouse package matrix](https://pgext.cloud/ext/pg_clickhouse)

`pg_clickhouse` 0.10.0 exposes ClickHouse tables to PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream targets PostgreSQL 13 or later and ClickHouse 23.3 or later; current Pigsty packages cover PostgreSQL 14–18. No preload is required for normal use; `session_preload_libraries` and `shared_preload_libraries` are optional connection-startup optimizations.

### Connect PostgreSQL to ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (
  driver 'binary',
  host 'localhost',
  dbname 'taxi',
  compression 'lz4'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

The required `driver` option is `binary` or `http`. Common server options include `host`, `port`, `dbname`, `compression`, `secure`, and `min_tls_version`; user mappings accept `user` and `password`. Version 0.10 deprecates and ignores `fetch_size` because both drivers now stream the same Native format.

`IMPORT FOREIGN SCHEMA` supports `LIMIT TO (...)` and `EXCEPT (...)`. Imported mixed-case identifiers remain quoted and must be referenced with matching quotes.

### Query and Write Foreign Tables

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM taxi.logs
GROUP BY node_id;

INSERT INTO taxi.nodes(node_id, name)
VALUES (9, 'west-node');

COPY taxi.nodes(node_id, name) FROM STDIN;
```

`SELECT`, `EXPLAIN`, prepared statements, `INSERT`, and `COPY` operate on foreign tables. In version 0.10 the binary driver flushes inserts in bounded 64 MiB batches, so `COPY` is no longer merely expanded into one statement per row. Use `EXPLAIN (VERBOSE)` to inspect remote SQL and verify which filters, joins, aggregates, and functions were pushed down.

### Direct Query and Command APIs

Version 0.10 adds typed arbitrary-query and command interfaces:

```sql
GRANT EXECUTE ON FUNCTION clickhouse_query(text, text) TO analyst;
GRANT EXECUTE ON PROCEDURE clickhouse_perform(text, text) TO operator;

SELECT *
FROM clickhouse_query(
  'taxi_srv',
  'SELECT region, count() FROM taxi GROUP BY region'
) AS t(region text, n bigint);

CALL clickhouse_perform(
  'taxi_srv',
  'OPTIMIZE TABLE taxi.nodes FINAL'
);

SELECT clickhouse_server_version('taxi_srv');
```

`clickhouse_query(server, sql)` returns rows using the caller-provided column definition, while `clickhouse_perform(server, sql)` discards any result. Both can run arbitrary remote SQL, so `EXECUTE` is revoked from `PUBLIC` and should be granted narrowly. `clickhouse_raw_query()` is deprecated in favor of these interfaces.

### Pushdown and Session Settings

Version 0.10 expands aggregate and function pushdown, improves aggregate execution over mixed local and foreign partitions, and fixes several PostgreSQL NULL-semantics mismatches. Subquery pushdown requires ClickHouse 25.8 or later; older servers evaluate those subqueries locally.

The default `pg_clickhouse.session_settings` preserves PostgreSQL-compatible behavior, including `join_use_nulls = 1`, `group_by_use_nulls = 1`, `final = 1`, and `transform_null_in = 0`. If it is overridden, retain the settings needed by the workload—especially `transform_null_in = 0`, which is required for safe `IN` pushdown.

### Upgrade and Operational Boundaries

```sql
ALTER EXTENSION pg_clickhouse UPDATE TO '0.10';
SELECT pgch_version();
```

The extension SQL version is `0.10`, while `pgch_version()` reports the full library version `0.10.0`. An installation upgraded from SQL version `0.3` must run `ALTER EXTENSION` after the new files are installed.

If `pg_clickhouse` is placed in `session_preload_libraries`, new sessions load it automatically. If it is placed in `shared_preload_libraries`, changing the library requires a PostgreSQL restart. Neither setting is mandatory, unlike extensions that register postmaster hooks.

Lightweight `UPDATE` and `DELETE` remain outside the documented write surface. Treat direct remote SQL as privileged, test pushdown with production-shaped NULL and type cases, and validate both PostgreSQL and ClickHouse versions before relying on a version-gated optimization.
