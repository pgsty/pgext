## Usage

Sources:

- [PGXN README](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/README.md)
- [PGXN reference](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/doc/pg_clickhouse.md)
- [PGXN tutorial](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/doc/tutorial.md)
- [PGXN changelog](https://api.pgxn.org/dist/pg_clickhouse/0.11.1/META.json)
- [PGXN control file](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/pg_clickhouse.control)
- [0.10 to 0.11 upgrade SQL](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/sql/pg_clickhouse--0.10--0.11.sql)

`pg_clickhouse` 0.11.1 exposes ClickHouse tables to PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream targets PostgreSQL 14 or later and ClickHouse 23.3 or later; package availability is tracked separately. No preload is required for normal use; `session_preload_libraries` and `shared_preload_libraries` are optional connection-startup optimizations.

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

The required `driver` option is `binary` or `http`. Common server options include `host`, `port`, `dbname`, `compression`, `secure`, and `min_tls_version`; user mappings accept `user` and `password`. Version 0.11 removes `fetch_size`; both drivers stream the Native format.

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

`clickhouse_query(server, sql)` returns rows using the caller-provided column definition, while `clickhouse_perform(server, sql)` discards any result. Both can run arbitrary remote SQL, so `EXECUTE` is revoked from `PUBLIC` and should be granted narrowly. `clickhouse_raw_query()` is removed in 0.11; migrate to these interfaces.

### Pushdown and Session Settings

Version 0.10 expands aggregate and function pushdown, improves aggregate execution over mixed local and foreign partitions, and fixes several PostgreSQL NULL-semantics mismatches. Subquery pushdown requires ClickHouse 25.8 or later; older servers evaluate those subqueries locally.

The default `pg_clickhouse.session_settings` preserves PostgreSQL-compatible behavior, including `join_use_nulls = 1`, `group_by_use_nulls = 1`, `final = 1`, and `transform_null_in = 0`. If it is overridden, retain the settings needed by the workload—especially `transform_null_in = 0`, which is required for safe `IN` pushdown.

### Upgrade and Operational Boundaries

```sql
ALTER EXTENSION pg_clickhouse UPDATE TO '0.11';
SELECT pgch_version();
```

The extension SQL version is `0.11`, while `pgch_version()` reports library version `0.11.1`. Run `ALTER EXTENSION` after installing the matching files to remove the obsolete function from the SQL catalog.

If `pg_clickhouse` is placed in `session_preload_libraries`, new sessions load it automatically. If it is placed in `shared_preload_libraries`, changing the library requires a PostgreSQL restart. Neither setting is mandatory, unlike extensions that register postmaster hooks.

Lightweight `UPDATE` and `DELETE` remain outside the documented write surface. Treat direct remote SQL as privileged, test pushdown with production-shaped NULL and type cases, and validate both PostgreSQL and ClickHouse versions before relying on a version-gated optimization.

Version 0.11 rejects text or JSON with invalid database encoding by default. The server option `encoding_check` accepts `fail`, `replace`, `remove` or `truncate`; the latter three alter returned data and require a deliberate policy. Import mappings also improve large integers, intervals and Array/Tuple/Map/Nested types. Existing foreign-table definitions do not automatically follow new inference rules; review their column mappings.

### 0.11.1 Settings and CloudNativePG

Version 0.11.1 documents `settings` options on servers and foreign tables. Table settings override matching server settings; session settings override both, and conflicting table values fall back with a NOTICE. The release also documents CloudNativePG ExtensionConfiguration and Database resources for mounting the matching image.
