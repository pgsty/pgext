## Usage

Sources:

- [Official upstream documentation](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/README.md)

`pg_track_slow_queries` — Tracks slow queries and their execution plans

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_track_slow_queries";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_track_slow_queries';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
