## 用法

来源：

- [官方扩展控制文件](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/pg_plsql_graphs.control)
- [官方上游文档](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/README.md)

`pg_plsql_graphs` — 捕获 PL/pgSQL 控制流与依赖图并导出为 DOT

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_plsql_graphs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_plsql_graphs';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
