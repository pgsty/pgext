## Usage

Sources:

- [Extension control file](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/pathman_sharding.control)
- [Common SQL objects](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/init.sql)
- [Hash-sharding SQL](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/hash.sql)
- [Range-sharding SQL](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/range.sql)
- [Required postgres_fdw patch](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/postgres_fdw_execute.patch)
- [Remote DDL and index-copy implementation](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/src/pathman_sharding.c)
- [ProcessUtility hook implementation](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/src/hooks.c)

`pathman_sharding` 0.1 is a historical PostgreSQL 10-era companion to `pg_pathman`. It creates `postgres_fdw` partitions and mirrors their physical tables and parent indexes onto remote PostgreSQL servers; it does not work with stock `postgres_fdw` and is not a modern built-in partitioning solution.

### Nonstandard Prerequisites

Install `pg_pathman`, `postgres_fdw`, and `pathman_sharding`, placing `pathman_sharding` in the same schema as `pg_pathman`. The repository's patch must also be applied to PostgreSQL's bundled `postgres_fdw`; it adds `postgres_fdw_execute_custom_command(text,text)`, which the extension calls for remote DDL. A stock PostgreSQL installation does not provide that function.

Each operator needs a working user mapping for every foreign server. The patched function executes commands through the current PostgreSQL role's mapping, so installation success alone does not prove that remote creation or deletion will work.

```sql
CREATE EXTENSION pg_pathman WITH SCHEMA public;
CREATE EXTENSION postgres_fdw WITH SCHEMA public;
CREATE EXTENSION pathman_sharding WITH SCHEMA public;

-- Assumes shard_a and shard_b are configured postgres_fdw servers.
CREATE TABLE events (id bigint NOT NULL, payload jsonb);

SELECT public.create_foreign_hash_partitions(
  'events'::regclass,
  'id',
  4,
  ARRAY['shard_a', 'shard_b'],
  false
);

SELECT partition, server
FROM public.pathman_foreign_partition_list;
```

Set the `partition_data` argument to true only after validating data movement and capacity; false creates the partitions but leaves existing rows in the parent.

### Main Objects

- `create_foreign_hash_partitions` creates hash partitions and distributes their names round-robin across foreign servers by default.
- `add_foreign_range_partition` adds one foreign partition to a parent already configured for range partitioning.
- `distribute_partitions_among_servers` is the default assignment procedure and can be replaced through the hash-creation function's procedure argument.
- `pathman_foreign_partition_list` joins pg_pathman metadata to foreign-table and server metadata.

### Remote DDL, Privilege, and Compatibility Boundaries

The ProcessUtility hook reacts to creation and deletion of pg_pathman foreign partitions. Creation sends remote table DDL and then attempts to copy every parent index; index-copy errors are caught and ignored. Deletion sends remote `DROP TABLE ... CASCADE`. Always compare local and remote tables and indexes after an operation, and coordinate backups and recovery for every shard.

The SQL grants public read access to `pathman_foreign_partition_list`, including foreign-server `srvoptions`. The postgres_fdw patch creates the raw command function without an explicit privilege restriction. Revoke `EXECUTE` from `PUBLIC` on `postgres_fdw_execute_custom_command(text,text)` and on the sharding helpers, then grant only to a tightly controlled administration role.

Upstream has no README, release history, upgrade scripts, or current compatibility matrix. The source uses PostgreSQL 10 utility-hook and relation APIs and the catalog only identifies PostgreSQL 10. Do not assume it builds on newer majors. Test patched-server builds, hook coexistence, role mappings, remote DDL failures, partial index creation, routing, data movement, and teardown in an isolated cluster before considering legacy maintenance use.
