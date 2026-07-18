## Usage

Sources:

- [Official extension control file](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/pg_ivfflat.control)
- [Official upstream documentation](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/README.md)
- [Official Rust package manifest](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/Cargo.toml)

`pg_ivfflat` — Experimental pgrx extension providing a Vector type and an in-memory IVFFlat index access method.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ivfflat";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ivfflat';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
