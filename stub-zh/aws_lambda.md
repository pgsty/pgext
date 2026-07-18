## 用法

来源：

- [官方上游文档](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda.html)
- [官方项目或服务商页面](https://aws.amazon.com/rds/postgresql/)

`aws_lambda` — Amazon RDS for PostgreSQL 中用于从 SQL 调用 AWS Lambda 函数的扩展，依赖 aws_commons 提供辅助结构。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`aws_commons`。
整理后的兼容版本集合为 `12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "aws_lambda";
SELECT extversion
FROM pg_extension
WHERE extname = 'aws_lambda';
```

这是 `AWS` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
