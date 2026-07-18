## Usage

Sources:

- [Official upstream documentation](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/README.md)

`pgafis` — AFIS Support for PostgreSQL

The reviewed catalog snapshot records version `1.2`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgafis";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgafis';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
