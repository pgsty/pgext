## 用法

来源：

- [官方扩展控制文件](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.control)
- [官方上游文档](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/README)

`pguri` — 提供可索引的 URI 与域名类型，以及 URI 文本搜索能力。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pguri";
SELECT extversion
FROM pg_extension
WHERE extname = 'pguri';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
