## Usage

Sources:

- [Official extension control file](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/pg_stat_usage.control)
- [Official upstream documentation](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/README.md)

`pg_stat_usage` — track usage statistics of stored procedure calls

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_stat_usage";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_usage';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
