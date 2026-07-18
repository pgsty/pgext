## 用法

来源：

- [官方扩展控制文件](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/pgcuid2.control)
- [官方上游文档](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/README.md)

`pgcuid2` — pgcuid2：基于 Rust cuid2 库在 PostgreSQL 中生成 CUID2 文本标识符。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgcuid2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcuid2';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
