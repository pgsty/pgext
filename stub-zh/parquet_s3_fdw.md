## 用法

来源：

- [官方扩展控制文件](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/parquet_s3_fdw.control)
- [官方上游文档](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/README.md)

`parquet_s3_fdw` — 用于从 PostgreSQL 访问本地文件系统和 Amazon S3 上 Parquet 文件的外部数据包装器。

已复核目录快照记录的版本为 `0.3`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "parquet_s3_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'parquet_s3_fdw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
