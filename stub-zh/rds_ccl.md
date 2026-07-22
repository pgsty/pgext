## 用法

来源：

- [阿里云 rds_ccl SQL 限流文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-ccl-extension-to-perform-sql-throttling-on-an-apsaradb-rds-for-postgresql-instance)
- [阿里云 RDS 产品概览](https://www.alibabacloud.com/help/en/rds/product-overview/what-is-apsaradb-rds-what-is-apsaradb-rds)

`rds_ccl` 是阿里云 ApsaraDB RDS for PostgreSQL 的服务扩展，用于限制选定规范化语句的并发执行数，并让超量工作排队。它是服务商专有组件，并非可移植到自管 PostgreSQL 的扩展。

### 核心流程

启用实例参数 `rds_enable_ccl=on`，并把 `compute_query_id=auto` 或 `on`，通过 RDS 扩展管理安装扩展，然后使用高权限 RDS 账号创建规则。下例把主库上某条语句限制为三个活动执行和两个等待者：

```sql
CREATE EXTENSION rds_ccl;

SELECT rds_add_ccl_rule(
    $$SELECT * FROM ccl_tbl;$$,
    1,
    3,
    2,
    true,
    'limit constant select',
    ''
);

SELECT * FROM rds_show_current_db_ccl_rule();
```

规则按规范化 `query_id` 匹配，因此只有字面量不同的语句可能共享一条规则。可用 `rds_get_query_id(text)` 计算标识符，或使用 `rds_add_ccl_rule_with_query_id(...)` 根据已知标识符创建规则。

### 规则操作

- `node_tag` 值 1、2 和 3 分别允许在主库、只读实例或两者上加载；只读实例仍需要显式调用 `rds_load_ccl_rule(integer)`。
- `rds_enabled_ccl_rule` 显示跨数据库的已启用规则，`rds_show_current_db_ccl_rule()` 显示当前数据库的完整详情。
- `rds_enable_ccl_rule(integer)`、`rds_disable_ccl_rule(integer)` 和 `rds_disable_all()` 控制持久规则状态。
- `rds_update_ccl_rule(integer, integer, integer)` 只更改并发和队列上限。`rds_unload_ccl_rule(integer, varchar)` 从内存卸载规则，`rds_del_ccl_rule(integer)` 则删除它。

### 服务边界

规则在主库创建。已启用的主库规则在创建或启用时自动加载，但内存中的规则不能跨实例重启或主备切换保留；之后必须重新加载。用于只读实例的规则必须具有兼容的 `node_tag`，并在该实例上加载。

过低的并发值会卡住正常流量，而超出 `max_waiting` 会中止新事务。应在与生产相同的数据库和 `search_path` 上下文中生成标识符，逐步上线限额，监控活动与排队语句，并保留经测试的禁用和卸载流程。可用性、版本、权限和行为仍受当前阿里云 RDS 内核与服务策略管辖。
