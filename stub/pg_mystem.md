## Usage

Sources:

- [Official upstream documentation](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/README.md)

`pg_mystem` — Runs Yandex Mystem worker processes to provide Russian-language lemmatization from PostgreSQL queries.

The reviewed catalog snapshot records version `1.0.1`, kind `preload`, and implementation language `C++`.

```sql
CREATE EXTENSION "pg_mystem";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_mystem';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
