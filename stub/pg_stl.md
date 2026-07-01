


## Usage

Sources: [pg_ts_analysis README](https://github.com/nadyaloseva/pg_ts_analysis), [SQL definitions](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl--1.0.sql), [control file](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl.control).

`pg_stl` provides time-series analysis functions for PostgreSQL: autocorrelation, partial autocorrelation, STL decomposition, and Holt-Winters forecasting. The upstream README and SQL definitions target PostgreSQL 16+.

### Autocorrelation

`acf_array(data double precision[], lags integer)` returns autocorrelation values for lags `1..lags`:

```sql
CREATE EXTENSION pg_stl;

SELECT acf_array(
  array_agg(revenue ORDER BY date)::double precision[],
  28
)
FROM daily_sales;
```

The README describes using peaks at lags such as `7`, `14`, and `21` as a signal for weekly seasonality. The function returns `NULL` when the series is too short, `lags < 1`, or `lags >= n`.

### Partial Autocorrelation

`pacf_array(data double precision[], lags integer)` returns partial autocorrelation values using the Durbin-Levinson recursion:

```sql
WITH series AS (
  SELECT array_agg(value ORDER BY ts)::double precision[] AS values
  FROM measurements
)
SELECT
  unnest(acf_array(values, 20)) AS acf,
  unnest(pacf_array(values, 20)) AS pacf
FROM series;
```

Use PACF when you want the direct lag relationship after accounting for shorter lags.

### STL Decomposition

`stl_decompose` decomposes a series into trend, seasonal, and residual arrays:

```sql
WITH data AS (
  SELECT array_agg(revenue ORDER BY month)::double precision[] AS values
  FROM monthly_revenue
),
decomposed AS (
  SELECT (stl_decompose(values, 12)).*
  FROM data
)
SELECT
  unnest(trend) AS trend,
  unnest(seasonal) AS seasonal,
  unnest(residual) AS residual
FROM decomposed;
```

Signature from the SQL definition:

```sql
stl_decompose(
  y double precision[],
  period integer,
  seasonal integer DEFAULT 7,
  robust boolean DEFAULT true,
  trend integer DEFAULT 0,
  low_pass integer DEFAULT 0,
  inner_iter integer DEFAULT 2,
  outer_iter integer DEFAULT 0
) RETURNS stl_result
```

Use the convenience functions when only one component is needed:

```sql
SELECT stl_trend(values, 12) FROM series;
SELECT stl_seasonal(values, 12) FROM series;
SELECT stl_residual(values, 12) FROM series;
```

### Ordered Collection Helper

The SQL file also defines `stl_collect_ordered(tbl regclass, val text, ord text)` to collect a column into an ordered `double precision[]`:

```sql
SELECT stl_decompose(
  stl_collect_ordered('monthly_revenue'::regclass, 'revenue', 'month'),
  12
);
```

### Holt-Winters Forecasting

`holt_winters_predict(seasonal_type text, period_length int, start_data_array real[])` forecasts one seasonal cycle ahead. `seasonal_type` is `'mult'` for multiplicative seasonality or `'add'` for additive seasonality:

```sql
SELECT *
FROM holt_winters_predict(
  'mult',
  4,
  (SELECT array_agg(revenue ORDER BY date)::real[] FROM sales)
);
```

The SQL implementation chooses smoothing coefficients automatically: first by 500 random initializations, then by refinement in `0.001` steps to minimize squared error. The helper `holt_winters_mse(...)` is present as the error-calculation routine used by the predictor.

### Caveats

- `stl_decompose` expects a `double precision[]` with no `NULL` values.
- The README states the series length must be at least `2 * period`.
- `seasonal` must be an odd integer greater than or equal to `3`.
- Holt-Winters expects a `real[]` input and supports only `'mult'` and `'add'` seasonal types.
