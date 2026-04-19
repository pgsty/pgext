## Usage

Sources: [official README](https://github.com/vibhorkum/pg_background/blob/master/README.md), [v1.9.2 release](https://github.com/vibhorkum/pg_background/releases/tag/v1.9.2)

`pg_background` executes SQL in PostgreSQL background worker processes. The worker runs inside the server and uses its own transaction, which makes it useful for asynchronous maintenance, autonomous side effects, and long-running tasks that should not block the caller.

```sql
CREATE EXTENSION pg_background;

SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table',
  65536,
  'count-large-table'
) AS h;

SELECT * FROM pg_background_result_v2(h.pid, h.cookie) AS (count bigint);
```

### Core API

- `pg_background_launch_v2(sql, queue_size, label)` launches a tracked worker and returns `(pid, cookie)`.
- `pg_background_submit_v2(sql, queue_size, label)` is fire-and-forget for side-effect SQL.
- `pg_background_result_v2(pid, cookie)` consumes the worker result set once.
- `pg_background_wait_v2(...)` and `pg_background_wait_v2_timeout(...)` wait for completion.
- `pg_background_cancel_v2(...)` stops execution; `pg_background_detach_v2(...)` stops tracking but lets work continue.
- `pg_background_list_v2()`, `pg_background_stats_v2()`, and `pg_background_get_progress_v2(...)` expose worker state and progress.

### Typical Patterns

Run maintenance without holding the client session open:

```sql
SELECT * FROM pg_background_submit_v2(
  'VACUUM (ANALYZE) public.events',
  65536,
  'vacuum-events'
);
```

Use an autonomous side effect from application SQL:

```sql
SELECT * FROM pg_background_submit_v2(
  $$INSERT INTO audit_log(ts, event) VALUES (clock_timestamp(), 'job queued')$$
);
```

### GUCs And Security

- `pg_background.max_workers` limits concurrent workers per session.
- `pg_background.default_queue_size` controls the shared-memory queue size.
- `pg_background.worker_timeout` sets an execution timeout; `0` means no limit.
- The extension creates a dedicated `pg_background_worker` NOLOGIN role and ships helper functions to grant or revoke execution privileges.

### Caveats

- Prefer the V2 API. The older V1 API is still present for compatibility but lacks cookie-based PID reuse protection.
- The `v1.9.2` release is a binary-only patch for assert-enabled PostgreSQL builds. The SQL extension version remains `1.9`, so there is no new SQL upgrade script or user-facing function delta from `1.9.1`.
