## 用法

来源：

- [官方上游文档](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-ml.html)

`aws_ml` — aws_ml：Amazon Aurora PostgreSQL 的机器学习扩展，可从 SQL 调用 Amazon Comprehend、SageMaker AI 和 Bedrock 服务。

已复核目录快照记录的版本为 `2.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`aws_commons`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "aws_ml";
SELECT extversion
FROM pg_extension
WHERE extname = 'aws_ml';
```

这是 `AWS` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
