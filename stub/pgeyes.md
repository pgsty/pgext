## Usage

Sources:

- [Official upstream documentation](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/README.md)

`pgeyes` — PL/pgSQL administrative helper extension with a pgeyes() environment check and foreign_db metadata table.

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgeyes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgeyes';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
