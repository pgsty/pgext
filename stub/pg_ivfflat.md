## Usage

Sources:

- [Official repository README](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/README.md)
- [PostgreSQL extension implementation](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/lib.rs)
- [Vector type implementation](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/vector_type.rs)
- [Distance operator implementation](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/operator.rs)

`pg_ivfflat` 0.1.0 is an experimental pgrx project intended to explore an IVFFlat index access method. The pinned PostgreSQL extension does not actually register an index access method: it provides a custom `vector` type and a cosine-distance operator only. Do not confuse this API with `pgvector` or assume the package name means an IVFFlat index is available.

### Available Workflow

```sql
CREATE EXTENSION pg_ivfflat;

CREATE TABLE items (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    embedding vector(3)
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id, embedding <=> '[3,1,2]'::vector(3) AS cosine_distance
FROM items
ORDER BY cosine_distance;
```

The `vector(n)` type stores JSON-array input as `float8` values and enforces dimensions in the range 1–65,535 when a type modifier is present. Its text output is a compact JSON array. `<=>` accepts two vectors and returns cosine distance as `double precision`.

### Experimental Boundaries

There is no `CREATE INDEX ... USING ivfflat` workflow, operator class, `lists` option, probes setting, or approximate-nearest-neighbor path in this source snapshot. The query above is an ordinary sort after evaluating distance for rows; confirm plans with `EXPLAIN` and expect full-table work.

The Cargo manifest enables only PostgreSQL 17 and uses pgrx 0.12.9. The generic names `vector` and `<=>` overlap APIs used by other vector extensions, so installing them in the same schema can conflict. The standalone `ivfflat` Rust crate in the repository is an in-memory experiment and is not exposed as a PostgreSQL index here. Treat this revision as development code, not a production vector-search implementation.
