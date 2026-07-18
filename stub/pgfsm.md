## Usage

Sources:

- [Official upstream documentation](https://michelp.github.io/pgfsm/)

`pgfsm` — Finite State Machine enforcement with triggers and util functions.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgfsm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfsm';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
