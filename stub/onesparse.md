## Usage
> Sources: [homepage](https://onesparse.com/), [release v1.0.0](https://github.com/OneSparse/OneSparse/releases/tag/v1.0.0), [control file at v1.0.0](https://raw.githubusercontent.com/OneSparse/OneSparse/v1.0.0/onesparse.control), [intro docs](https://onesparse.com/docs.html), [matrix docs](https://onesparse.com/test_matrix_header.html), [vector docs](https://onesparse.com/test_vector_header.html), [algorithm examples](https://onesparse.com/test_examples_header.html)

OneSparse is a PostgreSQL extension that binds SuiteSparse:GraphBLAS into Postgres and exposes sparse linear algebra and graph algorithms as new types, functions, and operators.
The docs treat `matrix` as the core type, with `vector` and `scalar` built on top of the same model. The `v1.0.0` release exists, while the extension control file at that tag still declares SQL `default_version = '0.1.0'`.

### Core Setup

```sql
CREATE EXTENSION onesparse;
SET search_path TO public,onesparse;

SELECT 'int32'::matrix;
SELECT 'int32'::vector;
SELECT 'int32:42'::scalar;
```

The docs site organizes the API around `matrix`, `vector`, and `scalar`, with interactive examples using casts and constructors.

### Matrix and Vector

The matrix page shows common operations such as construction, `print()`, `draw()`, assignment, extraction, `cast_to()`, resize, and aggregation. The vector page documents the matching vector API including `nvals()`, `size()`, `eadd()`, `emult()`, `reduce_scalar()`, `choose()`, and `apply()`.

```sql
SELECT print('int32(4:4)'::matrix);
SELECT draw('int32(4:4)[1:2:1 2:3:2 3:1:3]'::matrix);
SELECT eadd('int32[0:1 1:2 2:3]'::vector, 'int32[0:1 1:2 2:3]'::vector, 'plus_int32');
SELECT reduce_scalar('int32[0:1 1:2 2:3]'::vector, 'plus_monoid_int32');
```

### Graph Algorithms

The examples page uses Matrix Market input and graph visualization with `draw(...)`. The documented graph algorithms include:

- `bfs(graph, 1)` for level and parent BFS
- `sssp(cast_to(graph, 'int32'), 1::bigint, 1)` for single-source shortest path
- `pagerank(graph)` for ranking vertices by link structure
- `triangle_centrality(graph)` for triangle-based centrality
- `betweenness(graph, ARRAY[...])` and `square_clustering(graph)` for additional graph analysis

Representative example from the docs:

```sql
SELECT draw(triu(graph), (SELECT level FROM bfs(graph, 1)), false, false, true, 0.5)
FROM karate;
```

The same guide shows loading a graph with `mmread('/home/postgres/onesparse/demo/karate.mtx')`.
