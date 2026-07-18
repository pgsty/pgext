## Usage

Sources:

- [Official extension control file](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw.control)
- [Official upstream documentation](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/README.md)

`monetdb_fdw` — MonetDB foreign data wrapper

The reviewed catalog snapshot records version `0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "monetdb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'monetdb_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
