## Usage

Sources:

- [Official README](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/README.md)
- [Extension control file](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench.control)
- [Extension SQL](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench--1.0.sql)
- [Hook and perf implementation](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench.c)

`pg_microbench` instruments a Linux PostgreSQL backend with `perf_event_open` counters. It can measure planner, executor, utility, or explicit global SPI scopes and prints hardware-counter results as `NOTICE` messages for development experiments.

### Core Workflow

The operating system must permit the PostgreSQL account to open the required performance events. Prefer a narrowly scoped capability or security policy on an isolated host rather than weakening the host-wide perf policy.

```sql
CREATE EXTENSION pg_microbench;

REVOKE ALL ON FUNCTION pg_microbench_run(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_microbench_run(text) TO benchmark_admin;

SET pg_microbench.enable = on;
SET pg_microbench.log = on;
SET pg_microbench.scopes = 'planner,executor';
SELECT count(*) FROM generate_series(1, 100000);

SET pg_microbench.scopes = 'global';
SELECT pg_microbench_run('SELECT count(*) FROM generate_series(1, 100000)');
```

`pg_microbench.enable` opens or closes the per-backend perf file descriptors. `pg_microbench.log` controls notice output, and `pg_microbench.scopes` accepts a comma-separated subset of `global`, `planner`, `executor`, and `utility`. Reported fields include wall time, cycles, instructions, cache references and misses, L1 data and instruction events, branches, and branch misses.

### Benchmark Boundaries

The counters are session-local and the output is not persisted. Unsupported or restricted hardware events may be unavailable, and measurements vary with the CPU, kernel, PostgreSQL build, cache state, scheduler, and concurrent workload. Warm up the query, repeat trials, record the environment, and compare distributions rather than treating one notice as a guarantee.

`pg_microbench_run(text)` executes its SQL string synchronously with the caller's privileges and can contain writes or expensive work. Restrict it when only benchmark administrators should initiate such workload. Enabling hook scopes also adds measurement overhead to matching queries in that session. The pinned source explicitly targets Linux and PostgreSQL development master as of January 2026, so build against and test the exact server source; no broad released-major compatibility is promised. No preload is required because installing the extension loads the library in the session.
