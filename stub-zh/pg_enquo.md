## 用法

来源：

- [官方项目或服务商页面](https://enquo.org/)

`pg_enquo` — 为 PostgreSQL 提供可加密查询的数据类型与操作

已复核目录快照记录的版本为 `0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_enquo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_enquo';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
