## Usage

Sources:

- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/docs/risk.md)

`pgquant` computes simple/log returns, cumulative and annualized returns, and historical, Gaussian and Student-t VaR/ES. Version 0.1.0 is an early quantitative-analysis extension; volatility models and portfolio optimization remain planned.

### Core Workflow

```sql
CREATE EXTENSION pgquant;
SELECT pgquant_cumulative_simple_return(ARRAY[0.02,-0.01,0.03]::float8[]);
SELECT pgquant_var_historical(ARRAY[-0.03,0.01,-0.02,0.02]::float8[], 0.95);
SELECT pgquant_es_t(0.001, 0.015, 5.0, 0.95);
```

### Operational Boundaries

Upstream documents PostgreSQL 15–17. The control file permits non-superuser installation but does not mark the extension trusted. No preload is specified. Query-based functions require ordered symbol, date and double-precision price columns; the first observation per symbol has no return. Supply trusted SQL text. Confidence levels, distribution assumptions and input validation affect results; Student-t ES requires degrees of freedom greater than one.
