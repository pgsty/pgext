---
title: "postgis"
linkTitle: "postgis"
description: "PostGIS geometry and geography spatial types and functions"
weight: 1500
categories: ["GIS"]
languages: ["C"]
licenses: ["GPL-2.0"]
repos: ["PGDG"]
page_width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : PostGIS geometry and geography spatial types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1500** | {{< badge content="postgis" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis" >}} | `3.7.0` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "h3_postgis" >}} {{< ext "mobilitydb" >}} {{< ext "pg_eviltransform" >}} {{< ext "pgh_consistency" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output" >}} {{< ext "pgh_output_en_au" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_raster" >}} {{< ext "pghydro" >}} {{< ext "pgrouting" >}} {{< ext "pointcloud_postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "postgis_topology" >}} {{< ext "qdgc_postgis" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "h3" >}} {{< ext "mobilitydb" >}} {{< ext "ogr_fdw" >}} {{< ext "pointcloud" >}} {{< ext "q3c" >}} {{< ext "pg_sphere" >}} {{< ext "nominatim_fdw" >}} {{< ext "pg_polyline" >}} {{< ext "earthdistance" >}} {{< ext "pg_eviltransform" >}} {{< ext "pg_geohash" >}} |
|    **Siblings**   | {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} |

> [!Note] Pigsty fills EL8 gaps.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7.0` | {{< bg "18" "postgis37_18" "green" >}} {{< bg "17" "postgis37_17" "green" >}} {{< bg "16" "postgis37_16" "green" >}} {{< bg "15" "postgis37_15" "green" >}} {{< bg "14" "postgis37_14" "green" >}} | `postgis37_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7.0` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.7.0" "postgis37_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgis37_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgis37_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgis37_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgis37_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.7.0" "postgis37_14 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-postgis-3 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-postgis-3 : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis37_18` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.2 MiB | [postgis37_18-3.7.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgis37_18-3.7.0-1PGSTY.el8.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.0 MiB | [postgis37_18-3.7.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgis37_18-3.7.0-1PGSTY.el8.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis37_18-3.7.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis37_18-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis37_18-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis37_18-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis37_18-3.7.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis37_18-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis37_18-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis37_18-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis37_18-3.7.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis37_18-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis37_18-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis37_18-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis37_18-3.7.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis37_18-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis37_18-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_18` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_18-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis37_18-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-postgis-3` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 1.5 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 1.4 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 1.5 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 1.4 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 1.5 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 1.4 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 1.5 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 1.4 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 1.5 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 1.4 MiB | [postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis37_17` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.2 MiB | [postgis37_17-3.7.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgis37_17-3.7.0-1PGSTY.el8.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.0 MiB | [postgis37_17-3.7.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgis37_17-3.7.0-1PGSTY.el8.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis37_17-3.7.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis37_17-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis37_17-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis37_17-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis37_17-3.7.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis37_17-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis37_17-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis37_17-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis37_17-3.7.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis37_17-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis37_17-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis37_17-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis37_17-3.7.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis37_17-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis37_17-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_17` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_17-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis37_17-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-postgis-3` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 1.5 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 1.4 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 1.5 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 1.4 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 1.6 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 1.5 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 1.5 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 1.4 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 1.5 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 1.4 MiB | [postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis37_16` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.2 MiB | [postgis37_16-3.7.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgis37_16-3.7.0-1PGSTY.el8.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.0 MiB | [postgis37_16-3.7.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgis37_16-3.7.0-1PGSTY.el8.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis37_16-3.7.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis37_16-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis37_16-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis37_16-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis37_16-3.7.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis37_16-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis37_16-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis37_16-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis37_16-3.7.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis37_16-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis37_16-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis37_16-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis37_16-3.7.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis37_16-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis37_16-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_16` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_16-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis37_16-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-postgis-3` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 1.5 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 1.4 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 1.5 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 1.4 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 1.6 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 1.5 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 1.5 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 1.4 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 1.5 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 1.4 MiB | [postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis37_15` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.2 MiB | [postgis37_15-3.7.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgis37_15-3.7.0-1PGSTY.el8.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.0 MiB | [postgis37_15-3.7.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgis37_15-3.7.0-1PGSTY.el8.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis37_15-3.7.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis37_15-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis37_15-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis37_15-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis37_15-3.7.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis37_15-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis37_15-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis37_15-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.3 MiB | [postgis37_15-3.7.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis37_15-3.7.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.3 MiB | [postgis37_15-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis37_15-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis37_15-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis37_15-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis37_15-3.7.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis37_15-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis37_15-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_15` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_15-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis37_15-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-postgis-3` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 1.5 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 1.4 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 1.5 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 1.4 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 1.6 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 1.5 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 1.5 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 1.4 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 1.5 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 1.4 MiB | [postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.5 MiB | [postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.5 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis37_14` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.2 MiB | [postgis37_14-3.7.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgis37_14-3.7.0-1PGSTY.el8.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.0 MiB | [postgis37_14-3.7.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgis37_14-3.7.0-1PGSTY.el8.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis37_14-3.7.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis37_14-3.7.0-rc3_3PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis37_14-3.7.0-rc1_5PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis37_14-3.7.0-beta2_2PGDG.rhel9.8.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis37_14-3.7.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis37_14-3.7.0-rc3_3PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis37_14-3.7.0-rc1_5PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis37_14-3.7.0-beta2_2PGDG.rhel9.8.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.3 MiB | [postgis37_14-3.7.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis37_14-3.7.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.3 MiB | [postgis37_14-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis37_14-3.7.0-rc3_3PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis37_14-3.7.0-rc1_5PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis37_14-3.7.0-beta2_2PGDG.rhel10.2.x86_64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis37_14-3.7.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis37_14-3.7.0-rc3_3PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis37_14-3.7.0-rc1_5PGDG.rhel10.2.aarch64.rpm) |
| `postgis37_14` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 3.2 MiB | [postgis37_14-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis37_14-3.7.0-beta2_2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-postgis-3` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 1.5 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 1.4 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 1.5 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg13+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 1.4 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg13+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 1.6 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 1.5 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 1.5 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 1.4 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 1.5 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 1.4 MiB | [postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.7.0+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.4+dfsg-2.pgdg26.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://git.osgeo.org/gitea/postgis/postgis" title="Repository" icon="link" subtitle="git.osgeo.org/gitea/postgis/postgis" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgis-3.7.0.tar.gz postgis-3.7.0-en.pdf postgis37-filter-requires-perl-Pg.sh" />}}
{{< /cards >}}


```bash
pig build pkg postgis;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install postgis;		# install via package name, for the active PG version

pig install postgis -v 18;   # install for PG 18
pig install postgis -v 17;   # install for PG 17
pig install postgis -v 16;   # install for PG 16
pig install postgis -v 15;   # install for PG 15
pig install postgis -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION postgis;
```

## Usage

Sources:

- [PostGIS 3.7.0 release](https://postgis.net/2026/10/PostGIS-3.7.0/)
- [PostGIS 3.7 manual](https://postgis.net/docs/manual-3.7/)
- [Official release notes](https://postgis.net/docs/release_notes.html)
- [Current manual HTML](https://postgis.net/docs/postgis-en.html)

`postgis` 3.7.0 adds spatial types, indexes, and SQL functions to PostgreSQL. The main user-facing split is between `geometry` for planar/projected work and `geography` for spherical calculations on longitude/latitude data. This major release requires PostgreSQL 14 or later, GEOS 3.10 or later, and PROJ 6.1 or later.

### Basic setup

```sql
CREATE EXTENSION postgis;
SELECT PostGIS_Full_Version();
```

### Core types and functions

```sql
CREATE TABLE sensors (
  id bigserial PRIMARY KEY,
  geom geometry(Point, 4326),
  geog geography(Point, 4326)
);

SELECT ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326);
SELECT ST_Intersects(a.geom, b.geom) FROM a, b;
SELECT ST_DWithin(a.geom, b.geom, 100);
SELECT ST_Distance(a.geog, b.geog);
SELECT ST_Transform(geom, 3857) FROM sensors;
```

- constructors: `ST_MakePoint`, `ST_GeomFromText`, `ST_GeomFromGeoJSON`
- relationships: `ST_Intersects`, `ST_Contains`, `ST_Within`, `ST_DWithin`
- measurements and transforms: `ST_Distance`, `ST_Area`, `ST_Length`, `ST_Transform`
- processing: `ST_Buffer`, `ST_Intersection`, `ST_Union`

### Spatial indexes

```sql
CREATE INDEX idx_sensors_geom ON sensors USING GIST (geom);
```

The official manual continues to recommend GiST as the general-purpose spatial index, with BRIN and SP-GiST available for specific data distributions and tradeoffs.

### Caveats

- Use `geometry` in an appropriate projected SRID for planar distances and areas; use `geography` when you need meter-based spheroidal calculations.
- `PostGIS 3.7.0` is a major release. It drops PostgreSQL 12/13 and GEOS versions older than 3.10; review the topology tolerance changes and the move of `address_standardizer` and `tiger_geocoder` packaging before upgrading. After installing matching files, run `SELECT postgis_extensions_upgrade();` where the upstream upgrade procedure requires it.
