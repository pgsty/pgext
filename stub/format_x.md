## Usage

Sources:

- [Official extension control file](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/format_x.control)
- [Official PGXN distribution page](https://pgxn.org/dist/format_x/)
- [Official upstream README](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/README.md)

`format_x` — Polymorphic variadic string formatting with named lookups into composite, JSONB, and HSTORE values.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "format_x";
SELECT extversion
FROM pg_extension
WHERE extname = 'format_x';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
