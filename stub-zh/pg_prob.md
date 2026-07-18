## 用法

来源：

- [官方扩展控制文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/pg_prob.control)
- [官方 Rust 包清单](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/Cargo.toml)

`pg_prob` — 在 PostgreSQL 中提供概率分布数据类型、概率运算与 Monte Carlo 模拟。

已复核目录快照记录的版本为 `0.2.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_prob";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_prob';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
