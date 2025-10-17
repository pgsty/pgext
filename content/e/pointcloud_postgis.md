---
title: "pointcloud_postgis"
linkTitle: "pointcloud_postgis"
description: "integration for pointcloud LIDAR data and PostGIS geometry data"
weight: 1521
categories: ["Gis"]
width: full
---

integration for pointcloud LIDAR data and PostGIS geometry data

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1521** | {{< badge content="pointcloud_postgis" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< ext "pointcloud_postgis" "pointcloud" >}} | `1.2.5` | {{< category "GIS" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} {{< ext "pointcloud" >}} |
|   **See Also**    | {{< ext "postgis_raster" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} {{< ext "h3" >}} |
|    **Siblings**   | {{< ext "pointcloud" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pointcloud" >}} | `1.2.5` | {{< badge content="18" color="red" alt="pointcloud_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pointcloud_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pointcloud" >}} | `1.2.5` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pointcloud` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pointcloud_18" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pointcloud_18-1.2.5-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pointcloud_17" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pointcloud_17-1.2.5-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pointcloud_16" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pointcloud_16-1.2.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pointcloud_15" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pointcloud_15-1.2.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pointcloud_14" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pointcloud_14-1.2.5-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pointcloud_18" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pointcloud_18-1.2.5-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pointcloud_17" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pointcloud_17-1.2.5-2PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pointcloud_16" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pointcloud_16-1.2.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pointcloud_15" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pointcloud_15-1.2.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pointcloud_14" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pointcloud_14-1.2.5-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pointcloud_18" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pointcloud_18-1.2.5-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pointcloud_17" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pointcloud_17-1.2.5-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pointcloud_16" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pointcloud_16-1.2.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pointcloud_15" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pointcloud_15-1.2.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pointcloud_14" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pointcloud_14-1.2.5-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pointcloud_18" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pointcloud_18-1.2.5-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pointcloud_17" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pointcloud_17-1.2.5-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pointcloud_16" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pointcloud_16-1.2.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pointcloud_15" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pointcloud_15-1.2.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pointcloud_14" "1.2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pointcloud_14-1.2.5-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pointcloud" "1.2.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_18` | 1.2.5 | `el8.aarch64` | pgdg | 65.9 KiB | [pointcloud_18-1.2.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pointcloud_18-1.2.5-3PGDG.rhel8.aarch64.rpm) |
| `pointcloud_18` | 1.2.5 | `el8.x86_64` | pgdg | 67.8 KiB | [pointcloud_18-1.2.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pointcloud_18-1.2.5-3PGDG.rhel8.x86_64.rpm) |
| `pointcloud_18` | 1.2.5 | `el9.aarch64` | pgdg | 68.0 KiB | [pointcloud_18-1.2.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pointcloud_18-1.2.5-3PGDG.rhel9.aarch64.rpm) |
| `pointcloud_18` | 1.2.5 | `el9.x86_64` | pgdg | 69.2 KiB | [pointcloud_18-1.2.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pointcloud_18-1.2.5-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.5 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-18-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 97.6 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-18-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 97.9 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 94.8 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.1 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.2 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_17` | 1.2.5 | `el8.x86_64` | pgdg | 67.9 KiB | [pointcloud_17-1.2.5-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pointcloud_17-1.2.5-2PGDG.rhel8.x86_64.rpm) |
| `pointcloud_17` | 1.2.5 | `el8.aarch64` | pgdg | 65.8 KiB | [pointcloud_17-1.2.5-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pointcloud_17-1.2.5-2PGDG.rhel8.aarch64.rpm) |
| `pointcloud_17` | 1.2.5 | `el9.aarch64` | pgdg | 68.2 KiB | [pointcloud_17-1.2.5-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pointcloud_17-1.2.5-2PGDG.rhel9.aarch64.rpm) |
| `pointcloud_17` | 1.2.5 | `el9.x86_64` | pgdg | 69.5 KiB | [pointcloud_17-1.2.5-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pointcloud_17-1.2.5-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 97.7 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-17-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.6 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-17-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 104.2 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 107.4 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.2 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.0 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_16` | 1.2.5 | `el8.x86_64` | pgdg | 67.8 KiB | [pointcloud_16-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pointcloud_16-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_16` | 1.2.5 | `el8.aarch64` | pgdg | 65.7 KiB | [pointcloud_16-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pointcloud_16-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_16` | 1.2.5 | `el9.x86_64` | pgdg | 69.1 KiB | [pointcloud_16-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pointcloud_16-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_16` | 1.2.5 | `el9.aarch64` | pgdg | 67.7 KiB | [pointcloud_16-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pointcloud_16-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.6 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-16-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 97.7 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-16-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 104.2 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 107.3 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.1 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.2 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_15` | 1.2.5 | `el8.aarch64` | pgdg | 65.8 KiB | [pointcloud_15-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pointcloud_15-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_15` | 1.2.5 | `el8.x86_64` | pgdg | 67.9 KiB | [pointcloud_15-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pointcloud_15-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_15` | 1.2.5 | `el9.x86_64` | pgdg | 69.5 KiB | [pointcloud_15-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pointcloud_15-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_15` | 1.2.5 | `el9.aarch64` | pgdg | 68.4 KiB | [pointcloud_15-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pointcloud_15-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 98.0 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-15-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.7 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-15-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 104.4 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 107.6 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.4 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.5 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_14` | 1.2.5 | `el8.x86_64` | pgdg | 67.9 KiB | [pointcloud_14-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pointcloud_14-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_14` | 1.2.5 | `el8.aarch64` | pgdg | 65.8 KiB | [pointcloud_14-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pointcloud_14-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_14` | 1.2.5 | `el9.x86_64` | pgdg | 69.5 KiB | [pointcloud_14-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pointcloud_14-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_14` | 1.2.5 | `el9.aarch64` | pgdg | 68.2 KiB | [pointcloud_14-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pointcloud_14-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 98.0 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-14-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.8 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-14-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 107.6 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 104.5 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.5 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.4 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pointcloud_13` | 1.2.5 | `el8.aarch64` | pgdg | 65.8 KiB | [pointcloud_13-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pointcloud_13-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_13` | 1.2.5 | `el8.x86_64` | pgdg | 67.4 KiB | [pointcloud_13-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pointcloud_13-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_13` | 1.2.5 | `el9.aarch64` | pgdg | 68.1 KiB | [pointcloud_13-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pointcloud_13-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `pointcloud_13` | 1.2.5 | `el9.x86_64` | pgdg | 69.6 KiB | [pointcloud_13-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pointcloud_13-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-pointcloud` | 1.2.5 | `d12.aarch64` | pgdg | 94.6 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-13-pointcloud` | 1.2.5 | `d12.x86_64` | pgdg | 98.1 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-13-pointcloud` | 1.2.5 | `u22.aarch64` | pgdg | 104.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pointcloud` | 1.2.5 | `u22.x86_64` | pgdg | 107.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pointcloud` | 1.2.5 | `u24.x86_64` | pgdg | 96.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pointcloud` | 1.2.5 | `u24.aarch64` | pgdg | 94.4 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpointcloud/pointcloud" title="Repository" icon="github" subtitle="github.com/pgpointcloud/pointcloud" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pointcloud-1.2.5.tar.gz" >}}
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

