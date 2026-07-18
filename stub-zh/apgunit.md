## 用法

来源：

- [官方上游文档](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [官方项目或服务商页面](https://aws.amazon.com/rds/aurora/)

`apgunit` — apgunit：RDS/Aurora PostgreSQL 内部扩展，AWS 文档显示仅旧版 Aurora PostgreSQL 9.6 中列出且已不再支持。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "apgunit";
SELECT extversion
FROM pg_extension
WHERE extname = 'apgunit';
```

这是 `AWS` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
