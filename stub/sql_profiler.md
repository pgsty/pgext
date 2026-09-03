## Usage

Sources:

- [Official SQL Profiler documentation](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/)
- [Official installation guide](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/installing/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/sqlprofiler/using/)

`sql_profiler` captures bounded query traces for selected users and databases, including timing and explain data that can be loaded into a temporary table for analysis.

### Enablement

Install SQL Profiler 4.2.0, preload the hyphenated library name, restart, and create the underscore-named extension:

```ini
shared_preload_libraries = 'sql-profiler'
```

```sql
CREATE EXTENSION sql_profiler;
GRANT sql_profiler_admin TO profiler_user;
```

The `sql_profiler_admin` role lets a non-superuser call profiler functions; the surrounding PEM integration has additional account requirements.

### Capture and Read a Trace

`sp_activate` returns a trace ID. Use limits so a forgotten trace cannot grow without bound.

```sql
SELECT sp_activate(
  'checkout investigation',
  ''::oidvector,
  ''::oidvector,
  256,
  100,
  interval '15 minutes'
);

SELECT * FROM sp_active_traces();
SELECT sp_deactivate(1);
SELECT sp_load_trace(1, true);
```

Loaded rows appear in `_sp_tmp_tbl_sql_profiler`; `sp_traces_list` lists completed traces and `sp_cleanup` removes trace data.

### Operational Boundaries

Trace files can contain query text, parameters, plans, usernames, and database identifiers. Limit duration, minimum statement time, and maximum size; restrict the profiler role and filesystem; and clean up after analysis. Preloading adds hooks cluster-wide even when no trace is active, so validate overhead on representative workloads.

