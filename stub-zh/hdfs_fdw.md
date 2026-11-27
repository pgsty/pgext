## 用法

来源：

- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/README.md)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/hdfs_fdw.control)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/blob/v2.4.0/hdfs_fdw--2.0.5.sql)
- [Official source](https://github.com/EnterpriseDB/hdfs_fdw/releases/tag/v2.4.0)

`hdfs_fdw` 将 PostgreSQL 连接到 HiveServer2 或 Spark SQL 服务。发行版 2.4.0 支持 PostgreSQL 14–19，保留 SQL/control 版本 2.0.5。它需要兼容的 JVM/JDBC 客户端和可访问的远端服务；仅安装扩展不会提供 HDFS 或 Hive。

### 连接与查询

使用超级用户安装，无需共享预加载。准备好远端表后，替换示例端点和映射凭据：

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

用户映射必须提供 JDBC 认证路径所需的用户名。应根据远端服务明确选择 `auth_type`，不能依赖虚拟凭据。仅向预期角色授予服务器使用权和外部表访问权。远端凭据保存在用户映射中，需要限制访问。

### 选项与执行计划

`client_type` 选择 `hiveserver2` 或 `spark`。服务器设置包括 `connect_timeout`、`fetch_size`、`use_remote_estimate` 和 `log_remote_sql`。表选项 `dbname`、`table_name` 指定远端对象。2.4.0 仅允许这两个名称包含字母、数字和下划线，升级前需重命名不兼容的对象。

连接、聚合、排序和行数限制下推可以减少远端工作。通过 `EXPLAIN` 以及 `hdfs_fdw.enable_join_pushdown`、`hdfs_fdw.enable_aggregate_pushdown`、`hdfs_fdw.enable_order_by_pushdown`、`hdfs_fdw.enable_limit_pushdown` 检查计划。远端 SQL 语义、网络延迟和认证要求仍然适用。

### 升级边界

该版本修复 HiveQL 字符串字面量注入及 JVM 加载错误处理。安装新的库与 JDBC 组件后，让使用旧库的进程重新连接。已有 SQL 版本 2.0.5 保持不变。`log_remote_sql` 会将字面量写入服务器日志，必须限制日志访问。此次凭据遮蔽改进不意味着 SQL 日志适合承载秘密。
