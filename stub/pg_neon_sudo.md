## Usage

Sources:

- [Official extension control file](https://github.com/skyzh/pg_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.control)

`pg_neon_sudo` — Provides Neon-specific trusted sudo wrappers that run PostgreSQL Anonymizer dynamic-masking start and stop operations as the bootstrap superuser.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `anon`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_neon_sudo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_neon_sudo';
```

The upstream project is associated with `Neon`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
