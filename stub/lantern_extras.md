## Usage

Sources:

- [Official extension control file](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/lantern_extras.control)
- [Official upstream documentation](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/README.md)
- [Official Rust package manifest](https://github.com/lanterndata/lantern/blob/7df590ce19080df7dd8a8bbb0f9875fc324bb2ea/lantern_extras/Cargo.toml)

`lantern_extras` — Convenience functions for working with vector embeddings

The reviewed catalog snapshot records version `0.6.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "lantern_extras";
SELECT extversion
FROM pg_extension
WHERE extname = 'lantern_extras';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
