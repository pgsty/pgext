## Usage

Sources: [README](https://github.com/ClickHouse/pg_clickhouse/blob/master/README.md), [reference](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/pg_clickhouse.md), [tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/tutorial.md), [v0.2.0 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.2.0)

`pg_clickhouse` runs analytics queries on ClickHouse from PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream documents PostgreSQL 13+ and ClickHouse 23+ support.

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

User mapping options:

- `user`
- `password`

### Common operations

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.2';
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` also supports `LIMIT TO (...)` and `EXCEPT (...)`. The reference warns that imported mixed-case identifiers are double-quoted in PostgreSQL and must be queried with quotes.

### Version and pushdown notes

- The reference documents separate library and extension versions; `pgch_version()` was added in release `v0.2.0`.
- Patch-only releases update the library without requiring `ALTER EXTENSION`.
- Release `v0.2.0` added more pushdown for arrays, regex functions, `split_part()`, array operators, and current date/time expressions, plus the `pg_clickhouse.pushdown_regex` setting.

### Caveats

- This is positioned upstream as an analytics-first extension; the roadmap still lists broader DML support as future work.
- For full examples, follow the official tutorial, which creates a ClickHouse `taxi` database, imports it through `IMPORT FOREIGN SCHEMA`, and queries the resulting foreign tables.
