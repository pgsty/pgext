## 用法

来源：

- [官方扩展控制文件](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/constr_name_unif.control)
- [官方上游文档](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/doc/constr_name_unif.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/constr_name_unif/)

`constr_name_unif` — 统一 PostgreSQL 基表中各类约束的名称。

已复核目录快照记录的版本为 `1.0.1`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "constr_name_unif";
SELECT extversion
FROM pg_extension
WHERE extname = 'constr_name_unif';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
