## Usage

Sources:

- [Upstream source 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [Upstream source 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion_vector.control)

`pg_lion_vector` 0.3.0 provides filtered nearest-neighbor search for pgvector embeddings through the lionvec index access method.

### Core Workflow

The companion requires `pg_lion` and pgvector 0.7 or later. It supports `vector` and `halfvec`; use the corresponding distance operator class. A separate `lion` filter index can narrow candidates before nearest-neighbor ranking.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_vector;
CREATE TABLE lion_items (id integer, category integer, embedding vector(3));
CREATE INDEX ON lion_items USING lionvec (embedding vector_cosine_ops);
CREATE INDEX ON lion_items USING lion (category);
SELECT id FROM lion_items
WHERE category = 3
ORDER BY embedding <=> '[1,0,0]'::vector
LIMIT 10;
```

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Keep parent and companion files at the same version. Review upstream index options and recall behavior for the workload; an index choice is not a blanket exact-search or performance guarantee. Parent custom-WAL indexes retain their preload requirement.
