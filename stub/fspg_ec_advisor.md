## Usage

Sources:

- [Official documentation](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/README.md)
- [Control file](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/fspg_ec_advisor.control)
- [Version 1.0 SQL](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/fspg_ec_advisor--1.0.sql)
- [Decision guide](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/docs/sql-decision-guide.md)

`fspg_ec_advisor` 1.0 is a SQL-only PostgreSQL 17 workload advisor. It stores measurement baselines and operator feedback, then produces resource and scaling advice. SQL routines create advisory records and a delivery outbox; they do not resize infrastructure themselves.

### Enablement and Core Workflow

Install the extension files on a server that permits them. Configure and preload `pg_stat_statements`, restart when changing preload settings, and enable the dependency before the advisor:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION fspg_ec_advisor;
SELECT fspg_ec_advisor.capture();
SELECT * FROM fspg_ec_advisor.query_scores()
ORDER BY total_exec_ms DESC;
SELECT fspg_ec_advisor.latest_advice();
```

Objects reside in the fixed `fspg_ec_advisor` schema. The JSONB capture result and `capture_history` preserve evidence; `record_feedback()` and `recommendation_calibration` connect operator outcomes to recommendations. Run captures over a meaningful observation interval rather than treating a single instantaneous sample as a scaling decision.

### Scheduling and Integration

`run_advisory_cycle()` captures data, evaluates policy, updates advisory events, and queues delivery. An external scheduler can invoke it; optional `pg_cron` support is exposed through `schedule_advisory_cycle()` and its corresponding unschedule routine. Citus telemetry is optional. Policies live in `advisory_policy`; external collectors can supply infrastructure telemetry. Azure Monitor delivery uses an external dispatcher and configured cloud resources, not credentials embedded in the SQL extension.

### Privileges and Reset Boundaries

The control file does not require superuser installation, but the caller still needs database/schema permissions and access to the underlying statistics and operations. The session setting `fspg_ec_advisor.reset_stats` defaults off. Only an explicit opt-in enables `capture_and_reset()` to reset statement statistics after capture; database, I/O, and checkpointer statistics are not reset.

Upstream targets PostgreSQL 17. The package has no shared library or preload requirement of its own. Managed Azure PostgreSQL does not automatically permit arbitrary extension files: this source package is usable there only through a service-approved installation path.
