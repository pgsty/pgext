## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/tsvector2.control)
- [官方上游文档](https://pgxn.org/dist/tsvector2/1.0.0/)
- [官方 PGXN 分发页](https://pgxn.org/dist/tsvector2/)

`tsvector2` — PostgreSQL tsvector 的压缩替代类型，移除原始 1 MB 限制，并提供兼容的检索函数、GIN、GiST 及可选 RUM 索引。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tsvector2";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsvector2';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
