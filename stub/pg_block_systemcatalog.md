## Usage

Sources:

- [Official extension control file](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/pg_block_systemcatalog.control)

`pg_block_systemcatalog` — PostgreSQL extension to block references to system catalogs.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_block_systemcatalog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_block_systemcatalog';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
