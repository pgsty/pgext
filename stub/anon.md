## Usage

Sources:

- [PostgreSQL Anonymizer 3.1.3 README](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/README.md)
- [Masking functions](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/docs/masking_functions.md)
- [3.1.3 changelog](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/CHANGELOG.md)
- [Official documentation](https://postgresql-anonymizer.readthedocs.io/en/stable/)

`anon` is PostgreSQL Anonymizer. It applies declarative masking rules for protected query access, produces anonymized data sets, and provides pseudonymization and randomized-response helpers. Use it when realistic data must remain useful without exposing the original sensitive values; treat masking policy, role grants, and access to the unmasked database as part of the security boundary.

### Core Workflow

Load `anon` for sessions in the target database, install the extension, and enable transparent dynamic masking. New connections pick up database-level settings.

```sql
ALTER DATABASE app SET session_preload_libraries = 'anon';

\connect app
CREATE EXTENSION anon;
ALTER DATABASE app SET anon.transparent_dynamic_masking = true;
```

Mark a login as masked and attach masking rules to sensitive columns:

```sql
CREATE ROLE reporting LOGIN;
SECURITY LABEL FOR anon ON ROLE reporting IS 'MASKED';
GRANT pg_read_all_data TO reporting;

SECURITY LABEL FOR anon ON COLUMN customer.last_name
  IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
SECURITY LABEL FOR anon ON COLUMN customer.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

Queries made as `reporting` see the transformed values. Privileged users still see the originals, so do not grant masked roles a path around the policy.

### Masking Strategies

- **Dynamic masking** transforms results for roles labeled `MASKED` without rewriting the table.
- **Static masking** permanently rewrites selected data and is appropriate for disposable development or test copies.
- **Anonymous dumps and replicas** produce sanitized exports or downstream copies.
- **Masking views and data wrappers** expose a deliberately reduced or transformed projection.
- **Pseudonymization** uses deterministic transforms when joins or repeated values must remain consistent.

### Important Objects

- `anon.dummy_*`, `anon.random_*`, and `anon.partial(...)` generate or partially conceal values.
- `anon.hash(text)` and `anon.digest(text, text, text)` provide deterministic transformations. In 3.1.2 they were marked `RESTRICTED` to limit brute-force exposure.
- `anon.ldp_grrm(value, epsilon, max_v)` and `anon.ldp_grrm_pttt(value, truth_probability, max_v)` implement generalized randomized response for local differential privacy.
- `anon.ldp_truth_probability(...)` and `anon.ldp_lie_probability(...)` help inspect randomized-response probabilities.
- Security labels on roles and columns define who is masked and how each value is transformed.

### Operational Notes

`anon` is superuser-installed and non-relocatable. Test every policy with the same grants and connection path used by the intended consumer. Randomization is not automatically deterministic; use a confirmed pseudonymization function when stable equality is required. Static anonymization is destructive, so run it on a copy and verify constraints and application behavior afterward.

Version 3.1.3 reruns missing ARM builds and changes release metadata, with no new SQL workflow. The material delta since 3.1.1 is the 3.1.2 security hardening for `anon.hash` and `anon.digest`; deployments using those functions should upgrade rather than relying on the old labels.
