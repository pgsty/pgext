## 用法

来源：

- [v2.0 README](https://github.com/xocolatl/extra_window_functions/blob/v2.0/README.md)
- [Version 2.0 SQL](https://github.com/xocolatl/extra_window_functions/blob/v2.0/extra_window_functions--2.0.sql)
- [1.0-to-2.0 upgrade SQL](https://github.com/xocolatl/extra_window_functions/blob/v2.0/extra_window_functions--1.0--2.0.sql)

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
-- Equivalent to SQL Standard: NTH_VALUE(x, 3) FROM LAST IGNORE NULLS OVER w
SELECT nth_value_from_last_ignore_nulls(x, 3) OVER w FROM t WINDOW w AS (ORDER BY id);

-- Fill forward: carry last non-null value
SELECT lag_ignore_nulls(val, 1) OVER (ORDER BY ts) FROM measurements;
```

### 2.0 版本与窗口框架

2.0 新增 `count_ties()`、`avg_rank()`、`avg_percent_rank()`、`group_number(boolean)`、`run_length(anyelement)`、`run_position(anyelement)`、`ema(double precision, double precision)`、`interpolate(double precision)` 和 `most_common(anyelement)`，用于同值分组、连续区间、平滑和插值。这些函数的排序与窗口框架规则各不相同；应明确选择 ORDER BY 和框架，并查阅官方函数说明，不能把它们全部当成聚合函数。

lag 示例读取前一个非空值；如果还需要保留当前非空值，应在以当前行结尾的框架上使用 `last_value_ignore_nulls`。已有安装可执行 `ALTER EXTENSION extra_window_functions UPDATE TO '2.0'`。这个可重定位的 C 扩展无需预加载，安装通常需要超级用户。上游支持 PostgreSQL 9.6–19；PostgreSQL 19 为部分函数提供标准 IGNORE NULLS，而 FROM LAST 和新增的 2.0 函数仍有独立用途。
