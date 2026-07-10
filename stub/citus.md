


## Usage

Sources:

- [Citus v14.1.0 release](https://github.com/citusdata/citus/releases/tag/v14.1.0)
- [Citus v14.1.0 CHANGELOG](https://github.com/citusdata/citus/blob/v14.1.0/CHANGELOG.md)
- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Citus Utility Functions](https://docs.citusdata.com/en/stable/develop/api_udf.html)

Citus turns PostgreSQL into a distributed database by sharding tables across worker nodes while keeping PostgreSQL SQL, indexes, extensions, transactions, and operational tooling as the user-facing surface. It is commonly used for multi-tenant SaaS databases, real-time analytics, time-series/event workloads, and distributed microservice schemas.

The local Pigsty catalog packages Citus as `citus` and exposes the lead extension `citus`; the same package also contains `citus_columnar`. Citus is a preload extension, so every node must load the library before `CREATE EXTENSION`.

### Enable Citus

```conf
shared_preload_libraries = 'citus'
```

Restart PostgreSQL on the coordinator and workers, then create the extension in the database:

```sql
CREATE EXTENSION IF NOT EXISTS citus;
SELECT citus_version();
```

On a multi-node cluster, register the coordinator and workers from the coordinator:

```sql
SELECT citus_set_coordinator_host('coord-1', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

SELECT * FROM citus_get_active_worker_nodes();
```

### Distributed Tables

Distribute a table by a shard key. Rows with the same shard-key value are colocated on the same shard, so tenant-scoped joins and point lookups stay local.

```sql
CREATE TABLE events (
  tenant_id  bigint,
  event_id   bigserial,
  event_at   timestamptz DEFAULT now(),
  kind       text,
  payload    jsonb,
  PRIMARY KEY (tenant_id, event_id)
);

SELECT create_distributed_table('events', 'tenant_id');
```

You can tune the shard count and colocation explicitly:

```sql
SELECT create_distributed_table(
  'events',
  'tenant_id',
  shard_count  := 64,
  colocate_with := 'default'
);
```

Queries that filter on the distribution column can route to a single shard:

```sql
SELECT *
FROM events
WHERE tenant_id = 42
ORDER BY event_at DESC
LIMIT 50;
```

Cross-shard queries are planned as distributed tasks and run in parallel on the workers:

```sql
SELECT kind, count(*)
FROM events
WHERE event_at >= now() - interval '1 hour'
GROUP BY kind
ORDER BY count DESC;
```

### Reference Tables

Reference tables are fully replicated to all workers. They are useful for small lookup tables that must join with many distributed tables.

```sql
CREATE TABLE countries (
  code text PRIMARY KEY,
  name text NOT NULL
);

SELECT create_reference_table('countries');
```

### Schema-Based Sharding

Schema-based sharding is useful when each tenant or service owns its own schema. In v14.1.0, Citus adds support for running several schema-sharding DDLs from any node, including `CREATE SCHEMA`, `DROP SCHEMA`, `ALTER SCHEMA RENAME`, `ALTER SCHEMA OWNER`, and table-level DDL on distributed schemas.

```sql
CREATE SCHEMA tenant_42;
SELECT citus_schema_distribute('tenant_42');

CREATE TABLE tenant_42.orders (
  id bigserial PRIMARY KEY,
  amount numeric,
  created_at timestamptz DEFAULT now()
);
```

Use row-based distribution for shared tables and schema-based sharding for per-tenant schema layouts; do not mix the two models casually without checking colocation and SQL-support implications.

### Node and Shard Operations

```sql
-- Add or disable nodes.
SELECT * FROM citus_add_node('worker-3', 5432);
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- Drain and remove a node.
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- Rebalance shards.
SELECT citus_rebalance_start();
SELECT * FROM citus_rebalance_status();
SELECT rebalance_table_shards('events');

-- Inspect tables and shards.
SELECT * FROM citus_tables;
SELECT * FROM citus_shards;
```

### Backup Coordination

Citus v14.1.0 adds UDFs for blocking distributed 2PC commit decisions and schema/topology changes while taking coordinated disk snapshots. Use them only inside a controlled backup workflow, and always unblock the cluster after the snapshot step.

```sql
SELECT citus_cluster_changes_block();
SELECT * FROM citus_cluster_changes_block_status();

-- Take coordinated filesystem or volume snapshots here.

SELECT citus_cluster_changes_unblock();
```

Pair these functions with regular PostgreSQL backup discipline: consistent checkpoints, WAL archiving, snapshot ordering across nodes, and a tested restore procedure.

### Caveats

- Pigsty local metadata currently tracks Citus 14.x for PostgreSQL 16-18; Citus 14 dropped PostgreSQL 15 support.
- `shared_preload_libraries = 'citus'` must be set before extension creation. A plain `CREATE EXTENSION citus` is not enough on a fresh server.
- Choose the distribution column carefully. Primary keys and unique constraints on distributed tables generally need to include the distribution column.
- Cross-shard joins, repartition joins, distributed DDL, and multi-shard writes are powerful but have different planning and locking behavior from single-node PostgreSQL.
- Citus includes its own columnar storage surface through `citus_columnar`; Pigsty metadata marks it as conflicting with Hydra `columnar`.
- The cluster-change blocking functions are operational tools for backups. Do not leave a cluster blocked after a failed backup script.
