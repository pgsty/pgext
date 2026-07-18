## Usage

Sources:

- [Official extension control file](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr.control)
- [Official upstream documentation](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/README.md)

`evr` — PostgreSQL extension for RPM epoch/version/release datatype parsing.

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "evr";
SELECT extversion
FROM pg_extension
WHERE extname = 'evr';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
