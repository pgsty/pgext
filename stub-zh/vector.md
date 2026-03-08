

## 用法

适用于 PostgreSQL 的开源向量相似性搜索扩展。将向量与其他数据存储在一起，支持：

- 精确与近似最近邻搜索
- 单精度、半精度、二进制和稀疏向量
- L2 距离、内积、余弦距离、L1 距离、汉明距离和 Jaccard 距离
- 任何具有 PostgreSQL 客户端的[编程语言](https://github.com/pgvector/pgvector#languages)

此外还具备 [ACID](https://en.wikipedia.org/wiki/ACID) 合规性、时间点恢复（PITR）、JOIN 以及 PostgreSQL 的所有其他[优秀特性](https://www.postgresql.org/about/)

### 快速上手

启用扩展（在每个需要使用该扩展的数据库中执行一次）

```tsql
CREATE EXTENSION vector;
```

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

### 存储

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

```tsql
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

## 索引

默认情况下，pgvector 执行精确最近邻搜索，提供完美的召回率。

可以添加索引来使用近似最近邻搜索，以牺牲部分召回率换取速度提升。与普通索引不同，添加近似索引后查询结果可能会有所变化。

支持的索引类型：

- [HNSW](https://github.com/pgvector/pgvector#hnsw)
- [IVFFlat](https://github.com/pgvector/pgvector#ivfflat)

### HNSW

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

### IVFFlat

IVFFlat 索引将向量划分到多个列表中，然后搜索与查询向量最接近的那些列表的子集。相比 HNSW，它构建速度更快且占用内存更少，但查询性能较低（在速度与召回率的权衡上）。

实现良好召回率的三个关键：

1. 在表中已有一定数据*之后*再创建索引
2. 选择合适的列表数量——对于不超过 100 万行的数据，可从 `rows / 1000` 开始；超过 100 万行的数据则使用 `sqrt(rows)`
3. 查询时指定合适的[探针数](https://github.com/pgvector/pgvector#query-options-1)（值越高召回率越好，值越低速度越快）——可从 `sqrt(lists)` 开始

为所需的每种距离函数添加索引。

L2 距离

```sql
CREATE INDEX ON items USING ivfflat (embedding vector_l2_ops) WITH (lists = 100);
```

注意：`halfvec` 类型使用 `halfvec_l2_ops`（其他距离函数同理）

**内积**

```sql
CREATE INDEX ON items USING ivfflat (embedding vector_ip_ops) WITH (lists = 100);
```

**余弦距离**

```sql
CREATE INDEX ON items USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);
```

**汉明距离**

```sql
CREATE INDEX ON items USING ivfflat (embedding bit_hamming_ops) WITH (lists = 100);
```

支持的类型：

- `vector` - 最多 2,000 维
- `halfvec` - 最多 4,000 维
- `bit` - 最多 64,000 维

#### 查询选项

指定探针数（默认 1）

```sql
SET ivfflat.probes = 10;
```

该值越大，召回率越高，但速度会相应下降。可以将其设置为列表总数以实现精确最近邻搜索（此时查询规划器将不会使用索引）

在事务中使用 `SET LOCAL` 可以仅对单次查询生效

```sql
BEGIN;
SET LOCAL ivfflat.probes = 10;
SELECT ...
COMMIT;
```

#### 索引构建时间

对于大表，可以通过增加并行工作进程数来加速索引创建（默认 2 个）

```sql
SET max_parallel_maintenance_workers = 7; -- 加上 leader 进程
```

如果工作进程数较多，可能还需要同时增大 `max_parallel_workers`（默认 8）

#### 索引构建进度

查看[索引构建进度](https://www.postgresql.org/docs/current/progress-reporting.html#CREATE-INDEX-PROGRESS-REPORTING)

```sql
SELECT phase, round(100.0 * tuples_done / nullif(tuples_total, 0), 1) AS "%" FROM pg_stat_progress_create_index;
```

IVFFlat 的构建阶段：

1. `initializing`
2. `performing k-means`
3. `assigning tuples`
4. `loading tuples`

注意：`%` 仅在 `loading tuples` 阶段显示

### 过滤

有多种方式可以对带有 `WHERE` 子句的最近邻查询建立索引。

```sql
SELECT * FROM items WHERE category_id = 123 ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

首先可以在过滤列上创建索引。在很多场景下，这可以提供快速的精确最近邻搜索。PostgreSQL 为此提供了多种[索引类型](https://www.postgresql.org/docs/current/indexes-types.html)：B-tree（默认）、hash、GiST、SP-GiST、GIN 和 BRIN。

```sql
CREATE INDEX ON items (category_id);
```

对于多列过滤，可以考虑使用[多列索引](https://www.postgresql.org/docs/current/indexes-multicolumn.html)。

```sql
CREATE INDEX ON items (location_id, category_id);
```

精确索引适用于匹配行数占比较低的条件。否则，[近似索引](https://github.com/pgvector/pgvector#indexing)可能效果更好。

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops);
```

使用近似索引时，过滤是在索引扫描*之后*应用的。如果某个条件匹配 10% 的行，在 HNSW 默认 `hnsw.ef_search` 为 40 的情况下，平均只有 4 行会匹配。要获取更多行，请增大 `hnsw.ef_search`。

```sql
SET hnsw.ef_search = 200;
```

从 0.8.0 版本开始，可以启用[迭代索引扫描](https://github.com/pgvector/pgvector#iterative-index-scans)，在需要时自动扫描索引的更多部分。

```sql
SET hnsw.iterative_scan = strict_order;
```

如果只按少数几个不同的值过滤，可以考虑使用[部分索引](https://www.postgresql.org/docs/current/indexes-partial.html)。

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops) WHERE (category_id = 123);
```

如果按大量不同的值过滤，可以考虑使用[分区表](https://www.postgresql.org/docs/current/ddl-partitioning.html)。

```sql
CREATE TABLE items (embedding vector(3), category_id int) PARTITION BY LIST(category_id);
```

### 迭代索引扫描

使用近似索引时，带有过滤条件的查询可能返回较少的结果，因为过滤是在索引扫描*之后*应用的。从 0.8.0 版本开始，可以启用迭代索引扫描，它会自动扫描索引的更多部分，直到找到足够的结果（或达到 `hnsw.max_scan_tuples` 或 `ivfflat.max_probes` 的限制）。

迭代扫描可以使用严格排序或宽松排序。

严格排序确保结果按距离精确排序

```sql
SET hnsw.iterative_scan = strict_order;
```

宽松排序允许结果在距离排序上略有偏差，但能提供更好的召回率

```sql
SET hnsw.iterative_scan = relaxed_order;
# or
SET ivfflat.iterative_scan = relaxed_order;
```

使用宽松排序时，可以通过[物化 CTE](https://www.postgresql.org/docs/current/queries-with.html#QUERIES-WITH-CTE-MATERIALIZATION) 来获得严格排序

```sql
WITH relaxed_results AS MATERIALIZED (
    SELECT id, embedding <-> '[1,2,3]' AS distance FROM items WHERE category_id = 123 ORDER BY distance LIMIT 5
) SELECT * FROM relaxed_results ORDER BY distance + 0;
```

注意：PostgreSQL 17+ 需要 `+ 0`

对于按距离过滤的查询，请使用物化 CTE 并将距离过滤条件放在 CTE 外部以获得最佳性能（这是由于 PostgreSQL 执行器的[当前行为](https://www.postgresql.org/message-id/flat/CAOdR5yGUoMQ6j7M5hNUXrySzaqZVGf_Ne%2B8fwZMRKTFxU1nbJg%40mail.gmail.com)所致）

```sql
WITH nearest_results AS MATERIALIZED (
    SELECT id, embedding <-> '[1,2,3]' AS distance FROM items ORDER BY distance LIMIT 5
) SELECT * FROM nearest_results WHERE distance < 5 ORDER BY distance;
```

注意：其他过滤条件请放在 CTE 内部

#### 迭代扫描选项

由于扫描近似索引的大部分内容开销很大，因此提供了控制扫描何时终止的选项。

##### HNSW

指定最大访问元组数（默认 20,000）

```sql
SET hnsw.max_scan_tuples = 20000;
```

注意：该值为近似值，不影响初始扫描

指定最大可用内存量，以 `work_mem` 的倍数表示（默认 1）

```sql
SET hnsw.scan_mem_multiplier = 2;
```

注意：如果增大 `hnsw.max_scan_tuples` 未能改善召回率，可尝试增大此值

##### IVFFlat

指定最大探针数

```sql
SET ivfflat.max_probes = 100;
```

注意：如果该值低于 `ivfflat.probes`，则使用 `ivfflat.probes` 的值


--------

## 向量类型

### 半精度向量

使用 `halfvec` 类型存储半精度向量

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding halfvec(3));
```

### 半精度索引

以半精度对向量建立索引，可以获得更小的索引体积

```sql
CREATE INDEX ON items USING hnsw ((embedding::halfvec(3)) halfvec_l2_ops);
```

获取最近邻

```sql
SELECT * FROM items ORDER BY embedding::halfvec(3) <-> '[1,2,3]' LIMIT 5;
```

### 二进制向量

使用 `bit` 类型存储二进制向量（[示例](https://github.com/pgvector/pgvector-python/blob/master/examples/imagehash/example.py)）

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding bit(3));
INSERT INTO items (embedding) VALUES ('000'), ('111');
```

按汉明距离获取最近邻

```sql
SELECT * FROM items ORDER BY embedding <~> '101' LIMIT 5;
```

还支持 Jaccard 距离（`<%>`）

### 二进制量化

使用表达式索引进行二进制量化

```sql
CREATE INDEX ON items USING hnsw ((binary_quantize(embedding)::bit(3)) bit_hamming_ops);
```

按汉明距离获取最近邻

```sql
SELECT * FROM items ORDER BY binary_quantize(embedding)::bit(3) <~> binary_quantize('[1,-2,3]') LIMIT 5;
```

使用原始向量重新排序以提高召回率

```sql
SELECT * FROM (
    SELECT * FROM items ORDER BY binary_quantize(embedding)::bit(3) <~> binary_quantize('[1,-2,3]') LIMIT 20
) ORDER BY embedding <=> '[1,-2,3]' LIMIT 5;
```

### 稀疏向量

使用 `sparsevec` 类型存储稀疏向量

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding sparsevec(5));
```

插入向量

```sql
INSERT INTO items (embedding) VALUES ('{1:1,3:2,5:3}/5'), ('{1:4,3:5,5:6}/5');
```

格式为 `{索引1:值1,索引2:值2}/维度数`，索引从 1 开始，与 SQL 数组一致

按 L2 距离获取最近邻

```sql
SELECT * FROM items ORDER BY embedding <-> '{1:3,3:1,5:2}/5' LIMIT 5;
```

### 混合搜索

结合 PostgreSQL 的[全文搜索](https://www.postgresql.org/docs/current/textsearch-intro.html)实现混合搜索。

```sql
SELECT id, content FROM items, plainto_tsquery('hello search') query
    WHERE textsearch @@ query ORDER BY ts_rank_cd(textsearch, query) DESC LIMIT 5;
```

可以使用[倒数排名融合（RRF）](https://github.com/pgvector/pgvector-python/blob/master/examples/hybrid_search/rrf.py)或[交叉编码器](https://github.com/pgvector/pgvector-python/blob/master/examples/hybrid_search/cross_encoder.py)来合并结果。

### 子向量索引

使用表达式索引对子向量建立索引

```sql
CREATE INDEX ON items USING hnsw ((subvector(embedding, 1, 3)::vector(3)) vector_cosine_ops);
```

按余弦距离获取最近邻

```sql
SELECT * FROM items ORDER BY subvector(embedding, 1, 3)::vector(3) <=> subvector('[1,2,3,4,5]'::vector, 1, 3) LIMIT 5;
```

使用完整向量重新排序以提高召回率

```sql
SELECT * FROM (
    SELECT * FROM items ORDER BY subvector(embedding, 1, 3)::vector(3) <=> subvector('[1,2,3,4,5]'::vector, 1, 3) LIMIT 20
) ORDER BY embedding <=> '[1,2,3,4,5]' LIMIT 5;
```


--------

## 性能

### 调优

使用 [PgTune](https://pgtune.leopard.in.ua/) 等工具来设置 PostgreSQL 服务器参数的初始值。例如，`shared_buffers` 通常应设置为服务器内存的 25%。可以通过以下命令查找配置文件路径：

```sql
SHOW config_file;
```

查看各项参数设置：

```sql
SHOW shared_buffers;
```

修改后需要重启 PostgreSQL 才能生效。

### 数据加载

使用 `COPY` 批量加载数据（[示例](https://github.com/pgvector/pgvector-python/blob/master/examples/loading/example.py)）。

```sql
COPY items (embedding) FROM STDIN WITH (FORMAT BINARY);
```

为获得最佳性能，建议在加载完初始数据*之后*再创建索引。

### 索引

关于索引构建时间，参见 [HNSW](https://github.com/pgvector/pgvector#index-build-time) 和 [IVFFlat](https://github.com/pgvector/pgvector#index-build-time-1)。

在生产环境中，请使用并发方式创建索引以避免阻塞写操作。

```sql
CREATE INDEX CONCURRENTLY ...
```

### 查询

使用 `EXPLAIN (ANALYZE, BUFFERS)` 调试查询性能。

```sql
EXPLAIN (ANALYZE, BUFFERS) SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

#### 精确搜索

对于没有索引的查询，可以增大 `max_parallel_workers_per_gather` 来加速。

```sql
SET max_parallel_workers_per_gather = 4;
```

如果向量已归一化为单位长度（如 [OpenAI embeddings](https://platform.openai.com/docs/guides/embeddings/which-distance-function-should-i-use)），使用内积可获得最佳性能。

```tsql
SELECT * FROM items ORDER BY embedding <#> '[3,1,2]' LIMIT 5;
```

#### 近似搜索

对于使用 IVFFlat 索引的查询，可以增加倒排列表数量来加速（会牺牲部分召回率）。

```sql
CREATE INDEX ON items USING ivfflat (embedding vector_l2_ops) WITH (lists = 1000);
```

### 清理（VACUUM）

HNSW 索引的 VACUUM 操作可能耗时较长。可以先重建索引来加速。

```sql
REINDEX INDEX CONCURRENTLY index_name;
VACUUM table_name;
```

### 监控

使用 [pg_stat_statements](https://www.postgresql.org/docs/current/pgstatstatements.html) 监控性能（需确保已将其添加到 `shared_preload_libraries`）。

```sql
CREATE EXTENSION pg_stat_statements;
```

查询最耗时的 SQL：

```sql
SELECT query, calls, ROUND((total_plan_time + total_exec_time) / calls) AS avg_time_ms,
    ROUND((total_plan_time + total_exec_time) / 60000) AS total_time_min
    FROM pg_stat_statements ORDER BY total_plan_time + total_exec_time DESC LIMIT 20;
```

通过对比近似搜索与精确搜索的结果来监控召回率。

```sql
BEGIN;
SET LOCAL enable_indexscan = off; -- 使用精确搜索
SELECT ...
COMMIT;
```

### 扩展

pgvector 的扩展方式与 PostgreSQL 相同。

纵向扩展：增加单个实例的内存、CPU 和存储容量。使用现有工具进行[参数调优](https://github.com/pgvector/pgvector#tuning)和[性能监控](https://github.com/pgvector/pgvector#monitoring)。

横向扩展：使用[只读副本](https://www.postgresql.org/docs/current/hot-standby.html)，或使用 [Citus](https://github.com/citusdata/citus) 等方案进行分片（[示例](https://github.com/pgvector/pgvector-python/blob/master/examples/citus/example.py)）。


--------

## 语言支持

可以从任何具有 PostgreSQL 客户端的编程语言中使用 pgvector。甚至可以在一种语言中生成和存储向量，在另一种语言中进行查询。

语言 | 库 / 示例
--- | ---
Ada | [pgvector-ada](https://github.com/pgvector/pgvector-ada)
Algol | [pgvector-algol](https://github.com/pgvector/pgvector-algol)
C | [pgvector-c](https://github.com/pgvector/pgvector-c)
C++ | [pgvector-cpp](https://github.com/pgvector/pgvector-cpp)
C#, F#, Visual Basic | [pgvector-dotnet](https://github.com/pgvector/pgvector-dotnet)
COBOL | [pgvector-cobol](https://github.com/pgvector/pgvector-cobol)
Crystal | [pgvector-crystal](https://github.com/pgvector/pgvector-crystal)
D | [pgvector-d](https://github.com/pgvector/pgvector-d)
Dart | [pgvector-dart](https://github.com/pgvector/pgvector-dart)
Elixir | [pgvector-elixir](https://github.com/pgvector/pgvector-elixir)
Erlang | [pgvector-erlang](https://github.com/pgvector/pgvector-erlang)
Fortran | [pgvector-fortran](https://github.com/pgvector/pgvector-fortran)
Gleam | [pgvector-gleam](https://github.com/pgvector/pgvector-gleam)
Go | [pgvector-go](https://github.com/pgvector/pgvector-go)
Haskell | [pgvector-haskell](https://github.com/pgvector/pgvector-haskell)
Java, Kotlin, Groovy, Scala | [pgvector-java](https://github.com/pgvector/pgvector-java)
JavaScript, TypeScript | [pgvector-node](https://github.com/pgvector/pgvector-node)
Julia | [Pgvector.jl](https://github.com/pgvector/Pgvector.jl)
Lisp | [pgvector-lisp](https://github.com/pgvector/pgvector-lisp)
Lua | [pgvector-lua](https://github.com/pgvector/pgvector-lua)
Nim | [pgvector-nim](https://github.com/pgvector/pgvector-nim)
OCaml | [pgvector-ocaml](https://github.com/pgvector/pgvector-ocaml)
Pascal | [pgvector-pascal](https://github.com/pgvector/pgvector-pascal)
Perl | [pgvector-perl](https://github.com/pgvector/pgvector-perl)
PHP | [pgvector-php](https://github.com/pgvector/pgvector-php)
Prolog | [pgvector-prolog](https://github.com/pgvector/pgvector-prolog)
Python | [pgvector-python](https://github.com/pgvector/pgvector-python)
R | [pgvector-r](https://github.com/pgvector/pgvector-r)
Racket | [pgvector-racket](https://github.com/pgvector/pgvector-racket)
Raku | [pgvector-raku](https://github.com/pgvector/pgvector-raku)
Ruby | [pgvector-ruby](https://github.com/pgvector/pgvector-ruby), [Neighbor](https://github.com/ankane/neighbor)
Rust | [pgvector-rust](https://github.com/pgvector/pgvector-rust)
Swift | [pgvector-swift](https://github.com/pgvector/pgvector-swift)
Tcl | [pgvector-tcl](https://github.com/pgvector/pgvector-tcl)
Zig | [pgvector-zig](https://github.com/pgvector/pgvector-zig)


--------

## 常见问题

### 单张表最多能存储多少向量？

PostgreSQL 中非分区表默认上限为 32 TB。分区表可以有数千个该大小的分区。

### 是否支持复制？

是的，pgvector 使用预写式日志（WAL），支持流复制和时间点恢复（PITR）。

### 如果我想对超过 2,000 维的向量建立索引怎么办？

可以使用[半精度向量](https://github.com/pgvector/pgvector#half-precision-vectors)或[半精度索引](https://github.com/pgvector/pgvector#half-precision-indexing)来索引最多 4,000 维的向量，或使用[二进制量化](https://github.com/pgvector/pgvector#binary-quantization)来索引最多 64,000 维的向量。其他方案包括[子向量索引](https://github.com/pgvector/pgvector#indexing-subvectors)（需要模型支持）或[降维](https://en.wikipedia.org/wiki/Dimensionality_reduction)。

### 能否在同一列中存储不同维度的向量？

可以使用 `vector` 作为类型（而非 `vector(n)`）。

```sql
CREATE TABLE embeddings (model_id bigint, item_id bigint, embedding vector, PRIMARY KEY (model_id, item_id));
```

但只能对维度相同的行创建索引（使用[表达式索引](https://www.postgresql.org/docs/current/indexes-expressional.html)和[部分索引](https://www.postgresql.org/docs/current/indexes-partial.html)）：

```sql
CREATE INDEX ON embeddings USING hnsw ((embedding::vector(3)) vector_l2_ops) WHERE (model_id = 123);
```

查询方式：

```sql
SELECT * FROM embeddings WHERE model_id = 123 ORDER BY embedding::vector(3) <-> '[3,1,2]' LIMIT 5;
```

### 能否以更高精度存储向量？

可以使用 `double precision[]` 或 `numeric[]` 类型来存储更高精度的向量。

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding double precision[]);

-- PostgreSQL 数组使用 {} 而非 []
INSERT INTO items (embedding) VALUES ('{1,2,3}'), ('{4,5,6}');
```

可以选择添加[检查约束](https://www.postgresql.org/docs/current/ddl-constraints.html)，确保数据可以转换为 `vector` 类型并具有预期的维度。

```sql
ALTER TABLE items ADD CHECK (vector_dims(embedding::vector) = 3);
```

使用[表达式索引](https://www.postgresql.org/docs/current/indexes-expressional.html)建立索引（以较低精度）：

```sql
CREATE INDEX ON items USING hnsw ((embedding::vector(3)) vector_l2_ops);
```

查询方式：

```sql
SELECT * FROM items ORDER BY embedding::vector(3) <-> '[3,1,2]' LIMIT 5;
```

### 索引是否需要完全放入内存？

不需要，但与其他索引类型一样，索引能放入内存时性能通常更好。可以通过以下命令查看索引大小：

```sql
SELECT pg_size_pretty(pg_relation_size('index_name'));
```


--------

## 故障排查

### 为什么查询没有使用索引？

查询需要同时包含 `ORDER BY` 和 `LIMIT`，且 `ORDER BY` 必须直接使用距离操作符的结果（不能是表达式），并且必须为升序排列。

```sql
-- 使用索引
ORDER BY embedding <=> '[3,1,2]' LIMIT 5;

-- 不使用索引
ORDER BY 1 - (embedding <=> '[3,1,2]') DESC LIMIT 5;
```

可以通过以下方式引导查询规划器使用索引：

```sql
BEGIN;
SET LOCAL enable_seqscan = off;
SELECT ...
COMMIT;
```

另外，如果表很小，全表扫描可能反而更快。

### 为什么查询没有使用并行表扫描？

查询规划器在成本估算中不考虑[行外存储（TOAST）](https://www.postgresql.org/docs/current/storage-toast.html)，这可能导致串行扫描看起来成本更低。可以通过以下方式降低并行扫描的成本估算：

```sql
BEGIN;
SET LOCAL min_parallel_table_scan_size = 1;
SET LOCAL parallel_setup_cost = 1;
SELECT ...
COMMIT;
```

或者选择以内联方式存储向量：

```sql
ALTER TABLE items ALTER COLUMN embedding SET STORAGE PLAIN;
```

### 为什么添加 HNSW 索引后查询结果变少了？

结果数量受动态候选列表大小（`hnsw.ef_search`）的限制，默认值为 40。由于死元组或查询中的过滤条件，结果可能会更少。启用[迭代索引扫描](https://github.com/pgvector/pgvector#iterative-index-scans)有助于解决此问题。

另外请注意，`NULL` 向量不会被索引（余弦距离的零向量也不会被索引）。

### 为什么添加 IVFFlat 索引后查询结果变少了？

很可能是在数据量相对于列表数量不足时就创建了索引。请先删除索引，等表中有更多数据后再重新创建。

```sql
DROP INDEX index_name;
```

结果数量也可能受探针数（`ivfflat.probes`）的限制。启用[迭代索引扫描](https://github.com/pgvector/pgvector#iterative-index-scans)可以解决此问题。

另外请注意，`NULL` 向量不会被索引（余弦距离的零向量也不会被索引）。


--------

## 参考

### Vector 类型

每个向量占用 `4 * 维度数 + 8` 字节的存储空间。每个元素为单精度浮点数（类似于 PostgreSQL 中的 `real` 类型），所有元素必须为有限值（不允许 `NaN`、`Infinity` 或 `-Infinity`）。向量最多支持 16,000 维。

### Vector 操作符

操作符 | 描述 | 引入版本
--- | --- | ---
\+ | 逐元素加法 |
\- | 逐元素减法 |
\* | 逐元素乘法 | 0.5.0
\|\| | 拼接 | 0.7.0
<-> | 欧氏距离 |
<#> | 负内积 |
<=> | 余弦距离 |
<+> | 曼哈顿距离 | 0.7.0

### Vector 函数

函数 | 描述 | 引入版本
--- | --- | ---
binary_quantize(vector) → bit | 二进制量化 | 0.7.0
cosine_distance(vector, vector) → double precision | 余弦距离 |
inner_product(vector, vector) → double precision | 内积 |
l1_distance(vector, vector) → double precision | 曼哈顿距离 | 0.5.0
l2_distance(vector, vector) → double precision | 欧氏距离 |
l2_normalize(vector) → vector | 欧氏范数归一化 | 0.7.0
subvector(vector, integer, integer) → vector | 子向量 | 0.7.0
vector_dims(vector) → integer | 维度数 |
vector_norm(vector) → double precision | 欧氏范数 |

### Vector 聚合函数

函数 | 描述 | 引入版本
--- | --- | ---
avg(vector) → vector | 平均值 |
sum(vector) → vector | 求和 | 0.5.0

### Halfvec 类型

每个半精度向量占用 `2 * 维度数 + 8` 字节的存储空间。每个元素为半精度浮点数，所有元素必须为有限值（不允许 `NaN`、`Infinity` 或 `-Infinity`）。半精度向量最多支持 16,000 维。

### Halfvec 操作符

操作符 | 描述 | 引入版本
--- | --- | ---
\+ | 逐元素加法 | 0.7.0
\- | 逐元素减法 | 0.7.0
\* | 逐元素乘法 | 0.7.0
\|\| | 拼接 | 0.7.0
<-> | 欧氏距离 | 0.7.0
<#> | 负内积 | 0.7.0
<=> | 余弦距离 | 0.7.0
<+> | 曼哈顿距离 | 0.7.0

### Halfvec 函数

函数 | 描述 | 引入版本
--- | --- | ---
binary_quantize(halfvec) → bit | 二进制量化 | 0.7.0
cosine_distance(halfvec, halfvec) → double precision | 余弦距离 | 0.7.0
inner_product(halfvec, halfvec) → double precision | 内积 | 0.7.0
l1_distance(halfvec, halfvec) → double precision | 曼哈顿距离 | 0.7.0
l2_distance(halfvec, halfvec) → double precision | 欧氏距离 | 0.7.0
l2_norm(halfvec) → double precision | 欧氏范数 | 0.7.0
l2_normalize(halfvec) → halfvec | 欧氏范数归一化 | 0.7.0
subvector(halfvec, integer, integer) → halfvec | 子向量 | 0.7.0
vector_dims(halfvec) → integer | 维度数 | 0.7.0

### Halfvec 聚合函数

函数 | 描述 | 引入版本
--- | --- | ---
avg(halfvec) → halfvec | 平均值 | 0.7.0
sum(halfvec) → halfvec | 求和 | 0.7.0

### Bit 类型

每个位向量占用 `维度数 / 8 + 8` 字节的存储空间。更多信息参见 [PostgreSQL 文档](https://www.postgresql.org/docs/current/datatype-bit.html)。

### Bit 操作符

操作符 | 描述 | 引入版本
--- | --- | ---
<~> | 汉明距离 | 0.7.0
<%> | Jaccard 距离 | 0.7.0

### Bit 函数

函数 | 描述 | 引入版本
--- | --- | ---
hamming_distance(bit, bit) → double precision | 汉明距离 | 0.7.0
jaccard_distance(bit, bit) → double precision | Jaccard 距离 | 0.7.0

### Sparsevec 类型

每个稀疏向量占用 `8 * 非零元素数 + 16` 字节的存储空间。每个元素为单精度浮点数，所有元素必须为有限值（不允许 `NaN`、`Infinity` 或 `-Infinity`）。稀疏向量最多支持 16,000 个非零元素。

### Sparsevec 操作符

操作符 | 描述 | 引入版本
--- | --- | ---
<-> | 欧氏距离 | 0.7.0
<#> | 负内积 | 0.7.0
<=> | 余弦距离 | 0.7.0
<+> | 曼哈顿距离 | 0.7.0

### Sparsevec 函数

函数 | 描述 | 引入版本
--- | --- | ---
cosine_distance(sparsevec, sparsevec) → double precision | 余弦距离 | 0.7.0
inner_product(sparsevec, sparsevec) → double precision | 内积 | 0.7.0
l1_distance(sparsevec, sparsevec) → double precision | 曼哈顿距离 | 0.7.0
l2_distance(sparsevec, sparsevec) → double precision | 欧氏距离 | 0.7.0
l2_norm(sparsevec) → double precision | 欧氏范数 | 0.7.0
l2_normalize(sparsevec) → sparsevec | 欧氏范数归一化 | 0.7.0


--------

## 安装

### Linux 与 Mac

编译安装扩展（支持 PostgreSQL 13+）

```sh
cd /tmp
git clone --branch v0.8.2 https://github.com/pgvector/pgvector.git
cd pgvector
make
make install # 可能需要 sudo
```

如果遇到问题，请参阅[安装说明](https://github.com/pgvector/pgvector#installation-notes---linux-and-mac)

也可以通过 [Docker](https://github.com/pgvector/pgvector#docker)、[Homebrew](https://github.com/pgvector/pgvector#homebrew)、[PGXN](https://github.com/pgvector/pgvector#pgxn)、[APT](https://github.com/pgvector/pgvector#apt)、[Yum](https://github.com/pgvector/pgvector#yum)、[pkg](https://github.com/pgvector/pgvector#pkg)、[APK](https://github.com/pgvector/pgvector#apk) 或 [conda-forge](https://github.com/pgvector/pgvector#conda-forge) 安装，[Postgres.app](https://github.com/pgvector/pgvector#postgresapp) 和许多[托管服务商](https://github.com/pgvector/pgvector#hosted-postgres)已预装此扩展。此外还有 [GitHub Actions](https://github.com/pgvector/setup-pgvector) 的使用说明。

### Windows

确保已安装 [Visual Studio 中的 C++ 支持](https://learn.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-170#download-and-install-the-tools)，并以管理员身份运行 `x64 Native Tools Command Prompt for VS [版本]`。然后使用 `nmake` 编译：

```cmd
set "PGROOT=C:\Program Files\PostgreSQL\18"
cd %TEMP%
git clone --branch v0.8.2 https://github.com/pgvector/pgvector.git
cd pgvector
nmake /F Makefile.win
nmake /F Makefile.win install
```

如果遇到问题，请参阅[安装说明](https://github.com/pgvector/pgvector#installation-notes---windows)

也可以通过 [Docker](https://github.com/pgvector/pgvector#docker) 或 [conda-forge](https://github.com/pgvector/pgvector#conda-forge) 安装。

### 安装说明 - Linux 与 Mac

#### PostgreSQL 路径

如果机器上有多个 PostgreSQL 安装，请通过以下方式指定 [pg_config](https://www.postgresql.org/docs/current/app-pgconfig.html) 的路径：

```sh
export PG_CONFIG=/Library/PostgreSQL/18/bin/pg_config
```

然后重新执行安装步骤（如需要，在 `make` 前先执行 `make clean`）。如果 `make install` 需要 `sudo`，请使用：

```sh
sudo --preserve-env=PG_CONFIG make install
```

Mac 上常见的路径包括：

- EDB 安装器 - `/Library/PostgreSQL/18/bin/pg_config`
- Homebrew (arm64) - `/opt/homebrew/opt/postgresql@18/bin/pg_config`
- Homebrew (x86-64) - `/usr/local/opt/postgresql@18/bin/pg_config`

注意：请将 `18` 替换为实际的 PostgreSQL 服务器版本

#### 缺少头文件

如果编译报错 `fatal error: postgres.h: No such file or directory`，请确保服务器上已安装 PostgreSQL 开发文件。

对于 Ubuntu 和 Debian，使用：

```sh
sudo apt install postgresql-server-dev-18
```

注意：请将 `18` 替换为实际的 PostgreSQL 服务器版本

#### 缺少 SDK

如果在 Mac 上编译失败且输出包含 `warning: no such sysroot directory`，说明 PostgreSQL 安装指向了一个已不存在的路径。

```sh
pg_config --cppflags
```

重新安装 PostgreSQL 即可修复此问题。

#### 可移植性

默认情况下，pgvector 在某些平台上会使用 `-march=native` 编译以获得最佳性能。但如果要在其他机器上运行编译好的扩展，这可能导致 `Illegal instruction` 错误。

要以可移植方式编译，请使用：

```sh
make OPTFLAGS=""
```

### 安装说明 - Windows

#### 缺少头文件

如果编译报错 `Cannot open include file: 'postgres.h': No such file or directory`，请确保 `PGROOT` 设置正确。

#### 架构不匹配

如果编译报错 `error C2196: case value '4' already used`，请确保使用的是 `x64 Native Tools Command Prompt`。然后执行 `nmake /F Makefile.win clean` 并重新执行安装步骤。

#### 缺少符号

如果在 PostgreSQL 17.0-17.2 版本上链接失败并提示 `unresolved external symbol float_to_shortest_decimal_bufn`，请升级到 PostgreSQL 17.3+。

#### 权限问题

如果安装失败并提示 `Access is denied`，请以管理员身份重新执行安装步骤。

### 其他安装方式

#### Docker

获取 [Docker 镜像](https://hub.docker.com/r/pgvector/pgvector)：

```sh
docker pull pgvector/pgvector:pg18-trixie
```

该镜像在 [PostgreSQL 镜像](https://hub.docker.com/_/postgres)的基础上添加了 pgvector（请将 `18` 替换为实际的 PostgreSQL 服务器版本，使用方式相同）。

也可以手动构建镜像：

```sh
git clone --branch v0.8.2 https://github.com/pgvector/pgvector.git
cd pgvector
docker build --pull --build-arg PG_MAJOR=18 -t myuser/pgvector .
```

如果增大了 `maintenance_work_mem`，请确保 `--shm-size` 至少为该值，以避免并行 HNSW 索引构建时出错。

```sh
docker run --shm-size=1g ...
```

#### Homebrew

对于 Homebrew 安装的 PostgreSQL，可以使用：

```sh
brew install pgvector
```

注意：仅适用于 `postgresql@18` 和 `postgresql@17`

#### PGXN

从 [PostgreSQL 扩展网络](https://pgxn.org/dist/vector)安装：

```sh
pgxn install vector
```

#### APT

Debian 和 Ubuntu 软件包可从 [PostgreSQL APT 仓库](https://wiki.postgresql.org/wiki/Apt)获取。按照[配置说明](https://wiki.postgresql.org/wiki/Apt#Quickstart)操作后执行：

```sh
sudo apt install postgresql-18-pgvector
```

注意：请将 `18` 替换为实际的 PostgreSQL 服务器版本

#### Yum

RPM 软件包可从 [PostgreSQL Yum 仓库](https://yum.postgresql.org/)获取。按照对应发行版的[配置说明](https://www.postgresql.org/download/linux/redhat/)操作后执行：

```sh
sudo yum install pgvector_18
# 或
sudo dnf install pgvector_18
```

注意：请将 `18` 替换为实际的 PostgreSQL 服务器版本

#### pkg

安装 FreeBSD 软件包：

```sh
pkg install postgresql17-pgvector
```

或通过 port 安装：

```sh
cd /usr/ports/databases/pgvector
make install
```

#### APK

安装 Alpine 软件包：

```sh
apk add postgresql-pgvector
```

#### conda-forge

对于 Conda 安装的 PostgreSQL，从 [conda-forge](https://anaconda.org/conda-forge/pgvector) 安装：

```sh
conda install -c conda-forge pgvector
```

该方式由 [@mmcauliffe](https://github.com/mmcauliffe) [社区维护](https://github.com/conda-forge/pgvector-feedstock)

#### Postgres.app

下载 Postgres 15+ 的[最新版本](https://postgresapp.com/downloads.html)。

### 托管 PostgreSQL

pgvector 可在[这些服务商](https://github.com/pgvector/pgvector/issues/54)上使用。

### 升级

[安装](https://github.com/pgvector/pgvector#installation)最新版本（使用与初始安装相同的方式）。然后在每个需要升级的数据库中执行：

```sql
ALTER EXTENSION vector UPDATE;
```

可以通过以下命令查看当前数据库中的版本：

```sql
SELECT extversion FROM pg_extension WHERE extname = 'vector';
```
