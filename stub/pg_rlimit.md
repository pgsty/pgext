## Usage

Sources:

- [Official extension control file](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit.control)
- [Official upstream documentation](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/README.md)

`pg_rlimit` — setrlimit/getrlimit functions

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_rlimit";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_rlimit';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
