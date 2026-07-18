## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/index_analyzer.control)
- [Official upstream documentation](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/README.md)

`index_analyzer` — PL/pgSQL functions for evaluating existing and prospective indexes and checking foreign-key indexing.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "index_analyzer";
SELECT extversion
FROM pg_extension
WHERE extname = 'index_analyzer';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
