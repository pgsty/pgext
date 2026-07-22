## Usage

Sources:

- [Official README](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/README.md)
- [Version 0.1.0 SQL declarations](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/sql/pg_emd--0.1.0.sql)
- [PostgreSQL-facing Rust implementation](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/src/pg_extension.rs)
- [Cargo and PostgreSQL compatibility](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/Cargo.toml)

`pg_emd` 0.1.0 computes Earth Mover's Distance. The current `emd` implementation is an exact O(n) calculation for one-dimensional histograms whose array positions are bins. `emd_weighted` uses a randomized dynamic tree embedding for approximate distance between weighted multidimensional points. The extension is a research implementation, currently built only for PostgreSQL 17.

### One-Dimensional Histograms

```sql
CREATE EXTENSION pg_emd;

SELECT emd(
  ARRAY[0.5, 0.3, 0.2]::double precision[],
  ARRAY[0.2, 0.3, 0.5]::double precision[]
);
```

Both arrays must have the same length; an empty pair returns zero. Treat each value as mass at bin positions `0, 1, ...`. Validate that values are finite and nonnegative and that total masses are equal before calling, because the function does not enforce those distribution invariants.

### Weighted Multidimensional Points

```sql
SELECT emd_weighted(
  '[{"point":[1.0,2.0],"weight":0.7},
    {"point":[3.0,4.0],"weight":0.3}]'::json,
  '[{"point":[10.0,20.0],"weight":1.0}]'::json
);

SELECT tree_distance(
  ARRAY[0.0, 0.0]::double precision[],
  ARRAY[10.0, 10.0]::double precision[]
);
```

`tree_distance` exposes the underlying randomized tree metric for inspection; it is not exact Euclidean distance.

### Critical Planner and Input Caveats

The SQL declares all three functions `IMMUTABLE PARALLEL SAFE`, but `emd_weighted` and `tree_distance` construct tree embeddings with random grid offsets on each call. Their results can vary for identical arguments, so the volatility declaration is unsafe. Do not use those two functions in expression indexes, generated columns, partition bounds, cached constant expressions, or constraints unless the implementation and SQL declarations are corrected.

`emd_weighted` derives dimension from the first point and does not fully validate equal dimensions, finite coordinates, nonnegative weights, or equal total mass; nonnumeric coordinate entries are filtered during conversion. Validate and normalize JSON in the caller. Empty weighted input on either side returns zero, which may not match the application's missing-data semantics.

The README's performance and approximation figures are project measurements, not workload guarantees. Pin the build, validate numerical error and repeatability against an exact implementation, and test memory and latency with production distribution sizes before adoption.
