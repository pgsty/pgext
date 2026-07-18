## 用法

来源：

- [官方扩展控制文件](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/json_validator_ext.control)
- [官方上游文档](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/README.md)
- [官方 Rust 包清单](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/Cargo.toml)

`json_validator_ext` — 基于 Rust/pgrx 的 JSONB JSON Schema 校验扩展，提供 @@ 操作符和错误说明函数。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "json_validator_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'json_validator_ext';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
