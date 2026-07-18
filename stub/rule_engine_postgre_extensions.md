## Usage

Sources:

- [Official extension control file](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/rule_engine_postgre_extensions.control)
- [Official Rust package manifest](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/Cargo.toml)

`rule_engine_postgre_extensions` — Rust/pgrx rule-engine extension supporting GRL rules, RETE and forward/backward chaining, versioned rule repositories, triggers, and optional external integrations.

The reviewed catalog snapshot records version `2.0.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `pgcrypto`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rule_engine_postgre_extensions";
SELECT extversion
FROM pg_extension
WHERE extname = 'rule_engine_postgre_extensions';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
