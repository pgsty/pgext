## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/pg_tsparser.control)
- [Official upstream documentation](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/README.md)

`pg_tsparser` — Parser for PostgreSQL text search

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_tsparser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tsparser';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
