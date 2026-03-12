


## 用法

> [hll: 用于存储 HyperLogLog 数据的类型](https://github.com/citusdata/postgresql-hll)

`hll` 扩展提供了 HyperLogLog 数据类型，用于概率性的去重计数。它能够高效地近似 `COUNT(DISTINCT)` 操作，支持可配置的精度，并支持集合合并操作，使预聚合数据可以无损合并。

### 数据类型

- **`hll`** -- HyperLogLog 累加器，参数：`hll(log2m, regwidth, expthresh, sparseon)`
- **`hll_hashval`** -- 用于插入 HLL 结构的哈希值类型

### 核心函数

| 函数 | 描述 |
|------|------|
| `hll_empty()` | 创建空的 HLL |
| `hll_add(hll, hll_hashval)` | 向 HLL 添加一个哈希值 |
| `hll_cardinality(hll)` | 估算去重计数 |
| `hll_union(hll, hll)` | 合并两个 HLL |
| `hll_add_agg(hll_hashval)` | 将哈希值聚合为单个 HLL |
| `hll_union_agg(hll)` | 将多个 HLL 合并为一个 |
| `hll_print(hll)` | 显示 HLL 参数和内容 |

### 哈希函数

| 函数 | 输入类型 |
|------|---------|
| `hll_hash_boolean(boolean [, seed])` | boolean |
| `hll_hash_smallint(smallint [, seed])` | smallint |
| `hll_hash_integer(integer [, seed])` | integer |
| `hll_hash_bigint(bigint [, seed])` | bigint |
| `hll_hash_bytea(bytea [, seed])` | bytea |
| `hll_hash_text(text [, seed])` | text |
| `hll_hash_any(any [, seed])` | any（动态分发，较慢） |

### 运算符

| 运算符 | 功能 | 示例 |
|--------|------|------|
| `\|\|` | `hll_add` / `hll_union` | `users \|\| hll_hash_integer(123)` |
| `#` | `hll_cardinality` | `#users` |

### 示例：每日唯一用户跟踪

```sql
-- 存储每日唯一用户计数
CREATE TABLE daily_uniques (
    date  date UNIQUE,
    users hll
);

-- 聚合每日数据
INSERT INTO daily_uniques(date, users)
    SELECT date, hll_add_agg(hll_hash_integer(user_id))
    FROM facts GROUP BY 1;

-- 每周独立用户数（合并操作无损）
SELECT hll_cardinality(hll_union_agg(users))
FROM daily_uniques
WHERE date >= '2012-01-02' AND date <= '2012-01-08';

-- 按月统计
SELECT EXTRACT(MONTH FROM date) AS month,
       #hll_union_agg(users) AS approx_uniques
FROM daily_uniques
WHERE date >= '2012-01-01' AND date < '2013-01-01'
GROUP BY 1;

-- 7 天滑动窗口
SELECT date, #hll_union_agg(users) OVER seven_days
FROM daily_uniques
WINDOW seven_days AS (ORDER BY date ASC ROWS 6 PRECEDING);
```

### 配置参数

- **`log2m`**（4--31）：寄存器数量的以 2 为底的对数。控制精度，相对误差为 +/-1.04/sqrt(2^log2m)。默认值：11。
- **`regwidth`**（1--8）：每个寄存器的位数。与 log2m 配合调整最大基数估算。默认值：5。
- **`expthresh`**（-1 到 18）：控制 EXPLICIT 到 SPARSE 的转换。`-1` 为自动模式，`0` 跳过 EXPLICIT。默认值：-1。
- **`sparseon`**（0 或 1）：启用/禁用 SPARSE 表示。默认值：1。

同一 HLL 的所有输入必须使用相同的哈希种子。需要进行合并操作的 HLL 必须使用相同种子的哈希值填充。
