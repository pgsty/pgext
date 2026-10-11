## Usage

Sources:

- [docs/optimization.md](https://github.com/pranshu05/pgQuant/blob/2e9ccbef42a5e121224831b02854005e4736f457/docs/optimization.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/2e9ccbef42a5e121224831b02854005e4736f457/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/2e9ccbef42a5e121224831b02854005e4736f457/src/lib.rs)

- [docs/factors.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/factors.md)
- [docs/volatility.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/volatility.md)
- [docs/covariance.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/covariance.md)
- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/risk.md)
- [Portfolio optimization](https://github.com/pranshu05/pgQuant/blob/cbb28e372b248cff7c28f0b989f478180b1d99e9/docs/optimization.md)

`pgquant` computes simple/log returns, cumulative and annualized returns, and historical, Gaussian and Student-t VaR/ES. Version 0.5.1 also provides volatility, covariance and factor construction. WML momentum and unconstrained mean-variance optimization are implemented; long-only optimization and Black-Litterman models are implemented.

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

### WML Momentum Factor

`pgquant_construct_wml(query)` consumes symbol TEXT, date DATE, return DOUBLE PRECISION, size_bucket INT, mom_bucket INT and weight DOUBLE PRECISION columns. Momentum buckets are 1 for losers, 2 for neutral and 3 for winners; output is date and wml_return. Build a correctly typed size/momentum panel and supply trusted SQL. Version 0.3.0 includes WML, and unconstrained portfolio optimization is available.

### Unconstrained Portfolio Optimization

`pgquant_mv_optimize_unconstrained` takes expected returns, a flattened row-major covariance matrix, a risk-free rate and positive risk aversion. It returns risky-asset weights; they may be negative and need not sum to one. The residual allocation is to the risk-free asset. This is not a long-only optimizer.

### Portfolio Optimization in 0.5.1

long-only quadratic programming adds constrained quadratic programming. Black-Litterman functions compute posterior returns/covariance, and 0.5.1 adds Idzorek view-confidence calibration; risk parity remains planned. Supply dimensionally consistent arrays and valid covariance/confidence inputs. These numerical models do not establish the suitability of a portfolio.
