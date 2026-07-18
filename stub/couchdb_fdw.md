## Usage

Sources:

- [Official extension control file](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.control)
- [Official PGXN distribution page](https://pgxn.org/dist/couchdb_fdw/)
- [Official project or provider page](https://pgxn.org/extension/couchdb_fdw)

`couchdb_fdw` — Foreign data wrapper for querying CouchDB from PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "couchdb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'couchdb_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
