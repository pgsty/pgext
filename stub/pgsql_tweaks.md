## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/README.md)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/pgsql_tweaks.control)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/makefile)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/sql/out/versions/pgsql_tweaks--1.1.0.sql)

`pgsql_tweaks` 1.1.0 provides DBA-oriented helper functions and views in the `pgsql_tweaks` schema. The upstream test matrix covers PostgreSQL 14 through PostgreSQL 19rc1; Pigsty packages it for PostgreSQL 14-18.

### Enablement

```sql
CREATE EXTENSION pgsql_tweaks;
```

### Data Type Check Functions

```sql
SELECT pgsql_tweaks.is_date('2024-01-15');
SELECT pgsql_tweaks.is_time('10:30:00');
SELECT pgsql_tweaks.is_timestamp('2024-01-15 10:30:00');
SELECT pgsql_tweaks.is_real('3.14');
SELECT pgsql_tweaks.is_double_precision('3.14');
SELECT pgsql_tweaks.is_numeric('3.14');
SELECT pgsql_tweaks.is_bigint('9223372036854775807');
SELECT pgsql_tweaks.is_integer('42');
SELECT pgsql_tweaks.is_smallint('42');
SELECT pgsql_tweaks.is_boolean('true');
SELECT pgsql_tweaks.is_json('{"a":1}');
SELECT pgsql_tweaks.is_jsonb('{"a":1}');
SELECT pgsql_tweaks.is_hex('FF');
```

### System Information Functions

```sql
SELECT pgsql_tweaks.pg_schema_size('public');
SELECT * FROM pgsql_tweaks.role_inheritance();
```

### Encoding Functions

```sql
SELECT pgsql_tweaks.is_encoding('UTF8');
SELECT pgsql_tweaks.is_latin1('abc');
SELECT pgsql_tweaks.return_not_part_of_latin1('abc');
SELECT pgsql_tweaks.replace_latin1('abc', '?');
SELECT pgsql_tweaks.return_not_part_of_encoding('abc', 'UTF8');
SELECT pgsql_tweaks.replace_encoding('abc', 'UTF8', '?');
```

### Aggregate Functions

- `gap_fill`, for filling gaps in time series.
- `array_min`, `array_max`, `array_avg`, and `array_sum`.

### Format And Conversion Functions

```sql
SELECT pgsql_tweaks.date_de(current_date);
SELECT pgsql_tweaks.datetime_de(now());
SELECT pgsql_tweaks.to_unix_timestamp(now());
SELECT pgsql_tweaks.hex2bigint('FF');
```

### Utility Functions

```sql
SELECT pgsql_tweaks.is_empty('');
SELECT pgsql_tweaks.array_trim(ARRAY['a', '', 'b']);
SELECT pgsql_tweaks.get_markdown_doku_by_schema('pgsql_tweaks');
```

### System Information Views

- `pg_db_views`, `pg_foreign_keys`, `pg_functions`, `pg_active_locks`.
- `pg_table_matview_infos`, `pg_object_ownership`, `pg_partitioned_tables_infos`.
- `pg_unused_indexes`, `pg_bloat_info`, `pg_table_bloat`, `pg_missing_indexes`.
- `pg_role_permissions`, `pg_role_infos`.

### Statistic And Monitoring Views

- `statistics_top_ten_query_times`, `top_ten_query_average_time_in_seconds`.
- `statistics_top_ten_time_consuming_queries`, `statistics_top_ten_memory_usage_queries`.
- `statistics_top_ten_called_queries`, `statistics_top_ten_rows_returned_queries`.
- `statistics_top_ten_shared_block_hits_queries`, `statistics_top_ten_wal_records_generated_queries`.
- `statistics_query_activity`.
- `monitoring_wal`, `wal_archiving`, `monitoring_active_locks`, `monitoring_replication`.
- `monitoring_database_conflicts`, `monitoring_blocked_and_blocking_activity`.
- `monitoring_follower_wal_status`, `monitoring_vacuum`.

### Caveats

- Some views/functions depend on additional PostgreSQL supplied modules; upstream lists `pg_stat_statements` and `pgstattuple`.
- This distribution is pure SQL and needs no preloaded shared library. After installing matching files, use `ALTER EXTENSION pgsql_tweaks UPDATE TO '1.1.0'` for an existing installation. Helpers still require the relevant object or statistics privileges.
