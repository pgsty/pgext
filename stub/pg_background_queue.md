## Usage

Sources:

- [Official documentation](https://github.com/hank-cp/pg_background_queue/blob/8e6f1c5429df50d5475a0d3943753e5adc125be9/README.md)
- [Extension control file](https://github.com/hank-cp/pg_background_queue/blob/8e6f1c5429df50d5475a0d3943753e5adc125be9/pg_background_queue.control)
- [Official repository](https://github.com/hank-cp/pg_background_queue)

`pg_background_queue` Topic-aware background SQL task queue with concurrency limits and retries.

### Enablement

Merge `pg_background_queue` into the existing preload list, restart PostgreSQL, and then create `pg_background_queue` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_background_queue'
```

```sql
CREATE EXTENSION pg_background_queue;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Check active worker count
SELECT pg_background_queue_active_workers_count();

-- Ensure workers are running (❗️️SUGGESTING: trigger by pg_cron)
SELECT pg_background_queue_ensure_workers();

-- Calibrate worker count (❗️️SUGGESTING: trigger by pg_cron)
SELECT pg_background_queue_calibrate_workers_count();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_background_tasks` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pg_background_enqueue` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_background_queue_ensure_workers` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_background_queue_calibrate_workers_count` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_background_queue_active_workers_count` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_background_queue_task_state` | TYPE | User-facing data type created by the extension. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_background_queue` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
