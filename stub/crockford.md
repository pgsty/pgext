## Usage

Sources:

- [Official extension control file](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/crockford.control)
- [Official upstream documentation](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/README.markdown)

`crockford` — Crockford Base32 encoded unsigned integer data types

The reviewed catalog snapshot records version `0.8.34`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "crockford";
SELECT extversion
FROM pg_extension
WHERE extname = 'crockford';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
