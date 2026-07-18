## Usage

Sources:

- [Official extension control file](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/pgcuid2.control)
- [Official upstream documentation](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/README.md)

`pgcuid2` — Generate CUID2 text identifiers inside PostgreSQL using the cuid2 Rust crate.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgcuid2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcuid2';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
