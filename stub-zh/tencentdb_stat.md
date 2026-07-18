## 用法

来源：

- [官方上游文档](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_stat` — 用于采集 PostgreSQL QPS 统计信息的 TencentDB 厂商扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tencentdb_stat";
SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_stat';
```

这是 `Tencent Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
