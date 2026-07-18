## 用法

来源：

- [官方扩展控制文件](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql.control)
- [官方上游文档](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/README.md)

`dbms_sql` — 实现 Oracle DBMS_SQL 包 API 的子集，帮助将 Oracle 应用迁移到 PostgreSQL。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "dbms_sql";
SELECT extversion
FROM pg_extension
WHERE extname = 'dbms_sql';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
