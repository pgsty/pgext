## 用法

来源：

- [Amazon Aurora 机器学习文档](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-ml.html)

`aws_ml` 是 Amazon Aurora PostgreSQL 中用于从 SQL 调用 Amazon Comprehend、SageMaker AI 和 Bedrock 的扩展。它属于 AWS 托管能力，并不是可移植到自建 PostgreSQL 的扩展；所有参与的服务必须在同一个 AWS 区域中可用。

### 核心流程

先使用相应的 Aurora ML 功能，将服务所需的 IAM 角色与策略附加到 Aurora 集群。然后连接写实例并安装扩展；`CASCADE` 会同时安装 `aws_commons`。

```sql
CREATE EXTENSION IF NOT EXISTS aws_ml CASCADE;

SELECT aws_bedrock.invoke_model_get_embeddings(
  model_id := 'amazon.titan-embed-text-v1',
  content_type := 'application/json',
  json_key := 'embedding',
  model_input := '{"inputText":"PostgreSQL"}'
);
```

这个 Bedrock 示例要求扩展版本为 `2.0`、账户已获得模型访问权限，并且模型的请求与响应格式和给定 JSON 相匹配。版本 `1.0` 提供 `aws_comprehend.detect_sentiment` 与 `aws_sagemaker.invoke_endpoint`；版本 `2.0` 新增 `aws_bedrock.invoke_model` 与 `aws_bedrock.invoke_model_get_embeddings`。

### 安全与运维

安装会创建 `aws_ml` 管理角色，以及 `aws_comprehend`、`aws_sagemaker` 和 `aws_bedrock` 模式。默认会撤销 PUBLIC 的执行权限；管理员应只为必要函数授予模式使用权与执行权。

可用性取决于 Aurora PostgreSQL 版本和区域。IAM、服务配额、网络调用、模型费用、载荷限制与供应商升级均是运行依赖；在事务或延迟敏感查询中使用前，应在目标 Aurora 集群上验证错误路径与延迟。
