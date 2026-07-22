## 用法

来源：

- [官方 v2.28.1 control README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.28.1/_control/README.md)
- [官方 v2.28.1 安装 SQL](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.28.1/_control/install.sql)
- [官方 v2.28.1 发布树](https://github.com/dventimisupabase/pg_flight_recorder/tree/v2.28.1)
- [当前官方仓库树](https://github.com/dventimisupabase/pg_flight_recorder/tree/main)

`pgfr_control` 2.28.1 是 PostgreSQL Flight Recorder 历史版本中的纯 SQL 控制组件。它读取 `pgfr_record` 收集的快照，用于诊断 vacuum 行为、估算膨胀与增长、监控 OID 消耗并推荐 autovacuum scale factor。它只报告建议，不会自动应用设置。

### 安装标签版本组件

上游 v2.28.1 源码把 `pgfr_control` 作为 `_control` 下的安装脚本发布，而不是独立的 `CREATE EXTENSION` 包。请先安装并启用同版本的 `_record` 组件，再用 `psql` 运行标签版本的 control 脚本。

```sql
\i _record/install.sql
SELECT pgfr_record.enable();

\i _control/install.sql

SELECT *
FROM pgfr_control.vacuum_diagnostic('public.orders'::regclass);

SELECT *
FROM pgfr_control.vacuum_control_report(now() - interval '1 hour', now());
```

必须使用同一个 v2.28.1 标签中的文件。control 安装器会检查 `pgfr_record.config` 的 schema-version 行，如果 recorder core 不存在就会中止。

### 主要报告

- Vacuum 控制：`vacuum_control_mode(oid)` 返回 `normal`、`catch_up` 或 `safety`；`compute_recommended_scale_factor(oid)` 给出建议设置；`vacuum_diagnostic(oid)` 分类健康状态；`vacuum_control_report(start, end)` 覆盖所有已观测表。
- 死元组：`dead_tuple_growth_rate(oid, interval)`、`dead_tuple_trend(oid, interval)` 和 `time_to_budget_exhaustion(oid, budget)` 使用快照历史。
- 膨胀与大小：`estimate_table_bloat(oid)`、`bloat_report(interval)` 和 `table_size_growth_rate(oid, interval)` 无需运行 `pgstattuple` 即可生成派生估算。
- OID：`oid_consumption_rate(interval)` 与 `time_to_oid_exhaustion()` 估算集群级消耗。

`vacuum_diagnostic` 会输出包括 `NOT_SCHEDULED`、`RUNNING_BUT_LOSING`、`BLOCKED` 和 `HEALTHY` 在内的状态。应把它们视为需要调查的诊断信号，而不是盲目修改设置的命令。

### 版本与运维边界

结果取决于 `pgfr_record` 快照的可用性、采集间隔与保留时间。历史稀疏时会返回空值或低置信度速率；膨胀和耗尽值只是估算，不是物理检查结果或保证。请结合工作负载、锁、事务年龄及表级 reloption 上下文审查建议。

2.28.1 是固定的历史布局。当前上游代码树已不再把 `_control` 作为独立组件发布，因此不要把这些脚本与后续 recorder schema 混用。目录打包可能以不同方式包装 SQL，但上述权威标签源码流程并不能支持凭空编写 `CREATE EXTENSION pgfr_control` 的源码安装步骤。
