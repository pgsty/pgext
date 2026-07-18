## Usage

Sources:

- [Official extension control file](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/pgparts.control)
- [Official upstream documentation](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/README.md)

`pgparts` — SQL/PLpgSQL helper extension for creating and maintaining table partitions with triggers

The reviewed catalog snapshot records version `0.0.15`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pgparts";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgparts';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
