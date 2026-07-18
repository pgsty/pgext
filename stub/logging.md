## Usage

Sources:

- [Official extension control file](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.control)

`logging` — Administrative C function to enable or disable logging mode on relations.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "logging";
SELECT extversion
FROM pg_extension
WHERE extname = 'logging';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
