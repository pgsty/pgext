## Usage

Sources:

- [Official trimmed_aggregates README](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/README.md)
- [Version 2.0.0-dev extension SQL](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/sql/trimmed_aggregates--2.0.0-dev.sql)
- [Version 2.0.0-dev implementation](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/trimmed_aggregates.c)

`trimmed_aggregates` computes statistics after sorting a group and discarding configurable fractions from its low and high ends. The cataloged `2.0.0-dev` SQL uses three-argument overloads of familiar aggregate names; the older `avg_trimmed` examples in the README do not describe this preview API.

### Compute trimmed statistics

Pass the value, low-end cut fraction, and high-end cut fraction:

```sql
CREATE EXTENSION trimmed_aggregates;

SELECT avg(v, 0.10, 0.10)       AS trimmed_avg,
       stddev_samp(v, 0.10, 0.10) AS trimmed_stddev
FROM generate_series(1, 1000) AS g(v);

SELECT trimmed(v, 0.20, 0.10) AS all_statistics
FROM generate_series(1, 1000) AS g(v);
```

Version `2.0.0-dev` defines `avg`, `var`, `var_pop`, `var_samp`, `stddev`, `stddev_pop`, and `stddev_samp`, each with the signature `(value, low_cut, high_cut)`. The value may be `double precision`, `int`, `bigint`, or `numeric`.

`trimmed(value, low_cut, high_cut)` returns an array of seven values in this order: average, population variance, sample variance, variance, population standard deviation, sample standard deviation, and standard deviation. For this implementation, `variance` matches the sample-variance result and `stddev` matches the sample-standard-deviation result.

### Cut rules and resource use

Each cut must be non-null and satisfy `0 <= cut < 1`; their sum must be less than `1`. Pass constants for both cuts within a group. The transition state records cut values when the group state is first created and does not provide row-by-row semantics for changing cut expressions. Null input values are skipped.

All variants retain and sort the complete non-null input set for every group. Version 2 adds combine, serialize, and deserialize functions and marks the aggregates parallel-safe, but parallel execution only distributes state; it does not remove the total memory and sort cost. High-cardinality grouping, skewed large groups, or many concurrent queries can exhaust memory.

This is a development-version API whose SQL differs materially from the README's stable examples. Pin the exact revision, qualify calls where overload resolution matters, and regression-test empty groups, small retained samples, numeric precision, parallel plans, and worst-case group sizes before use.
