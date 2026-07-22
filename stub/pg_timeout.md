## Usage

Sources:

- [Official README](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/README.md)
- [Official extension SQL](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout--1.0.sql)
- [Official worker implementation](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout.c)

`pg_timeout` is an archived background worker that terminates cluster sessions whose `pg_stat_activity.state` has remained `idle` beyond a global threshold. Current PostgreSQL releases provide the built-in `idle_session_timeout`; prefer that supported setting unless maintaining an older server.

### Enablement

The behavior comes from preloading the library, not from calling its SQL function. Configure it and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_timeout'
pg_timeout.naptime = 30
pg_timeout.idle_session_timeout = 300
```

You may create the extension after restart to register its `pg_timeout_main()` SQL object, but normal users should never invoke that worker entry point directly:

```sql
CREATE EXTENSION pg_timeout;
SHOW pg_timeout.naptime;
SHOW pg_timeout.idle_session_timeout;
```

`pg_timeout.naptime` is the scan interval in seconds; `pg_timeout.idle_session_timeout` is the idle threshold in seconds. Both are reloadable.

### Termination Semantics

The worker connects to the local `postgres` database, scans the cluster-wide `pg_stat_activity` view, logs matching session metadata, and calls `pg_terminate_backend()` for every session with state exactly `idle` and an older `state_change`. It deliberately does not terminate `idle in transaction` sessions.

### Operational Boundaries

The threshold is cluster-wide: there is no per-role, per-database, application, or allowlist exception, so pools, administrators, replication tooling, and maintenance clients can all be disconnected. Termination is sampled at the nap interval and is not exact to the second. The worker uses elevated termination authority and logs usernames, databases, application names, and hostnames. Test reconnect behavior and exclusions before enabling it, monitor restart loops, and use `idle_in_transaction_session_timeout` separately for abandoned transactions. On supported PostgreSQL versions, replace this extension with role/database-scoped built-in timeout settings.
