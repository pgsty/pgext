## Usage

Sources:

- [Official upstream documentation](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/readme.md)

`town` — Minimal PL/pgSQL helper that creates JSONB-and-array time-series tables with GIN and GiST indexes.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `btree_gist`, `plpgsql`.

```sql
CREATE EXTENSION "town";
SELECT extversion
FROM pg_extension
WHERE extname = 'town';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
