## 用法

来源：

- [官方上游来源](https://docs.cloud.google.com/alloydb/docs/use-index-advisor)

`google_db_advisor` — google_db_advisor：AlloyDB 索引顾问扩展，用于分析工作负载并推荐可提升查询性能的索引。

已复核目录快照记录的版本为 `unknown`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "google_db_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'google_db_advisor';
```

这是 `Google Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
