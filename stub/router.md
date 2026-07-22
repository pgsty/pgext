## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/README.md)
- [Version 0.0.1 SQL API](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router--0.0.1.sql)
- [Shard metadata implementation](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/mdbr/shard.c)
- [Planner-hook implementation](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/planner/planner.c)
- [Extension control file](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router.control)

`router` is an unfinished proof-of-concept request router for OLTP sharding. The canonical extension name is `router`, although stale comments in its SQL file say `mdb_router`. Its planner hook and in-memory metadata code are unsafe for production data or credentials; evaluate only in a disposable, isolated lab.

### Intended Interface

The SQL surface includes `mdbr_add_shard`, `mdbr_reset_meta`, `create_sharding_key`, `assign_key_range_2_shard`, `show_shkey`, and `add_local_table`. A minimal lab setup illustrates the intended model:

```sql
CREATE EXTENSION router;

SELECT mdbr_add_shard(
  'shard1', '127.0.0.1', '5432', 'sharddb', 'router_user', 'lab-only-password'
);
SELECT create_sharding_key('customer_key', 'customer_id');
SELECT assign_key_range_2_shard('shard1', 'customer_key', 1, 100000);
SELECT show_shkey('customer_key');
```

The upstream demonstration uses preload-oriented startup and a make-driven cluster harness; creating SQL objects alone does not establish a supported operating procedure. Metadata is kept in shared structures and is not durably coordinated or restored after restart.

### Critical Safety Boundaries

Shard hosts, users, and plaintext passwords are copied into fixed-size shared-memory buffers with unbounded string operations. Oversized input can corrupt memory, and debug logging can disclose connection parameters including the password. Never provide real credentials or expose the metadata functions to untrusted callers.

The code has small fixed limits for shards, sharding keys, columns, and key ranges, alongside unfinished and version-copied planner paths. The README itself lists bad memory management and missing metadata consistency. No transaction, failover, restart-recovery, schema-change, or routing-correctness guarantees can be inferred. Use `router` only to inspect historical design ideas, never as a production sharding layer.
