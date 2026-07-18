## 用法

来源：

- [官方扩展控制文件](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck.control)
- [官方上游文档](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/README.md)

`pgsynck` — 仅解析 SQL 文本而不执行，并返回逐条语句的语法诊断信息

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgsynck";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgsynck';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
