

## 用法

> [mysql_fdw: 查询 MySQL 服务器的外部数据包装器](https://github.com/EnterpriseDB/mysql_fdw)

### 创建服务器

```sql
CREATE EXTENSION mysql_fdw;

CREATE SERVER mysql_server
  FOREIGN DATA WRAPPER mysql_fdw
  OPTIONS (host '127.0.0.1', port '3306');
```

**服务器选项：** `host`（默认 `127.0.0.1`）、`port`（默认 `3306`）、`secure_auth`（默认 `true`）、`init_command`、`use_remote_estimate`（默认 `false`）、`reconnect`（默认 `false`）、`sql_mode`（默认 `ANSI_QUOTES`）、`fetch_size`（默认 `100`）、`character_set`、`truncatable`（默认 `true`），以及 SSL 选项（`ssl_key`、`ssl_cert`、`ssl_ca`、`ssl_capath`、`ssl_cipher`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR pguser
  SERVER mysql_server
  OPTIONS (username 'mysqluser', password 'mysqlpass');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE warehouse (
  warehouse_id int,
  warehouse_name text,
  warehouse_created timestamp
)
SERVER mysql_server
OPTIONS (dbname 'mydb', table_name 'warehouse');
```

**表选项：** `dbname`（必填，MySQL 数据库名）、`table_name`（默认为外部表名）、`fetch_size`（覆盖服务器设置）、`max_blob_size`、`truncatable`（默认 `true`）。

### CRUD 操作

```sql
INSERT INTO warehouse VALUES (1, 'UPS', current_date);
SELECT * FROM warehouse ORDER BY warehouse_id;
UPDATE warehouse SET warehouse_name = 'NEW_NAME' WHERE warehouse_id = 1;
DELETE FROM warehouse WHERE warehouse_id = 3;
```

### 下推特性

mysql_fdw 通过多种下推机制优化查询：

- **WHERE 子句** 下推以最小化数据传输
- **列** 下推以仅获取请求的列
- **JOIN** 下推用于同一 MySQL 服务器上外部表之间的连接
- **聚合** 下推支持 `min`、`max`、`sum`、`avg`、`count`
- **ORDER BY** 和 **LIMIT/OFFSET** 下推以减少网络流量
