## Usage

Sources:

- [pgfr_record v2.29.2 README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/README.md)
- [pgfr_record control file](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/extension.control)
- [pg_flight_recorder v2.29.2 reference](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 release notes](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_record is the collection half of PostgreSQL Flight Recorder. It periodically samples PostgreSQL activity, waits, locks, replication, vacuum, and related health data into bounded in-database buffers, then retains snapshots for incident analysis. Use it when short-lived database conditions must survive long enough for later diagnosis.

### Install and Enable Recording

pgfr_record requires pg_cron:

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    SELECT pgfr_record.enable();

enable() installs and schedules the collector jobs. It also reports configuration warnings; review them rather than treating a successful call as proof that every metric is being collected.

### Inspect Recorder Health

    SELECT * FROM pgfr_record.health_check();
    SELECT * FROM pgfr_record.ring_buffer_health();
    SELECT * FROM pgfr_record.list_profiles();

Use set_mode or apply_profile to select the intended collection profile before an incident:

    SELECT pgfr_record.set_mode('normal');

The available collection modes are normal, light, emergency, and kill. Profile names and sampling intervals can evolve, so list the installed profiles rather than hard-coding an undocumented name.

### Recorded Data Index

- deltas: interval changes for cumulative PostgreSQL counters.
- recent_activity and recent_waits: sampled sessions and wait events.
- recent_locks and recent_idle: lock and idle-session observations.
- recent_replication and recent_vacuum: replication and maintenance state.
- archiver_status: WAL archive health.
- snapshot and ring-buffer tables: retained history used by pgfr_analyze.

Many working buffers are UNLOGGED to reduce write amplification. They are not crash-durable and are not replicated like ordinary logged tables; durable snapshots provide the longer-lived analysis surface.

### Administration Functions

- pgfr_record.enable(): create or activate scheduled collectors.
- pgfr_record.disable(): stop scheduled collection.
- pgfr_record.health_check(): report collector and configuration health.
- pgfr_record.set_mode(...): change collection mode.
- pgfr_record.apply_profile(...): apply a predefined profile.
- pgfr_record.list_profiles(): enumerate available profiles.
- pgfr_record.ring_buffer_health(): inspect capacity and retention pressure.
- pgfr_record.cleanup(...): remove retained history according to the API.

### Retention and Overhead

The default design keeps short ring-buffer history and longer durable snapshots, commonly around 7 and 30 days depending on the installed profile. Verify actual table sizes, job schedules, and retention settings in the installed version.

The recorder creates roughly ten pg_cron jobs. pg_cron.log_run can generate thousands of rows per day; disable that logging or purge cron history when the extra audit trail is unnecessary. Sampling also adds SQL, storage, and catalog traffic, so measure overhead on the target workload.

Version 2.29.2 handles managed-service roles that cannot UPDATE cron.job: jobs can still be scheduled, while the optional nodename normalization is skipped with a warning.

### Caveats

- pg_stat_statements enriches several analyses but is optional; enable and size it separately when needed.
- Collection cannot reconstruct time periods that were never sampled. Enable and validate the recorder before an incident.
- UNLOGGED buffers can be truncated after crash recovery.
- Recorder tables can contain query text, role names, client data, and operational details. Apply appropriate privileges and retention controls.
