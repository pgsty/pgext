

## 用法

> [postgres_fdw: 远程 PostgreSQL 服务器的外部数据包装器](https://www.postgresql.org/docs/current/postgres-fdw.html)

### 创建服务器

```sql
CREATE EXTENSION postgres_fdw;

CREATE SERVER remote_pg FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '192.168.1.10', port '5432', dbname 'remotedb');
```

**服务器选项：** 任何 libpq 连接参数（`host`、`port`、`dbname` 等），以及 `use_remote_estimate`（默认 `false`）、`fdw_startup_cost`（默认 `100`）、`fdw_tuple_cost`（默认 `0.2`）、`extensions`（两端都安装的扩展列表，逗号分隔）、`fetch_size`（默认 `100`）、`batch_size`（默认 `1`）、`keep_connections`（默认 `on`）、`parallel_commit`（默认 `false`）、`parallel_abort`（默认 `false`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR local_user SERVER remote_pg
  OPTIONS (user 'remote_user', password 'secret');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE remote_table (
  id integer NOT NULL,
  name text,
  value numeric
)
SERVER remote_pg
OPTIONS (schema_name 'public', table_name 'source_table');
```

**表/列选项：** `schema_name`（默认：本地模式名）、`table_name`（默认：本地表名）、`column_name`（按列设置，远程列名）、`updatable`（默认 `true`）、`truncatable`（默认 `true`）、`async_capable`（默认 `false`）。

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA remote_schema
  FROM SERVER remote_pg
  INTO local_schema;

-- 导入特定表
IMPORT FOREIGN SCHEMA remote_schema
  LIMIT TO (table1, table2)
  FROM SERVER remote_pg
  INTO local_schema;
```

**导入选项：** `import_collate`（默认 `true`）、`import_default`（默认 `false`）、`import_generated`（默认 `true`）、`import_not_null`（默认 `true`）。

### CRUD 操作

```sql
SELECT * FROM remote_table WHERE id > 100;
INSERT INTO remote_table VALUES (1, 'test', 42.0);
UPDATE remote_table SET value = 100 WHERE id = 1;
DELETE FROM remote_table WHERE id = 1;
TRUNCATE remote_table;
```

### 查询优化

postgres_fdw 自动下推 WHERE 子句、同一服务器上表之间的 JOIN、聚合函数、ORDER BY 和 LIMIT/OFFSET。使用以下命令查看远程查询：

```sql
EXPLAIN VERBOSE SELECT * FROM remote_table WHERE id = 1;
```

使用 `extensions` 选项允许这些扩展的函数/操作符下推：

```sql
ALTER SERVER remote_pg OPTIONS (ADD extensions 'postgis,hstore');
```

### 异步执行

启用跨多个外部服务器的并发扫描：

```sql
ALTER FOREIGN TABLE remote_table OPTIONS (ADD async_capable 'true');
```

### 连接管理

```sql
SELECT * FROM postgres_fdw_get_connections(true);   -- 列出连接
SELECT postgres_fdw_disconnect('remote_pg');         -- 关闭特定连接
SELECT postgres_fdw_disconnect_all();                -- 关闭所有连接
```

### 事务行为

如果本地事务为 SERIALIZABLE，远程事务使用 SERIALIZABLE；否则使用 REPEATABLE READ。目前不支持两阶段提交。
