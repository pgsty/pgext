## 用法

来源：

- [官方上游 README](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/README.md)
- [官方扩展控制文件 (pg_otel_tracer.control)](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/pg_otel_tracer.control)
- [官方扩展 SQL (pg_otel_tracer--0.1.0.sql)](https://github.com/mstryoda/pgtrace/blob/77c2efdef80c0c50627589a8ad05f1058cef1ac5/sql/pg_otel_tracer--0.1.0.sql)

`pg_otel_tracer` — OpenTelemetry 跟踪扩展用于 PostgreSQL。从 SQL 注释中提取 W3C traceparent，并通过 OTLP/HTTP 导出查询生命周期跨度。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_otel_tracer;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
