## 用法

- 来源：[README](https://github.com/tds-fdw/tds_fdw/blob/master/README.md), [foreign server](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignServerCreation.md), [foreign table](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignTableCreation.md), [user mapping](https://github.com/tds-fdw/tds_fdw/blob/master/UserMappingCreation.md), [foreign schema](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignSchemaImporting.md), [variables](https://github.com/tds-fdw/tds_fdw/blob/master/Variables.md)

`tds_fdw` 是一个 foreign data wrapper，用于通过 FreeTDS 等 DB-Library 实现查询 Sybase 和 Microsoft SQL Server 等 TDS 数据库。

### 创建服务器

```sql
CREATE EXTENSION tds_fdw;

CREATE SERVER mssql_svr
  FOREIGN DATA WRAPPER tds_fdw
  OPTIONS (servername '127.0.0.1', port '1433',
           database 'tds_fdw_test', tds_version '7.1');
```

**服务器选项：** `servername`（服务器地址或 DSN，支持逗号分隔 failover list）、`port`、`database`、`dbuse`（0 表示 direct connection，非 0 表示 dbuse()）、`tds_version`（protocol version）、`language`、`character_set`、`msg_handler`（`notice` 或 `blackhole`）、`sqlserver_ansi_mode`、`fdw_startup_cost`、`fdw_tuple_cost`。

### 创建用户映射

```sql
CREATE USER MAPPING FOR postgres
  SERVER mssql_svr
  OPTIONS (username 'sa', password 'secret');
```

对于 Azure SQL databases，`username` 选项使用 `username@servername` 格式。

### 创建外部表

直接映射远端表：

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

**表选项：** `table_name` 或 `query`（二者必选其一，且互斥）、`schema_name`、`match_column_names`（按名称还是按位置映射）、`use_remote_estimate`、`local_tuple_estimate`、`row_estimate_method`（`execute` 或 `showplan_all`）。

**列选项：** `column_name`（远端列名不同于本地列名时使用）。

### 查询与调试

```sql
SELECT * FROM mssql_table WHERE id > 100;

-- View the remote query sent to SQL Server
EXPLAIN (VERBOSE) SELECT * FROM mssql_table WHERE id > 100;
```

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA dbo
  FROM SERVER mssql_svr
  INTO public
  OPTIONS (import_default 'true');
```

**导入选项：** `import_default`、`import_not_null`，以及 `keep_custom_types`，后者用于在已存在匹配 PostgreSQL domains 时保留 Sybase user-defined types。

### Planner 与运行时说明

上游 README 说明当前版本不支持 join pushdown 或写操作。当 `match_column_names` 启用时，它支持 `WHERE` 和列下推。

使用 PostgreSQL `SET` 设置诊断 memory-stat variables，例如：

```sql
SET tds_fdw.show_finished_memory_stats = 1;
```

可用变量包括 `tds_fdw.show_before_row_memory_stats`、`tds_fdw.show_after_row_memory_stats` 和 `tds_fdw.show_finished_memory_stats`。

Pigsty package metadata 是来自 PGDG 的版本 `2.0.5`，覆盖 PostgreSQL 14-18。上游文档说该 FDW 应支持 PostgreSQL 9.2+，当前 build matrix 包含 PostgreSQL 13-18，但此 stub 遵循 `db/extension.csv` 中的已打包 PostgreSQL 版本。
