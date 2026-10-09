## 用法

来源：

- [README.md](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/README.md)
- [Makefile](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/Makefile)
- [License](https://github.com/pgexperts/pg_plan_filter/blob/c1afc73dacae8688456bdeef72ed6a3030c500ba/License)

`plan_filter` 是 pg_plan_filter 1.0.0 的可加载库，根据配置的计划估算成本拒绝语句或事务。支持 PostgreSQL 14–18，没有需要创建的 SQL 扩展。

### 核心用法

```ini
shared_preload_libraries = 'plan_filter'
plan_filter.statement_cost_limit = 100000.0
plan_filter.transaction_cost_limit = 500000.0
plan_filter.filter_select_only = false
```

```sql
SELECT current_setting('plan_filter.module_loaded', true) IS NOT NULL AS active;
```

### 运行边界

将库追加到已有预加载列表并重启；开发和测试也可使用 `LOAD`。两个成本上限都默认为零，即不限制。`plan_filter.statement_cost_limit` 限制单条计划，也可能拒绝高成本的普通 EXPLAIN。`plan_filter.transaction_cost_limit` 对每次执行记账，包括反复执行预备语句；回滚到保存点不会返还预算。普通 EXPLAIN 不消耗事务预算，包含执行计划的工具命令则会计入。

`plan_filter.filter_select_only` 为真时两个上限只过滤 SELECT，但 SELECT 仍可能产生副作用。`plan_filter.module_loaded` 是只读的加载标志。修改上限需要超级用户或明确授予的参数设置权限，拒绝操作的 SQLSTATE 为 54001。成本是估算值而非运行时间，普通角色可以降低规划器成本设置绕过阈值。因此它适合防止意外负载，不能单独承担对抗恶意会话的安全边界。
