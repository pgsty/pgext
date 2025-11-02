---
title: "pointcloud_postgis"
linkTitle: "pointcloud_postgis"
description: "integration for pointcloud LIDAR data and PostGIS geometry data"
weight: 1521
categories: ["GIS"]
width: full
---

integration for pointcloud LIDAR data and PostGIS geometry data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1521** | {{< badge content="pointcloud_postgis" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< ext "pointcloud_postgis" "pointcloud" >}} | `1.2.5` | {{< category "GIS" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} {{< ext "pointcloud" >}} |
|   **See Also**    | {{< ext "postgis_raster" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} {{< ext "h3" >}} |
|    **Siblings**   | {{< ext "pointcloud" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pointcloud" >}} | `1.2.5` | {{< bg "18" "pointcloud_18*" "green" >}} {{< bg "17" "pointcloud_17*" "green" >}} {{< bg "16" "pointcloud_16*" "green" >}} {{< bg "15" "pointcloud_15*" "green" >}} {{< bg "14" "pointcloud_14*" "green" >}} {{< bg "13" "pointcloud_13*" "green" >}} | `pointcloud_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pointcloud" >}} | `1.2.5` | {{< bg "18" "postgresql-18-pointcloud" "green" >}} {{< bg "17" "postgresql-17-pointcloud" "green" >}} {{< bg "16" "postgresql-16-pointcloud" "green" >}} {{< bg "15" "postgresql-15-pointcloud" "green" >}} {{< bg "14" "postgresql-14-pointcloud" "green" >}} {{< bg "13" "postgresql-13-pointcloud" "green" >}} | `postgresql-$v-pointcloud` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-13-pointcloud : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpointcloud/pointcloud" title="Repository" icon="github" subtitle="github.com/pgpointcloud/pointcloud" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pointcloud-1.2.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pointcloud_postgis; # get pointcloud_postgis source code
pig build dep pointcloud_postgis; # install build dependencies
pig build pkg pointcloud_postgis; # build extension rpm or deb
pig build ext pointcloud_postgis; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pointcloud_postgis; # install by extension name, for the current active PG version
pig ext install pointcloud; # install via package alias, for the active PG version
pig ext install pointcloud_postgis -v 18;   # install for PG 18
pig ext install pointcloud_postgis -v 17;   # install for PG 17
pig ext install pointcloud_postgis -v 16;   # install for PG 16
pig ext install pointcloud_postgis -v 15;   # install for PG 15
pig ext install pointcloud_postgis -v 14;   # install for PG 14
pig ext install pointcloud_postgis -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pointcloud_postgis;
```

