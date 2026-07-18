## Usage

Sources:

- [Official extension control file](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/fmodel_rust_postgres.control)
- [Official upstream documentation](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/README.md)
- [Official Rust package manifest](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/Cargo.toml)

`fmodel_rust_postgres` — Run f{model} event-sourced domain models inside PostgreSQL as a pgrx extension.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "fmodel_rust_postgres";
SELECT extversion
FROM pg_extension
WHERE extname = 'fmodel_rust_postgres';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
