## 用法

来源：

- [官方扩展控制文件](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/pg_web.control)
- [官方上游文档](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/README.md)

`pg_web` — 通过 PostgreSQL 后台工作进程运行内嵌 HTTP 服务器

已复核目录快照记录的版本为 `0.1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_web";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_web';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
