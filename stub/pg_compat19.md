## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/pg_compat19.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/pg_compat19--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/meson.build)

`pg_compat19` `1.0` provides sql compatibility helpers for the alohadb postgresql 18 base. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_compat19;
SELECT random('2026-01-01'::date, '2026-12-31'::date);
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Installed into pg_catalog by a superuser, this module adds error_on_null, tuple-ID helpers, date/time random ranges, range subtraction and UUID/bytea casts. It depends on matching AlohaDB internal functions. Remove the compatibility extension before upgrading to an AlohaDB kernel based on PostgreSQL 19, where these objects are built in; review dependent objects before removal.

This module corresponds to the former upstream name `alohadb_pg19compat`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
