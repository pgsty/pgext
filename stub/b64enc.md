## Usage

Sources:

- [Official extension control file](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/b64enc.control)
- [Official Rust package manifest](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/Cargo.toml)
- [Official upstream README](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/README.md)

`b64enc` — Rust/C extension that defines a b64enc base type with URL-safe Base64 input and output functions.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.

```sql
CREATE EXTENSION "b64enc";
SELECT extversion
FROM pg_extension
WHERE extname = 'b64enc';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
