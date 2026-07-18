## Usage

Sources:

- [Official Rust package manifest](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/Cargo.toml)
- [Official upstream README](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/README.md)

`accumulo_access_pg` — PostgreSQL extension for parsing and evaluating Accumulo Access Expressions.

The reviewed catalog snapshot records version `0.1.5`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "accumulo_access_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'accumulo_access_pg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
