## Usage

- Source: [Codeberg repo](https://codeberg.org/kop/pg_isok), [documentation home](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/index.html), [doc source](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/isok.xml)
- Isok is a query-centered monitoring extension for PostgreSQL. It reports changes to previously seen questionable data patterns, not just the existence of the rows.

```sql
CREATE SCHEMA isok;
CREATE EXTENSION pg_isok SCHEMA isok;
```

## Core Workflow

The extension centers on two tables:

- `ISOK_QUERIES`, which stores the monitoring queries
- `ISOK_RESULTS`, which stores the discovered issues and their resolution state

Run the monitor with `run_isok_queries()`:

```sql
SELECT * FROM run_isok_queries();
SELECT * FROM run_isok_queries($$VALUES ('new_countries')$$) AS problems;
```

Rows in `ISOK_RESULTS` can be resolved or deferred so later runs no longer report them as active problems.

## Notes

The documentation describes Isok as a "soft trigger" style tool for data cleanup and integrity review. It installs on PostgreSQL 10 or later and can be built as pure SQL for managed environments.
