## Usage

Sources:

- [README.md](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/README.md)
- [hnsw--0.1.0.sql](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/hnsw--0.1.0.sql)
- [Makefile](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/Makefile)
- [hnsw.control](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/hnsw.control)

`hnsw` is an older PostgreSQL ANN prototype using real arrays and an in-memory index built on demand. It is separate from pgvector’s similarly named index access method.

### Core Workflow

```sql
CREATE EXTENSION hnsw;
CREATE TABLE embeddings (id integer PRIMARY KEY, payload real[]);
CREATE INDEX ON embeddings USING hnsw(payload) WITH (maxelements=100000, dims=3, m=32);
SELECT id FROM embeddings ORDER BY payload <-> ARRAY[1,2,3]::real[] LIMIT 10;
```

### Operational Boundaries

Create the extension as a superuser. Its SQL defines `l2_distance(real[], real[])`, the `<->` operator, the hnsw access method and its real-array operator class. Set `maxelements` to bound the index and `dims` to match the vector dimension; `ef` controls considered neighbors. Query using distance ordering with a limit.

The upstream source last changed in 2023 and declares no current PostgreSQL-major matrix. It has no preload declaration and is relocatable. Its in-memory design needs explicit capacity planning and restart/recovery validation; do not assume pgvector’s durability or operational contract. Avoid installing another extension with a conflicting access-method name in the same database.
