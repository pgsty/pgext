## Usage

Sources:

- [Official extension control file](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/avocado.control)
- [Official Rust package manifest](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/Cargo.toml)
- [Official project or provider page](https://avocadodb.ai)

`avocado` — AvocadoDB PostgreSQL extension for deterministic context compilation for AI agents.

The reviewed catalog snapshot records version `2.2.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "avocado";
SELECT extversion
FROM pg_extension
WHERE extname = 'avocado';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
