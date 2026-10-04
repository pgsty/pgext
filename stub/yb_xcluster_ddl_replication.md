## Usage

Sources:

- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.control](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.control)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/README.md](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/README.md)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication--1.0.sql](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication--1.0.sql)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.c](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.c)
- [LICENSE.md](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/LICENSE.md)

`yb_xcluster_ddl_replication` is a YugabyteDB-specific extension used by xCluster automatic DDL replication. It records source DDL and target replay state; the external xCluster handler performs the replay.

### Core Workflow

Enable automatic xCluster replication through YugabyteDB's management workflow first. At the cited revision, YugabyteDB includes the library in its default `shared_preload_libraries`; xCluster setup creates the extension in participating databases on both universes. Then inspect the database's replication role:

```sql
SELECT yb_xcluster_ddl_replication.get_replication_role();
```

### Operational Boundaries

The fixed schema contains `ddl_queue`, `replicated_ddls` and event-trigger handlers. Only `get_replication_role()` is granted to PUBLIC; internal tables and replay settings are not an application API. Target-side user DDL is restricted. Extensions whose scripts contain complex DDL such as table creation must be installed before adding the database to automatic replication.

The reviewed control version is 1.0 at the cited source revision; this is not a portable vanilla PostgreSQL extension or an independent replication service.
