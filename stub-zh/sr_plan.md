## 用法

来源：

- [官方 README](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/README.md)
- [1.2 扩展 SQL](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/init.sql)
- [配置实现](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/sr_plan.c)

`sr_plan` 1.2 捕获 PostgreSQL 执行计划，并可强制之后匹配的查询复用已启用的保存计划。它类似 outline 或计划基线工具，预加载后会改变全局优化器行为；只能用于受控查询，并应预先准备回退方案。

### 核心流程

预加载模块并重启 PostgreSQL，再在目标数据库中创建扩展。

```conf
shared_preload_libraries = 'sr_plan'
```

```sql
CREATE EXTENSION sr_plan;

SET sr_plan.write_mode = true;
SELECT * FROM orders WHERE customer_id = 42;
SET sr_plan.write_mode = false;

SELECT query_hash, plan_hash, enable, query
FROM sr_plans
ORDER BY query_hash;

UPDATE sr_plans SET enable = true WHERE query_hash = 12345;
```

捕获模式会保存之后所有查询的计划，包括重复项。启用前应检查具体记录，不要无差别启用所有捕获计划。

### 主要对象与设置

- `sr_plans`：保存查询哈希、查询标识、计划哈希、计划文本、启用状态以及引用的表/索引 OID。
- `sr_plan.write_mode`：启用计划捕获，属于超级用户设置。
- `sr_plan.enabled`：启用优化器钩子，属于超级用户设置。
- `sr_plan.log_usage`：控制使用缓存计划时的消息级别。
- `_p(anyelement)`：标记计划匹配时应视为参数的常量。
- `show_plan(query_hash, index, format)`：以 `text`、`json`、`xml` 或 `yaml` 显示保存计划。

扩展创建的事件触发器会在引用的表或索引被删除时清理保存计划。但模式、统计信息和数据分布变化仍可能让曾经合适的强制计划变得有害，因此应监控并清理 `sr_plans`。若同时预加载 `pg_stat_statements`，上游要求它在预加载列表中位于 `sr_plan` 之后。项目已无人维护，使用前必须针对准确的 PostgreSQL 大版本验证兼容性。
