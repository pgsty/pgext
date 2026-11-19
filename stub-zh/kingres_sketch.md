## 用法

来源：

- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/README.md)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/kingres_sketch.control)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/kingres_sketch--1.0.sql)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_sketch` 1.0 提供可合并的 HyperLogLog 去重计数、Space-Saving 高频项摘要与 t-digest 分位数。固定 schema 避免与其他 sketch 扩展冲突。所审查源码面向 PostgreSQL 18/19，无需预加载或重启。

### 构建与合并摘要

```sql
CREATE EXTENSION kingres_sketch;
SELECT kingres_sketch.approx_count_distinct(x)
FROM generate_series(1, 10000) AS t(x);
SELECT kingres_sketch.hll_cardinality(kingres_sketch.hll_add_agg(x))
FROM generate_series(1, 10000) AS t(x);
```

### 类型与误差边界

类型 `kingres_sketch.hll`、`kingres_sketch.topn` 与 `kingres_sketch.tdigest` 可以存储及合并。`hll_add_agg`/`hll_union_agg`、`topn_add_agg`/`topn_union_agg` 与 t-digest 聚合家族支持部分聚合。使用 schema 限定名，或将扩展 schema 放在应用 schema 之后的搜索路径中。

HLL 精度范围为 4–18；合并时精度必须一致，否则先用 `hll_fold` 降低精度。Top-N 容量范围为 1–100000，合并容量必须一致；`topn_list` 返回频率估计及误差界。T-digest 压缩参数范围为 10–10000。这些是近似摘要，不是精确计数或有保证的分位数。Sketch 类型没有比较运算符或默认运算符类，不能直接作为分组键或索引键。本版本不提供跨分片组的分布式下推。
