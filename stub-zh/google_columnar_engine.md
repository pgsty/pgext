## 用法

来源：

- [官方上游来源](https://docs.cloud.google.com/alloydb/docs/columnar-engine/about?hl=en)

`google_columnar_engine` — AlloyDB 的列式存储、规划与执行引擎，用于加速 HTAP/OLAP 的扫描、连接和聚合。

已复核目录快照记录的版本为 `unknown`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "google_columnar_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'google_columnar_engine';
```

这是 `Google Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
