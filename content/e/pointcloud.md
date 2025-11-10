---
title: "pointcloud"
linkTitle: "pointcloud"
description: "data type for lidar point clouds"
weight: 1520
categories: ["GIS"]
width: full
---

[**pointcloud**](https://github.com/pgpointcloud/pointcloud)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1520** | {{< badge content="pointcloud" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< ext "pointcloud" >}} | `1.2.5` | {{< category "GIS" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pointcloud_postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} |
|    **Siblings**   | {{< ext "pointcloud_postgis" >}} |


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


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pointcloud_13` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.4 KiB | [pointcloud_13-1.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pointcloud_13-1.2.5-1PGDG.rhel8.x86_64.rpm) |
| `pointcloud_13` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.8 KiB | [pointcloud_13-1.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pointcloud_13-1.2.5-1PGDG.rhel8.aarch64.rpm) |
| `pointcloud_13` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.6 KiB | [pointcloud_13-1.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pointcloud_13-1.2.5-1PGDG.rhel9.x86_64.rpm) |
| `pointcloud_13` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.1 KiB | [pointcloud_13-1.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pointcloud_13-1.2.5-1PGDG.rhel9.aarch64.rpm) |
| `pointcloud_13` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.6 KiB | [pointcloud_13-1.2.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pointcloud_13-1.2.5-3PGDG.rhel10.x86_64.rpm) |
| `pointcloud_13` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.9 KiB | [pointcloud_13-1.2.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pointcloud_13-1.2.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pointcloud` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.1 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg12+1_amd64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.6 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg12+1_arm64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.4 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg13+1_amd64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.8 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg13+1_arm64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.3 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pointcloud` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.4 KiB | [postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpointcloud/postgresql-13-pointcloud_1.2.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpointcloud/pointcloud" title="Repository" icon="github" subtitle="github.com/pgpointcloud/pointcloud" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pointcloud-1.2.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pointcloud; # get pointcloud source code
pig build dep pointcloud; # install build dependencies
pig build pkg pointcloud; # build extension rpm or deb
pig build ext pointcloud; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pointcloud; # install by extension name, for the current active PG version
pig ext install pointcloud; # install via package alias, for the active PG version
pig ext install pointcloud -v 18;   # install for PG 18
pig ext install pointcloud -v 17;   # install for PG 17
pig ext install pointcloud -v 16;   # install for PG 16
pig ext install pointcloud -v 15;   # install for PG 15
pig ext install pointcloud -v 14;   # install for PG 14
pig ext install pointcloud -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pointcloud;
```

