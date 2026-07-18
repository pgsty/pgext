## 用法

来源：

- [官方扩展控制文件](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/pgosql.control)
- [官方上游文档](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/README.md)

`pgosql` — 提供 PL/SQL 风格兼容性与 Oracle 类 sys 目录视图的 C 过程语言处理器。

已复核目录快照记录的版本为 `9.4`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgosql";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgosql';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
