## 用法

来源：

- [Citus v14.2.0 发行说明](https://github.com/citusdata/citus/releases/tag/v14.2.0)
- [Citus v14.2.0 CHANGELOG](https://github.com/citusdata/citus/blob/v14.2.0/CHANGELOG.md)
- [Citus v14.2.0 控制文件](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/distributed/citus.control)
- [什么是 Citus？](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Citus 实用函数](https://docs.citusdata.com/en/stable/develop/api_udf.html)

Citus 通过把表分片到工作节点，将 PostgreSQL 转换为分布式数据库，同时仍以 PostgreSQL SQL、索引、扩展、事务和运维工具作为面向用户的接口。它常用于多租户 SaaS 数据库、实时分析、时序/事件工作负载和分布式微服务模式。

本地 Pigsty 目录将 Citus 打包为 `citus`，并提供主扩展 `citus`；同一软件包还包含 `citus_columnar`。Citus 是预加载扩展，因此每个节点都必须在执行 `CREATE EXTENSION` 之前加载其库。

### 启用 Citus

```conf
shared_preload_libraries = 'citus'
```

在协调节点和工作节点上重启 PostgreSQL，然后在数据库中创建扩展：

```sql
CREATE EXTENSION IF NOT EXISTS citus;
SELECT citus_version();
```

在多节点集群中，从协调节点注册协调节点自身与各工作节点：

```sql
SELECT citus_set_coordinator_host('coord-1', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

SELECT * FROM citus_get_active_worker_nodes();
```

### 分布式表

按分片键分布表。分片键值相同的行会被共置在同一分片上，使限定租户的连接和点查询可以在本地完成。

```sql
CREATE TABLE events (
  tenant_id  bigint,
  event_id   bigserial,
  event_at   timestamptz DEFAULT now(),
  kind       text,
  payload    jsonb,
  PRIMARY KEY (tenant_id, event_id)
);
```

分布该表，并显式调整分片数与共置设置：

```sql
SELECT create_distributed_table(
  'events',
  'tenant_id',
  shard_count  := 64,
  colocate_with := 'none'
);
```

显式选择分片数时，应使用 `colocate_with := 'none'` 启动新的共置组。要与现有分布式表共置，请指定该表的名称，并让它的分片数决定布局。

过滤分布列的查询可以路由到单个分片：

```sql
SELECT *
FROM events
WHERE tenant_id = 42
ORDER BY event_at DESC
LIMIT 50;
```

跨分片查询会被规划为分布式任务，并在工作节点上并行运行：

```sql
SELECT kind, count(*)
FROM events
WHERE event_at >= now() - interval '1 hour'
GROUP BY kind
ORDER BY count DESC;
```

### 参考表

参考表会完整复制到所有工作节点。它们适用于需要与许多分布式表连接的小型查找表。

```sql
CREATE TABLE countries (
  code text PRIMARY KEY,
  name text NOT NULL
);

SELECT create_reference_table('countries');
```

### 基于模式的分片

当每个租户或服务拥有独立模式时，基于模式的分片很有用。Citus 支持从任意节点执行模式分片 DDL，包括 `CREATE SCHEMA`、`DROP SCHEMA`、`ALTER SCHEMA RENAME`、`ALTER SCHEMA OWNER`，以及分布式模式上的表级 DDL。

```sql
CREATE SCHEMA tenant_42;
SELECT citus_schema_distribute('tenant_42');

CREATE TABLE tenant_42.orders (
  id bigserial PRIMARY KEY,
  amount numeric,
  created_at timestamptz DEFAULT now()
);
```

共享表应使用基于行的分布，每租户模式布局应使用基于模式的分片；在未检查共置与 SQL 支持影响之前，不要随意混用这两种模型。

### 节点与分片操作

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

### 备份协调

Citus v14.1.0 新增了 UDF，可在制作协调一致的磁盘快照时阻止分布式 2PC 提交决策和模式/拓扑变更。只能在受控备份流程中使用这些函数，并且在完成快照步骤后必须解除集群阻塞。

```sql
SELECT citus_cluster_changes_block();
SELECT * FROM citus_cluster_changes_block_status();

-- Take coordinated filesystem or volume snapshots here.

SELECT citus_cluster_changes_unblock();
```

应将这些函数与常规 PostgreSQL 备份规范配合使用：一致检查点、WAL 归档、跨节点快照顺序，以及经过测试的恢复流程。

### 版本 14.2 运维

Citus 14.2 新增了仅限超级用户使用的 `citus_internal.distribute_object()` 修复辅助函数，用于处理元数据未正确传播的受支持数据库对象。应将它视为针对性的恢复操作，而不是常规分布 API。

该版本还新增了 `citus.allow_unsafe_insert_select_pushdown`，用于显式选择在共置表上启用批量 `INSERT ... SELECT` 下推，并改进了单分片存储过程执行。除非已经依据发行说明中的限制检查工作负载，否则应保持关闭不安全下推设置。

### 注意事项

- Pigsty 本地元数据当前跟踪面向 PostgreSQL 16-18 的 Citus 14.x；Citus 14 已停止支持 PostgreSQL 15。
- 必须在创建扩展之前设置 `shared_preload_libraries = 'citus'`。在全新服务器上，仅执行 `CREATE EXTENSION citus` 并不充分。
- 应仔细选择分布列。分布式表上的主键和唯一约束通常需要包含分布列。
- 跨分片连接、重分区连接、分布式 DDL 和多分片写入功能强大，但其规划与锁行为不同于单节点 PostgreSQL。
- Citus 通过 `citus_columnar` 提供自己的列存储接口；Pigsty 元数据将其标记为与 Hydra `columnar` 冲突。
- 集群变更阻塞函数是用于备份的运维工具。备份脚本失败后，不要让集群一直处于阻塞状态。
