## Usage

Sources:

- [Official README](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/README.md)
- [Control file](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect.control)
- [Version 1.0 SQL](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect--1.0.sql)
- [Planner hook and configuration](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect.c)

`drift_detect` detects structural changes in registered query plans and stores events while notifying listeners. For example, it can report a switch from an index scan to a sequential scan. The documented source snapshot declares extension version 1.0 and PostgreSQL 14–17 support; the repository has no published release.

### Enable and Register a Query

Add the library to `shared_preload_libraries`, preserving existing entries, and select the one database the worker will monitor. Restart PostgreSQL, then create the extension there as an administrator. PL/pgSQL is needed by the management functions.

```ini
shared_preload_libraries = 'drift_detect'
drift_detect.database = 'appdb'
```

```sql
CREATE EXTENSION drift_detect;
```

For an existing orders table, register the parameterized SQL as the application sends it, supplying representative parameters for the baseline plan. The returned integer identifies the tracked query.

```sql
SELECT drift.track_query(
    'get user orders',
    'SELECT * FROM orders WHERE user_id = $1 ORDER BY created_at DESC LIMIT 20',
    '{"params": [1]}'::jsonb
);
LISTEN plan_drift;
SELECT * FROM drift.drift_summary();
```

### Objects and Plan Matching

`drift.track_query` registers a query and its initial baseline; its optional fourth argument chooses a notification channel. `drift.untrack_query` disables tracking. `drift.acknowledge_drift` accepts an event ID and the acknowledging user’s label; it marks the event handled rather than changing the query. `drift.drift_summary` lists unacknowledged events.

`drift.tracked_queries`, `drift.plan_snapshots`, and `drift.plan_drift_events` store registrations, baselines, and events. `drift.reload_cache` reloads enabled registrations. `drift.compute_query_id` and `drift.fingerprint` are diagnostic helpers. Matching uses PostgreSQL’s normalized query ID; structural fingerprints describe plan nodes, relations, indexes, and strategies, excluding cost and row estimates.

### Configuration and Boundaries

`drift_detect.enabled` defaults to true. `drift_detect.snapshot_interval_secs` defaults to 60 seconds, `drift_detect.notify_channel` to `plan_drift`, and `drift_detect.min_table_growth_ratio` to 3.0. These settings take effect on configuration reload. `drift_detect.max_tracked` defaults to 500; changing it or `drift_detect.database` requires restart.

The worker monitors one database, and deployment must permit installing and preloading the library. SQL management functions use caller privileges; arrange schema, table, and function access explicitly for non-administrators. Notifications signal a plan-shape change, not proof of a performance regression. Keep the persisted events for investigation and manage their retention according to workload; this snapshot does not document automatic retention.
