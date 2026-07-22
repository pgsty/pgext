## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/bitmap/bitmap_counter.control)
- [Official bitmap estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/bitmap/README.md)

`bitmap_counter` implements approximate distinct counting with a self-learning bitmap. It exposes the `bitmap_estimator` state type for stored or procedural state and a `bitmap_distinct` aggregate for direct queries.

### Core Workflow

```sql
CREATE EXTENSION bitmap_counter;

SELECT bitmap_distinct(i, 0.01, 100000)
FROM generate_series(1, 100000) AS s(i);
```

The two tuning arguments are the target error and expected distinct count. The one-argument `bitmap_distinct(anyelement)` overload defaults to 0.025 and 1,000,000, so explicit values are preferable when the expected set is smaller or less precision is acceptable.

### API

- `bitmap_init(real, int)` creates a `bitmap_estimator`, and `bitmap_size(real, int)` estimates its size.
- `bitmap_add_item(bitmap_estimator, anyelement)` updates a state; `bitmap_get_estimate(bitmap_estimator)` reads the estimate.
- `bitmap_get_error(bitmap_estimator)` and `bitmap_get_ndistinct(bitmap_estimator)` expose configured state properties.
- `bitmap_reset(bitmap_estimator)` clears a state, and `length(bitmap_estimator)` reports its storage size.

### Operational Notes

The version 1.3.5 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Stored estimator states may be several kilobytes, and frequent row updates can amplify MVCC bloat. Size the state conservatively, batch updates, and treat results as estimates rather than exact cardinalities.
