## Usage

Sources:

- [Official README](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/README.md)
- [Extension control file](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/pgext/pgtempus.control)
- [Versioned install SQL](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/pgext/pgtempus--2.1.sql)

`pgtempus` loads multimodal transport graphs into a PostgreSQL backend and computes one-to-many or many-to-one paths with time, fare, transfer, and mode-change costs.

### Enablement

Install the 2.1 native library and extension files, then create the relocatable extension in the target schema:

```sql
CREATE SCHEMA tempus;
CREATE EXTENSION pgtempus SCHEMA tempus;
```

The control file does not require superuser installation or preloading. Upstream does not publish a current PostgreSQL-major support matrix, so validate 2.1 against the exact server build before deployment.

### Build a Session Graph

`tempus_build_multimodal_graph` accepts SQL queries describing transport modes, nodes, and arcs. The query strings execute in the caller's database context.

```sql
SELECT tempus_build_multimodal_graph(
  'city',
  'SELECT id, name, category, traffic_rules FROM transport_modes',
  'SELECT id, parking_transport_modes, x, y, z FROM nodes',
  'SELECT id, node_from_id, node_to_id, is_pt, traffic_rules FROM arcs'
);

SELECT * FROM tempus_loaded_graphs();
DELETE FROM session_graph WHERE id = 'city';
```

Graphs live in backend memory and disappear when the session ends. Reuse a pooled session deliberately; do not assume another backend can see a loaded graph.

### Costs and Routing

Use `tempus_set_static_road_section_costs`, `tempus_set_pt_section_timetable`, `tempus_set_arcs_sequence_costs`, and related functions to attach generalized costs. `tempus_one_to_many_paths` handles departure-time searches, while `tempus_many_to_one_paths` handles arrive-before searches.

The API accepts caller-provided SQL text and can read large graph sets into backend memory. Restrict execution to trusted roles, validate the supplied queries, and size sessions for the graph and result cardinality. Version 2.1 is inactive upstream; treat modern PostgreSQL compatibility and restore behavior as deployment tests, not assumptions.

