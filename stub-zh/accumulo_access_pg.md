## 用法

来源：

- [官方 Rust 包清单](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/Cargo.toml)
- [官方上游 README](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/README.md)

`accumulo_access_pg` — 用于在 PostgreSQL 中解析和评估 Accumulo 访问表达式，可用于行级安全策略。

已复核目录快照记录的版本为 `0.1.5`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "accumulo_access_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'accumulo_access_pg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
