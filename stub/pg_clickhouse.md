## Usage

Sources: [README](https://github.com/ClickHouse/pg_clickhouse/blob/main/README.md), [reference](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/pg_clickhouse.md), [tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/tutorial.md), [v0.3.2 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.3.2), [changelog](https://github.com/ClickHouse/pg_clickhouse/blob/main/CHANGELOG.md)

`pg_clickhouse` runs analytics queries on ClickHouse from PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream documents PostgreSQL 13+ and ClickHouse 23+ support; the current version is 0.3.2.

### Connect PostgreSQL to ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (driver 'binary', host 'localhost', dbname 'taxi');

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

Server options documented upstream:

- `driver`: required, `binary` or `http`
- `host`
- `port`
- `dbname`
- `fetch_size`: HTTP streaming batch size; `0` disables streaming
- `compression`: binary-driver compression, `none`, `lz4`, or `zstd`; v0.3.2 defaults to `lz4`
- `secure`: explicit TLS mode, `on`, `off`, or `auto`
- `min_tls_version`: minimum TLS protocol version for both binary and HTTP drivers

User mapping options:

- `user`
- `password`

### Common operations

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.3';
SELECT pgch_version();
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` also supports `LIMIT TO (...)` and `EXCEPT (...)`. The reference warns that imported mixed-case identifiers are double-quoted in PostgreSQL and must be queried with quotes.

### Query and write notes

`SELECT`, `EXPLAIN`, prepared statements, `INSERT`, and `COPY` can operate on `pg_clickhouse` foreign tables. Use `EXPLAIN (VERBOSE)` to inspect the remote SQL that will be sent to ClickHouse.

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM logs
GROUP BY node_id;

INSERT INTO nodes(node_id, name, region, arch, os)
VALUES (9, 'west-node', 'us-west-2', 'amd64', 'Linux');
```

`COPY` into a foreign table is documented, but upstream notes that it currently uses `INSERT` statements because FDW batch insertion is still future work.

### Version and pushdown notes

- The reference documents separate library and extension versions; `pgch_version()` reports the loaded library version.
- Release `v0.3.2` is binary-only for existing SQL version `0.3`; it does not require `ALTER EXTENSION UPDATE` when upgrading from another 0.3 build.
- Release `v0.3.2` adds the `compression`, `secure`, and `min_tls_version` server options, adds `regexp_match()` pushdown, and adds PostgreSQL 19beta1 distribution support.
- Release `v0.3.2` also fixes regular-expression flag pushdown and avoids pushing down regex calls when the regex argument is not a constant.
- Release `v0.3.1` is binary-only for existing SQL version `0.3`; it does not require `ALTER EXTENSION UPDATE` when upgrading from `v0.3.0`.
- Release `v0.3.1` replaces the client library with `ClickHouse/clickhouse-c`, streams result blocks, and adds rectangular multidimensional array support for `SELECT` and `INSERT`.
- Release `v0.3.1` also adds pushdown for `pg_re2` 0.3.0 functions such as `re2extractallgroupshorizontal`, `re2extractallgroupsvertical`, `re2regexpquotemeta`, and `re2splitbyregexp`, and fixes `UInt16` casts to PostgreSQL `int4`.
- Release `v0.3.0` uses SQL version `0.3`; run `ALTER EXTENSION pg_clickhouse UPDATE TO '0.3'` to apply its SQL-level privilege change.
- Release `v0.3.0` adds pushdown for `re2` functions, `soundex()`, two-argument `levenshtein()`, compatible `to_char(timestamp[tz], fmt)`, selected builtin functions, and JSON/JSONB path operations.
- ClickHouse `JSON` maps to PostgreSQL `jsonb` or `json`; the binary driver's `JSON` mapping requires ClickHouse 24.10 or later.
- `pg_clickhouse.pushdown_regex` controls built-in PostgreSQL regex pushdown. Upstream recommends considering the `re2` extension for regex work that should push down directly.

### Caveats

- In 0.3.0, `clickhouse_raw_query(text, text)` is no longer executable by `PUBLIC`; grant it only to roles that need ad-hoc ClickHouse queries.
- This is positioned upstream as an analytics-first extension; lightweight `DELETE` and `UPDATE` support remain on the roadmap.
- For full examples, follow the official tutorial, which creates a ClickHouse `taxi` database, imports it through `IMPORT FOREIGN SCHEMA`, and queries the resulting foreign tables.
