## Usage

Sources:

- [Official extension control file](https://github.com/CrunchyData/pgmonitor-extension/blob/e6859da8306cca419e16fc509d888e1da36f172e/pgmonitor.control)
- [Official upstream documentation](https://github.com/CrunchyData/pgmonitor-extension/blob/e6859da8306cca419e16fc509d888e1da36f172e/README.md)

`pgmonitor` — PostgreSQL metrics extension for external collectors, with optional background worker maintenance.

The reviewed catalog snapshot records version `2.2.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgmonitor";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmonitor';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
