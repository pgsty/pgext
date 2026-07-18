## 用法

来源：

- [官方上游文档](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/fuzzywuzzy/)

`fuzzywuzzy` — fuzzywuzzy：全文搜索相关 PostgreSQL 扩展；上游说明为“为 PostgreSQL 提供 fuzzywuzzy 相似度评分”。

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`fuzzystrmatch`。

```sql
CREATE EXTENSION "fuzzywuzzy";
SELECT extversion
FROM pg_extension
WHERE extname = 'fuzzywuzzy';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
