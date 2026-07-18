## Usage

Sources:

- [Official extension control file](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/rag_bge_small_en_v15.control)
- [Official upstream documentation](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [Official Rust package manifest](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/Cargo.toml)

`rag_bge_small_en_v15` — rag_bge_small_en_v15: local bge-small-en-v1.5 embeddings in PostgreSQL

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rag_bge_small_en_v15";
SELECT extversion
FROM pg_extension
WHERE extname = 'rag_bge_small_en_v15';
```

This is a provider-specific component for `Neon`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
