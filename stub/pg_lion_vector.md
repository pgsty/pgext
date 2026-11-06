## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector.control)
- [Extension SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector--0.3.0.sql)

`pg_lion_vector` 0.3.0 adds the `lionvec` nearest-neighbor access method for pgvector `vector` and `halfvec` values. It uses the parent `pg_lion` library and requires pgvector 0.7 or later.

### Core Workflow

Create the parent and `vector` extensions first. A separate `lion` index can filter candidates before distance ranking.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_vector;
CREATE TABLE lion_items (id integer, category integer, embedding vector(3));
INSERT INTO lion_items VALUES (1, 3, '[1,0,0]'), (2, 3, '[0,1,0]'), (3, 4, '[0,0,1]');
CREATE INDEX lion_embedding ON lion_items USING lionvec (embedding vector_cosine_ops);
CREATE INDEX lion_category ON lion_items USING lion (category);
ALTER FUNCTION cosine_distance(vector, vector) COST 100;
VACUUM ANALYZE lion_items;
SELECT id FROM lion_items
WHERE category = 3
ORDER BY embedding <=> '[1,0,0]'::vector
LIMIT 10;
```

Upstream recommends `COST 100` for the distance function used by the chosen operator class so the planner can compare scan and sort costs. This setting is not preserved by `pg_dump` and may be reset by an update of `vector`.

### Operator Classes and Tuning

| Type | L2 (`<->`) | Negative inner product (`<#>`) | Cosine (`<=>`) |
| --- | --- | --- | --- |
| `vector` | `vector_l2_ops` (default) | `vector_ip_ops` | `vector_cosine_ops` |
| `halfvec` | `halfvec_l2_ops` (default) | `halfvec_ip_ops` | `halfvec_cosine_ops` |

`partitions` controls the index build's clustering; `maintenance_work_mem` must hold an adequate sample. `pg_lion.vector_margin` and `pg_lion.vector_patience` control approximate search. Small candidate sets under a supported lion filter can use exact scoring. `pg_lion.vector_postfilter` controls predicates the lion indexes cannot answer; review the reference before changing it.

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Only one `vector` or `halfvec` column, up to 8128 dimensions, is supported; this is generally approximate search. A filter does not turn every query into exact nearest-neighbor search. Heavy churn may require occasional `REINDEX`. Parent custom-WAL indexes retain their preload requirement. Keep parent and companion files at the same version.
