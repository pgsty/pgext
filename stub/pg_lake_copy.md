## Usage

Sources:

- [Official upstream documentation](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)

`pg_lake_copy` — Copy to/from data lake files

The reviewed catalog snapshot records version `3.4`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_lake_engine`, `pg_lake_iceberg`, `pg_lake_table`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lake_copy";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_copy';
```

The upstream project is associated with `Snowflake`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
