## Usage

Sources:

- [Official upstream README](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/README.md)
- [Official extension control file (pg_query_stats.control)](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/pg_query_stats.control)
- [Official extension SQL (pg_query_stats--1.0.0.sql)](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/pg_query_stats--1.0.0.sql)

`pg_query_stats` — **pg_query_stats** is a lightweight PostgreSQL extension that collects and exposes execution statistics for SQL queries. It helps developers and DBAs monitor query performance with minimal overhead and simple architecture. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_query_stats;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_query_stats()` is an extension function and returns `SETOF`.
- `pg_query_stats_reset()` is an extension function and returns `void`.
- `pg_query_stats` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
