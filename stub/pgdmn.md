## Usage

Sources:

- [Official README](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/README.md)
- [Extension control file](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/pgdmn.control)
- [pgrx manifest](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/Cargo.toml)

`pgdmn` evaluates Decision Model and Notation (DMN) models and Friendly Enough Expression Language (FEEL) expressions inside PostgreSQL.

### Enablement

Version 0.1.0 has pgrx features for PostgreSQL 14–18. Build and install for the exact major, then create the non-relocatable extension as a superuser:

```sql
CREATE EXTENSION pgdmn;
```

It does not require preloading. The extension embeds a vendored dsntk 0.3 engine; treat upgrades as application-level rule-engine changes and test saved models before rollout.

### Evaluate FEEL

`feel_eval` returns JSONB and accepts an optional JSONB context. Typed variants fail when the result cannot be represented by the requested PostgreSQL type.

```sql
SELECT feel_eval('1 + 2');
SELECT feel_eval('x * 2', '{"x":21}'::jsonb);
SELECT feel_eval_numeric('x * 2', '{"x":21}'::jsonb);
```

Typed functions include `feel_eval_text`, `feel_eval_numeric`, `feel_eval_bool`, `feel_eval_date`, `feel_eval_timestamp`, and `feel_eval_interval`.

### Load and Evaluate DMN

`dmn_load` parses DMN XML into a `dmnmodel`. `dmn_eval` resolves a named decision, business knowledge model, or decision service with JSONB inputs.

```sql
WITH model AS (
  SELECT dmn_load($dmn$<definitions>...</definitions>$dmn$) AS value
)
SELECT dmn_eval(
  value,
  'Eligibility',
  '{"Age":30,"Income":75000}'::jsonb
)
FROM model;
```

Use `dmn_invocables`, `dmn_info`, `dmn_name`, `dmn_namespace`, and `dmn_xml` to inspect a model. Validate untrusted XML and bound its size before evaluation; model parsing and rule execution consume backend CPU and memory in the calling session.

