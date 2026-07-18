## Usage

Sources:

- [Official extension control file](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/pg_emd.control)
- [Official upstream documentation](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/README.md)
- [Official Rust package manifest](https://github.com/xizhao/pg_emd/blob/653ea98a82c4b0d71c7c526ac3f9df4d0a0a16bf/Cargo.toml)

`pg_emd` — Fast approximate Earth Mover's Distance using dynamic tree embeddings

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_emd";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_emd';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
