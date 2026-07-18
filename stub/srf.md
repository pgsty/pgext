## Usage

Sources:

- [Official extension control file](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.control)

`srf` — Minimal C example implementing an inclusive integer set-returning function.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "srf";
SELECT extversion
FROM pg_extension
WHERE extname = 'srf';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
