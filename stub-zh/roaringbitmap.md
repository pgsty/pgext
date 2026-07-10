


## 用法

来源：

- [PGXN pg_roaringbitmap 1.2.0](https://pgxn.org/dist/pg_roaringbitmap/1.2.0/)
- [pg_roaringbitmap README](https://github.com/ChenHuajun/pg_roaringbitmap)
- [pg_roaringbitmap CHANGELOG](https://github.com/ChenHuajun/pg_roaringbitmap/blob/master/CHANGELOG.md)

`pg_roaringbitmap` 安装的 PostgreSQL 扩展名是 `roaringbitmap`，它提供基于 Roaring Bitmap 的压缩位图类型和集合运算函数。适合紧凑存储整数集合、快速并集/交集、用户分群、faceting 和位图聚合。

v1.2.0 增加了 `rb_runoptimize()` / `rb64_runoptimize()` 用于缩小位图二进制尺寸，保留旧拼写 `rb_exsit` 以兼容老版本，加入 PostgreSQL 19 兼容，并在 receive 函数中校验不可信位图输入。

### 创建扩展

```sql
CREATE EXTENSION IF NOT EXISTS roaringbitmap;

SET roaringbitmap.output_format = 'array';
SELECT rb_build(ARRAY[1, 2, 3, 4, 5]);
```

`roaringbitmap.output_format` 可设为 `bytea` 或 `array`。默认输出格式是 `bytea`，更适合大位图；`array` 适合交互式查看。

### 数据类型

- `roaringbitmap` 存储无符号 32 位整数集合，逻辑范围为 `[0, 4294967296)`。
- `roaringbitmap64` 存储无符号 64 位整数集合，使用 `rb64_` 函数族。

```sql
CREATE TABLE cohorts (
  segment text PRIMARY KEY,
  users32 roaringbitmap,
  users64 roaringbitmap64
);
```

### 构建与转换

```sql
INSERT INTO cohorts(segment, users32)
VALUES ('trial', rb_build(ARRAY[1, 2, 3, 100, 200]));

INSERT INTO cohorts(segment, users32)
SELECT 'active', rb_build_agg(id)
FROM generate_series(1, 100000) AS id;

SELECT rb_cardinality(users32) FROM cohorts WHERE segment = 'active';
SELECT rb_to_array(users32) FROM cohorts WHERE segment = 'trial';
SELECT rb_iterate(users32) FROM cohorts WHERE segment = 'trial';
```

64 位值使用 `rb64_build()`、`rb64_build_agg()`、`rb64_to_array()` 和 `rb64_iterate()`。

### 集合运算

```sql
SELECT rb_build(ARRAY[1,2,3]) | rb_build(ARRAY[3,4,5]);  -- 并集
SELECT rb_build(ARRAY[1,2,3]) & rb_build(ARRAY[3,4,5]);  -- 交集
SELECT rb_build(ARRAY[1,2,3]) # rb_build(ARRAY[3,4,5]);  -- 异或
SELECT rb_build(ARRAY[1,2,3]) - rb_build(ARRAY[3,4,5]);  -- 差集

SELECT rb_build(ARRAY[1,2,3]) | 9;                       -- 添加元素
SELECT rb_build(ARRAY[1,2,3]) - 2;                       -- 移除元素
```

也支持包含与重叠操作符：

```sql
SELECT rb_build(ARRAY[1,2,3]) @> rb_build(ARRAY[2,3]);
SELECT rb_build(ARRAY[2,3]) <@ rb_build(ARRAY[1,2,3]);
SELECT rb_build(ARRAY[1,2,3]) && rb_build(ARRAY[3,4,5]);
```

### 基数与范围函数

```sql
SELECT rb_and_cardinality(a.users32, b.users32);
SELECT rb_or_cardinality(a.users32, b.users32);
SELECT rb_xor_cardinality(a.users32, b.users32);
SELECT rb_andnot_cardinality(a.users32, b.users32);

SELECT rb_range(users32, 100, 1000);
SELECT rb_range_cardinality(users32, 100, 1000);
SELECT rb_fill(users32, 100, 200);
SELECT rb_clear(users32, 100, 200);
SELECT rb_flip(users32, 100, 200);

SELECT rb_min(users32), rb_max(users32), rb_rank(users32, 500), rb_index(users32, 500);
SELECT rb_jaccard_dist(a.users32, b.users32);
```

64 位范围函数使用 `rb64_` 前缀。自 v1.1.0 起，多个 `rb64_` range/select 函数中 `range_end = 0` 表示无限上界。

### 聚合函数

```sql
SELECT rb_build_agg(user_id) FROM events;
SELECT rb_or_agg(users32) FROM cohorts;
SELECT rb_and_agg(users32) FROM cohorts;
SELECT rb_xor_agg(users32) FROM cohorts;

SELECT rb64_build_agg(user_id::bigint) FROM events;
SELECT rb64_or_agg(users64) FROM cohorts;
```

### 优化序列化尺寸

```sql
UPDATE cohorts
SET users32 = rb_runoptimize(users32);

UPDATE cohorts
SET users64 = rb64_runoptimize(users64);
```

`rb_runoptimize()` 和 `rb64_runoptimize()` 可在适合的数据分布下减小位图序列化尺寸。不要在高频写入路径中未经测试直接使用。

### 注意事项

- Pigsty 使用扩展文件名 `roaringbitmap.md`；上游包名是 `pg_roaringbitmap`。
- 源码构建需要 PostgreSQL 头文件和 CRoaring 依赖。README 提到发布前回归测试覆盖 PostgreSQL 13 及以上版本。
- `Makefile_native` 可以用 SIMD 指令编译，在匹配 CPU 上可能更快；但这种二进制放到缺少对应 CPU 特性的机器上可能触发 `SIGILL`。
- 大位图建议使用 `bytea` 输出；人工查看时再切到 `array` 输出。
