## 用法

来源：

- [官方来源](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/README.md)
- [官方来源](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/pg_otel_exporter.control)
- [官方来源](https://github.com/its-the-shrimp/pg_otel_exporter/releases/tag/v0.0.3)
- [官方来源](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/LICENSE)
- [官方来源](https://github.com/its-the-shrimp/pg_otel_exporter/blob/v0.0.3/Cargo.toml)

`pg_otel_exporter` 0.0.3 通过 OTLP/HTTP 导出语句追踪、服务端日志与采样的 PostgreSQL 指标。发布制品和 Cargo feature 支持 PostgreSQL 14–19；README 中更宽泛的 13–19 表述没有对应的构建依据。

### 启用遥测

将库加入预加载列表并重启，以启用指标工作进程及辅助进程日志采集。以超级用户在每个数据库创建扩展。启用导出前，应设置由自己控制的采集端点。

```ini
shared_preload_libraries = 'pg_otel_exporter'
pg_otel_exporter.endpoint = 'http://127.0.0.1:4318'
```

```sql
CREATE EXTENSION pg_otel_exporter;
BEGIN;
CALL traceparent('00-0af7651916cd43dd8448eb8c8defd531-00f067aa0ba902b7-01');
SELECT 1;
COMMIT;
```

### 信号与配置

`traceparent(text)` 将当前事务中的后续语句关联到 W3C trace；在事务外调用不生效。携带相同上下文的 SQL 注释优先。提交、回滚或准备事务后清除上下文。执行器和 utility hook 生成 span 并反映嵌套语句关系，事务控制语句不生成 span。

`pg_otel_exporter.traces`、`pg_otel_exporter.metrics` 与 `pg_otel_exporter.logs` 选择信号；`pg_otel_exporter.sdk_disabled` 禁用导出。`pg_otel_exporter.traces_sampler_arg`、`pg_otel_exporter.metric_export_interval`、`pg_otel_exporter.log_level` 与 `pg_otel_exporter.bsp_max_export_batch_size` 控制采样和批处理。文档说明这些设置可由用户调整，应相应管理端点和角色策略。

### 数据边界

Span 包含完整 SQL 文本、数据库名与角色名，日志记录包含服务端消息。这些数据可能将查询字面量及敏感内容发送给采集端。采样遥测不是已提交的审计日志。该项目与目录中另一个 `pg_otel` 扩展不同。
