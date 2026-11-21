## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.0/README.md)
- [pg_regresql.control](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.0/pg_regresql.control)
- [pg_regresql--2.2.sql](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.0/pg_regresql--2.2.sql)

`pg_regresql` makes query planning use catalog statistics instead of rescaling them from physical file size. This supports plan regression tests in databases with representative statistics but little data. Distribution 2.2.0 installs SQL version 2.2 in the fixed `pg_catalog` schema.

### Activate the Planner Hook

The hook supports PostgreSQL 15 or later. Load it as a superuser in the test session:

```sql
LOAD 'pg_regresql';
CREATE TABLE test_orders (created_at date);
EXPLAIN SELECT * FROM test_orders WHERE created_at > DATE '2024-06-01';
```

For new sessions on a test instance, add the library to the existing session preload list:

```ini
session_preload_libraries = 'pg_regresql'
```

The hook replaces table and index page/tuple estimates with `pg_class` values and computes visibility from catalog statistics. Inject useful statistics before comparing plans; an empty table alone is not a production workload.

### Restore Statistics

On PostgreSQL 16 and 17, create the extension before loading statistics emitted by a recent pg_dump:

```sql
CREATE EXTENSION pg_regresql;
SELECT pg_restore_relation_stats(
    'schemaname', 'public', 'relname', 'test_orders',
    'relpages', 123513::integer, 'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);
```

Version 2.2 adds `pg_restore_extended_stats()` and `pg_clear_extended_stats()` for those servers, and supports the JSON input emitted by pg_dump 19 for `pg_ndistinct` and `pg_dependencies`. On PostgreSQL 16 remove the dump's unsupported transaction_timeout setting. PostgreSQL 18 already supplies relation/attribute restoration in core; extended restoration is in core from 19. The extension does not install these backports on 18 or later.

### Upgrade and Cleanup

From SQL version 2.0, upstream requires dropping and recreating the extension. From 2.1, use `ALTER EXTENSION pg_regresql UPDATE`. Review dependencies before removing it. The extension is not relocatable and its install requires a superuser.

The replacement input functions named `regresql_pg_*_in` are deliberately detached from the extension on 16/17. `DROP EXTENSION` leaves them behind so the installed casts keep working; dropping the extension is therefore not a complete return to the original catalog state. Use this tool in controlled planning tests, and do not assume compatibility with other planner hooks has been tested.
