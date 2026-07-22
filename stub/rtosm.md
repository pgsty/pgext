## Usage

Sources:

- [Official extension control file](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm.control)
- [Official README](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/README.md)
- [Official extension SQL](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm--0.1.sql)

`rtosm` 0.1 implements error-bounded, real-time simplification for spatial objects stored in the OpenStreetMap Rails API database used by `openstreetmap-website`. It builds EB-tree metadata for ways and relations, then returns a bounded set of node IDs for a bounding box. It is not a generic PostGIS geometry simplifier.

### Prerequisites and Build

Install into a database that already has the exact OpenStreetMap API schema. The extension SQL directly references objects such as `nwr_enum`, `current_nodes`, `way_nodes`, `relation_members`, and `relation_tags`; `CREATE EXTENSION rtosm` will fail in an ordinary PostgreSQL database. Install the declared `intarray` dependency first.

```sql
CREATE EXTENSION intarray;
CREATE EXTENSION rtosm;

SELECT build_way_trees();
SELECT build_relation_trees();
SELECT build_indexes();
SELECT tileid_c(30000);

SELECT unnest(
    vquery_c(-0.15, 51.48, -0.05, 51.55, 5000, 0.2)
) AS node_id;
```

The four build calls can scan and materialize substantial portions of the OSM dataset. Run them in a maintenance window, monitor space and locks, and validate results on a copy before using them against a live API database.

### Main Objects

- Persistent metadata tables: `tile_nums`, `node_vis_errors`, `way_trees`, and `relation_trees`.
- Builders: `build_way_trees()`, `build_relation_trees()`, `build_indexes()`, and tile helpers including `tileid_c(integer)`.
- Query: `vquery_c(x1, y1, x2, y2, k, e) RETURNS bigint[]`; the first four arguments define the bounding box in degrees, `k` limits returned nodes, and `e` is the error limit.
- Lower-level C and PL/pgSQL helpers implement EB-tree construction, spatial filtering, and sequence editing.

The main result contains node IDs needed to assemble simplified OSM objects. The caller must join them back to the OSM schema and construct the desired response or visualization.

### Schema and Maintenance Boundaries

Version 0.1 is historical research code from 2016 and assumes the contemporary Rails schema and its coordinate conventions. Verify every referenced type, table, column, and member encoding against the deployed `openstreetmap-website` revision before installation.

The source exposes editing helpers but does not document an automatic trigger integration with the Rails write path. Do not assume `way_trees`, `relation_trees`, or `node_vis_errors` stay synchronized after OSM edits; define and test the update/rebuild procedure for the application. Back up these derived tables or make rebuild time part of recovery planning.
