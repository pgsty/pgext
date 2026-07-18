## Usage

Sources:

- [Official extension control file](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp.control)

`pg_udp` — Send UDP datagrams from a PostgreSQL function.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_udp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_udp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
