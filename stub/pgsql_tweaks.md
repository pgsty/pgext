


## Usage

> [pgsql_tweaks: PostgreSQL views and functions for DBA daily jobs](https://codeberg.org/pgsql_tweaks/pgsql_tweaks)

All objects are created in the `pgsql_tweaks` schema. Full documentation: [rtfm.pgsql-tweaks.org](https://rtfm.pgsql-tweaks.org)

### Data Type Check Functions

```sql
SELECT pgsql_tweaks.is_date('2024-01-15');       -- true
SELECT pgsql_tweaks.is_integer('42');             -- true
SELECT pgsql_tweaks.is_numeric('3.14');           -- true
SELECT pgsql_tweaks.is_json('{"a":1}');           -- true
SELECT pgsql_tweaks.is_jsonb('{"a":1}');          -- true
SELECT pgsql_tweaks.is_boolean('true');            -- true
SELECT pgsql_tweaks.is_timestamp('2024-01-15 10:30:00');  -- true
SELECT pgsql_tweaks.is_hex('FF');                 -- true
```

### System Information Functions

```sql
SELECT pgsql_tweaks.pg_schema_size('public');     -- schema size in bytes
```

### Aggregate Functions

- `gap_fill` -- fill gaps in time series
- `array_min`, `array_max`, `array_avg`, `array_sum` -- array aggregates

### Conversion Functions

```sql
SELECT pgsql_tweaks.to_unix_timestamp(now());
SELECT pgsql_tweaks.hex2bigint('FF');
```

### Utility Functions

```sql
SELECT pgsql_tweaks.is_empty('');                 -- true
SELECT pgsql_tweaks.array_trim(ARRAY['a','','b']);
```

### System Information Views

- `pg_db_views`, `pg_foreign_keys`, `pg_functions`, `pg_active_locks`
- `pg_table_matview_infos`, `pg_object_ownership`, `pg_unused_indexes`
- `pg_bloat_info`, `pg_missing_indexes`, `pg_role_permissions`

### Monitoring Views

- `monitoring_wal`, `monitoring_active_locks`, `monitoring_replication`
- `monitoring_database_conflicts`, `monitoring_vacuum`
- `statistics_top_ten_query_times`, `statistics_query_activity`
