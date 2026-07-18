## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/range_partitioning/)

`range_partitioning` — Manages trigger-based static range partitions for PostgreSQL-XL, including split and merge operations.

The reviewed catalog snapshot records version `1.2.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "range_partitioning";
SELECT extversion
FROM pg_extension
WHERE extname = 'range_partitioning';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
