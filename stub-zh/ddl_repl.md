## 用法

来源：

- [官方扩展控制文件](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.control)
- [官方上游文档](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/README.md)

`ddl_repl` — 通过实用命令钩子在多个 PostgreSQL 实例间复制 DDL 与 DCL 命令

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "ddl_repl";
SELECT extversion
FROM pg_extension
WHERE extname = 'ddl_repl';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
