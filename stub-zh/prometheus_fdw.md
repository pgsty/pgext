## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/prometheus_fdw/)
- [官方项目或服务商页面](https://tembo.io/blog/monitoring-with-prometheus-fdw)

`prometheus_fdw` — prometheus_fdw 是一个 Rust/pgrx 外部数据包装器，可通过 PostgreSQL 外部表查询 Prometheus 指标。

已复核目录快照记录的版本为 `0.2.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "prometheus_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'prometheus_fdw';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
