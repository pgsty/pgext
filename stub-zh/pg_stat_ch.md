
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pg_stat_ch;
> SELECT pg_stat_ch_version();
> SELECT * FROM pg_stat_ch_stats();
> ```
>
> 来源：[README](https://github.com/ClickHouse/pg_stat_ch)，[博客文章](https://clickhouse.com/blog/pg_stat_ch-postgres-extension-stats-to-clickhouse)

`pg_stat_ch` 会在 PostgreSQL 中捕获每条查询的执行遥测信息，并实时把原始事件导出到 ClickHouse。上游项目将其与 `pg_stat_statements` 对比：后者在 PostgreSQL 内聚合统计，而它把原始事件发送到 ClickHouse 进行下游分析。

## 架构

README 将数据流描述为：

```text
PostgreSQL hooks -> shared memory queue -> background worker -> ClickHouse
```

上游明确强调的设计目标包括：

- 查询路径上不发生网络 I/O
- 通过固定大小环形缓冲区限制内存
- 导出原始事件，而不是在本地聚合
- 当队列溢出或 ClickHouse 不可用时优雅降级

## 配置

扩展必须预加载，并配置 ClickHouse 连接参数：

```ini
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

在 PostgreSQL 重启并完成 ClickHouse schema 初始化后：

```sql
CREATE EXTENSION pg_stat_ch;
```

## SQL API

README 记录了以下 SQL 函数：

- `pg_stat_ch_version()`
- `pg_stat_ch_stats()`
- `pg_stat_ch_reset()`

`pg_stat_ch_stats()` 会暴露队列和导出器计数器，便于确认事件是否已被捕获并刷出。

## 捕获内容

当前 README 说明支持捕获：

- 查询耗时与返回行数
- 缓冲区使用量与 WAL 使用量
- CPU 时间
- DML、DDL 和通用语句
- SQLSTATE 与错误级别
- PostgreSQL 15+ 的 JIT 信息
- PostgreSQL 18+ 的并行 worker 统计
- 诸如应用名和客户端 IP 等客户端上下文信息

项目目前声明对 PostgreSQL 16、17 和 18 提供完整支持。
