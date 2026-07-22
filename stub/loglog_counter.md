## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/loglog_counter.control)
- [Official LogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/README.md)

`loglog_counter` implements the LogLog cardinality estimator as a `loglog_estimator` type, state functions, and aggregates. It is intended for approximate distinct counts where fixed memory matters more than exactness.

### Core Workflow

```sql
CREATE EXTENSION loglog_counter;

SELECT loglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

The one-argument `loglog_distinct(anyelement)` overload uses a 2.5% default error rate. The explicit overload lets the query select a different error target.

### API

- `loglog_init(real)` creates a `loglog_estimator`, and `loglog_size(real)` reports the corresponding size.
- `loglog_add_item(loglog_estimator, anyelement)` updates a state; `loglog_get_estimate(loglog_estimator)` reads the result.
- `loglog_reset(loglog_estimator)` clears the state, and `length(loglog_estimator)` reports its storage size.
- `loglog_distinct(anyelement, real)` performs configurable aggregation; `loglog_distinct(anyelement)` uses the default.

### Operational Notes

The version 1.2.4 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Estimator states may occupy several kilobytes; frequent updates to stored states produce MVCC row versions and can cause bloat. Use the lowest acceptable precision, batch updates, and measure actual error before relying on an estimate.
