## Usage

Sources:

- [Official upstream README](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/README.md)
- [Official extension control file (pg_turso.control)](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/extension/pg_turso.control)
- [Official extension SQL (pg_turso--1.0.sql)](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/extension/pg_turso--1.0.sql)

`pg_turso` — Postgres output plugin for replicating data to Turso. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION pg_turso;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `turso_generate_create_table_for_mv(mv_name text)` is an extension function and returns `text`.
- `turso_generate_create_table_for_table(table_name text)` is an extension function and returns `text`.
- `turso_migrate_mv_schema(mv_name text)` is an extension function and returns `text`.
- `turso_migrate_table_schema(table_name text)` is an extension function and returns `text`.
- `turso_schedule_mv_replication(view_name text, refresh_interval text)` is an extension function and returns `integer`.
- `turso_schedule_table_replication(table_name text, refresh_interval text)` is an extension function and returns `integer`.
- `turso_send(url text, token text, data text)` is an extension function and returns `text`.
- `turso_replicate_mv` is an extension procedure.
- `turso_replicate_table` is an extension procedure.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `pg_cron`.
- The control file marks the extension as relocatable.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
