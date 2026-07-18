## Usage

Sources:

- [Official extension control file](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/git_fdw.control)
- [Official upstream documentation](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/git_fdw/)

`git_fdw` — PostgreSQL Git Foreign Data Wrapper

The reviewed catalog snapshot records version `1.1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "git_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'git_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
