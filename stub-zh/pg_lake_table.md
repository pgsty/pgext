## 用法

来源：

- [官方上游 README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)

`pg_lake_table` — 数据湖表和 Iceberg 表

已复核目录快照记录的版本为 `3.4`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`btree_gist`, `pg_lake_engine`, `pg_lake_iceberg`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_lake_table";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_table';
```

该上游项目与 `Snowflake` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
