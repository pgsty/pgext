## Usage

Sources:

- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/README.md)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/hdfs_fdw.control)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/hdfs_fdw--2.0.5.sql)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/releases/tag/v2.4.0)

`hdfs_fdw` connects PostgreSQL to HiveServer2 or a Spark SQL service. Distribution 2.4.0 supports PostgreSQL 14–19 and retains SQL/control version 2.0.5. It needs the compatible JVM/JDBC client and a reachable remote service; installing the extension alone does not provide HDFS or Hive.

### Connect and Query

Install as a superuser. Shared preload is not required. With a remote table already present, replace the example endpoint and mapping credentials:

```sql
CREATE EXTENSION hdfs_fdw;
CREATE SERVER hive_server FOREIGN DATA WRAPPER hdfs_fdw
  OPTIONS (host '127.0.0.1', port '10000', client_type 'hiveserver2');
CREATE USER MAPPING FOR CURRENT_USER SERVER hive_server
  OPTIONS (username 'hive_user', password 'replace_with_remote_password');
CREATE FOREIGN TABLE hive_events (event_id bigint, event_type text)
  SERVER hive_server OPTIONS (dbname 'default', table_name 'events');
SELECT event_type, count(*) FROM hive_events GROUP BY event_type;
```

The mapping must provide the username required by the JDBC authentication path. Select `auth_type` deliberately for the remote service; do not depend on dummy credentials. Grant server usage and foreign-table access only to intended roles. Remote credentials remain in user mappings and need restricted access.

### Options and Plans

`client_type` selects `hiveserver2` or `spark`. Server settings include `connect_timeout`, `fetch_size`, `use_remote_estimate` and `log_remote_sql`. Table options `dbname` and `table_name` select the remote object. Version 2.4.0 accepts only alphanumeric characters and underscores in those two names; rename incompatible objects before upgrading.

Join, aggregate, ordering and limit pushdown reduce remote work. Use `EXPLAIN` and the `hdfs_fdw.enable_join_pushdown`, `hdfs_fdw.enable_aggregate_pushdown`, `hdfs_fdw.enable_order_by_pushdown` and `hdfs_fdw.enable_limit_pushdown` settings to examine plans. Remote SQL semantics, network latency and authentication still apply.

### Upgrade Boundaries

This release fixes HiveQL string-literal injection and JVM-load error handling. Install the updated library/JDBC components and reconnect processes using the old library. An existing SQL version 2.0.5 remains unchanged. `log_remote_sql` includes literal values in server logs; restrict log access. Credential masking improves this release, but it does not make SQL logging suitable for secrets.
