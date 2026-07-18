## 用法

来源：

- [官方扩展控制文件](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/varlena_alignment.control)
- [官方 Rust 包清单](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/Cargo.toml)

`varlena_alignment` — 基于 pgrx 的 PostgreSQL 测试扩展，用于记录 pgrx set_varsize 辅助函数的 varlena 对齐行为。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "varlena_alignment";
SELECT extversion
FROM pg_extension
WHERE extname = 'varlena_alignment';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
