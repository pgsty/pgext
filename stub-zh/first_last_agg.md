

## 用法

> [first_last_agg: PostgreSQL 的 first 和 last 聚合函数](https://github.com/wulczer/first_last_agg)

提供 `first` 和 `last` 聚合函数，返回分组中的第一个或最后一个值，适用于任意元素类型。

```sql
CREATE EXTENSION first_last_agg;
```

### 函数

| 函数 | 说明 |
|---|---|
| `first(anyelement)` | 返回分组中的第一个值 |
| `last(anyelement)` | 返回分组中的最后一个值 |

### 示例

```sql
-- 获取第一个和最后一个值（使用 ORDER BY 以获得确定性结果）
SELECT id, first(val ORDER BY ts), last(val ORDER BY ts)
FROM events
GROUP BY id;

-- 不使用 ORDER BY 时，分组内的顺序是不确定的
SELECT department, first(salary ORDER BY hire_date)
FROM employees
GROUP BY department;
```
