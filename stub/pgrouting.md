

## Usage

> [pgRouting - Routing on PostgreSQL](https://github.com/pgRouting/pgrouting)

pgRouting extends the PostGIS/PostgreSQL geospatial database to provide geospatial routing and other network analysis functionality.

This library contains the following features:

- All Pairs Shortest Path (Floyd-Warshall, Johnson)
- A* algorithm (with bidirectional variant)
- Dijkstra algorithms (cost, cost matrix, driving distance, K shortest paths, via routing, nearest)
- Bidirectional Dijkstra
- Traveling Salesman Problem (TSP)
- Network flow (max flow, Boykov-Kolmogorov, Edmonds-Karp, push-relabel)
- Spanning trees (Kruskal, Prim with BFS/DFS/driving distance variants)
- Graph components (connected, strong, biconnected, articulation points, bridges)
- Turn Restriction Shortest Path (TRSP)
- WithPoints routing (arbitrary locations on edges)
- Graph contraction and utility functions

### Getting Started

Enable the extension (requires PostGIS):

```sql
CREATE EXTENSION pgrouting CASCADE;
```

### Graph Representation

pgRouting represents graphs using SQL queries that return edge data. The standard edge query format:

```sql
SELECT id, source, target, cost, reverse_cost FROM edges;
```

| Column | Type | Description |
|--------|------|-------------|
| `id` | ANY-INTEGER | Edge identifier |
| `source` | ANY-INTEGER | Starting vertex identifier |
| `target` | ANY-INTEGER | Ending vertex identifier |
| `cost` | ANY-NUMERICAL | Weight (source to target); negative values exclude the edge |
| `reverse_cost` | ANY-NUMERICAL | Weight (target to source); default -1 (non-existent) |

### Simple Example Without Geometry

Create a graph and find the shortest path:

```sql
CREATE TABLE wiki (
  id SERIAL,
  source INTEGER,
  target INTEGER,
  cost INTEGER
);

INSERT INTO wiki (source, target, cost) VALUES
  (1, 2, 7),  (1, 3, 9), (1, 6, 14),
  (2, 3, 10), (2, 4, 15),
  (3, 6, 2),  (3, 4, 11),
  (4, 5, 6),
  (5, 6, 9);

SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost FROM wiki',
  1, 5, false);
```

--------

## Function Families

### Dijkstra - Shortest Path

The core routing function. Supports one-to-one, one-to-many, many-to-one, many-to-many, and combinations signatures.

```sql
pgr_dijkstra(Edges SQL, start_vid,  end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vid,  end_vids, [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vids, [directed])
pgr_dijkstra(Edges SQL, Combinations SQL,     [directed])
```

Returns: `(seq, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

**One to One:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10, true);
```

**One to Many:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, ARRAY[10, 17]);
```

**Many to Many (undirected):**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  ARRAY[6, 1], ARRAY[10, 17],
  directed => false);
```

**Combinations:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  'SELECT source, target FROM combinations',
  false);
```

#### Dijkstra Cost

Returns only aggregate cost without path details:

```sql
pgr_dijkstraCost(Edges SQL, start_vid, end_vid, [directed])
```

Returns: `(start_vid, end_vid, agg_cost)`

#### Dijkstra Cost Matrix

Generate a cost matrix for a set of vertices:

```sql
pgr_dijkstraCostMatrix(Edges SQL, vids, [directed])
```

#### Dijkstra Via

Route through an ordered sequence of vertices:

```sql
pgr_dijkstraVia(Edges SQL, via_vertices, [directed, strict, U_turn_on_edge])
```

#### Dijkstra Near

Find the nearest vertex to a set of targets:

```sql
pgr_dijkstraNear(Edges SQL, start_vid, end_vids, [directed])
```

### A* - Shortest Path

Uses the A* heuristic algorithm. Requires additional coordinate columns (`x1`, `y1`, `x2`, `y2`) in the edges query.

```sql
pgr_aStar(Edges SQL, start_vid, end_vid, [directed, heuristic, factor, epsilon])
```

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `directed` | BOOLEAN | `true` | Graph direction |
| `heuristic` | INTEGER | `5` | Distance heuristic (0-5) |
| `factor` | FLOAT | `1` | Units manipulation |
| `epsilon` | FLOAT | `1` | Approximation factor |

```sql
SELECT * FROM pgr_aStar(
  'SELECT id, source, target, cost, reverse_cost, x1, y1, x2, y2 FROM edges',
  6, 12,
  directed => true, heuristic => 2);
```

Also available: `pgr_aStarCost`, `pgr_aStarCostMatrix`

### Bidirectional Algorithms

Bidirectional variants search from both ends simultaneously:

- `pgr_bdDijkstra`, `pgr_bdDijkstraCost`, `pgr_bdDijkstraCostMatrix`
- `pgr_bdAstar`, `pgr_bdAstarCost`, `pgr_bdAstarCostMatrix`

```sql
SELECT * FROM pgr_bdDijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10);
```

### K Shortest Paths (Yen's Algorithm)

Find the K shortest paths between two vertices:

```sql
pgr_KSP(Edges SQL, start_vid, end_vid, K, [directed, heap_paths])
```

Returns: `(seq, path_id, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_KSP(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 17, 2);
```

### Driving Distance

Find all vertices reachable within a given distance:

```sql
pgr_drivingDistance(Edges SQL, start_vid, distance, [directed])
pgr_drivingDistance(Edges SQL, start_vids, distance, [directed, equicost])
```

Returns: `(seq, depth, start_vid, pred, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_drivingDistance(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  11, 3.0);
```

### Traveling Salesman Problem

**Matrix-based TSP:**

```sql
pgr_TSP(Matrix SQL, [start_id, end_id])
```

Returns: `(seq, node, cost, agg_cost)`

```sql
SELECT * FROM pgr_TSP(
  $$SELECT * FROM pgr_dijkstraCostMatrix(
    'SELECT id, source, target, cost, reverse_cost FROM edges',
    ARRAY[1, 3, 5, 6, 7, 8, 9, 10, 11, 15, 16, 17],
    directed => false)$$,
  start_id => 1);
```

**Euclidean TSP** (uses coordinates directly):

```sql
pgr_TSPeuclidean(Coordinates SQL, [start_id, end_id])
```

### Network Flow

Compute maximum flow and related properties:

```sql
-- Maximum flow
pgr_maxFlow(Edges SQL, source, sink)

-- Specific algorithms
pgr_boykovKolmogorov(Edges SQL, source, sink)
pgr_edmondsKarp(Edges SQL, source, sink)
pgr_pushRelabel(Edges SQL, source, sink)

-- Edge-disjoint paths
pgr_edgeDisjointPaths(Edges SQL, source, sink)

-- Maximum cardinality matching
pgr_maxCardinalityMatch(Edges SQL, [directed])
```

Flow edges SQL uses `capacity` and `reverse_capacity` instead of `cost`/`reverse_cost`.

### Spanning Trees

**Kruskal's algorithm:**

```sql
pgr_kruskal(Edges SQL)         -- minimum spanning tree
pgr_kruskalBFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDD(Edges SQL, root_vid, distance)
```

**Prim's algorithm:**

```sql
pgr_prim(Edges SQL)            -- minimum spanning tree
pgr_primBFS(Edges SQL, root_vid, [max_depth])
pgr_primDFS(Edges SQL, root_vid, [max_depth])
pgr_primDD(Edges SQL, root_vid, distance)
```

### Graph Components

```sql
-- Connected components (undirected)
pgr_connectedComponents(Edges SQL)

-- Strongly connected components (directed)
pgr_strongComponents(Edges SQL)

-- Biconnected components
pgr_biconnectedComponents(Edges SQL)

-- Articulation points (cut vertices)
pgr_articulationPoints(Edges SQL)

-- Bridges (cut edges)
pgr_bridges(Edges SQL)
```

### Turn Restriction Shortest Path (TRSP)

Route with forbidden path restrictions:

```sql
pgr_trsp(Edges SQL, Restrictions SQL, start_vid, end_vid, [directed])
pgr_trspVia(Edges SQL, Restrictions SQL, via_vertices, [directed, strict, U_turn_on_edge])
pgr_trsp_withPoints(Edges SQL, Restrictions SQL, Points SQL, start_pid, end_pid, [options])
```

Restrictions SQL format:

| Column | Type | Description |
|--------|------|-------------|
| `path` | ARRAY[ANY-INTEGER] | Sequence of forbidden edge IDs |
| `cost` | ANY-NUMERICAL | Cost of the forbidden path |

### WithPoints - Routing from Arbitrary Locations

Route between points located on edges (not just vertices):

```sql
pgr_withPoints(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCost(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCostMatrix(Edges SQL, Points SQL, pids, [options])
pgr_withPointsKSP(Edges SQL, Points SQL, start_pid, end_pid, K, [options])
pgr_withPointsDD(Edges SQL, Points SQL, start_pid, distance, [options])
```

Points SQL format:

| Column | Type | Default | Description |
|--------|------|---------|-------------|
| `pid` | ANY-INTEGER | | Point identifier |
| `edge_id` | ANY-INTEGER | | Closest edge |
| `fraction` | ANY-NUMERICAL | | Position on edge (0-1) |
| `side` | CHAR | `b` | `r`(right), `l`(left), `b`(both) |

### Graph Contraction

Simplify graphs by contracting vertices:

```sql
pgr_contraction(Edges SQL, contraction_order, [options])
```

### Utility Functions

```sql
-- Extract vertices from edge data
pgr_extractVertices(Edges SQL)

-- Find edges near a point
pgr_findCloseEdges(Edges SQL, point, tolerance, [options])

-- Separate crossing geometries
pgr_separateCrossing(Edges SQL)

-- Separate touching geometries
pgr_separateTouching(Edges SQL)

-- Version info
pgr_version()
pgr_full_version()
```

--------

## Working with Geometries

### Building a Routing Topology

Extract vertices from spatial edges and build topology:

```sql
-- Extract vertices from edge geometries
SELECT * INTO vertices
FROM pgr_extractVertices('SELECT id, geom FROM edges ORDER BY id');

-- Set source vertex
UPDATE edges AS e
SET source = v.id, x1 = x, y1 = y
FROM vertices AS v
WHERE ST_StartPoint(e.geom) = v.geom;

-- Set target vertex
UPDATE edges AS e
SET target = v.id, x2 = x, y2 = y
FROM vertices AS v
WHERE ST_EndPoint(e.geom) = v.geom;
```

### Setting Costs from Geometry Length

```sql
UPDATE edges SET
  cost = sign(cost) * ST_Length(geom),
  reverse_cost = sign(reverse_cost) * ST_Length(geom);
```

### Getting Route Geometry

Combine routing results with edge geometries:

```sql
SELECT seq, node, edge, cost, agg_cost, geom
FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10
) AS r
LEFT JOIN edges AS e ON r.edge = e.id;
```

--------

## Performance Tips

Bound queries to the area of interest to reduce processed edges:

```sql
SELECT * FROM pgr_dijkstra($$
  SELECT id, source, target, cost, reverse_cost
  FROM edges
  WHERE geom && (
    SELECT ST_Buffer(ST_Union(geom), 1)
    FROM edges WHERE source IN (7) OR target IN (8))$$,
  7, 8);
```

--------

## All Pairs Shortest Path

For computing distances between all pairs of vertices:

```sql
-- Floyd-Warshall (no edge id required)
pgr_floydWarshall(Edges SQL, [directed])

-- Johnson (no edge id required)
pgr_johnson(Edges SQL, [directed])
```

Returns: `(start_vid, end_vid, agg_cost)`

```sql
SELECT * FROM pgr_floydWarshall(
  'SELECT id, source, target, cost, reverse_cost FROM edges');
```
