## Usage

Sources:

- [Official usage documentation](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/doc/pgbouncer_wrapper.md)
- [Version 1.2.2 control file](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/pgbouncer_wrapper.control)
- [Views, foreign server, and control functions](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/sql/pgbouncer_wrapper.sql)

`pgbouncer_wrapper` 1.2.2 exposes the PgBouncer administration console through PostgreSQL views and SQL control functions backed by `dblink_fdw`. It is useful for SQL-based monitoring and carefully controlled operations, but it delegates the PgBouncer admin identity to database callers and must be treated as a privileged management interface.

### Connect the Wrapper

The install SQL creates a foreign server named `pgbouncer` with Unix-socket host `/tmp`, port `6432`, admin database `pgbouncer`, plus a `PUBLIC` user mapping using user `pgbouncer`:

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pgbouncer_wrapper;

SELECT * FROM pgbouncer.version;
SELECT database, "user", cl_active, cl_waiting, sv_active, sv_idle
FROM pgbouncer.pools;
SELECT database, total_query_count, avg_query_time
FROM pgbouncer.stats;
```

If the admin console uses another socket, port, user, or authentication method, alter the `pgbouncer` foreign server and user mapping as a privileged migration before querying. Test the exact PgBouncer version because each view declares a fixed column layout for a `SHOW` command.

### Monitoring and Control Objects

Monitoring views include `clients`, `servers`, `pools`, `databases`, `stats`, `stats_averages`, `stats_totals`, `totals`, `config`, `lists`, `mem`, `sockets`, `active_sockets`, `users`, `version`, and DNS/state views.

Control functions include `disable`, `enable`, `kill`, `pause`, `reconnect`, `reload`, `resume`, `shutdown`, `suspend`, `wait_close`, and `pgbouncer.set`. These send commands directly to the admin console; `kill`, `shutdown`, `pause`, and `suspend` can interrupt service.

### Privilege and Compatibility Boundaries

The install creates `CREATE USER MAPPING FOR PUBLIC`, so every database role later granted access to the `pgbouncer` schema can use the same mapped PgBouncer identity. Revoke default function execution, avoid broad schema usage, and grant read-only views separately from control functions. Never expose `pgbouncer.fds` casually: its declared output includes password and SCRAM key fields, and the underlying `SHOW FDS` command can block PgBouncer's event loop.

`pgbouncer.set(key, value)` formats the key as raw command text, so only trusted, validated keys may reach it. Treat all control-function arguments as administrator input.

PgBouncer changes `SHOW` columns across releases. A mismatch can make a view fail or misrepresent fields; validate every view against the deployed PgBouncer documentation and run a failover/restart drill before automating control calls. The wrapper is not a telemetry cache: every query opens a live `dblink` request and inherits admin-console availability and latency.
