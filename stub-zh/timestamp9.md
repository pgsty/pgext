
## 用法

> [timestamp9: PostgreSQL 的纳秒精度时间戳类型](https://github.com/optiver/timestamp9)

`timestamp9` 扩展提供纳秒精度的时间戳类型，底层以 64 位整数存储，即自 UNIX 纪元以来的纳秒数。

```sql
CREATE EXTENSION timestamp9;
```

### 数据类型

`timestamp9` 支持从 1970-01-01 到 2262-04-12 的时间戳，精度为纳秒。

### 输入格式

```sql
-- 标准 PostgreSQL 格式
SELECT '2019-09-19 08:30:05'::timestamp9;

-- 带时区的完整纳秒精度
SELECT '2019-09-19 08:30:05.123456789 +0200'::timestamp9;

-- 从 bigint 转换，自纪元以来的纳秒数
SELECT 1568878205123456789::bigint::timestamp9;
```

### 类型转换

- 与 `timestamp` 和 `timestamptz` 之间的相互转换
- 与 `date` 之间的相互转换

转换过程中会保留纳秒精度。

### 运算符

```sql
-- 比较
SELECT '2019-09-19'::timestamp9 > '2019-09-18'::timestamp9;  -- true

-- 与 interval 的算术运算
SELECT '2019-09-19 23:00:00.123456789'::timestamp9 + interval '1 day';

-- 减法
SELECT '2019-09-20'::timestamp9 - '2019-09-19'::timestamp9;
```

### 函数

```sql
SELECT greatest('2019-09-19'::timestamp9, '2019-09-20'::timestamp9);
```

### 索引支持

支持 Btree 和 Hash 索引。
