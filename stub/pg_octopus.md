## Usage

Sources:

- [Official extension control file](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/pg_octopus.control)

`pg_octopus` — Runs a background worker that periodically connects to configured PostgreSQL nodes and records their health status.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_octopus";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_octopus';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
