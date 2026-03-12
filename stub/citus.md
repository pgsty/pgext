

## Usage

> [citus: Distributed PostgreSQL for multi-tenant and real-time analytics](https://github.com/citusdata/citus)

Citus transforms PostgreSQL into a distributed database, enabling horizontal scaling by sharding tables across multiple nodes. It supports multi-tenant SaaS workloads, real-time analytics, and high-throughput transactional use cases while preserving the full PostgreSQL feature set.

**Key Documentation:**

- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Distributed Tables](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-distributed-table)
- [Reference Tables](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-reference-table)
- [Sharding Models](https://docs.citusdata.com/en/stable/sharding/data_modeling.html)
- [Colocation](https://docs.citusdata.com/en/stable/sharding/data_modeling.html#colocation)
- [Distributed Queries](https://docs.citusdata.com/en/stable/develop/reference_sql.html)
- [Cluster Management](https://docs.citusdata.com/en/stable/admin_guide/cluster_management.html)
- [Configuration Reference](https://docs.citusdata.com/en/stable/develop/api_guc.html)
- [Columnar Storage](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)

### Getting Started

Enable the extension and add worker nodes:

```sql
CREATE EXTENSION citus;

-- Add worker nodes to the cluster
SELECT citus_set_coordinator_host('coord-host', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

-- Verify the cluster
SELECT * FROM citus_get_active_worker_nodes();
```

### Creating Distributed Tables

Distribute a table by a chosen distribution column (shard key). Rows with the same key value are colocated on the same shard.

```sql
CREATE TABLE events (
    tenant_id   INT,
    event_id    BIGSERIAL,
    event_time  TIMESTAMPTZ DEFAULT now(),
    event_type  TEXT,
    payload     JSONB,
    PRIMARY KEY (tenant_id, event_id)
);

-- Hash-distribute by tenant_id (default: 32 shards)
SELECT create_distributed_table('events', 'tenant_id');
```

You can control the shard count:

```sql
SELECT create_distributed_table('events', 'tenant_id', shard_count := 64);
```

### Reference Tables

Small lookup tables that need to be joined with distributed tables should be created as reference tables. They are replicated in full to every node.

```sql
CREATE TABLE countries (
    code CHAR(2) PRIMARY KEY,
    name TEXT NOT NULL
);

SELECT create_reference_table('countries');
```

Reference tables can be joined with any distributed table without restrictions.

### Colocation

Tables distributed on the same column type and shard count are automatically colocated, meaning rows with matching distribution keys are stored on the same node. This enables efficient local joins.

```sql
CREATE TABLE tenants (
    id   INT PRIMARY KEY,
    name TEXT
);
SELECT create_distributed_table('tenants', 'id');

CREATE TABLE orders (
    tenant_id  INT REFERENCES tenants(id),
    order_id   BIGSERIAL,
    amount     NUMERIC,
    PRIMARY KEY (tenant_id, order_id)
);
SELECT create_distributed_table('orders', 'tenant_id');

-- This join is pushed down to each node (no cross-shard traffic)
SELECT t.name, sum(o.amount)
FROM tenants t JOIN orders o ON t.id = o.tenant_id
GROUP BY t.name;
```

You can also explicitly specify colocation groups:

```sql
SELECT create_distributed_table('orders', 'tenant_id',
    colocate_with := 'tenants');
```

### Distributed Queries

Citus pushes queries down to individual shards when possible. Queries that filter on the distribution column are routed to a single shard:

```sql
-- Single-shard query (fast, touches one node)
SELECT * FROM events WHERE tenant_id = 42;
```

Cross-shard queries are parallelized across all workers:

```sql
-- Parallel aggregation across all shards
SELECT event_type, count(*), avg(payload->>'duration')::numeric
FROM events
WHERE event_time > now() - INTERVAL '1 hour'
GROUP BY event_type
ORDER BY count DESC
LIMIT 10;
```

### Distributed Joins

Joins between colocated tables on the distribution column are executed locally on each shard:

```sql
-- Colocated join: efficient, no data movement
SELECT e.*, o.amount
FROM events e JOIN orders o
    ON e.tenant_id = o.tenant_id
WHERE e.tenant_id = 42;
```

Joins with reference tables work from any distributed table:

```sql
SELECT e.*, c.name AS country_name
FROM events e JOIN countries c ON e.payload->>'country' = c.code;
```

### Node Management

```sql
-- Add a new node
SELECT * FROM citus_add_node('worker-3', 5432);

-- Remove a node (moves shards to other nodes first)
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- Disable a node temporarily
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- View current cluster state
SELECT * FROM citus_get_active_worker_nodes();
```

### Shard Rebalancing

After adding or removing nodes, rebalance shards to distribute data evenly:

```sql
-- Rebalance all distributed tables
SELECT citus_rebalance_start();

-- Monitor rebalance progress
SELECT * FROM citus_rebalance_status();

-- Rebalance a specific table
SELECT rebalance_table_shards('events');
```

### Shard Management

```sql
-- View shard placements
SELECT * FROM citus_shards;

-- View shard sizes
SELECT table_name, shard_count, citus_total_relation_size(table_name::text)
FROM citus_tables;

-- Move a specific shard to another node
SELECT citus_move_shard_placement(shard_id, 'source-host', 5432, 'dest-host', 5432);
```

### Configuration Parameters

Key GUC parameters for tuning Citus:

```sql
-- Number of parallel connections per node for multi-shard queries
SET citus.max_adaptive_executor_pool_size = 4;

-- Shard replication factor (default 1; set to 2 for HA without streaming replication)
SET citus.shard_replication_factor = 1;

-- Control executor behavior
SET citus.multi_shard_modify_mode = 'parallel';   -- or 'sequential'
SET citus.enable_repartition_joins = on;           -- enable repartitioned joins

-- Task assignment policy
SET citus.task_assignment_policy = 'round-robin';  -- or 'greedy', 'first-replica'

-- Log distributed queries
SET citus.log_multi_join_order = on;
SET citus.explain_all_tasks = on;
```

### Example: Multi-Tenant SaaS Schema

A typical multi-tenant schema distributes all tenant-scoped tables by `tenant_id`:

```sql
CREATE TABLE tenants (
    tenant_id   INT PRIMARY KEY,
    name        TEXT,
    plan        TEXT DEFAULT 'free',
    created_at  TIMESTAMPTZ DEFAULT now()
);
SELECT create_distributed_table('tenants', 'tenant_id');

CREATE TABLE users (
    tenant_id   INT,
    user_id     BIGSERIAL,
    email       TEXT,
    PRIMARY KEY (tenant_id, user_id)
);
SELECT create_distributed_table('users', 'tenant_id');

CREATE TABLE projects (
    tenant_id   INT,
    project_id  BIGSERIAL,
    name        TEXT,
    owner_id    BIGINT,
    PRIMARY KEY (tenant_id, project_id)
);
SELECT create_distributed_table('projects', 'tenant_id');

-- Shared lookup: replicated to every node
CREATE TABLE plans (
    name TEXT PRIMARY KEY,
    max_users INT,
    max_projects INT
);
SELECT create_reference_table('plans');

-- All joins scoped to a tenant are colocated and fast
SELECT u.email, p.name AS project
FROM users u
JOIN projects p ON u.tenant_id = p.tenant_id AND u.user_id = p.owner_id
WHERE u.tenant_id = 7;
```

### Example: Real-Time Analytics

Distributed aggregation for dashboards and analytics:

```sql
CREATE TABLE page_views (
    site_id      INT,
    url          TEXT,
    view_time    TIMESTAMPTZ DEFAULT now(),
    user_agent   TEXT,
    country      CHAR(2)
);
SELECT create_distributed_table('page_views', 'site_id');

-- Real-time rollup: parallelized across shards
SELECT
    date_trunc('minute', view_time) AS minute,
    count(*)                        AS views,
    count(DISTINCT country)         AS countries
FROM page_views
WHERE site_id = 1
  AND view_time > now() - INTERVAL '1 hour'
GROUP BY minute
ORDER BY minute DESC;

-- Top pages across all sites (cross-shard parallel query)
SELECT url, count(*) AS views
FROM page_views
WHERE view_time > now() - INTERVAL '24 hours'
GROUP BY url
ORDER BY views DESC
LIMIT 20;
```
