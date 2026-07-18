## 用法

来源：

- [官方上游文档](https://pgxn.org/dist/multicdr_fdw/1.2.2/doc/multicdr_fdw.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/multicdr_fdw/)
- [官方上游 README](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/README.md)

`multicdr_fdw` — 用于将通话详单（CDR）文件流作为外部 SQL 表读取的外部数据包装器。

已复核目录快照记录的版本为 `1.2.2`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "multicdr_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'multicdr_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
