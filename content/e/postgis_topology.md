---
title: "postgis_topology"
linkTitle: "postgis_topology"
description: "PostGIS topology spatial types and functions"
weight: 1501
categories: ["GIS"]
width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : PostGIS topology spatial types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1501** | {{< badge content="postgis_topology" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_topology" "postgis" >}} | `3.6.2` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `topology` |
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | `postgis` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.2` | {{< bg "18" "postgis36_18" "green" >}} {{< bg "17" "postgis36_17" "green" >}} {{< bg "16" "postgis36_16" "green" >}} {{< bg "15" "postgis36_15" "green" >}} {{< bg "14" "postgis36_14" "green" >}} | `postgis36_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.2` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.2" "postgis36_14 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.2" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.2" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://git.osgeo.org/gitea/postgis/postgis" title="Repository" icon="link" subtitle="git.osgeo.org/gitea/postgis/postgis" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install postgis;		# install via package name, for the active PG version
pig install postgis_topology;		# install by extension name, for the current active PG version

pig install postgis_topology -v 18;   # install for PG 18
pig install postgis_topology -v 17;   # install for PG 17
pig install postgis_topology -v 16;   # install for PG 16
pig install postgis_topology -v 15;   # install for PG 15
pig install postgis_topology -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgis_topology CASCADE; -- requires postgis
```



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
