## 用法

来源：

- [官方上游文档](https://cloud.tencent.com/document/product/409/86587)
- [官方项目或服务商页面](https://cloud.tencent.com/product/postgres)

`tencentdb_failover_slot` — 腾讯云数据库专有内核扩展：把逻辑复制槽状态同步到备节点，使逻辑订阅可在主备切换后连续运行。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tencentdb_failover_slot";
SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_failover_slot';
```

这是 `Tencent Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
