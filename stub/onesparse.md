## Usage
> Sources: [homepage](https://onesparse.github.io/OneSparse/), [intro docs](https://onesparse.github.io/OneSparse/docs.html), [Matrix](https://onesparse.github.io/OneSparse/test_matrix_header.html), [Vector](https://onesparse.github.io/OneSparse/test_vector_header.html), and [Algorithms](https://onesparse.github.io/OneSparse/test_examples_header.html).

OneSparse is a PostgreSQL extension that binds SuiteSparse:GraphBLAS into Postgres and exposes sparse linear algebra and graph algorithms as new types, functions, and operators.
The docs treat `matrix` as the core type, with `vector` and `scalar` built on top of the same model.

### Core Setup

```sql
CREATE EXTENSION onesparse;
SET search_path TO public,onesparse;

SELECT 'int32'::matrix;
SELECT 'int32'::vector;
SELECT 'int32:42'::scalar;
```

The intro docs note that OneSparse keeps its API in the `onesparse` schema, and the matrix/vector pages show the same `search_path` pattern for interactive use.

### Matrix and Vector

The matrix page shows common operations such as constructing, printing, drawing, resizing, casting, and aggregating matrices.
The vector page shows the matching vector API, including `nvals`, `size`, `set_element`, `get_element`, `eadd`, `emult`, `reduce_scalar`, `choose`, and `apply`.

```sql
SELECT print('int32(4:4)'::matrix);
SELECT draw('int32(4:4)[1:2:1 2:3:2 3:1:3]'::matrix);
SELECT eadd('int32[0:1 1:2 2:3]'::vector, 'int32[0:1 1:2 2:3]'::vector, 'plus_int32');
SELECT reduce_scalar('int32[0:1 1:2 2:3]'::vector, 'plus_monoid_int32');
```

### Graph Algorithms

The getting-started docs use graph examples built from Matrix Market files and random graphs.
They highlight these algorithms:

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

The same guide also shows graph loading with `mmread(...)` and graph visualization with `draw(...)`.

### Scope

The documentation set is broad. This stub captures the core interface and the main examples that repeat across the intro, matrix, vector, and algorithms pages, without reproducing the full GraphBLAS catalog.
