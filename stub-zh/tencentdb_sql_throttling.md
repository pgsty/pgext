## 用法

来源：

- [官方项目或服务商页面](https://cloud.tencent.com/product/postgres)

`tencentdb_sql_throttling` — 腾讯云数据库专有预加载扩展：按规范化 SQL 文本或 query_id 限制主实例与只读节点上的 SQL 并发。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tencentdb_sql_throttling";
SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_sql_throttling';
```

这是 `Tencent Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
