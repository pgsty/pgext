## 用法

来源：

- [官方上游文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-encdb-extension-to-encrypt-sensitive-columns)
- [官方项目或服务商页面](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_encdb` — 阿里云 RDS 敏感列查询结果加密扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "rds_encdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_encdb';
```

这是 `Alibaba Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
