

## 用法

> [roaringbitmap: PostgreSQL 的压缩位图数据类型](https://github.com/ChenHuajun/pg_roaringbitmap)

`roaringbitmap` 扩展提供了压缩位图数据类型，用于对整数集合进行高效的集合运算。

```sql
CREATE EXTENSION roaringbitmap;
SET roaringbitmap.output_format = 'array';

SELECT rb_build('{1,2,3,4,5}'::int[]);  -- {1,2,3,4,5}
```

### 数据类型

- **`roaringbitmap`**：32 位位图，范围 [0, 4294967296)
- **`roaringbitmap64`**：64 位位图，范围 [0, 18446744073709551615)

输出格式通过以下方式控制：`SET roaringbitmap.output_format = 'array'` 或 `'bytea'`

### 运算符

| 运算符 | 说明 |
|----------|-------------|
| `\|` | 按位或（并集） |
| `&` | 按位与（交集） |
| `#` | 按位异或 |
| `-` | 差集（与非） |
| `\|`（与 int） | 添加元素 |
| `-`（与 int） | 移除元素 |
| `@>` | 包含 |
| `<@` | 被包含于 |
| `&&` | 重叠 |
| `=`, `<>` | 相等/不等 |

### 核心函数

```sql
-- 构建
SELECT rb_build(ARRAY[1,2,3,4,5]);

-- 分析
SELECT rb_cardinality(rb_build('{1,2,3}'::int[]));   -- 3
SELECT rb_to_array(rb_build('{1,2,3}'::int[]));      -- {1,2,3}
SELECT rb_iterate(rb_build('{1,2,3}'::int[]));        -- 返回集合

-- 集合运算基数
SELECT rb_and_cardinality(a, b);
SELECT rb_or_cardinality(a, b);
SELECT rb_xor_cardinality(a, b);
SELECT rb_andnot_cardinality(a, b);

-- 范围操作
SELECT rb_range(bitmap, 2, 5);   -- 提取范围 [2, 5)
SELECT rb_fill(bitmap, 0, 10);   -- 添加范围 [0, 10)
SELECT rb_clear(bitmap, 3, 7);   -- 移除范围 [3, 7)
SELECT rb_flip(bitmap, 0, 10);   -- 翻转范围 [0, 10)

-- 元素访问
SELECT rb_min(bitmap);            -- 最小元素
SELECT rb_max(bitmap);            -- 最大元素
SELECT rb_rank(bitmap, 5);        -- 统计 <= 5 的元素数量
SELECT rb_index(bitmap, 3);       -- 元素的从零开始的位置

-- 工具函数
SELECT rb_is_empty(bitmap);
SELECT rb_jaccard_dist(a, b);     -- Jaccard 相似度
```

### 聚合函数

```sql
SELECT rb_build_agg(id) FROM table;       -- 从行构建位图
SELECT rb_or_agg(bitmap) FROM table;      -- 合并所有位图（并集）
SELECT rb_and_agg(bitmap) FROM table;     -- 合并所有位图（交集）
SELECT rb_xor_agg(bitmap) FROM table;     -- 合并所有位图（异或）
```

64 位版本使用 `rb64_` 前缀。
