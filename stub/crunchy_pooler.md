## Usage

Sources:

- [Official upstream documentation](https://docs.crunchybridge.com/how-to/pgbouncer)
- [Official project or provider page](https://www.crunchydata.com/products/crunchy-bridge)

`crunchy_pooler` — Crunchy Bridge helper extension that lets PgBouncer authenticate users through PostgreSQL.

The reviewed catalog snapshot records version `unknown`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "crunchy_pooler";
SELECT extversion
FROM pg_extension
WHERE extname = 'crunchy_pooler';
```

This is a provider-specific component for `Crunchy Data`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
