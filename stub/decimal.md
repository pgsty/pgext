## Usage

Sources:

- [Official extension control file](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.control)
- [Official upstream documentation](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/pgdecimal/)

`decimal` — Decimal32 and decimal64 decimal floating-point data types with arithmetic and casts

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "decimal";
SELECT extversion
FROM pg_extension
WHERE extname = 'decimal';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
