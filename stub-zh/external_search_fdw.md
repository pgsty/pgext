## 用法

来源：

- [官方上游来源](https://docs.cloud.google.com/alloydb/docs/elastic-search?hl=en)

`external_search_fdw` — AlloyDB 扩展，用于通过外部数据包装器从 PostgreSQL 只读查询 Elasticsearch 数据。

已复核目录快照记录的版本为 `unknown`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "external_search_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'external_search_fdw';
```

这是 `Google Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
