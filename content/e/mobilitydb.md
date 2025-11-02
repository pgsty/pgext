---
title: "mobilitydb"
linkTitle: "mobilitydb"
description: "MobilityDB geospatial trajectory data management & analysis platform"
weight: 1650
categories: ["GIS"]
width: full
---

MobilityDB geospatial trajectory data management & analysis platform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1650** | {{< badge content="mobilitydb" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< ext "mobilitydb" >}} | `1.3.0` | {{< category "GIS" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "h3_postgis" >}} {{< ext "timescaledb" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |

> [!Note] need another schema


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **Debian** | {{< badge content="PGDG" link="/e/mobilitydb" >}} | `1.3.0` | {{< bg "18" "postgresql-18-mobilitydb" "green" >}} {{< bg "17" "postgresql-17-mobilitydb" "green" >}} {{< bg "16" "postgresql-16-mobilitydb" "green" >}} {{< bg "15" "postgresql-15-mobilitydb" "green" >}} {{< bg "14" "postgresql-14-mobilitydb" "green" >}} {{< bg "13" "postgresql-13-mobilitydb" "green" >}} | `postgresql-$v-mobilitydb` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 709.5 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 642.0 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 710.6 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 651.7 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.8 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 572.2 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 709.8 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 641.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 709.4 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 651.3 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | 1.2.0 | `u22.x86_64` | pgdg | 574.0 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | 1.2.0 | `u22.aarch64` | pgdg | 535.8 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 572.0 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 708.5 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 642.8 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 709.7 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 653.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | 1.2.0 | `u22.x86_64` | pgdg | 574.2 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | 1.2.0 | `u22.aarch64` | pgdg | 535.7 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.6 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 572.2 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 708.7 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 643.2 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 708.9 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 653.4 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | 1.2.0 | `u22.x86_64` | pgdg | 573.5 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | 1.2.0 | `u22.aarch64` | pgdg | 536.0 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.5 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 572.6 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 708.7 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 641.6 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 709.9 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 652.5 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | 1.2.0 | `u22.x86_64` | pgdg | 573.2 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | 1.2.0 | `u22.aarch64` | pgdg | 535.6 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.3 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 572.0 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-mobilitydb` | 1.3.0 | `d12.x86_64` | pgdg | 709.6 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-13-mobilitydb` | 1.3.0 | `d12.aarch64` | pgdg | 640.8 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-13-mobilitydb` | 1.3.0 | `d13.x86_64` | pgdg | 709.4 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-13-mobilitydb` | 1.3.0 | `d13.aarch64` | pgdg | 652.2 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-13-mobilitydb` | 1.2.0 | `u22.x86_64` | pgdg | 574.3 KiB | [postgresql-13-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-mobilitydb` | 1.2.0 | `u22.aarch64` | pgdg | 535.0 KiB | [postgresql-13-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-mobilitydb` | 1.3.0 | `u24.x86_64` | pgdg | 609.7 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-mobilitydb` | 1.3.0 | `u24.aarch64` | pgdg | 571.6 KiB | [postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-13-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MobilityDB/MobilityDB" title="Repository" icon="github" subtitle="github.com/MobilityDB/MobilityDB" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install mobilitydb; # install by extension name, for the current active PG version
pig ext install mobilitydb; # install via package alias, for the active PG version
pig ext install mobilitydb -v 18;   # install for PG 18
pig ext install mobilitydb -v 17;   # install for PG 17
pig ext install mobilitydb -v 16;   # install for PG 16
pig ext install mobilitydb -v 15;   # install for PG 15
pig ext install mobilitydb -v 14;   # install for PG 14
pig ext install mobilitydb -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION mobilitydb;
```

