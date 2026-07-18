## Usage

Sources:

- [Official extension control file](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/pg_flashback.control)
- [Official upstream documentation](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/README.md)
- [Official Rust package manifest](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/Cargo.toml)

`pg_flashback` — Table-level point-in-time restore and time-travel queries for PostgreSQL using trigger or WAL capture.

The reviewed catalog snapshot records version `0.1.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_flashback";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_flashback';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
