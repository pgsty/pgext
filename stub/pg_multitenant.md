## Usage

Sources:

- [Official upstream documentation](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/README.md)
- [Official Rust package manifest](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/Cargo.toml)

`pg_multitenant` — Multi-tenant shared-schema helpers using PostgreSQL row-level security

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_multitenant";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_multitenant';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
