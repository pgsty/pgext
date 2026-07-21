## Usage

Sources:

- [FbSQL 0.1.0 README](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/README.md)
- [FbSQL 0.1.0 changes](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/Changes)
- [Extension control file](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/fbsql.control)
- [PGXN release](https://pgxn.org/dist/fbsql/0.1.0/)

`fbsql` is a proof-of-concept statistical-modeling DSL that keeps fitting and prediction relational: SQL queries go in and rows come back, while models are described with R formula syntax. Release 0.1.0 implements generalized linear models through PL/R for fitting and pure PL/pgSQL for prediction.

### Prerequisites

FbSQL was developed and tested with PostgreSQL 16 and requires PL/R 8.4.0 or newer plus R. `plr` is an untrusted language, so a superuser must install the dependency and extension.

```sql
CREATE EXTENSION fbsql CASCADE;
SELECT fbsql.version();
```

Grant regular users only the function access and source-data privileges they require.

### Core Workflow

Fit a binomial churn model and retain the returned relation:

```sql
CREATE TEMP TABLE churn_model AS
SELECT *
FROM fbsql.fit_glm(
  relation => $$
    SELECT churn_flag, age, gender
    FROM customer
    WHERE created_at >= DATE '2025-01-01'
      AND created_at <  DATE '2026-01-01'
  $$,
  formula => 'churn_flag ~ age + gender',
  family => 'binomial'
);
```

Prediction accepts a query for new rows and a query returning the saved model. Because it returns `SETOF record`, supply the output columns at the call site:

```sql
SELECT customer_id, churn_flag_predicted
FROM fbsql.predict_glm(
  relation => $$SELECT customer_id, age, gender FROM customer_2026$$,
  model    => $$SELECT * FROM churn_model$$
) AS p(
  customer_id bigint,
  age integer,
  gender text,
  churn_flag_predicted double precision
);
```

### Important Objects

- `fbsql.fit_glm(relation, formula, family)` returns one row per model term, repeated fit statistics, and `metadata jsonb` containing the information needed for prediction.
- `fbsql.predict_glm(relation, model, on_new_levels)` appends `<response>_predicted` to the input rows. `on_new_levels` is `error` by default or `na` to produce a null prediction for unseen factor levels.
- `fbsql.version()` reports the extension version.

### Supported Surface and Caveats

Version 0.1.0 supports Gaussian models with the identity link and binomial models with the logit link, using numeric and factor predictors. Fitting applies complete-case analysis and reports used and dropped row counts; prediction returns `NULL` when a predictor is null. Prediction uses stored coefficients and metadata and does not invoke R at runtime.

Interactions, custom contrasts, offsets, weights, prediction intervals, additional families and links, and distributed fitting are not supported. The `relation` and `model` parameters contain SQL text: construct them from trusted SQL, not unsanitized user input, and review the executing role's privileges.
