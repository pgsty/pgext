## Usage

Sources:

- [Official extension control file](https://github.com/aquameta/pg_catalog_get_defs/blob/master/pg_catalog_get_defs.control)

`pg_catalog_get_defs` — Adds PL/pgSQL functions that reconstruct CREATE TYPE definitions from pg_catalog metadata.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pg_catalog_get_defs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_catalog_get_defs';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
