## 用法

来源：

- [EDB OTEL 官方文档](https://www.enterprisedb.com/docs/pg_extensions/otel/)
- [官方安装指南](https://www.enterprisedb.com/docs/pg_extensions/otel/installing/)
- [官方 SQL API 指南](https://www.enterprisedb.com/docs/pg_extensions/otel/using/)

`edb_otel` 把 PostgreSQL metric 与 trace 导出到 OpenTelemetry endpoint，并提供其他扩展可调用的集中 telemetry routing API。

### 启用

EDB OTEL 2.0.0 通过 EDB repository 分发。配置 collector endpoint、预加载库、重启并创建扩展：

```ini
shared_preload_libraries = 'edb_otel'
edb_otel.metrics_endpoint = 'http://otel-collector:4318'
edb_otel.traces_endpoint = 'http://otel-collector:4318'
edb_otel.enable_tracing = true
```

```sql
CREATE EXTENSION edb_otel;
```

2.0 与 1.0 API 不兼容；依赖扩展代码与 SQL 调用方必须一起更新。

### 上报 Metric 与 Trace

`edb_otel.report_metric` 发送 counter、gauge、up/down counter 或 histogram value，并可附加 JSON label。Tracing 使用显式 span lifecycle。

```sql
SELECT edb_otel.report_metric(
  'orders', 'processed', 1, 1::bigint, '{"region":"us-east"}'
);

SELECT edb_otel.tracing_init();
SELECT edb_otel.start_span('checkout', 'validate-order');
```

使用返回的 span key 调用 `edb_otel.tracing_set_attribute`、`edb_otel.tracing_end_span` 与 `edb_otel.tracing_get_http_trace_context`。

### 运维边界

Metric 通过 shared-memory queue 传递；`edb_otel.metrics_queue_size` 耗尽后，新 metric 会被丢弃，直到 worker 完成排空。超过 1,024 字节的 label 或无效 JSON 也会被丢弃并写入日志。Trace export 在每个 backend 内使用同步 `SimpleSpanProcessor`，因此 collector 延迟或失败会影响 span 完成，已经结束但导出失败的 span 不会重试。

