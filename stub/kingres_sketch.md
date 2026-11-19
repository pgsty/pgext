## Usage

Sources:

- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/README.md)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/kingres_sketch.control)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_sketch/kingres_sketch--1.0.sql)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_sketch` 1.0 supplies mergeable HyperLogLog distinct counts, Space-Saving frequent-item summaries and t-digest percentiles. Its fixed schema avoids conflicts with other sketch extensions. The reviewed source targets PostgreSQL 18/19 and needs no preload or restart.

### Build and combine sketches

```sql
CREATE EXTENSION kingres_sketch;
SELECT kingres_sketch.approx_count_distinct(x)
FROM generate_series(1, 10000) AS t(x);
SELECT kingres_sketch.hll_cardinality(kingres_sketch.hll_add_agg(x))
FROM generate_series(1, 10000) AS t(x);
```

### Types and error bounds

Types `kingres_sketch.hll`, `kingres_sketch.topn` and `kingres_sketch.tdigest` can be stored and merged. `hll_add_agg`/`hll_union_agg`, `topn_add_agg`/`topn_union_agg` and the t-digest aggregate family support partial aggregation. Use schema-qualified names or add the extension schema after the application schema in the search path.

HLL precision is 4–18; sketches must have matching precision to merge, or first use `hll_fold` to reduce precision. Top-N capacity is 1–100000 and merge capacities must match; `topn_list` reports frequency estimates and error bounds. T-digest compression is 10–10000. These are approximate summaries, not exact counts or guaranteed percentiles. Sketches have no comparison operators/default operator classes and cannot directly be grouping or index keys. Distributed pushdown across shard groups is outside this release.
