


## Usage

> [anon: Anonymization & Data Masking for PostgreSQL](https://gitlab.com/dalibo/postgresql_anonymizer/)

`postgresql_anonymizer` (extension name: `anon`) masks or replaces personally identifiable information (PII) using a declarative approach. Masking rules are defined directly in the database schema using security labels.

```sql
CREATE EXTENSION IF NOT EXISTS anon CASCADE;
SELECT anon.init();
```

### Declaring Masking Rules

```sql
SECURITY LABEL FOR anon ON COLUMN player.name
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN player.id
  IS 'MASKED WITH VALUE NULL';
```

### Static Masking

Permanently replace PII in the database:

```sql
SECURITY LABEL FOR anon ON COLUMN customer.full_name
  IS 'MASKED WITH FUNCTION anon.fake_first_name() || '' '' || anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN customer.birth
  IS 'MASKED WITH FUNCTION anon.random_date_between(''1920-01-01''::DATE, now())';

SELECT anon.anonymize_database();
-- Also available: anon.anonymize_table(), anon.anonymize_column()
```

### Dynamic Masking

Hide PII from specific roles while others see original data:

```sql
SELECT anon.start_dynamic_masking();

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';

SECURITY LABEL FOR anon ON COLUMN people.lastname
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN people.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

When `skynet` queries the table, masked data is returned automatically.

### Anonymous Dumps

```bash
pg_dump_anon.sh -h localhost -p 5432 -U bob bob_db > dump.sql
```

### Common Masking Functions

| Function | Description |
|----------|-------------|
| `anon.fake_first_name()` | Random first name |
| `anon.fake_last_name()` | Random last name |
| `anon.fake_company()` | Random company name |
| `anon.random_date_between(d1, d2)` | Random date in range |
| `anon.random_zip()` | Random zip code |
| `anon.partial(value, prefix, padding, suffix)` | Partial scrambling |
| `anon.random_string(n)` | Random string of length n |
| `anon.random_int_between(i1, i2)` | Random integer in range |
