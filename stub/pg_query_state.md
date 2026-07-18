## Usage

Sources:

- [Pinned current upstream README](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/README.md)
- [Pinned extension control file](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/pg_query_state.control)
- [Pinned executor and messaging implementation](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/pg_query_state.c)
- [PostgreSQL 18 custom-signal patch](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/patches/custom_signals_18.0.patch)
- [PostgreSQL 18 runtime-EXPLAIN patch](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/patches/runtime_explain_18.0.patch)

pg_query_state version 1.2 extracts the live execution state of a query running in another backend. It can return nested query frames and parallel-worker plans with current row, loop, timing, and buffer statistics, and it offers numeric and repeatedly printed progress estimates.

### Patched-server setup and query

This is not a drop-in extension. Apply both upstream patches for the exact PostgreSQL major, rebuild and deploy the server, then preload the module and restart:

```conf
shared_preload_libraries = 'pg_query_state'
pg_query_state.enable = on
pg_query_state.enable_timing = off
pg_query_state.enable_buffers = off
```

```sql
CREATE EXTENSION pg_query_state;

SELECT pid, frame_number, query_text, plan, leader_pid
FROM pg_query_state(
    12345,
    format => 'json'
);

SELECT pg_progress_bar(12345);
```

Replace the PID with a currently executing backend. A caller must be superuser or have the same effective role as the target backend.

### Invasive monitoring boundary

The server patches add custom process signals and runtime EXPLAIN state to PostgreSQL internals. The preloaded module installs executor hooks, allocates shared memory, sends signals to the leader and parallel workers, and transfers query text and plan state through shared-memory queues. It always requests row instrumentation when enabled; timing and buffer collection add further overhead. Validate the complete patched server build, not just the shared library, for every minor update and extension interaction.

Output contains live query text, literals, object names, nested function queries, and plan details. Same-role access can therefore reveal another session's data-bearing SQL. Restrict role sharing and function execution, protect monitoring output, and test timeout and cancellation behavior so an unresponsive target cannot stall monitors.

Progress is a heuristic derived from plan rows and is unavailable or misleading for uncountable operations, inaccurate estimates, blocking waits, and many write plans. The current repository contains patches through PostgreSQL 18 and was active in 2026, but that does not make a patched production server low-risk. Deploy first to a dedicated diagnostics build and measure overhead with realistic parallel and nested workloads.
