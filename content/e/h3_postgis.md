---
title: "h3_postgis"
linkTitle: "h3_postgis"
description: "H3 PostGIS integration"
weight: 1531
categories: ["GIS"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_h3**](https://github.com/postgis/h3-pg) : H3 PostGIS integration


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1531** | {{< badge content="h3_postgis" link="https://github.com/postgis/h3-pg" >}} | {{< ext "h3_postgis" "pg_h3" >}} | `4.5.0` | {{< category "GIS" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "h3" >}} {{< ext "postgis" >}} {{< ext "postgis_raster" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "qdgc" >}} {{< ext "pg_geohash" >}} {{< ext "pgrouting" >}} {{< ext "q3c" >}} {{< ext "pg_polyline" >}} {{< ext "pg_eviltransform" >}} {{< ext "earthdistance" >}} {{< ext "mobilitydb" >}} |
|    **Siblings**   | {{< ext "h3" >}} |

> [!Note] RPM and DEB 4.5.0 require h3, postgis and postgis_raster; point coordinates use longitude, latitude.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_h3` | `h3`, `postgis`, `postgis_raster` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.0` | {{< bg "18" "h3-pg_18" "green" >}} {{< bg "17" "h3-pg_17" "green" >}} {{< bg "16" "h3-pg_16" "green" >}} {{< bg "15" "h3-pg_15" "green" >}} {{< bg "14" "h3-pg_14" "green" >}} | `h3-pg_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.0` | {{< bg "18" "postgresql-18-h3" "green" >}} {{< bg "17" "postgresql-17-h3" "green" >}} {{< bg "16" "postgresql-16-h3" "green" >}} {{< bg "15" "postgresql-15-h3" "green" >}} {{< bg "14" "postgresql-14-h3" "green" >}} | `postgresql-$v-h3` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "h3-pg_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-18-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-17-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-16-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-15-h3 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 4.5.0" "postgresql-14-h3 : AVAIL 3" "green" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgis/h3-pg" title="Repository" icon="github" subtitle="github.com/postgis/h3-pg" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="h3-pg-4.5.0.tar.gz h3-4.5.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_h3;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_h3;		# install via package name, for the active PG version
pig install h3_postgis;		# install by extension name, for the current active PG version

pig install h3_postgis -v 18;   # install for PG 18
pig install h3_postgis -v 17;   # install for PG 17
pig install h3_postgis -v 16;   # install for PG 16
pig install h3_postgis -v 15;   # install for PG 15
pig install h3_postgis -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION h3_postgis CASCADE; -- requires h3, postgis, postgis_raster
```

## Usage

Sources:

- [4.5.0 PostGIS API](https://github.com/postgis/h3-pg/blob/v4.5.0/docs/api.md)
- [Dependencies and extension definition](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/CMakeLists.txt)
- [4.5.0 migration SQL](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/sql/updates/h3_postgis--4.2.3--4.5.0.sql)
- [4.5.0 release](https://github.com/postgis/h3-pg/releases/tag/v4.5.0)

`h3_postgis` bridges H3 cells with PostGIS geometry, geography, and raster. It requires `h3`, `postgis`, and `postgis_raster`, including for geometry-only use. Input geometries must use SRID 4326 with longitude, latitude coordinates; the functions do not reproject inputs.

### Convert Points and Cells

```sql
CREATE EXTENSION h3_postgis CASCADE;
SET h3.strict = true;

SELECT h3_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(-122.0553238, 37.3615593), 4326), 9
);
SELECT h3_cell_to_geometry('85283473fffffff'::h3index);
SELECT h3_cell_to_boundary_geometry('85283473fffffff'::h3index);
```

Transform other coordinate systems to SRID 4326 with PostGIS before calling the H3 functions. Setting an SRID label alone does not transform coordinates.

### Core API

| Task | Functions |
| --- | --- |
| Point to cell | `h3_latlng_to_cell(geometry, integer)`, `h3_latlng_to_cell(geography, integer)` |
| Cell center | `h3_cell_to_geometry`, `h3_cell_to_geography` |
| Cell boundary | `h3_cell_to_boundary_geometry`, `h3_cell_to_boundary_geography` |
| Polygon coverage | `h3_polygon_to_cells`, `h3_cells_to_multi_polygon_geometry`, `h3_cells_to_multi_polygon_geography` |
| Continuous raster summaries | `h3_raster_summary`, `h3_raster_summary_stats_agg` |
| Categorical raster summaries | `h3_raster_class_summary`, `h3_raster_class_summary_item_agg` |

The geometry-to-resolution `@` operator also maps a location to an H3 cell. Validate polygon inputs with `ST_IsValid()`; repairs through `ST_MakeValid()` can change topology and produce geometry collections, so retain and review polygonal components before coverage calculations. Invalid polygons have undefined behavior.

### Summarize Raster Data

```sql
SELECT (summary).h3,
       (h3_raster_summary_stats_agg((summary).stats)).*
FROM (
    SELECT h3_raster_summary(rast, 8) AS summary
    FROM rasters
) AS r
GROUP BY (summary).h3;
```

The default summary chooses a method; explicit clip, centroid, and subpixel variants let you control how raster pixels are assigned to cells. Review the choice against raster resolution and the requested H3 resolution.

### Upgrade and Boundaries

```sql
ALTER EXTENSION h3 UPDATE TO '4.5.0';
ALTER EXTENSION h3_postgis UPDATE TO '4.5.0';
```

The base extension update performs affected btree rebuilds and distance-dependent refreshes, so plan its maintenance window before updating the companion. Version 4.5.0 fixes restricted-search-path maintenance on PostgreSQL 17+, expression-index dump/restore, and several geometry and polygonization errors. Both extension versions should match. Install the matching 4.5.0 package files before updating either extension.

`h3.extend_antimeridian` should normally remain false for planar overlays. Both extensions are relocatable; ensure the schemas containing H3 and PostGIS objects are visible when issuing unqualified SQL. Neither extension requires shared preloading.
