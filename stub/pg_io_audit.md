## Usage

Sources:

- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/README.md)
- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/pg_io_audit.control)
- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/LICENSE)
- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/pg_io_audit--1.0.sql)
- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/patches/README.md)
- [Official source](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/docs/limitations.md)

pg_io_audit is a research module for PostgreSQL 17 with the project's storage-manager hook patch. It is not loadable on an ordinary unpatched server. Counters describe requests to the storage manager, not confirmed physical I/O completion.

### Enable and inspect

```ini
shared_preload_libraries = 'pg_io_audit'
pg_io_audit.enabled = on
pg_io_audit.trace_size = 1024
```

Restart the patched server, then install the SQL interface as an administrator. Monitoring functions and views are limited to pg_monitor or superusers; reset is superuser-only.

```sql
CREATE EXTENSION pg_io_audit;
SELECT * FROM pg_io_audit_stats;
SELECT * FROM pg_io_audit_trace;
SELECT pg_io_audit_reset();
```

### Interpretation

Statistics group calls and blocks by operation and fork. The optional bounded trace records recent relation-file events on a best-effort basis. enabled can be reloaded; trace_size requires restart. Shared-memory counters and traces disappear after restart or crash. WAL, SLRU, temporary files and direct directory operations are outside these hooks.
