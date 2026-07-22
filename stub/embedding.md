## Usage

Sources:

- [Upstream README and deprecation notice](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/README.md)
- [Extension control file](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/embedding.control)
- [Version 0.3.6 SQL](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/embedding--0.3.6.sql)

`embedding` version `0.3.6` stores vectors as `real[]` and adds an HNSW access method for approximate nearest-neighbor search. Operators include Euclidean distance `<->`, cosine distance `<=>`, and Manhattan distance `<~>`.

### Example

```sql
CREATE EXTENSION embedding;
CREATE TABLE documents (id integer PRIMARY KEY, embedding real[]);
INSERT INTO documents VALUES (1, ARRAY[1, 2, 3]), (2, ARRAY[4, 5, 6]);
SELECT id FROM documents ORDER BY embedding <-> ARRAY[3, 3, 3] LIMIT 1;
CREATE INDEX ON documents USING hnsw (embedding)
  WITH (dims=3, m=3, efconstruction=5, efsearch=5);
```

The upstream repository is archived and explicitly says `pg_embedding` is deprecated, with migration to `pgvector` strongly recommended. Keep vector dimensions consistent; `dims` is mandatory for an index and must equal the array element count.

Use dense, NULL-free arrays with finite coordinates. The C distance path reads the raw `real[]` payload without checking for null elements, and cosine distance divides by both vector norms, so a zero vector produces an undefined/NaN result. Validate HNSW recall and tune its build/query parameters against representative data. Use this only to support an existing deployment while planning migration, not for a new system.
