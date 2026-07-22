## Usage

Sources:

- [Extension control file](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/conda_pgsql_rust_ext.control)
- [Rust implementation](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/src/lib.rs)
- [Upstream regression examples](https://github.com/JeanChristopheMorinPerso/data_experiments/tree/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/tests/pg_regress/sql)

`conda_pgsql_rust_ext` provides a `condaversion` type backed by `rattler_conda_types`. Use it when Conda package versions must be parsed, compared, indexed, or aggregated according to Conda version ordering instead of ordinary text ordering.

### Core Workflow

Create the extension, store versions in a typed column, and use the normal comparison and aggregate syntax:

```sql
CREATE EXTENSION conda_pgsql_rust_ext;

CREATE TABLE package_release (
    package text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO package_release VALUES
    ('demo', '1.1dev1'),
    ('demo', '1.1.0'),
    ('demo', '1!0.4.1');

SELECT version::text
FROM package_release
WHERE package = 'demo'
ORDER BY version;

SELECT min(version), max(version)
FROM package_release;
```

The type accepts Conda version strings, retains the original spelling for text output, and supplies equality, ordering, hashing, plus `min(condaversion)` and `max(condaversion)` aggregates. Invalid Conda syntax raises an input error.

### Operational Notes

The reviewed upstream revision reports crate version `0.0.0`. The control file marks the extension non-relocatable, untrusted, and superuser-only; it does not declare a preload requirement. Treat the early version number as an API-stability warning, and test comparison behavior with the exact version forms used by your package metadata before adopting it for constraints or unique indexes.
