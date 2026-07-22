## Usage

Sources:

- [Official pg_shardman README](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/readme.md)
- [Version 1.0 extension SQL](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman--1.0.sql)
- [Extension control file](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman.control)

`pg_shardman` is an experimental PostgreSQL 10 cluster manager for hash-sharded tables built on `pg_pathman`, `postgres_fdw`, and logical replication. The project says it does not fully reach its goals; several important workflows require patched PostgreSQL and patched dependencies, so it should be treated as historical lab software rather than a supported sharding solution.

### Cluster outline

Install the extension and its dependencies on every server. Add `pg_pathman` to `shared_preload_libraries` and restart each instance. Designate one metadata-only shardlord and configure all nodes:

```conf
shardman.shardlord = on
shardman.shardlord_connstring = 'host=shardlord dbname=app'
shardman.sync_replication = off
```

Set `shardman.shardlord = off` on workers. The shardlord currently cannot also store normal application data. Create `pg_shardman` on all nodes, then run cluster-management calls on the shardlord:

```sql
SELECT shardman.add_node(
    'host=worker1 dbname=app user=postgres',
    'host=worker1 dbname=app user=app',
    'rack_a'
);

CREATE TABLE public.events (
    event_id bigint NOT NULL,
    payload  jsonb
);

SELECT shardman.create_hash_partitions(
    'public.events'::regclass,
    'event_id',
    32,
    0
);
```

The initial partition count cannot be changed later, so choose it for future node growth. Replication groups constrain where replicas may be placed. The default group creates all-to-all logical replication channels—even at redundancy zero—so use deliberate group names when replication is not wanted.

### Management surface

- `shardman.create_shared_table` replicates a small reference table from one master node.
- `shardman.mv_partition`, `shardman.mv_replica`, `shardman.rebalance`, and `shardman.rebalance_replicas` move or redistribute data.
- `shardman.set_redundancy` increases a table's replica target; `shardman.ensure_redundancy` fills missing copies.
- `shardman.nodes`, `shardman.tables`, `shardman.partitions`, `shardman.replicas`, and `shardman.replication_lag` expose metadata and lag.
- `shardman.forall` broadcasts DDL, while `shardman.recovery` attempts to reconcile nodes with shardlord metadata.

### Critical limitations

Vanilla PostgreSQL 10 cannot `COPY FROM` the foreign partitions used here; the documented path requires patched PostgreSQL and `pg_pathman`. Sane cross-node transactions require a patched `postgres_fdw` with two-phase commit and distributed snapshots. Sharded-table `ALTER TABLE`, global primary keys, and foreign keys to sharded tables are mostly unsupported, and cross-node OLAP does not get parallel foreign scans.

Failure detection and recovery are manual. With multiple replicas, promotion is random and divergent replica states are not reconciled. `shardman.rm_node(..., true)` can discard data, and `shardman.rm_table` has no confirmation prompt. Shared-table transactions may not read their own replicated writes. Before any experiment, read the full upstream limitations, isolate the cluster, and rehearse node loss, shardlord recovery, and data reconciliation.
