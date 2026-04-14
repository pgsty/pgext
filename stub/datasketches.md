## Usage

> Sources: [README](https://raw.githubusercontent.com/apache/datasketches-postgresql/master/README.md), [Apache DataSketches site](https://datasketches.apache.org)
> PostgreSQL extension for approximate analytics sketches and aggregates.

```sql
CREATE EXTENSION datasketches;
```

The extension supports CPC, HLL, Theta, Array Of Doubles, KLL, Quantiles, and Frequent Strings sketches.

### Sketch Families

- CPC for compact distinct counting.
- HLL for HyperLogLog-style distinct counting.
- Theta for distinct counting with set operations such as union, intersection, and A-not-B.
- Array Of Doubles for tuple sketches with arrays of double values per key.
- KLL for quantiles, ranks, PMF, and CDF estimation.
- Quantiles sketch for long-term support of distribution estimates.
- Frequent strings for tracking the heaviest items by count or weight.

### Examples

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

### Core Operations

- Build sketches with `*_sketch_build(...)`.
- Merge or aggregate them with `*_sketch_union(...)`, `*_sketch_merge(...)`, and sketch-specific set-operation helpers.
- Read estimates with `*_sketch_get_estimate(...)` and distribution helpers such as `kll_float_sketch_get_quantile(...)`.

### Notes

- The README says the extension targets PostgreSQL 9.6 and higher and depends on Boost 1.75 and DataSketches C++ core 5.0.0 or later.
- The upstream examples emphasize additive analytics in data cubes, not exact replacement for normal aggregates.
