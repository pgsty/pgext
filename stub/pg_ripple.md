## Usage

Sources:

- [Official extension control file](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/pg_ripple.control)
- [Official upstream documentation](https://trickle-labs.github.io/pg-ripple/)
- [Official Rust package manifest](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/Cargo.toml)

`pg_ripple` — Provides an RDF knowledge-graph engine with SPARQL, SHACL, Datalog reasoning, federation, and graph analytics inside PostgreSQL 18.

The reviewed catalog snapshot records version `0.128.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ripple";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ripple';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
