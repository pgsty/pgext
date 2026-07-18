## 用法

来源：

- [官方上游文档](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.rds.version.html)

`rds_tools` — AWS RDS for PostgreSQL 管理函数扩展，用于引擎版本查询、密码加密审计、autovacuum 诊断和存储监控

已复核目录快照记录的版本为 `1.9`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "rds_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_tools';
```

这是 `AWS` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
