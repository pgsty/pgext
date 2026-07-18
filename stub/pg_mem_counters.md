## Usage

Sources:

- [Pinned upstream README](https://github.com/topharley/pg_mem_counters/blob/3a6a6f4e2e0b62fa23ef1041a7ae357d016e600f/README.md)
- [Pinned extension control file](https://github.com/topharley/pg_mem_counters/blob/3a6a6f4e2e0b62fa23ef1041a7ae357d016e600f/pg_mem_counters.control)

`pg_mem_counters` maintains named total-hit and requests-per-minute counters in PostgreSQL shared memory. It is useful for lightweight process-local instrumentation where durable, transactional accounting is not required.

The library must be preloaded and PostgreSQL restarted before the extension is created:

```conf
shared_preload_libraries = 'pg_mem_counters'
```

```sql
CREATE EXTENSION pg_mem_counters;

SELECT inc_mem_counter('api_requests', 1);
SELECT inc_mem_counter('api_requests');
SELECT get_mem_counter_rpm('api_requests');
SELECT * FROM mem_counters();
```

`inc_mem_counter(counter,increment)` adds to and returns the total; omitting the increment reads the total. `get_mem_counter_rpm(counter)` returns the recent per-minute count, and `mem_counters()` lists each name with its total and RPM value.

The state is held in shared memory, not in durable database tables, so do not use it for billing, audit, quotas, or any counter that must survive restart or participate in rollback. The reviewed README documents only PostgreSQL 10 and 11, and the control version is `1.0`; test memory sizing, name cardinality, concurrency, restart behavior, and current-server ABI before deployment.
