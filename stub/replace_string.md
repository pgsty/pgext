## Usage

Sources:

- [Official extension control file](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/replace_string.control)
- [Official upstream documentation](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/README.md)
- [Official Rust package manifest](https://github.com/Bradyphrenia/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/Cargo.toml)

`replace_string` — Rust/pgrx function for literal string replacement

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "replace_string";
SELECT extversion
FROM pg_extension
WHERE extname = 'replace_string';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
