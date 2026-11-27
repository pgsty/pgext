## 用法

来源：

- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/README.md)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/mysql_fdw.control)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/mysql_fdw--1.2.sql)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/releases/tag/REL-2_10_0)

`mysql_fdw` 2.10.0 通过 PostgreSQL 14–19 访问 MySQL 表，SQL/control 版本为 1.2。服务器需要兼容的 MySQL 客户端库。使用超级用户安装，无需共享预加载。

### 连接与映射

为已有表创建服务器和权限受限的远端账户映射：

```sql
CREATE EXTENSION mysql_fdw;
CREATE SERVER mysql_server FOREIGN DATA WRAPPER mysql_fdw
  OPTIONS (host '127.0.0.1', port '3306');
CREATE USER MAPPING FOR CURRENT_USER SERVER mysql_server
  OPTIONS (username 'mysql_user', password 'replace_with_remote_password');
CREATE FOREIGN TABLE warehouse (warehouse_id integer, warehouse_name text)
  SERVER mysql_server OPTIONS (dbname 'mydb', table_name 'warehouse');
SELECT * FROM warehouse;
```

必须指定 `dbname`，`table_name` 默认使用外部表名。该 FDW 也支持 `IMPORT FOREIGN SCHEMA`。仅向预期角色授予服务器和表访问权，保护映射凭据，并在不可信网络上使用验证证书的 TLS。

### 写入与下推

`UPDATE` 和 `DELETE` 使用外部表第一列作为行标识，该列必须唯一确定远端行。2.10.0 在该列属于复合键时拒绝这两种操作，因为旧版本可能影响共享该列值的多行。本地约束应与远端数据一致。远端写入有独立的事务及故障边界。

该 FDW 可下推过滤、连接、聚合、排序及行数限制，依赖索引或下推行为前应检查 `EXPLAIN`。`fetch_size`、`use_remote_estimate`、`reconnect` 影响远端访问。排序优化信任本地 NOT NULL 声明，如果远端列可为空而本地声明不一致，排序结果可能不同。

### 升级与安全变化

2.10.0 修复远端及本地 SQL 注入、生成列 INSERT 处理、重复参数查询以及二进制和文本缓冲区读取。`init_command`、`mysql_default_file` 和 SSL 路径及选项仅限超级用户设置。其他用户创建且包含受限选项的已有外部服务器，需要超级用户修改或重建。

远端 `sql_mode` 始终加入 `NO_BACKSLASH_ESCAPES`。导入的默认值现在是字符串字面量，导入后需在外部表上重新设置 `CURRENT_TIMESTAMP` 等表达式默认值。安装新库并重连已有会话，SQL 版本 1.2 不变。恢复应用访问前，应复核可写表的键、默认值、转义行为和影响行数。
