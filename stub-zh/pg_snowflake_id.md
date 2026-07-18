## 用法

来源：

- [官方扩展控制文件](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/pg_snowflake_id.control)
- [官方上游文档](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/README.md)

`pg_snowflake_id` — pg_snowflake_id：基于 pgrx/Rust 的 Snowflake 风格 64 位 ID 生成函数扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_snowflake_id";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_snowflake_id';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
