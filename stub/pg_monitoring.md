## Usage

Sources:

- [Official README](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/README)
- [Version 0.2 SQL](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring--0.2.sql)
- [Release notes](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/Changes)
- [Extension control file](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring.control)

`pg_monitoring` 0.2 is a historical pure-SQL collection of replication-lag and aggregate load functions. It was written against pre-PostgreSQL-10 statistics and WAL function names and contains data-quality defects in its table-load calculations. It is useful as a reviewable example, not as a current monitoring contract.

### Installed Functions

```sql
CREATE EXTENSION pg_monitoring;

SELECT * FROM pg_monitoring_lag_info();
SELECT pg_monitoring_time_since_replay();
SELECT * FROM pg_monitoring_raw_table_load();
SELECT * FROM pg_monitoring_load_across_tables();
SELECT * FROM pg_monitoring_load_across_databases();
```

- `pg_monitoring_lag_info()` reads `pg_stat_replication` and estimates byte lag.
- `pg_monitoring_time_since_replay()` reports seconds since the last replayed transaction on a standby.
- `pg_monitoring_raw_table_load()` aggregates table scan and modification counters.
- `pg_monitoring_load_across_tables()` mutates a one-row baseline table and returns a delta since the prior call.
- `pg_monitoring_load_across_databases()` returns a current cluster-wide sum of backends, transactions, reads, and hits; it is not a delta.

### Version and Correctness Defects

The lag function references `replay_location` and `pg_current_xlog_location()`, names replaced by `replay_lsn` and `pg_current_wal_lsn()` in PostgreSQL 10. It also converts the high WAL component with a multiplier based on 255 rather than the 256 required for a 32-bit hexadecimal boundary. Do not expect the function to execute or report correct bytes on current PostgreSQL.

`pg_monitoring_raw_table_load()` assigns its delete counter from `n_tup_upd` instead of `n_tup_del`, so deletes are wrong and updates are counted twice in the total. The delta function is explicitly documented as unsafe for concurrent callers: every poll replaces one shared baseline, so two collectors, a failed scrape, a statistics reset, or a restart produces misleading intervals. The install table `pg_monitoring_last_run` is not used by these functions.

### Safe Migration

Do not expose these results as alerts or service-level indicators without rewriting and testing the SQL for the target major version. Prefer direct current statistics views and LSN helpers such as `pg_wal_lsn_diff`, and let the monitoring system keep per-target timestamps, baselines, reset detection, and rate calculations. Grant a collector only the required monitoring privileges, verify primary-versus-standby NULL behavior, and test failover, statistics resets, counter wrap/reset, concurrent scrapes, and version upgrades against known workloads.
