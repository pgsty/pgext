## Usage

Sources:

- [Official upstream documentation](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/README.md)

`units` — Unit conversion table and PL/pgSQL conversion function for length units.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "units";
SELECT extversion
FROM pg_extension
WHERE extname = 'units';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
