## 用法

来源：

- [官方资料](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/README.md)
- [官方资料](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/pg_qi.control)
- [官方资料](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/LICENSE)
- [官方资料](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/sql/pg_qi--0.0.1.sql)

pg_qi 0.0.1 是用于 pg_duckdb 集成的 pre-alpha 影子分类器，实际意图始终为 pass_through，并保留下游计划。自动路由、遥测导出和策略管理尚未实现。

### 启用诊断

按 pg_duckdb、pg_qi 的顺序预加载并重启，再以超级用户安装 pg_qi。已测试 PostgreSQL 16–18；PostgreSQL 19 测试固定于 RC2，不能据此声明正式版支持。集成要求真实且可重置的 duckdb.force_execution 设置，以及文档指定的依赖修订。

```ini
shared_preload_libraries = 'pg_duckdb,pg_qi'
```

```sql
CREATE EXTENSION pg_qi;
SELECT qi.version(), qi.is_preloaded();
SHOW pg_qi.enabled;
SHOW pg_qi.shadow_reason;
SHOW pg_qi.shadow_classifier_status;
SHOW pg_qi.shadow_observed_engine;
SHOW pg_qi.shadow_verification;
```

### 解释边界

使用 SHOW 查看已完成的后端局部影子记录，以免生成另一份计划。unproved_outermost 不能证明计划直接来自客户端；未发现引擎节点仍表示 unknown/unverified。pg_qi_spike_test 是显式条件构建的测试辅助扩展，不属于正常安装，也不是生产路由器。
