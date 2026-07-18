## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/README.md)
- [Version 1.0 SQL objects](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots--1.0.sql)
- [Launcher implementation](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots_launcher.c)

`synchronize_logical_slots` is an abandoned EPAS 12-only mechanism for copying logical replication slot state to synchronous streaming standbys. It depends on `dblink`, installs a background-worker launcher, and must be built and installed on the primary and each synchronous standby.

```sql
CREATE EXTENSION synchronize_logical_slots CASCADE;
GRANT EXECUTE ON FUNCTION primary_checkpoint() TO replication_agent;
```

Configure `$libdir/synchronize_logical_slots_launcher` in `shared_preload_libraries` on primary and standbys, enable `hot_standby_feedback`, grant the replication user `pg_read_all_stats`, provide the required `pg_hba.conf` entries, and restart every node. Restrict the replication role and the `primary_checkpoint()` grant carefully.

The project handles synchronous standbys only. After promotion, upstream requires operators to remove unused logical slots manually; stale slots can retain WAL indefinitely. This predates maintained PostgreSQL failover-slot facilities and has no support for current releases. Prefer native, version-supported mechanisms. If legacy EPAS 12 forces its use, test slot advancement, failover, network partitions, repeated promotion, WAL retention, and recovery from inconsistent state under production-like load.
