## 用法

来源：

- [官方扩展控制文件](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/pg_splade.control)
- [官方上游文档](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/README.md)
- [官方 Rust 包清单](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/Cargo.toml)

`pg_splade` — 在 PostgreSQL 内生成 SPLADE 稀疏向量嵌入

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_splade";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_splade';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
