## Usage

Sources:

- [src/pg_fdw/README.md](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/README.md)
- [src/pg_fdw/springtail_fdw--1.0.sql](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/springtail_fdw--1.0.sql)
- [src/pg_fdw/CMakeLists.txt](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/CMakeLists.txt)
- [LICENSE.md](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/LICENSE.md)
- [src/pg_fdw/springtail_fdw.control](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/springtail_fdw.control)

`springtail_fdw` 1.0 is the SQL-facing storage adapter in Springtail. It depends on a configured Springtail deployment; installing the extension alone does not create a read-scaling cluster.

### Core Workflow

```sql
CREATE EXTENSION springtail_fdw;
CREATE SERVER springtail_fdw_server FOREIGN DATA WRAPPER springtail_fdw;
```

### Operational Boundaries

Configure `springtail_fdw.config_file_path` to the deployment's system.json before using the adapter. The upstream installation guide targets PostgreSQL 16, registers server `springtail_fdw_server`, and imports a Springtail namespace with IMPORT FOREIGN SCHEMA. The XID-manager daemon must be running with an XID at least as high as the table's XID. The broader deployment has additional daemon, Redis and replication configuration requirements. Installation defaults to superuser-only. The source is governed by Elastic License 2.0, including its hosted-service restrictions.
