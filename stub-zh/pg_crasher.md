## 用法

来源：

- [官方扩展控制文件](https://github.com/davidcrawford/pg_crasher/blob/master/pg_crasher.control)
- [官方上游文档](https://github.com/davidcrawford/pg_crasher/blob/master/README)

`pg_crasher` — 提供一个故意让 PostgreSQL 后端崩溃的测试函数，用于验证客户端错误处理。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_crasher";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_crasher';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
