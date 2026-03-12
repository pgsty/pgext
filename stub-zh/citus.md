

## 用法

> [citus: 面向多租户和实时分析的分布式 PostgreSQL](https://github.com/citusdata/citus)

Citus 将 PostgreSQL 转变为分布式数据库，通过在多个节点间分片表来实现水平扩展。它支持多租户 SaaS 工作负载、实时分析和高吞吐量事务场景，同时保留完整的 PostgreSQL 功能集。

**核心文档：**

- [什么是 Citus？](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [分布式表](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-distributed-table)
- [引用表](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-reference-table)
- [分片模型](https://docs.citusdata.com/en/stable/sharding/data_modeling.html)
- [协同定位](https://docs.citusdata.com/en/stable/sharding/data_modeling.html#colocation)
- [分布式查询](https://docs.citusdata.com/en/stable/develop/reference_sql.html)
- [集群管理](https://docs.citusdata.com/en/stable/admin_guide/cluster_management.html)
- [配置参考](https://docs.citusdata.com/en/stable/develop/api_guc.html)
- [列式存储](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)

### 快速开始

启用扩展并添加工作节点：

```sql
CREATE EXTENSION citus;

-- 将协调节点添加到集群
SELECT citus_set_coordinator_host('coord-host', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

-- 验证集群状态
SELECT * FROM citus_get_active_worker_nodes();
```

### 创建分布式表

通过指定分布列（分片键）来分布表。具有相同键值的行会被协同定位在同一个分片上。

```sql
CREATE TABLE events (
    tenant_id   INT,
    event_id    BIGSERIAL,
    event_time  TIMESTAMPTZ DEFAULT now(),
    event_type  TEXT,
    payload     JSONB,
    PRIMARY KEY (tenant_id, event_id)
);

-- 按 tenant_id 进行哈希分布（默认：32 个分片）
SELECT create_distributed_table('events', 'tenant_id');
```

可以控制分片数量：

```sql
SELECT create_distributed_table('events', 'tenant_id', shard_count := 64);
```

### 引用表

需要与分布式表进行连接的小型查找表应创建为引用表。引用表会被完整复制到每个节点。

```sql
CREATE TABLE countries (
    code CHAR(2) PRIMARY KEY,
    name TEXT NOT NULL
);

SELECT create_reference_table('countries');
```

引用表可以与任何分布式表进行无限制的连接。

### 协同定位

基于相同列类型和分片数量进行分布的表会自动协同定位，即具有匹配分布键的行存储在同一节点上。这使得高效的本地连接成为可能。

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

-- 此连接被下推到各节点执行（无跨分片流量）
SELECT t.name, sum(o.amount)
FROM tenants t JOIN orders o ON t.id = o.tenant_id
GROUP BY t.name;
```

也可以显式指定协同定位组：

```sql
SELECT create_distributed_table('orders', 'tenant_id',
    colocate_with := 'tenants');
```

### 分布式查询

Citus 尽可能将查询下推到各个分片。过滤分布列的查询会被路由到单个分片：

```sql
-- 单分片查询（快速，只访问一个节点）
SELECT * FROM events WHERE tenant_id = 42;
```

跨分片查询会在所有工作节点上并行执行：

```sql
-- 跨所有分片的并行聚合
SELECT event_type, count(*), avg(payload->>'duration')::numeric
FROM events
WHERE event_time > now() - INTERVAL '1 hour'
GROUP BY event_type
ORDER BY count DESC
LIMIT 10;
```

### 分布式连接

协同定位表之间基于分布列的连接在每个分片上本地执行：

```sql
-- 协同定位连接：高效，无数据移动
SELECT e.*, o.amount
FROM events e JOIN orders o
    ON e.tenant_id = o.tenant_id
WHERE e.tenant_id = 42;
```

与引用表的连接可从任何分布式表发起：

```sql
SELECT e.*, c.name AS country_name
FROM events e JOIN countries c ON e.payload->>'country' = c.code;
```

### 节点管理

```sql
-- 添加新节点
SELECT * FROM citus_add_node('worker-3', 5432);

-- 移除节点（先将分片迁移到其他节点）
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- 临时禁用节点
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- 查看当前集群状态
SELECT * FROM citus_get_active_worker_nodes();
```

### 分片再平衡

添加或移除节点后，通过再平衡分片来均匀分配数据：

```sql
-- 再平衡所有分布式表
SELECT citus_rebalance_start();

-- 监控再平衡进度
SELECT * FROM citus_rebalance_status();

-- 再平衡特定表
SELECT rebalance_table_shards('events');
```

### 分片管理

```sql
-- 查看分片分布
SELECT * FROM citus_shards;

-- 查看分片大小
SELECT table_name, shard_count, citus_total_relation_size(table_name::text)
FROM citus_tables;

-- 将特定分片迁移到另一个节点
SELECT citus_move_shard_placement(shard_id, 'source-host', 5432, 'dest-host', 5432);
```

### 配置参数

用于调优 Citus 的关键 GUC 参数：

```sql
-- 多分片查询时每个节点的并行连接数
SET citus.max_adaptive_executor_pool_size = 4;

-- 分片复制因子（默认 1；无流复制时可设为 2 以实现高可用）
SET citus.shard_replication_factor = 1;

-- 控制执行器行为
SET citus.multi_shard_modify_mode = 'parallel';   -- 或 'sequential'
SET citus.enable_repartition_joins = on;           -- 启用重分区连接

-- 任务分配策略
SET citus.task_assignment_policy = 'round-robin';  -- 或 'greedy'、'first-replica'

-- 记录分布式查询日志
SET citus.log_multi_join_order = on;
SET citus.explain_all_tasks = on;
```

### 示例：多租户 SaaS 模式

典型的多租户模式将所有租户相关的表按 `tenant_id` 进行分布：

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

-- 共享查找表：复制到每个节点
CREATE TABLE plans (
    name TEXT PRIMARY KEY,
    max_users INT,
    max_projects INT
);
SELECT create_reference_table('plans');

-- 所有限定在单个租户范围内的连接都是协同定位的，查询速度快
SELECT u.email, p.name AS project
FROM users u
JOIN projects p ON u.tenant_id = p.tenant_id AND u.user_id = p.owner_id
WHERE u.tenant_id = 7;
```

### 示例：实时分析

用于仪表板和分析的分布式聚合：

```sql
CREATE TABLE page_views (
    site_id      INT,
    url          TEXT,
    view_time    TIMESTAMPTZ DEFAULT now(),
    user_agent   TEXT,
    country      CHAR(2)
);
SELECT create_distributed_table('page_views', 'site_id');

-- 实时汇总：跨分片并行执行
SELECT
    date_trunc('minute', view_time) AS minute,
    count(*)                        AS views,
    count(DISTINCT country)         AS countries
FROM page_views
WHERE site_id = 1
  AND view_time > now() - INTERVAL '1 hour'
GROUP BY minute
ORDER BY minute DESC;

-- 所有站点的热门页面（跨分片并行查询）
SELECT url, count(*) AS views
FROM page_views
WHERE view_time > now() - INTERVAL '24 hours'
GROUP BY url
ORDER BY views DESC
LIMIT 20;
```
