---
title: "postgis_tiger_geocoder"
linkTitle: "postgis_tiger_geocoder"
description: "PostGIS tiger geocoder and reverse geocoder"
weight: 1504
categories: ["Gis"]
width: full
---

PostGIS tiger geocoder and reverse geocoder

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1504** | {{< badge content="postgis_tiger_geocoder" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_tiger_geocoder" "postgis" >}} | `3.6.0` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} {{< ext "fuzzystrmatch" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_topology" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgis35_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/postgis" >}} | `3.6.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-postgis-3 postgresql-$v-postgis-3-scripts` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "postgis36_18" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgis36_18-3.6.0-1PGDG.rhel8.1.x86_64.rpm" >}} | {{< pkg "postgis36_17" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgis36_17-3.6.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgis36_16" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgis36_16-3.6.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgis36_15" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgis36_15-3.6.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgis36_14" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgis36_14-3.6.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "postgis36_18" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.0-1PGDG.rhel8.1.aarch64.rpm" >}} | {{< pkg "postgis36_17" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgis36_16" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgis36_15" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgis36_14" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "postgis36_18" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgis36_17" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgis36_16" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgis36_15" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgis36_14" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-4PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "postgis36_18" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgis36_17" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgis36_16" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgis36_15" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgis36_14" "3.6.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-4PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-postgis-3" "3.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_18` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_18-3.6.0-1PGDG.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgis36_18-3.6.0-1PGDG.rhel8.1.x86_64.rpm) |
| `postgis36_18` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_18-3.6.0-1PGDG.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgis36_18-3.6.0-1PGDG.rhel8.1.aarch64.rpm) |
| `postgis36_18` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-1PGDG.rhel9.1.aarch64.rpm) |
| `postgis36_18` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis36_18-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_18` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_18-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_18` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_18-3.6.0-1PGDG.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis36_18-3.6.0-1PGDG.rhel9.1.x86_64.rpm) |
| `postgresql-18-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-18-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.3 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-18-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.5 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.6 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.4 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.7 MiB | [postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-18-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_17` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_17-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgis36_17-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_17` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_17-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgis36_17-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_17` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_17` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_17` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_17-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis36_17-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_17` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_17-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis36_17-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-17-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.3 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-17-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.6 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.7 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.4 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.6 MiB | [postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-17-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_16` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_16-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgis36_16-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_16` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_16-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgis36_16-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_16` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_16` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis36_16-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_16` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_16-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_16` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_16-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis36_16-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-16-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.3 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-16-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.7 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.6 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.6 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.4 MiB | [postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-16-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_15` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_15-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgis36_15-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_15` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_15-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgis36_15-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_15` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_15` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgis36_15-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_15` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_15-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_15` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_15-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgis36_15-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.2 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-15-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-15-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.5 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.6 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.3 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.4 MiB | [postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-15-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_14` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_14-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgis36_14-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_14` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_14-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgis36_14-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_14` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_14` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgis36_14-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgis36_14` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_14-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_14` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_14-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgis36_14-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.2 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-14-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-14-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.6 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.5 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.4 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.3 MiB | [postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-14-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgis36_13` | 3.6.0 | `el8.x86_64` | pgdg | 5.1 MiB | [postgis36_13-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgis36_13-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `postgis36_13` | 3.6.0 | `el8.aarch64` | pgdg | 5.0 MiB | [postgis36_13-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgis36_13-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `postgis36_13` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_13-3.6.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgis36_13-3.6.0-4PGDG.rhel9.aarch64.rpm) |
| `postgis36_13` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_13-3.6.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgis36_13-3.6.0-4PGDG.rhel9.x86_64.rpm) |
| `postgis36_13` | 3.6.0 | `el9.x86_64` | pgdg | 4.2 MiB | [postgis36_13-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgis36_13-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgis36_13` | 3.6.0 | `el9.aarch64` | pgdg | 4.2 MiB | [postgis36_13-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgis36_13-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-postgis-3` | 3.6.0 | `d12.aarch64` | pgdg | 3.2 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg12+1_arm64.deb) |
| `postgresql-13-postgis-3` | 3.6.0 | `d12.x86_64` | pgdg | 3.3 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg12+1_amd64.deb) |
| `postgresql-13-postgis-3` | 3.6.0 | `u22.x86_64` | pgdg | 3.6 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-postgis-3` | 3.6.0 | `u22.aarch64` | pgdg | 3.5 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-postgis-3` | 3.6.0 | `u24.x86_64` | pgdg | 3.3 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-postgis-3` | 3.6.0 | `u24.aarch64` | pgdg | 5.4 MiB | [postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis/postgresql-13-postgis-3_3.6.0+dfsg-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

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
pig ext install postgis_tiger_geocoder; # install by extension name, for the current active PG version
pig ext install postgis; # install via package alias, for the active PG version
pig ext install postgis_tiger_geocoder -v 18;   # install for PG 18
pig ext install postgis_tiger_geocoder -v 17;   # install for PG 17
pig ext install postgis_tiger_geocoder -v 16;   # install for PG 16
pig ext install postgis_tiger_geocoder -v 15;   # install for PG 15
pig ext install postgis_tiger_geocoder -v 14;   # install for PG 14
pig ext install postgis_tiger_geocoder -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION postgis_tiger_geocoder CASCADE SCHEMA tiger;
```

