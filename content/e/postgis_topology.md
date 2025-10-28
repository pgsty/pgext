---
title: "postgis_topology"
linkTitle: "postgis_topology"
description: "PostGIS topology spatial types and functions"
weight: 1501
categories: ["GIS"]
width: full
---

PostGIS topology spatial types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1501** | {{< badge content="postgis_topology" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_topology" "postgis" >}} | `3.6.0` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< bg "18" "postgis35_18*" "green" >}} {{< bg "17" "postgis35_17*" "green" >}} {{< bg "16" "postgis35_16*" "green" >}} {{< bg "15" "postgis35_15*" "green" >}} {{< bg "14" "postgis35_14*" "green" >}} {{< bg "13" "postgis35_13*" "green" >}} | `postgis35_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< bg "18" "postgresql-18-postgis-3 postgresql-18-postgis-3-scripts" "green" >}} {{< bg "17" "postgresql-17-postgis-3 postgresql-17-postgis-3-scripts" "green" >}} {{< bg "16" "postgresql-16-postgis-3 postgresql-16-postgis-3-scripts" "green" >}} {{< bg "15" "postgresql-15-postgis-3 postgresql-15-postgis-3-scripts" "green" >}} {{< bg "14" "postgresql-14-postgis-3 postgresql-14-postgis-3-scripts" "green" >}} {{< bg "13" "postgresql-13-postgis-3 postgresql-13-postgis-3-scripts" "green" >}} | `postgresql-$v-postgis-3 postgresql-$v-postgis-3-scripts` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 1" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 1" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 3.6.0" "postgis36_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.0" "postgis36_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 3.6.0" "postgresql-18-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-17-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-16-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-15-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-14-postgis-3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.6.0" "postgresql-13-postgis-3 : AVAIL 1" "blue" >}} |


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
pig ext install postgis_topology; # install by extension name, for the current active PG version
pig ext install postgis; # install via package alias, for the active PG version
pig ext install postgis_topology -v 18;   # install for PG 18
pig ext install postgis_topology -v 17;   # install for PG 17
pig ext install postgis_topology -v 16;   # install for PG 16
pig ext install postgis_topology -v 15;   # install for PG 15
pig ext install postgis_topology -v 14;   # install for PG 14
pig ext install postgis_topology -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION postgis_topology CASCADE SCHEMA topology;
```

