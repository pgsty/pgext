

## 用法

> [tds_fdw: 查询 TDS 数据库（Sybase 或 Microsoft SQL Server）的外部数据包装器](https://github.com/tds-fdw/tds_fdw)

### 创建服务器

```sql
CREATE EXTENSION tds_fdw;

CREATE SERVER mssql_svr
  FOREIGN DATA WRAPPER tds_fdw
  OPTIONS (servername '127.0.0.1', port '1433',
           database 'tds_fdw_test', tds_version '7.1');
```

**服务器选项：** `servername`（服务器地址或 DSN，支持逗号分隔的故障转移列表）、`port`、`database`、`dbuse`（0 为直接连接，非0 使用 dbuse()）、`tds_version`（协议版本）、`language`、`character_set`、`msg_handler`（`notice` 或 `blackhole`）、`sqlserver_ansi_mode`、`fdw_startup_cost`、`fdw_tuple_cost`。

### 创建用户映射

```sql
CREATE USER MAPPING FOR postgres
  SERVER mssql_svr
  OPTIONS (username 'sa', password 'secret');
```

对于 Azure SQL 数据库，`username` 选项使用 `username@servername` 格式。

### 创建外部表

直接映射远程表：

```sql
CREATE FOREIGN TABLE mssql_table (
  id integer,
  name varchar(255),
  value numeric(10,2)
)
SERVER mssql_svr
OPTIONS (schema_name 'dbo', table_name 'mytable');
```

或使用自定义 SQL 查询：

```sql
CREATE FOREIGN TABLE mssql_query (
  id integer,
  name varchar(255),
  total numeric(10,2)
)
SERVER mssql_svr
OPTIONS (query 'SELECT id, name, SUM(amount) AS total FROM orders GROUP BY id, name');
```

**表选项：** `table_name` 或 `query`（二选一，互斥）、`schema_name`、`match_column_names`（按名称匹配还是按位置匹配）、`use_remote_estimate`、`row_estimate_method`（`execute` 或 `showplan_all`）。

**列选项：** `column_name`（本地列名不同于远程时使用）。

### 查询和调试

```sql
SELECT * FROM mssql_table WHERE id > 100;

-- 查看发送到 SQL Server 的远程查询
EXPLAIN (VERBOSE) SELECT * FROM mssql_table WHERE id > 100;
```

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA dbo
  FROM SERVER mssql_svr
  INTO public;
```
