## 用法

来源：

- [官方 pg_shardman README](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/readme.md)
- [1.0 版扩展 SQL](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman--1.0.sql)
- [扩展控制文件](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman.control)

`pg_shardman` 是面向 PostgreSQL 10 的实验性集群管理器，基于 `pg_pathman`、`postgres_fdw` 和逻辑复制管理哈希分片表。项目明确表示尚未完全实现目标；多个重要流程需要打补丁的 PostgreSQL 和依赖，因此它应被视为历史实验软件，而不是受支持的分片方案。

### 集群概要

在每台服务器上安装扩展及其依赖，将 `pg_pathman` 加入 `shared_preload_libraries`，然后重启所有实例。指定一个只保存元数据的 shardlord，并在所有节点配置：

```conf
shardman.shardlord = on
shardman.shardlord_connstring = 'host=shardlord dbname=app'
shardman.sync_replication = off
```

工作节点上的 `shardman.shardlord = off`。当前 shardlord 不能同时存储普通应用数据。在全部节点创建 `pg_shardman` 后，到 shardlord 执行集群管理调用：

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

初始分区数以后无法修改，因此应按未来节点增长来选择。复制组限制副本可放置的位置。默认组即使在冗余度为零时也会建立全互联逻辑复制通道；不需要复制时应明确规划不同的组名。

### 管理接口

- `shardman.create_shared_table` 从一个主节点复制小型参考表。
- `shardman.mv_partition`、`shardman.mv_replica`、`shardman.rebalance` 和 `shardman.rebalance_replicas` 用于移动或重新分布数据。
- `shardman.set_redundancy` 提高表的副本目标，`shardman.ensure_redundancy` 补齐缺少的副本。
- `shardman.nodes`、`shardman.tables`、`shardman.partitions`、`shardman.replicas` 和 `shardman.replication_lag` 提供元数据与延迟信息。
- `shardman.forall` 广播 DDL，`shardman.recovery` 尝试按 shardlord 元数据协调节点状态。

### 关键限制

标准 PostgreSQL 10 无法对这里使用的外部分区执行 `COPY FROM`；文档所述路径需要打补丁的 PostgreSQL 和 `pg_pathman`。可靠的跨节点事务需要带两阶段提交和分布式快照补丁的 `postgres_fdw`。分片表的 `ALTER TABLE`、全局主键和指向分片表的外键大多不受支持，跨节点 OLAP 也无法获得并行外部扫描。

故障检测与恢复完全依赖人工。存在多个副本时会随机提升副本，且不会协调副本间的分歧状态。`shardman.rm_node(..., true)` 可能丢弃数据，`shardman.rm_table` 也没有确认提示。共享表事务可能读不到自身刚复制的写入。开始任何实验前，都应通读上游限制，隔离集群，并演练节点丢失、shardlord 恢复和数据协调。
