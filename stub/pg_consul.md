## Usage

Sources:

- [Official extension control file](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/pg_consul.control)

`pg_consul` — C++ PostgreSQL extension exposing SQL functions for Consul status, catalog, and key-value APIs.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_consul";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_consul';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
