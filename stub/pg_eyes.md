## Usage

Sources:

- [Official README](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/README.md)
- [Version 1.4 SQL](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/sql/pg_eyes--1.4.sql)
- [Extension control file](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/pg_eyes.control)

`pg_eyes` 1.4 is a pure-SQL monitoring toolkit in the fixed `eyes` schema. It exposes activity, lock, size, catalog, settings, replication, and `pg_stat_statements` data as ready-to-scrape metrics and views. Upstream's 1.x branch targets PostgreSQL 9.6, so verify every object on newer catalogs before use.

### Setup and Basic Queries

`pg_stat_statements` must be installed and loaded correctly, normally through `shared_preload_libraries`, before its statistics are useful. The extension itself requires superuser installation and is not relocatable.

```conf
shared_preload_libraries = 'pg_stat_statements'
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_eyes;

SELECT stat_name, stat_value
FROM eyes.get_activity();

SELECT * FROM eyes.active_connections;
SELECT * FROM eyes.active_queryes;
SELECT * FROM eyes.blocked_queryes;
SELECT * FROM eyes.top_queryes LIMIT 20;
```

The spelling `active_queryes`, `blocked_queryes`, and `top_queryes` is part of the upstream SQL API.

### Important Objects

- `eyes.get_activity()` returns metric-name/value rows for sessions, transactions, database and background-writer counters, WAL, replication, and response time.
- `eyes.get_activity(text)` runs custom scalar metric SQL stored in the table `eyes.get_activity` for the selected group.
- `eyes.get_pg_stat_activity()` and `eyes.get_pg_stat_statements()` are `SECURITY DEFINER` wrappers owned by `postgres`.
- `eyes.db_objects`, `eyes.db_tables`, `eyes.db_tables_size`, `eyes.db_indexes`, `eyes.db_indexes_size`, `eyes.db_attributes`, `eyes.db_functions`, and `eyes.db_settings` expose inventory and configuration details.

Many counters are cumulative and must be converted to rates by the monitoring system. Metrics ending in `_time` use milliseconds according to the upstream README.

### Security and Compatibility

The definer-rights wrappers intentionally expose complete activity and statement text to non-superusers who can execute them. Query text, client addresses, settings paths, and function definitions can be sensitive; explicitly audit and narrow `EXECUTE` and schema privileges after installation.

`eyes.get_activity(text)` executes the `stat_query` column as the function owner. Never grant untrusted roles write access to the `eyes.get_activity` table, and review every custom metric before insertion. Older upgrade paths may also retain definer-rights administrative functions; inspect effective grants with `\df+ eyes.*`.

Version 1.4 SQL refers to PostgreSQL 9.6-era catalog columns such as the historical `pg_stat_statements.total_time`. It is not evidence of compatibility with all later majors. Test installation and every selected view against the exact PostgreSQL and `pg_stat_statements` versions, or port the SQL before production deployment.
