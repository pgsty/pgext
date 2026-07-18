## Usage

Sources:

- [Official extension control file](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp.control)

`pg_nanp` — Validated and consistently formatted North American Numbering Plan data type

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_nanp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_nanp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
