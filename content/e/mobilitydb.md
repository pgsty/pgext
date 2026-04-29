---
title: "mobilitydb"
linkTitle: "mobilitydb"
description: "MobilityDB geospatial trajectory data management & analysis platform"
weight: 1650
categories: ["GIS"]
width: full
---

[**mobilitydb**](https://github.com/MobilityDB/MobilityDB) : MobilityDB geospatial trajectory data management & analysis platform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1650** | {{< badge content="mobilitydb" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< ext "mobilitydb" >}} | `1.3.0` | {{< category "GIS" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|    **Need By**    | {{< ext "mobilitydb_datagen" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "h3_postgis" >}} {{< ext "timescaledb" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |
|    **Siblings**   | {{< ext "mobilitydb_datagen" >}} |

> [!Note] need another schema


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `mobilitydb` | `postgis` |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "postgresql-18-mobilitydb" "green" >}} {{< bg "17" "postgresql-17-mobilitydb" "green" >}} {{< bg "16" "postgresql-16-mobilitydb" "green" >}} {{< bg "15" "postgresql-15-mobilitydb" "green" >}} {{< bg "14" "postgresql-14-mobilitydb" "green" >}} | `postgresql-$v-mobilitydb` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.3 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 709.5 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.8 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 642.0 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.7 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 710.6 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 657.6 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 651.7 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.8 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.8 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.2 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.5 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.4 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.7 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.2 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 716.0 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 709.8 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 641.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 714.9 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.4 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 651.3 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 574.0 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.8 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.5 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 581.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.0 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.8 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.6 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.3 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.2 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.5 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.9 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 642.8 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 717.0 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.7 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.4 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 653.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 574.2 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.7 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.7 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.6 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.4 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.2 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.2 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.6 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.7 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.7 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 643.2 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 715.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 708.9 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.3 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 653.4 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 573.5 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 536.0 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.5 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.4 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.6 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.0 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 612.7 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.3 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.4 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 716.4 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.7 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.3 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 641.6 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.6 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.9 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 657.1 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 652.5 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 573.2 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.6 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.3 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.0 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.0 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.4 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.5 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.2 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MobilityDB/MobilityDB" title="Repository" icon="github" subtitle="github.com/MobilityDB/MobilityDB" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install mobilitydb;		# install via package name, for the active PG version

pig install mobilitydb -v 18;   # install for PG 18
pig install mobilitydb -v 17;   # install for PG 17
pig install mobilitydb -v 16;   # install for PG 16
pig install mobilitydb -v 15;   # install for PG 15
pig install mobilitydb -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mobilitydb CASCADE; -- requires postgis
```



## Usage

> [mobilitydb: Temporal and spatio-temporal data management for PostgreSQL](https://github.com/MobilityDB/MobilityDB)

MobilityDB extends PostgreSQL and PostGIS with temporal and spatio-temporal data types, enabling efficient storage, indexing, and querying of moving object data such as vehicle trajectories, sensor readings, and time-varying attributes.

**Key Documentation:**

- [MobilityDB Manual](https://docs.mobilitydb.com/MobilityDB/master/)
- [Temporal Types](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-types)
- [Temporal Operations](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-operations)
- [Spatial-Temporal Types](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#spatial-temporal-types)
- [Indexing](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#indexing)
- [MobilityDB Workshop](https://mobilitydb.com/documentation/)
- [API Reference](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html)

### Getting Started

MobilityDB requires PostGIS. Enable both extensions:

```sql
CREATE EXTENSION PostGIS;
CREATE EXTENSION MobilityDB;
```

### Temporal Types

MobilityDB provides temporal variants of base types:

| Temporal Type | Base Type | Description |
|---------------|-----------|-------------|
| `tbool`       | `boolean` | Time-varying boolean |
| `tint`        | `integer` | Time-varying integer |
| `tfloat`      | `float`   | Time-varying float |
| `ttext`       | `text`    | Time-varying text |
| `tgeompoint`  | `geometry(Point)` | Time-varying geometric point |
| `tgeogpoint`  | `geography(Point)` | Time-varying geographic point |

### Temporal Subtypes

Each temporal type can be represented in different subtypes depending on how values change over time:

| Subtype | Description | Example |
|---------|-------------|---------|
| **Instant** | Single value at a single timestamp | `'25.5@2025-01-01 08:00'` |
| **Sequence** | Continuous values over a time interval | `'[25.5@08:00, 28.1@09:00, 30.0@10:00]'` |
| **SequenceSet** | Set of non-overlapping sequences | `'{[25.5@08:00, 28.1@09:00], [30.0@11:00, 31.2@12:00]}'` |

Sequences use brackets to indicate inclusive `[` or exclusive `(` bounds, just like PostgreSQL range types.

### Creating Temporal Values

**Instant values:**

```sql
SELECT tfloat '25.5@2025-06-01 08:00:00+00';
SELECT tgeompoint 'SRID=4326;Point(2.3522 48.8566)@2025-06-01 08:00:00+00';
```

**Sequence values (continuous interpolation):**

```sql
SELECT tfloat '[20.0@2025-06-01 08:00, 25.5@2025-06-01 09:00, 22.0@2025-06-01 10:00]';
```

**Discrete sequences (stepwise interpolation):**

```sql
SELECT tint 'Interp=Step;[10@2025-06-01 08:00, 20@2025-06-01 09:00, 15@2025-06-01 10:00]';
```

**SequenceSet values:**

```sql
SELECT tfloat '{[20.0@08:00, 25.5@09:00], [22.0@11:00, 28.0@12:00]}';
```

**Constructing from components:**

```sql
SELECT tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00');
SELECT tgeompoint_seq(ARRAY[
    tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
    tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:30+00'),
    tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 09:00+00')
]);
```

### Temporal Operations

**Extracting values at a specific time:**

```sql
SELECT valueAtTimestamp(temp, '2025-06-01 08:30:00+00')
FROM (SELECT tfloat '[20.0@08:00, 30.0@09:00]' AS temp) t;
-- Returns 25.0 (linear interpolation)
```

**Restricting to a time period:**

```sql
SELECT atTime(trip, tstzspan '[2025-06-01 08:00, 2025-06-01 09:00]')
FROM trips;
```

**Getting the time span of a temporal value:**

```sql
SELECT duration(trip), startTimestamp(trip), endTimestamp(trip)
FROM trips;
```

**Temporal comparisons:**

```sql
-- Time periods when temperature exceeded 30 degrees
SELECT atValue(temperature, true)
FROM (SELECT tfloat '[20@08:00, 35@09:00, 25@10:00]' #> 30.0 AS temperature) t;
```

### Spatial-Temporal Operations

**Trajectory: extract the spatial path as a geometry:**

```sql
SELECT ST_AsText(trajectory(trip))
FROM trips
WHERE vehicle_id = 42;
```

**Speed calculation:**

```sql
-- Speed in units per second (m/s for geographic points)
SELECT speed(trip)
FROM trips
WHERE vehicle_id = 42;
```

**Length of trajectory:**

```sql
SELECT length(trip)
FROM trips
WHERE vehicle_id = 42;
```

**Space-time bounding box (stbox):**

```sql
-- Get the space-time bounding box
SELECT stbox(trip)
FROM trips;

-- Construct an stbox for querying
SELECT stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

**Spatial restriction: values within an area:**

```sql
-- Portions of a trip within a polygon
SELECT atGeometry(trip, ST_Buffer(ST_Point(2.35, 48.86, 4326), 0.01))
FROM trips;
```

**Distance between two temporal points:**

```sql
SELECT distance(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

**Nearest approach distance and time:**

```sql
SELECT nearestApproachDistance(t1.trip, t2.trip),
       nearestApproachInstant(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

### Indexing

MobilityDB supports GiST and SP-GiST indexes for efficient temporal and spatio-temporal queries.

**SP-GiST index for temporal types (time dimension):**

```sql
CREATE INDEX ON measurements USING spgist(temperature);
```

**GiST index for spatio-temporal types (space + time):**

```sql
CREATE INDEX ON trips USING gist(trip);
```

These indexes accelerate bounding box queries, temporal overlap checks, and spatial-temporal intersection:

```sql
-- Uses GiST index for space-time filtering
SELECT vehicle_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

### Example: Vehicle Tracking

A complete example storing and querying vehicle GPS trajectories:

```sql
CREATE TABLE vehicles (
    vehicle_id  INT PRIMARY KEY,
    plate       TEXT,
    type        TEXT
);

CREATE TABLE trips (
    trip_id     BIGSERIAL PRIMARY KEY,
    vehicle_id  INT REFERENCES vehicles(vehicle_id),
    trip        tgeompoint,
    trip_date   DATE
);

CREATE INDEX ON trips USING gist(trip);

-- Insert a trip as a sequence of GPS points
INSERT INTO trips (vehicle_id, trip, trip_date) VALUES (
    1,
    tgeompoint_seq(ARRAY[
        tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
        tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:15+00'),
        tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 08:30+00'),
        tgeompoint_inst(ST_Point(2.3488, 48.8534, 4326), '2025-06-01 08:45+00')
    ]),
    '2025-06-01'
);

-- Where was vehicle 1 at 08:20?
SELECT valueAtTimestamp(trip, '2025-06-01 08:20+00')
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- What was the average speed?
SELECT twAvg(speed(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Total distance traveled
SELECT length(trip)
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Get the full trajectory as a LineString
SELECT ST_AsGeoJSON(trajectory(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';
```

### Example: Spatio-Temporal Intersection Query

Find all trips that passed through a specific area during a given time window:

```sql
-- Define area of interest: a circle around the Eiffel Tower
WITH area AS (
    SELECT ST_Buffer(ST_Point(2.2945, 48.8584, 4326)::geography, 500)::geometry AS geom
)
SELECT t.vehicle_id,
       t.trip_date,
       atGeometry(t.trip, a.geom) AS trip_in_area,
       length(atGeometry(t.trip, a.geom)) AS distance_in_area
FROM trips t, area a
WHERE t.trip && stbox(
    a.geom,
    tstzspan '[2025-06-01 07:00+00, 2025-06-01 10:00+00]'
)
  AND eIntersects(t.trip, a.geom)
ORDER BY t.trip_date;
```

### Aggregate Functions

MobilityDB provides temporal aggregates:

```sql
-- Time-weighted average of a temporal float
SELECT twAvg(temperature) FROM sensor_data WHERE sensor_id = 1;

-- Merge multiple temporal points into one
SELECT tUnion(trip) FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Centroid of a set of temporal points at each timestamp
SELECT tCentroid(trip) FROM trips WHERE trip_date = '2025-06-01';
```
