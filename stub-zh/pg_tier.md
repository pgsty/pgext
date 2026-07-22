## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/README.md)
- [扩展 control 文件](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/pg_tier.control)
- [SQL API 与分层工作流](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/sql/pg_tier.sql)
- [Rust 实现](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/src/lib.rs)
- [Cargo 元数据](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/Cargo.toml)

`pg_tier` 通过必需的 `parquet_s3_fdw` 依赖，把本地表移动到兼容 S3 的存储，再以外部表形式在原表名下暴露数据。

### 核心流程

```sql
CREATE EXTENSION pg_tier CASCADE;
SELECT tier.set_tier_config(
  :'bucket_name', :'access_key', :'secret_key', :'region'
);
SELECT tier.enable_tiering('archive_source');
SELECT tier.execute_s3_tiering('archive_source');
```

示例使用 psql 变量，因此文档中没有硬编码真实凭据；但变量展开后仍会把值发送给 PostgreSQL，语句日志可能捕获它们。实际使用应优先采用客户端协议参数，并配置合适的日志控制。

### 安全与限制

- `tier.set_tier_config` 把 bucket、access key、secret key 和 region 以文本形式存入 `tier.server_credential`，并创建 `FOR public` 的 FDW user mapping。应限制扩展表与管理函数权限、保护备份，并使用仅有最小 bucket 权限的专用身份。
- 配置函数创建固定名称 `pg_tier_s3_srv` 的服务器。重复配置或已有同名对象可能失败，投入运行前应测试完整生命周期与恢复流程。
- 启用分层会重命名本地表、以原名称创建外部表，并在扩展目录中记录两者；执行阶段写入远端数据，而且通常会清空已重命名的本地表。
- 应验证依赖的凭据、端点和 TLS 行为。已复核构建只覆盖 PostgreSQL 14 至 16，没有 PostgreSQL 17 或更高版本的 feature。

上游明确将 0.0.5 标为 beta，并称其可能存在错误与并发问题，不适合生产。分层会改变表的存储与对象身份，并在 PostgreSQL 与 S3 之间引入部分失败边界。必须测试行数、约束、索引、权限、并发写入、回滚、备份、恢复与故障恢复；远端数据和替换外部表验证完整前，应保留独立核实的副本。
