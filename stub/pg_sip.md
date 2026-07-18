## Usage

Sources:

- [Official extension control file](https://github.com/quentusrex/pg_sip/blob/3fb695cc3e10748e683db2ad8b60b8037fbb6c45/pg_sip.control)

`pg_sip` — Early-stage SIP extension exposing only a library version helper

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_sip";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sip';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
