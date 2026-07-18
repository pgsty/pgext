## Usage

Sources:

- [Official extension control file](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/www_fdw.control)
- [Official upstream documentation](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/www_fdw/)

`www_fdw` — Foreign data wrapper for accessing web services from PostgreSQL.

The reviewed catalog snapshot records version `0.1.9`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "www_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'www_fdw';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
