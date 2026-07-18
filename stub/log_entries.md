## Usage

Sources:

- [Official extension control file](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries.control)

`log_entries` — Row trigger that records who changed a row and when

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "log_entries";
SELECT extversion
FROM pg_extension
WHERE extname = 'log_entries';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
