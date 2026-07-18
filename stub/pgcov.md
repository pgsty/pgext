## Usage

Sources:

- [Official extension control file](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov.control)
- [Official upstream documentation](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/README.md)

`pgcov` — Tracks test coverage in PL/PgSQL functions and requires shared_preload_libraries.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgcov";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcov';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
