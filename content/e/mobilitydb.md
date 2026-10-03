---
title: "mobilitydb"
linkTitle: "mobilitydb"
description: "MobilityDB geospatial trajectory data management & analysis platform"
weight: 1650
categories: ["GIS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**mobilitydb**](https://github.com/MobilityDB/MobilityDB) : MobilityDB geospatial trajectory data management & analysis platform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1650** | {{< badge content="mobilitydb" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< ext "mobilitydb" >}} | `1.3.1` | {{< category "GIS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|    **Need By**    | {{< ext "mobilitydb_datagen" >}} |
|   **See Also**    | {{< ext "h3" >}} {{< ext "pgrouting" >}} {{< ext "postgis" >}} {{< ext "pg_polyline" >}} {{< ext "q3c" >}} {{< ext "pg_sphere" >}} {{< ext "pointcloud" >}} {{< ext "pg_geohash" >}} {{< ext "qdgc" >}} {{< ext "pg_eviltransform" >}} |
|    **Siblings**   | {{< ext "mobilitydb_datagen" >}} |

> [!Note] Pigsty 1.3.1 includes the security fix; upgrading from 1.2 to 1.3 requires upstream backup/restore.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `mobilitydb` | `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "mobilitydb_18" "green" >}} {{< bg "17" "mobilitydb_17" "green" >}} {{< bg "16" "mobilitydb_16" "green" >}} {{< bg "15" "mobilitydb_15" "green" >}} {{< bg "14" "mobilitydb_14" "green" >}} | `mobilitydb_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "postgresql-18-mobilitydb" "green" >}} {{< bg "17" "postgresql-17-mobilitydb" "green" >}} {{< bg "16" "postgresql-16-mobilitydb" "green" >}} {{< bg "15" "postgresql-15-mobilitydb" "green" >}} {{< bg "14" "postgresql-14-mobilitydb" "green" >}} | `postgresql-$v-mobilitydb` | `postgresql-$v-postgis-3` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 2" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 2" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mobilitydb_18` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 789.4 KiB | [mobilitydb_18-1.3.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mobilitydb_18-1.3.1-1PGSTY.el8.x86_64.rpm) |
| `mobilitydb_18` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 737.8 KiB | [mobilitydb_18-1.3.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mobilitydb_18-1.3.1-1PGSTY.el8.aarch64.rpm) |
| `mobilitydb_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 690.3 KiB | [mobilitydb_18-1.3.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mobilitydb_18-1.3.1-1PGSTY.el9.x86_64.rpm) |
| `mobilitydb_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.5 KiB | [mobilitydb_18-1.3.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mobilitydb_18-1.3.1-1PGSTY.el9.aarch64.rpm) |
| `mobilitydb_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 707.8 KiB | [mobilitydb_18-1.3.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mobilitydb_18-1.3.1-1PGSTY.el10.x86_64.rpm) |
| `mobilitydb_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 681.8 KiB | [mobilitydb_18-1.3.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mobilitydb_18-1.3.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-mobilitydb` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 716.4 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.9 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.3 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 709.5 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 648.1 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.1 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.8 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 642.0 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 715.6 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.8 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.7 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 710.6 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 660.0 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.1 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 657.6 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 651.7 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 667.3 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 656.4 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 664.2 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.4 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.8 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 653.3 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 581.0 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.8 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.2 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 661.3 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.8 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.5 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.4 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mobilitydb` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 648.9 KiB | [postgresql-18-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.8 KiB | [postgresql-18-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.7 KiB | [postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.2 KiB | [postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-18-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mobilitydb_17` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 789.5 KiB | [mobilitydb_17-1.3.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mobilitydb_17-1.3.1-1PGSTY.el8.x86_64.rpm) |
| `mobilitydb_17` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 737.4 KiB | [mobilitydb_17-1.3.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mobilitydb_17-1.3.1-1PGSTY.el8.aarch64.rpm) |
| `mobilitydb_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 690.6 KiB | [mobilitydb_17-1.3.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mobilitydb_17-1.3.1-1PGSTY.el9.x86_64.rpm) |
| `mobilitydb_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.4 KiB | [mobilitydb_17-1.3.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mobilitydb_17-1.3.1-1PGSTY.el9.aarch64.rpm) |
| `mobilitydb_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 707.8 KiB | [mobilitydb_17-1.3.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mobilitydb_17-1.3.1-1PGSTY.el10.x86_64.rpm) |
| `mobilitydb_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 681.5 KiB | [mobilitydb_17-1.3.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mobilitydb_17-1.3.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-mobilitydb` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 713.4 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.6 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 716.0 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 709.8 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 648.0 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.6 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 641.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 715.9 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.6 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 714.9 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.4 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 657.9 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.3 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 651.3 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 667.0 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 574.0 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 663.5 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.8 KiB | [postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 664.3 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.3 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.5 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.9 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 653.4 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.4 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 581.1 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.0 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 661.4 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.6 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.8 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mobilitydb` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 649.0 KiB | [postgresql-17-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.9 KiB | [postgresql-17-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.6 KiB | [postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.3 KiB | [postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-17-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mobilitydb_16` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 789.3 KiB | [mobilitydb_16-1.3.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mobilitydb_16-1.3.1-1PGSTY.el8.x86_64.rpm) |
| `mobilitydb_16` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 737.2 KiB | [mobilitydb_16-1.3.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mobilitydb_16-1.3.1-1PGSTY.el8.aarch64.rpm) |
| `mobilitydb_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 690.3 KiB | [mobilitydb_16-1.3.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mobilitydb_16-1.3.1-1PGSTY.el9.x86_64.rpm) |
| `mobilitydb_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.3 KiB | [mobilitydb_16-1.3.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mobilitydb_16-1.3.1-1PGSTY.el9.aarch64.rpm) |
| `mobilitydb_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 708.0 KiB | [mobilitydb_16-1.3.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mobilitydb_16-1.3.1-1PGSTY.el10.x86_64.rpm) |
| `mobilitydb_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 681.4 KiB | [mobilitydb_16-1.3.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mobilitydb_16-1.3.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-mobilitydb` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 715.3 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.6 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.2 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.5 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 647.8 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.8 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.9 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 642.8 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 716.2 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.3 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 717.0 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.7 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 658.0 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.0 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.4 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 653.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 667.0 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 574.2 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 656.0 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.7 KiB | [postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 664.0 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 619.0 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.7 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.6 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 652.9 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.8 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.4 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.2 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 661.2 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.2 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.2 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mobilitydb` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 648.8 KiB | [postgresql-16-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.6 KiB | [postgresql-16-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.6 KiB | [postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.0 KiB | [postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-16-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mobilitydb_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 788.9 KiB | [mobilitydb_15-1.3.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mobilitydb_15-1.3.1-1PGSTY.el8.x86_64.rpm) |
| `mobilitydb_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 737.2 KiB | [mobilitydb_15-1.3.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mobilitydb_15-1.3.1-1PGSTY.el8.aarch64.rpm) |
| `mobilitydb_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 690.9 KiB | [mobilitydb_15-1.3.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mobilitydb_15-1.3.1-1PGSTY.el9.x86_64.rpm) |
| `mobilitydb_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 675.9 KiB | [mobilitydb_15-1.3.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mobilitydb_15-1.3.1-1PGSTY.el9.aarch64.rpm) |
| `mobilitydb_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 707.0 KiB | [mobilitydb_15-1.3.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mobilitydb_15-1.3.1-1PGSTY.el10.x86_64.rpm) |
| `mobilitydb_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 681.5 KiB | [mobilitydb_15-1.3.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mobilitydb_15-1.3.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-mobilitydb` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 715.5 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.3 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 715.7 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.7 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 647.2 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 647.9 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 643.2 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 715.7 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.4 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 715.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 708.9 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 658.5 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.0 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.3 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 653.4 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 666.7 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 573.5 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 656.0 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 536.0 KiB | [postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 663.9 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.5 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 662.3 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.7 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.4 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.6 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 661.3 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 621.6 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.0 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 612.7 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mobilitydb` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 648.9 KiB | [postgresql-15-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.2 KiB | [postgresql-15-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.3 KiB | [postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.4 KiB | [postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-15-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mobilitydb_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 788.9 KiB | [mobilitydb_14-1.3.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mobilitydb_14-1.3.1-1PGSTY.el8.x86_64.rpm) |
| `mobilitydb_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 737.3 KiB | [mobilitydb_14-1.3.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mobilitydb_14-1.3.1-1PGSTY.el8.aarch64.rpm) |
| `mobilitydb_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 690.4 KiB | [mobilitydb_14-1.3.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mobilitydb_14-1.3.1-1PGSTY.el9.x86_64.rpm) |
| `mobilitydb_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.3 KiB | [mobilitydb_14-1.3.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mobilitydb_14-1.3.1-1PGSTY.el9.aarch64.rpm) |
| `mobilitydb_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 706.4 KiB | [mobilitydb_14-1.3.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mobilitydb_14-1.3.1-1PGSTY.el10.x86_64.rpm) |
| `mobilitydb_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 681.8 KiB | [mobilitydb_14-1.3.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mobilitydb_14-1.3.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-mobilitydb` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 714.0 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 716.5 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 716.4 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 708.7 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 648.1 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.1 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 648.3 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 641.6 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg12+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 716.3 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.3 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 716.6 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 709.9 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 659.6 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.3 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 657.1 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 652.5 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg13+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 666.9 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 573.2 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 656.2 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 535.6 KiB | [postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 664.2 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.3 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 618.2 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 609.3 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 652.9 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.3 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 580.0 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 572.0 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 661.2 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.5 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 622.4 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 613.0 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mobilitydb` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 649.1 KiB | [postgresql-14-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.1-1PGSTY~resolute_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.3 KiB | [postgresql-14-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 580.5 KiB | [postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~rc1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-mobilitydb` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 572.2 KiB | [postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mobilitydb/postgresql-14-mobilitydb_1.3.0~alpha-3.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MobilityDB/MobilityDB" title="Repository" icon="github" subtitle="github.com/MobilityDB/MobilityDB" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mobilitydb-1.3.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg mobilitydb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install mobilitydb;		# install via package name, for the active PG version

pig install mobilitydb -v 18;   # install for PG 18
pig install mobilitydb -v 17;   # install for PG 17
pig install mobilitydb -v 16;   # install for PG 16
pig install mobilitydb -v 15;   # install for PG 15
pig install mobilitydb -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'postgis-3';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mobilitydb CASCADE; -- requires postgis
```

## Usage

Sources:

- [MobilityDB v1.3.1 README](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/README.md)
- [Extension control file](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb.in.control)
- [Version 1.3 migration manual](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/introduction.xml)
- [Temporal spatial API](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/temporal_spatial_p1.xml)
- [Version 1.3.1 release and upgrade](https://github.com/MobilityDB/MobilityDB/releases/tag/v1.3.1)
- [1.3.0 to 1.3.1 SQL migration](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb--1.3.0--1.3.1.sql)

`mobilitydb` 1.3.1 extends PostgreSQL and PostGIS with temporal values and moving-object trajectories. It supports storing changing attributes, reconstructing positions at a timestamp, and indexing space-time bounds. This patch fixes a backend-crashing binary-input vulnerability; installations on 1.3.0 should upgrade.

### Enable the Extension

The release requires PostgreSQL 14 or later and PostGIS 3 or later; it also adds PostgreSQL 19 build support. Package availability is tracked separately. Upstream requires loading the matching PostGIS library and recommends this lock allocation:

```conf
shared_preload_libraries = 'postgis-3'
max_locks_per_transaction = 128
```

Append the PostGIS library to the existing preload list, restart PostgreSQL, and enable both extensions in the target database with an authorized administrative role:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
```

### Store and Query a Trajectory

The example uses projected coordinates and complete UTC timestamps. Choose the coordinate reference system appropriate for the application; geographic coordinates require different distance semantics.

```sql
CREATE TABLE trips (
    trip_id bigint PRIMARY KEY,
    trip tgeompoint NOT NULL
);

INSERT INTO trips VALUES (
    1,
    tgeompoint 'SRID=3857;[Point(0 0)@2026-01-01 08:00:00+00,
                         Point(1000 0)@2026-01-01 09:00:00+00]'
);

SELECT valueAtTimestamp(trip, '2026-01-01 08:30:00+00'),
       ST_AsText(trajectory(trip)),
       length(trip),
       speed(trip)
FROM trips;

CREATE INDEX trips_space_time_idx ON trips USING gist (trip);

SELECT trip_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(-100, -100, 1100, 100, 3857),
    tstzspan '[2026-01-01 08:00:00+00, 2026-01-01 09:00:00+00]'
);
```

The bounding-box operator supplies an indexable filter. Apply the appropriate exact temporal or spatial predicate afterward when bounding overlap is insufficient.

### Type and Function Index

- `tbool`, `tint`, `tfloat`, and `ttext`: time-varying scalar values.
- `tgeompoint` and `tgeogpoint`: moving geometry or geography points; `tnpoint` represents a network point when that optional family is built.
- `tgeometry` and `tgeography`: arbitrary changing spatial values with discrete or step interpolation.
- `tcbuffer`, `tpose`, and `trgeometry`: optional experimental spatial families in the 1.3 line; do not assume every build contains them.
- Instant, sequence, and sequence-set representations describe one timestamp, one sequence, or multiple non-overlapping sequences. Linear interpolation is type-dependent.
- `valueAtTimestamp`, `startTimestamp`, `endTimestamp`, and `duration`: inspect temporal extent and values.
- `atTime` and `atGeometry`: restrict values to a time domain or geometry.
- `trajectory`, `length`, and `speed`: inspect the spatial path and motion.
- `twAvg` and `tUnion`: time-weighted summaries and temporal aggregation.
- GiST and SP-GiST operator classes accelerate supported temporal and space-time bounding queries.

### Upgrade and Safety Boundaries

Install the new library and SQL files, then update each database:

```sql
ALTER EXTENSION mobilitydb UPDATE TO '1.3.1';
SELECT extversion FROM pg_extension WHERE extname = 'mobilitydb';
```

- Version 1.3.1 fixes CVE-2026-102639: malformed WKB temporal, set, or span input could read beyond the input buffer and crash a backend. The SQL migration alone does not replace the vulnerable library; reconnect or restart processes that loaded the older binary.
- The migration removes five same-base-type `<->` operators and their `set_distance` functions because they conflict with operators supplied by `btree_gist`. Review dependent objects before updating and use the appropriate `btree_gist` operators where needed.
- Upgrading from the 1.2 line to 1.3 changes the temporal binary format and requires the upstream backup-and-restore procedure. An in-place 1.3.0-to-1.3.1 SQL update does not replace that major-line migration.
- Coordinate systems, interpolation, gaps, inclusive bounds, and units affect results. Validate them against the data model rather than treating every trajectory as a continuous geographical line.
