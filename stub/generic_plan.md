## Usage

Sources:

- [Official extension control file](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/generic_plan.control)
- [Official upstream documentation](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/README.md)

`generic_plan` — Generate a generic plan for a parameterized SQL statement.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "generic_plan";
SELECT extversion
FROM pg_extension
WHERE extname = 'generic_plan';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
