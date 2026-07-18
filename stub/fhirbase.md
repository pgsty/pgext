## Usage

Sources:

- [Official extension control file](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/fhirbase.control)
- [Official upstream documentation](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/README.md)

`fhirbase` — Incomplete PostgreSQL extension scaffold for a FHIR database.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "fhirbase";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhirbase';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
