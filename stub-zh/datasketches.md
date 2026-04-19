## 用法

来源：[README](https://github.com/apache/datasketches-postgresql/blob/master/README.md), [latest release 1.7.0](https://github.com/apache/datasketches-postgresql/releases/tag/1.7.0), [Apache DataSketches](https://datasketches.apache.org/)

`datasketches` 为 PostgreSQL 增加近似分析 sketch 类型与聚合。上游 README 列出 CPC、HLL、Theta、Array Of Doubles、KLL、Quantiles 与 Frequent Strings sketch；GitHub 最新正式 release 为 `1.7.0`，默认分支已经推进到 `1.8.0-SNAPSHOT`。

```sql
CREATE EXTENSION datasketches;
```

### 核心 Sketch 家族

- `cpc_sketch` 与 `hll_sketch` 用于近似去重计数。
- `theta_sketch` 用于去重计数，并支持 union、intersection、A-not-B 等集合运算。
- `aod_sketch` 用于按标识符维护 double 数组形式的 tuple 指标。
- `kll_*_sketch` 与 `quantiles_*_sketch` 用于分位数、rank、PMF 与 CDF。
- `frequent_strings_sketch` 用于检测 high-hitter。

### 常见模式

从原始值构建 sketch：

```sql
SELECT cpc_sketch_build(1);
SELECT kll_float_sketch_build(value) FROM normal;
```

使用一次性近似聚合：

```sql
SELECT cpc_sketch_distinct(id) FROM random_ints_100m;
```

在分组或 cube 维度之间合并 sketch：

```sql
SELECT cpc_sketch_get_estimate(cpc_sketch_union(sketch)) FROM cpc_sketch_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(sketch)) FROM hll_sketch_test;
SELECT kll_float_sketch_get_quantile(kll_float_sketch_merge(sketch), 0.5)
FROM kll_float_sketch_test;
```

对 Theta sketch 执行集合运算：

```sql
SELECT theta_sketch_get_estimate(theta_sketch_intersection(sketch1, sketch2))
FROM theta_set_op_test;
```

查找超过阈值的高频项：

```sql
SELECT frequent_strings_sketch_result_no_false_negatives(
  frequent_strings_sketch_build(9, value),
  1000000
)
FROM zipf_1p1_8k_100m;
```

### 注意事项

- 上游文档要求 PostgreSQL 9.6+，并依赖 Boost 1.75.0 与 DataSketches C++ core 5.0.0 或更高版本。
- 这些结构是可跨维度合并的近似结构，不是 `COUNT(DISTINCT ...)` 或精确直方图的直接替代。
