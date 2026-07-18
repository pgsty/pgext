## 用法

来源：

- [官方扩展控制文件](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/pg_bigmr.control)
- [官方上游文档](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/README.md)
- [官方 Rust 包清单](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/Cargo.toml)

`pg_bigmr` — pg_bigmr：全文搜索相关 PostgreSQL 扩展；上游说明为“基于 bigram 的文本相似度计算和索引搜索”。

已复核目录快照记录的版本为 `0.1.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_bigmr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bigmr';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
