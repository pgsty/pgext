## 用法

来源：

- [官方扩展控制文件](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/rule_engine_postgre_extensions.control)
- [官方 Rust 包清单](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/Cargo.toml)

`rule_engine_postgre_extensions` — 基于 Rust/pgrx 的规则引擎扩展，支持 GRL、RETE 与正向/反向链式推理、规则版本库、触发器及可选外部集成。

已复核目录快照记录的版本为 `2.0.0`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`pgcrypto`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "rule_engine_postgre_extensions";
SELECT extversion
FROM pg_extension
WHERE extname = 'rule_engine_postgre_extensions';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
