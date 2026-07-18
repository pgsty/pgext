## 用法

来源：

- [官方扩展控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/pg_lake_benchmark.control)
- [官方上游文档](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/README.md)

`pg_lake_benchmark` — 为 pg_lake、Iceberg 和普通堆表提供 ClickBench、TPC-H 与 TPC-DS 数据生成及基准查询。

已复核目录快照记录的版本为 `3.4`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_lake`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_lake_benchmark";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_benchmark';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
