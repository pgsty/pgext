## Usage

Sources:

- [Official extension control file](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/pg_bigmr.control)
- [Official upstream documentation](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/README.md)
- [Official Rust package manifest](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/Cargo.toml)

`pg_bigmr` — text similarity measurement and index searching based on bigrams

The reviewed catalog snapshot records version `0.1.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_bigmr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bigmr';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
