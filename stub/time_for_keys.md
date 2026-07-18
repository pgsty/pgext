## Usage

Sources:

- [Official upstream documentation](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/README.md)

`time_for_keys` — C and PL/pgSQL helpers that enforce temporal foreign keys between valid-time tables with constraint triggers.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "time_for_keys";
SELECT extversion
FROM pg_extension
WHERE extname = 'time_for_keys';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
