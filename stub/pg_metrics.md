## Usage

Sources:

- [Extension SQL definition](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics--1.0.sql)
- [Shared-memory implementation](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.c)
- [Extension control file](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.control)

`pg_metrics` stores named integer counters in PostgreSQL shared memory. Backends can atomically add to a counter, while monitoring sessions enumerate current values and capacity. It is a small building block for custom in-database metrics, not a general time-series store.

### Core Workflow

The library must be initialized during postmaster startup. Add it to `shared_preload_libraries`, optionally set the counter capacity, and restart PostgreSQL before creating the SQL extension.

```conf
shared_preload_libraries = 'pg_metrics'
pg_metrics.max = 100
```

```sql
CREATE EXTENSION pg_metrics;

SELECT counter_add('jobs_processed', 1);
SELECT * FROM metrics();
SELECT * FROM metrics_stats();
```

`counter_add` returns the previous counter value and then applies the increment. Negative increments are accepted by the `int8` signature and can reduce a counter.

### User-Facing Objects

- `metric_type`: enum currently containing only `COUNTER`.
- `counter_add(counter text, increment int8)`: creates a named counter on first use and atomically adds the increment.
- `metrics()`: returns `metric_name`, `metric_type`, and `counter_value` for each shared counter.
- `metrics_stats()`: returns `max_metrics` and `num_metrics`.
- `pg_metrics.max`: postmaster setting for the maximum number of counters; the source default is `50`, with a minimum of `10`.

### Operational Notes

Preloading and a server restart are mandatory: without initialized shared memory, `counter_add` returns `NULL`. Names longer than 127 bytes are rejected, and adding a new counter after the shared hash reaches `pg_metrics.max` also returns `NULL`. Values live only in shared memory and are reset when the shared-memory state is recreated; there is no persistence or reset SQL API in version `1.0`. The counters are shared across backends, while the SQL objects must exist in each database that calls them. Export or snapshot values externally if historical metrics are required.
