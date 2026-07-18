## Usage

Sources:

- [Official extension control file](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/jsonb_opx.control)
- [Official upstream README](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/README.md)

`jsonb_opx` — hstore-style functions and operators for jsonb

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "jsonb_opx";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_opx';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
