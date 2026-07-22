## Usage

Sources:

- [Official control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/pcsa/pcsa_counter.control)
- [Official PCSA README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/pcsa/README.md)

`pcsa_counter` implements probabilistic counting with stochastic averaging as a `pcsa_estimator` type and aggregate. It provides approximate distinct counts with configurable bitmap count and key size.

### Core Workflow

```sql
CREATE EXTENSION pcsa_counter;

SELECT pcsa_distinct(i, 32, 3)
FROM generate_series(1, 100000) AS s(i);
```

The arguments select the number of bitmaps and key size. The one-argument `pcsa_distinct(anyelement)` overload defaults to 64 bitmaps and key size 4; tune explicitly when a smaller state or different accuracy is appropriate.

### API

- `pcsa_init(int, int)` creates a `pcsa_estimator`; `pcsa_size(int, int)` reports the state size.
- `pcsa_add_item(pcsa_estimator, anyelement)` updates a state, and `pcsa_get_estimate(pcsa_estimator)` reads its estimate.
- `pcsa_reset(pcsa_estimator)` clears a state, while `length(pcsa_estimator)` reports its storage size.
- `pcsa_distinct(anyelement, int, int)` is the configurable aggregate; `pcsa_distinct(anyelement)` uses the documented defaults.

### Operational Notes

The version 1.3.3 control file is relocatable and declares no prerequisite extension or preload requirement. The upstream repository is archived and read-only. Larger configurations consume more memory, and repeatedly updating a stored estimator can cause MVCC bloat. Benchmark the selected parameters against representative cardinalities and do not treat estimates as exact results.
