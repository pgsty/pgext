## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/LICENSE)
- [extensions/pg_prob/pg_prob.control](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_prob/pg_prob.control)
- [extensions/pg_prob/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_prob/Cargo.toml)
- [extensions/pg_prob/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_prob/README.md)
- [extensions/pg_prob/sql/pg_prob--0.3.0--0.3.1.sql](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_prob/sql/pg_prob--0.3.0--0.3.1.sql)
- [extensions/pg_prob/example.md](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_prob/example.md)

`pg_prob` version `0.3.1` provides the `pgprob.dist` probabilistic type for Monte Carlo modeling in SQL. Distributions can be constructed, combined arithmetically, aggregated across rows, sampled, and summarized without moving uncertain business data out of PostgreSQL.

### Core Workflow

```sql
CREATE EXTENSION pg_prob;

WITH estimate AS (
    SELECT pgprob.pert(5, 10, 20)
         + pgprob.pert(10, 15, 30) AS total_days
)
SELECT
    pgprob.summarize(total_days, 10000, 42) AS summary,
    pgprob.prob_below(total_days, 40, 10000, 42) AS finish_by_day_40
FROM estimate;
```

Constructors include `literal`, `normal`, `uniform`, `triangular`, `beta`, `lognormal`, `pert`, `poisson`, and `exponential`. Sampling and statistics include `sample`, `samples`, `summarize`, `mean`, `variance`, `stddev`, `percentile`, `prob_below`, `prob_above`, and `prob_between`. Aggregates include `sum`, `avg`, `min`, `max`, and distribution-fitting variants; correlated sampling and conditional constructors support more advanced models.

### Operational Notes

Monte Carlo answers are estimates: sample count controls both error and execution cost, and an explicit seed is required for reproducible runs. Arithmetic on `dist` propagates modeled uncertainty, but model assumptions and correlations remain the caller's responsibility. The extension is fixed to the `pgprob` schema, is non-relocatable, and is not superuser-only. Before using results for financial or operational decisions, validate parameterization against observed data and benchmark large sample counts under the target workload.

### Version 0.3.0 Boundary

This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Changes in This Release

Source feature declarations cover PostgreSQL 14–18. The repository release is 0.7.0, while each extension retains its own SQL version; this does not establish Pigsty package availability.

Version 0.3.1 marks random sampling and statistics functions `VOLATILE`, preventing an unseeded constant-argument call from reusing one draw throughout a query. Explicit seeds remain reproducible. `pert` now uses standard PERT Beta parameters when the mode is the interval midpoint. After installing matching libraries, run `ALTER EXTENSION pg_prob UPDATE TO '0.3.1'` on 0.3.0 databases so existing SQL functions receive the new volatility markings.
