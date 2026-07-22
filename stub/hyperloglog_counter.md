## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/hyperloglog/hyperloglog_counter.control)
- [Official HyperLogLog README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/hyperloglog/README.md)

`hyperloglog_counter` provides approximate distinct counting with the HyperLogLog algorithm. It supplies a `hyperloglog_estimator` state type plus aggregate and state-management functions; use it for memory-bounded cardinality estimates, not exact counts.

### Core Workflow

```sql
CREATE EXTENSION hyperloglog_counter;

SELECT hyperloglog_distinct(i, 0.01)
FROM generate_series(1, 100000) AS s(i);
```

The one-argument `hyperloglog_distinct(anyelement)` overload uses a 2% default error rate. Pass an explicit error rate to control the memory-versus-accuracy tradeoff.

### API

- `hyperloglog_init(real)` creates a `hyperloglog_estimator`; `hyperloglog_size(real)` reports the required size for an error rate.
- `hyperloglog_add_item(hyperloglog_estimator, anyelement)` updates a state, and `hyperloglog_get_estimate(hyperloglog_estimator)` reads its estimate.
- `hyperloglog_reset(hyperloglog_estimator)` clears a state, while `length(hyperloglog_estimator)` reports its storage size.
- `hyperloglog_distinct(anyelement, real)` is the configurable aggregate; `hyperloglog_distinct(anyelement)` uses the default error rate.

### Operational Notes

The version 1.2.6 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. A stored estimator can consume several kilobytes, and repeated updates can create MVCC bloat. Choose the least memory-intensive acceptable error rate and validate estimates on the workload's real cardinality distribution.
