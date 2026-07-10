


## 用法

来源：

- [Citus v14.1.0 release](https://github.com/citusdata/citus/releases/tag/v14.1.0)
- [Citus v14.1.0 CHANGELOG](https://github.com/citusdata/citus/blob/v14.1.0/CHANGELOG.md)
- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Citus Utility Functions](https://docs.citusdata.com/en/stable/develop/api_udf.html)

Citus 通过把表分片到多个 worker 节点，把 PostgreSQL 扩展成分布式数据库，同时保留 PostgreSQL SQL、索引、扩展、事务和运维工具作为主要使用界面。它常用于多租户 SaaS、实时分析、时间序列/事件数据和分布式微服务 schema。

Pigsty 本地 catalog 中包名和主扩展名都是 `citus`；同一个包还包含 `citus_columnar`。Citus 是 preload 扩展，每个节点都需要先加载库，再创建扩展。

### 启用 Citus

```conf
shared_preload_libraries = 'citus'
```

在 coordinator 和 worker 上重启 PostgreSQL，然后在数据库中创建扩展：

```sql
CREATE EXTENSION IF NOT EXISTS citus;
SELECT citus_version();
```

多节点集群中，在 coordinator 上注册 coordinator 与 worker：

```sql
SELECT citus_set_coordinator_host('coord-1', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

SELECT * FROM citus_get_active_worker_nodes();
```

### 分布式表

分布式表通过分片键分布。相同分片键值的行会协同定位到同一个 shard，因此租户内 join 和点查可以保持本地执行。

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

也可以显式设置 shard 数量和 colocation：

```sql
SELECT create_distributed_table(
  'events',
  'tenant_id',
  shard_count  := 64,
  colocate_with := 'default'
);
```

带分布列过滤的查询可以路由到单个 shard：

```sql
SELECT *
FROM events
WHERE tenant_id = 42
ORDER BY event_at DESC
LIMIT 50;
```

跨 shard 查询会被规划为分布式任务，在 worker 上并行执行：

```sql
SELECT kind, count(*)
FROM events
WHERE event_at >= now() - interval '1 hour'
GROUP BY kind
ORDER BY count DESC;
```

### 引用表

引用表会完整复制到所有 worker，适合需要和多个分布式表 join 的小型维表。

```sql
CREATE TABLE countries (
  code text PRIMARY KEY,
  name text NOT NULL
);

SELECT create_reference_table('countries');
```

### 基于 Schema 的分片

当每个租户或服务拥有独立 schema 时，可以使用 schema-based sharding。Citus v14.1.0 增加了从任意节点执行多类 schema 分片 DDL 的支持，包括 `CREATE SCHEMA`、`DROP SCHEMA`、`ALTER SCHEMA RENAME`、`ALTER SCHEMA OWNER`，以及分布式 schema 内的表级 DDL。

```sql
CREATE SCHEMA tenant_42;
SELECT citus_schema_distribute('tenant_42');

CREATE TABLE tenant_42.orders (
  id bigserial PRIMARY KEY,
  amount numeric,
  created_at timestamptz DEFAULT now()
);
```

共享表通常使用行分布模型；每租户 schema 布局可以使用 schema 分片。不要在没有检查 colocation 和 SQL 支持限制的情况下随意混用两种模型。

### 节点与 Shard 运维

```sql
-- 添加或禁用节点。
SELECT * FROM citus_add_node('worker-3', 5432);
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- 排空并移除节点。
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- 重平衡 shard。
SELECT citus_rebalance_start();
SELECT * FROM citus_rebalance_status();
SELECT rebalance_table_shards('events');

-- 查看表和 shard。
SELECT * FROM citus_tables;
SELECT * FROM citus_shards;
```

### 备份协调

Citus v14.1.0 增加了用于临时阻塞分布式 2PC 提交决策以及 schema/topology 变更的 UDF，方便在协调磁盘快照时获得一致性窗口。它们只应该出现在受控备份流程里，快照步骤结束后必须解除阻塞。

```sql
SELECT citus_cluster_changes_block();
SELECT * FROM citus_cluster_changes_block_status();

-- 此处执行协调后的文件系统或卷快照。

SELECT citus_cluster_changes_unblock();
```

这些函数仍然需要配合常规 PostgreSQL 备份纪律：一致 checkpoint、WAL 归档、跨节点快照顺序，以及经过验证的恢复流程。

### 注意事项

- Pigsty 本地元数据当前跟踪 PostgreSQL 16-18 上的 Citus 14.x；Citus 14 已移除 PostgreSQL 15 支持。
- 创建扩展前必须设置 `shared_preload_libraries = 'citus'`。新实例上单纯执行 `CREATE EXTENSION citus` 不够。
- 分布列选择很关键。分布式表上的主键和唯一约束通常需要包含分布列。
- 跨 shard join、repartition join、分布式 DDL 和多 shard 写入虽然强大，但规划、锁和失败语义都不同于单机 PostgreSQL。
- Citus 通过 `citus_columnar` 提供自己的列式存储界面；Pigsty 元数据中它与 Hydra `columnar` 冲突。
- cluster-change blocking 函数是备份运维工具。备份脚本失败后不要让集群长期保持阻塞状态。
