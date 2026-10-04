## Usage

Sources:

- [docs/volatility.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/volatility.md)
- [docs/covariance.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/covariance.md)
- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/risk.md)

`pgquant` computes simple/log returns, cumulative and annualized returns, and historical, Gaussian and Student-t VaR/ES. Version 0.1.5 adds volatility and covariance models; factor construction and portfolio optimization remain planned.

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
