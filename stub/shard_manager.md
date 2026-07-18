## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/shard_manager/)

`shard_manager` — SQL/PLpgSQL toolkit for schema-based sharding, shard maps, template-table cloning, and 64-bit distributed ID generation.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "shard_manager";
SELECT extversion
FROM pg_extension
WHERE extname = 'shard_manager';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
