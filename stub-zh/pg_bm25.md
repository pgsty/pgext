## 用法

来源：

- [官方扩展控制文件](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/pg_bm25.control)
- [官方上游文档](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/README.md)
- [官方 Rust 包清单](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/Cargo.toml)

`pg_bm25` — pg_bm25：全文搜索相关 PostgreSQL 扩展；上游说明为“使用 BM25 为 PostgreSQL 提供全文搜索”。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_bm25";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bm25';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
