## Usage

Sources:

- [Official README](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/README.md)
- [Version 1.0 load script](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table--1.0.sql)
- [ProcessUtility hook implementation](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table.c)

`pg_force_unlogged_create_table` installs a utility hook that changes ordinary `CREATE TABLE` and `CREATE TABLE AS` commands into `UNLOGGED` table creation. It is intended for environments that deliberately prefer load speed over crash durability and replication coverage.

### Core Workflow

Creating the extension executes `LOAD`, so the hook becomes active immediately in that session. Put the library in `shared_preload_libraries` and restart PostgreSQL when every new session should inherit the behavior.

```sql
CREATE EXTENSION pg_force_unlogged_create_table;
CREATE TABLE staging (id bigint, payload jsonb);

SELECT relpersistence
FROM pg_class
WHERE oid = 'staging'::regclass;
```

The expected `relpersistence` is `u`. Explicit `CREATE TEMP TABLE` and `CREATE UNLOGGED TABLE` commands are left unchanged, and `CREATE MATERIALIZED VIEW` is not affected. The extension defines no SQL functions or types.

### Durability Boundary

Unlogged relations are truncated after an unclean shutdown and are not written to WAL for physical replication. Indexes on an unlogged table are unlogged as well. Use this only for disposable, reconstructible data, and verify that backup, failover, logical processing, and recovery procedures do not assume these tables are durable.

Session loading affects only statements executed after the library is loaded. Dropping the extension removes its membership objects but does not convert existing unlogged tables back to logged tables; convert or recreate those relations explicitly if durability requirements change.
