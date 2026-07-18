## Usage

Sources:

- [Official extension control file](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/pg_splade.control)
- [Official upstream documentation](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/README.md)
- [Official Rust package manifest](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/Cargo.toml)

`pg_splade` — Generate SPLADE sparse-vector embeddings inside PostgreSQL

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_splade";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_splade';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
