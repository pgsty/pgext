## Usage

Sources:

- [README.md](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/README.md)
- [monotonic_chi2.control](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2.control)
- [monotonic_chi2--1.0.sql](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2--1.0.sql)
- [monotonic_chi2.c](https://github.com/bojan-janjic-nz/pg_monotonic_chi2/blob/d14e656b20345d3af38ffc9198c15f16efe6d041/monotonic_chi2.c)

`monotonic_chi2` 1.0 fits chi-square bins for a numeric feature and a binary target. It merges low-support bins, applies adjacent ChiMerge, and optionally applies PAVA to make event rates monotonic. The repository is named `pg_monotonic_chi2`; the SQL extension name is `monotonic_chi2`.

### Enable and Use

The extension contains C functions and PL/pgSQL helpers. Installation requires a superuser; it is relocatable and needs no preload or restart. Upstream does not state a supported PostgreSQL major-version range.

```sql
CREATE EXTENSION monotonic_chi2;

SELECT monotonic_chi2_cutpoints(
    ARRAY[10, 20, 30, 40]::float8[],
    ARRAY[1, 5, 10, 20]::int8[],
    ARRAY[19, 15, 10, 0]::int8[],
    alpha => 0.05, max_bins => 3, min_pct => 0.05,
    monotonic => true
);

SELECT monotonic_chi2_bin(20::float8, ARRAY[10, 20, 30]::float8[]);
```

The second query assigns the value to bin 2. Boundaries are right-inclusive: each interval is `(previous_cut, cut]`, and values above the last cut go into the final bin.

### Table Helpers

For an existing table with numeric `age` and binary `defaulted` columns:

```sql
SELECT fit_chi2_bins('public.customers'::regclass, 'age', 'defaulted',
                    monotonic => false);
SELECT * FROM summarize_chi2_bins('public.customers'::regclass,
                                 'age', 'defaulted', monotonic => false);
```

`fit_chi2_bins` groups equal finite feature values before prebinning, so ties stay together. `summarize_chi2_bins` returns `bin_id`, `min_val`, `max_val`, `n_rows`, `events`, `non_events`, and `event_rate`; its optional `cutpoints` argument reuses a fitted set. Use `target_threshold` for a continuous target: values above the threshold are events. Without it, target values must be 0/1 or Boolean.

### Parameters and Boundaries

- `alpha` must be in `(0, 1)`, `max_bins` at least 1, `min_pct` in `[0, 0.5)`, and `n_prebins` at least 1. Their defaults are 0.05, 6, 0.05, and 100.
- `monotonic => false` skips PAVA, which is useful for U-shaped or multi-peaked relationships.
- Fitting excludes null targets and null, NaN, or infinite feature values. An empty fitting input returns an empty cutpoint array. Assignment with an empty array uses bin 1; null or NaN values return NULL. The summary puts null/NaN feature values in bin 0 and excludes null targets.
- Helpers read the supplied table with the caller's privileges. Choose and retain cutpoints on training data before applying them to scoring data; the extension does not validate a predictive model for you.
