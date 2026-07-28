## Usage

Sources:

- [Official database.dev package page](https://database.dev/dventimi/pg_flight_recorder)

`dventimi@pg_flight_recorder` — Server-side flight recorder for PostgreSQL performance monitoring. Use it when collecting or interpreting the corresponding PostgreSQL statistics. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `flight_recorder._check_and_adjust_mode` is an extension function.
- `flight_recorder._check_catalog_ddl_locks` is an extension function.
- `flight_recorder._check_circuit_breaker` is an extension function.
- `flight_recorder._check_schema_size` is an extension function.
- `flight_recorder._check_statements_health` is an extension function.
- `flight_recorder._collect_config_snapshot` is an extension function.
- `flight_recorder._collect_db_role_config_snapshot` is an extension function.
- `flight_recorder._collect_index_stats` is an extension function.
- `flight_recorder._collect_table_stats` is an extension function.
- `flight_recorder._get_config` is an extension function.
- `flight_recorder._get_ring_buffer_slots` is an extension function.
- `flight_recorder._get_ring_retention_interval` is an extension function.
- `flight_recorder._get_setting_from_snapshots` is an extension function.
- `flight_recorder._has_pg_stat_statements` is an extension function.

### Requirements and Caveats

- The catalog records version `2.26.3`.
- Install the confirmed extension dependencies first: `pg_cron`.
- Upstream material contains an explicit deprecation boundary.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
