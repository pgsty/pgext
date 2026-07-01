

## 用法

来源：[README](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_5_3/README.md)、[REL-5_5_3 release](https://github.com/EnterpriseDB/mongo_fdw/releases/tag/REL-5_5_3)

### 创建服务器

```sql
CREATE EXTENSION mongo_fdw;

CREATE SERVER mongo_server FOREIGN DATA WRAPPER mongo_fdw
  OPTIONS (address '127.0.0.1', port '27017');
```

**服务器选项：** `address`（默认 `127.0.0.1`）、`port`（默认 `27017`）、`authentication_database`、`replica_set`、`read_preference`（`primary`、`secondary`、`primaryPreferred`、`secondaryPreferred`、`nearest`）、`ssl`（默认 `false`）、`pem_file`、`pem_pwd`、`ca_file`、`ca_dir`、`crl_file`、`weak_cert_validation`、`use_remote_estimate`（默认 `false`）、`enable_join_pushdown`（默认 `true`）、`enable_aggregate_pushdown`（默认 `true`）、`enable_order_by_pushdown`（默认 `true`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR pguser SERVER mongo_server
  OPTIONS (username 'mongouser', password 'mongopass');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE warehouse (
  _id name,
  warehouse_id int,
  warehouse_name text,
  warehouse_created timestamptz
)
SERVER mongo_server
OPTIONS (database 'mydb', collection 'warehouse');
```

**重要：** 第一列必须是 `name` 类型的 `_id`（MongoDB 的对象标识符）。

**表选项：** `database`（默认 `test`）、`collection`（默认为表名）、`enable_join_pushdown`、`enable_aggregate_pushdown`、`enable_order_by_pushdown`。

### CRUD 操作

```sql
SELECT warehouse_id, warehouse_name FROM warehouse WHERE warehouse_id > 10;
INSERT INTO warehouse VALUES ('new_id', 100, 'New Warehouse', now());
UPDATE warehouse SET warehouse_name = 'Updated' WHERE warehouse_id = 100;
DELETE FROM warehouse WHERE warehouse_id = 100;
```

### 下推特性

mongo_fdw 将 WHERE 子句、同一服务器上外部表之间的 JOIN、聚合函数、ORDER BY、LIMIT 和 OFFSET 下推到 MongoDB，以实现高效查询执行。诊断远端执行计划时，可使用 `mongo_fdw.enable_join_pushdown`、`mongo_fdw.enable_aggregate_pushdown`、`mongo_fdw.enable_order_by_pushdown` 和 `mongo_fdw.log_remote_query` GUC。

### 版本说明

`mongo_fdw` 5.5.3，上游 tag 为 `REL-5_5_3`，增加 PostgreSQL 18 支持，为 MongoDB 8 更新 bundled `mongoc-driver` 和 `json-c` libraries，增加 `mongo_fdw.log_remote_query` debug GUC，并修复 nested-field、WHERE、ORDER BY 和 unsafe join-pushdown 场景。本线中上游已停止 PostgreSQL 12 支持。

### 注意事项

- BSON 仅支持 UTF-8；确保 PostgreSQL 数据库使用 UTF-8 编码
- 包含大写字母或点号（用于嵌套文档）的列名需要双引号
- PostgreSQL 默认将列名限制为 63 个字符
