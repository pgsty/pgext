


## Usage

Sources: [official README](https://github.com/vibhorkum/pg_background/blob/master/README.md), [v2.0 release notes](https://github.com/vibhorkum/pg_background/releases/tag/v2.0), [migration guide](https://github.com/vibhorkum/pg_background/blob/v2.0/docs/MIGRATION.md).

`pg_background` executes SQL inside PostgreSQL background worker processes. Workers run independent transactions inside the server, which is useful for asynchronous maintenance, autonomous side effects, bounded long-running tasks, and progress-tracked jobs.

Version 2.0 makes the unsuffixed API canonical. The old `_v2` names remain deprecated aliases through the 2.x line, but new code should use names such as `pg_background_launch`, `pg_background_result`, and `pg_background_run`.

### One-Shot Execution

Use `pg_background_run` when the SQL has side effects and you only need execution metadata:

```sql
CREATE EXTENSION pg_background;

SELECT completed, has_error, sqlstate, error_message,
       row_count, command_tag, elapsed_ms, timed_out
FROM pg_background_run(
  'INSERT INTO audit_log(ts, who) VALUES (clock_timestamp(), current_user)',
  queue_size := 0,
  timeout_ms := 30000,
  label := 'audit-login'
);
```

### Launch And Fetch Results

Use the launch/result pattern when the background SQL returns rows:

```sql
SELECT * FROM pg_background_launch(
  'SELECT count(*) FROM large_table',
  queue_size := 65536,
  label := 'count-large-table'
) AS h;

SELECT * FROM pg_background_result(h.pid, h.cookie) AS (count bigint);
```

Results can be consumed once. Keep both `pid` and `cookie`; the cookie protects later calls from PID reuse.

### Fire And Forget

For side effects where no result rows need to be consumed:

```sql
SELECT * FROM pg_background_submit(
  $$VACUUM (ANALYZE) public.events$$,
  queue_size := 65536,
  label := 'vacuum-events'
);
```

### Core API

- `pg_background_launch(sql, queue_size, label)` launches a worker and returns `pg_background_handle(pid, cookie)`.
- `pg_background_submit(sql, queue_size, label)` launches fire-and-forget work and returns a handle.
- `pg_background_result(pid, cookie)` consumes result rows once.
- `pg_background_result_info(pid, cookie)` returns completion and row-count metadata without consuming rows.
- `pg_background_error_info(pid, cookie)` returns structured SQL error details.
- `pg_background_wait(pid, cookie, timeout_ms DEFAULT 0)` waits for completion; `timeout_ms <= 0` blocks indefinitely.
- `pg_background_cancel(pid, cookie, grace_ms DEFAULT 0)` requests cooperative cancellation.
- `pg_background_detach(pid, cookie)` stops tracking a worker while letting it continue.
- `pg_background_outcome(pid, cookie)` returns a combined status snapshot without raising on missing state.
- `pg_background_list` and `pg_background_activity` are monitoring views; `pg_background_stats()` returns session counters.

Convenience helpers include `pg_background_run_query`, `pg_background_drain`, `pg_background_wait_any`, `pg_background_cancel_by_label`, and `pg_background_purge`.

### Progress Reporting

Report progress from inside the worker SQL, then poll it from the launcher:

```sql
SELECT * FROM pg_background_launch($$
  SELECT pg_background_report_progress(0, 'starting');
  SELECT pg_sleep(1);
  SELECT pg_background_report_progress(100, 'done');
$$) AS h;

SELECT * FROM pg_background_get_progress(h.pid, h.cookie);
```

`pg_background_report_progress` is the 2.0 name; the earlier `pg_background_progress` name was hard-renamed.

### GUCs And Loading

`pg_background` does not require `shared_preload_libraries`. Preloading is optional and mainly useful when you want its GUCs available before the extension is first loaded in a session.

```sql
SET pg_background.max_workers = 10;
SET pg_background.default_queue_size = '256KB';
SET pg_background.worker_timeout = '5min';
```

- `pg_background.max_workers` defaults to `16`.
- `pg_background.default_queue_size` defaults to `65536` bytes.
- `pg_background.worker_timeout` defaults to `0`, meaning no execution timeout.

### Caveats

- Pigsty packages `pg_background` 2.0 for PostgreSQL 14-18; upstream 2.0 also validates PostgreSQL 19 beta.
- Upgrades from pre-1.8 installs must first reach the 1.8/1.10 release line before updating to 2.0.
- The original v1 PID-only API was removed. Unsuffixed names now have cookie-protected semantics and return/use `(pid, cookie)`.
- `pg_background_cancel_v2_grace` and `pg_background_wait_v2_timeout` are folded into `pg_background_cancel(..., grace_ms)` and `pg_background_wait(..., timeout_ms)`.
- `pg_background_status_v2` was removed; use `pg_background_outcome(pid, cookie)`.
