## Usage

Sources:

- [Official documentation](https://github.com/spa-28/pg_logtap/blob/cf59eb5c7dbcc2442fcf6ef1b6dfeea7fd0a5246/README.md)
- [Extension control file](https://github.com/spa-28/pg_logtap/blob/cf59eb5c7dbcc2442fcf6ef1b6dfeea7fd0a5246/pg_logtap.control)
- [Official repository](https://github.com/spa-28/pg_logtap)

`pg_logtap` Structured PostgreSQL log capture and delivery to Vector, HTTP, TCP, or files.

### Enablement

Merge `pg_logtap` into the existing preload list, restart PostgreSQL, and then create `pg_logtap` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_logtap'
```

```sql
CREATE EXTENSION pg_logtap;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM pg_logtap_delivery;          -- monitoring view, one row (see below)
SELECT pg_logtap_stats();                  -- same counters as compact text
SELECT unnest(pg_logtap_dump(100));        -- last events as JSON, non-destructive
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_logtap_delivery` | VIEW | Inspection or query view created by the extension. |
| `pg_logtap_stats_json` | FUNCTION | Callable function from the reviewed install surface. |
| `grants` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_logtap_stats_t` | TYPE | User-facing data type created by the extension. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_logtap` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
