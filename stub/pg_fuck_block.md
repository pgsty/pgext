## Usage

Sources:

- [Official extension control file](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.control)
- [Official upstream documentation](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/README)

`pg_fuck_block` — C extension for reading all tuples from a relation block

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_fuck_block";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fuck_block';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
