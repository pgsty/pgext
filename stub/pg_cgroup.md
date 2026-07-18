## Usage

Sources:

- [Official extension control file](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.control)
- [Official upstream documentation](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/README.md)

`pg_cgroup` — Provides SQL-callable functions to create, configure, and attach processes to Linux cgroups.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_cgroup";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cgroup';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
