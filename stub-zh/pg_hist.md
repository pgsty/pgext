## 用法

来源：

- [官方扩展控制文件](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/pg_hist.control)
- [官方上游文档](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/README.md)

`pg_hist` — 在 PostgreSQL 中快速计算一维或多维的加权与非加权直方图

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_hist";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_hist';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
