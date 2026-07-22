## 用法

来源：

- [官方 README](https://github.com/umitanuki/s3_fdw/blob/ee5b25a10d0d760423571fbe3895345cb0644df5/doc/s3_fdw.md)
- [官方扩展 SQL](https://github.com/umitanuki/s3_fdw/blob/ee5b25a10d0d760423571fbe3895345cb0644df5/s3_fdw.sql)
- [官方 PGXN 分发页面](https://pgxn.org/dist/s3_fdw/)

`s3_fdw` 是一个实验性的只读外部数据包装器，它使用 PostgreSQL `COPY` 解析选项，将一个 Amazon S3 对象流式映射为一张外部表。复核的版本是 2011 年发布的 `0.1.0`；其传输与认证设计已经过时，应将它视为历史代码，而不是生产级 S3 集成方案。

### 核心流程

创建扩展、注册外部数据包装器，并在用户映射中保存 S3 凭据：

```sql
CREATE EXTENSION s3_fdw;

CREATE SERVER amazon_s3
  FOREIGN DATA WRAPPER s3_fdw;

CREATE USER MAPPING FOR CURRENT_USER
  SERVER amazon_s3
  OPTIONS (
    accesskey 'example-access-key',
    secretkey 'example-secret-key'
  );
```

将一个对象映射为外部表。表字段定义必须与对象内容及所选的 `COPY` 选项一致：

```sql
CREATE FOREIGN TABLE s3_events (
  event_time timestamptz,
  payload text
)
SERVER amazon_s3
OPTIONS (
  hostname 's3-ap-northeast-1.amazonaws.com',
  bucketname 'example-bucket',
  filename 'events.tsv',
  delimiter E'\t'
);

SELECT * FROM s3_events;
```

### 选项与行为

- `hostname`、`bucketname` 与 `filename` 分别标识 S3 端点、存储桶和对象；这三个外部表选项均为必填项。
- `accesskey` 与 `secretkey` 是用户映射选项，供旧式 S3 请求签名器使用。
- 其他外部表选项会传给 PostgreSQL 的 `COPY` 输入解析器，可使用 `delimiter` 等设置描述对象格式。
- 该包装器只实现扫描回调，不支持 `INSERT`、`UPDATE` 或 `DELETE`，也没有服务器级选项。

### 运维注意事项

实现会构造 `http://` URL，并使用旧式 AWS 访问密钥签名；它没有现代 HTTPS、区域、会话令牌或 IAM 角色处理能力。密钥保存在 PostgreSQL 目录中可见的用户映射选项里。它还依赖派生的辅助进程与 Unix FIFO，上游文档也明确称该版本不稳定且行为不可预测。不要通过这段代码暴露凭据或敏感数据。现代系统应改用仍在维护、支持当前 TLS 与 AWS 认证机制的 S3 FDW 或数据导入流水线。
