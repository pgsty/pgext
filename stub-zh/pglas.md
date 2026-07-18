## 用法

来源：

- [官方扩展控制文件](https://github.com/shestero/pglas/blob/efc9f4e89585ac773550c2093782d65d1049a466/pglas.control)

`pglas` — 通过 PostgreSQL 函数读取 LAS 2.0 测井文件。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pglas";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglas';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
