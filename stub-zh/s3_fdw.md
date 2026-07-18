## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/s3_fdw/)

`s3_fdw` — 通过 PostgreSQL COPY 格式读取 Amazon S3 文件的外部数据包装器。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "s3_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 's3_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
