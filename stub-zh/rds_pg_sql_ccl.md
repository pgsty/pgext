## 用法

来源：

- [华为云 RDS for PostgreSQL rds_pg_sql_ccl 指南](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0073.html)

`rds_pg_sql_ccl` 是华为云 RDS for PostgreSQL 的 SQL 并发控制内核扩展。它既可以限制 SQL 总并发量，也可以按 query ID 设置独立的运行与等待上限。它由服务商托管，不是可自行安装的社区软件包。

### 可用性与启用

服务商指南列出的版本包括 PostgreSQL 17 全部小版本、16.2+、15.4+、14.8+、13.11+、12.15+ 和 11.20+。先检查实际实例，再通过华为控制函数创建内核扩展并启用内核参数：

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_pg_sql_ccl';

SELECT control_extension('create', 'rds_pg_sql_ccl');

SHOW rds_pg_sql_ccl.enable_ccl;
SHOW rds_pg_sql_ccl.max_concurrent_sql;
SHOW rds_pg_sql_ccl.max_enabled_rules;
```

通过 RDS 参数控制把 `rds_pg_sql_ccl.enable_ccl` 设为 on。`rds_pg_sql_ccl.max_concurrent_sql` 设置优先级更高的实例级总上限；小于等于零表示不限制。`rds_pg_sql_ccl.max_enabled_rules` 限制同时生效的规则数量。

### 按查询规则

根据 SQL 文本创建规则，显式启用，然后检查活动规则集合：

```sql
WITH new_rule AS (
  SELECT rds_pg_sql_ccl.add_ccl_rule_by_query(
    'SELECT pg_sleep(1)',
    2,
    10,
    'public'
  ) AS rule_id
)
SELECT rds_pg_sql_ccl.enable_ccl_rule(rule_id)
FROM new_rule;

SELECT * FROM rds_pg_sql_ccl.get_all_enabled_rule;
SELECT * FROM rds_pg_sql_ccl.get_activity_query_status;
SELECT * FROM rds_pg_sql_ccl.get_current_db_ccl_rule;
```

运行上限是 `max_concurrency`；`max_waiting` 控制额外允许等待的执行数量。新建规则在调用 `enable_ccl_rule` 前不会生效。

### 重要对象

- `get_query_id(query_string, search_path)` 为不含 bind variable 的 SQL 计算 query ID。
- `add_ccl_rule_by_query(...)` 根据 SQL 文本创建规则；`add_ccl_rule_by_queryid(...)` 根据已知 query ID 创建规则。
- `enable_ccl_rule(bigint)`、`disable_ccl_rule(bigint)`、`disable_all_ccl_rule()`、`delete_ccl_rule(bigint)` 和 `update_ccl_rule(...)` 管理规则状态。
- `get_all_enabled_rule`、`get_activity_query_status` 和 `get_current_db_ccl_rule` 分别暴露已启用规则、等待活动和当前数据库的规则定义。

### 绑定变量与生命周期

`get_query_id` 和基于文本的规则创建无法识别带 bind variable 的 prepared statement。应先执行 prepared statement，从 `pg_stat_statements` 获取其 `queryid`，再调用 `add_ccl_rule_by_queryid`。必须针对目标引擎与 search path 验证 query ID 和规范化文本。

实例重启后不会保留 enabled 状态，必须重新启用规则。read replica 会同步 primary 规则并启用它们，但规则只作用于启用后才开始的语句。规则在单个数据库内按 query ID 唯一，不同数据库可以复用相同 query ID。

### 运维说明

并发限制可能拒绝或排队生产查询，因此参数应来自负载和资源测量。实例级总上限优先于按查询规则。安装、参数、引擎升级和可用性都应通过华为云 RDS 管理；依赖该功能前，应测试故障切换、重启后的重新启用、replica 行为、队列饱和以及应用超时。
