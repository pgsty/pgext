## 用法

来源：

- [pgvector v0.8.5 README](https://github.com/pgvector/pgvector/blob/v0.8.5/README.md)
- [pgvector v0.8.5 CHANGELOG](https://github.com/pgvector/pgvector/blob/v0.8.5/CHANGELOG.md)
- [Changes from v0.8.4 to v0.8.5](https://github.com/pgvector/pgvector/compare/v0.8.4...v0.8.5)

`pgvector` 在 PostgreSQL 内部提供了向量相似性搜索功能。扩展名称为 `vector`，而 Pigsty 将其打包为 `pgvector`。它支持精确搜索、使用 HNSW 和 IVFFlat 索引的近似最近邻搜索以及多种向量表示形式，包括密集型、半精度、二进制和稀疏嵌入。

版本 `0.8.5` 在构建小表上的 IVFFlat 索引时减少了内存使用。它保留了当前 README 中记录的 0.8.x HNSW 迭代扫描和维护改进。

### 创建和查询向量

```sql
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE items (
  id bigserial PRIMARY KEY,
  embedding vector(3)
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 5;
```

常用的距离操作符：

- `<->` 表示 L2 距离
- `<#>` 表示负内积
- `<=>` 表示余弦距离
- `<+>` 表示 L1 距离
- `<~>` 表示二进制向量的汉明距离
- `<%>` 表示二进制向量的 Jaccard 距离

由于 PostgreSQL 索引扫描顺序为升序，`<#>` 返回的是负内积；在显示实际内积时需要乘以 `-1`。

### 向量类型

```sql
CREATE TABLE embeddings (
  id bigserial PRIMARY KEY,
  dense      vector(768),
  half_dense halfvec(768),
  binary_sig bit(1024),
  sparse     sparsevec(100000)
);
```

`vector` 是标准的单精度类型。使用 `halfvec` 可以减少存储和内存压力，使用 `bit` 用于二进制签名，使用 `sparsevec` 用于高维稀疏向量。

可以对向量列使用聚合函数如 `avg()` 和 `sum()`：

```sql
SELECT avg(embedding) FROM items;
```

### HNSW 索引

HNSW 提供了速度和召回率之间的强权衡，并不需要训练步骤。

```sql
CREATE INDEX items_embedding_hnsw
ON items USING hnsw (embedding vector_l2_ops);

SET hnsw.ef_search = 100;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

选择与距离匹配的操作符类：

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
CREATE INDEX ON embeddings USING hnsw (half_dense halfvec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (sparse sparsevec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (binary_sig bit_hamming_ops);
```

有用的调优设置包括 `hnsw.ef_search`、`hnsw.iterative_scan`、`hnsw.max_scan_tuples` 和 `hnsw.scan_mem_multiplier`。

### IVFFlat 索引

IVFFlat 在索引创建前需要代表性的数据，因为它在构建时会训练聚类列表。

```sql
CREATE INDEX items_embedding_ivfflat
ON items USING ivfflat (embedding vector_l2_ops)
WITH (lists = 100);

SET ivfflat.probes = 10;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

增加 `lists` 以适应更大的表，并通过增加 `ivfflat.probes` 来提高召回率。对于过滤查询，请测试是否使用精确的 btree 过滤、部分向量索引或分区可以提供更好的计划。

### 混合搜索

普通的 PostgreSQL 过滤器可以与向量排序结合使用：

```sql
CREATE INDEX ON items (tenant_id);

SELECT *
FROM items
WHERE tenant_id = 42
ORDER BY embedding <=> '[0.1,0.2,0.3]'
LIMIT 20;
```

对于混合搜索，可以将 `pgvector` 与 PostgreSQL 全文搜索、三元组搜索或外部排名表达式结合起来：

```sql
SELECT id,
       ts_rank_cd(text_tsv, plainto_tsquery('database')) AS text_score,
       1 - (embedding <=> '[0.1,0.2,0.3]') AS vector_score
FROM docs
WHERE text_tsv @@ plainto_tsquery('database')
ORDER BY vector_score DESC
LIMIT 20;
```

### 维护

```sql
VACUUM items;
REINDEX INDEX CONCURRENTLY items_embedding_hnsw;
ANALYZE items;
```

HNSW 索引可能很大且构建成本高昂。使用 `maintenance_work_mem` 进行构建，监控构建通知，并在索引膨胀或召回漂移时安排 `REINDEX`。

### 注意事项

- 版本 `0.8.5` 是一个专注于 IVFFlat 构建内存的补丁；它不会改变 SQL 查询表面。当数据库报告较旧的 SQL 版本时，在安装新的扩展文件后运行 `ALTER EXTENSION vector UPDATE`。
- 使用与查询操作符匹配的操作符类。余弦索引不会加速 L2 `ORDER BY` 操作。
- 近似索引以牺牲精确召回率为代价换取速度。请使用代表性数据和查询过滤器验证召回率。
- 在加载数据后构建 IVFFlat 索引。如果数据分布发生显著变化，请重新构建索引。
- 当使用 HNSW 并且存在大量写入和 vacuum 活动时，保持 pgvector 更新；`0.8.x` 系列包括重要的 HNSW 维护修复。

这里仍采用负内积操作符；版本对比基线是 `0.8.4`。
