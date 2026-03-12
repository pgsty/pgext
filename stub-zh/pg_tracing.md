


## 用法

> [pg_tracing: PostgreSQL 分布式追踪](https://github.com/DataDog/pg_tracing)

pg_tracing 生成服务端 Span 用于分布式追踪。它为规划器、执行器、计划节点、嵌套查询、触发器、并行工作进程和事务提交创建 Span。

### 通过 SQLCommenter 传播追踪上下文

将追踪上下文作为 SQL 注释传递。带有采样标志的查询将生成 Span：

```sql
/*traceparent='00-00000000000000000000000000000123-0000000000000123-01'*/ SELECT 1;
```

### 通过 GUC 传播追踪上下文

```sql
BEGIN;
SET LOCAL pg_tracing.trace_context='traceparent=''00-00000000000000000000000000000005-0000000000000005-01''';
UPDATE pgbench_accounts SET abalance=1 WHERE aid=1;
COMMIT;
```

### 独立采样

无需外部追踪上下文，随机采样查询：

```sql
SET pg_tracing.sample_rate = 1.0;  -- 追踪所有查询
SELECT 1;
```

### 查看 Span

```sql
-- 消费 Span（从缓冲区移除）
SELECT trace_id, parent_id, span_id, span_start, span_end, span_type, span_operation
FROM pg_tracing_consume_spans ORDER BY span_start;

-- 查看 Span（非破坏性）
SELECT * FROM pg_tracing_peek_spans;

-- 导出为 OTLP JSON
SELECT pg_tracing_json_spans();
```

### 统计信息

```sql
SELECT * FROM pg_tracing_info;
SELECT pg_tracing_reset();
```

### 发送 Span 到 OpenTelemetry Collector

在 `postgresql.conf` 中配置：

```
pg_tracing.otel_endpoint = http://127.0.0.1:4318/v1/traces
pg_tracing.otel_naptime = 2000
```

### 关键 GUC 参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_tracing.max_span` | 10000 | 共享内存中的最大 Span 数 |
| `pg_tracing.track` | `all` | 追踪哪些语句 |
| `pg_tracing.sample_rate` | 0 | 随机采样查询的比例 |
| `pg_tracing.otel_endpoint` | (空) | 用于 Span 导出的 OTLP HTTP 端点 |
