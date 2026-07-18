## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace.control)
- [官方上游文档](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/README)

`pg_backtrace` — pg_backtrace：用于在 PostgreSQL 错误和信号处理时输出调用栈回溯的扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_backtrace";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_backtrace';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
