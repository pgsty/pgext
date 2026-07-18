## Usage

Sources:

- [Official extension control file](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp.control)
- [Official upstream documentation](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/README.md)

`pgseccomp` — provide seccomp syscall filtering for PostgreSQL

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pgseccomp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgseccomp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
