## Usage

Sources:

- [Official upstream documentation](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/README.md)

`pg_reversi` — Reversi game implemented as a PostgreSQL extension.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pg_reversi";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_reversi';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
