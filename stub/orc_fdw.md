## Usage

Sources:

- [Official extension control file](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw.control)
- [Official upstream documentation](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/README.md)

`orc_fdw` — Foreign data wrapper for reading ORC formatted files

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "orc_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'orc_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
