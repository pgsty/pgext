## Usage

Sources:

- [Official extension control file](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.control)

`zlib` — zlib decompress function for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "zlib";
SELECT extversion
FROM pg_extension
WHERE extname = 'zlib';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
