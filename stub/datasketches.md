## Usage

Sources: [README](https://github.com/apache/datasketches-postgresql/blob/master/README.md), [latest release 1.7.0](https://github.com/apache/datasketches-postgresql/releases/tag/1.7.0), [Apache DataSketches](https://datasketches.apache.org/)

`datasketches` adds approximate analytics sketch types and aggregates to PostgreSQL. The upstream README lists CPC, HLL, Theta, Array Of Doubles, KLL, Quantiles, and Frequent Strings sketches; the 1.7.0 release is the latest published GitHub release, while the default branch has already moved on to `1.8.0-SNAPSHOT`.

```sql
CREATE EXTENSION datasketches;
```

### Core Sketch Families

- `cpc_sketch` and `hll_sketch` for approximate distinct counting.
- `theta_sketch` for distinct counting plus set operations such as union, intersection, and A-not-B.
- `aod_sketch` for tuple-style metrics keyed by identifiers with arrays of doubles.
- `kll_*_sketch` and `quantiles_*_sketch` for quantiles, ranks, PMF, and CDF.
- `frequent_strings_sketch` for heavy-hitter detection.

### Common Patterns

Build a sketch from raw values:

```sql
SELECT cpc_sketch_build(1);
SELECT kll_float_sketch_build(value) FROM normal;
```

Use one-shot approximate aggregates:

```sql
SELECT cpc_sketch_distinct(id) FROM random_ints_100m;
```

Merge sketches across groups or cube dimensions:

```sql
SELECT cpc_sketch_get_estimate(cpc_sketch_union(sketch)) FROM cpc_sketch_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(sketch)) FROM hll_sketch_test;
SELECT kll_float_sketch_get_quantile(kll_float_sketch_merge(sketch), 0.5)
FROM kll_float_sketch_test;
```

Run set operations on Theta sketches:

```sql
SELECT theta_sketch_get_estimate(theta_sketch_intersection(sketch1, sketch2))
FROM theta_set_op_test;
```

Find frequent items above a threshold:

```sql
SELECT frequent_strings_sketch_result_no_false_negatives(
  frequent_strings_sketch_build(9, value),
  1000000
)
FROM zipf_1p1_8k_100m;
```

### Caveats

- Upstream documents PostgreSQL 9.6+ plus Boost 1.75.0 and DataSketches C++ core 5.0.0 or later as build dependencies.
- These are approximate structures meant to be mergeable across dimensions; they are not exact replacements for `COUNT(DISTINCT ...)` or exact histograms.
