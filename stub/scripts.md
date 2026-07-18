## Usage

Sources:

- [Official upstream documentation](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/README.md)

`scripts` — PhaseZero aftermarket application schema packaged as a PostgreSQL extension.

The reviewed catalog snapshot records version `1.4`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "scripts";
SELECT extversion
FROM pg_extension
WHERE extname = 'scripts';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
