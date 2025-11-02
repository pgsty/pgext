---
title: "postgis_raster"
linkTitle: "postgis_raster"
description: "PostGIS raster types and functions"
weight: 1502
categories: ["GIS"]
width: full
---

PostGIS raster types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1502** | {{< badge content="postgis_raster" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_raster" "postgis" >}} | `3.6.0` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|    **Need By**    | {{< ext "h3_postgis" >}} |
|   **See Also**    | {{< ext "pointcloud_postgis" >}} {{< ext "pointcloud" >}} {{< ext "pgrouting" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} {{< ext "pg_polyline" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< bg "18" "postgis36_18*" "green" >}} {{< bg "17" "postgis36_17*" "green" >}} {{< bg "16" "postgis36_16*" "green" >}} {{< bg "15" "postgis36_15*" "green" >}} {{< bg "14" "postgis36_14*" "green" >}} {{< bg "13" "postgis36_13*" "green" >}} | `postgis36_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< bg "18" "postgresql-18-postgis-3 postgresql-18-postgis-3-scripts" "green" >}} {{< bg "17" "postgresql-17-postgis-3 postgresql-17-postgis-3-scripts" "green" >}} {{< bg "16" "postgresql-16-postgis-3 postgresql-16-postgis-3-scripts" "green" >}} {{< bg "15" "postgresql-15-postgis-3 postgresql-15-postgis-3-scripts" "green" >}} {{< bg "14" "postgresql-14-postgis-3 postgresql-14-postgis-3-scripts" "green" >}} {{< bg "13" "postgresql-13-postgis-3 postgresql-13-postgis-3-scripts" "green" >}} | `postgresql-$v-postgis-3 postgresql-$v-postgis-3-scripts` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://git.osgeo.org/gitea/postgis/postgis" title="Repository" icon="link" subtitle="git.osgeo.org/gitea/postgis/postgis" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install postgis_raster; # install by extension name, for the current active PG version
pig ext install postgis; # install via package alias, for the active PG version
pig ext install postgis_raster -v 18;   # install for PG 18
pig ext install postgis_raster -v 17;   # install for PG 17
pig ext install postgis_raster -v 16;   # install for PG 16
pig ext install postgis_raster -v 15;   # install for PG 15
pig ext install postgis_raster -v 14;   # install for PG 14
pig ext install postgis_raster -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION postgis_raster;
```

