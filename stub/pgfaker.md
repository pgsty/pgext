## Usage

Sources:

- [Upstream README](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/README.md)
- [Extension control file](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/pgfaker.control)
- [Rust package manifest](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/Cargo.toml)
- [SQL function implementation](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/src/lib.rs)

`pgfaker` version `0.0.3` installs English-US fake-data generators in the fixed `pgfaker` schema. It can produce company names and slogans, person names and affixes, phone numbers, domains, email addresses, and usernames. The current package targets PostgreSQL 14–18.

### Generate disposable sample rows

```sql
CREATE EXTENSION pgfaker;

SELECT pgfaker.company(),
       pgfaker.email(),
       pgfaker.phone();

SELECT g AS sample_id,
       pgfaker.person_full_name()
FROM generate_series(1, 5) AS g;
```

Results are random and the extension exposes no seed or reproducibility interface. Generated values are not guaranteed to be unique, valid, deliverable, culturally representative, privacy-safe, or suitable for constraints in an application schema. Use them for disposable demonstrations and tests, add explicit uniqueness/validation handling, and never use them as passwords, API tokens, identity proof, or a way to anonymize production data. Avoid running bulk generators in tables that contain real records unless the operation is isolated and reversible.
