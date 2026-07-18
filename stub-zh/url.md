## 用法

来源：

- [官方扩展控制文件](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url.control)
- [官方上游文档](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/README.md)

`url` — 提供用于 URL 解析、编码解码以及查询字符串键值处理的 SQL 和 PL/pgSQL 类型与辅助函数。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "url";
SELECT extversion
FROM pg_extension
WHERE extname = 'url';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
