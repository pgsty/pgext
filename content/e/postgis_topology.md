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
| **1501** | {{< badge content="postgis_topology" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_topology" "postgis" >}} | `3.6.1` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `postgis` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.1` | {{< bg "18" "postgis36_18" "green" >}} {{< bg "17" "postgis36_17" "green" >}} {{< bg "16" "postgis36_16" "green" >}} {{< bg "15" "postgis36_15" "green" >}} {{< bg "14" "postgis36_14" "green" >}} {{< bg "13" "postgis36_13" "green" >}} | `postgis36_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.1` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} {{< bg "13" "postgresql-13-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.1" "postgis36_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.1" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.1" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |


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
pig install postgis_topology -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgis_topology CASCADE; -- requires postgis
```
