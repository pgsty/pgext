## Usage

Sources:

- [Official extension control file](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/json_validator_ext.control)
- [Official upstream documentation](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/README.md)
- [Official Rust package manifest](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/Cargo.toml)

`json_validator_ext` — Validate JSONB values against JSON Schema using Rust and pgrx

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "json_validator_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'json_validator_ext';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
