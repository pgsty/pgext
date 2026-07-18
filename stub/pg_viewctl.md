## Usage

Sources:

- [Official extension control file](https://github.com/boringSQL/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/pg_viewctl.control)
- [Official upstream documentation](https://github.com/boringSQL/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/README.md)

`pg_viewctl` — View dependency management and safe schema evolution utilities

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_viewctl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_viewctl';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
