## Usage

Sources:

- [Official upstream documentation](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/README.md)

`pg_substrait` — Execute Substrait protobuf or JSON query plans through PostgreSQL functions

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_substrait";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_substrait';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
