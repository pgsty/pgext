## Usage

Sources:

- [Official extension control file](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_fdw.control)
- [Official upstream documentation](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/README.md)

`douban_fdw` — Foreign data wrapper for querying Douban public movie-ranking APIs

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `Go`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "douban_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'douban_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
