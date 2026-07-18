## Usage

Sources:

- [Official extension control file](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/nextid.control)
- [Official upstream documentation](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/README.md)

`nextid` — Generate Instagram-style sharded integer IDs.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "nextid";
SELECT extversion
FROM pg_extension
WHERE extname = 'nextid';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
