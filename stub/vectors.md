## Usage

Sources:

- [Official extension control file](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/vectors.control)
- [Official upstream documentation](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/README.md)
- [Official project or provider page](https://docs.vectorchord.ai/getting-started/overview.html)

`vectors` — Vector database plugin for PostgreSQL, written in Rust, specifically designed for LLM

The reviewed catalog snapshot records version `0.4.0`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "vectors";
SELECT extversion
FROM pg_extension
WHERE extname = 'vectors';
```

The upstream project is associated with `TensorChord`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
