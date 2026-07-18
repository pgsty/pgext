## Usage

Sources:

- [Official extension control file](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw.control)
- [Official upstream documentation](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/README)

`sybase_fdw` — C foreign data wrapper for accessing Sybase databases from PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "sybase_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'sybase_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
