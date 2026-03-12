

## 用法

> [pg_duration: PostgreSQL 的 ISO 8601 持续时间类型](https://github.com/jkosh44/pg_duration)

`pg_duration` 扩展提供了一个 `duration` 类型，用 8 字节以微秒为单位存储经过时间，比内置的 `interval` 类型更简单且比较行为更一致。

```sql
CREATE EXTENSION pg_duration;
```

### 数据类型

`duration` 类型表示绝对经过时间，不包含日历组件（没有月或日）。有效输入接受任何不超过小时单位的 PostgreSQL interval 语法。

```sql
SELECT '01:30:00'::duration;
SELECT '2 hours 30 minutes'::duration;
```

### 运算符

- **算术运算**：duration 之间的 `+`、`-`；与 `float8` 的 `*`、`/`；一元 `-`
- **比较运算**：`<`、`<=`、`>`、`>=`、`=`、`<>`

### 函数

```sql
-- 从各组件构造
SELECT make_duration(hours => 2, mins => 30, secs => 15.5);

-- 检查是否有限
SELECT isfinite('01:00:00'::duration);

-- 截断到指定精度
SELECT date_trunc('hour', '02:45:30'::duration);

-- 提取子字段
SELECT date_part('minute', '02:45:30'::duration);
SELECT extract_duration('hour', '02:45:30'::duration);
```

### 类型转换

duration 可隐式转换为 `interval`。将 `interval` 转换为 `duration` 需要显式类型转换。

### 聚合与索引

支持标准聚合函数（`avg`、`count`、`max`、`min`、`sum`）以及 B-tree 和 Hash 索引。
