## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime.c)

`pg_realtime` `1.0` provides trigger-based change events and notifications. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_realtime;
CREATE TABLE realtime_demo(id int PRIMARY KEY);
SELECT pg_realtime_subscribe('realtime_demo');
INSERT INTO realtime_demo VALUES (1);
SELECT * FROM pg_realtime_poll(0);
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Subscription installs row triggers; pg_realtime_unsubscribe removes a subscription and pg_realtime_cleanup trims events. Polling reads an event table and is distinct from consuming LISTEN/NOTIFY. The serializer represents non-null column values as JSON strings. Review event-table grants and payload sensitivity, and arrange cleanup; notifications are not a durable exactly-once delivery system.

This module corresponds to the former upstream name `alohadb_realtime`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
