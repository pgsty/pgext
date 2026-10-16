## 用法

来源：

- [README.md](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/README.md)
- [Cargo.toml](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/Cargo.toml)
- [src/lib.rs](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/src/lib.rs)
- [sql/pg_otel--0.1.0--0.2.0.sql](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/sql/pg_otel--0.1.0--0.2.0.sql)
- [sql/setup.sql](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/sql/setup.sql)
- [pg_otel.control](https://github.com/neteye-platform/pg_otel/blob/490e8dea7860f71a86950902fef426bc6e02edbf/pg_otel.control)

`pg_otel` 0.2.0 通过 OTLP 后台工作进程导出查询及计划节点的 span，目前属于实验性扩展。必须在服务器启动时载入库并重启，然后显式启用跟踪；最短时长默认值 -1 会关闭测量与导出。

### 核心用法

```ini
shared_preload_libraries = 'pg_otel'
pg_otel.otlp_endpoint = 'http://localhost:4317'
pg_otel.otlp_protocol = 'grpc'
pg_otel.min_duration_ms = 100
```

```sql
CREATE EXTENSION pg_otel;
SELECT pg_otel_dropped_spans();
```

### 运行边界

所检查的 Cargo 清单提供 PostgreSQL 18 和 19 目标，默认使用 19；这些构建特性不等于生产认证。库会拒绝后续 `LOAD`。`CREATE EXTENSION` 安装 SQL 接口；既有 0.1.0 数据库须在安装新版二进制并重启后执行 `ALTER EXTENSION pg_otel UPDATE`。所检查的仓库没有明确许可证。

`pg_otel.otlp_endpoint`、`pg_otel.otlp_protocol` 和超时设置用于配置采集端。`pg_otel.queue_size_kb` 在启动时预留共享内存，`pg_otel.max_plan_spans` 限制计划细节。队列满时丢弃 span，可用 `pg_otel_dropped_spans()` 检查。应使用事务级 `pg_otel.traceparent`，避免连接池复用后关联到上一个请求。

测量增加逐节点计时开销。`pg_otel.query_text` 默认输出规范化 SQL，`raw` 可能泄露敏感文本。采集端认证属于受限设置。跟踪记录可能不完整或被丢弃，调用者也能通过未采样的父上下文抑制导出，因此不能将它用作持久审计日志。
