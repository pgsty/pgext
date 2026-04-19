## Usage

Source: [README](https://github.com/datasentinel/pg_datasentinel/blob/main/README.md), [Release 1.0](https://github.com/datasentinel/pg_datasentinel/releases/tag/1.0)

`pg_datasentinel` adds observability views on top of PostgreSQL activity, maintenance, temporary-file, checkpoint, wraparound, and container resource data. It must be preloaded because it allocates shared memory and hooks into activity logging.

### Required setup

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_datasentinel'
log_autovacuum_min_duration = 0
log_temp_files = 0
log_checkpoints = on

CREATE EXTENSION pg_datasentinel;
```

### Main views

- `ds_stat_activity`: extends `pg_stat_activity` with backend RSS, optional PSS, temp-file bytes, and `plan_id` on PostgreSQL 18+.
- `ds_container_resources`: reports cgroup CPU and memory limits plus current usage.
- `ds_wraparound_risk`: estimates XID and MXID ETA to aggressive vacuum and wraparound from hourly snapshots.
- `ds_xid_snapshots`: raw snapshot history used by `ds_wraparound_risk`.
- `ds_vacuum_activity`, `ds_analyze_activity`, `ds_tempfile_activity`, `ds_checkpoint_activity`: shared-memory ring buffers for maintenance and checkpoint events.
- `ds_activity_summary`: one-row summary of ring-buffer occupancy and timestamps.

### Useful GUCs

- `pg_datasentinel.enabled`: enables or disables capture.
- `pg_datasentinel.max_entries`: ring-buffer capacity per activity stream; restart required.
- `pg_datasentinel.maintenance_force_verbose`: adds `VERBOSE` to manual `VACUUM` and `ANALYZE` so they are captured.
- `pg_datasentinel.ignore_system_schemas`: suppresses `pg_catalog` and `information_schema` noise.
- `pg_datasentinel.enable_pss_memory`: reads `/proc/*/smaps_rollup` for per-backend PSS.

### Caveats

- PostgreSQL 15+ is required upstream.
- Memory and temp-file enrichment depends on Linux `/proc`; container metrics depend on cgroups.
- Vacuum and analyze parsing depends on English log keywords, so translated server message locales are not fully supported.
- `plan_id` is only populated on PostgreSQL 18+, and is most useful when paired with the official `pg_store_plans` fork linked from the README.
