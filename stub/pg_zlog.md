## Usage

Sources:

- [Upstream architecture and configuration](https://github.com/cruzdb/pg_zlog/blob/5d428db977b0eec10c07f155a30db54b2c1af51c/README.md)
- [Extension control file](https://github.com/cruzdb/pg_zlog/blob/5d428db977b0eec10c07f155a30db54b2c1af51c/pg_zlog.control)

`pg_zlog` is a research extension for replicating selected tables through a ZLog shared log stored on Ceph. A node rolls the log forward before accessing a registered table, and metadata records the expected Ceph cluster and pool identifiers.

```sql
CREATE EXTENSION pg_zlog;
SELECT pgzlog_add_cluster('ceph', '/etc/ceph/ceph.conf');
SELECT pgzlog_add_pool('ceph', 'pg_pool');
SELECT pgzlog_add_log('pg_pool', 'pg_log', 'seq-host', 5678);
SELECT pgzlog_replicate_table('pg_log', 'coordinates');
```

This project is abandoned and depends on obsolete ZLog/Ceph integration. It is not a substitute for supported PostgreSQL physical or logical replication. Use only in an isolated historical test environment; verify extension preload requirements, Ceph credentials, cluster identity, replay ordering, DDL behavior, failure recovery, and data consistency before evaluating any result.
