## 用法

> 来源: [README](https://raw.githubusercontent.com/apache/datasketches-postgresql/master/README.md), [Apache DataSketches 网站](https://datasketches.apache.org)
> PostgreSQL 扩展，用于近似分析草图与聚合。

```sql
CREATE EXTENSION datasketches;
```

该扩展支持 CPC、HLL、Theta、Array Of Doubles、KLL、Quantiles 以及 Frequent Strings 草图。

### 草图类型

- CPC 用于紧凑的去重计数。
- HLL 用于 HyperLogLog 风格的去重计数。
- Theta 用于去重计数，并支持 union、intersection 和 A-not-B 等集合运算。
- Array Of Doubles 用于每个键对应一个 double 数组的 tuple sketch。
- KLL 用于分位数、排名、PMF 和 CDF 估计。
- Quantiles 草图用于分布估计。
- Frequent strings 用于按计数或权重追踪最重项。

### 示例

```sql
SELECT cpc_sketch_to_string(cpc_sketch_build(1));
SELECT cpc_sketch_distinct(id) FROM random_ints_100m;
SELECT cpc_sketch_get_estimate(cpc_sketch_union(sketch)) FROM cpc_sketch_test;
SELECT theta_sketch_get_estimate(theta_sketch_union(sketch)) FROM theta_sketch_test;
SELECT theta_sketch_get_estimate(theta_sketch_intersection(sketch1, sketch2)) FROM theta_set_op_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(sketch)) FROM hll_sketch_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(hll_sketch_build(1), hll_sketch_build(2)));
SELECT kll_float_sketch_get_quantile(kll_float_sketch_merge(sketch), 0.5) FROM kll_float_sketch_test;
SELECT frequent_strings_sketch_result_no_false_negatives(frequent_strings_sketch_build(9, value), 1000000) FROM zipf_1p1_8k_100m;
```

### 核心操作

- 使用 `*_sketch_build(...)` 构建草图。
- 使用 `*_sketch_union(...)`、`*_sketch_merge(...)` 以及各类草图特定的集合运算辅助函数进行合并或聚合。
- 使用 `*_sketch_get_estimate(...)` 以及 `kll_float_sketch_get_quantile(...)` 等函数读取估计值。

### 说明

- README 说明该扩展面向 PostgreSQL 9.6 及以上版本，并依赖 Boost 1.75 和 DataSketches C++ core 5.0.0 或更高版本。
- 上游示例强调的是面向数据立方体的增量分析，而不是普通聚合的精确替代品。
