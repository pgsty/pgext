## Usage

Sources:

- [Official extension control file](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra.control)

`jsonb_extra` — Extra JSONB functions

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "jsonb_extra";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_extra';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
