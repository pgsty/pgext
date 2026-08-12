## 用法

来源：

- [pg_search v0.25.2 README](https://github.com/paradedb/paradedb/blob/v0.25.2/pg_search/README.md)
- [pg_search v0.25.2 发行说明](https://github.com/paradedb/paradedb/releases/tag/v0.25.2)
- [pg_search v0.25.2 变更日志](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/changelog/0.25.2.mdx)
- [pg_search v0.25.1 迁移说明](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/changelog/0.25.1.mdx)
- [创建 ParadeDB 索引](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/indexing/create-index.mdx)
- [全文匹配操作符](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/full-text/match.mdx)
- [BM25 评分](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/sorting/score.mdx)
- [高亮与摘要](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/full-text/highlight.mdx)
- [索引向量](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/indexing/indexing-vectors.mdx)
- [查询向量](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/vector/querying.mdx)
- [混合搜索概述](https://github.com/paradedb/paradedb/blob/v0.25.2/docs/documentation/hybrid/overview.mdx)

`pg_search` 0.25.2 为 PostgreSQL 增加 ParadeDB 的全文、结构化、向量和混合搜索索引。版本 0.25 使用 `paradedb` 索引访问方法；旧的 `bm25` 访问方法名称仍保留为兼容别名。该扩展依赖 `vector`，上游支持 PostgreSQL 15-18，且必须通过 `shared_preload_libraries` 加载。

### 安装并构建索引

```conf
shared_preload_libraries = 'pg_search'
```

重启 PostgreSQL，然后创建扩展以及一张拥有稳定唯一键的表：

```sql
CREATE EXTENSION pg_search CASCADE;

CREATE TABLE documents (
  id          bigint PRIMARY KEY,
  title       text,
  body        text,
  category    text,
  embedding   vector(768)
);

CREATE INDEX documents_search_idx ON documents
USING paradedb (
  id,
  title,
  body,
  category,
  embedding vector_cosine_ops
)
WITH (key_field = 'id');
```

`key_field` 必须是第一个索引列，并能唯一标识每一行。文本键必须以不分词的方式建立索引。一张表只能拥有一个 ParadeDB 索引，因此应将所有需要搜索的字段都纳入该索引。

### 全文搜索

使用 `|||` 匹配任意词元，使用 `&&&` 要求匹配所有词元：

```sql
SELECT id, title, pdb.score(id) AS score
FROM documents
WHERE body ||| 'postgresql search'
ORDER BY score DESC, id;

SELECT id, pdb.snippet(body) AS excerpt
FROM documents
WHERE body &&& 'postgresql indexing';
```

`pdb.score(key_field)` 提供当前行的相关性分数。`pdb.snippet(indexed_text_column)` 返回高亮摘要。这些辅助函数只有在 ParadeDB 搜索谓词驱动的查询中才有意义。

### 向量搜索

向量索引在 0.25 系列中处于 beta 阶段，使用 pgvector 的 `vector` 类型。创建索引时应选择操作符类；更改距离度量需要重建索引。

```sql
SELECT id, title, embedding <=> $1::vector AS distance
FROM documents
WHERE id @@@ pdb.all()
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

支持的索引操作符类为 `vector_l2_ops`、`vector_ip_ops` 和 `vector_cosine_ops`。0.25 向量索引不为 `halfvec`、`sparsevec` 或 `bit` 列建立索引。

### 混合搜索

单个 ParadeDB 索引可以组合词法谓词、结构化过滤与向量排序。需要更复杂的融合时，应使用文档中的 RRF 和加权混合搜索函数，而不是直接将量纲不同的分数相加。

```sql
SELECT id, title, pdb.score(id) AS lexical_score
FROM documents
WHERE body ||| 'postgresql extension'
  AND category === 'database'
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

### 版本 0.25.2 与注意事项

- 版本 0.25 将主要索引访问方法从 `bm25` 重命名为 `paradedb`。现有的 `USING bm25` 定义仍受支持，但新示例应使用 `USING paradedb`。
- 版本 0.25.1 支持确定性的向量并列结果排序，并将倒数排名融合查询的向量分支下推到索引中。它还新增 `paradedb.vector_clustering_threshold`，默认值为 500，并将向量索引构建并行度上限设为四个工作进程。
- 版本 0.25.1 移除了 `paradedb.vector_cluster_probe_epsilon`，并更改了向量索引的边界门控。从 0.25.0 升级数据库后，必须对所有包含向量字段的 ParadeDB 索引执行 `REINDEX`；对于这些索引，仅安装新的共享库并执行 `ALTER EXTENSION` 并不充分。
- 0.25.2 是稳定性与正确性版本：它修复带向量列的无字段 `more_like_this`、通用预备计划中的 `pdb.fuzzy`、遗留动态过滤器、多种并行子计划和 MPP 计划形态错误，并收紧 typemod 定义的访问控制。除了继承自 0.25.0 的向量索引重建要求外，没有新增索引迁移。
- `CREATE EXTENSION pg_search CASCADE` 可以安装所需的 `vector` 扩展，但仍须先为所有服务器进程配置预加载并重启。仅通过 `LOAD` 或 `session_preload_libraries` 加载并不充分。
- 使用不同字段选项重建索引后，查询计划、分词和排名都可能变化。在上线前，请使用符合生产形态的数据测试相关性与向量召回率。
