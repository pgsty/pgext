## Usage

Sources:

- [Official upstream README](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/README.md)
- [Official extension control file (pg_index_stats.control)](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/pg_index_stats.control)
- [Official extension SQL (pg_index_stats--0.2.sql)](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/pg_index_stats--0.2.sql)

`pg_index_stats` — Lightweight extension for PostgreSQL that generates extended statistics based on index definitions. It introduces dependency of the statistics on the corresponding index. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE TABLE test(x integer, y integer);
CREATE INDEX ON test (x,y);
CREATE EXTENSION pg_index_stats;
SELECT pg_index_stats_rebuild();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_index_stats_build(idxname text, mode text DEFAULT 'mcv, ndistinct')` is an extension function and returns `boolean`.
- `pg_index_stats_rebuild()` is an extension function and returns `integer`.
- `pg_index_stats_remove()` is an extension function and returns `integer`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
