## Usage

Sources:

- [Official README](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/README.md)
- [Extension control file](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/pg_logqueryid.control)
- [PGXN distribution and releases](https://pgxn.org/dist/pg_logqueryid/)

`pg_logqueryid` writes the `pg_stat_statements` query identifier next to plans emitted by `auto_explain`, allowing operators to correlate a logged plan with the normalized statement statistics. It adds no queryable SQL API or custom GUC of its own.

### Core Workflow

Load and configure `pg_stat_statements` and `auto_explain`, enable log collection, then load `pg_logqueryid` in a session or through a preload setting.

```conf
logging_collector = on
shared_preload_libraries = 'pg_stat_statements,auto_explain,pg_logqueryid'
pg_stat_statements.track = 'all'
auto_explain.log_min_duration = '1s'
```

```sql
SELECT queryid, query
FROM pg_stat_statements
WHERE queryid = 5917340101676597114;
```

The module can instead be activated for selected connections with `LOAD 'pg_logqueryid'` or `session_preload_libraries`. It remains inactive if the two prerequisite modules are not loaded and configured.

### Version and Logging Boundaries

The cataloged control file reports `0.0.1`, while PGXN publishes later distributions through `1.0.1`; pin the actual installed artifact rather than inferring its version from the repository README. Upstream reports validation through PostgreSQL 18, but hook interaction and log format still need an exact-build test.

On PostgreSQL 16 and later, `auto_explain` can already print the query ID when `auto_explain.log_verbose` is enabled, so this extension may be redundant. Logging every plan with a zero duration threshold can impose heavy overhead and expose SQL text or sensitive values. Start with a selective threshold and sampling policy, secure log access and retention, and verify that the logged query ID matches the intended `pg_stat_statements` entry.
