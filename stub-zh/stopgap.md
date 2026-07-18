## 用法

来源：

- [官方上游文档](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [官方 Rust 包清单](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/Cargo.toml)

`stopgap` — 基于 PLTS 管理 TypeScript/JavaScript 函数包版本、部署、激活与回滚的 pgrx 扩展。

已复核目录快照记录的版本为 `0.1.3`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`plts`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "stopgap";
SELECT extversion
FROM pg_extension
WHERE extname = 'stopgap';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
