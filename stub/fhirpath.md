## Usage

Sources:

- [Official extension control file](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath.control)
- [Official upstream documentation](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/README.md)

`fhirpath` — data type for jsonb inspection

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "fhirpath";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhirpath';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
