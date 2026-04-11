
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_clickhouse;
> CREATE SERVER taxi_srv FOREIGN DATA WRAPPER clickhouse_fdw
>   OPTIONS(driver 'binary', host 'localhost', dbname 'taxi');
> CREATE USER MAPPING FOR CURRENT_USER SERVER taxi_srv OPTIONS (user 'default');
> IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
> ```
>
> Sources: [README](https://github.com/ClickHouse/pg_clickhouse), [Reference](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/pg_clickhouse.md), [Tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/tutorial.md)

`pg_clickhouse` is a PostgreSQL extension for running analytics queries on ClickHouse directly from PostgreSQL, including a foreign data wrapper and query pushdown support. The current upstream docs state support for PostgreSQL 13+ and ClickHouse 23+.

## Getting Started

The upstream project documents two common starting points:

- use the published Docker image `ghcr.io/clickhouse/pg_clickhouse:18`
- build with `make` / `make install` or install from PGXN

Once installed, enable the extension:

```sql
CREATE EXTENSION pg_clickhouse;
```

Or install it into a dedicated schema:

```sql
CREATE SCHEMA ch;
CREATE EXTENSION pg_clickhouse WITH SCHEMA ch;
```

## Connecting to ClickHouse

The reference docs show the normal flow as:

```sql
CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (driver 'binary', host 'localhost', dbname 'taxi');

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

Documented server options include:

- `driver`, required, either `binary` or `http`
- `dbname`
- `fetch_size`
- `host`
- `port`

## What the Docs Emphasize

The README positions pg_clickhouse around transparent pushdown for analytic workloads. It links both a tutorial and a SQL reference:

- the tutorial walks through connecting PostgreSQL to a ClickHouse sample database and querying imported tables
- the reference documents extension lifecycle commands, foreign server options, and SQL objects exposed by the extension

The project README also includes TPC-H benchmark examples showing where pushdown can significantly reduce runtime.

## Operational Notes

The reference docs describe versioning separately for:

- the library version, visible via `pgch_version()` or `pg_get_loaded_modules()`
- the extension version tracked by PostgreSQL catalogs and extension upgrade scripts

Minor and major upgrades can require `ALTER EXTENSION pg_clickhouse UPDATE`.
