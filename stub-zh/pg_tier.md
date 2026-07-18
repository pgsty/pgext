## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/README.md)
- [扩展 control 文件](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/pg_tier.control)
- [Rust 实现](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/src/lib.rs)
- [Cargo 元数据](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/Cargo.toml)

`pg_tier` 通过必需的 `parquet_s3_fdw` 依赖，把本地表移动到兼容 S3 的存储，再以外部表形式在原表名下暴露数据。

```sql
CREATE EXTENSION pg_tier CASCADE;
SELECT tier.set_tier_config(
  'archive-bucket', 'ACCESS_KEY', 'SECRET_KEY', 'REGION'
);
SELECT tier.enable_tiering('archive_source');
SELECT tier.execute_s3_tiering('archive_source');
```

上面的凭据字符串只是占位符。不要把真实云密钥放入查询文本、日志、shell 历史或可广泛读取的目录中。应使用仅有最小权限的专用 bucket 身份，并验证依赖的凭据与 TLS 行为。

上游明确将 0.0.5 标为 beta，并称其可能存在错误与并发问题，不适合生产。分层会改变表的存储与对象身份，并在 PostgreSQL 与 S3 之间引入部分失败边界。必须测试行数、约束、索引、权限、并发写入、回滚、备份、恢复与故障恢复；远端数据和替换外部表验证完整前，应保留独立核实的副本。
