## 用法

来源：

- [官方扩展控制文件](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice.control)
- [官方上游文档](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/README.md)

`pg_nice` — 使用 Linux ionice 降低当前 PostgreSQL 后端进程的 I/O 调度优先级

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_nice";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_nice';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
