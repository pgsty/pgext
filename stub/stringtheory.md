## Usage

Sources:

- [Official upstream documentation](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/README.md)

`stringtheory` — C++ extension providing SIMD-optimized exact equality and substring-search functions for PostgreSQL text values.

The reviewed catalog snapshot records version `1.0.2`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "stringtheory";
SELECT extversion
FROM pg_extension
WHERE extname = 'stringtheory';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
