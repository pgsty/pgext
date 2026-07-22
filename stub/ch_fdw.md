## Usage

Sources:

- [Official README](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/README.md)
- [Official FDW handler](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/ch_fdw.c)
- [Official extension control file](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/ch_fdw.control)

`ch_fdw` is an experimental, read-only foreign data wrapper that lets PostgreSQL query ClickHouse over its native TCP protocol. The implementation combines PostgreSQL C internals with Go through CGo and supports selected filter and aggregate pushdown.

### Core Workflow

Create the extension, define a server, then map a PostgreSQL foreign table to an existing ClickHouse table with compatible column names and types:

```sql
CREATE EXTENSION ch_fdw;

CREATE SERVER clickhouse_server
  FOREIGN DATA WRAPPER ch_fdw
  OPTIONS (host '127.0.0.1', port '9000');

CREATE FOREIGN TABLE click_events (
  event_name text,
  event_count bigint,
  observed_at timestamptz
)
SERVER clickhouse_server
OPTIONS (table 'events');

SELECT event_name, sum(event_count)
FROM click_events
WHERE observed_at >= now() - interval '1 day'
GROUP BY event_name;
```

Use `EXPLAIN (VERBOSE)` to confirm which filters and aggregates are actually sent to ClickHouse; unsupported expressions remain local or may fail at deparse/conversion boundaries.

### Implemented Boundary

- The handler supplies scan, rescan, analyze, explain, and upper-path callbacks. It has no foreign-modification callbacks, so `INSERT`, `UPDATE`, and `DELETE` are not supported.
- Upstream documents `WHERE` and aggregate pushdown. Source comments exclude joins, parameters, ordering pushdown, grouping sets, ordered-set aggregates, aggregate `FILTER`, and several expression forms.
- Server options used by the reviewed code are `host` and `port`; the validator is a no-op, so misspelled or unsafe options are not rejected during DDL.

### Operational Caveats

- Upstream labels the project experimental, says it has no active MessageBird development or use, and reports testing only with PostgreSQL 13, ClickHouse 20.3.19.4, and Ubuntu 20.04. Treat the repository as a reference implementation, not a maintained compatibility promise.
- The reviewed connection code forces ClickHouse driver's debug mode and uses the default ClickHouse user when no username is supplied; its option parsing does not establish a documented production credential workflow. Inspect logs and code before using sensitive endpoints.
- Query pushdown changes semantics across PostgreSQL and ClickHouse type, time-zone, collation, null, numeric, and function rules. Compare pushed and non-pushed results with representative edge cases.
- CGo, C memory, Go runtime, and PostgreSQL memory contexts share one backend process. Stress cancellation, reconnect, errors, large results, repeated scans, and backend shutdown on the exact server build.
