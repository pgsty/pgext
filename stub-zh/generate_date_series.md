## 用法

来源：

- [官方函数文档](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/doc/generate_date_series.md)
- [扩展 SQL](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/sql/generate_date_series.sql)
- [C 实现](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/src/generate_date_series.c)

`generate_date_series` 增加了一个闭区间日期集合生成器，步长为整数天。当查询只需要简单的升序或降序日历日期序列时，它可以避免通过时间戳或间隔转换日期。

### 核心流程

```sql
CREATE EXTENSION generate_date_series;

SELECT date_val
FROM generate_date_series('2026-07-01'::date, '2026-07-07'::date) AS d(date_val);

SELECT date_val
FROM generate_date_series('2026-07-31'::date, '2026-07-01'::date, -7) AS d(date_val);
```

双参数调用默认使用一天步长。三参数形式接受正数或负数天数；只有序列恰好到达结束日期时才会包含该日期。

### 函数契约

- `generate_date_series(start_date date, end_date date, step_days integer DEFAULT 1) -> setof date` 是唯一面向用户的函数。
- 正步长会在当前日期小于或等于结束日期期间持续返回行。
- 负步长会在当前日期大于或等于结束日期期间持续返回行。
- 零步长会报出 `step size cannot equal zero`；方向背离结束日期的步长不返回任何行。

### 运维说明

版本 `1.0.0` 可重定位，且未声明预加载要求。上游构建面向 PostgreSQL `9.2` 或更高版本。该函数为 `IMMUTABLE` 且 `STRICT`，所以任一参数为 SQL `NULL` 时都不返回任何行。

扩展会直接按整数步长推进 PostgreSQL 内部日期值。面向用户的查询应限制生成范围，以免意外产生大量结果；依赖极端日期值前，应测试 PostgreSQL 支持日期边界附近的行为。
