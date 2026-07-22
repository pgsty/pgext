## Usage

Sources:

- [Official README](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/README.md)
- [SQL function definitions](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/sql/migrations/002_functions.sql)
- [Background-worker implementation](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/src/bgworker.c)

`pg_timers` schedules one-shot SQL actions for an exact timestamp or after an interval. A single preloaded background worker wakes for the next due row, executes the action as the user who scheduled it, and records fired, failed, or cancelled status in the `timers.timers` table.

### Preload and Scheduling

Choose the one database the worker will serve, preload the library, restart PostgreSQL, and create the non-relocatable extension in that database:

```conf
shared_preload_libraries = 'pg_timers'
pg_timers.database = 'app'
```

```sql
CREATE EXTENSION pg_timers;

SELECT timers.schedule_at(
  '2026-08-01 09:00:00+08',
  $$DELETE FROM sessions WHERE expires_at < now()$$,
  0,
  30000
);

SELECT timers.schedule_in(
  interval '5 minutes',
  $$SELECT app.expire_session(42)$$
);
```

The optional `shard_key` supports Citus distribution and defaults to 0. `timeout_ms` sets a per-action statement timeout; zero means no limit. Scheduling participates in the caller's transaction, so a rolled-back insert never reaches the worker.

### State and Control

`timers.timers` uses status 0 for pending, 1 for fired, 2 for failed, and 3 for cancelled. `timers.cancel` changes a still-pending row to cancelled and is protected against the worker by row locking. Success records `fired_at`; failure records the error text.

```sql
SELECT timers.cancel(timer_id, shard_key)
FROM timers.timers
WHERE status = 0
  AND id = 42;
```

`timers.fire` and `timers.fire_all_pending` synchronously execute pending actions while ignoring their scheduled time. They are intended for transactional pgTAP tests. A role with those privileges and visibility of the table can fire another user's timer early, so do not grant them to production application roles.

### Security, Reliability, and Capacity

All functions and the table are revoked from `PUBLIC` at installation. Grant scheduling, cancellation, and read access explicitly. Actions run as the recorded `session_user`, not as the background-worker superuser, but anyone allowed to schedule can arrange arbitrary future SQL under their own identity.

Each action runs in a subtransaction. Database changes and the status transition commit atomically with the worker batch, so a crash before commit leaves the timer pending for retry. Non-transactional external side effects can still be duplicated and must be idempotent.

One worker processes actions sequentially in one configured database. `pg_timers.max_timers_per_tick` limits each batch, while `pg_timers.check_interval_ms` optionally caps sleep time; both are reloadable. Slow actions block later timers, and completed rows require a retention policy.

The current source reports development version 0.0.0 even though the README shows an example 0.1.0 image tag. Pin an actual release artifact and reconcile that version before deployment. The documented build matrix covers PostgreSQL 15 through 18; verify the precise artifact, Citus placement, restart/failover behavior, grants, clock assumptions, and workload latency on the target environment.
