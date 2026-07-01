


## 用法

来源：[README](https://github.com/ClickHouse/pg_stat_ch/blob/main/README.md)、[release 0.3.6](https://github.com/ClickHouse/pg_stat_ch/releases/tag/v0.3.6)

`pg_stat_ch` 会捕获 PostgreSQL 的逐查询执行遥测，并通过共享内存队列和后台 worker 将原始事件导出到 ClickHouse。上游将它定位为 `pg_stat_statements` 的原始事件替代方案：聚合与看板交给 ClickHouse 处理，而不是留在 PostgreSQL 内部。

```sql
CREATE EXTENSION pg_stat_ch;
```

### 所需配置

`pg_stat_ch` 必须被 preload，并指向一个 ClickHouse 数据库：

```conf
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

README 说明 PostgreSQL 16、17 和 18 已完整支持；最新 release 是 2026-04-15 发布的 `0.3.6`。

### SQL API

- `pg_stat_ch_version()` 返回扩展版本。
- `pg_stat_ch_stats()` 暴露队列与导出器计数器。
- `pg_stat_ch_reset()` 清空队列计数器。
- `pg_stat_ch_flush()` 立即触发一次导出 flush。

```sql
SELECT pg_stat_ch_version();
SELECT * FROM pg_stat_ch_stats();
SELECT pg_stat_ch_flush();
```

### 重要 GUC

- `pg_stat_ch.enabled` 控制是否采集。
- `pg_stat_ch.queue_capacity` 和 `pg_stat_ch.string_area_size` 用于调整共享内存缓冲区大小。
- `pg_stat_ch.flush_interval_ms` 和 `pg_stat_ch.batch_max` 控制导出节奏和批量大小。
- `pg_stat_ch.log_min_elevel` 控制捕获哪些错误。

### 捕获内容

- 查询耗时、返回行数、缓冲区使用、WAL 使用量和 CPU 时间。
- DML、DDL 和 utility 语句。
- SQLSTATE 和错误级别。
- PostgreSQL 15+ 的 JIT 指标。
- PostgreSQL 18+ 的并行 worker 统计。
- 应用名、客户端 IP，以及在上游截断限制内的查询文本。

Release `0.3.6` 将 normalize cache 替换为以 queryId 为键的 LRU cache，并将 OpenTelemetry SDK 替换为 direct-protobuf exporter 加 sampling support。现有 SQL 用法仍围绕事件表、汇总视图，以及上面的 reset/version 辅助函数。

### 注意事项

- 该设计在队列溢出时会主动丢弃事件，而不是阻塞前台查询路径。
- ClickHouse schema 创建是必需的部署步骤；上游 quickstart 脚本会自动加载，但手工部署必须单独加载 schema。
- Pigsty 当前这次 RPM refresh 中，`pg_stat_ch` 是 PostgreSQL 16-18 且仅覆盖 EL9/EL10；Debian/Ubuntu 包仍覆盖活跃目标上的 PostgreSQL 16-18。
