## 用法

来源：

- [官方 README 与弃用声明](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/README.md)
- [1.2 扩展 SQL](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/sql/pg_shard.sql)
- [官方变更日志](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/CHANGELOG.md)

`pg_shard` 1.2 是 Citus Data 已退役的分片扩展，用于把哈希分区表分布到多个 PostgreSQL 工作节点。上游已宣布其生命周期结束，并将开源 Citus 扩展列为功能超集；本文只用于维护或迁移存量部署，不应据此新建集群。

### 集群设置

在协调节点上预加载模块，并在数据目录下的文件中每行列出一个工作节点主机与端口。协调节点必须能通过 TCP 连接工作节点且无需交互式认证，每个工作节点上也必须预先存在同名数据库。

```conf
shared_preload_libraries = 'pg_shard'
```

```text
# pg_worker_list.conf
worker-101  5432
worker-102  5432
```

重启协调节点，创建扩展和原型表，再创建分片。

```sql
CREATE EXTENSION pg_shard;

CREATE TABLE customer_reviews (
  customer_id text NOT NULL,
  review_rating integer
);

SELECT master_create_distributed_table('customer_reviews', 'customer_id');
SELECT master_create_worker_shards('customer_reviews', 16, 2);
```

### 主要操作与限制

- `master_create_distributed_table(...)`：选择哈希分区列。
- `master_create_worker_shards(...)`：按指定分片数与复制因子创建分片。
- `master_copy_shard_placement(...)`：修复失效副本；修复要求所有工作节点都安装 `pg_shard`。
- `pgs_distribution_metadata.partition` 与 `pgs_distribution_metadata.shard`：暴露分布元数据。

协调节点上的原型表用于保存元数据和模式，而不是用户行。1.2 版本要求 `UPDATE` 与 `DELETE` 的 `WHERE` 子句包含分区列。工作节点认证、故障处理、分片修复和迁移都需要明确运维流程。还要注意预加载与重启要求，并在修改继承来的集群前规划迁移到受支持的 Citus。
