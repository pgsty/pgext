## Usage

Sources:

- [Official documentation](https://github.com/Srujan-rai/Pgtrace/blob/0e97f8c4528008ef39d72181b5cf48db05ab74df/README.md)
- [Extension control file](https://github.com/Srujan-rai/Pgtrace/blob/0e97f8c4528008ef39d72181b5cf48db05ab74df/pgtrace.control)
- [Official repository](https://github.com/Srujan-rai/Pgtrace)

`pgtrace` Event-driven query, lock, failure, and latency observability from PostgreSQL executor hooks.

### Enablement

Merge `pgtrace` into the existing preload list, restart PostgreSQL, and then create `pgtrace` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pgtrace'
```

```sql
CREATE EXTENSION pgtrace;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Get count of currently tracked queries
SELECT pgtrace_query_count();

-- Clear all query stats
SELECT pgtrace_reset();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pgtrace_query_stats` | VIEW | Inspection or query view created by the extension. |
| `pgtrace_alien_queries` | VIEW | Inspection or query view created by the extension. |
| `pgtrace_failing_queries` | VIEW | Inspection or query view created by the extension. |
| `pgtrace_latency_histogram` | VIEW | Inspection or query view created by the extension. |
| `pgtrace_metrics` | VIEW | Inspection or query view created by the extension. |
| `pgtrace_query_count` | FUNCTION | Callable function from the reviewed install surface. |
| `pgtrace_reset` | FUNCTION | Callable function from the reviewed install surface. |
| `pgtrace_slow_queries` | VIEW | Inspection or query view created by the extension. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16; do not infer unlisted majors.
- Preloading `pgtrace` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
