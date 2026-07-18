## Usage

Sources:

- [Pinned upstream README](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/README.md)
- [Version 0.1.0 SQL API](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/sql/pgr_dmmsy--0.1.0.sql)
- [Pinned PostgreSQL C wrapper](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/src/dmmsy.c)
- [Pinned algorithm implementation](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/src/dmmsy_algorithm.c)
- [Pinned parameter tests](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/test/sql/test_parameters.sql)

pgr_dmmsy version 0.1.0 implements a directed or undirected non-negative single-source shortest-path function with a pgRouting-style edge query. It requires PostgreSQL 14 or newer and pgRouting 3.x. The function currently works as a point-to-point path query; the documented all-destinations form is not implemented correctly.

### Point-to-point path

The edge SQL must return id BIGINT, source BIGINT, target BIGINT, and cost FLOAT8 in exactly that order:

```sql
CREATE EXTENSION pgrouting;
CREATE EXTENSION pgr_dmmsy;

CREATE TABLE edges (
  id bigint,
  source bigint,
  target bigint,
  cost double precision
);

INSERT INTO edges VALUES
  (1, 1, 2, 1.5),
  (2, 2, 3, 2.0),
  (3, 3, 4, 0.5);

SELECT seq, path_seq, node, edge, cost, agg_cost
FROM pgr_dmmsy(
  'SELECT id, source, target, cost FROM edges',
  1,
  4
)
ORDER BY path_seq;
```

NULL edge fields and negative costs raise errors. directed false adds a reverse edge for each input row. An unknown or unreachable target returns no rows.

### Current implementation boundaries

The README says a NULL target returns a distance map for all reachable vertices, but the C wrapper constructs output only when target is non-NULL and non-negative. The pinned regression test explicitly expects NULL target to return an empty set. Always provide a concrete target until the wrapper has a tested all-destinations output path.

output_predecessors and constant_degree are accepted but unused by the current source. max_levels can stop phase processing before the target is settled and then return no path. The repository's implementation notes and current algorithm comments also disagree about whether this is a simplified Dijkstra-based implementation or the advertised asymptotic DMMSY algorithm. Treat the complexity claim as unverified and benchmark correctness and performance against pgr_dijkstra on your graphs.

The wrapper reads the complete edge query into backend memory before computing. It checks only cost less than zero, so callers should reject NaN and infinity explicitly. Arbitrary edges_sql executes with caller privileges; allow only trusted SQL templates, qualify every object, and do not concatenate user input.

The project is active but version 0.1.0, has no declared license, and CI covers only a subset of supported majors. Validate equal-cost ties, parallel edges, self-loops, zero weights, huge or sparse vertex IDs, disconnected graphs, cancellation, memory pressure, directed false, every tuning parameter, and comparison with a trusted implementation before production use.
