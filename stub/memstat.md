## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat.control)
- [Official upstream documentation](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/README.md)

`memstat` — Memory context statistics for local and instance backends

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "memstat";
SELECT extversion
FROM pg_extension
WHERE extname = 'memstat';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
