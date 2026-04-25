---
title: "postgis"
linkTitle: "postgis"
description: "PostGIS geometry and geography spatial types and functions"
weight: 1500
categories: ["GIS"]
width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : PostGIS geometry and geography spatial types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1500** | {{< badge content="postgis" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis" >}} | `3.6.3` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "h3_postgis" >}} {{< ext "mobilitydb" >}} {{< ext "pgrouting" >}} {{< ext "pointcloud_postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "postgis_topology" >}} {{< ext "pg_eviltransform" >}} |
|   **See Also**    | {{< ext "pointcloud" >}} {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "geoip" >}} {{< ext "pg_polyline" >}} {{< ext "earthdistance" >}} {{< ext "ogr_fdw" >}} {{< ext "tzf" >}} |
|    **Siblings**   | {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgis36_18" "green" >}} {{< bg "17" "postgis36_17" "green" >}} {{< bg "16" "postgis36_16" "green" >}} {{< bg "15" "postgis36_15" "green" >}} {{< bg "14" "postgis36_14" "green" >}} | `postgis36_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis36_18` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_18-3.6.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgis36_18-3.6.3-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_18` | `3.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [postgis36_18-3.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgis36_18-3.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_18-3.6.0-1PGDG.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgis36_18-3.6.0-1PGDG.rhel8.1.x86_64.rpm) |
| `postgis36_18` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_18-3.6.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.3-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_18` | `3.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.1 MiB | [postgis36_18-3.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_18-3.6.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.0-6PGDG.rhel8.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_18-3.6.0-1PGDG.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.0-1PGDG.rhel8.1.aarch64.rpm) |
| `postgis36_18` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.3-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_18` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.2-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_18` | `3.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.1-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-6PGDG.rhel9.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-1PGDG.rhel9.1.x86_64.rpm) |
| `postgis36_18` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.3-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_18` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.2-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_18` | `3.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.1-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-6PGDG.rhel9.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-1PGDG.rhel9.1.aarch64.rpm) |
| `postgis36_18` | `3.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_18-3.6.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis36_18-3.6.3-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_18` | `3.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_18-3.6.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis36_18-3.6.2-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_18` | `3.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_18-3.6.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis36_18-3.6.1-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis36_18-3.6.0-4PGDG.rhel10.x86_64.rpm) |
| `postgis36_18` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis36_18-3.6.0-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_18` | `3.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis36_18-3.6.3-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_18` | `3.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis36_18-3.6.2-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_18` | `3.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis36_18-3.6.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis36_18-3.6.0-4PGDG.rhel10.aarch64.rpm) |
| `postgis36_18` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis36_18-3.6.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | `3.6.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis36_17` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_17-3.6.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgis36_17-3.6.3-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_17` | `3.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [postgis36_17-3.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgis36_17-3.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_17-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgis36_17-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_17` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_17-3.6.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.3-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_17` | `3.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.1 MiB | [postgis36_17-3.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_17-3.6.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.0-6PGDG.rhel8.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_17-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_17` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.3-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_17` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.2-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_17` | `3.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.1-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-6PGDG.rhel9.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_17` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.3-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_17` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.2-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_17` | `3.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.1-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-6PGDG.rhel9.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_17` | `3.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_17-3.6.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis36_17-3.6.3-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_17` | `3.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_17-3.6.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis36_17-3.6.2-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_17` | `3.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_17-3.6.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis36_17-3.6.1-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis36_17-3.6.0-4PGDG.rhel10.x86_64.rpm) |
| `postgis36_17` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis36_17-3.6.0-1PGDG.rhel10.x86_64.rpm) |
| `postgis36_17` | `3.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis36_17-3.6.3-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_17` | `3.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis36_17-3.6.2-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_17` | `3.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis36_17-3.6.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis36_17-3.6.0-4PGDG.rhel10.aarch64.rpm) |
| `postgis36_17` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis36_17-3.6.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.5 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | `3.6.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.7 MiB | [postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis36_16` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_16-3.6.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgis36_16-3.6.3-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_16` | `3.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [postgis36_16-3.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgis36_16-3.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_16-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgis36_16-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_16` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_16-3.6.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.3-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_16` | `3.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.1 MiB | [postgis36_16-3.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_16-3.6.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.0-6PGDG.rhel8.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_16-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_16` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.3-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_16` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.2-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_16` | `3.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.1-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-6PGDG.rhel9.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_16` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.3-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_16` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.2-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_16` | `3.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.1-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-6PGDG.rhel9.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_16` | `3.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_16-3.6.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis36_16-3.6.3-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_16` | `3.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_16-3.6.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis36_16-3.6.2-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_16` | `3.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_16-3.6.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis36_16-3.6.1-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis36_16-3.6.0-4PGDG.rhel10.x86_64.rpm) |
| `postgis36_16` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis36_16-3.6.0-1PGDG.rhel10.x86_64.rpm) |
| `postgis36_16` | `3.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis36_16-3.6.3-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_16` | `3.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis36_16-3.6.2-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_16` | `3.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis36_16-3.6.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis36_16-3.6.0-4PGDG.rhel10.aarch64.rpm) |
| `postgis36_16` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis36_16-3.6.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.7 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.6 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | `3.6.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis36_15` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_15-3.6.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgis36_15-3.6.3-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_15` | `3.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [postgis36_15-3.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgis36_15-3.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_15-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgis36_15-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_15` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_15-3.6.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.3-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_15` | `3.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.1 MiB | [postgis36_15-3.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_15-3.6.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.0-6PGDG.rhel8.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_15-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_15` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.3 MiB | [postgis36_15-3.6.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.3-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_15` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.2-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_15` | `3.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.1-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-6PGDG.rhel9.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_15` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.3-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_15` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.2-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_15` | `3.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.1-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-6PGDG.rhel9.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_15` | `3.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_15-3.6.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis36_15-3.6.3-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_15` | `3.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_15-3.6.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis36_15-3.6.2-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_15` | `3.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_15-3.6.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis36_15-3.6.1-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis36_15-3.6.0-4PGDG.rhel10.x86_64.rpm) |
| `postgis36_15` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgis36_15-3.6.0-1PGDG.rhel10.x86_64.rpm) |
| `postgis36_15` | `3.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis36_15-3.6.3-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_15` | `3.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis36_15-3.6.2-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_15` | `3.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis36_15-3.6.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis36_15-3.6.0-4PGDG.rhel10.aarch64.rpm) |
| `postgis36_15` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgis36_15-3.6.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.4 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.5 MiB | [postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | `3.6.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.5 MiB | [postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis36_14` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_14-3.6.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgis36_14-3.6.3-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_14` | `3.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [postgis36_14-3.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgis36_14-3.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.1 MiB | [postgis36_14-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgis36_14-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_14` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_14-3.6.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.3-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_14` | `3.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.1 MiB | [postgis36_14-3.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_14-3.6.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.0-6PGDG.rhel8.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.0 MiB | [postgis36_14-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_14` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.3 MiB | [postgis36_14-3.6.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.3-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_14` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.2-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_14` | `3.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.1-1PGDG.rhel9.7.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-6PGDG.rhel9.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_14` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.3-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_14` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.2-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_14` | `3.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.1-1PGDG.rhel9.7.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-6PGDG.rhel9.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_14` | `3.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_14-3.6.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis36_14-3.6.3-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_14` | `3.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_14-3.6.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis36_14-3.6.2-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_14` | `3.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.3 MiB | [postgis36_14-3.6.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis36_14-3.6.1-1PGDG.rhel10.1.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis36_14-3.6.0-4PGDG.rhel10.x86_64.rpm) |
| `postgis36_14` | `3.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgis36_14-3.6.0-1PGDG.rhel10.x86_64.rpm) |
| `postgis36_14` | `3.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis36_14-3.6.3-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_14` | `3.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis36_14-3.6.2-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_14` | `3.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis36_14-3.6.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis36_14-3.6.0-4PGDG.rhel10.aarch64.rpm) |
| `postgis36_14` | `3.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgis36_14-3.6.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-postgis-3` | `3.6.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg13+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg13+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 3.6 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 3.5 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | `3.6.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.3+dfsg-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | `3.6.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.2+dfsg-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

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

pig install postgis -v 18;   # install for PG 18
pig install postgis -v 17;   # install for PG 17
pig install postgis -v 16;   # install for PG 16
pig install postgis -v 15;   # install for PG 15
pig install postgis -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgis;
```

## Usage

Source: [Official manual](https://postgis.net/documentation/manual/), [current manual HTML](https://postgis.net/docs/postgis-en.html), [release notes](https://postgis.net/docs/release_notes.html), [patch release announcement](https://postgis.net/2026/04/PostGIS-Patch-Releases/)

`postgis` adds spatial types, indexes, and SQL functions to PostgreSQL. The main user-facing split is between `geometry` for planar/projected work and `geography` for spherical calculations on longitude/latitude data.

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
- `PostGIS 3.6.3` is a patch release dated 2026-04-14. The release notes describe fixes and a security hardening change, not a new stub-level usage surface, so this refresh mostly trims and aligns the stub with the current manual.
