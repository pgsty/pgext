## Usage

Sources:

- [Official repository README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [Extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/pgx_trace_hooks.control)
- [Trace hook implementation](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/src/lib.rs)

`pgx_trace_hooks` is a diagnostic preload extension that logs entry into PostgreSQL's executor and utility hooks. Use it in a controlled test server to observe whether statements reach `ExecutorStart`, `ExecutorRun`, `ExecutorFinish`, `ExecutorEnd`, and `ProcessUtility`; it records hook-stage names, not query plans, timings, or statement attribution.

### Core Workflow

The library rejects ordinary runtime loading, so configure it in `shared_preload_libraries` and restart PostgreSQL before creating the SQL extension:

```conf
shared_preload_libraries = 'pgx_trace_hooks'
```

```sql
CREATE EXTENSION pgx_trace_hooks;

SELECT 1;
CREATE TEMP TABLE trace_probe(id integer);
```

Inspect the PostgreSQL server log after executing representative statements. A query can emit executor-stage entries, while utility commands can emit `ProcessUtility`. The implementation preserves and calls the previously installed callback for every hook, or falls back to PostgreSQL's standard handler when none was installed.

### Traced Hooks

- `ExecutorStart` marks executor initialization.
- `ExecutorRun` marks tuple-producing execution.
- `ExecutorFinish` marks executor completion work.
- `ExecutorEnd` marks executor teardown.
- `ProcessUtility` marks utility-statement processing.

### Operational Notes

`pgx_trace_hooks` is non-relocatable and requires superuser privileges to create. `_PG_init` raises an error unless the library is being processed through `shared_preload_libraries`; changing that setting requires a server restart. Hook messages use PostgreSQL's `INFO` level and can multiply log volume for ordinary workloads, so keep this extension out of busy production systems unless log cost and data exposure have been explicitly evaluated.

The hooks chain to the callbacks that were present when this library initialized. Preload order can therefore affect the resulting chain; test it with every other hook-using extension in the final order. The reviewed source belongs to a PostgreSQL 15-era research repository and exposes no SQL API beyond loading the library.
