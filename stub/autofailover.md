## Usage

Sources:

- [Official extension control file](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover.control)

`autofailover` — C extension exposing autofailover functions for role and WAL status checks, promotion, and synchronous replication setting changes.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "autofailover";
SELECT extversion
FROM pg_extension
WHERE extname = 'autofailover';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
