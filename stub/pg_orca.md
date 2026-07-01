


## Usage

Sources: [pgorca README](https://github.com/quantumiodb/pgorca), [entrypoint/GUC source](https://github.com/quantumiodb/pgorca/blob/main/pg_orca.cpp), [control file](https://github.com/quantumiodb/pgorca/blob/main/pg_orca.control).

`pg_orca` plugs the ORCA cost-based optimizer from the Greenplum/Apache Cloudberry lineage into a standard PostgreSQL 18 server. The upstream README describes this project as PostgreSQL 18-only, and the local package metadata also marks it for PG18 only.

### Enable ORCA For A Session

`CREATE EXTENSION` loads the library in the current session, so the `pg_orca.*` GUCs and planner hook are available immediately:

```sql
CREATE EXTENSION pg_orca;

SET pg_orca.enable_orca = on;

EXPLAIN
SELECT *
FROM orders
WHERE customer_id = 42
  AND created_at >= now() - interval '30 days';
```

If ORCA cannot handle a query, the README says it falls back to the standard PostgreSQL planner automatically. Turn on fallback logging while validating a workload:

```sql
SET pg_orca.trace_fallback = on;
```

### Preload For New Connections

For automatic planner-hook loading in later sessions, upstream recommends `session_preload_libraries`, not `shared_preload_libraries`:

```sql
ALTER DATABASE mydb SET session_preload_libraries = 'pg_orca';
ALTER DATABASE mydb SET pg_orca.enable_orca = on;
```

Existing sessions are unaffected until they reconnect. If another session preload library is already configured, include both values explicitly:

```sql
ALTER DATABASE mydb
SET session_preload_libraries = 'pg_orca,pg_stat_statements';
```

Role-local and cluster-wide scopes are also valid:

```sql
ALTER ROLE bench SET session_preload_libraries = 'pg_orca';

ALTER SYSTEM SET session_preload_libraries = 'pg_orca';
SELECT pg_reload_conf();
```

### Tuning And Diagnostics

The README documents these main GUCs:

- `pg_orca.enable_orca`: enable ORCA; default `off`.
- `pg_orca.trace_fallback`: log fallback to the standard planner; default `off`.
- `optimizer_segments`: segment count for costing; default `1`.
- `optimizer_sort_factor`: sort cost scaling; default `1.0`.
- `optimizer_metadata_caching`: cache relation metadata; default `on`.
- `optimizer_mdcache_size`: metadata cache size in KB; default `16384`.
- `optimizer_search_strategy_path`: optional custom search strategy XML path.

The entrypoint source also defines additional ORCA tuning and debug GUCs such as `optimizer_join_order`, `pg_orca.join_order_dynamic_threshold`, `pg_orca.cost_model`, and `optimizer_print_*`. Treat those as workload/debug knobs and validate plans before keeping them in a persistent database setting.

### Caveats

- PostgreSQL 18 only.
- Use `session_preload_libraries = 'pg_orca'` for automatic loading in new sessions.
- ORCA is disabled by default after loading; set `pg_orca.enable_orca = on`.
- Fallback to the PostgreSQL planner is expected for unsupported queries; enable `pg_orca.trace_fallback` when checking coverage.
