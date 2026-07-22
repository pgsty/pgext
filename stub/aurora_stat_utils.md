## Usage

Sources:

- [AWS documentation for aurora_wait_report](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora_wait_report.html)
- [AWS documentation for aurora_stat_system_waits](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora_stat_system_waits.html)
- [AWS Aurora PostgreSQL extension-version matrix](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`aurora_stat_utils` is an AWS-provided Aurora PostgreSQL extension whose documented user-facing helper, `aurora_wait_report([time])`, measures instance-level wait activity across an interval. It is provider-only: install and use it on Aurora PostgreSQL, not on a community PostgreSQL server.

### Core Workflow

```sql
CREATE EXTENSION aurora_stat_utils;

-- Use the default 10-second observation interval.
SELECT *
FROM aurora_wait_report();

-- Observe a specific 60-second interval.
SELECT *
FROM aurora_wait_report(60);
```

The optional argument is seconds and defaults to 10. The call takes an initial statistics snapshot, waits for the requested interval, takes another snapshot, and returns the difference.

### Result and Interpretation

The report returns `type_name`, `event_name`, `waits`, `wait_time`, `ms_per_wait`, `waits_per_xact`, and `ms_per_xact`. It derives the interval from `aurora_stat_system_waits()` and `pg_stat_database` statistics.

The result covers the connected DB instance and includes activity from all databases on that instance; it is not scoped to the current session or query. On the writer, AWS also reports committed transactions and TPS as a notice. Correlate a hot wait with workload, Performance Insights, and PostgreSQL activity before assigning causality.

`aurora_stat_system_waits()` exposes cumulative waits and microsecond wait time, whereas `aurora_wait_report([time])` reports interval wait time in milliseconds. Do not compare the numeric columns without accounting for that unit and scope difference.

### Operational Boundaries

The function deliberately occupies the calling backend for the observation interval. Use a dedicated diagnostic session, choose a bounded duration, and avoid placing it in latency-sensitive application transactions or frequent monitoring scrapes.

Cumulative system-wait statistics reset when the DB instance restarts, and short intervals can be noisy. Save the observation window and engine/instance context with any report.

AWS currently lists extension version `1.0` across supported Aurora PostgreSQL releases, but availability is tied to the exact Aurora engine minor version and Region rollout. Check the live AWS extension matrix and `pg_available_extension_versions` on the target cluster before automation.
