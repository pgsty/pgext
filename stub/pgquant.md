## Usage

Sources:

- [docs/factors.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/factors.md)
- [docs/volatility.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/volatility.md)
- [docs/covariance.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/covariance.md)
- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/risk.md)

`pgquant` computes simple/log returns, cumulative and annualized returns, and historical, Gaussian and Student-t VaR/ES. Version 0.2.0 also provides volatility, covariance and factor construction. Momentum and portfolio optimization remain planned.

### Core Workflow

```sql
CREATE EXTENSION pgquant;
SELECT pgquant_cumulative_simple_return(ARRAY[0.02,-0.01,0.03]::float8[]);
SELECT pgquant_var_historical(ARRAY[-0.03,0.01,-0.02,0.02]::float8[], 0.95);
SELECT pgquant_es_t(0.001, 0.015, 5.0, 0.95);
```

### Operational Boundaries

Upstream documents PostgreSQL 14–17. The control file permits non-superuser installation but does not mark the extension trusted. No preload is specified. Query-based functions require ordered symbol, date and double-precision price columns; the first observation per symbol has no return. Supply trusted SQL text. Confidence levels, distribution assumptions and input validation affect results; Student-t ES requires degrees of freedom greater than one.

### Volatility and Covariance

`pgquant_rolling_vol` uses an ordered returns query; `pgquant_ewma_vol` returns an array and `pgquant_ewma_lambda_mle` estimates its decay parameter. `pgquant_garch11_normal` and `pgquant_garch11_t` fit GARCH parameters and require at least ten observations. `pgquant_sample_cov` and `pgquant_shrinkage_cov_lw` return asset-pair covariance from complete shared dates, taking a query with symbol, date and double-precision return columns. Input ordering, missing data and distribution assumptions affect the estimates.

```sql
SELECT pgquant_ewma_vol(ARRAY[0.02,-0.01,0.03,-0.02]::float8[], 0.94);
```

### Factor Construction

`pgquant_portfolio_sort(query, num_buckets)` takes exactly symbol TEXT, date DATE and characteristic_value DOUBLE PRECISION columns and returns symbol, date and bucket; bucket 1 contains the lowest values. `pgquant_construct_smb(query)` and `pgquant_construct_hml(query)` consume symbol, date, return DOUBLE PRECISION, size_bucket INT, btm_bucket INT and weight DOUBLE PRECISION. Size buckets are 1/2; book-to-market buckets are 1/2/3. Returns are date with smb_return or hml_return. Prepare a correctly typed, aligned panel and supply trusted SQL text.

```sql
SELECT * FROM pgquant_portfolio_sort(
  $$ SELECT * FROM (VALUES
     ('A'::text, DATE '2026-01-01', 100.0::float8),
     ('B'::text, DATE '2026-01-01', 200.0::float8)
  ) AS panel(symbol, date, characteristic_value) $$, 2);
```
