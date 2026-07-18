## 用法

来源：

- [官方扩展控制文件](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw.control)
- [官方上游文档](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/README)

`sybase_fdw` — 用于从 PostgreSQL 访问 Sybase 数据库的 C 外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "sybase_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'sybase_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
