## 用法

来源：

- [v3.0.2 README](https://github.com/tvondra/count_distinct/blob/v3.0.2/README.md)
- [v3.0.2 aggregate definitions](https://github.com/tvondra/count_distinct/blob/v3.0.2/sql/count_distinct--3.0.2.sql)
- [Control file](https://github.com/tvondra/count_distinct/blob/v3.0.2/count_distinct.control)

提供 `COUNT(DISTINCT ...)` 的替代实现，避免排序并支持并行聚合。

```sql
CREATE EXTENSION count_distinct;
```

### 函数

| 函数 | 描述 |
|---|---|
| `count_distinct(value anyelement)` | 计算去重计数（`COUNT(DISTINCT ...)` 的替代方案） |
| `array_agg_distinct(value anyelement)` | 将去重值聚合为数组 |
| `count_distinct_elements(value anyarray)` | 计算输入数组中去重元素的数量 |
| `array_agg_distinct_elements(value anyarray)` | 将输入数组中的去重元素聚合为数组 |

### 示例

```sql
CREATE TABLE test_table (id INT, val INT);
INSERT INTO test_table
SELECT mod(i, 1000), (1000 * random())::int
FROM generate_series(1, 10000) s(i);

-- Instead of:  SELECT id, COUNT(DISTINCT val) FROM test_table GROUP BY 1;
-- Use:
SELECT id, count_distinct(val) FROM test_table GROUP BY 1;

-- Aggregate distinct values into an array
SELECT id, array_agg_distinct(val) FROM test_table GROUP BY 1;

-- Count distinct elements across arrays
SELECT count_distinct_elements(ARRAY[1, 2, 2, 3]);
```

### 类型、内存与版本

3.0.2 接受按值传递的定长类型，例如整数，以及这些类型的数组；它不能普遍替代文本值的去重聚合。先对大值做哈希可能发生碰撞，因此会把结果变成估计值。这些聚合将状态保留在内存中，无法可靠执行 `work_mem` 限制；高基数或大量并发分组可能耗尽内存。应使用实际负载与内置聚合比较正确性、执行计划和内存，而不是预设它一定更快。控制文件允许重定位且无需预加载；安装 C 扩展通常需要超级用户。
