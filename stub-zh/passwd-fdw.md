## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/passwd-fdw/passwd-fdw-0.4.0/passwd-fdw.control)
- [官方上游文档](https://pgxn.org/dist/passwd-fdw/0.4.0/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/passwd-fdw/)

`passwd-fdw` — passwd-fdw：外部数据访问类 PostgreSQL 扩展；用于以 FDW 方式访问 Unix/Linux 用户与用户组数据库。

已复核目录快照记录的版本为 `0.4.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "passwd-fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'passwd-fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
