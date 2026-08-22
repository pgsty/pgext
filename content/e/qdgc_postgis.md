---
title: "qdgc_postgis"
linkTitle: "qdgc_postgis"
description: "Add PostGIS geometry and geography bindings plus area-to-cell fills for QDGC."
weight: 1710
categories: ["GIS"]
languages: ["SQL"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**qdgc**](https://pgxn.org/dist/qdgc/0.1.0/) : Add PostGIS geometry and geography bindings plus area-to-cell fills for QDGC.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1710** | {{< badge content="qdgc_postgis" link="https://pgxn.org/dist/qdgc/0.1.0/" >}} | {{< ext "qdgc_postgis" "qdgc" >}} | `0.1.0` | {{< category "GIS" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "qdgc" >}} {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "pgrouting" >}} {{< ext "q3c" >}} {{< ext "pg_polyline" >}} {{< ext "pg_eviltransform" >}} {{< ext "earthdistance" >}} {{< ext "mobilitydb" >}} |
|    **Siblings**   | {{< ext "qdgc" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `qdgc` | `qdgc`, `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "qdgc_18" "green" >}} {{< bg "17" "qdgc_17" "green" >}} {{< bg "16" "qdgc_16" "green" >}} {{< bg "15" "qdgc_15" "green" >}} {{< bg "14" "qdgc_14" "green" >}} | `qdgc_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-qdgc" "green" >}} {{< bg "17" "postgresql-17-qdgc" "green" >}} {{< bg "16" "postgresql-16-qdgc" "green" >}} {{< bg "15" "postgresql-15-qdgc" "green" >}} {{< bg "14" "postgresql-14-qdgc" "green" >}} | `postgresql-$v-qdgc` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/qdgc/0.1.0/" title="Repository" icon="link" subtitle="pgxn.org/dist/qdgc/0.1.0/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="qdgc-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg qdgc;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install qdgc;		# install via package name, for the active PG version
pig install qdgc_postgis;		# install by extension name, for the current active PG version

pig install qdgc_postgis -v 18;   # install for PG 18
pig install qdgc_postgis -v 17;   # install for PG 17
pig install qdgc_postgis -v 16;   # install for PG 16
pig install qdgc_postgis -v 15;   # install for PG 15
pig install qdgc_postgis -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION qdgc_postgis CASCADE; -- requires qdgc, postgis
```

## Usage

Sources:

- [PGXN qdgc 0.1.0 release](https://pgxn.org/dist/qdgc/0.1.0/)
- [Official 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [Official qdgc_postgis control file](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis.control)
- [Official qdgc_postgis 0.1.0 extension SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis--0.1.0.sql)

`qdgc_postgis` 0.1.0 is the PostGIS companion to the pure-SQL `qdgc` core. It converts QDGC cells to and from PostGIS points and polygons, measures cell area on the WGS84 spheroid, and fills arbitrary geometries with QDGC cells. The extension requires both `qdgc` and `postgis`; it does not replace either one.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION qdgc;
CREATE EXTENSION qdgc_postgis;

SELECT qdgc_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(31.4, 2.7), 4326),
    5
);

SELECT qdgc_cell_to_geometry('E031N02ADBAC');
SELECT qdgc_cell_to_boundary_geometry('E031N02ADBAC');
SELECT qdgc_cell_area_km2('E031N02ADBAC');
```

The point overload transforms geometry with a nonzero, non-4326 SRID to EPSG:4326. An SRID of zero is assumed to already contain longitude and latitude.

### Fill an Area of Interest

Estimate the result size before producing a deep fill:

```sql
WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT qdgc_estimate_cell_count(geom, 7)
FROM area;

WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT cell
FROM area
CROSS JOIN LATERAL qdgc_polygon_to_cells(
    geom,
    7,
    'intersects'
) AS cell;
```

The predicate can be:

- `intersects`, the default, for cells intersecting the geometry;
- `centroid`, for cells whose center lies inside the geometry;
- `contains`, for cells wholly contained by the geometry.

The implementation descends a pruning quadtree instead of testing every cell in the geometry's full envelope. Multi-part geometries are filled per part and their cell sets are combined.

### Important Objects

- `qdgc_latlng_to_cell(geometry, level)` and its `geography` overload encode PostGIS points.
- `qdgc_cell_to_geometry` and `qdgc_cell_to_geography` return the cell centroid.
- `qdgc_cell_to_boundary_geometry` and `qdgc_cell_to_boundary_geography` return the rectangular cell boundary.
- `qdgc_cell_area_km2` measures the boundary geography on the WGS84 spheroid.
- `qdgc_polygon_to_cells` fills an area using one of the three documented predicates.
- `qdgc_estimate_cell_count` provides a cheap, envelope-capped guard before materializing a fill.

### Operational Notes

- `qdgc_postgis.control` declares `requires = 'qdgc,postgis'` and `relocatable = true`. Install PostGIS with an appropriately privileged role before delegating use of the companion extension.
- No `shared_preload_libraries`, `LOAD`, or restart is required. The extension is SQL-only, but its PostGIS dependency includes native code.
- Install `qdgc`, `qdgc_postgis`, and their callable dependencies into schemas visible on the active `search_path`, because the relocatable SQL calls functions by unqualified name.
- Upstream tests PostgreSQL 13 through 17. Do not infer PostgreSQL 18 support from the absence of compiled code.
- Deep area fills can still produce enormous sets even with pruning. Treat `qdgc_estimate_cell_count` as an operational guard and apply application-specific limits before executing `qdgc_polygon_to_cells`.

