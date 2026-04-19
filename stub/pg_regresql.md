## Usage

Sources: [extension README](https://github.com/boringSQL/regresql/blob/master/pg_ext/README.md), [control file](https://github.com/boringSQL/regresql/blob/master/pg_ext/pg_regresql.control), [portable stats article](https://boringsql.com/posts/portable-stats/)

`pg_regresql` is the PostgreSQL planner-hook part of RegreSQL. It makes the planner trust `pg_class` statistics instead of rescaling estimates from the physical file size, which is useful for plan regression testing with injected production statistics.

### Activate the hook

```sql
LOAD 'pg_regresql';

EXPLAIN SELECT ...;
```

For a whole test instance, upstream recommends:

```conf
session_preload_libraries = 'pg_regresql'
```

### What it overrides

The upstream README says the hook runs after `estimate_rel_size()` and replaces planner estimates with catalog values:

- `rel->pages` from `pg_class.relpages`
- `rel->tuples` from `pg_class.reltuples`
- `rel->allvisfrac` from `pg_class.relallvisible / relpages`
- `IndexOptInfo->pages` from index `pg_class.relpages`
- `IndexOptInfo->tuples` from index `pg_class.reltuples`

### Typical workflow

```sql
SELECT pg_restore_relation_stats(
    'schemaname', 'public',
    'relname', 'test_orders',
    'relpages', 123513::integer,
    'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);

LOAD 'pg_regresql';

EXPLAIN SELECT * FROM test_orders WHERE created_at > '2024-06-01';
```

This is meant for reproducing production plans locally and for migration or upgrade plan-regression tests.

### Notes

- The control file currently declares `default_version = '2.0'`.
- The public repository tags visible upstream are still `v2.0.0-alpha*`, so the packaged `2.0.0` target is ahead of a clearly tagged final GitHub release.
- Upstream documents PostgreSQL 14+ compatibility for the extension.
