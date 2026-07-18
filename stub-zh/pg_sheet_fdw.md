## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_sheet_fdw/)

`pg_sheet_fdw` — 将本地 Excel XLSX 工作表作为 PostgreSQL 外部表读取

已复核目录快照记录的版本为 `0.1.2`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_sheet_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sheet_fdw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
