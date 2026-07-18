## Usage

Sources:

- [Official extension control file](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t.control)

`plpython3t` — A SQL-only wrapper that registers PostgreSQL PL/Python 3 handlers as a trusted procedural language named plpython3t.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpython3u`.

```sql
CREATE EXTENSION "plpython3t";
SELECT extversion
FROM pg_extension
WHERE extname = 'plpython3t';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
