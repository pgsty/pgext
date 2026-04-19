
## Usage

> Sources: [overview](https://postgresql-anonymizer.readthedocs.io/en/stable/), [static masking](https://postgresql-anonymizer.readthedocs.io/en/stable/static_masking/), [dynamic masking](https://postgresql-anonymizer.readthedocs.io/en/stable/dynamic_masking/), [anonymous dumps](https://postgresql-anonymizer.readthedocs.io/en/stable/anonymous_dumps/), [masking functions](https://postgresql-anonymizer.readthedocs.io/en/stable/masking_functions/)

`anon` applies declarative masking rules with `SECURITY LABEL FOR anon`. The official docs center on three user-facing flows: permanent masking, masked roles, and anonymized dumps.

### Initialize and Declare Rules

```sql
CREATE EXTENSION IF NOT EXISTS anon CASCADE;
SELECT anon.init();

SECURITY LABEL FOR anon ON COLUMN customer.full_name
IS 'MASKED WITH FUNCTION anon.dummy_name()';

SECURITY LABEL FOR anon ON COLUMN customer.employer
IS 'MASKED WITH FUNCTION anon.dummy_company_name()';

SECURITY LABEL FOR anon ON COLUMN customer.phone
IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

### Static Masking

Static masking rewrites the stored data in place:

```sql
SELECT anon.anonymize_database();
-- See also: anon.anonymize_table(), anon.anonymize_column()
```

The static-masking docs also cover shuffling, noise injection, and parallel masking for larger datasets.

### Dynamic Masking

Dynamic masking hides values only from roles labeled as masked:

```sql
ALTER DATABASE demo SET session_preload_libraries = 'anon';
ALTER DATABASE demo SET anon.transparent_dynamic_masking TO true;

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';
GRANT pg_read_all_data TO skynet;

SECURITY LABEL FOR anon ON COLUMN people.lastname
IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
```

When `skynet` queries the table, masked values are returned instead of the originals.

### Anonymous Dumps and Pseudonymization

The current docs recommend transparent anonymous dumps through a masked role and `pg_dump`. Older helpers `pg_dump_anon.sh` and `pg_dump_anon` are explicitly marked deprecated.

For stable key remapping in dumps, the docs call out:

- `anon.pseudo_shift(bigint)`
- `anon.pseudo_xor(bigint)`
- `anon.set_shift()`

### Common Functions and Caveats

Common masking helpers in the function catalog include:

- `anon.dummy_first_name()`
- `anon.dummy_last_name()`
- `anon.dummy_company_name()`
- `anon.random_zip()`
- `anon.random_date_between(date, date)`
- `anon.partial(value, prefix, mask, suffix)`

Caveats from the official docs:

- dynamic masking needs preload/configuration before masked-role sessions use it
- static masking destroys the original values
- pseudonymization is not anonymization
