## 用法

来源：

- [Amazon RDS 官方 S3 函数参考](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.Reference.html#USER_PostgreSQL.S3Import.create_s3_uri)
- [Amazon RDS 官方 S3 导入指南](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)

`aws_commons` 是 Amazon RDS for PostgreSQL 的辅助扩展，用于构造供 `aws_s3` 使用的 S3 URI 与凭据复合值。它只由服务商提供：这些对象来自 RDS，而不是社区 PostgreSQL 软件包；版本 `1.2` 是否可用取决于 RDS 引擎版本。

### 核心流程

使用 RDS 特权账户启用服务商扩展，然后构造 S3 对象描述：

```sql
CREATE EXTENSION aws_commons;
CREATE EXTENSION aws_s3;

SELECT aws_commons.create_s3_uri(
  'example-bucket',
  'imports/events.csv',
  'us-east-1'
);
```

将该值传给 `aws_s3` 导入调用；导入选项使用 PostgreSQL `COPY` 语法：

```sql
SELECT aws_s3.table_import_from_s3(
  'public.events',
  '',
  '(format csv, header true)',
  aws_commons.create_s3_uri(
    'example-bucket',
    'imports/events.csv',
    'us-east-1'
  )
);
```

### 重要对象

- `aws_commons.create_s3_uri(bucket, file_path, region)` 返回标识单个 S3 对象的 `aws_commons._s3_uri_1` 值。
- `aws_commons.create_aws_credentials(access_key, secret_key, session_token)` 返回 `aws_commons._aws_credentials_1` 值，供接收凭据的 `aws_s3` 重载使用。

### 运维说明

应优先为 RDS 实例附加 IAM 角色，避免在 SQL 中出现访问密钥。如果必须显式传递凭据，应使用临时凭据、限制 SQL 可见性与日志记录，并且不要持久化返回的复合值。存储桶策略、IAM 权限、网络连通性、对象格式与目标表都是独立的前置条件。该扩展无法安装在自管 PostgreSQL 上，函数也可能随服务商版本变化，因此应查阅目标 RDS 大版本和小版本对应的文档。
