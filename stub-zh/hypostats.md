## 用法

来源：

- [官方 Rust 包清单](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/Cargo.toml)

`hypostats` — 将 pg_statistic 统计信息行导出/导入为 JSON，用于统计信息实验。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "hypostats";
SELECT extversion
FROM pg_extension
WHERE extname = 'hypostats';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
