## Usage

Sources:

- [Official extension control file](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/pgbrtypes.control)
- [Official upstream documentation](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/README.md)

`pgbrtypes` — Brazilian CPF and CNPJ data types with validation, formatting, casts, and indexing support

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgbrtypes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgbrtypes';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
