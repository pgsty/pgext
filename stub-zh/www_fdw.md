## 用法

来源：

- [官方扩展控制文件](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/www_fdw.control)
- [官方上游文档](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/www_fdw/)

`www_fdw` — 从 PostgreSQL 访问 Web 服务的外部数据包装器。

已复核目录快照记录的版本为 `0.1.9`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "www_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'www_fdw';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
