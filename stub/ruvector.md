## Usage

Sources:

- [Official extension README](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/README.md)
- [Official extension control file](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/ruvector.control)
- [Official extension build manifest](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/Cargo.toml)

`ruvector` adds a fixed-dimension vector type, distance operators, and a RuVector HNSW access method to PostgreSQL. Use it for nearest-neighbor retrieval after confirming that the exact packaged build exposes the functions you need: the reviewed extension control version is `0.3.0`, while the current Rust crate manifest reports `2.0.1` and gates many advanced modules behind compile-time features.

### Core Workflow

Create the extension, store vectors with an explicit dimension, and order by the chosen distance operator:

```sql
CREATE EXTENSION ruvector;

CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    content text NOT NULL,
    embedding ruvector(3) NOT NULL
);

INSERT INTO documents (content, embedding) VALUES
    ('alpha', '[0.1,0.2,0.3]'),
    ('beta',  '[0.8,0.1,0.2]');

CREATE INDEX documents_embedding_idx
    ON documents USING ruhnsw (embedding ruvector_l2_ops);

SELECT id, content,
       embedding <-> '[0.2,0.2,0.3]'::ruvector AS distance
FROM documents
ORDER BY embedding <-> '[0.2,0.2,0.3]'::ruvector
LIMIT 5;
```

The query must contain the same `ORDER BY` distance expression used to rank rows. Validate recall, latency, and index behavior with representative data rather than treating an approximate-nearest-neighbor result as exact.

### Important Objects

- `ruvector(n)` stores vectors whose dimension is fixed by the type modifier.
- `<->` computes L2 distance for ordering; the upstream extension also defines function forms for L2, cosine, inner-product, and Manhattan calculations.
- `ruhnsw` is the HNSW index access method, and `ruvector_l2_ops` selects its L2 operator class.
- Advanced graph, attention, solver, local-embedding, and mathematical functions are feature-dependent. Inspect `pg_proc` on the installed build before relying on a function shown by newer upstream documentation.

### Operational Notes

The reviewed upstream README declares PostgreSQL 14 through 17 support; test the package against the exact server major version. The control file marks the extension trusted and non-relocatable, but object privileges still need normal review.

HNSW indexes trade build time, memory, disk space, and exactness for search speed. Rebuild and recall-test them after major upgrades or bulk data changes, and include both the table and index in backup/restore exercises. Local embedding or AI functions, when compiled in, load models and run inference inside database backends, so measure backend memory, CPU use, statement timeouts, and concurrency before enabling them for production traffic.
