
## 用法

### 快速上手

创建一个 3 维向量列

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
```

插入向量

```sql
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
```

按 L2 距离获取最近邻

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

还支持内积（`<#>`）、余弦距离（`<=>`）和 L1 距离（`<+>`）

注意：`<#>` 返回的是负内积，因为 PostgreSQL 的索引扫描仅支持 `ASC` 排序的操作符


--------

### 增删改查

创建一个带有向量列的新表

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
```

或者向已有表添加向量列

```sql
ALTER TABLE items ADD COLUMN embedding vector(3);
```

还支持[半精度](https://github.com/pgvector/pgvector#half-precision-vectors)、[二进制](https://github.com/pgvector/pgvector#binary-vectors)和[稀疏](https://github.com/pgvector/pgvector#sparse-vectors)向量

**插入向量**

```sql
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
```

或使用 `COPY` 批量导入向量（[示例](https://github.com/pgvector/pgvector-python/blob/master/examples/loading/example.py)）

```sql
COPY items (embedding) FROM STDIN WITH (FORMAT BINARY);
```

**更新插入（Upsert）向量**

```sql
INSERT INTO items (id, embedding) VALUES (1, '[1,2,3]'), (2, '[4,5,6]')
    ON CONFLICT (id) DO UPDATE SET embedding = EXCLUDED.embedding;
```

**更新向量**

```sql
UPDATE items SET embedding = '[1,2,3]' WHERE id = 1;
```

**删除向量**

```sql
DELETE FROM items WHERE id = 1;
```


--------

### 查询

获取某个向量的最近邻

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

支持的距离函数：

- `<->` - L2 距离
- `<#>` - （负）内积
- `<=>` - 余弦距离
- `<+>` - L1 距离
- `<~>` - 汉明距离（二进制向量）
- `<%>` - Jaccard 距离（二进制向量）

获取某行的最近邻

```sql
SELECT * FROM items WHERE id != 1 ORDER BY embedding <-> (SELECT embedding FROM items WHERE id = 1) LIMIT 5;
```

获取指定距离范围内的行

```sql
SELECT * FROM items WHERE embedding <-> '[3,1,2]' < 5;
```

注意：配合 `ORDER BY` 和 `LIMIT` 使用才能利用索引


#### 距离计算

获取距离值

```sql
SELECT embedding <-> '[3,1,2]' AS distance FROM items;
```

对于内积，需要乘以 -1（因为 `<#>` 返回的是负内积）

```sql
SELECT (embedding <#> '[3,1,2]') * -1 AS inner_product FROM items;
```

对于余弦相似度，使用 1 减去余弦距离

```sql
SELECT 1 - (embedding <=> '[3,1,2]') AS cosine_similarity FROM items;
```

#### 聚合

计算向量平均值

```sql
SELECT AVG(embedding) FROM items;
```

按分组计算向量平均值

```sql
SELECT category_id, AVG(embedding) FROM items GROUP BY category_id;
```


--------

### HNSW 索引

默认情况下，pgvector 执行精确最近邻搜索，提供完美的召回率。

可以添加索引来使用近似最近邻搜索，以牺牲部分召回率换取速度提升。与普通索引不同，添加近似索引后查询结果可能会有所变化。

支持的索引类型：

- [HNSW](https://github.com/pgvector/pgvector#hnsw)
- [IVFFlat](https://github.com/pgvector/pgvector#ivfflat)

HNSW 索引会创建一个多层图结构。相比 IVFFlat，它具有更好的查询性能（在速度与召回率的权衡上），但构建时间更长且占用更多内存。此外，由于没有像 IVFFlat 那样的训练步骤，HNSW 索引可以在表中没有数据时就创建。

为所需的每种距离函数添加索引。

L2 距离

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops);
```

注意：`halfvec` 类型使用 `halfvec_l2_ops`，`sparsevec` 类型使用 `sparsevec_l2_ops`（其他距离函数同理）

**内积**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
```

**余弦距离**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
```

**L1 距离**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
```

**汉明距离**

```sql
CREATE INDEX ON items USING hnsw (embedding bit_hamming_ops);
```

**Jaccard 距离**

```sql
CREATE INDEX ON items USING hnsw (embedding bit_jaccard_ops);
```

支持的类型：

- `vector` - 最多 2,000 维
- `halfvec` - 最多 4,000 维
- `bit` - 最多 64,000 维
- `sparsevec` - 最多 1,000 个非零元素

#### 索引选项

指定 HNSW 参数

- `m` - 每层的最大连接数（默认 16）
- `ef_construction` - 构建图时动态候选列表的大小（默认 64）

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops) WITH (m = 16, ef_construction = 64);
```

`ef_construction` 值越大，召回率越高，但索引构建时间和插入耗时也会相应增加。

#### 查询选项

指定搜索时动态候选列表的大小（默认 40）

```sql
SET hnsw.ef_search = 100;
```

该值越大，召回率越高，但查询速度会相应下降。

在事务中使用 `SET LOCAL` 可以仅对单次查询生效

```sql
BEGIN;
SET LOCAL hnsw.ef_search = 100;
SELECT ...
COMMIT;
```

#### 索引构建时间

当图结构能完全放入 `maintenance_work_mem` 时，索引构建速度会显著提升

```sql
SET maintenance_work_mem = '8GB';
```

当图结构无法再放入内存时，会显示如下提示

```
NOTICE:  hnsw graph no longer fits into maintenance_work_mem after 100000 tuples
DETAIL:  Building will take significantly more time.
HINT:  Increase maintenance_work_mem to speed up builds.
```


注意：不要将 `maintenance_work_mem` 设置得过高以至于耗尽服务器内存

与其他索引类型一样，在加载完初始数据之后再创建索引会更快

还可以通过增加并行工作进程数来加速索引创建（默认 2 个）

```sql
SET max_parallel_maintenance_workers = 7; -- 加上 leader 进程
```

如果工作进程数较多，可能需要同时增大 `max_parallel_workers`（默认 8）

[索引选项](https://github.com/pgvector/pgvector#index-options)对构建时间也有显著影响（除非召回率偏低，否则使用默认值即可）

#### 索引构建进度

查看[索引构建进度](https://www.postgresql.org/docs/current/progress-reporting.html#CREATE-INDEX-PROGRESS-REPORTING)

```sql
SELECT phase, round(100.0 * blocks_done / nullif(blocks_total, 0), 1) AS "%" FROM pg_stat_progress_create_index;
```

HNSW 的构建阶段：

1. `initializing`
2. `loading tuples`
