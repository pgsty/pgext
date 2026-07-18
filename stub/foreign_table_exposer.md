## Usage

Sources:

- [Official extension control file](https://github.com/komamitsu/foreign_table_exposer/blob/1c57a554d1c50c872fac3b103317b11dec10e95f/foreign_table_exposer.control)
- [Official upstream documentation](https://pgxn.org/dist/foreign_table_exposer/1.0.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/foreign_table_exposer/)

`foreign_table_exposer` — Exposes foreign tables as regular tables for catalog-based clients.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "foreign_table_exposer";
SELECT extversion
FROM pg_extension
WHERE extname = 'foreign_table_exposer';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
