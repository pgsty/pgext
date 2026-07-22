## 用法

来源：

- [AWS 配置指南](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda.html)
- [AWS 函数与参数参考](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda-functions.html)
- [AWS 调用示例](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda-examples.html)

`aws_lambda` 是 Amazon RDS for PostgreSQL 中用于从 SQL 调用 AWS Lambda 函数的扩展。它是 AWS 托管的供应商专有组件：SQL 接口生效前，必须先在 RDS 实例上配置网络和 IAM。

### 核心流程

为 Lambda 配置出站 HTTPS 访问，附加仅对允许函数授予 `lambda:InvokeFunction` 的 RDS IAM 角色，并用 `Lambda` 功能将角色关联到数据库实例。然后以拥有 `rds_superuser` 权限的用户连接，并连同 `aws_commons` 依赖安装扩展：

```sql
CREATE EXTENSION IF NOT EXISTS aws_lambda CASCADE;
```

构造函数标识并执行同步调用：

```sql
SELECT *
FROM aws_lambda.invoke(
  aws_commons.create_lambda_function_arn(
    'arn:aws:lambda:us-west-1:444455556666:function:simple'
  ),
  '{"body":"Hello from PostgreSQL"}'::json,
  'RequestResponse'
);
```

结果包含 `status_code`、`payload`、`executed_version` 和 `log_result`。只向非超级用户授予所需的 SQL 权限：

```sql
GRANT USAGE ON SCHEMA aws_lambda TO app_user;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA aws_lambda TO app_user;
```

### 主要接口

- `aws_commons.create_lambda_function_arn(function_name, region)` 返回 `aws_lambda.invoke` 接受的复合函数标识。
- `aws_lambda.invoke` 提供 `json` 和 `jsonb` 重载，函数名既可使用文本名称/ARN，也可使用 `aws_commons` 复合值。
- `invocation_type` 区分大小写：`RequestResponse` 等待并返回载荷，`Event` 将调用异步入队，`DryRun` 只检查访问权限而不执行函数。
- `log_type = 'Tail'` 以 Base64 返回最后 4 KB 执行日志；默认值为 `None`。
- `context` 传递客户端上下文，`qualifier` 选择 Lambda 版本或别名。

### 配置与注意事项

`aws_lambda.connect_timeout_ms` 和 `aws_lambda.request_timeout_ms` 是动态超时参数。`aws_lambda.endpoint_override` 是静态参数，修改后需要重启数据库。

AWS 当前记录的支持范围为 PostgreSQL 12.6+、13.2+、14.1+，以及所有 RDS PostgreSQL 15–18 版本。可用性取决于具体 RDS 引擎构建，而不是可移植的社区软件包。

SQL 授权不能替代 IAM 授权：数据库实例角色仍决定可运行哪些 Lambda 函数。私有实例需要 NAT 或 Lambda VPC 端点；后者还要求设置 `rds.custom_dns_resolution = 1` 并重启。载荷、返回日志、网络延迟、Lambda 重试和异步副作用都属于外部服务行为，而不是 PostgreSQL 事务工作。
