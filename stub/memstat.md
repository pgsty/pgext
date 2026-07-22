## Usage

Sources:

- [Official README](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/README.md)
- [Extension SQL](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat--1.0.sql)
- [Collector implementation](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat.c)

memstat reports PostgreSQL memory-context allocation statistics. It can inspect the current backend without preload, or collect per-backend summaries for the whole instance when loaded at server start. Use it for focused memory diagnostics, not permanent high-frequency monitoring.

### Core Workflow

For the current session, create the extension and inspect its context tree:

```sql
CREATE EXTENSION memstat;

SELECT name, level, totalspace, freespace
FROM local_memory_stats()
ORDER BY totalspace DESC;
```

For all backends, put `memstat` last in `shared_preload_libraries`, restart PostgreSQL, and query the instance view:

```ini
shared_preload_libraries = 'other_library,memstat'
memstat.period = 10s
```

```sql
SELECT * FROM memory_stats ORDER BY totalspace DESC;
```

### Important Objects

- `local_memory_stats` returns context name, nesting level, block and chunk counts, total bytes, and free bytes for the current backend.
- `instance_memory_stats` adds the backend PID and reads statistics for live backends; preload is required.
- `memory_stats` aggregates the instance data into a per-backend summary.
- `memstat.period` sets the minimum seconds between instance-statistic collections.

### Operational Notes

The preload hook collects memory statistics at the beginning of queries and can be expensive on a busy server. Increase `memstat.period`, sample briefly, and remove the module after diagnosis. Upstream says it must be last in the preload list because it wraps shared-memory setup. Context names and layouts are PostgreSQL implementation details and can change across major versions.
