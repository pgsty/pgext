## Usage

Sources:

- [Official upstream documentation](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/README.md)

`pg_keeper` — Background-worker based clustering module for monitoring streaming replication, changing synchronous replication mode and promoting a standby on failure.

The reviewed catalog snapshot records version `2.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_keeper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_keeper';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
