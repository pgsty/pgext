

## 用法

[pg_duckdb 文档](https://github.com/duckdb/pg_duckdb/tree/main/docs)

| 主题                                                                                                   | 说明                                                       |
|:-------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------|
| [**函数**](https://github.com/duckdb/pg_duckdb/blob/main/docs/functions.md)                            | 所有可用函数的完整参考                                      |
| [**语法指南与注意事项**](https://github.com/duckdb/pg_duckdb/blob/main/docs/gotchas_and_syntax.md)      | 常见 SQL 模式与注意事项速查                                 |
| [**类型**](https://github.com/duckdb/pg_duckdb/blob/main/docs/types.md)                                | 支持的数据类型与类型映射                                    |
| [**扩展**](https://github.com/duckdb/pg_duckdb/blob/main/docs/extensions.md)                           | DuckDB 扩展的安装与使用                                     |
| [**配置**](https://github.com/duckdb/pg_duckdb/blob/main/docs/settings.md)                             | 配置选项与参数                                              |
| [**事务**](https://github.com/duckdb/pg_duckdb/blob/main/docs/transactions.md)                         | 事务行为与限制                                              |



### 快速上手

使用 pig 安装 pg_duckdb：

```bash
pig repo set
pig install pg_duckdb
```

编辑 `postgresql.conf`，然后重启使配置生效

```ini
shared_preload_libraries = 'pg_duckdb'
duckdb.allow_community_extensions = true
```


### 加速查询

通过 DuckDB 可以直接查询现有的 PostgreSQL 表，无需做任何修改：

```sql
-- pgbench -is 1000  # 使用 pgbench 初始化一些测试数据
CREATE EXTENSION pg_duckdb;

-- 默认行为，使用 PostgreSQL 原生引擎
SET duckdb.force_execution = false;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;

-- 现在查询将由 pg_duckdb 执行
SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
```

在本地笔记本电脑的 4 核虚拟机上，查询耗时从 8 秒降至 4 秒：

```
postgres@el9:5432/postgres=# SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
SET
Time: 0.206 ms
                                              QUERY PLAN
------------------------------------------------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0) (actual time=0.001..0.001 rows=0 loops=1)
   DuckDB Execution Plan:

 ┌─────────────────────────────────────┐
 │┌───────────────────────────────────┐│
 ││    Query Profiling Information    ││
 │└───────────────────────────────────┘│
 └─────────────────────────────────────┘
 EXPLAIN ANALYZE  SELECT count(*) AS count FROM pgduckdb.public.pgbench_accounts
 ┌────────────────────────────────────────────────┐
 │┌──────────────────────────────────────────────┐│
 ││               Total Time: 3.89s              ││
 │└──────────────────────────────────────────────┘│
 └────────────────────────────────────────────────┘
 ┌───────────────────────────┐
 │           QUERY           │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │      EXPLAIN_ANALYZE      │
 │    ────────────────────   │
 │           0 rows          │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │        Aggregates:        │
 │        count_star()       │
 │                           │
 │           1 row           │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         TABLE_SCAN        │
 │    ────────────────────   │
 │           Table:          │
 │      pgbench_accounts     │
 │                           │
 │      100,000,000 rows     │
 │          (3.88s)          │
 └───────────────────────────┘
```



### 数据湖

以下示例演示如何使用本地 MinIO 实例操作数据湖：

```sql
SELECT duckdb.create_simple_secret(
    type := 'S3', key_id := 's3user_data', secret := 'S3User.Data',
    endpoint := 'https://sss.pigsty:9000', url_style := 'path'
);
```
