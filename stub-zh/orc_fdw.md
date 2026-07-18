## 用法

来源：

- [官方扩展控制文件](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw.control)
- [官方上游文档](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/README.md)

`orc_fdw` — orc_fdw：用于访问或集成 ORC 格式文件的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "orc_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'orc_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
