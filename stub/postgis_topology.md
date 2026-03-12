

## Usage

> [PostGIS Topology: Topological data model support for PostGIS](https://github.com/postgis/postgis)

PostGIS Topology implements the SQL/MM topology model for PostgreSQL. It represents spatial data as a graph of nodes, edges, and faces, ensuring that shared boundaries are stored only once and geometric consistency is enforced.

- [Topology Reference](https://postgis.net/docs/Topology.html)

### Setup

```sql
CREATE EXTENSION postgis_topology;
```

This creates the `topology` schema with management tables and functions.

--------

## Creating a Topology

A topology is a named schema with a defined SRID and tolerance for snapping:

```sql
-- Create a topology named 'city_topo' with SRID 4326 and 0.00001 degree tolerance
SELECT topology.CreateTopology('city_topo', 4326, 0.00001);
```

List all topologies:

```sql
SELECT * FROM topology.topology;
```

--------

## Building Topology Primitives

### Adding Edges

Edges are the fundamental building blocks. Nodes are created automatically at edge endpoints.

```sql
-- Add isolated nodes
SELECT topology.ST_AddIsoNode('city_topo', NULL, ST_Point(-73.98, 40.75));

-- Add an edge between two points (creates nodes if needed)
SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.99, 40.74), ST_Point(-73.98, 40.74)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.98, 40.74), ST_Point(-73.98, 40.75)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.98, 40.75), ST_Point(-73.99, 40.75)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.99, 40.75), ST_Point(-73.99, 40.74)));
```

### Inspecting Primitives

```sql
-- List nodes
SELECT node_id, ST_AsText(geom) FROM city_topo.node;

-- List edges
SELECT edge_id, ST_AsText(geom) FROM city_topo.edge_data;

-- List faces (face_id 0 is the universal face)
SELECT face_id, ST_AsText(mbr) FROM city_topo.face;

-- Get the geometry of a face
SELECT topology.ST_GetFaceGeometry('city_topo', 1);
```

--------

## TopoGeometry

A TopoGeometry is a spatial object defined by references to topology primitives rather than coordinates. This ensures shared boundaries stay consistent.

### Creating a TopoGeometry Layer

```sql
-- Create a table with a topogeometry column
CREATE TABLE city_topo.blocks (
    id serial PRIMARY KEY,
    name text
);

-- Register a topogeometry column (returns a layer_id)
SELECT topology.AddTopoGeometryColumn('city_topo', 'city_topo', 'blocks', 'topo', 'POLYGON');
```

### Assigning TopoGeometry Values

```sql
-- Create a TopoGeometry from a regular geometry (snaps to existing topology)
INSERT INTO city_topo.blocks (name, topo) VALUES
    ('Block A', topology.toTopoGeom(
        ST_GeomFromText('POLYGON((-73.99 40.74,-73.98 40.74,-73.98 40.75,-73.99 40.75,-73.99 40.74))', 4326),
        'city_topo', 1));

-- Convert a TopoGeometry back to a regular geometry
SELECT name, topo::geometry FROM city_topo.blocks;
```

### Validating Topology

```sql
-- Validate the entire topology
SELECT * FROM topology.ValidateTopology('city_topo');

-- Check for topology errors
SELECT error FROM topology.ValidateTopology('city_topo')
WHERE error IS NOT NULL;
```

--------

## Cleanup

```sql
-- Drop a topology and all its objects
SELECT topology.DropTopology('city_topo');
```
