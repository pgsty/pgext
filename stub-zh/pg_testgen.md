## 用法

来源：

- [官方扩展控制文件](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen.control)
- [官方上游文档](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/README.md)

`pg_testgen` — 为 PostgreSQL 开发生成随机整数、字符串及批量测试数据行

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_testgen";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_testgen';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
