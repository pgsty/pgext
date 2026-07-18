## Usage

Sources:

- [Official extension control file](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/pgrx_jsonb.control)
- [Official upstream documentation](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/README.md)
- [Official Rust package manifest](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/Cargo.toml)

`pgrx_jsonb` — Proof-of-concept pgrx functions for iterating PostgreSQL JSONB internals and reconstructing JSON values

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgrx_jsonb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrx_jsonb';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
