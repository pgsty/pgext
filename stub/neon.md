## Usage

Sources:

- [Official Neon extension documentation](https://neon.com/docs/extensions/neon)
- [Official extension README](https://github.com/neondatabase/neon/blob/842a5091d5db4c23aeb29aea070c37ad06b12d63/pgxn/neon/README.md)
- [Official public control file](https://github.com/neondatabase/neon/blob/842a5091d5db4c23aeb29aea070c37ad06b12d63/pgxn/neon/neon.control)

`neon` exposes Neon compute statistics and service-internal functions, most notably Local File Cache metrics. It is part of Neon's modified PostgreSQL compute stack; installing its SQL objects on an ordinary PostgreSQL server does not provide Neon storage, WAL safekeepers, or a Local File Cache.

### Core Workflow

On a Neon database, install the extension in the database that will expose the view, run a representative workload, and then inspect cache behavior:

```sql
CREATE EXTENSION neon;

SELECT file_cache_misses,
       file_cache_hits,
       file_cache_used,
       file_cache_writes,
       file_cache_hit_ratio
FROM neon_stat_file_cache;
```

`file_cache_hits` counts requests served by the Local File Cache after missing PostgreSQL shared buffers. `file_cache_misses` counts requests that must reach Neon storage. `file_cache_hit_ratio` is the percentage of Local File Cache hits among hits plus misses; it is not the PostgreSQL shared-buffer hit ratio.

### Important Objects

- `neon_stat_file_cache` is the documented user-facing view for cache misses, hits, accesses, writes, and hit ratio.
- `pg_cluster_size` reports cluster size using Neon storage semantics.
- `approximate_working_set_size` and `approximate_working_set_size_seconds` estimate working-set size.
- `get_local_cache_state`, `get_prewarm_info`, and `prewarm_local_cache` support Local File Cache state capture and prewarming.
- `backpressure_lsns`, `backpressure_throttling_time`, `neon_backend_perf_counters`, and `neon_perf_counters` expose service internals primarily used by Neon monitoring.

Grant monitoring access through `pg_monitor` where the service supports it. Do not expose prewarm or internal control functions broadly merely because they appear in the extension.

### Lifecycle and Deployment Boundaries

The statistics cover the whole compute, not one database or table, and live only for the compute lifetime. They reset when a compute stops, restarts, or is resized, so measure only after a representative workload has run. Small working sets may be served entirely from shared buffers and produce few Local File Cache observations.

The core `neon` shared library is loaded during server startup and implements Neon's storage manager, pageserver communication, WAL proposer, control-plane hooks, and file cache. That preload and its service configuration are managed by Neon; a customer should not add or tune it as a normal portable extension.

There is a version provenance mismatch: the catalog records 1.9, while the cited public upstream control file at the reviewed commit declares 1.6. The managed service may deploy private or later packaging, so use `pg_extension.extversion` on the target Neon compute and do not infer the service build from the public control file alone.
