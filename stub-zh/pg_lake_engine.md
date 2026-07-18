## 用法

来源：

- [官方上游文档](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)

`pg_lake_engine` — 用于数据湖查询的查询引擎

已复核目录快照记录的版本为 `3.4`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_extension_base`, `pg_map`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_lake_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_engine';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
