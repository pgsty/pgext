## Usage

Sources:

- [Official upstream README](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/README.md)
- [Official extension control file (pgvector.control)](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/pgvector.control)
- [Official implementation source](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/src/lib.rs)

`pgvector` — Custom vector type with Euclidean distance, cosine similarity, and prototype IVFFlat helpers. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgvector;

CREATE TABLE test_vectors (
    id SERIAL PRIMARY KEY,
    embedding PGVector
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `create_vector` is an extension function.
- `euclidean_distance` is an extension function.
- `ivfflat_incex_create` is an extension function.
- `ivfflat_index_search` is an extension function.
- `vector_cosine_similarity` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
