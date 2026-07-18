## 用法

来源：

- [官方扩展控制文件](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/pgc_fdw.control)

`pgc_fdw` — 通过派生自 postgres_fdw、由 FoundationDB 提供查询缓存的外部数据包装器访问远程 PostgreSQL 表

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgc_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgc_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
