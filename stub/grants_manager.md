## Usage

Sources:

- [Official extension control file](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/grants_manager.control)
- [Official upstream documentation](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/README.md)

`grants_manager` — Declarative PL/pgSQL toolkit for snapshotting, reporting, and aligning database object grants.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "grants_manager";
SELECT extversion
FROM pg_extension
WHERE extname = 'grants_manager';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
