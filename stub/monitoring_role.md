## Usage

Sources:

- [Official extension control file](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/monitoring_role.control)
- [Official upstream documentation](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/README.md)

`monitoring_role` — Security-definer monitoring wrappers for a non-privileged role

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "monitoring_role";
SELECT extversion
FROM pg_extension
WHERE extname = 'monitoring_role';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
