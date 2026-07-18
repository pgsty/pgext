## Usage

Sources:

- [Official extension control file](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table.control)
- [Official upstream documentation](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/README.md)

`pg_force_unlogged_create_table` — force CREATE TABLE to create UNLOGGED tables

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_force_unlogged_create_table";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_force_unlogged_create_table';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
