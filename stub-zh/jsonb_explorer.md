## 用法

来源：

- [官方扩展控制文件](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer.control)
- [官方上游 README](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/README.md)

`jsonb_explorer` — 输出 jsonb 树结构和路径的 C 扩展函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "jsonb_explorer";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_explorer';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
