## 用法

来源：

- [官方扩展控制文件](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime.control)
- [官方上游文档](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/README.md)

`pg_sysdatetime` — 提供 SQL Server 风格的高精度系统时间函数，主要用于旧版 Windows 上的 PostgreSQL。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_sysdatetime";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sysdatetime';
```

该上游项目与 `2ndQuadrant` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
