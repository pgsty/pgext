## 用法

来源：

- [官方扩展控制文件](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/pgzan.control)

`pgzan` — pgzan 是一个概念验证型 pgrx 扩展，用于在 PostgreSQL 行级安全策略中评估 Zanzibar 风格的授权规则。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgzan";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzan';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
