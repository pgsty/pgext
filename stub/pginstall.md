## Usage

Sources:

- [Official extension control file](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/pginstall.control)

`pginstall` — PostgreSQL extension installer that hooks CREATE EXTENSION to fetch and install extension archives.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pginstall";
SELECT extversion
FROM pg_extension
WHERE extname = 'pginstall';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
