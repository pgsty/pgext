## Usage

Sources:

- [Upstream aggregate documentation](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/README.md)
- [Extension control file](https://github.com/tvondra/trimmed_aggregates/blob/bc8c1918f5d21885378cabbd070cb1b8eea02e5f/trimmed_aggregates.control)
- [PGXN distribution page](https://pgxn.org/dist/trimmed_aggregates/)

`trimmed_aggregates` computes averages, variance, and standard deviation after discarding configurable fractions from the low and high ends. The combined `trimmed` aggregate returns all seven statistics while sharing the collected state.

```sql
CREATE EXTENSION trimmed_aggregates;
SELECT avg_trimmed(v, 0.10, 0.10)
FROM generate_series(1, 1000) AS g(v);
SELECT trimmed(v, 0.20, 0.10)
FROM generate_series(1, 1000) AS g(v);
```

These aggregates retain and sort the complete input set, so memory and runtime grow with group size; many concurrent or high-cardinality groups can exhaust a backend. Validate cut fractions, NULL and small-group behavior, apply statement limits, and benchmark worst-case groups. Catalog version `2.0.0-dev` is a preview build and should be pinned and regression-tested before use.
