## Usage

Sources:

- [Official source](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/README.md)
- [Official source](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/LICENSE)
- [Official source](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/pg/pg_pmu.c)

pg_pmu is a headless PostgreSQL 19/master module. Its module identity declares version 0.1.0. It adds hardware instruction and cycle counts to EXPLAIN ANALYZE on Linux and macOS, including participating parallel workers.

### Enable and use

Add the library to the existing preload list and restart. There is no CREATE EXTENSION step. PMU requires ANALYZE, which really executes the statement; use read-only examples or the normal safeguards for a modifying statement.

```ini
shared_preload_libraries = 'pg_pmu'
```

```sql
EXPLAIN (ANALYZE, PMU, TIMING off) SELECT count(*) FROM pg_class;
EXPLAIN (ANALYZE, PMU, VERBOSE, FORMAT JSON, TIMING off)
SELECT count(*) FROM pg_class;
```

### Interpretation and limits

Linux requires kernel.perf_event_paranoid at most 2 or CAP_PERFMON; containers and virtual machines may still deny counters. Failure to obtain or retain counters fails the query rather than returning a plausible count. Linux user-space counts and macOS kernel-inclusive counts are not comparable. Worker counts are already included in execution totals; VERBOSE separates their contribution.

Counts are per statement, not per plan node, and include measurement overhead. Cached prepared plans lack planning counts. auto_explain is outside the interface. pg_pmu.max_nested defaults to 16 and is fixed at server start; exhausted slots or hardware resources raise errors.
