## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/probabilistic/probabilistic_counter.control)
- [Official probabilistic estimator README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/probabilistic/README.md)

`probabilistic_counter` implements the Flajolet-Martin probabilistic counter as a `probabilistic_estimator` state type and aggregate. It trades exactness for a compact, configurable distinct-count estimate.

### Core Workflow

```sql
CREATE EXTENSION probabilistic_counter;

SELECT probabilistic_distinct(i, 4, 32)
FROM generate_series(1, 100000) AS s(i);
```

The explicit arguments are the state byte count and number of salts. The one-argument `probabilistic_distinct(anyelement)` overload uses 4 bytes and 32 salts.

### API

- `probabilistic_init(int, int)` creates a `probabilistic_estimator`; `probabilistic_size(int, int)` reports its size.
- `probabilistic_add_item(probabilistic_estimator, anyelement)` updates a state, and `probabilistic_get_estimate(probabilistic_estimator)` returns the estimate.
- `probabilistic_reset(probabilistic_estimator)` clears a state, while `length(probabilistic_estimator)` reports its storage size.
- `probabilistic_distinct(anyelement, int, int)` is the configurable aggregate; `probabilistic_distinct(anyelement)` uses the defaults.

### Operational Notes

The version 1.3.3 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Larger states can improve behavior at the cost of memory; stored states can also contribute to MVCC bloat when updated frequently. Test the estimator on representative distributions and retain exact counting for correctness-sensitive work.
