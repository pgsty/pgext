## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/pg_unified.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/pg_unified--1.1.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_unified/meson.build)

`pg_unified` `1.1` provides buffered events and batch rows through one view. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_unified;
CREATE TABLE unified_demo_batch(url text);
SELECT unified_register('demo', 'unified_demo_batch');
SELECT unified_ingest('demo', '{"url":"/home"}');
SELECT * FROM demo_unified;
SELECT unified_flush('demo');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

The registered view combines batch rows with unflushed events. unified_flush moves queued rows, unified_purge removes eligible history and unified_stream_events exposes the buffer. Schedule flushing explicitly and keep event JSON compatible with the batch table. Inspect privileges on the generated view and underlying tables; this is a database-local buffer, not a distributed streaming broker.

This module corresponds to the former upstream name `alohadb_unified`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
