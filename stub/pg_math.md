

## Usage

> [pg_math: statistical distribution functions using GSL for PostgreSQL](https://github.com/chanukyasds/pg_math)

Provides CDF (cumulative distribution) and RDF (random distribution) functions for 31 statistical distributions using the GNU Scientific Library (GSL).

```sql
CREATE EXTENSION pg_math;
```

### Supported Distributions

Gaussian, Unit Gaussian, Gaussian Tail, Bivariate Gaussian, F-Distribution, Exponential, Laplace, Exponential Power, Cauchy, Rayleigh, Rayleigh Tail, Landau, Gamma, Flat (Uniform), Lognormal, Chi-squared, T-Distribution, Beta, Logistic, Pareto, Weibull, Type-1 Gumbel, Type-2 Gumbel, Poisson, Bernoulli, Binomial, Negative Binomial, Pascal, Geometric, Hypergeometric, Logarithmic.

### Function Naming Convention

- `rdf_<distribution>(...)` -- random distribution function (PDF value)
- `cdf_<distribution>_p(...)` -- cumulative distribution P-value
- `cdf_<distribution>_q(...)` -- cumulative distribution Q-value (1-P)
- `cdf_<distribution>_pinv(...)` -- inverse CDF P
- `cdf_<distribution>_qinv(...)` -- inverse CDF Q

### Examples

```sql
-- Gaussian distribution
SELECT rdf_gaussian(1.5, 2.0);            -- PDF at x=1.5, sigma=2.0
SELECT cdf_gaussian_p(1.5, 2.0);          -- CDF P-value

-- Unit Gaussian (standard normal)
SELECT cdf_unit_gaussian_p(1.96);          -- ~0.975

-- Chi-squared distribution
SELECT cdf_chisq_p(3.84, 1.0);            -- ~0.95

-- T-distribution
SELECT cdf_tdist_pinv(0.975, 10.0);       -- critical value for 95% CI with df=10

-- Poisson distribution
SELECT rdf_poisson(5, 3.0);               -- P(X=5) with lambda=3

-- Beta distribution
SELECT rdf_beta(0.5, 2.0, 5.0);           -- PDF at x=0.5, a=2, b=5
```
