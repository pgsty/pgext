## Usage

Sources:

- [Official extension control file](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/ast_postgis.control)
- [Official upstream documentation](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/README.md)
- [Official consistency-check implementation](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/sql/consistency_functions.sql)

`ast_postgis` 1.1.1 adds OMT-G-inspired spatial domains and database-side integrity checks to PostGIS. Use it when a spatial model needs simple-geometry, network, topology, or aggregation rules enforced near the data. It requires `postgis` and installs event triggers, so installation and DDL integration require an appropriately privileged administrator.

### Core Workflow

Create the extension, define columns with its domains, then attach a relationship trigger:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ast_postgis;

CREATE TABLE bus_stop (
  stop_id integer PRIMARY KEY,
  geom ast_point
);

CREATE TABLE route_segment (
  segment_id integer PRIMARY KEY,
  geom ast_uniline
);

CREATE TRIGGER route_segment_nodes
AFTER INSERT OR UPDATE ON route_segment
FOR EACH STATEMENT
EXECUTE PROCEDURE ast_arcnodenetwork(
  'route_segment', 'geom', 'bus_stop', 'geom'
);
```

Relationship triggers validate the affected model after writes and raise an error when the configured rule is violated.

### Important Objects

- Geometry domains include `ast_polygon`, `ast_line`, `ast_point`, `ast_node`, `ast_isoline`, `ast_planarsubdivision`, `ast_tin`, `ast_sample`, `ast_uniline`, and `ast_biline`; `ast_tesselation` wraps PostGIS raster.
- `ast_spatialrelationship()` supports `contains`, `containsproperly`, `covers`, `coveredby`, `crosses`, `disjoint`, `distant`, `intersects`, `near`, `overlaps`, `touches`, and `within` trigger rules.
- `ast_arcnodenetwork()`, `ast_arcarcnetwork()`, and `ast_aggregation()` enforce network and part/whole relationships.
- `ast_isSpatialRelationshipValid(...)`, `ast_isNetworkValid(...)`, and `ast_isSpatialAggregationValid(...)` inspect existing data and write violations to `ast_violation_log`.

### Caveats

The extension implements constraints with statement-level triggers, event triggers, PL/pgSQL, and dynamically constructed SQL. Restrict its checker and trigger-management APIs to trusted roles, pass only administrator-controlled identifiers, and test write concurrency and large-table cost. The reviewed `ast_isSpatialAggregationValid(...)` source contains a hard-coded query against `tablea` and `tableb`, so do not trust that checker without patching and regression-testing it for the actual model. Verify every domain and relation on a copy before enabling enforcement on existing data, and review `ast_violation_log` retention because validation records are included in extension configuration dumps.
