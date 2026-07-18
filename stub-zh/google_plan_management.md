## 用法

来源：

- [Google Cloud 查询计划管理指南](https://docs.cloud.google.com/alloydb/docs/query-plan-management)
- [Google Cloud 扩展配置指南](https://docs.cloud.google.com/alloydb/docs/instance-configure-extensions)
- [Google Cloud 数据库标志参考](https://docs.cloud.google.com/alloydb/docs/reference/database-flags)

google_plan_management 是 AlloyDB for PostgreSQL 的供应商专有预览扩展。它记录查询模板、计划与执行统计，并允许获授权操作员批准或拒绝 AlloyDB 后续执行时可使用的计划。

### 启用与检查

首先将 AlloyDB 数据库标志 google_plan_management.enabled 设为 on。Google 文档说明修改该标志会重启实例。随后连接主实例并执行：

```sql
CREATE EXTENSION google_plan_management;
GRANT google_plan_management_role TO plan_operator;
SELECT * FROM google_plan_management.tracked_plans_view;
SELECT * FROM google_plan_management.managed_plans_view;
```

请将 plan_operator 替换为已有且严格控制的数据库角色。调用 approve_plan 或 deny_plan 前，应审查已跟踪计划与实测行为。

### 注意事项

- 这是 AlloyDB 功能，不是可移植的开源 PostgreSQL 扩展。可用性与行为取决于所选 AlloyDB 版本及 Google Cloud 服务条款。
- 该功能处于预览阶段并受预正式发布条款约束。支持可能有限，接口也可能变化。
- 功能启用后，计划跟踪默认开启，并会增加一定工作负载开销。仓库存有查询文本、计划与执行统计，应据此限制管理角色。
- 批准计划会改变未来优化器选择。若批准多个计划，AlloyDB 会选择估算成本最低者，但它未必在所有参数或数据分布下最快。
- 查询计划管理不支持分区表或分组集，只能在主实例上工作，最多存储 100,000 个唯一计划，且没有保留策略。
- 应保留回滚流程：google_plan_management.enable_steer_plans 标志可临时停止强制执行，同时保留观测数据。
