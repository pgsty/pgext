## Usage

Sources:

- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/README.md)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/pg_react.control)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/LICENSE)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/sql/pg_react--0.46.2.sql)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/reference/release-identity.md)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/getting-started.md)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/installation.md)
- [Official source](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/operations.md)

pg_react turns relational conditions into durable rule work. The source SQL default is 0.46.2, while the officially qualified learning profile remains pg_react 0.46.0, pg_trickle 0.108.0, PostgreSQL 18.3 and Linux amd64. Cargo version 0.48.0 and the optional synthetic MDM adapter do not establish a newer qualified deployment.

### Prepare and inspect

Preload pg_trickle and pg_react in that order and restart. Use the installation guide's trigger CDC and disabled pg_trickle scheduler settings, and configure managed workers for each database with a suitable execution role. Installing SQL objects does not alone start a working deployment.

```ini
shared_preload_libraries = 'pg_trickle,pg_react'
```

```sql
SELECT extname, extversion FROM pg_extension
WHERE extname IN ('pg_react','pg_trickle') ORDER BY extname;
SELECT pgreact.doctor();
SELECT pgreact_api.managed_status();
```

### Rule workflow

Create an application view whose rows represent a condition and choose a stable semantic key. Construct pgreact.rule(), validate it, preview it, inspect findings and deploy with the reviewed token. CONSTRAINT records matches; COMMAND requires a typed consequence function. The following minimal declaration assumes the view and its order_id key already exist.

```sql
SELECT pgreact.validate(pgreact.rule(
  name => 'review-required',
  condition => 'rule_def.review_required'::regclass,
  semantic_key => 'order_id'::name
));
SELECT pgreact.run();
SELECT * FROM pgreact.matches WHERE name = 'review-required';
SELECT * FROM pgreact.work WHERE name = 'review-required';
```

### Execution boundary

Validation does not deploy the rule. Review tokens bind a plan and state but grant no permissions. Default bootstrap seeds existing matches without treating them as new work. Commit source changes so workers can observe them; pgreact.run() coordinates work rather than waiting for every consequence to finish. Database-local effects share the bookkeeping transaction; external outbox delivery is at least once and requires deduplication. Inspect pgreact.attempts and operations guidance before retries or removals.
