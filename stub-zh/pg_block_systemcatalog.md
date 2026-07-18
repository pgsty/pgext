## 用法

来源：

- [官方扩展控制文件](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/pg_block_systemcatalog.control)

`pg_block_systemcatalog` — pg_block_systemcatalog：安全类 PostgreSQL 扩展；上游说明为“阻止引用系统目录的 PostgreSQL 扩展”。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_block_systemcatalog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_block_systemcatalog';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
