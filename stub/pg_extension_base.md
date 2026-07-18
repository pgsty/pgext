## Usage

Sources:

- [Official extension control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/pg_extension_base.control)
- [Official upstream documentation](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/README.md)

`pg_extension_base` — Extension development kit by Snowflake

The reviewed catalog snapshot records version `3.4`, kind `preload`, and implementation language `C`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_extension_base";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_extension_base';
```

The upstream project is associated with `Snowflake`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
