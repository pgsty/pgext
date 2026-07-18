## Usage

Sources:

- [Official extension control file](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson.control)
- [Official upstream documentation](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/README.md)

`pg_simdjson` — Prototype SIMD-accelerated parser that converts JSON text to jsonb

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "pg_simdjson";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_simdjson';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
