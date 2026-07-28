## Usage

Sources:

- [Official upstream README](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/README.md)
- [Official extension control file (pg_pulse.control)](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/pg_pulse.control)
- [Official extension SQL (pg_pulse--1.0.sql)](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/pg_pulse--1.0.sql)

`pg_pulse` — A lock-free, in-memory Active Session History (ASH) sampler for PostgreSQL 14–18 — inspired by Oracle's ASH, built for low-overhead, continuous session-level monitoring in production. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_pulse;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_pulse_conn_ring_sample(OUT sample_time timestamptz, OUT total_connections int, OUT active_connections int, OUT idle_connections int, OUT idle_in_transaction_connections int, OUT idle_in_transaction_aborted_connections int, OUT other_connections int)` is an extension function and returns `SETOF`.
- `pg_pulse_purge_history(retention_days int DEFAULT NULL)` is an extension function and returns `void`.
- `pg_pulse_connections_live` is an extension-defined view.
- `pg_pulse_live` is an extension-defined view.
- `pg_pulse_connection_history` is a table installed or managed by the extension.
- `pg_pulse_history` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
