## Usage

Sources:

- [pg_search v0.25.1 README](https://github.com/paradedb/paradedb/blob/v0.25.1/pg_search/README.md)
- [pg_search v0.25.1 release](https://github.com/paradedb/paradedb/releases/tag/v0.25.1)
- [pg_search v0.25.1 changelog](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/changelog/0.25.1.mdx)
- [Create a ParadeDB index](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/indexing/create-index.mdx)
- [Full-text match operators](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/full-text/match.mdx)
- [BM25 scoring](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/sorting/score.mdx)
- [Highlighting and snippets](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/full-text/highlight.mdx)
- [Index vectors](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/indexing/indexing-vectors.mdx)
- [Query vectors](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/vector/querying.mdx)
- [Hybrid-search overview](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/hybrid/overview.mdx)

`pg_search` adds ParadeDB's full-text, structured, vector, and hybrid search index to PostgreSQL. Version 0.25 uses the `paradedb` index access method; the older `bm25` access-method name remains a compatibility alias. The extension requires `vector`, supports PostgreSQL 15-18 upstream, and must be loaded through `shared_preload_libraries`.

### Install and Build an Index

```conf
shared_preload_libraries = 'pg_search'
```

Restart PostgreSQL, then create the extension and a table with a stable unique key:

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

The `key_field` must be the first indexed column and uniquely identify every row. A text key must be indexed without tokenization. A table can have only one ParadeDB index, so include every searchable field in that index.

### Full-Text Search

Use `|||` to match any token and `&&&` to require all tokens:

```sql
SELECT id, title, pdb.score(id) AS score
FROM documents
WHERE body ||| 'postgresql search'
ORDER BY score DESC, id;

SELECT id, pdb.snippet(body) AS excerpt
FROM documents
WHERE body &&& 'postgresql indexing';
```

`pdb.score(key_field)` exposes the relevance score for the current row. `pdb.snippet(indexed_text_column)` returns a highlighted excerpt. These helpers are meaningful only in a query driven by a ParadeDB search predicate.

### Vector Search

Vector indexing is beta in the 0.25 line and uses the `vector` type from pgvector. Choose the operator class when the index is created; changing the metric requires rebuilding the index.

```sql
SELECT id, title, embedding <=> $1::vector AS distance
FROM documents
WHERE id @@@ pdb.all()
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

Supported index operator classes are `vector_l2_ops`, `vector_ip_ops`, and `vector_cosine_ops`. The 0.25 vector index does not index `halfvec`, `sparsevec`, or `bit` columns.

### Hybrid Search

A single ParadeDB index can combine lexical predicates, structured filters, and vector ordering. For more elaborate fusion, use the documented RRF and weighted hybrid-search functions instead of adding scores from unrelated scales directly.

```sql
SELECT id, title, pdb.score(id) AS lexical_score
FROM documents
WHERE body ||| 'postgresql extension'
  AND category === 'database'
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

### Version 0.25.1 and Caveats

- Version 0.25 renamed the primary index access method from `bm25` to `paradedb`. Existing `USING bm25` definitions remain supported, but new examples should use `USING paradedb`.
- Version 0.25.1 supports deterministic vector tie breakers and pushes the vector arm of reciprocal-rank-fusion queries into the index. It also adds `paradedb.vector_clustering_threshold`, whose default is 500, and caps vector-index build parallelism at four workers.
- Version 0.25.1 removes `paradedb.vector_cluster_probe_epsilon` and changes the vector-index bounds gate. After upgrading a database from 0.25.0, `REINDEX` every ParadeDB index that contains a vector field; installing the new shared library and running `ALTER EXTENSION` alone is not sufficient for those indexes.
- `CREATE EXTENSION pg_search CASCADE` can install the required `vector` extension, but every server process still needs the preload configuration and restart first.
- Query plans, tokenization, and ranking can change when an index is rebuilt with different field options. Test relevance and vector recall with production-shaped data before rollout.
