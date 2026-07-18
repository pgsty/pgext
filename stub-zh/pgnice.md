## 用法

来源：

- [官方扩展控制文件](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/pgnice.control)
- [官方上游文档](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/README.md)
- [官方 Rust 包清单](https://github.com/mselee/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/Cargo.toml)

`pgnice` — pgnice 是一个基于 pgrx 的 PostgreSQL 扩展，用于调整当前后端进程的 nice、I/O 优先级和资源限制。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgnice";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgnice';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
