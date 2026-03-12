

## 用法

> [count_distinct: COUNT(DISTINCT ...) 的高性能替代方案](https://github.com/tvondra/count_distinct)

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
FROM generate_series(1, 10000000) s(i);

-- 替代：SELECT id, COUNT(DISTINCT val) FROM test_table GROUP BY 1;
-- 使用：
SELECT id, count_distinct(val) FROM test_table GROUP BY 1;

-- 将去重值聚合为数组
SELECT id, array_agg_distinct(val) FROM test_table GROUP BY 1;

-- 计算数组中的去重元素数量
SELECT count_distinct_elements(ARRAY[1, 2, 2, 3]);
```
