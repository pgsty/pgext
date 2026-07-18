## 用法

来源：

- [官方扩展控制文件](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji.control)
- [官方上游文档](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/README.md)

`emoji` — emoji：纯 SQL PostgreSQL 扩展，可将 bytea/text 编码为 emoji，或从 emoji 解码。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "emoji";
SELECT extversion
FROM pg_extension
WHERE extname = 'emoji';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
