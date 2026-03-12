

## 用法

> [extra_window_functions: PostgreSQL 的附加窗口函数](https://github.com/xocolatl/extra_window_functions)

提供模拟 SQL 标准中 PostgreSQL 语法尚不支持的窗口函数，以及 `flip_flop` 等新颖函数。

```sql
CREATE EXTENSION extra_window_functions;
```

### 模拟 SQL 标准的函数

| 函数 | 说明 |
|---|---|
| `lag_ignore_nulls(expr [, offset [, default]])` | 跳过 NULL 值的 LAG |
| `lead_ignore_nulls(expr [, offset [, default]])` | 跳过 NULL 值的 LEAD |
| `first_value_ignore_nulls(expr)` | 跳过 NULL 的 FIRST_VALUE |
| `last_value_ignore_nulls(expr)` | 跳过 NULL 的 LAST_VALUE |
| `nth_value_from_last(expr, offset)` | 从窗口帧末尾计数的 NTH_VALUE |
| `nth_value_ignore_nulls(expr, offset)` | 跳过 NULL 的 NTH_VALUE |
| `nth_value_from_last_ignore_nulls(expr, offset)` | 从末尾计数且跳过 NULL 的 NTH_VALUE |

### 扩展 SQL 标准的函数（带默认值）

| 函数 | 说明 |
|---|---|
| `first_value_ignore_nulls(expr, default)` | 超出窗口帧时返回默认值的 FIRST_VALUE |
| `last_value_ignore_nulls(expr, default)` | 超出窗口帧时返回默认值的 LAST_VALUE |
| `nth_value_from_last(expr, offset, default)` | 带默认值的从末尾计数 NTH_VALUE |
| `nth_value_ignore_nulls(expr, offset, default)` | 跳过 NULL 且带默认值的 NTH_VALUE |
| `nth_value_from_last_ignore_nulls(expr, offset, default)` | 组合从末尾计数、跳过 NULL、带默认值 |

### 非标准函数

| 函数 | 说明 |
|---|---|
| `flip_flop(expr [, expr])` | 触发器运算符：在第一个表达式为真之前返回 false，之后返回 true 直到第二个表达式匹配 |

### 示例

```sql
-- 等价于 SQL 标准: NTH_VALUE(x, 3) FROM LAST IGNORE NULLS OVER w
SELECT nth_value_from_last_ignore_nulls(x, 3) OVER w FROM t WINDOW w AS (ORDER BY id);

-- 前向填充：携带最近一个非空值
SELECT lead_ignore_nulls(val, 1) OVER (ORDER BY ts) FROM measurements;
```
