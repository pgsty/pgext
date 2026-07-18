## Usage

Sources:

- [Official extension control file](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/pg_lexo.control)
- [Official upstream documentation](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/README.md)
- [Official Rust package manifest](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/Cargo.toml)

`pg_lexo` — Provides a Base62 lexicographic position type and helper functions for stable list ordering and reordering without bulk row updates.

The reviewed catalog snapshot records version `0.6.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lexo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lexo';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
