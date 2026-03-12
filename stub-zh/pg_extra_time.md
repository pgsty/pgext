

## 用法

> [pg_extra_time: PostgreSQL 的额外日期/时间函数与运算符](https://github.com/bigsmoke/pg_extra_time)

### 转换为秒（浮点数）

```sql
SELECT to_float('1970-01-01 00:00:00+0'::timestamptz);  -- 0.0
SELECT to_float('1 day 1 sec'::interval);                -- 86401.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tstzrange);  -- 130.0
```

也支持类型转换语法：

```sql
SELECT '1970-01-01 01:03:01+00'::timestamptz::float;    -- 3181.00
SELECT '1 day 1 sec 200 ms'::interval::float;            -- 86401.2
```

### 转换为天数

```sql
SELECT days('[2024-06-06,2024-06-08 06:00]'::tstzrange);  -- 3.25（小数天数）
SELECT whole_days('[2024-06-06,2024-06-08 18:00]'::tstzrange);  -- 2（仅整天数）
SELECT days('10 days 12 hours'::interval);                -- 10.5
SELECT whole_days('10 days 20 hours'::interval);          -- 10
```

### 从范围中提取时间间隔

```sql
SELECT to_interval('[2024-01-01,2024-01-05]'::tstzrange);  -- 4 days
```

### Each 函数（从范围生成序列）

```sql
SELECT * FROM each_subperiod('[2024-01-01,2024-01-05]'::tstzrange, '1 day'::interval);
```

### 运算符

```sql
-- 范围持续时间转换为浮点数（秒）
SELECT '[epoch,1970-01-01T01:03:01+00]'::tstzrange::float;
-- 时间间隔转换为浮点数（秒）
SELECT '10 seconds 100 milliseconds'::interval::float;  -- 10.100
```
