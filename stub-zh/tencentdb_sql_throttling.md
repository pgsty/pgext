## 用法

来源：

- [腾讯云 PostgreSQL SQL 限流文档](https://intl.cloud.tencent.com/zh/document/product/409/64750)
- [腾讯云 PostgreSQL 产品页](https://cloud.tencent.com/product/postgres)

`tencentdb_sql_throttling` 是腾讯云托管扩展，按规范化查询 ID 限制并发执行数。它不是可移植的社区包：可用性、支持的内核构建、预加载修改、权限、HA 行为和升级均由腾讯云控制。

```sql
CREATE EXTENSION tencentdb_sql_throttling;
SELECT tencentdb_sql_throttling.add_rule_with_queryid(
  497939862935121343, 0, 10, true
);
SELECT * FROM tencentdb_sql_throttling.rules;
```

需要由腾讯云支持把 `tencentdb_sql_throttling` 加入 `shared_preload_libraries`，这会重启实例。官方文档要求至少达到所列腾讯云内核修订版，并说明无法直接配置扩展查询协议语句。应从 `pg_stat_activity` 或 `pg_stat_statements` 获取查询 ID；`work_node` 为 0 时作用于主节点与只读节点，1 仅主节点，2 仅只读节点。

规则只有载入内存后才影响流量。重启、HA 事件或大版本升级会清除内存规则；应将规则持久化到 `tencentdb_sql_throttling.persistent_rules_table` 并显式重新加载。`drop_rule`、`drop_all_rules`、`change_rule_status` 及 load/dump 辅助函数会区分内存与持久状态；每次控制操作后都要核对两处状态。

应从保守的 `max_concurrency` 开始，监控 `current_concurrency`、`total_hit_count`、`reject_count`、延迟和应用错误，并保留紧急回滚路径。SQL、模式、规划器或版本变化后的查询 ID 改变，可能导致规则漏匹配或作用于非预期规范化查询，因此必须与厂商演练故障切换和升级行为。
