## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/superloglog_counter.control)
- [Official SuperLogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/README.md)

`superloglog_counter` implements SuperLogLog, a LogLog variant with truncation and restriction rules, as a `superloglog_estimator` type and aggregate. It is for approximate, memory-bounded distinct counting.

### Core Workflow

```sql
CREATE EXTENSION superloglog_counter;

SELECT superloglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

The one-argument `superloglog_distinct(anyelement)` overload uses a 2.5% default error rate. Pass an explicit rate to tune memory and accuracy.

### API

- `superloglog_init(real)` creates a `superloglog_estimator`, and `superloglog_size(real)` reports its size.
- `superloglog_add_item(superloglog_estimator, anyelement)` updates a state; `superloglog_get_estimate(superloglog_estimator)` reads the estimate.
- `superloglog_reset(superloglog_estimator)` clears a state, while `length(superloglog_estimator)` reports its storage size.
- `superloglog_distinct(anyelement, real)` is configurable; `superloglog_distinct(anyelement)` uses the default error rate.

### Operational Notes

The version 1.2.3 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Estimator states may occupy several kilobytes, and repeated table updates can cause MVCC bloat. Choose the lowest acceptable precision, batch state updates, and verify estimates with real data before operational use.
