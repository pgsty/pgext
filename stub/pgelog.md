
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION IF NOT EXISTS dblink;
> CREATE EXTENSION IF NOT EXISTS pg_variables;
> CREATE EXTENSION pgelog;
> SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
> ```
>
> Source: [README](https://github.com/anfiau/pgelog)

`pgelog` writes log records into PostgreSQL tables using pseudo-autonomous transactions implemented through `dblink`. The key goal is that log entries survive even when the caller's main transaction rolls back.

## Prerequisites

The README requires:

- PostgreSQL 11 or newer
- `dblink`
- `pg_variables`
- local passwordless `dblink` access, typically via a `peer` local entry in `pg_hba.conf`

It also warns that each session may open one extra connection for `dblink`, so `max_connections` should be sized accordingly.

## Objects

The extension creates:

- `pgelog_params` for configuration
- `pgelog_logs` as the base log table
- `pgelog_vw_logs` as a log view with timing information

The log table/view stores fields such as timestamp, log type, source function, phase, message text, transaction id, SQLSTATE, SQLERRM, and connection name.

## Basic Logging

Write a log entry:

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
```

Read the latest log:

```sql
SELECT log_stamp, log_info
FROM pgelog_logs
ORDER BY log_stamp DESC
LIMIT 1;
```

## PL/pgSQL Exception Logging

The README includes a larger PL/pgSQL example that catches exceptions, collects diagnostics, writes a `FAIL` log entry through `pgelog_to_log(...)`, and then re-raises the exception. This is the main pattern for capturing rollback-resistant failure logs.

## Configuration

Configuration parameters are managed with:

```sql
SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

The README documents `pgelog_ttl_minutes` and other parameters through the `pgelog_params` table and helper functions.
