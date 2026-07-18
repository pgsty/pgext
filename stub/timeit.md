## Usage

Sources:

- [Official extension control file](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.control)
- [Official upstream documentation](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/README.md)

`timeit` — High-resolution benchmarking of PostgreSQL built-in functions and SQL expressions.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "timeit";
SELECT extversion
FROM pg_extension
WHERE extname = 'timeit';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
