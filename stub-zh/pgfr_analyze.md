## 用法

来源：

- [官方扩展控制文件](https://github.com/dventimisupabase/pg_flight_recorder/blob/34517280f70b67ae8c8f99d18515550b629c9cd2/pgfr_analyze/extension.control)
- [官方项目或服务商页面](https://dventimisupabase.github.io/pg_flight_recorder/)
- [官方主文档](https://database.dev/dventimi/pgfr_analyze)

`pgfr_analyze` — 为 pgfr_record 采集的数据提供报告、异常检测、事故取证和容量规划分析函数。

已复核目录快照记录的版本为 `2.29.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgfr_record`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgfr_analyze";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfr_analyze';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
