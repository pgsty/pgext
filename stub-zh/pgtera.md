## 用法

来源：

- [官方扩展控制文件](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/pgtera.control)
- [官方上游文档](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/README.md)
- [官方 Rust 包清单](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/Cargo.toml)

`pgtera` — 在 PostgreSQL 内使用 Tera 渲染 HTML 模板。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgtera";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtera';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
