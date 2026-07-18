## 用法

来源：

- [官方扩展控制文件](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.control)

`pg_metrics` — 在共享内存中收集 PostgreSQL 后端的命名计数器，并通过 SQL 指标函数暴露。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_metrics";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_metrics';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
