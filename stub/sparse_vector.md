## Usage

Sources:

- [Pinned official README](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/README.md)
- [Pinned extension SQL](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/sparse_vector--0.1.sql)

The `sparse_vector` extension, packaged by the upstream project as `pg_sparse_vector`, stores only the nonzero positions of a `float4[]` vector and provides dot-product and cosine-similarity calculations. It is a compact calculation type; it does not include nearest-neighbor indexes or vector-distance operators.

### Core Workflow

Values must be constructed from arrays because textual input is not implemented:

```sql
CREATE EXTENSION sparse_vector;

CREATE TABLE sparse_data (
  id bigint GENERATED ALWAYS AS IDENTITY,
  embedding sparse_vector
);

INSERT INTO sparse_data (embedding)
VALUES (sparse_vector(ARRAY[1.1, 0, 5, 0, 3]::float4[]));

SELECT id,
       cosine_similarity(
         sparse_vector(ARRAY[1, 0, 4, 0, 0]::float4[]),
         embedding
       ) AS similarity
FROM sparse_data;
```

When every stored vector and query vector is normalized, use the dot product to avoid recomputing norms:

```sql
SELECT id,
       dot_product(
         sparse_vector_norm(ARRAY[1, 0, 4, 0, 0]::float4[]),
         embedding
       ) AS similarity
FROM sparse_data;
```

### Important Objects

- `sparse_vector(float4[])`: builds a sparse value by retaining nonzero elements and their zero-based indexes.
- `sparse_vector_norm(float4[])`: builds a value after dividing elements by the array's Euclidean norm.
- `dot_product(sparse_vector, sparse_vector)`: returns a `float4` dot product over matching stored indexes.
- `cosine_similarity(sparse_vector, sparse_vector)`: returns dot product divided by both computed norms.

### Caveats

The representation does not store the original dense dimension; trailing zeros are lost, and vectors created from different array lengths are not rejected. Ensure dimensions in application metadata. Use one-dimensional, non-null `float4[]` inputs; the C implementation reads the array payload directly and does not document multidimensional or null-element handling.

An all-zero vector has zero norm, so cosine similarity can produce a non-finite result. Normalizing such a vector has no meaningful direction. Text casts such as `'(0, 1)'::sparse_vector` fail by design; use the constructors.

There are no comparison operators, aggregates, casts back to arrays, index operator classes, or approximate-nearest-neighbor search. Similarity queries therefore scan and calculate unless the application adds another retrieval strategy. No preload or restart is required.
