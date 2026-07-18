## Usage

Sources:

- [Official extension control file](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/postr.control)
- [Official upstream documentation](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/README.md)
- [Official Rust package manifest](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/Cargo.toml)

`postr` — Embeds Ruby as an untrusted PostgreSQL procedural language supporting functions, DO blocks, triggers, and set-returning functions.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postr";
SELECT extversion
FROM pg_extension
WHERE extname = 'postr';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
