## Usage

Sources:

- [Official extension control file](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/pg_knowledge_graph.control)
- [Official README](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/README.md)
- [Official schema SQL](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/sql/pg_knowledge_graph--0.1.0.sql)

`pg_knowledge_graph` 0.1.0 stores entities and directed, weighted relationships in ordinary PostgreSQL tables and exposes traversal, ranking, component, vector, hybrid-search, RAG-context, and approximate quantized-search functions. It depends on `vector` and uses fixed 1536-dimensional embeddings in the upstream schema.

### Core Workflow

Install pgvector first. The pgrx installation supplies the extension functions, but the repository's own Docker initializer states that the graph tables must be created separately. Apply the cited schema SQL or create the two tables explicitly before inserting data.

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

Create the indexes from the official schema SQL, including the HNSW index on `embedding`, before using vector search at scale.

### Function Groups

- Inspection: `kg_version` and `kg_stats`.
- Traversal: `kg_bfs`, `kg_dfs`, and `kg_shortest_path` over directed edges.
- Whole-graph algorithms: `kg_pagerank`, `kg_louvain`, `kg_connected_components`, and `kg_strongly_connected_components`.
- Retrieval: `kg_vector_search`, `kg_hybrid_search`, and `kg_get_context`.
- Approximate retrieval: `kg_quantized_search` with `int8`, `int4`, or `binary`, plus `kg_quantize_info`.

`kg_hybrid_search` combines vector and graph scores using caller-provided `alpha` and `beta`. `kg_get_context` returns an N-hop neighborhood for retrieval-augmented generation; it does not call an embedding or language-model service.

### Operational Boundaries

The Rust functions query `kg_entities` and `kg_relations` through SPI and load graph or embedding data into the calling backend for several algorithms. PageRank, component detection, community detection, and quantized search are recomputed on demand rather than maintained incrementally. Benchmark memory, execution time, and statement timeouts on realistic graph sizes.

The functions use unqualified table names, so keep the official tables in the expected search path and prevent untrusted users from substituting look-alike relations. The extension is superuser-only, untrusted, and non-relocatable according to its control file.

Quantized similarity is approximate. The README's compression and recall figures are project measurements, not correctness guarantees; evaluate recall using your embeddings. Version 0.1.0 declares PostgreSQL 16, 17, and 18 build features and should be treated as an early release until independently load- and failure-tested.
