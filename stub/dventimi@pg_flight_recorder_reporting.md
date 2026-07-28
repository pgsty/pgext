## Usage

Sources:

- [Official database.dev package page](https://database.dev/dventimi/pg_flight_recorder_reporting)

`dventimi@pg_flight_recorder_reporting` — Reporting and analysis functions for pg_flight_recorder. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder_reporting";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `flight_recorder_reporting._diagnose_regression_causes` is an extension function.
- `flight_recorder_reporting.anomaly_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.blast_radius(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.blast_radius_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TEXT`.
- `flight_recorder_reporting.capacity_report(p_time_window INTERVAL DEFAULT interval '24 hours')` is an extension function and returns `TEXT`.
- `flight_recorder_reporting.capacity_summary(p_time_window INTERVAL DEFAULT interval '24 hours')` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.check_alerts(p_lookback_interval INTERVAL DEFAULT '1 hour')` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.config_at(p_timestamp TIMESTAMPTZ, p_category TEXT DEFAULT NULL)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.config_changes(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.config_health_check()` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.db_role_config_at(p_timestamp TIMESTAMPTZ, p_database TEXT DEFAULT NULL, p_role TEXT DEFAULT NULL, p_prefix TEXT DEFAULT NULL)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.db_role_config_changes(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.db_role_config_summary()` is an extension function and returns `TABLE`.
- `flight_recorder_reporting.detect_query_storms(p_lookback INTERVAL DEFAULT NULL, p_threshold_multiplier NUMERIC DEFAULT NULL)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The catalog records version `2.26.3`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
