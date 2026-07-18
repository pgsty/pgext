## Usage

Sources:

- [Official extension control file](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/pg/pg_ask.control)
- [Official upstream documentation](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/README.md)

`pg_ask` — Postgres Extension which uses AI to create SQL from Natural language.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ask";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ask';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
