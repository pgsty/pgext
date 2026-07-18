## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/pg_tsparser.control)
- [官方上游文档](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/README.md)

`pg_tsparser` — pg_tsparser：全文搜索相关 PostgreSQL 扩展；上游说明为“PostgreSQL 全文搜索解析器”。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_tsparser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tsparser';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
