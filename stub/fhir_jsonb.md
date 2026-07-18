## Usage

Sources:

- [Official extension control file](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb.control)

`fhir_jsonb` — Synthetic FHIR JSONB data generator and query benchmark fixture for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "fhir_jsonb";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhir_jsonb';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
