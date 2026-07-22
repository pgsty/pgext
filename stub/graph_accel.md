## Usage

Sources:

- [Official README for version 0.5.0](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/README.md)
- [Official extension control file](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/graph_accel.control)
- [Official load implementation](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/src/load.rs)

`graph_accel` is a read-only, in-memory traversal accelerator for graphs stored by the Apache AGE extension. Use it when an AGE graph must serve repeated neighborhood, shortest-path, degree, or induced-subgraph queries without repeatedly executing Cypher. The cache belongs to one PostgreSQL backend; it is not shared across sessions.

### Core Workflow

Install Apache AGE first, create `graph_accel` as a superuser, select the AGE graph, and load its adjacency data into the current backend:

```sql
CREATE EXTENSION age;
CREATE EXTENSION graph_accel;

SET graph_accel.source_graph = 'my_graph';
SET graph_accel.node_id_property = 'concept_id';

SELECT * FROM graph_accel_load('my_graph');
SELECT * FROM graph_accel_neighborhood('concept-42', 3, 'both');
SELECT * FROM graph_accel_path('concept-42', 'concept-99', 10, 'out');
SELECT * FROM graph_accel_degree(20);
SELECT * FROM graph_accel_status();
```

`graph_accel_load()` reports node count, edge count, and load time. `graph_accel_neighborhood()` returns each reachable node with its label, application identifier, distance, and path edge metadata. `graph_accel_path()` finds an unweighted shortest path subject to its hop, direction, and optional confidence filters. `graph_accel_subgraph()` returns an induced subgraph around seed identifiers.

### Configuration and Cache Invalidation

- `graph_accel.source_graph` selects the AGE graph; it can also be passed to `graph_accel_load()`.
- `graph_accel.node_id_property` selects the node property exposed as the application identifier.
- `graph_accel.node_labels` and `graph_accel.edge_types` restrict loaded labels and edge types; `*` means all.
- `graph_accel.max_memory_mb` limits the cache in each backend, not cluster-wide memory.
- `graph_accel.auto_reload` and `graph_accel.reload_debounce_sec` control generation-based reload behavior.

AGE writes bypass ordinary PostgreSQL triggers used by this extension. After changing graph data, the writer must explicitly advance the generation:

```sql
SELECT graph_accel_invalidate('my_graph');
```

The function sends a notification, but other sessions reload only when they next check and observe a generation mismatch. Without cooperative invalidation they can continue serving stale cached topology.

### Operational Notes

The reviewed README covers PostgreSQL 13-18, but its prebuilt artifact is tied to the exact PostgreSQL/Apache AGE ABI of the documented image; rebuild for other combinations. Memory use multiplies with the number of backends that load a graph. The accelerator does not replace AGE as the source of truth and does not provide weighted shortest paths or write operations.
