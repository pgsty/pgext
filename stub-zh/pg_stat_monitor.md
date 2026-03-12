


## 用法

> [pg_stat_monitor: PostgreSQL 查询性能监控工具](https://github.com/percona/pg_stat_monitor)

pg_stat_monitor 是 `pg_stat_statements` 的高级替代方案，它将统计信息聚合到可配置的时间桶中，提供查询来源信息、实际参数捕获和查询计划详情。

### 查询统计信息

```sql
-- 基本查询监控
SELECT application_name, userid::regrole AS user_name,
       datname AS database_name, substr(query, 0, 50) AS query,
       calls, client_ip
FROM pg_stat_monitor;

-- 基于时间桶的分析
SELECT bucket, bucket_start_time, query, calls,
       mean_exec_time, total_exec_time
FROM pg_stat_monitor
ORDER BY total_exec_time DESC;

-- 显示查询计划
SELECT query, query_plan FROM pg_stat_monitor
WHERE query_plan IS NOT NULL;
```

### 主要特性

- **时间桶**：统计信息按可配置的时间间隔分组，提供更精确的分析
- **客户端 IP 追踪**：每个查询记录发起的客户端 IP 地址
- **实际参数**：可选捕获真实查询参数值而非占位符
- **查询计划**：每个查询附带其实际执行计划
- **热点查询追踪**：识别每个时间桶中最耗资源的查询
- **直方图支持**：通过直方图函数提供时间分布可视化

### 函数

```sql
-- 重置所有统计信息
SELECT pg_stat_monitor_reset();

-- 查看内部信息
SELECT * FROM pg_stat_monitor_info;
```

### 配置

关键参数（在 `postgresql.conf` 中设置）：

| 参数 | 描述 |
|-----------|-------------|
| `pg_stat_monitor.pgsm_max` | 最大追踪语句数 |
| `pg_stat_monitor.pgsm_query_max_len` | 最大查询长度 |
| `pg_stat_monitor.pgsm_bucket_time` | 时间桶持续时间（秒） |
| `pg_stat_monitor.pgsm_max_buckets` | 最大时间桶数量 |
| `pg_stat_monitor.pgsm_enable_query_plan` | 启用查询计划捕获 |
| `pg_stat_monitor.pgsm_track` | 追踪级别：`top`、`all` 或 `none` |
