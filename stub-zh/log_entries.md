## 用法

来源：

- [官方扩展控制文件](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries.control)

`log_entries` — 通过行级触发器记录谁在何时修改了数据行。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "log_entries";
SELECT extversion
FROM pg_extension
WHERE extname = 'log_entries';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
