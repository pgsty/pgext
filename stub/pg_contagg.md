## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_contagg/pg_contagg.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_contagg/pg_contagg--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_contagg/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_contagg/pg_contagg.c)

`pg_contagg` `1.0` provides materialized aggregate management and time-series functions. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_contagg;
SELECT * FROM analytics_cont_agg_status();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

analytics_create_cont_agg registers a materialized view; analytics_refresh_cont_agg performs a full REFRESH MATERIALIZED VIEW. The refresh interval is stored metadata, not proof of automatic incremental maintenance. Projection triggers also perform full refreshes. Limit query-definition access to trusted owners and schedule refreshes explicitly.

This module corresponds to the former upstream name `alohadb_analytics`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
