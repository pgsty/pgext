## 用法

来源：

- [Azure Database for PostgreSQL 官方指南](https://learn.microsoft.com/en-us/azure/postgresql/extensions/how-to-use-pgdiskann)
- [Microsoft Research DiskANN 官方项目](https://www.microsoft.com/en-us/research/project/project-akupara-approximate-nearest-neighbor-search-for-large-scale-semantic-search/)

`pg_diskann` 是 Azure Database for PostgreSQL Flexible Server 为 `vector` 列提供 DiskANN 近似最近邻索引的扩展。它由云服务商托管，并非有公开目录源码的 PostgreSQL 扩展；可用性、受支持版本和运维限制均由 Azure 控制。

### 启用与查询

先在 Azure 服务器上允许 `pg_diskann`，然后在每个数据库中连同其 `vector` 依赖一起创建：

```sql
CREATE EXTENSION IF NOT EXISTS pg_diskann CASCADE;

CREATE TABLE demo (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding public.vector(3)
);

CREATE INDEX demo_embedding_diskann_idx
ON demo USING diskann (embedding vector_cosine_ops);

SELECT id, embedding
FROM demo
ORDER BY embedding <=> '[2.0,3.0,4.0]'
LIMIT 5;
```

该索引是近似索引。应使用符合生产形态的向量、过滤条件和并发度测量召回率与延迟，并用 `EXPLAIN (ANALYZE, BUFFERS)` 验证索引使用情况。

### 索引调优

索引存储参数包括 `max_neighbors`、`l_value_ib`、`product_quantized`、`pq_param_num_chunks` 和 `pq_param_training_samples`。产品量化从版本 `0.6` 起可用，可以减少内存；Azure 指南还要求维度超过 2,000、最高 16,000 时启用它：

```sql
CREATE INDEX demo_embedding_diskann_pq_idx
ON demo USING diskann (embedding vector_cosine_ops)
WITH (
  product_quantized = true,
  pq_param_num_chunks = 0,
  pq_param_training_samples = 0
);
```

`diskann.l_value_is` 在查询工作量与召回率之间取舍。`diskann.iterative_search` 接受 `relaxed_order`、`strict_order` 或 `off`；宽松顺序可能不按精确距离输出，而非迭代搜索返回的结果数不能超过搜索列表长度。对量化搜索得到的较大内部候选集使用全精度向量重新排序，可以改善最终顺序。

### 运维边界

构建索引可能消耗大量内存、存储和时间；Azure 文档把 `maintenance_work_mem`、表级 `parallel_workers` 和服务器工作进程上限列为构建控制项。修改 `max_worker_processes` 需要重启。不要为了强制索引而全局关闭 `enable_seqscan`：局部设置也会影响同一事务中的所有表，并可能掩盖规划或召回问题。使用版本特定选项前，应在目标 Azure 区域确认实际提供的扩展版本是否与目录中的 `0.6.5` 一致；索引重建、备份恢复、副本、故障转移和升级行为也必须通过 Azure 支持的流程测试。
