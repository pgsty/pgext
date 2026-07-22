## 用法

来源：

- [官方扩展控制文件](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/pg_knowledge_graph.control)
- [官方 README](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/README.md)
- [官方模式 SQL](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/sql/pg_knowledge_graph--0.1.0.sql)

`pg_knowledge_graph` 0.1.0 使用普通 PostgreSQL 表保存实体以及有向加权关系，并提供遍历、排名、连通分量、向量检索、混合检索、RAG 上下文和近似量化检索函数。它依赖 `vector`，上游模式固定使用 1536 维嵌入。

### 核心流程

先安装 pgvector。pgrx 安装会提供扩展函数，但仓库自己的 Docker 初始化器明确说明图表需要单独创建。插入数据前，应执行所引模式 SQL，或显式创建下面两张表。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_knowledge_graph;

CREATE TABLE kg_entities (
    id bigserial PRIMARY KEY,
    entity_type text NOT NULL,
    name text NOT NULL,
    properties jsonb NOT NULL DEFAULT '{}',
    embedding vector(1536),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE kg_relations (
    id bigserial PRIMARY KEY,
    source_id bigint NOT NULL REFERENCES kg_entities(id) ON DELETE CASCADE,
    target_id bigint NOT NULL REFERENCES kg_entities(id) ON DELETE CASCADE,
    rel_type text NOT NULL,
    weight float8 NOT NULL DEFAULT 1.0,
    properties jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO kg_entities (entity_type, name)
VALUES ('person', 'Alice'), ('person', 'Bob');

INSERT INTO kg_relations (source_id, target_id, rel_type, weight)
VALUES (1, 2, 'knows', 1.0);

SELECT * FROM kg_bfs(1, 2);
SELECT kg_shortest_path(1, 2, 5);
SELECT * FROM kg_pagerank(0.85, 100) ORDER BY score DESC;
```

大规模使用向量检索前，还应创建官方模式 SQL 中的索引，包括 `embedding` 上的 HNSW 索引。

### 函数组

- 检查：`kg_version` 和 `kg_stats`。
- 遍历：沿有向边工作的 `kg_bfs`、`kg_dfs` 和 `kg_shortest_path`。
- 全图算法：`kg_pagerank`、`kg_louvain`、`kg_connected_components` 和 `kg_strongly_connected_components`。
- 检索：`kg_vector_search`、`kg_hybrid_search` 和 `kg_get_context`。
- 近似检索：支持 `int8`、`int4` 或 `binary` 的 `kg_quantized_search`，以及 `kg_quantize_info`。

`kg_hybrid_search` 使用调用者提供的 `alpha` 与 `beta` 组合向量分数和图分数。`kg_get_context` 返回用于检索增强生成的 N 跳邻域，但不会调用嵌入服务或语言模型服务。

### 运维边界

Rust 函数通过 SPI 查询 `kg_entities` 与 `kg_relations`，部分算法会把图或嵌入数据加载到调用后端。PageRank、连通分量、社区检测与量化检索均按请求重新计算，而不是增量维护。应针对真实图规模测量内存、执行时间和语句超时。

函数使用未限定模式的表名，因此官方表必须位于预期搜索路径中，并应防止不受信用户替换成同名关系。根据控制文件，该扩展只能由超级用户安装、不受信且不可重定位。

量化相似度是近似值。README 中的压缩率与召回率是项目测量结果，不是正确性保证；应使用自己的嵌入评估召回率。版本 0.1.0 声明了 PostgreSQL 16、17 与 18 的构建特性，在完成独立负载和故障测试前，应把它视为早期版本。
