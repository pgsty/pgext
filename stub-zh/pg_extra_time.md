
## 用法

> 来源：[pg_extra_time upstream README](https://github.com/bigsmoke/pg_extra_time)、[PGXN pg_extra_time](https://pgxn.org/dist/pg_extra_time/)、[local metadata](../db/extension.csv)。

`pg_extra_time` 提供一组小型 SQL 函数和类型转换，用于处理 PostgreSQL 核心函数本身不太方便表达的日期/时间、interval 和 range 计算。

```sql
CREATE EXTENSION pg_extra_time;
```

### 转换为秒（浮点数）

对时间戳、时间戳范围和 interval，可以使用 `to_float(...)` 或显式转换到 `float`/`double precision`。时间戳按 Unix epoch 起算；范围和 interval 按持续时间秒数计算。

```sql
SELECT to_float('1970-01-01 00:00:00+0'::timestamptz);  -- 0.0
SELECT to_float('1970-01-01 00:00:00+0'::timestamp);    -- 0.0
SELECT to_float('1 day 1 sec'::interval);                -- 86401.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tstzrange);  -- 130.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tsrange);    -- 130.0
```

也支持类型转换语法：

```sql
SELECT '1970-01-01 01:03:01+00'::timestamptz::float;    -- 3181.00
SELECT '1 day 1 sec 200 ms'::interval::float;            -- 86401.2
SELECT '[epoch,1970-01-01T01:03:01+00]'::tstzrange::float;  -- 3181.00
```

### 转换为天数

需要保留小数时使用 `days(...)`；需要完整天数的整数结果时使用 `whole_days(...)`。

```sql
SELECT days('[2024-06-06,2024-06-09)'::daterange);       -- 2
SELECT days('[2024-06-06,2024-06-08 06:00]'::tstzrange);  -- 3.25 (fractional days)
SELECT whole_days('[2024-06-06,2024-06-08 18:00]'::tstzrange);  -- 2
SELECT days('10 days 12 hours'::interval);                -- 10.5
SELECT whole_days('10 days 20 hours'::interval);          -- 10
```

`whole_days(interval)` 处理负 interval 时，会先对绝对天数向下取整，再应用符号。

### 统计日期部件

`date_part_parts(part, subpart, timestamp with time zone, timezone)` 返回给定时间戳和时区下，一个较大的日期部件包含多少个较小的日期部件。这对处理夏令时等导致“一天不总是 24 小时”的计算很有用。

```sql
SELECT date_part_parts('month', 'days', '2024-02-12'::timestamptz, 'UTC');  -- 29
SELECT date_part_parts('year', 'days', '2024-08-23'::timestamptz, 'UTC');   -- 366
```

### 构造并拆分范围

使用 `make_tstzrange` 或 `make_tsrange` 从时间戳和 interval 构造范围，也支持负 interval。

```sql
SELECT make_tstzrange('2024-01-05 00:00+00'::timestamptz, '-4 days'::interval);
SELECT make_tsrange('2024-01-01 00:00'::timestamp, '4 days'::interval, '[)');
```

`each_subperiod(tstzrange, interval, round_remainder integer DEFAULT 0)` 将时间戳范围拆分成按 interval 大小切分的片段。余数策略为：`1` 向上补齐到完整片段，`0` 保留最后一个不完整片段，`-1` 丢弃余数。

```sql
SELECT *
FROM each_subperiod(
  '[2023-01-01,2023-04-02)'::tstzrange,
  '1 month'::interval,
  0
);
```

### 提取 interval 与求余

`to_interval(tstzrange)` 使用月、日和微秒单位从时间戳范围中提取 interval。`to_interval(tstzrange, interval[])` 接受按从大到小顺序排列的显式单位，并通过丢弃余数向下取整。

```sql
SELECT to_interval('[2024-01-01,2024-01-05]'::tstzrange);  -- 4 days
SELECT to_interval(
  '[2024-01-01,2024-04-13 01:10]'::tstzrange,
  ARRAY['1 mon'::interval, '1 day'::interval, '1 hour'::interval]
);
```

当需要余数时，使用 `%` 或 `modulo(...)`。

```sql
SELECT '10 seconds 100 milliseconds'::interval % '3 seconds'::interval;
SELECT '[2024-01-01,2024-01-10)'::tstzrange % '4 days'::interval;
```

### 注意事项

对于无界范围，`to_float(tstzrange)` 和 `to_float(tsrange)` 返回正无穷或负无穷；对于空范围返回 `0`。扩展有意不提供整数转换；需要整数天数时请使用 `whole_days(...)`。`extract_days(interval)` 和 `extract_interval(tstzrange, interval[])` 等已弃用别名仍为兼容性保留，但上游推荐改用 `whole_days(...)` 和 `to_interval(...)`。

### 参考

常用公开函数：

| 函数 | 用途 |
|----------|-----|
| `current_timezone()` | 返回当前活跃的 `pg_timezone_names` 行 |
| `date_part_parts(...)` | 统计较大日期部件中包含的较小日期部件数量 |
| `days(...)` | 根据输入类型返回小数或整数天数 |
| `whole_days(...)` | 从 interval 或时间戳范围中取完整天数 |
| `to_float(...)` | 从时间戳、时间戳范围或 interval 得到秒数 |
| `to_interval(...)` | 从 `tstzrange` 提取 interval |
| `make_tsrange(...)` / `make_tstzrange(...)` | 从时间戳加 interval 构造范围 |
| `each_subperiod(...)` | 将 `tstzrange` 拆分为子范围 |
| `modulo(...)` / `%` | interval 或范围相除后的余数 |
