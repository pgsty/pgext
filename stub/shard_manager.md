## Usage

Sources:

- [Upstream usage documentation](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/doc/shard_manager.md)
- [Version 0.0.1 SQL implementation](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/sql/shard_manager.sql)
- [Extension control file](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/shard_manager.control)

`shard_manager` is an abandoned PL/pgSQL toolkit for schema-per-shard layouts and 64-bit time/shard/sequence IDs. It records templates in `shard_table`, physical mappings in `shard_map`, and generator settings in `shard_config`; it does not route queries or create remote foreign tables.

```sql
CREATE SCHEMA shard;
CREATE EXTENSION shard_manager WITH SCHEMA shard;
SELECT shard.register_base_table('comm', 'yell', 'id');
SELECT shard.create_next_shard('comm', 'localhost');
SELECT shard.init_shard_tables('comm', 1);
```

Configure `epoch`, `shard_count`, and `ids_per_ms` before initializing any shard. The extension rounds the latter two to powers of two and refuses relevant changes after initialization because changing bit allocation can create collisions. Applications must keep `shard_map` globally consistent and perform their own routing.

Management functions are `SECURITY DEFINER` and create schemas, sequences, tables, defaults, and grants. Some dynamic SQL concatenates schema names without identifier quoting. Allow only tightly controlled identifiers and roles, audit every definer function and search path, and do not grant the administrative helpers to untrusted users.

Table cloning with `LIKE ... INCLUDING ALL` is not a complete distributed-schema or migration system. Test constraints, sequences, indexes, triggers, ownership, DDL rollout, concurrent shard creation, failover, clock behavior, ID exhaustion, backup/restore, and application routing before relying on this historical design.
