## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/adaptive/adaptive_counter.control)
- [Official adaptive estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/adaptive/README.md)

`adaptive_counter` provides an `adaptive_estimator` type, functions, and aggregates for approximate distinct counting with adaptive sampling. Use it when a bounded-memory estimate is acceptable; it does not replace exact `count(DISTINCT ...)` when exact results are required.

### Core Workflow

Create the extension, then select an error target and expected maximum cardinality for the aggregate:

```sql
CREATE EXTENSION adaptive_counter;

SELECT adaptive_distinct(i, 0.01, 100000)
FROM generate_series(1, 100000) AS s(i);
```

The one-argument `adaptive_distinct(anyelement)` overload uses an error target of 0.025 and an expected cardinality of 1,000,000. Pass explicit values when those defaults would allocate a larger estimator than the workload needs.

### API

- `adaptive_init(error real, ndistinct int)` creates an `adaptive_estimator`; `adaptive_size(error real, ndistinct int)` reports the corresponding size.
- `adaptive_add_item(adaptive_estimator, anyelement)` updates a state, while `adaptive_get_estimate(adaptive_estimator)` reads its current estimate.
- `adaptive_get_error(adaptive_estimator)`, `adaptive_get_ndistinct(adaptive_estimator)`, and `adaptive_get_item_size(adaptive_estimator)` expose state parameters.
- `adaptive_merge(adaptive_estimator, adaptive_estimator)` combines compatible states; `adaptive_reset(adaptive_estimator)` clears a state and `length(adaptive_estimator)` reports its storage size.

### Operational Notes

The version 1.3.3 control file marks the extension relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Estimator values can occupy several kilobytes; repeatedly updating one inside a table creates new row versions under MVCC and may cause bloat. Use the lowest acceptable precision and realistic cardinality bound, batch state updates, and validate error behavior with representative data.
