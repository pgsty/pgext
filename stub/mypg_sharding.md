## Usage

Sources:

- [Official README](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/README.md)
- [Extension control file](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.control)
- [Extension SQL](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding--0.0.1.sql)
- [C implementation](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.c)

`mypg_sharding` is an abandoned prototype derived from Postgres Pro's pg_shardman work. The pinned 0.0.1 source does not contain an installable, internally consistent extension, so there is no safe end-user sharding workflow to document or run.

### Intended Workflow

The source intends to preload a library, create metadata in the fixed `mypg` schema, define node groups, register nodes backed by `postgres_fdw`, create hash partitions, and rebalance them. Its visible configuration is:

```conf
shared_preload_libraries = 'mypg_sharding'
mypg.is_master = on
mypg.node_name = 'node1'
```

The proposed SQL surface includes `create_nodegroup`, `register_node`, `rebalance_nodegroup`, and `create_hash_partitions`, plus internal connection, broadcast, copy, and partition-replacement helpers. These names describe unfinished intent; they are not a supported API.

### Blocking Source Defects

The installation script contains invalid PL/pgSQL syntax, creates `nodegroup`, `node`, and `tableinfo` but later refers to absent objects such as `mypg.nodeinfo`, and retains unresolved `shardman` names. It also calls pg_pathman-style partition helpers that are not declared as dependencies, while the control file lists only `postgres_fdw`. Consequently, `CREATE EXTENSION` cannot complete against the pinned source without redesign and repair.

The C code targets 2018-era server internals and has no maintained compatibility or upgrade path. Do not preload the library or expose its SQL in a production cluster. A future revival should first produce a coherent schema, declare every dependency, add deterministic install and upgrade tests, define authentication and failure semantics, and test routing, rebalance, backup, failover, and recovery on each supported PostgreSQL major.
