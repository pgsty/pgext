## 用法

来源：

- [在 ApsaraDB RDS 主实例上启用 DuckDB](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/enable-duckdb-for-the-master-instance)
- [AP 加速引擎（`rds_duckdb`）](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/how-to-use-duckdb-to-speed-up-queries/)
- [DuckDB 分析加速概览](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/duckdb-analytics-acceleration/)

`rds_duckdb` 1.3.2.3 是阿里云 ApsaraDB RDS for PostgreSQL 扩展。它为 PostgreSQL 表维护 DuckDB 列存副本，并将符合条件的分析型 `SELECT` 下推到 DuckDB 向量化引擎。它适用于 RDS 实例上的扫描、连接与聚合，不是面向社区 PostgreSQL 的可移植 DuckDB 扩展。

### 核心流程

先在 RDS 控制台为实例启用 DuckDB 加速。托管流程会配置必要的实例参数，包括 `wal_level=logical` 以及在 `shared_preload_libraries` 中加入 `rds_duckdb`，过程中会发生短暂断连。然后用高权限账号连接：

```sql
CREATE EXTENSION IF NOT EXISTS rds_duckdb;

SELECT rds_duckdb.create_duckdb_table('public.orders');
-- Batch form:
SELECT rds_duckdb.create_duckdb_tables(
  '{public.customers,public.order_items}'
);
```

按会话或单条语句启用下推，并用 `EXPLAIN` 确认计划中出现 `Custom Scan (DuckDBScan)`：

```sql
SET rds_duckdb.execution = on;
EXPLAIN SELECT customer_id, sum(total)
FROM orders
GROUP BY customer_id;

/*+ set(rds_duckdb.execution on) */
EXPLAIN SELECT count(*) FROM orders;
```

### 控制与观测

- `rds_duckdb.execution` 启用或禁用查询下推；语句级 hint 只支持设置这一项。
- `rds_duckdb.plan_cost_threshold` 让低成本查询留在 PostgreSQL，避免 DuckDB 转发开销。
- `rds_duckdb.worker_threads` 与 `rds_duckdb.memory_limit` 限制会话级 DuckDB 资源。
- `rds_duckdb.enable_fallback` 控制无法下推时是否回退 PostgreSQL；排障时可临时关闭以暴露原因。
- `rds_duckdb.wait_sync_timeout` 限制等待增量同步的时间。
- `rds_duckdb.duckdb_sync_stat` 报告表的同步状态与错误，需要高权限账号查询。

### 依赖与注意事项

当前厂商文档要求主实例为 PostgreSQL 13–18、RDS 小版本不低于 20260130；只读实例加速及更新特性的要求更窄。目录中的扩展版本与 RDS 内核小版本是两个不同标识。

只有符合条件的只读 `SELECT` 会下推。DML、DDL、不支持的 SQL、包含非 DuckDB 表的连接，或尚未同步的表会回退 PostgreSQL，除非显式禁用 fallback。同步表需要主键或合适的 replica identity。列存副本占用额外存储，并可能落后于 PostgreSQL 数据；应明确选择一致性与超时设置。可用性、重启、恢复、复制、升级和计费行为仍由阿里云服务控制。
