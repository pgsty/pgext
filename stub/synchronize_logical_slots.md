## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/README.md)
- [Version 1.0 SQL objects](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots--1.0.sql)
- [Launcher implementation](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots_launcher.c)

`synchronize_logical_slots` is an abandoned EPAS 12-only mechanism for copying logical replication slot state to synchronous streaming standbys. It depends on `dblink`, installs a background-worker launcher, and must be built and installed on the primary and each synchronous standby.

### Installation boundary

```sql
CREATE EXTENSION synchronize_logical_slots CASCADE;
GRANT EXECUTE ON FUNCTION primary_checkpoint() TO replication_agent;
```

Configure `$libdir/synchronize_logical_slots_launcher` in `shared_preload_libraries` on primary and standbys. The launcher connects to `postgres` by default; set `sync_logical_slot.database` to the database that contains the extension and the logical slots when it differs. Reserve capacity in `max_worker_processes`, enable `hot_standby_feedback` on standbys, grant the replication user `pg_read_all_stats`, provide the required `pg_hba.conf` entries, and restart every node. The worker polls approximately every 60 seconds.

### Destructive and security boundaries

On a synchronous standby, reconciliation creates or advances missing slots, drops and recreates a slot whose plugin differs, and drops local logical slots absent from the primary's list. Do not mix unrelated logical slots on a managed standby, and test the exact slot inventory before enabling the worker. After promotion, upstream still requires manual removal of unused slots; stale slots can retain WAL indefinitely.

`primary_checkpoint()` is `SECURITY DEFINER` and calls unqualified extension functions. Keep its grant limited to a narrowly controlled replication role and harden the function search path before use. The synchronization function catches errors and returns them as text, so launcher activity alone is not proof of successful slot advancement; monitor slot LSNs and server logs.

The project supports synchronous standbys only and predates maintained PostgreSQL failover-slot facilities. Prefer native, version-supported mechanisms. If legacy EPAS 12 forces its use, test slot advancement, failover, network partitions, unrelated-slot preservation, repeated promotion, WAL retention, and recovery from inconsistent state under production-like load.
