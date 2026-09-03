## Usage

Sources:

- [Official README](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/README.md)
- [Extension control file](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/pg_statkit.control)
- [pgrx manifest](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/Cargo.toml)

`pg_statkit` provides descriptive statistics and clinical effect-size functions over PostgreSQL arrays without exporting source data to a separate analysis process.

### Enablement

Release 1.1.0 builds and tests PostgreSQL 13–17 with pgrx 0.18.1. Install the matching native artifact, then create the non-relocatable, untrusted extension as a superuser:

```sql
CREATE EXTENSION pg_statkit;
```

It does not require preloading or a restart.

### Descriptive Statistics

Functions accept `double precision[]`; explicitly cast array literals because decimal literals otherwise form `numeric[]`.

```sql
SELECT statkit_mean('{1,2,3,4,5}'::float8[]);
SELECT statkit_median('{1,2,3,4,5}'::float8[]);
SELECT statkit_iqr('{1,2,3,4,5}'::float8[]);
SELECT statkit_percentile('{1,2,3,4,5}'::float8[], 40);

SELECT patient_id,
       statkit_median(array_agg(value::float8)) AS median_value
FROM measurements
GROUP BY patient_id;
```

The surface also includes variance, population/sample standard deviation, MAD, coefficient of variation, and standard error.

### Effect Sizes and Interpretation

`statkit_risk_difference`, `statkit_risk_ratio`, and `statkit_odds_ratio` accept a 2×2 table as four `bigint` counts. `statkit_cohens_d`, `statkit_hedges_g`, and `statkit_glass_delta` accept two samples.

```sql
SELECT statkit_odds_ratio(a => 8, b => 2, c => 2, d => 8);

SELECT statkit_cohens_d(
  array_agg(value::float8) FILTER (WHERE arm = 'treatment'),
  array_agg(value::float8) FILTER (WHERE arm = 'control')
)
FROM trial;
```

The odds-ratio function applies a 0.5 correction when any cell is zero. Undefined inputs return `NULL`. Version 1.1.0 does not provide confidence intervals, hypothesis tests, or assumption checks; use a full statistical package when those are required for research conclusions.

