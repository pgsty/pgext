## Usage

Sources:

- [Official documentation](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension/blob/1cbd9d46a401b22070549543e493f19845eb3ad2/README.md)
- [Extension control file](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension/blob/1cbd9d46a401b22070549543e493f19845eb3ad2/pg_trace.control)
- [Official repository](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension)

`pg_trace` Query execution tracing with shared-memory buffers, persistence, and retention controls.

### Enablement

Install the files for the intended server, then create `pg_trace` in the target database:

```sql
CREATE EXTENSION pg_trace;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- All commands in one session
SELECT pg_trace_start();
SELECT 1+1;
SELECT pg_trace_stop();
SELECT * FROM pg_trace_queries;
SELECT pg_trace_flush();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_trace_start` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_trace_stop` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_trace_flush` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_trace_log` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pg_trace_queries` | VIEW | Inspection or query view created by the extension. |
| `pg_trace_cleanup_old` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_trace_clear` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_trace_hourly_stats` | VIEW | Inspection or query view created by the extension. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 16, 17, 18; do not infer unlisted majors.
