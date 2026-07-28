## Usage

Sources:

- [Official upstream README](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/README.md)
- [Official extension control file (approx_count.control)](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/approx_count.control)
- [Official extension SQL (approx_count--1.0.sql)](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/sql/approx_count--1.0.sql)

`approx_count` — Fast approximate row counts for tables and indexes in PostgreSQL 14+, read from pg_class.reltuples instead of disk-heavy exact COUNT(*) scans, with a governed ANALYZE refresh when the statistics go stale. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION approx_count;

-- On a ~300M-row table an exact count scans for minutes:
SELECT count(*) FROM events;     -- minutes, heavy I/O
-- approx_count reads the planner's cached estimate:
SELECT approx_count.approx_count('events');   -- ~0.2 ms
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `IF` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
