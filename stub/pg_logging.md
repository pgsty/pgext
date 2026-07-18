## Usage

Sources:

- [Upstream README](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/README.md)
- [Extension install SQL](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/main.sql)
- [Shared-buffer implementation](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/pg_logging.c)

pg_logging captures PostgreSQL error-reporting events in a shared-memory ring buffer and exposes them as rows. It can support short-lived diagnostics or an external collector that needs structured log level, message, query, detail, and position fields.

### Server configuration

The module refuses normal late loading, so preload it and restart PostgreSQL. Buffer size is allocated at startup:

```ini
shared_preload_libraries = 'pg_logging'
pg_logging.buffer_size = 10240
pg_logging.enabled = on
```

Install the SQL objects, inspect records without advancing the shared read position, then flush them deliberately:

```sql
CREATE EXTENSION pg_logging;

SELECT level::error_level, message, position
FROM get_log(false)
ORDER BY position;

SELECT flush_log();
```

Calling get_log() with its default true argument advances the shared read position. A consumer can instead save position and call the integer overload to resume from a known point.

### Caveats

- The ring buffer overwrites old unfetched records when it wraps; it is not durable audit storage.
- Reading with flush enabled or calling flush_log changes shared consumer state. Coordinate multiple collectors so one does not discard another's unread range.
- Captured fields can contain SQL text, bind-related context, object names, user data, and error details. Restrict function access and apply a short, explicit retention policy in downstream storage.
- Larger buffer settings consume shared memory and require a restart. Measure logging overhead and overflow under peak error rates.
- Upstream does not publish a current PostgreSQL major-version compatibility matrix; validate this server-internal module on the exact target build.
