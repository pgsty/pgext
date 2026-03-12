


## 用法

> [omni_aws: Amazon Web Services API（S3）](https://docs.omnigres.org/omni_aws/s3/)

`omni_aws` 扩展为 PostgreSQL 提供 S3 兼容的 API 函数。

### S3 API 函数

**`s3_create_bucket(bucket, region)`** -- 创建 S3 存储桶。

**`s3_list_objects_v2(bucket, path, continuation_token, delimiter, encoding_type, fetch_owner, max_keys, start_after, region)`** -- 列出存储桶中的对象，支持分页。

**`s3_put_object(bucket, path, payload, content_type, region)`** -- 上传对象。默认内容类型为 `application/octet-stream`。

### 执行

**单个请求：**

```sql
SELECT * FROM omni_aws.aws_execute(
    access_key_id => 'AKID',
    secret_access_key => 'SECRET',
    request => omni_aws.s3_put_object('my-bucket', '/path/to/file', 'content'::bytea)
);
```

**批量请求：** 将 `requests` 作为数组传入以实现高效批量操作。返回带 `error` 列的表。

### 预签名 URL

```sql
SELECT omni_aws.s3_presigned_url(
    bucket => 'my-bucket',
    path => '/my-file',
    access_key_id => 'AKID',
    secret_access_key => 'SECRET',
    expires => 604800,  -- 7 天（默认）
    region => 'us-east-1'
);
```

### 自定义端点

- **DigitalOcean Spaces：** `digitalocean_s3_endpoint()`
- **通用端点：** `s3_endpoint(url)`，支持 `${region}` 和 `${bucket}` 变量，适用于 MinIO 等兼容服务
