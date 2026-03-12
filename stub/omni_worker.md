


## Usage

> [omni_worker: Generalized worker pool](https://docs.omnigres.org/omni_worker/intro/)

The `omni_worker` extension provides a generalized PostgreSQL worker pool that executes arbitrary workloads within individual backend contexts.

### SQL Handler

The built-in SQL handler schedules SQL statement execution in background workers.

**Fire-and-forget:**

```sql
SELECT omni_worker.sql('INSERT INTO logs VALUES (now(), $1)');
```

Returns `null` immediately upon scheduling.

**With timeout:**

```sql
SELECT omni_worker.sql('VACUUM ANALYZE my_table', wait_ms => 1000);
```

Returns:
- `true` -- Execution completed successfully
- `false` -- Execution completed with failure
- `null` -- Execution did not complete within the timeout

The SQL handler is enabled by default but can be disabled by removing its entry from the `handlers` table.
