## 用法

来源：

- [官方 README](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/README.md)
- [FDW 选项与回调实现](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/hive_fdw.c)
- [导入外部模式文档](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/IMPORT_FOREIGN_SCHEMA.md)

`hive_fdw` 版本 `4.0` 通过 JDBC 为 PostgreSQL 提供对 HiveServer2 的只读访问。它会在 PostgreSQL 进程内启动 JVM，并要求 OpenJDK 11 与匹配的 Hive 客户端 JAR。

### 运行时前提

启动 PostgreSQL 前，在服务器环境中设置 `PGHOME`、`HIVECLIENT_JAR_HOME` 和 `HIVE_FDW_CLASSPATH`。这些变量在 JVM 初始化时读取，因此修改后需要重启承载 FDW 的服务器进程。

### 核心流程

```sql
CREATE EXTENSION hive_fdw;

CREATE SERVER hive_serv FOREIGN DATA WRAPPER hive_fdw
OPTIONS (host 'localhost', port '10000');

CREATE USER MAPPING FOR CURRENT_USER SERVER hive_serv
OPTIONS (username 'hive_user', password 'secret');

CREATE FOREIGN TABLE orders (id integer)
SERVER hive_serv
OPTIONS (schema 'example', table 'orders');

SELECT * FROM orders LIMIT 5;
```

实现接受服务器上的 `host` 与 `port`、用户映射上的 `username` 与 `password`，以及外部表上的 `schema`、`table` 或自定义 `query`。README 正文把表选项称为 `schema_name` 与 `table_name`，但版本 `4.0` 的源码和示例实际校验的拼写是 `schema` 与 `table`。

### 导入与下推

`IMPORT FOREIGN SCHEMA` 可以根据 Hive 元数据创建本地外部表定义。该 FDW 还实现了连接下推。依赖下推或导入定义前，应使用 `EXPLAIN` 检查计划并验证类型映射。

### 运维说明

回调表没有插入、更新或删除方法，因此应把外部表视为只读。用户映射密码存储在 PostgreSQL 系统目录中，每个查询都跨越 JDBC/JVM 边界；应限制目录访问，使用最小权限 Hive 账号，并为远端延迟和 HiveServer2 故障做好规划。
