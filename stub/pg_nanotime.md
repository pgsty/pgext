## Usage

Sources:

- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/README.md)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/pg_nanotime.control)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/sql/pg_nanotime--0.1.0.sql)
- [Official source](https://github.com/bunko666-cloud/pg_nanotime/blob/668e19bd4bfe1de77a95a94e69fd9b3493c21e46/docs/TECHNICAL.md)

`pg_nanotime` provides nanosecond timestamp and interval types on PostgreSQL 17–18. Version 0.1.0 is a source extension; precision does not imply nanosecond clock accuracy.

### Core Workflow

```sql
CREATE EXTENSION pg_nanotime;
CREATE TABLE events (at nanotimestamptz DEFAULT clock_ns(), body text);
INSERT INTO events (body) VALUES ('received');
SELECT at, at::timestamptz FROM events;
```

`nanotimestamptz` and `nanotimestamp` use 16 bytes with and without time zones. `nanotimestamptz8` uses 8 bytes with a narrower 1677–2262 range. `nanointerval` represents fixed elapsed durations, unlike calendar-aware PostgreSQL intervals. Conversion to built-in timestamps loses sub-microsecond precision.

### Clocks and Boundaries

`clock_ns()` reads the wall clock; `now_ns()` caches an extension clock reading for the transaction, which is not the same instant as PostgreSQL's transaction timestamp. `clock_ns_monotonic()` orders calls within one backend. `clock_ns_global()` orders calls across backends and requires this configuration followed by a server restart:

```ini
shared_preload_libraries = 'pg_nanotime'
```

Preserve other libraries already in the setting. Ordinary types and clocks do not require preloading. Global monotonic state is in shared memory and does not survive restart; it is not a durable sequence or distributed ordering mechanism. Install as a superuser. The extension is relocatable and is not marked trusted.
