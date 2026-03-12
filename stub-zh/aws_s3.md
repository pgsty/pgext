

## 用法

> [aws_s3: 从 S3 导入/导出数据的 PostgreSQL 扩展](https://github.com/chimpler/postgres-aws-s3)

### 设置凭证

通过 PostgreSQL 会话变量配置 AWS 凭证：

```sql
SET aws_s3.access_key_id TO 'your_access_key';
SET aws_s3.secret_key TO 'your_secret_key';
SET aws_s3.session_token TO 'optional_session_token';  -- 使用临时凭证时
```

本地开发使用 LocalStack：

```sql
SET aws_s3.endpoint_url TO 'http://localhost:4566';
```

### 从 S3 导入数据

```sql
CREATE EXTENSION aws_s3;

CREATE TABLE animals (
  name text,
  age int
);

SELECT aws_s3.table_import_from_s3(
  'animals',
  '',
  '(FORMAT CSV, DELIMITER '','', HEADER true)',
  'my-bucket',
  'animals.csv',
  'us-east-1'
);

SELECT * FROM animals;
```

**参数：** 表名、列列表（空字符串表示全部）、COPY 选项、S3 桶名、S3 键名、AWS 区域。

### 导出数据到 S3

```sql
SELECT * FROM aws_s3.query_export_to_s3(
  'SELECT * FROM animals',
  'my-bucket',
  'export/animals.csv',
  'us-east-1',
  options := 'FORMAT CSV, HEADER true'
);
```

**参数：** SQL 查询、S3 桶名、S3 键名、AWS 区域、COPY 选项。

### 特性

- 带有 `Content-Encoding=gzip` 元数据的文件在导入时自动解压
- 凭证可以作为函数参数直接传递或通过会话变量设置
- 使用与 PostgreSQL 相同的 COPY 格式选项（CSV、TEXT、BINARY，以及所有相关设置如 DELIMITER、HEADER、NULL 等）
