## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering.c)

`pg_tiering` `1.0` provides rule-driven partition tablespace movement. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_tiering;
SELECT * FROM pg_tiering_status();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Rules in pg_tiering_rules identify parent tables, age thresholds and destination tablespaces. pg_tiering_check_now runs checks manually; periodic work requires preloading pg_tiering, restart, pg_tiering.database and pg_tiering.check_interval. Moving a relation uses ALTER TABLE SET TABLESPACE and consumes locks and I/O. Verify free space, ownership and partition-age selection before enabling rules; this is local tablespace movement, not object-storage archiving.

This module corresponds to the former upstream name `alohadb_tiering`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
