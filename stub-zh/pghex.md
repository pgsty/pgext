## 用法

来源：

- [官方扩展控制文件](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/pghex.control)
- [官方上游文档](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/README.md)
- [官方 Rust 包清单](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/Cargo.toml)

`pghex` — 为 PostgreSQL 添加六边形网格坐标类型、运算符及邻居、距离、范围和画线等操作的 pgrx 扩展。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pghex";
SELECT extversion
FROM pg_extension
WHERE extname = 'pghex';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
