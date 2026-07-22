## Usage

Sources:

- [Official README](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/README.md)
- [Official extension SQL](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/sql/pg_octopus.sql)
- [Official health-check implementation](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/src/pg_octopus.c)

`pg_octopus` is an experimental background worker that records whether it can establish a libpq connection to a list of PostgreSQL host/port pairs. It checks TCP/database login reachability only; it does not assess replication state, read/write readiness, lag, quorum, or application health.

### Enable the Worker

The worker must be preloaded and PostgreSQL restarted. Its implementation connects to the local database named `postgres`, so create the extension and node table there:

```ini
shared_preload_libraries = 'pg_octopus'
pg_octopus.health_check_period = 10000
pg_octopus.health_check_timeout = 2000
pg_octopus.health_check_max_retries = 2
pg_octopus.health_check_retry_delay = 1000
```

```sql
CREATE EXTENSION pg_octopus;

INSERT INTO octopus.nodes(node_name, node_port)
VALUES ('db-a.internal', 5432), ('db-b.internal', 5432);

SELECT node_name, node_port, health_status
FROM octopus.nodes
ORDER BY node_name, node_port;
```

`health_status` is `1` for reachable, `0` for unreachable after retries, and `-1` before a result is known.

### Configuration Index

- `pg_octopus.health_check_period` sets the round duration in milliseconds.
- `pg_octopus.health_check_timeout` sets each connection timeout in milliseconds.
- `pg_octopus.health_check_max_retries` limits retries after the first attempt.
- `pg_octopus.health_check_retry_delay` sets the delay between attempts.
- `pg_octopus.nodes_table` changes the table name read and updated by the worker.

Keep the timeout plus all retry delays and timeouts below the period so rounds do not continuously overrun.

### Operational Boundaries

Remote checks always target database `postgres` and do not put a user or password in the connection string; libpq defaults, the worker environment, and `pg_hba.conf` therefore determine authentication. A successful connection proves only that this path and credential set can log in. The worker also has `BGW_NEVER_RESTART`, so a crash leaves checks stopped until PostgreSQL restarts. The reviewed source is an old Citus demonstration and uses dynamically constructed SQL for the configured table and node values. Restrict writes to the nodes table, monitor the worker process and timestamp results externally, test reload behavior, and do not use `health_status` alone for automated failover.
