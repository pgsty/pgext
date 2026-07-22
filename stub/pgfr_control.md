## Usage

Sources:

- [Official v2.28.1 control README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.28.1/_control/README.md)
- [Official v2.28.1 install SQL](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.28.1/_control/install.sql)
- [Official v2.28.1 release tree](https://github.com/dventimisupabase/pg_flight_recorder/tree/v2.28.1)
- [Current official repository tree](https://github.com/dventimisupabase/pg_flight_recorder/tree/main)

`pgfr_control` 2.28.1 is the historical SQL-only control component of PostgreSQL Flight Recorder. It reads snapshots collected by `pgfr_record` to diagnose vacuum behavior, estimate bloat and growth, monitor OID consumption, and recommend autovacuum scale factors. It reports recommendations but never applies settings automatically.

### Install the Tagged Component

The upstream v2.28.1 source ships `pgfr_control` as an install script under `_control`, not as a standalone `CREATE EXTENSION` package. Install and enable the matching `_record` component first, then run the tagged control script with `psql`.

```sql
\i _record/install.sql
SELECT pgfr_record.enable();

\i _control/install.sql

SELECT *
FROM pgfr_control.vacuum_diagnostic('public.orders'::regclass);

SELECT *
FROM pgfr_control.vacuum_control_report(now() - interval '1 hour', now());
```

Use files from the same v2.28.1 tag. The control installer checks for the `pgfr_record.config` schema-version row and aborts if the recorder core is absent.

### Main Reports

- Vacuum control: `vacuum_control_mode(oid)` returns `normal`, `catch_up`, or `safety`; `compute_recommended_scale_factor(oid)` proposes a setting; `vacuum_diagnostic(oid)` classifies health; `vacuum_control_report(start, end)` covers all observed tables.
- Dead tuples: `dead_tuple_growth_rate(oid, interval)`, `dead_tuple_trend(oid, interval)`, and `time_to_budget_exhaustion(oid, budget)` use snapshot history.
- Bloat and size: `estimate_table_bloat(oid)`, `bloat_report(interval)`, and `table_size_growth_rate(oid, interval)` derive estimates without running `pgstattuple`.
- OIDs: `oid_consumption_rate(interval)` and `time_to_oid_exhaustion()` estimate cluster-wide consumption.

`vacuum_diagnostic` emits states including `NOT_SCHEDULED`, `RUNNING_BUT_LOSING`, `BLOCKED`, and `HEALTHY`. Treat these as diagnostic signals to investigate, not commands to change settings blindly.

### Version and Operational Boundaries

Results depend on the availability, interval, and retention of `pgfr_record` snapshots. Sparse history returns null or low-confidence rates; bloat and exhaustion values are estimates rather than physical inspection or guarantees. Review recommendations against workload, lock, transaction-age, and table-specific reloption context.

Version 2.28.1 is a pinned historical layout. The current upstream tree no longer ships `_control` as a separate component, so do not mix these scripts with a later recorder schema. Catalog packaging may wrap the SQL differently, but the authoritative tagged source workflow above does not justify fabricating `CREATE EXTENSION pgfr_control` for a source install.
