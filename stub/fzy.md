## Usage

Sources:

- [Official upstream documentation](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/README.md)

`fzy` — FZY algorithm for PostgreSQL

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "fzy";
SELECT extversion
FROM pg_extension
WHERE extname = 'fzy';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
