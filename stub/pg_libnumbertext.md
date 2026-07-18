## Usage

Sources:

- [Official extension control file](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/pg_libnumbertext.control)
- [Official upstream documentation](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/README.md)

`pg_libnumbertext` — Converts numbers to words in multiple languages through libnumbertext.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "pg_libnumbertext";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_libnumbertext';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
