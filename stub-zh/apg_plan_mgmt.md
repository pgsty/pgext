## 用法

来源：

- [Aurora PostgreSQL 查询计划管理指南](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Optimize.html)
- [查询计划管理函数参考](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Optimize.Functions.html)
- [Aurora PostgreSQL 查询计划管理发布说明](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/auroraqpm.updates.html)
- [Aurora PostgreSQL 扩展版本矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apg_plan_mgmt` 是 Amazon Aurora PostgreSQL 托管的查询计划管理扩展。它捕获优化器计划，将其记录到 `apg_plan_mgmt.dba_plans`，并允许管理员批准、拒绝、优选、禁用、验证或删除受管计划。目前 Aurora PostgreSQL 18.3 列出的版本为 3.0；其他引擎版本提供的扩展版本不同。

### 启用并检查计划捕获

Aurora 管理员必须先将集群参数 `rds.enable_plan_management` 设为 `1`，挂载所需的自定义参数组，重启实例，再以 `rds_superuser` 创建扩展。较为保守的手工捕获会话如下：

```sql
CREATE EXTENSION apg_plan_mgmt;

SET apg_plan_mgmt.capture_plan_baselines = manual;
EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SET apg_plan_mgmt.capture_plan_baselines = off;

SELECT sql_hash, plan_hash, status, enabled
FROM apg_plan_mgmt.dba_plans;
```

只有在审查捕获到的基线并用代表性参数完成测试后，才应设置 `apg_plan_mgmt.use_plan_baselines`。仅向确实需要查看或管理计划的用户授予 `apg_plan_mgmt` 角色。

### 运维边界

`Approved`、`Unapproved`、`Rejected` 和 `Preferred` 等计划状态会直接影响优化器可以选择的执行计划。应把状态变更当作生产配置变更，并保留回滚路径。

`apg_plan_mgmt.evolve_plan_baselines` 会以调用者权限通过 `EXPLAIN ANALYZE` 运行捕获的语句。因此语句可能执行 DML 或调用易变函数；应在非生产副本上评估候选计划，或者只传入副作用已明确的语句。可用性和可安装版本由 Aurora 引擎版本控制，并不存在可移植的社区软件包。

AWS 的 3.0 版本发布说明记录了计划未能捕获问题的修复，并新增 `evolve_plan_baselines` 对 `sql_hash` 变化的支持。升级后应重新验证捕获与演进流程，不要假定现有基线无需变更即可继续使用。
