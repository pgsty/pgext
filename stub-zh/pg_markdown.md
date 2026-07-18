## 用法

来源：

- [官方扩展控制文件](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/pg_markdown.control)
- [官方上游文档](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/README.md)

`pg_markdown` — PostgreSQL 的预发布 Markdown 转 HTML 扩展。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_markdown";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_markdown';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
