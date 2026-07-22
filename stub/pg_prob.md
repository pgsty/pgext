## Usage

Sources:

- [Official pg_prob README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/README.md)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/pg_prob.control)
- [End-to-end upstream example](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/example.md)

`pg_prob` version `0.2.0` adds the `pgprob.dist` probabilistic type for Monte Carlo modeling in SQL. Distributions can be constructed, combined arithmetically, aggregated across rows, sampled, and summarized without moving uncertain business data out of PostgreSQL.

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
