

## Usage

> [pg_background: Execute SQL in background worker processes](https://github.com/vibhorkum/pg_background)

Execute arbitrary SQL commands in **background worker processes** within PostgreSQL. Unlike `dblink` (which creates a separate connection), `pg_background` workers run **inside** the database server in **independent transactions**.

**Use Cases:**
- Background maintenance (VACUUM, ANALYZE, REINDEX)
- Asynchronous audit logging
- Long-running ETL pipelines
- Independent notification delivery
- Parallel query patterns

### Quick Start (V2 API)

```sql
CREATE EXTENSION pg_background;

-- Launch a background job
SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table'
) AS handle;
--   pid  |      cookie
-- -------+-------------------
--  12345 | 1234567890123456

-- Retrieve results (one-time consumption)
SELECT * FROM pg_background_result_v2(12345, 1234567890123456) AS (count BIGINT);

-- Fire-and-forget (no result needed)
SELECT * FROM pg_background_submit_v2(
  'INSERT INTO audit_log (ts, event) VALUES (now(), ''system_check'')'
) AS handle;
```


## V2 API Reference

| Function | Returns | Description |
|----------|---------|-------------|
| `pg_background_launch_v2(sql, queue_size)` | `pg_background_handle` | Launch worker, return cookie-protected handle |
| `pg_background_submit_v2(sql, queue_size)` | `pg_background_handle` | Fire-and-forget (no result consumption) |
| `pg_background_result_v2(pid, cookie)` | `SETOF record` | Retrieve results (one-time consumption) |
| `pg_background_detach_v2(pid, cookie)` | `void` | Stop tracking worker (worker continues) |
| `pg_background_cancel_v2(pid, cookie)` | `void` | Request cancellation |
| `pg_background_cancel_v2_grace(pid, cookie, grace_ms)` | `void` | Cancel with grace period |
| `pg_background_wait_v2(pid, cookie)` | `void` | Block until worker completes |
| `pg_background_wait_v2_timeout(pid, cookie, timeout_ms)` | `bool` | Wait with timeout |
| `pg_background_list_v2()` | `SETOF record` | List known workers in current session |
| `pg_background_stats_v2()` | `pg_background_stats` | Session statistics (v1.8+) |
| `pg_background_progress(pct, msg)` | `void` | Report progress from worker (v1.8+) |
| `pg_background_get_progress_v2(pid, cookie)` | `pg_background_progress` | Get worker progress (v1.8+) |

### Cancel vs Detach

| Operation | Stops Execution | Removes Tracking |
|-----------|-----------------|------------------|
| `cancel_v2()` | Yes (best-effort) | No |
| `detach_v2()` | No | Yes |

- Use **`cancel_v2()`** to **stop work** (terminate execution, prevent commit)
- Use **`detach_v2()`** to **stop tracking** (free bookkeeping while worker continues)

### Worker Lifecycle

```sql
-- Cancel a running job
SELECT pg_background_cancel_v2(pid, cookie);

-- Wait for completion
SELECT pg_background_wait_v2(pid, cookie);

-- Wait with timeout (returns true if completed)
SELECT pg_background_wait_v2_timeout(pid, cookie, 5000);

-- List active workers
SELECT * FROM pg_background_list_v2() AS (
  pid int4, cookie int8, launched_at timestamptz,
  user_id oid, queue_size int4, state text,
  sql_preview text, last_error text, consumed bool
);
```

Worker states: `running`, `stopped`, `canceled`, `error`

### Progress Reporting (v1.8+)

```sql
-- From within worker SQL
SELECT pg_background_progress(50, 'Halfway done');

-- From launcher (check progress)
SELECT * FROM pg_background_get_progress_v2(pid, cookie);
```


## GUC Settings (v1.8+)

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_background.max_workers` | 16 | Max concurrent workers per session |
| `pg_background.default_queue_size` | 65536 | Default shared memory queue size |
| `pg_background.worker_timeout` | 0 | Worker execution timeout (0 = no limit) |


## V1 API (Legacy)

The v1 API is retained for backward compatibility but lacks cookie-based PID reuse protection:

```sql
SELECT pg_background_launch('VACUUM VERBOSE my_table') AS pid \gset
SELECT * FROM pg_background_result(:pid) AS (result TEXT);
SELECT pg_background_detach(:pid);
```

The V2 API is recommended for production use due to cookie-based PID reuse protection.


## Security Model

- Extension is installed by superusers, with **no PUBLIC grants** by default
- A dedicated `pg_background_worker` NOLOGIN role is created
- Helper functions manage privileges: `grant_pg_background_privileges(role, include_v1)`
- Workers execute as the **launching user** (not superuser)
