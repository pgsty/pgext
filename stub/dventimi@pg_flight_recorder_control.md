## Usage

Sources:

- [Official database.dev package page](https://database.dev/dventimi/pg_flight_recorder_control)

`dventimi@pg_flight_recorder_control` — Control functions for pg_flight_recorder. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder_control";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `flight_recorder._get_table_autovacuum_settings` is an extension function.
- `flight_recorder.compute_recommended_scale_factor(p_relid OID)` is an extension function and returns `TABLE`.
- `flight_recorder.dead_tuple_growth_rate(p_relid OID, p_window INTERVAL)` is an extension function and returns `NUMERIC`.
- `flight_recorder.dead_tuple_trend(p_relid OID, p_window INTERVAL)` is an extension function and returns `NUMERIC`.
- `flight_recorder.time_to_budget_exhaustion(p_relid OID, p_budget BIGINT)` is an extension function and returns `INTERVAL`.
- `flight_recorder.vacuum_control_mode(p_relid OID)` is an extension function and returns `TABLE`.
- `flight_recorder.vacuum_control_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TABLE`.
- `flight_recorder.vacuum_diagnostic(p_relid OID)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.bloat_report` is an extension function.
- `flight_recorder_reporting.dead_tuple_growth_rate(p_relid OID, p_window INTERVAL)` is an extension function and returns `NUMERIC`.
- `flight_recorder_reporting.estimate_table_bloat(p_relid OID DEFAULT NULL)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.oid_consumption_rate(p_window INTERVAL)` is an extension function and returns `NUMERIC`.
- `flight_recorder_reporting.table_size_growth_rate(p_relid OID, p_window INTERVAL)` is an extension function and returns `NUMERIC`.
- `flight_recorder_reporting.time_to_budget_exhaustion(p_relid OID, p_budget BIGINT)` is an extension function and returns `INTERVAL`.

### Requirements and Caveats

- The catalog records version `2.26.3`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
