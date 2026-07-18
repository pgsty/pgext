## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/README.md)
- [Version 0.0.1 SQL API](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router--0.0.1.sql)
- [Planner-hook implementation](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/planner/planner.c)
- [Extension control file](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router.control)

`router` is an experimental request router for OLTP sharding. Its C library installs planner behavior and exposes metadata-mutating functions for shards, sharding keys, key ranges, and local tables.

```sql
CREATE EXTENSION router;
SELECT mdbr_add_shard('shard1', '127.0.0.1', '5432', 'sharddb', 'router_user', 'replace-me');
SELECT create_sharding_key('customer_key', 'customer_id');
SELECT assign_key_range_2_shard('shard1', 'customer_key', 1, 100000);
```

The SQL file's load hint calls the project `mdb_router`, while the actual control and SQL filenames use `router`. Upstream provides a make-driven demonstration cluster rather than a supported installation and operational procedure; do not infer production semantics from the example above.

The README explicitly lists bad memory management, absent metadata consistency control, and shared-memory metadata loss after restart. Source comments also describe unfinished paths. Treat version 0.0.1 as abandoned research code, never as a production sharding layer; isolate any evaluation from real data and credentials.
