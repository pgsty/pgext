## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs.c)

`pg_codefs` `1.0` provides versioned database-backed file workspace. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_codefs;
SELECT codefs_write('src/demo.txt', 'hello', 'main', 'initial');
SELECT codefs_read('src/demo.txt');
SELECT * FROM codefs_history('src/demo.txt');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Relative paths, versions, branches, locks and policy tables form the SQL workspace. Disk synchronization and policy workers require shared_preload_libraries containing pg_codefs and a restart; configure pg_codefs.database, pg_codefs.sync_interval and pg_codefs.policy_interval for that optional mode. Policy changes are superuser-controlled. Disk synchronization is an external effect and needs filesystem isolation; protect stored source and history with grants.

This module corresponds to the former upstream name `alohadb_codefs`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
