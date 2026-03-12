


## 用法

> [plan_filter: 根据执行计划过滤语句](https://github.com/pgexperts/pg_plan_filter)

`pg_plan_filter` 模块在执行前根据配置的条件检查语句，如果违反条件则抛出错误。这允许管理员阻止在生产数据库上执行某些查询。

目前支持的唯一条件是语句计划的最大允许估算成本。

### 配置

在 `postgresql.conf` 中添加：

```ini
shared_preload_libraries = 'plan_filter'
plan_filter.statement_cost_limit = 100000.0
```

`statement_cost_limit` 必须设置为非零值才能启用过滤。默认值为 `0`（不过滤）。

### GUC 参数

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|------|
| `plan_filter.statement_cost_limit` | float | `0` | 最大允许估算计划成本。`0` 表示禁用过滤 |
| `plan_filter.limit_select_only` | bool | `false` | 为 true 时仅过滤 `SELECT` 语句 |

### 示例

全局阻止高成本查询：

```ini
plan_filter.statement_cost_limit = 100000.0
```

仅限制 SELECT 语句（注意：SELECT != READONLY，因为 SELECT 也可以修改数据）：

```ini
plan_filter.limit_select_only = true
```

当模块以非零 `statement_cost_limit` 运行时，也会阻止对高成本查询执行 `EXPLAIN`。临时绕过过滤：

```sql
BEGIN;
SET LOCAL plan_filter.statement_cost_limit = 0;
EXPLAIN SELECT ...;
COMMIT;
```

按用户覆盖限制：

```sql
ALTER USER can_run_expensive SET plan_filter.statement_cost_limit = 0;
ALTER USER only_cheap_queries SET plan_filter.statement_cost_limit = 10000;
```

### 注意事项

`statement_cost_limit` 基于**估算成本**取消计划。PostgreSQL 规划器返回的成本估算可能与实际查询执行时间无关。请准备好应对误报取消的情况，并将限制值设置得宽裕一些。
