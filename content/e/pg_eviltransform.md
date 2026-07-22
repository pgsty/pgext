---
title: "pg_eviltransform"
linkTitle: "pg_eviltransform"
description: "Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform"
weight: 1580
categories: ["GIS"]
width: full
---

[**pg_eviltransform**](https://github.com/aiyou178/pg_eviltransform) : Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1580** | {{< badge content="pg_eviltransform" link="https://github.com/aiyou178/pg_eviltransform" >}} | {{< ext "pg_eviltransform" >}} | `0.0.4` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `eviltransform_internal` |
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "earthdistance" >}} {{< ext "tzf" >}} {{< ext "geoip" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_eviltransform` | `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "pg_eviltransform_18" "green" >}} {{< bg "17" "pg_eviltransform_17" "green" >}} {{< bg "16" "pg_eviltransform_16" "green" >}} {{< bg "15" "pg_eviltransform_15" "green" >}} {{< bg "14" "pg_eviltransform_14" "green" >}} | `pg_eviltransform_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "postgresql-18-eviltransform" "green" >}} {{< bg "17" "postgresql-17-eviltransform" "green" >}} {{< bg "16" "postgresql-16-eviltransform" "green" >}} {{< bg "15" "postgresql-15-eviltransform" "green" >}} {{< bg "14" "postgresql-14-eviltransform" "green" >}} | `postgresql-$v-eviltransform` | `postgresql-$v-postgis` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_18` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 918.9 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_18-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 819.2 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_18-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 926.7 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_18-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 871.2 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_18-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 926.8 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_18-0.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 849.0 KiB | [pg_eviltransform_18-0.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_18-0.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-eviltransform` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 737.3 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 611.9 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 737.2 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 612.7 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 816.9 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 723.7 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 810.9 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 715.0 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 806.7 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 713.3 KiB | [postgresql-18-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_17` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 915.6 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_17-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 816.2 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_17-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 923.9 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_17-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 866.7 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_17-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 923.3 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_17-0.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 848.4 KiB | [pg_eviltransform_17-0.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_17-0.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-eviltransform` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 735.1 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 609.8 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 734.9 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 610.7 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 816.1 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 721.5 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 808.9 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 712.9 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 804.1 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 711.4 KiB | [postgresql-17-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_16` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 914.0 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_16-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 814.8 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_16-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 921.5 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_16-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 865.5 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_16-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 921.3 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_16-0.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 848.6 KiB | [pg_eviltransform_16-0.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_16-0.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-eviltransform` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 733.5 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 610.0 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 733.8 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 610.0 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 814.0 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 720.0 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 807.5 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 712.6 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 803.0 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 709.9 KiB | [postgresql-16-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_15` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 904.9 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_15-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 805.9 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_15-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 911.3 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_15-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 855.5 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_15-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 912.0 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_15-0.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 844.5 KiB | [pg_eviltransform_15-0.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_15-0.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-eviltransform` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 728.0 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 605.4 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 727.9 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 605.7 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 805.6 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 714.0 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 800.4 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 706.0 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 796.1 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 704.0 KiB | [postgresql-15-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_14` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 901.1 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_14-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 803.3 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_14-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 908.9 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_14-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 852.8 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_14-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 908.4 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_14-0.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 843.1 KiB | [pg_eviltransform_14-0.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_14-0.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-eviltransform` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 725.3 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 603.7 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 725.1 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 604.2 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 804.6 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 711.2 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 798.0 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 704.6 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 793.9 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 703.3 KiB | [postgresql-14-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aiyou178/pg_eviltransform" title="Repository" icon="github" subtitle="github.com/aiyou178/pg_eviltransform" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_eviltransform-0.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_eviltransform;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_eviltransform;		# install via package name, for the active PG version

pig install pg_eviltransform -v 18;   # install for PG 18
pig install pg_eviltransform -v 17;   # install for PG 17
pig install pg_eviltransform -v 16;   # install for PG 16
pig install pg_eviltransform -v 15;   # install for PG 15
pig install pg_eviltransform -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_eviltransform CASCADE; -- requires postgis
```

## Usage

Sources:

- [Official v0.0.4 README](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/README.md)
- [v0.0.4 release notes](https://github.com/aiyou178/pg_eviltransform/releases/tag/v0.0.4)
- [v0.0.4 control file](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/pg_eviltransform.control)
- [v0.0.4 upgrade SQL](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/pg_eviltransform--0.0.3--0.0.4.sql)

`pg_eviltransform` extends PostGIS with coordinate transformations involving China's GCJ-02 and BD-09 systems. Version `0.0.4` also adds exact Jenks natural-break classification through `ST_JenksBins` array and aggregate overloads.

### Coordinate Transformation

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_eviltransform;

-- WGS84 to GCJ-02 using a readable coordinate-system name.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326),
    'GCJ02'
);

-- BD-09 to Web Mercator.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002),
    3857
);
```

Custom SRIDs are `990001` for GCJ-02 and `990002` for BD-09. When neither endpoint uses a custom system, `ST_EvilTransform` delegates to PostGIS `ST_Transform`; otherwise it converts through WGS84 (`4326`) when necessary.

### Jenks Natural Breaks

```sql
-- Array form; NULL elements are ignored.
SELECT ST_JenksBins(ARRAY[1, 2, NULL, 10, 11]::numeric[], 2);

-- Streaming aggregate form for a large table.
SELECT ST_JenksBins(value, 7)
FROM measurements;

-- Return lower rather than upper bin edges.
SELECT ST_JenksBins(value, 7, true)
FROM measurements;
```

Array inputs support `numeric`, `double precision`, `real`, `bigint`, `integer`, and `smallint`. Aggregate inputs are `numeric` or `double precision`; cast other numeric columns when needed.

### API Index and Caveats

- `ST_EvilTransform(geometry, integer|text)` and `ST_EvilTransform(geometry, text, integer|text)`: four overloads corresponding to the PostGIS `ST_Transform` interface.
- `ST_JenksBins(values[], breaks [, invert])`: classifies an array and returns `double precision[]` edges.
- `ST_JenksBins(value, breaks [, invert])`: streaming aggregate that avoids materializing `array_agg`.
- PostGIS is a runtime prerequisite and must be installed before `pg_eviltransform`.
- Jenks inputs must be finite and `breaks` must be at least one. `numeric` values are converted to finite `f64`, so returned edges are floating-point values.
- When the distinct value count does not exceed `breaks`, the result is the sorted set of unique values; no valid input rows return `NULL`.
