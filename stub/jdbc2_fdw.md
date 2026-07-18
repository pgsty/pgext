## Usage

Sources:

- [Official extension control file](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/jdbc2_fdw.control)
- [Official upstream documentation](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/README.md)

`jdbc2_fdw` — Experimental foreign data wrapper for JDBC-accessible databases with remote filter pushdown.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "jdbc2_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'jdbc2_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
