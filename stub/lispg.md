## Usage

Sources:

- [Official extension control file](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg.control)

`lispg` — Lisp-like expression type and evaluator for PostgreSQL

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "lispg";
SELECT extversion
FROM pg_extension
WHERE extname = 'lispg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
