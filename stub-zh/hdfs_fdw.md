

## 用法

> [hdfs_fdw: 远程 HDFS 服务器的外部数据包装器](https://github.com/EnterpriseDB/hdfs_fdw)

### 创建服务器

```sql
CREATE EXTENSION hdfs_fdw;

CREATE SERVER hdfs_server FOREIGN DATA WRAPPER hdfs_fdw
  OPTIONS (host '127.0.0.1', port '10000', client_type 'hiveserver2');
```

**服务器选项：** `host`（默认 `localhost`）、`port`（默认 `10000`）、`client_type`（`hiveserver2` 或 `spark`，默认 `hiveserver2`）、`auth_type`（`NOSASL` 或 `LDAP`）、`connect_timeout`（默认 300）、`fetch_size`（默认 10000）、`log_remote_sql`（默认 `false`）、`use_remote_estimate`（默认 `false`）、`enable_join_pushdown`（默认 `true`）、`enable_aggregate_pushdown`（默认 `true`）、`enable_order_by_pushdown`（默认 `true`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR postgres SERVER hdfs_server
  OPTIONS (username 'hive_user', password 'hive_password');
```

对于 NOSASL 认证，完全省略 OPTIONS 子句。

### 创建外部表

```sql
CREATE FOREIGN TABLE weblogs (
  client_ip text,
  http_status_code text,
  uri text,
  request_count bigint
)
SERVER hdfs_server
OPTIONS (dbname 'default', table_name 'weblogs');
```

**表选项：** `dbname`（默认 `default`）、`table_name`（默认为外部表名）、`enable_join_pushdown`、`enable_aggregate_pushdown`、`enable_order_by_pushdown`。

### 查询

```sql
SELECT client_ip, count(*) FROM weblogs GROUP BY client_ip ORDER BY count(*) DESC LIMIT 10;
```

### Spark 示例

```sql
CREATE SERVER spark_server FOREIGN DATA WRAPPER hdfs_fdw
  OPTIONS (host '127.0.0.1', port '10000', client_type 'spark');

CREATE USER MAPPING FOR postgres SERVER spark_server
  OPTIONS (username 'spark_user', password 'spark_pass');

CREATE FOREIGN TABLE spark_table (
  id int,
  name text,
  value double precision
)
SERVER spark_server
OPTIONS (dbname 'default', table_name 'my_table');
```

### 下推特性

hdfs_fdw 将 WHERE 子句、JOIN、聚合函数、ORDER BY 和 LIMIT/OFFSET 下推到远程 Hive/Spark 服务器。在会话级别控制下推：

```sql
SET hdfs_fdw.enable_join_pushdown = on;
SET hdfs_fdw.enable_aggregate_pushdown = on;
SET hdfs_fdw.enable_order_by_pushdown = on;
SET hdfs_fdw.enable_limit_pushdown = on;
```
