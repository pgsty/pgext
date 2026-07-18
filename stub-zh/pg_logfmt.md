## 用法

来源：

- [官方扩展控制文件](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/pg_logfmt.control)

`pg_logfmt` — pg_logfmt：将 logfmt 字符串解析为 JSONB 和键集合。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_logfmt";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_logfmt';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
