## 用法

来源：

- [官方扩展控制文件](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/pg_zen_engine.control)

`pg_zen_engine` — 使用 Zen Engine 在 PostgreSQL 中对 JSONB 数据执行 JSON Decision Model 决策图

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_zen_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_zen_engine';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
