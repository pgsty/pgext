## 用法

来源：

- [pgvector v0.8.6 README](https://github.com/pgvector/pgvector/blob/v0.8.6/README.md)
- [pgvector v0.8.6 CHANGELOG](https://github.com/pgvector/pgvector/blob/v0.8.6/CHANGELOG.md)
- [从 v0.8.5 到 v0.8.6 的变更](https://github.com/pgvector/pgvector/compare/v0.8.5...v0.8.6)

`pgvector` 在 PostgreSQL 内提供向量相似性搜索。扩展名为 `vector`，Pigsty 将其打包为 `pgvector`。它支持精确搜索、使用 HNSW 与 IVFFlat 索引的近似最近邻搜索，以及用于稠密、半精度、二进制和稀疏嵌入的多种向量表示。

版本 `0.8.6` 是一个专注于正确性的修复版本，同时保留了当前 README 中介绍的 0.8.x HNSW 迭代扫描与维护改进。

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

- `<->`：L2 距离
- `<#>`：负内积
- `<=>`：余弦距离
- `<+>`：L1 距离
- `<~>`：二进制向量的汉明距离
- `<%>`：二进制向量的杰卡德距离

由于 PostgreSQL 索引按升序扫描，`<#>` 返回负内积；显示实际内积时应乘以 `-1`。

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

`vector` 是标准单精度类型。可使用 `halfvec` 降低存储和内存压力，使用 `bit` 表示二进制签名，并使用 `sparsevec` 表示高维稀疏向量。

`avg()` 和 `sum()` 等聚合函数可以用于向量列：

```sql
SELECT avg(embedding) FROM items;
```

### HNSW 索引

HNSW 在速度与召回率之间提供良好的权衡，且不需要训练步骤。

```sql
CREATE INDEX items_embedding_hnsw
ON items USING hnsw (embedding vector_l2_ops);

SET hnsw.ef_search = 100;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

请选择与距离匹配的操作符类：

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
CREATE INDEX ON embeddings USING hnsw (half_dense halfvec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (sparse sparsevec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (binary_sig bit_hamming_ops);
```

常用调优设置包括 `hnsw.ef_search`、`hnsw.iterative_scan`、`hnsw.max_scan_tuples` 和 `hnsw.scan_mem_multiplier`。

### IVFFlat 索引

IVFFlat 在构建索引时需要训练聚类列表，因此创建索引之前必须已有具备代表性的数据。

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

表越大，应增加 `lists`；要提高召回率，则增加 `ivfflat.probes`。对于带过滤条件的查询，应测试精确的 btree 过滤器、部分向量索引或分区中哪一种能产生更好的执行计划。

### 过滤与混合搜索

普通 PostgreSQL 过滤条件可以与向量排序组合使用：

```sql
ALTER TABLE items ADD COLUMN tenant_id bigint;
CREATE INDEX ON items (tenant_id);

SELECT *
FROM items
WHERE tenant_id = 42
ORDER BY embedding <=> '[0.1,0.2,0.3]'
LIMIT 20;
```

进行混合搜索时，可以将 `pgvector` 与 PostgreSQL 全文搜索、三元组搜索或外部排序表达式组合使用：

```sql
CREATE TABLE docs (
  id bigint PRIMARY KEY,
  body text NOT NULL,
  text_tsv tsvector GENERATED ALWAYS AS
    (to_tsvector('english', body)) STORED,
  embedding vector(3)
);

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

HNSW 索引可能很大，构建成本也可能很高。构建时使用 `maintenance_work_mem`，监控构建通知，并在索引膨胀或召回率漂移成为问题时安排 `REINDEX`。

### 注意事项

- 版本 `0.8.6` 修复了 32 位系统上的 IVFFlat 构建溢出、将数组转换为 `sparsevec` 时未强制执行非零元素上限，以及嵌套循环中的 IVFFlat 扫描内存增长问题。它没有新增 SQL 功能界面。安装新的扩展文件后，如果数据库报告的是较旧 SQL 版本，请运行 `ALTER EXTENSION vector UPDATE`。
- 使用与查询操作符匹配的操作符类。余弦索引无法加速 L2 `ORDER BY`。
- 近似索引以精确召回率换取速度。请使用有代表性的数据和查询过滤条件验证召回率。
- 应在数据加载后构建 IVFFlat。如果数据分布发生显著变化，请重建索引。
- 在高强度写入和 vacuum 活动场景中使用 HNSW 时，请及时更新 pgvector；`0.8.x` 系列包含重要的 HNSW 维护修复。
