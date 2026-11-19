## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config.c)

`pg_remote_config` `1.0` provides remote configuration status and management tables. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_remote_config;
SELECT * FROM remote_config_status();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

The remote mode requires preloading pg_remote_config and a restart. Configure pg_remote_config.mode, pg_remote_config.host, pg_remote_config.port and the private pg_remote_config.password through administrator-controlled configuration. This is a server control-plane integration: changing the remote source can change server configuration. The status SQL alone does not enable remote management. Keep management tables, tokens and network access restricted.

This module corresponds to the former upstream name `alohadb_remote_config`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
