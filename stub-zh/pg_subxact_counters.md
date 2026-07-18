## 用法

来源：

- [官方扩展控制文件](https://github.com/bdrouvot/pg_subxact_counters/blob/e24d6effb4691fa2670f926e0eb0c56106392b73/c/pg_subxact_counters.control)

`pg_subxact_counters` — pg_subxact_counters：提供 PostgreSQL 子事务活动的全局计数器。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_subxact_counters";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_subxact_counters';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
