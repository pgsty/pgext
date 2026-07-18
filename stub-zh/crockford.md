## 用法

来源：

- [官方扩展控制文件](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/crockford.control)
- [官方上游文档](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/README.markdown)

`crockford` — 提供 Crockford Base32 编码的无符号整数数据类型。

已复核目录快照记录的版本为 `0.8.34`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "crockford";
SELECT extversion
FROM pg_extension
WHERE extname = 'crockford';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
