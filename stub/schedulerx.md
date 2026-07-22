## Usage

Sources:

- [Official README](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/README.md)
- [Extension SQL for version 1.0](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/schedulerx--1.0.sql)
- [Background-worker implementation](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/scheduler.c)

`schedulerx` is a PostgreSQL 9.5-era background-worker scheduler for stateless SQL jobs. One central scheduler process starts database workers for one-shot, periodic, or notification-driven commands, including multiple parallel workers for the same job. It intentionally does not model dependency graphs or retry workflows.

### Preload and Installation

The library registers a background worker and requests shared memory in its startup hook, so add it to `shared_preload_libraries` and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'schedulerx'
```

```sql
CREATE EXTENSION schedulerx;
```

Extension creation creates `job_schedule`, its enforcement triggers, helper functions, and a `jobscheduler` security label that activates the current database. Use `enable_scheduler` and `disable_scheduler` to change that database-level label.

### Registering Jobs

Register a one-shot job with an absolute start time or a periodic job with a repeat interval:

```sql
SELECT register_job_at(
  clock_timestamp() + interval '1 minute',
  'CALL app.refresh_once()',
  current_user,
  'refresh-once'
);

SELECT register_periodic_job(
  clock_timestamp(),
  interval '5 minutes',
  'CALL app.drain_queue()',
  current_user,
  'drain-queue'
);
```

Both helpers insert a row into `job_schedule`. Important columns include `suspended`, `job_start`, `job_repeat_after`, `job_user`, `job_cmd`, `max_workers`, `run_workers`, `idle_timeout`, `life_time`, `job_start_timeout`, and `job_timeout`. A nonempty `job_name` must be unique.

`register_listener_job` is intended to start work after `NOTIFY` on `listen_channel`, optionally coalescing notifications with `job_start_delay`. The published 1.0 SQL checks an undefined identifier inside that helper; test it on the exact build and prefer a reviewed direct `job_schedule` insert or a local fix before relying on notification jobs.

### Privileges and Limits

Ordinary users may schedule commands as themselves. Creating, updating, or deleting a job for another role requires the current database owner or a superuser. The stored SQL runs with `job_user` privileges, so restrict access to the scheduling table and helper functions and audit command text.

Jobs are stateless executions. The scheduler does not provide “retry N times,” dependency chaining, or “run B when A fails” semantics. Concurrent executions can overlap when timing and worker settings allow, so job SQL must be idempotent or implement its own locking. Capacity is bounded by PostgreSQL background-worker resources as well as per-job settings.

Version 1.0 is old and its control file does not declare the operational preload requirement. Test startup, restart behavior, failover, timeout enforcement, configuration reload, and job recovery on the target PostgreSQL version. Removing the preload entry requires another restart after jobs are disabled and the extension is removed.
