


## 用法

来源：

- [pgvector v0.8.4 release](https://github.com/pgvector/pgvector/releases/tag/v0.8.4)
- [pgvector v0.8.4 README](https://github.com/pgvector/pgvector/blob/v0.8.4/README.md)
- [pgvector v0.8.4 CHANGELOG](https://github.com/pgvector/pgvector/blob/v0.8.4/CHANGELOG.md)
- [本地包元数据](../db/extension.csv)

`pgvector` 在 PostgreSQL 内提供向量相似性搜索。扩展名是 `vector`，Pigsty 中的包名是 `pgvector`。它支持精确搜索、基于 HNSW 与 IVFFlat 的近似最近邻搜索，并提供 dense、half-precision、binary、sparse 等多种向量表示。

v0.8.4 是 0.8.x HNSW / vacuum 修复之后的维护版本。维护写入较多的 HNSW 索引时，应优先使用它而不是更早的 0.8.x 构建。

### 创建与查询向量

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

常用距离操作符：

- `<->` 表示 L2 距离
- `<#>` 表示负内积
- `<=>` 表示余弦距离
- `<+>` 表示 L1 距离
- `<~>` 表示二进制向量上的 Hamming 距离
- `<%>` 表示二进制向量上的 Jaccard 距离

由于 PostgreSQL 索引扫描按升序工作，`<#>` 返回负内积；展示实际内积时需要乘以 `-1`。

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

`vector` 是标准单精度类型。`halfvec` 可降低存储和内存压力，`bit` 适合二进制签名，`sparsevec` 适合高维稀疏向量。

向量列可以使用 `avg()`、`sum()` 等聚合：

```sql
SELECT avg(embedding) FROM items;
```

### HNSW 索引

HNSW 提供较好的速度/召回权衡，并且不需要训练步骤。

```sql
CREATE INDEX items_embedding_hnsw
ON items USING hnsw (embedding vector_l2_ops);

SET hnsw.ef_search = 100;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

操作符类必须匹配距离类型：

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
CREATE INDEX ON embeddings USING hnsw (half_dense halfvec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (sparse sparsevec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (binary_sig bit_hamming_ops);
```

常用调优参数包括 `hnsw.ef_search`、`hnsw.iterative_scan`、`hnsw.max_scan_tuples` 和 `hnsw.scan_mem_multiplier`。

### IVFFlat 索引

IVFFlat 构建时需要训练聚类列表，因此应在载入有代表性的数据后再创建。

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

大表通常需要提高 `lists`；提高 `ivfflat.probes` 可以提升召回率。带过滤条件的查询要实测普通 btree 过滤、partial vector index 和分区表哪种计划更好。

### 过滤与混合搜索

向量排序可以和普通 PostgreSQL 过滤结合：

```sql
CREATE INDEX ON items (tenant_id);

SELECT *
FROM items
WHERE tenant_id = 42
ORDER BY embedding <=> '[0.1,0.2,0.3]'
LIMIT 20;
```

混合搜索可以把 `pgvector` 与 PostgreSQL 全文检索、trigram 检索或自定义排序表达式结合：

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

HNSW 索引可能很大，构建成本也高。构建时配置 `maintenance_work_mem`，观察 build notice；当索引膨胀或召回漂移重要时，安排 `REINDEX`。

### 注意事项

- Pigsty 本地元数据可能落后于上游版本；本文档按上游 pgvector 0.8.4 编写，本地包行可能要等包 catalog 刷新后才显示同一版本。
- 查询操作符必须匹配索引操作符类。cosine 索引不会加速 L2 `ORDER BY`。
- 近似索引用速度换取非精确召回。应使用代表性数据和过滤条件验证 recall。
- IVFFlat 应在导入数据后构建；数据分布明显变化后需要重建。
- 如果 HNSW 表有大量写入和 vacuum 活动，应保持 pgvector 在较新版本；0.8.x 包含重要的 HNSW 维护修复。
