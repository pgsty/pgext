---
title: "pointcloud"
linkTitle: "pointcloud"
description: "data type for lidar point clouds"
weight: 1520
categories: ["GIS"]
width: full
---

[**pointcloud**](https://github.com/pgpointcloud/pointcloud) : data type for lidar point clouds


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1520** | {{< badge content="pointcloud" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< ext "pointcloud" >}} | `1.2.5` | {{< category "GIS" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pointcloud_postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} |
|    **Siblings**   | {{< ext "pointcloud_postgis" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pointcloud` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.5` | {{< bg "18" "pointcloud_18" "green" >}} {{< bg "17" "pointcloud_17" "green" >}} {{< bg "16" "pointcloud_16" "green" >}} {{< bg "15" "pointcloud_15" "green" >}} {{< bg "14" "pointcloud_14" "green" >}} | `pointcloud_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.5` | {{< bg "18" "postgresql-18-pointcloud" "green" >}} {{< bg "17" "postgresql-17-pointcloud" "green" >}} {{< bg "16" "postgresql-16-pointcloud" "green" >}} {{< bg "15" "postgresql-15-pointcloud" "green" >}} {{< bg "14" "postgresql-14-pointcloud" "green" >}} | `postgresql-$v-pointcloud` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2.5" "pointcloud_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "pointcloud_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.2.5" "postgresql-18-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-17-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-16-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-15-pointcloud : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.5" "postgresql-14-pointcloud : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_18` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.8 KiB | [pointcloud_18-1.2.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pointcloud_18-1.2.5-3PGDG.rhel8.x86_64.rpm) |
| `pointcloud_18` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.9 KiB | [pointcloud_18-1.2.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pointcloud_18-1.2.5-3PGDG.rhel8.aarch64.rpm) |
| `pointcloud_18` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.2 KiB | [pointcloud_18-1.2.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pointcloud_18-1.2.5-3PGDG.rhel9.x86_64.rpm) |
| `pointcloud_18` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.0 KiB | [pointcloud_18-1.2.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pointcloud_18-1.2.5-3PGDG.rhel9.aarch64.rpm) |
| `pointcloud_18` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.1 KiB | [pointcloud_18-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pointcloud_18-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_18` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.4 KiB | [pointcloud_18-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pointcloud_18-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 97.6 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.5 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.2 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.1 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 97.9 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 94.8 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.1 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.2 KiB | [postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-18-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_17` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.9 KiB | [pointcloud_17-1.2.5-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pointcloud_17-1.2.5-2PGDG.rhel8.x86_64.rpm) |
| `pointcloud_17` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.8 KiB | [pointcloud_17-1.2.5-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pointcloud_17-1.2.5-2PGDG.rhel8.aarch64.rpm) |
| `pointcloud_17` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.5 KiB | [pointcloud_17-1.2.5-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pointcloud_17-1.2.5-2PGDG.rhel9.x86_64.rpm) |
| `pointcloud_17` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pointcloud_17-1.2.5-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pointcloud_17-1.2.5-2PGDG.rhel9.aarch64.rpm) |
| `pointcloud_17` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.2 KiB | [pointcloud_17-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pointcloud_17-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_17` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.5 KiB | [pointcloud_17-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pointcloud_17-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 97.7 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.6 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.1 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.0 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.4 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.2 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.0 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.2 KiB | [postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-17-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_16` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.8 KiB | [pointcloud_16-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pointcloud_16-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_16` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.7 KiB | [pointcloud_16-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pointcloud_16-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_16` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.1 KiB | [pointcloud_16-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pointcloud_16-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_16` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pointcloud_16-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pointcloud_16-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `pointcloud_16` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.2 KiB | [pointcloud_16-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pointcloud_16-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_16` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.5 KiB | [pointcloud_16-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pointcloud_16-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 97.7 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.6 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.2 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.1 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.3 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.2 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.1 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.2 KiB | [postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-16-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_15` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.9 KiB | [pointcloud_15-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pointcloud_15-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_15` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.8 KiB | [pointcloud_15-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pointcloud_15-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_15` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.5 KiB | [pointcloud_15-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pointcloud_15-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_15` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.4 KiB | [pointcloud_15-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pointcloud_15-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `pointcloud_15` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.7 KiB | [pointcloud_15-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pointcloud_15-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_15` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 70.0 KiB | [pointcloud_15-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pointcloud_15-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.0 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.7 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.7 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.2 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.6 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.4 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.4 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.5 KiB | [postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-15-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_14` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.9 KiB | [pointcloud_14-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pointcloud_14-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_14` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.8 KiB | [pointcloud_14-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pointcloud_14-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_14` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.5 KiB | [pointcloud_14-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pointcloud_14-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_14` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pointcloud_14-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pointcloud_14-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `pointcloud_14` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.7 KiB | [pointcloud_14-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pointcloud_14-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_14` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 70.0 KiB | [pointcloud_14-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pointcloud_14-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.0 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.8 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.6 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.2 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.6 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.5 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.4 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.5 KiB | [postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-14-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpointcloud/pointcloud" title="Repository" icon="github" subtitle="github.com/pgpointcloud/pointcloud" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pointcloud-1.2.5.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pointcloud;		# install via package name, for the active PG version

pig install pointcloud -v 18;   # install for PG 18
pig install pointcloud -v 17;   # install for PG 17
pig install pointcloud -v 16;   # install for PG 16
pig install pointcloud -v 15;   # install for PG 15
pig install pointcloud -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pointcloud;
```



## Usage

> [pgpointcloud/pointcloud: A PostgreSQL extension for storing point cloud (LIDAR) data](https://github.com/pgpointcloud/pointcloud)
>
> [Documentation](https://pgpointcloud.github.io/pointcloud/)

PostgreSQL Pointcloud stores point cloud (LIDAR) data in PostgreSQL. It introduces
two new data types: `PcPoint` for individual points and `PcPatch` for collections
of points. Data is organized by schema documents that describe the dimensions and
encoding of each point.

```sql
CREATE EXTENSION pointcloud;
-- For PostGIS integration:
CREATE EXTENSION pointcloud_postgis;
```


## Concepts

### Schemas

PostgreSQL Pointcloud uses a "schema document" to describe the contents of any
particular LIDAR point. Each point contains a number of dimensions, and each
dimension can be of any data type, with scaling and/or offsets applied to move
between the actual value and the value stored in the database. The schema
document format is the same one used by the [PDAL](https://pdal.io/) library.

Here is a simple 4-dimensional schema document:

```sql
INSERT INTO pointcloud_formats (pcid, srid, schema) VALUES (1, 4326,
'<?xml version="1.0" encoding="UTF-8"?>
<pc:PointCloudSchema xmlns:pc="http://pointcloud.org/schemas/PC/1.1"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <pc:dimension>
    <pc:position>1</pc:position>
    <pc:size>4</pc:size>
    <pc:description>X coordinate as a long integer. You must use the
                    scale and offset information of the header to
                    determine the double value.</pc:description>
    <pc:name>X</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>2</pc:position>
    <pc:size>4</pc:size>
    <pc:description>Y coordinate as a long integer. You must use the
                    scale and offset information of the header to
                    determine the double value.</pc:description>
    <pc:name>Y</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>3</pc:position>
    <pc:size>4</pc:size>
    <pc:description>Z coordinate as a long integer. You must use the
                    scale and offset information of the header to
                    determine the double value.</pc:description>
    <pc:name>Z</pc:name>
    <pc:interpretation>int32_t</pc:interpretation>
    <pc:scale>0.01</pc:scale>
  </pc:dimension>
  <pc:dimension>
    <pc:position>4</pc:position>
    <pc:size>2</pc:size>
    <pc:description>The intensity value is the integer representation
                    of the pulse return magnitude. This value is optional
                    and system specific. However, it should always be
                    included if available.</pc:description>
    <pc:name>Intensity</pc:name>
    <pc:interpretation>uint16_t</pc:interpretation>
    <pc:scale>1</pc:scale>
  </pc:dimension>
  <pc:metadata>
    <Metadata name="compression">dimensional</Metadata>
  </pc:metadata>
</pc:PointCloudSchema>');
```

Schema documents are stored in the `pointcloud_formats` table, along with a
`pcid` or "pointcloud identifier". Each object just has a `pcid`, which serves
as a key to find the schema in `pointcloud_formats`. This is similar to the way
the `srid` is resolved for spatial reference system support in PostGIS.

### Point Cloud Objects

**PcPoint**: The basic point cloud type. Every point has a (large) number of
dimensions, but at a minimum an X and Y coordinate. Points can be rendered in
JSON form using `PC_AsText(pcpoint)`:

```json
{
    "pcid" : 1,
      "pt" : [0.01, 0.02, 0.03, 4]
}
```

**PcPatch**: A collection of `PcPoint` grouped together. Instead of storing
billions of individual point records, LIDAR data is represented as a smaller
collection of `PcPatch` records. Patches are rendered with `PC_AsText(pcpatch)`:

```json
{
    "pcid" : 1,
     "pts" : [
              [0.02, 0.03, 0.05, 6],
              [0.02, 0.03, 0.05, 8]
             ]
}
```

### Tables

```sql
-- This requires the schema entry above so that pcid==1 exists.

-- A table of points
CREATE TABLE points (
    id SERIAL PRIMARY KEY,
    pt PCPOINT(1)
);

-- A table of patches
CREATE TABLE patches (
    id SERIAL PRIMARY KEY,
    pa PCPATCH(1)
);
```

Two system-provided tables are available:

- `pointcloud_formats`: holds all the pcid entries and schema documents
- `pointcloud_columns`: a view displaying all columns that contain point cloud objects

```sql
SELECT * FROM pointcloud_columns;

 schema |    table    | column | pcid | srid |  type
--------+-------------+--------+------+------+---------
 public | points      | pt     |    1 | 4326 | pcpoint
 public | patches     | pa     |    1 | 4326 | pcpatch
```

### Compressions

Compression is declared in the `<pc:metadata>` block of the schema document:

```xml
<pc:metadata>
  <Metadata name="compression">dimensional</Metadata>
</pc:metadata>
```

Three supported compression methods:

- **None**: stores points and patches as byte arrays using type and formats described in the schema document.
- **Dimensional**: stores patches as collections of dimensional data arrays with "appropriate" compression applied. Makes the most sense for smaller patch sizes.
- **LAZ** (LASZip): requires Pointcloud built with laz-perf support.

If no compression is declared, `none` is assumed.

Dimensional compression uses three schemes internally: run-length encoding (for low variability), common bits removal (for narrow bit range variability), and raw deflate compression via zlib.


## Functions: Schema

### PC_SchemaGetNDims

`PC_SchemaGetNDims(pcid integer) returns integer` -- Return the number of dimensions.

```sql
SELECT PC_SchemaGetNDims(1);
-- 18
```

### PC_SchemaIsValid

`PC_SchemaIsValid(xml text) returns boolean` -- Return true if the pointcloud schema is valid.

```sql
SELECT PC_SchemaIsValid(schema) FROM pointcloud_formats LIMIT 1;
-- t
```


## Functions: PcPoint

### PC_MakePoint

`PC_MakePoint(pcid integer, vals float8[]) returns pcpoint` -- Construct a new pcpoint from a pcid and array of doubles.

```sql
SELECT PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0]);

-- 010100000064CEFFFF94110000703000000400
```

Insert test values:

```sql
INSERT INTO points (pt)
SELECT PC_MakePoint(1, ARRAY[x,y,z,intensity])
FROM (
  SELECT
  -127+a/100.0 AS x,
    45+a/100.0 AS y,
         1.0*a AS z,
          a/10 AS intensity
  FROM generate_series(1,100) AS a
) AS values;
```

### PC_AsText (point)

`PC_AsText(p pcpoint) returns text` -- Returns a JSON version of the point data.

```sql
SELECT PC_AsText('010100000064CEFFFF94110000703000000400'::pcpoint);
-- {"pcid":1,"pt":[-127,45,124,4]}
```

### PC_Get

`PC_Get(pt pcpoint) returns float8[]` -- Returns values of all dimensions in an array.

```sql
SELECT PC_Get('010100000064CEFFFF94110000703000000400'::pcpoint);
-- {-127,45,124,4}
```

`PC_Get(pt pcpoint, dimname text) returns numeric` -- Returns the value of a named dimension.

```sql
SELECT PC_Get('010100000064CEFFFF94110000703000000400'::pcpoint, 'Intensity');
-- 4
```

### PC_MemSize (point)

`PC_MemSize(pt pcpoint) returns int4` -- Returns the memory size of a pcpoint.

```sql
SELECT PC_MemSize(PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0]));
-- 25
```

### PC_PCId (point)

`PC_PCId(p pcpoint) returns integer` -- Returns the pcid schema number of this point.

```sql
SELECT PC_PCId('010100000064CEFFFF94110000703000000400'::pcpoint);
-- 1
```


## Functions: PcPatch

### PC_Patch

`PC_Patch(pts pcpoint[]) returns pcpatch` -- Aggregate function that collects pcpoint values into a pcpatch.

```sql
INSERT INTO patches (pa)
SELECT PC_Patch(pt) FROM points GROUP BY id/10;
```

### PC_MakePatch

`PC_MakePatch(pcid integer, vals float8[]) returns pcpatch` -- Construct a new pcpatch from a pcid and array of doubles (array size must be a multiple of dimension count).

```sql
SELECT PC_AsText(PC_MakePatch(1, ARRAY[-126.99,45.01,1,0, -126.98,45.02,2,0, -126.97,45.03,3,0]));

-- {"pcid":1,"pts":[
--  [-126.99,45.01,1,0],[-126.98,45.02,2,0],[-126.97,45.03,3,0]
-- ]}
```

### PC_AsText (patch)

`PC_AsText(p pcpatch) returns text` -- Returns a JSON version of the patch data.

```sql
SELECT PC_AsText(pa) FROM patches LIMIT 1;

-- {"pcid":1,"pts":[
--  [-126.99,45.01,1,0],[-126.98,45.02,2,0],[-126.97,45.03,3,0],
--  [-126.96,45.04,4,0],[-126.95,45.05,5,0],[-126.94,45.06,6,0],
--  [-126.93,45.07,7,0],[-126.92,45.08,8,0],[-126.91,45.09,9,0]
-- ]}
```

### PC_Summary

`PC_Summary(p pcpatch) returns text` -- Returns a JSON formatted summary of the patch data.

```sql
SELECT PC_Summary(pa) FROM patches LIMIT 1;

-- {"pcid":1, "npts":9, "srid":4326, "compr":"dimensional",
--  "dims":[{"pos":0,"name":"X","size":4,"type":"int32_t","compr":"sigbits",
--           "stats":{"min":-126.99,"max":-126.91,"avg":-126.95}}, ...]}
```

### PC_NumPoints

`PC_NumPoints(p pcpatch) returns integer` -- Returns the number of points in a patch.

```sql
SELECT PC_NumPoints(pa) FROM patches LIMIT 1;
-- 9
```

### PC_PCId (patch)

`PC_PCId(p pcpatch) returns integer` -- Returns the pcid schema number of the patch.

### PC_MemSize (patch)

`PC_MemSize(p pcpatch) returns int4` -- Returns the memory size of a pcpatch.

### PC_Explode

`PC_Explode(p pcpatch) returns SetOf[pcpoint]` -- Set-returning function converting a patch to individual point records.

```sql
SELECT PC_AsText(PC_Explode(pa)), id
FROM patches WHERE id = 7;

              pc_astext               | id
--------------------------------------+----
 {"pcid":1,"pt":[-126.5,45.5,50,5]}   |  7
 {"pcid":1,"pt":[-126.49,45.51,51,5]} |  7
 {"pcid":1,"pt":[-126.48,45.52,52,5]} |  7
 ...
```

### PC_PointN

`PC_PointN(p pcpatch, n int4) returns pcpoint` -- Returns the n-th point (1-based). Negative n counts from the end.

### PC_Range

`PC_Range(p pcpatch, start int4, n int4) returns pcpatch` -- Returns a patch containing n points starting from the start-th point (1-based).

### PC_Union

`PC_Union(p pcpatch[]) returns pcpatch` -- Aggregate function that merges pcpatch entries into a single pcpatch.

```sql
SELECT PC_NumPoints(PC_Union(pa)) FROM patches;
-- 100
```

### PC_Intersects (patch-patch)

`PC_Intersects(p1 pcpatch, p2 pcpatch) returns boolean` -- Returns true if the bounds of p1 intersect the bounds of p2.

### PC_PatchAvg

`PC_PatchAvg(p pcpatch, dimname text) returns numeric` -- Returns the average value of a named dimension across all points.

```sql
SELECT PC_PatchAvg(pa, 'intensity') FROM patches WHERE id = 7;
-- 5.0000000000000000
```

`PC_PatchAvg(p pcpatch) returns pcpoint` -- Returns a PcPoint with average values of each dimension.

```sql
SELECT PC_AsText(PC_PatchAvg(pa)) FROM patches WHERE id = 7;
-- {"pcid":1,"pt":[-126.46,45.54,54.5,5]}
```

### PC_PatchMin

`PC_PatchMin(p pcpatch, dimname text) returns numeric` -- Returns the minimum value of a named dimension.

`PC_PatchMin(p pcpatch) returns pcpoint` -- Returns a PcPoint with the minimum values of each dimension.

### PC_PatchMax

`PC_PatchMax(p pcpatch, dimname text) returns numeric` -- Returns the maximum value of a named dimension.

`PC_PatchMax(p pcpatch) returns pcpoint` -- Returns a PcPoint with the maximum values of each dimension.

### PC_FilterGreaterThan

`PC_FilterGreaterThan(p pcpatch, dimname text, float8 value) returns pcpatch` -- Filter points with values greater than the given value.

```sql
SELECT PC_AsText(PC_FilterGreaterThan(pa, 'y', 45.57))
FROM patches WHERE id = 7;

-- {"pcid":1,"pts":[[-126.42,45.58,58,5],[-126.41,45.59,59,5]]}
```

### PC_FilterLessThan

`PC_FilterLessThan(p pcpatch, dimname text, float8 value) returns pcpatch` -- Filter points with values less than the given value.

### PC_FilterEquals

`PC_FilterEquals(p pcpatch, dimname text, float8 value) returns pcpatch` -- Filter points with values equal to the given value.

### PC_FilterBetween

`PC_FilterBetween(p pcpatch, dimname text, float8 value1, float8 value2) returns pcpatch` -- Filter points with values between (excluding) value1 and value2.

### PC_Sort

`PC_Sort(p pcpatch, dimnames text[]) returns pcpatch` -- Returns a copy of the patch lexicographically sorted along the given dimensions.

### PC_IsSorted

`PC_IsSorted(p pcpatch, dimnames text[], strict boolean default true) returns boolean` -- Checks whether a pcpatch is sorted lexicographically. The `strict` option checks for no duplicates.

### PC_SetPCId

`PC_SetPCId(p pcpatch, pcid int4, def float8 default 0.0) returns pcpatch` -- Sets the schema on a PcPatch. For dimensions in the new schema but not in the old, the value `def` is used (default `0.0`).

### PC_Transform

`PC_Transform(p pcpatch, pcid int4, def float8 default 0.0) returns pcpatch` -- Returns a new patch with data transformed based on the target schema. Unlike `PC_SetPCId`, this may change patch data if interpretations, scales or offsets differ.

### PC_Compress

`PC_Compress(p pcpatch, global_compression_scheme text, compression_config text) returns pcpatch` -- Compress a patch with a manually specified scheme.

Allowed global compression schemes: `auto`, `laz`, `dimensional`. For dimensional, the config is a comma-separated list of per-dimension compressions: `auto`, `zlib`, `sigbits`, `rle`.

### PC_Uncompress

`PC_Uncompress(p pcpatch) returns pcpatch` -- Returns an uncompressed version (compression type `none`). Must be the outer function in your query to return uncompressed data on the wire.


## Functions: WKB

### PC_AsBinary (point)

`PC_AsBinary(p pcpoint) returns bytea` -- Return the OGC "well-known binary" format for the point.

### PC_EnvelopeAsBinary

`PC_EnvelopeAsBinary(p pcpatch) returns bytea` -- Return the OGC WKB for the 2D bounds of the patch.

Note: `PC_Envelope` is a deprecated alias for `PC_EnvelopeAsBinary`.

### PC_BoundingDiagonalAsBinary

`PC_BoundingDiagonalAsBinary(p pcpatch) returns bytea` -- Return the OGC WKB for the bounding diagonal of the patch.


## Functions: PostGIS Integration

The `pointcloud_postgis` extension adds functions for using Pointcloud with PostGIS, converting `PcPoint` and `PcPatch` to Geometry and doing spatial filtering.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pointcloud;
CREATE EXTENSION pointcloud_postgis;
```

### Geometry Cast

`Geometry(pcpoint) returns geometry` / `pcpoint::geometry` -- Casts PcPoint to PostGIS geometry, mapping x/y/z/m.

```sql
SELECT ST_AsText(PC_MakePoint(1, ARRAY[-127, 45, 124.0, 4.0])::geometry);
-- POINT Z (-127 45 124)
```

### PC_EnvelopeGeometry

`PC_EnvelopeGeometry(pcpatch) returns geometry` -- Returns the 2D bounds as a PostGIS Polygon 2D.

```sql
SELECT ST_AsText(PC_EnvelopeGeometry(pa)) FROM patches LIMIT 1;
-- POLYGON((-126.99 45.01,-126.99 45.09,-126.91 45.09,-126.91 45.01,-126.99 45.01))
```

Useful for creating an index:

```sql
CREATE INDEX ON patches USING GIST(PC_EnvelopeGeometry(patch));
```

### PC_BoundingDiagonalGeometry

`PC_BoundingDiagonalGeometry(pcpatch) returns geometry` -- Returns the bounding diagonal as a LineString (2D/Z/M/ZM based on available dimensions).

```sql
SELECT ST_AsText(PC_BoundingDiagonalGeometry(pa)) FROM patches;
-- LINESTRING Z (-126.99 45.01 1,-126.91 45.09 9)
```

Useful for creating a ND index:

```sql
CREATE INDEX ON patches USING GIST(PC_BoundingDiagonalGeometry(patch) gist_geometry_ops_nd);
```

### PC_Intersection

`PC_Intersection(pcpatch, geometry) returns pcpatch` -- Returns a PcPatch containing only points that intersect the geometry.

```sql
SELECT PC_AsText(PC_Explode(PC_Intersection(
      pa,
      'SRID=4326;POLYGON((-126.451 45.552, -126.42 47.55, -126.40 45.552, -126.451 45.552))'::geometry
)))
FROM patches WHERE id = 7;

             pc_astext
--------------------------------------
 {"pcid":1,"pt":[-126.44,45.56,56,5]}
 {"pcid":1,"pt":[-126.43,45.57,57,5]}
 {"pcid":1,"pt":[-126.42,45.58,58,5]}
 {"pcid":1,"pt":[-126.41,45.59,59,5]}
```

### PC_Intersects (patch-geometry)

`PC_Intersects(p pcpatch, g geometry) returns boolean` / `PC_Intersects(g geometry, p pcpatch) returns boolean` -- Returns true if the bounds of the patch intersect the geometry.

```sql
SELECT PC_Intersects('SRID=4326;POINT(-126.451 45.552)'::geometry, pa)
FROM patches WHERE id = 7;
-- t
```


## Functions: Utils

### PC_Version / PC_Lib_Version / PC_Script_Version

```sql
SELECT PC_Version();           -- 1.2.5
SELECT PC_Lib_Version();       -- 1.2.5 2346cc2
SELECT PC_Script_Version();    -- 1.2.5
SELECT PC_Full_Version();
-- POINTCLOUD="1.2.5 2346cc2" PGSQL="170" LIBXML2="2.14.3 LAZPERF enabled=false
```

### PC_Lazperf_Enabled

`PC_Lazperf_Enabled() returns boolean` -- Returns true if LAZperf compression support is available.


## Loading Data with PDAL

[PDAL](https://pdal.io/) is used to load LIDAR files into PostgreSQL Pointcloud.
A PDAL pipeline is a JSON file declaring readers, filters, and writers.

Example pipeline to load a LAS file:

```json
{
  "pipeline":[
    {
      "type":"readers.las",
      "filename":"/home/lidar/st-helens-small.las",
      "spatialreference":"EPSG:26910"
    },
    {
      "type":"filters.chipper",
      "capacity":400
    },
    {
      "type":"writers.pgpointcloud",
      "connection":"host='localhost' dbname='pc' user='lidar' password='lidar' port='5432'",
      "table":"sthsm",
      "compression":"dimensional",
      "srid":"26910"
    }
  ]
}
```

Execute with:

```bash
pdal pipeline --input pipelinefile.json
```

The `filters.chipper` groups unordered points into compact patches for efficient storage.

### PDAL Writer Options

- **connection**: PostgreSQL connection string
- **table**: Table to write patches to
- **schema**: Schema to create the table in (optional)
- **column**: Column name for patches (default: `pa`)
- **compression**: Patch compression format (default: `dimensional`)
- **overwrite**: Replace any existing table (default: `true`)
- **srid**: Spatial reference id (default: `4326`)
- **pcid**: An existing PCID to use (optional)
- **pre_sql** / **post_sql**: SQL to execute before/after pipeline (optional)

### PDAL Reader Options

- **connection**: PostgreSQL connection string
- **table**: Table to read patches from
- **schema**: Schema to read from (optional)
- **column**: Column name to read (default: `pa`)
- **where**: SQL where clause to constrain the query (optional)
- **spatialreference**: Override the database SRID (optional)

Example pipeline to read and export:

```json
{
  "pipeline":[
    {
      "type":"readers.pgpointcloud",
      "connection":"host='localhost' dbname='pc' user='lidar' password='lidar' port='5432'",
      "table":"sthsm",
      "column":"pa",
      "spatialreference":"EPSG:26910",
      "where":"PC_Intersects(pa, ST_MakeEnvelope(560037.36, 5114846.45, 562667.31, 5118943.24, 26910))"
    },
    {
      "type":"writers.text",
      "filename":"/home/lidar/st-helens-small-out.txt"
    }
  ]
}
```
