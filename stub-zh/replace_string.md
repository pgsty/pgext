## 用法

来源：

- [官方扩展控制文件](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/replace_string.control)
- [官方上游文档](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/README.md)
- [官方 Rust 包清单](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/Cargo.toml)

`replace_string` — 基于 Rust/pgrx 的字面字符串替换函数

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "replace_string";
SELECT extversion
FROM pg_extension
WHERE extname = 'replace_string';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
