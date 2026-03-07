
## 用法

[`pg_mooncake`](https://github.com/Mooncake-Labs/pg_mooncake) 是一个 Postgres 扩展，能够为表创建 Iceberg 格式的列存镜像，设计为 `pg_duckdb` 的子扩展。

pg_mooncake 文档：https://docs.mooncake.dev/


### 快速上手

使用 pig 安装 pg_duckdb 和 pg_mooncake：

```bash
pig repo set
pig install pg_duckdb pg_mooncake
```

编辑 postgresql.conf，然后重启使配置生效

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
duckdb.allow_community_extensions = true
wal_level = logical
```



### 入门示例

- [教程](https://docs.mooncake.dev/pg/get-started/Hello-world)

```sql
-- 连同 pg_duckdb 一起创建扩展
CREATE EXTENSION pg_mooncake CASCADE;

-- 接下来，创建一张普通的 Postgres 表 trades：
CREATE TABLE trades(
  id bigint PRIMARY KEY,
  symbol text,
  time timestamp,
  price real
);

-- 然后，创建一个列存镜像 trades_iceberg，与 trades 保持同步：
CALL mooncake.create_table('trades_iceberg', 'trades');

-- 现在，向 trades 中插入一些数据：
INSERT INTO trades VALUES
    (1,  'AMD', '2024-06-05 10:00:00', 119),
    (2, 'AMZN', '2024-06-05 10:05:00', 207),
    (3, 'AAPL', '2024-06-05 10:10:00', 203),
    (4, 'AMZN', '2024-06-05 10:15:00', 210);

-- 最后，使用 duckdb 进行查询
EXPLAIN
    SELECT avg(price) FROM trades_iceberg WHERE symbol = 'AMZN';
```

执行计划中将显示 DuckDBScan：

```bash
                         QUERY PLAN
------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0)
   DuckDB Execution Plan:

 ┌───────────────────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │    Aggregates: avg(#0)    │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         PROJECTION        │
 │    ────────────────────   │
 │   CAST(price AS DOUBLE)   │
 │                           │
 │          ~0 rows          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │       MOONCAKE_SCAN       │
 │    ────────────────────   │
 │   Table: trades_iceberg   │
 │     Projections: price    │
 │                           │
 │          Filters:         │
 │       symbol='AMZN'       │
 │                           │
 │          ~0 rows          │
 └───────────────────────────┘
```
