## Usage

Sources:

- [Official extension control file](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/base32_4b.control)
- [Official upstream README](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/README.md)

`base32_4b` — Base32 data type stored in four bytes with comparison operators.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "base32_4b";
SELECT extversion
FROM pg_extension
WHERE extname = 'base32_4b';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
