## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/pg_backtrace.control)
- [Official upstream documentation](https://github.com/postgrespro/pg_backtrace/blob/de245a5f6b42ba870259ba68b4fbdd0c8b7bae69/README)

`pg_backtrace` — Show backtrace for errors and signals

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_backtrace";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_backtrace';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
