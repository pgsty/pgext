## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/oai_fdw/oai_fdw-1.13.0/oai_fdw.control)
- [官方上游文档](https://pgxn.org/dist/oai_fdw/1.13.0/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/oai_fdw/)

`oai_fdw` — 用于访问 OAI-PMH 元数据仓库的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.13`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "oai_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oai_fdw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
