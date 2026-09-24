## Usage

Sources:

- [README](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/README.md)
- [Control file](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/pg_lockwatch.control)
- [Cargo.toml](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/Cargo.toml)
- [src/lib.rs](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/src/lib.rs)

`pg_lockwatch` samples blocking backends and wait chains, estimates lock-contention risk, and emits alerts. The reviewed 0.1.0 source is experimental; its scores are monitoring signals, not predictions with a proven accuracy guarantee.

### Core Workflow

Add `pg_lockwatch` to `shared_preload_libraries`, select the worker database using `lockwatch.database`, and restart PostgreSQL. Install as a superuser in that database.

```sql
CREATE EXTENSION pg_lockwatch;
SELECT lockwatch_sample_now();
SELECT * FROM lockwatch_risks;
SELECT * FROM lockwatch_history;
LISTEN lockwatch_alert;
```

### Configuration

`lockwatch.sample_interval_ms` sets the sampling interval and `lockwatch.risk_threshold` sets the alert threshold. Weights include `lockwatch.weight_velocity`, `lockwatch.weight_duration`, `lockwatch.weight_lock_mode`, and `lockwatch.weight_cascade_depth`. A notification is also recorded in the history table. The documented release workflow targets PostgreSQL 15–18.

### Limits

Shared-memory arrays have a compiled capacity; raising `lockwatch.max_tracked_blockers` or `lockwatch.history_window` cannot enlarge them. Fingerprints hash raw query text rather than normalized queries. Duration baselines start with the first observed blocker, and `lockwatch_history.resolved_as` is not back-filled.

Upstream asks users to validate live integration before real workloads. No license file or Cargo license declaration was found in the reviewed revision; do not infer redistribution permission from public source visibility.
