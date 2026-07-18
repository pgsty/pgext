## Usage

Sources:

- [Official upstream documentation](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/README.md)

`uints` — C extension providing unsigned 16-bit and 32-bit integer types with casts, operators, and numeric support functions.

The reviewed catalog snapshot records version `0.9`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "uints";
SELECT extversion
FROM pg_extension
WHERE extname = 'uints';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
