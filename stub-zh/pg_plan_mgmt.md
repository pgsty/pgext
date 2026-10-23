## 用法

来源：

- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/README.md)
- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/pg_plan_mgmt.control)
- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/pg_plan_mgmt--1.0.sql)

`pg_plan_mgmt` 使用 PostgreSQL 19 的 `pg_plan_advice` 模块捕获并选择查询计划基线。这是实验性初版：SQL 接口、配置和存储格式均可能变化，上游不建议用于生产。

### 捕获与检查

加载两个模块，保留配置中已有的其他库，然后重启：

```ini
shared_preload_libraries = 'pg_plan_advice,pg_plan_mgmt'
```

```sql
CREATE EXTENSION pg_plan_mgmt;
SET pg_plan_mgmt.capture_plan_baselines = 'manual';
SELECT count(*) FROM pg_class;
SELECT query_id, plan_no, status FROM plan_mgmt.plans;
SET pg_plan_mgmt.capture_plan_baselines = 'off';
SET pg_plan_mgmt.use_plan_baselines = on;
```

捕获仅限超级用户会话，也允许其随后切换角色。首次捕获的计划自动获准，后续计划初始为未批准状态。已启用的优选计划优先于普通已批准计划；没有适用基线时，优化器可以回退到通常的最低成本计划。

### 管理与使用边界

`plan_mgmt.set_plan_status`、`plan_mgmt.set_plan_enabled`、`plan_mgmt.delete_plan` 和 `plan_mgmt.validate_plans` 用于管理计划；`plan_mgmt.add_plan` 导入 SQL 语句与计划轮廓。校验可以更新、禁用或删除失效条目，调用前应明确选择行为。

固定模式为 `plan_mgmt`，扩展不可迁移。SQL 接口默认撤销公共访问权限；授予管理函数权限会影响整个数据库的计划。捕获的 SQL 可能含字面值，应限制访问并保护持久化文件。持久化和过期清理工作进程要求共享预加载。会话加载可以使用内存基线，但这些基线在服务器重启后丢失。默认保留期限为 32 天，配置可限制语句和计划数量。
